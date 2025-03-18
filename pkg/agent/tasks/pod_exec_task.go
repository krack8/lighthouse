package tasks

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/krack8/lighthouse/pkg/common/consts"
	"github.com/krack8/lighthouse/pkg/common/k8s"
	_log "github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"google.golang.org/grpc"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"sync"
	"time"
)

type PodExecStream struct {
	stdinWriter     *io.PipeWriter
	stdinReader     *io.PipeReader
	stdoutReader    *io.PipeReader
	stdoutWriter    *io.PipeWriter
	stderrReader    *io.PipeReader
	stderrWriter    *io.PipeWriter
	terminalQueue   *TerminalSizeQueue
	execStream      remotecommand.Executor
	execStreamCtx   context.Context
	closeExecStream context.CancelFunc
	wg              sync.WaitGroup
	grpcStream      grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]
}

type TerminalSizeQueue struct {
	sizeChan chan *remotecommand.TerminalSize
}

var taskPodExecStreams = make(map[string]*PodExecStream)
var mutex sync.Mutex
var heartbeatTimers = make(map[string]*time.Timer)

func PodExecTask(taskID string, payloadType string, input string, command []byte, grpcStream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) error {
	inputPayload := k8s.PodExecInputParams{}

	err := json.Unmarshal([]byte(input), &inputPayload)
	if err != nil {
		_log.Logger.Errorw(fmt.Sprintf("Closing Stream! Input payload error: %s", err.Error()), "TaskID", taskID, "Payload", payloadType)
		return err
	}

	if payloadType == consts.TaskPodExecInitConn {
		_log.Logger.Infow(fmt.Sprintf("Recieved initiate connection request from controller"), "TaskID", taskID, "Payload", payloadType)
		streamExec, err := initExecStream(taskID, inputPayload.NamespaceName, inputPayload.PodName, inputPayload.ContainerName, k8s.GetKubeRestConfig(), grpcStream)
		if err != nil {
			_log.Logger.Errorw(fmt.Sprintf("Closing Stream! Failed to initiate exec stream: %s", err.Error()), "TaskID", taskID, "Payload", payloadType)
			return nil
		}

		streamExec.initHeartbeat(taskID)
		streamExec.wg.Add(1)
		go streamExec.startPodExecStreamReader(taskID)

	} else {
		streamExec, err := getExecStream(taskID)
		if err != nil {
			_log.Logger.Errorw(fmt.Sprintf("Closing Stream! Failed to get exec stream: %s", err.Error()), "TaskID", taskID, "Payload", payloadType, "Command", string(command))
			return nil
		}

		if payloadType == consts.TaskPodExecCloseConn {
			_log.Logger.Infow(fmt.Sprintf("Recieved close connection request from controller"), "TaskID", taskID, "Payload", payloadType)
			streamExec.closeExecStream()

		} else if payloadType == consts.TaskPodExecHeartbeat {
			streamExec.checkHeartbeat(taskID)

		} else if payloadType == consts.TaskPodExecCommand {
			// Writing commands to pod exec stream
			go streamExec.writeCommandToPodExecStream(taskID, command)
		}
	}
	return nil
}

func initExecStream(taskID, namespace, podName, containerName string, config *rest.Config, grpcStream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) (*PodExecStream, error) {
	mutex.Lock()
	defer mutex.Unlock()

	stdinReader, stdinWriter := io.Pipe()
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	shellCmd := "/bin/bash"
	if !isShellAvailable(config, namespace, podName, containerName, "/bin/bash") {
		shellCmd = "/bin/sh"
	}

	reqURL := k8s.GetKubeClientSet().CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(namespace).
		Name(podName).
		SubResource("exec").
		Param("container", containerName).
		Param("stdin", "true").
		Param("stdout", "true").
		Param("stderr", "true").
		Param("command", shellCmd).
		Param("tty", "true").
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   []string{}, // Start shell session
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec).URL()

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", reqURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create SPDY executor: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	stream := &PodExecStream{
		stdinWriter:     stdinWriter,
		stdinReader:     stdinReader,
		stdoutReader:    stdoutReader,
		execStreamCtx:   ctx,
		closeExecStream: cancel,
		stdoutWriter:    stdoutWriter,
		stderrReader:    stderrReader,
		stderrWriter:    stderrWriter,
		execStream:      exec,
		terminalQueue:   &TerminalSizeQueue{sizeChan: make(chan *remotecommand.TerminalSize, 1)},
		grpcStream:      grpcStream,
	}

	taskPodExecStreams[taskID] = stream

	// Start exec stream
	go stream.startExecStream(taskID)

	return stream, nil
}

func getExecStream(taskID string) (*PodExecStream, error) {
	if stream, exists := taskPodExecStreams[taskID]; exists {
		return stream, nil
	}
	return nil, errors.New("exec stream not found")
}

func isShellAvailable(config *rest.Config, namespace, pod, container, shell string) bool {
	reqURL := k8s.GetKubeClientSet().CoreV1().RESTClient().Post().
		Resource("pods").
		Namespace(namespace).
		Name(pod).
		SubResource("exec").
		Param("container", container).
		Param("stdin", "false").
		Param("stdout", "true").
		Param("stderr", "true").
		Param("command", shell).
		Param("tty", "false").
		VersionedParams(&v1.PodExecOptions{
			Container: container,
			Command:   []string{},
			Stdin:     false,
			Stdout:    true,
			Stderr:    true,
			TTY:       false,
		}, scheme.ParameterCodec).URL()

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", reqURL)
	if err != nil {
		return false
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	err = exec.StreamWithContext(context.Background(), remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
		Tty:    false,
	})

	if err != nil {
		return false
	}
	return true
}

func (t *PodExecStream) writeCommandToPodExecStream(taskID string, command []byte) {
	if _, err := getExecStream(taskID); err != nil {
		return
	}
	var msg map[string]uint16

	if err := json.Unmarshal(command, &msg); err != nil {
		_, err = t.stdinWriter.Write(command)
		if err != nil {
			_log.Logger.Errorw(fmt.Sprintf("Closing Stream! Unable to write message to exec stream: %s", err.Error()), "TaskID", taskID)
			t.closeExecStream()
			return
		}

	} else {
		t.terminalQueue.sizeChan <- &remotecommand.TerminalSize{
			Width:  msg["cols"],
			Height: msg["rows"],
		}
	}
}

func (t *PodExecStream) startPodExecStreamReader(taskID string) {
	defer t.wg.Done()
	buf := make([]byte, 2048)
	for {
		n, err := t.stdoutReader.Read(buf)
		if err != nil {
			if _, err = getExecStream(taskID); err != nil {
				return
			}
			_log.Logger.Errorw(fmt.Sprintf("Closing Stream! Unable to read from exec stream: %s", err.Error()), "TaskID", taskID)
			t.closeExecStream()
			return
		}
		if n > 0 {
			output := buf[:n]
			resultMsg := &pb.TaskStreamRequest{
				Payload: &pb.TaskStreamRequest_ExecResp{
					ExecResp: &pb.TerminalExecResponse{
						TaskId:  taskID,
						Success: true,
						Output:  output,
					},
				},
			}

			err = t.grpcStream.Send(resultMsg)
			if err != nil {
				_log.Logger.Errorw(fmt.Sprintf("Closing Stream! Failed to send response to controller: %s", err.Error()), "TaskID", taskID)
				t.closeExecStream()
				return
			}
		}
	}
}

func (t *PodExecStream) startExecStream(taskID string) {
	defer func() {
		t.Close(taskID)
	}()

	err := t.execStream.StreamWithContext(t.execStreamCtx, remotecommand.StreamOptions{
		Stdin:             t.stdinReader,
		Stdout:            t.stdoutWriter,
		Stderr:            t.stderrWriter,
		Tty:               true,
		TerminalSizeQueue: t.terminalQueue,
	})
	if err != nil {
		return
	}
}

func (t *TerminalSizeQueue) Next() *remotecommand.TerminalSize {
	size := <-t.sizeChan
	if size == nil {
		return nil
	}
	return size
}

func (t *PodExecStream) Close(taskID string) {
	mutex.Lock()
	defer mutex.Unlock()
	delete(taskPodExecStreams, taskID)
	delete(heartbeatTimers, taskID)

	_ = t.stdoutWriter.Close()
	_ = t.stdinWriter.Close()
	_ = t.stderrWriter.Close()

	t.wg.Wait()

	_ = t.stdoutReader.Close()
	_ = t.stdinReader.Close()
	_ = t.stderrReader.Close()

	if t.terminalQueue != nil && t.terminalQueue.sizeChan != nil {
		select {
		case <-t.terminalQueue.sizeChan:
		default:
			close(t.terminalQueue.sizeChan)
		}
	}

	resultMsg := &pb.TaskStreamRequest{
		Payload: &pb.TaskStreamRequest_ExecResp{
			ExecResp: &pb.TerminalExecResponse{
				TaskId:  taskID,
				Success: false,
				Output:  []byte("Connection Closed!"),
			},
		},
	}

	_ = t.grpcStream.Send(resultMsg)
	_log.Logger.Infow(fmt.Sprintf("Closed all connection!"), "TaskID", taskID)
}

func (t *PodExecStream) initHeartbeat(taskID string) {
	mutex.Lock()
	defer mutex.Unlock()
	heartbeatTimers[taskID] = time.AfterFunc(10*time.Second, func() {})
}

func (t *PodExecStream) checkHeartbeat(taskID string) {
	mutex.Lock()
	defer mutex.Unlock()

	if timer, exists := heartbeatTimers[taskID]; exists {
		timer.Stop()

		heartbeatTimers[taskID] = time.AfterFunc(10*time.Second, func() {
			if _, err := getExecStream(taskID); err != nil {
				return
			}
			_log.Logger.Errorw(fmt.Sprintf("Closing Stream! Heartbeat Timeout"), "TaskID", taskID)
			t.closeExecStream()
		})
	}
}
