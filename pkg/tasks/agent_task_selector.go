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
		//Certificate
	case k8s.GetCertificateListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCertificateListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetCertificateDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCertificateDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployCertificateInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployCertificateInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteCertificateInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteCertificateInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ClusterRole
	case k8s.GetClusterRoleListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetClusterRoleListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetClusterRoleDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetClusterRoleDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployClusterRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployClusterRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteClusterRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteClusterRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ClusterRoleBinding
	case k8s.GetClusterRoleBindingListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetClusterRoleBindingListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetClusterRoleBindingDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetClusterRoleBindingDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployClusterRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployClusterRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteClusterRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteClusterRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ConfigMap
	case k8s.GetConfigMapListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetConfigMapListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetConfigMapDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetConfigMapDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployConfigMapInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployConfigMapInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteConfigMapInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteConfigMapInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ControllerRevision
	case k8s.GetControllerRevisionListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetControllerRevisionListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetControllerRevisionDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetControllerRevisionDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployControllerRevisionInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployControllerRevisionInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteControllerRevisionInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteControllerRevisionInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//CRD
	case k8s.GetCrdListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCrdListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetCrdDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCrdDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployCrdInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployCrdInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteCrdInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteCrdInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//cronJob
	case k8s.GetCronJobListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCronJobListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetCronJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCronJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployCronJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployCronJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteCronJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteCronJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//customResource
	case k8s.GetCustomResourceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCustomResourceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetCustomResourceDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetCustomResourceDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployCustomResourceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployCustomResourceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteCustomResourceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteCustomResourceInputParams) (interface{}, error))
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
