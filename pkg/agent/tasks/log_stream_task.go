package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/consts"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	_log "github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"google.golang.org/grpc"
	"io"
	corev1 "k8s.io/api/core/v1"
	"sync"
	"time"
)

// var logsMutex sync.Mutex
var logTaskMap = make(map[string]*logTask)
var logTaskMapMutex = sync.Mutex{}

const (
	PodLogBoolTrue     = "y"
	PodLogBoolFalse    = "n"
	TailLinesThreshold = int64(2500)
)

type logTask struct {
	cancel    context.CancelFunc
	heartbeat *time.Timer
}

func initLogStreamTask(logsCancel context.CancelFunc, podLogsTaskId string) {
	newLogTask := logTask{logsCancel, nil}
	logTaskMapMutex.Lock()
	logTaskMap[podLogsTaskId] = &newLogTask
	logTaskMapMutex.Unlock()
}

func LogStreamTask(podLogsTask *pb.PodLogsStream, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) error {
	if podLogsTask.Payload == consts.LogsTaskCancel {
		cancelTask(podLogsTask.Id, stream)
	} else if podLogsTask.Payload == consts.LogsTaskHeartbeat {
		keepAliveLogs(podLogsTask.Id, stream)
	} else {
		logsCtx, logsCancel := context.WithCancel(stream.Context())
		initLogStreamTask(logsCancel, podLogsTask.Id)
		keepAliveLogs(podLogsTask.Id, stream)
		go func(taskID, taskPayload string, task *pb.PodLogsStream, ctx context.Context) {
			var input k8s.GetPodLogsInputParams
			err := json.Unmarshal([]byte(task.Input), &input)
			if err != nil {
				_log.Logger.Errorw("Failed to unmarshal pod logs input", "err", err)
			}
			logResult := &pb.LogsResult{TaskId: taskID, Cancel: false}
			podLogOptions := corev1.PodLogOptions{Follow: true}
			if input.Container != "" {
				podLogOptions.Container = input.Container
			}
			zero := int64(0)
			if input.TailLines != nil && *input.TailLines != zero {
				if *input.TailLines > TailLinesThreshold {
					threshold := TailLinesThreshold
					podLogOptions.TailLines = &threshold
				} else {
					podLogOptions.TailLines = input.TailLines
				}
			}
			if input.SinceSeconds != nil && *input.SinceSeconds > zero {
				podLogOptions.SinceSeconds = input.SinceSeconds
			}
			// checking timestamps
			switch input.Timestamps {
			case PodLogBoolTrue:
				podLogOptions.Timestamps = true
			case PodLogBoolFalse:
				podLogOptions.Timestamps = false
			}

			req := k8s.GetKubeClientSet().CoreV1().Pods(input.NamespaceName).GetLogs(input.Pod, &podLogOptions)
			podLogs, err := req.Stream(ctx)
			if err != nil {
				_log.Logger.Errorw("failed to get pod logs of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
				cancelTask(taskID, stream)
			}
			defer func() {
				if podLogs != nil {
					_ = podLogs.Close()
				}
			}()

			buf := make([]byte, 2048)
			for {
				select {
				case <-ctx.Done():
					_log.Logger.Infow("log streaming task cancelled", "logs-stream", "cancelled")
					if podLogs != nil {
						_ = podLogs.Close()
					}
					return // Exit the goroutine
				default:
					numBytes, err := podLogs.Read(buf)
					if err != nil {
						logTaskMapMutex.Lock()
						_, ok := logTaskMap[taskID]
						logTaskMapMutex.Unlock()
						if !ok {
							return
						}
						if err == io.EOF {
							_log.Logger.Infow("End of pod logs stream", "taskID", taskID)
						} else {
							_log.Logger.Errorw("failed to read logs of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
						}
						cancelTask(taskID, stream)
						return // Exit the goroutine
					}

					logResult.Output = buf[:numBytes]
					if err := stream.Send(&pb.TaskStreamRequest{
						Payload: &pb.TaskStreamRequest_LogsResult{
							LogsResult: logResult,
						},
					}); err != nil {
						_log.Logger.Errorw("failed to send log message of namespace: "+input.NamespaceName+", pod: "+input.Pod, "pod-logs-err", err)
						cancelTask(taskID, stream)
					}
					time.Sleep(500 * time.Millisecond)
				}
			}

		}(podLogsTask.Id, podLogsTask.Payload, podLogsTask, logsCtx)
	}
	return nil
}

func keepAliveLogs(taskID string, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) {
	_log.Logger.Infow(fmt.Sprintf("logs streaming keep alive func task with ID: %s", taskID), "keep-alive", taskID)
	logsTimeout := 10 * time.Second

	logTaskMapMutex.Lock()
	defer logTaskMapMutex.Unlock()

	task, ok := logTaskMap[taskID]
	if !ok {
		_log.Logger.Warnw(fmt.Sprintf("No active task found with ID: %s", taskID), "keep-alive", taskID)
		if err := stream.Send(&pb.TaskStreamRequest{
			Payload: &pb.TaskStreamRequest_LogsResult{
				LogsResult: &pb.LogsResult{TaskId: taskID, Output: nil, Cancel: true},
			},
		}); err != nil {
			_log.Logger.Errorw("failed to send log stream closed message [No active task found]", "logs-timeout", taskID)
		}
		return
	}

	resetHeartbeat(task, taskID, logsTimeout, stream)
}

func resetHeartbeat(task *logTask, taskID string, logsTimeout time.Duration, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) {
	if task.heartbeat != nil {
		task.heartbeat.Stop()
	}

	task.heartbeat = time.AfterFunc(logsTimeout, func() {
		_log.Logger.Infow(fmt.Sprintf("logs timeout for task: %s", taskID), "logs-timeout", taskID)
		cancelTask(taskID, stream)
		_log.Logger.Infow(fmt.Sprintf("logs timeout for task: %s", taskID), "logs-timeout", taskID)
		task.cancel()
		logTaskMapMutex.Lock()
		delete(logTaskMap, taskID)
		logTaskMapMutex.Unlock()
		_log.Logger.Infow(fmt.Sprintf("Cancelled task with ID: %s", taskID), "logs-timeout", taskID)
		if err := stream.Send(&pb.TaskStreamRequest{
			Payload: &pb.TaskStreamRequest_LogsResult{
				LogsResult: &pb.LogsResult{TaskId: taskID, Output: nil, Cancel: true},
			},
		}); err != nil {
			_log.Logger.Errorw("failed to send log stream closed message", "log-cancel", taskID)
		}
	})
}

func cancelTask(taskID string, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) {
	_log.Logger.Infow(fmt.Sprintf("Cancelling func task with ID: %s", taskID), "log-cancel", taskID)
	logTaskMapMutex.Lock()
	defer logTaskMapMutex.Unlock()
	if task, ok := logTaskMap[taskID]; ok {
		task.cancel()
		if task.heartbeat != nil {
			task.heartbeat.Stop()
		}
		delete(logTaskMap, taskID)
		_log.Logger.Infow("Cancelled task", "log-cancel", taskID)
		if err := stream.Send(&pb.TaskStreamRequest{
			Payload: &pb.TaskStreamRequest_LogsResult{
				LogsResult: &pb.LogsResult{TaskId: taskID, Output: nil, Cancel: true},
			},
		}); err != nil {
			_log.Logger.Errorw("failed to send log stream closed message", "log-cancel", taskID)
		}
	} else {
		_log.Logger.Warnw("No active task found", "log-cancel", taskID)
		if err := stream.Send(&pb.TaskStreamRequest{
			Payload: &pb.TaskStreamRequest_LogsResult{
				LogsResult: &pb.LogsResult{TaskId: taskID, Output: nil, Cancel: true},
			},
		}); err != nil {
			_log.Logger.Errorw("failed to send log stream closed message [No active task found]", "log-cancel", taskID)
		}
	}
}
