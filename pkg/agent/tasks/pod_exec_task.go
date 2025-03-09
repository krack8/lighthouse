package tasks

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	"strings"
	"sync"
)

type PodExecStream struct {
	stdinWriter   *io.PipeWriter
	stdinReader   *io.PipeReader
	stdoutReader  *io.PipeReader
	stdoutWriter  *io.PipeWriter
	stderrReader  *io.PipeReader
	stderrWriter  *io.PipeWriter
	terminalQueue *TerminalSizeQueue
	execSession   remotecommand.Executor
	mutex         sync.Mutex
	ctx           context.Context
	cancel        context.CancelFunc
	wg            sync.WaitGroup
}

type TerminalSizeQueue struct {
	sizeChan chan *remotecommand.TerminalSize
}

var taskPodExecStreams = make(map[string]*PodExecStream)
var taskMutex sync.Mutex

func PodExecTask(taskId string, input string, command []byte, isCloseConn bool, stream grpc.BidiStreamingClient[pb.TaskStreamRequest, pb.TaskStreamResponse]) error {
	payload := k8s.PodExecInputParams{}

	err := json.Unmarshal([]byte(input), &payload)
	if err != nil {
		_log.Logger.Error("Agent input error: ", err.Error())
		return err
	}

	streamExec, isNewStream, err := getExecStream(taskId, payload.NamespaceName, payload.PodName, payload.ContainerName, k8s.GetKubeRestConfig())
	if err != nil {
		return errors.New(fmt.Sprintf("failed to get exec stream: %v", err))
	}

	if isCloseConn == true {
		log.Println("[DEBUG] Received stop connection message from output, closing streams")
		streamExec.Close()
		delete(taskPodExecStreams, taskId)
		return nil
	}

	// Goroutine to handle input streaming
	go func(streamExec *PodExecStream, command []byte) {
		log.Println("[DEBUG] Input command:", string(command))
		var msg map[string]uint16

		if err := json.Unmarshal(command, &msg); err != nil {
			if streamExec.stdinWriter == nil {
				log.Println("[DEBUG] No in writer found")
				return
			}

			_, err = streamExec.stdinWriter.Write(command)
			if err != nil {
				log.Println("[DEBUG] Error writing to stdin:", err)
				return
			}

		} else {
			streamExec.terminalQueue.sizeChan <- &remotecommand.TerminalSize{
				Width:  msg["cols"],
				Height: msg["rows"],
			}
			log.Println("[DEBUG] Updated terminal size:", msg["cols"], "x", msg["rows"])
		}
	}(streamExec, command)

	// Goroutine to handle output streaming
	if isNewStream == true {
		//streamExec.wg.Add(1)
		go func(streamExec *PodExecStream) {
			//defer streamExec.wg.Done()
			buf := make([]byte, 2048)
			for {
				n, err := streamExec.stdoutReader.Read(buf)
				if err != nil {
					if err == io.EOF {
						log.Println("[DEBUG] End of file ...")
						return
					} else if errors.Is(err, io.ErrClosedPipe) {
						log.Println("[DEBUG] Attempted to read from a closed pipe")
						return
					}

					log.Println("[DEBUG] Error reading stdout:", err)
					return
				}
				if n > 0 {
					output := buf[:n]
					outputStr := strings.TrimSpace(string(output))

					log.Println("Pod Output:", outputStr)
					resultMsg := &pb.TaskStreamRequest{
						Payload: &pb.TaskStreamRequest_ExecResp{
							ExecResp: &pb.TerminalExecResponse{
								TaskId:  taskId,
								Success: true,
								Output:  output,
							},
						},
					}

					err = stream.Send(resultMsg)
					if err != nil {
						log.Println("Error sending output to gRPC:", err)
						break
					}
				}
			}
		}(streamExec)
	}
	return nil
}

func getExecStream(taskID, namespace, podName, containerName string, config *rest.Config) (*PodExecStream, bool, error) {
	taskMutex.Lock()
	defer taskMutex.Unlock()

	if stream, exists := taskPodExecStreams[taskID]; exists {
		return stream, false, nil
	}

	stdinReader, stdinWriter := io.Pipe()
	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	shellCmd := "/bin/bash"
	if !isShellAvailable(config, namespace, podName, containerName, "/bin/bash") {
		log.Println("Falling back to /bin/sh")
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
		return nil, false, fmt.Errorf("failed to create SPDY executor: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	stream := &PodExecStream{
		stdinWriter:   stdinWriter,
		stdinReader:   stdinReader,
		stdoutReader:  stdoutReader,
		ctx:           ctx,
		cancel:        cancel,
		stdoutWriter:  stdoutWriter,
		stderrReader:  stderrReader,
		stderrWriter:  stderrWriter,
		execSession:   exec,
		terminalQueue: &TerminalSizeQueue{sizeChan: make(chan *remotecommand.TerminalSize, 1)},
	}

	taskPodExecStreams[taskID] = stream

	// Start exec stream
	stream.wg.Add(1)
	go stream.startExecStream()

	return stream, true, nil
}

func (t *PodExecStream) startExecStream() {
	defer t.wg.Done()
	err := t.execSession.StreamWithContext(t.ctx, remotecommand.StreamOptions{
		Stdin:             t.stdinReader,
		Stdout:            t.stdoutWriter,
		Stderr:            t.stderrWriter,
		Tty:               true,
		TerminalSizeQueue: t.terminalQueue,
	})
	if err != nil {
		log.Println("[DEBUG ]Error in exec stream:", err)
		return
	}
}

func (t *TerminalSizeQueue) Next() *remotecommand.TerminalSize {
	size := <-t.sizeChan
	if size == nil {
		return nil
	}
	log.Println(fmt.Sprintf("terminal size to width: %d height: %d", size.Width, size.Height))
	return size
}

func (t *PodExecStream) Close() {
	t.mutex.Lock()
	defer t.mutex.Unlock()

	t.cancel()

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
	log.Println("[DEBUG] Closed all pod exec streams successfully")
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
		log.Println("Shell", shell, "is not available:", err)
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
		_log.Logger.Error("Error executing shell: ", shell, ":", err)
		return false
	}
	return true
}
