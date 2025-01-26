package tasks

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/krack8/lighthouse/pkg/common/pb"
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
)

var ErrTaskNotExistsRegistryMsg = "task does not exists"
var ErrTaskNotFoundMsg = "task not found"

func TaskSelector(task *pb.Task) (interface{}, error) {
	var res interface{}
	var err error
	newTask := GetTask(task.Name)
	if newTask == nil {
		return nil, errors.New(ErrTaskNotExistsRegistryMsg)
	}
	switch v := newTask.TaskInput.(type) {
	case k8s.GetNamespaceInputParams:
		log.Logger.Infow("Bhua:", v)
	case k8s.GetNamespaceListInputParams:
		log.Logger.Infow("task: "+task.Name+" started.", " task ID#:", task.Id)
		input := k8s.GetNamespaceListInputParams{}
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}

		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNamespaceListInputParams) (interface{}, error))
		if !exists {
			return nil, errors.New(ErrTaskNotFoundMsg)
		}
		res, err = execute(context.Background(), k8s.GetNamespaceListInputParams{})
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
