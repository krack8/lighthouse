package tasks

import (
	"bytes"
	"context"
	"encoding/json"
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
	"log"
	"sync"
	"time"
)

type PodExecStream struct {
	stdinWriter      *io.PipeWriter
	stdinReader      *io.PipeReader
	stdoutReader     *io.PipeReader
	stdoutWriter     *io.PipeWriter
	stderrReader     *io.PipeReader
	stderrWriter     *io.PipeWriter
	terminalQueue    *TerminalSizeQueue
	execStream       remotecommand.Executor
	execStreamCtx    context.Context
	execStreamCancel context.CancelFunc
	wg               sync.WaitGroup
}

type TerminalSizeQueue struct {
	sizeChan chan *remotecommand.TerminalSize
}

var taskPodExecStreams = make(map[string]*PodExecStream)
var mutex sync.Mutex
var heartbeatTimers = make(map[string]*time.Timer)

func PodExecTask(taskId string, payloadType string, input string, command []byte, grpcStream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) error {
	inputPayload := k8s.PodExecInputParams{}

	err := json.Unmarshal([]byte(input), &inputPayload)
	if err != nil {
		_log.Logger.Error(fmt.Sprintf("Closing Stream! Input payload error: %s", err.Error()), "TaskID", taskId, "Payload", payloadType)
		return err
	}

	streamExec, err := getExecStream(taskId, inputPayload.NamespaceName, inputPayload.PodName, inputPayload.ContainerName, k8s.GetKubeRestConfig())
	if err != nil {
		_log.Logger.Error(fmt.Sprintf("Closing Stream! Failed to get pod exec stream: %s", err.Error()), "TaskID", taskId, "Payload", payloadType)
		return err
	}

	if payloadType == consts.TaskPodExecCloseConn {
		log.Println("[DEBUG] Received stop connection message from output, closing streams", "TaskID", taskId, "Payload", payloadType)
		streamExec.Close(taskId)
		return nil

	} else if payloadType == consts.TaskPodExecHeartbeat {
		log.Println("[DEBUG] Received heart beat message", "TaskID", taskId, "Payload", payloadType)
		streamExec.checkHeartbeat(taskId)
		return nil

	} else if payloadType == consts.TaskPodExecInitConn {
		log.Println("[DEBUG] Received init connection", "TaskID", taskId, "Payload", payloadType)
		go streamExec.startPodExecStreamReader(taskId, grpcStream)
	}

	// Writing commands to pod exec stream
	go streamExec.writeCommandToPodExecStream(taskId, command)
	return nil
}

func getExecStream(taskID, namespace, podName, containerName string, config *rest.Config) (*PodExecStream, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if stream, exists := taskPodExecStreams[taskID]; exists {
		return stream, nil
	}

	stdinReader, stdinWriter := io.Pipe()
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	shellCmd := "/bin/bash"
	if !isShellAvailable(config, namespace, podName, containerName, "/bin/bash") {
		log.Println("[DEBUG] Falling back to /bin/sh")
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
		stdinWriter:      stdinWriter,
		stdinReader:      stdinReader,
		stdoutReader:     stdoutReader,
		execStreamCtx:    ctx,
		execStreamCancel: cancel,
		stdoutWriter:     stdoutWriter,
		stderrReader:     stderrReader,
		stderrWriter:     stderrWriter,
		execStream:       exec,
		terminalQueue:    &TerminalSizeQueue{sizeChan: make(chan *remotecommand.TerminalSize, 1)},
	}

	taskPodExecStreams[taskID] = stream

	// Start exec stream
	stream.wg.Add(1)
	go stream.startExecStream()

	return stream, nil
}

func (t *PodExecStream) writeCommandToPodExecStream(taskID string, command []byte) {
	var msg map[string]uint16

	if err := json.Unmarshal(command, &msg); err != nil {
		_, err = t.stdinWriter.Write(command)
		if err != nil {
			_log.Logger.Error(fmt.Sprintf("Closing Stream! Unable to write message to pod exec stream: %s", err.Error()), "TaskID", taskID)
			t.Close(taskID)
			return
		}

	} else {
		t.terminalQueue.sizeChan <- &remotecommand.TerminalSize{
			Width:  msg["cols"],
			Height: msg["rows"],
		}
	}
}

func (t *PodExecStream) startPodExecStreamReader(taskID string, grpcStream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) {
	buf := make([]byte, 2048)
	for {
		n, err := t.stdoutReader.Read(buf)
		if err != nil {
			_log.Logger.Error(fmt.Sprintf("Closing Stream! Unable to read from pod exec stream: %s", err.Error()), "TaskID", taskID)
			t.Close(taskID)
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

			err = grpcStream.Send(resultMsg)
			if err != nil {
				_log.Logger.Error(fmt.Sprintf("Closing Stream! Failed to send response to controller: %s", err.Error()), "TaskID", taskID)
				t.Close(taskID)
				return
			}
		}
	}
}

func (t *PodExecStream) startExecStream() {
	defer t.wg.Done()
	err := t.execStream.StreamWithContext(t.execStreamCtx, remotecommand.StreamOptions{
		Stdin:             t.stdinReader,
		Stdout:            t.stdoutWriter,
		Stderr:            t.stderrWriter,
		Tty:               true,
		TerminalSizeQueue: t.terminalQueue,
	})
	if err != nil {
		log.Println("[DEBUG] Error in exec stream:", err)
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

	t.execStreamCancel()

	if t.stdoutReader != nil {
		_ = t.stdoutReader.Close()
		t.stdoutReader = nil
	}

	if t.stdinWriter != nil {
		_ = t.stdinWriter.Close()
		t.stdinWriter = nil
	}

	if t.stdoutWriter != nil {
		_ = t.stdoutWriter.Close()
		t.stdoutWriter = nil
	}

	if t.stderrWriter != nil {
		_ = t.stderrWriter.Close()
		t.stderrWriter = nil
	}

	t.wg.Wait()

	if t.stdinReader != nil {
		_ = t.stdinReader.Close()
		t.stdinReader = nil
	}

	if t.stderrReader != nil {
		_ = t.stderrReader.Close()
		t.stderrReader = nil
	}

	if t.terminalQueue != nil && t.terminalQueue.sizeChan != nil {
		select {
		case <-t.terminalQueue.sizeChan:
		default:
			close(t.terminalQueue.sizeChan)
		}
	}

	delete(taskPodExecStreams, taskID)
	_log.Logger.Infow(fmt.Sprintf("Closed all connection!"), "TaskID", taskID)
}

func (t *PodExecStream) checkHeartbeat(taskID string) {
	mutex.Lock()
	defer mutex.Unlock()

	if timer, exists := heartbeatTimers[taskID]; exists {
		timer.Stop()
	}

	heartbeatTimers[taskID] = time.AfterFunc(10*time.Second, func() {
		_log.Logger.Error(fmt.Sprintf("Closing Stream! Heartbeat Timeout"), "TaskID", taskID)
		t.Close(taskID)
	})
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
