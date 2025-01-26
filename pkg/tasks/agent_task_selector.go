package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
)

var ErrTaskNotExistsRegistry = errors.New("task does not exists")
var ErrTaskNotFound = errors.New("task not found")
var ErrUnexpectedTask = errors.New("unexpected task")

func logTaskStarted(task *pb.Task) {
	log.Logger.Infow("Task: "+task.Name+" started.", "task ID#", task.Id)
}
func TaskSelector(task *pb.Task) (interface{}, error) {
	var res interface{}
	var err error
	newTask := GetTask(task.Name)
	if newTask == nil {
		return nil, ErrTaskNotExistsRegistry
	}
	switch input := newTask.TaskInput.(type) {
	//namespace
	case k8s.GetNamespaceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNamespaceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetNamespaceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNamespaceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetNamespaceNamesInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNamespaceNamesInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployNamespaceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployNamespaceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteNamespaceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteNamespaceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	default:
		return nil, ErrUnexpectedTask
	}
}
