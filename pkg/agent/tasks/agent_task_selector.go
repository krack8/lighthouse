package tasks

import (
	"context"
	"encoding/json"
	"errors"
	k8s2 "github.com/krack8/lighthouse/pkg/common/k8s"
	"github.com/krack8/lighthouse/pkg/common/log"
	"github.com/krack8/lighthouse/pkg/common/pb"
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
	case k8s2.GetNamespaceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetNamespaceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetNamespaceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetNamespaceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetNamespaceNamesInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetNamespaceNamesInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployNamespaceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployNamespaceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteNamespaceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteNamespaceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//Certificate
	case k8s2.GetCertificateListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCertificateListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetCertificateDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCertificateDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployCertificateInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployCertificateInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteCertificateInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteCertificateInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ClusterRole
	case k8s2.GetClusterRoleListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetClusterRoleListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetClusterRoleDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetClusterRoleDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployClusterRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployClusterRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteClusterRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteClusterRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ClusterRoleBinding
	case k8s2.GetClusterRoleBindingListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetClusterRoleBindingListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetClusterRoleBindingDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetClusterRoleBindingDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployClusterRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployClusterRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteClusterRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteClusterRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ConfigMap
	case k8s2.GetConfigMapListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetConfigMapListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetConfigMapDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetConfigMapDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployConfigMapInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployConfigMapInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteConfigMapInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteConfigMapInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ControllerRevision
	case k8s2.GetControllerRevisionListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetControllerRevisionListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetControllerRevisionDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetControllerRevisionDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployControllerRevisionInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployControllerRevisionInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteControllerRevisionInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteControllerRevisionInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//CRD
	case k8s2.GetCrdListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCrdListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetCrdDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCrdDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployCrdInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployCrdInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteCrdInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteCrdInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//cronJob
	case k8s2.GetCronJobListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCronJobListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetCronJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCronJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployCronJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployCronJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteCronJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteCronJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//customResource
	case k8s2.GetCustomResourceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCustomResourceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetCustomResourceDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetCustomResourceDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployCustomResourceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployCustomResourceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteCustomResourceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteCustomResourceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//daemonSet
	case k8s2.GetDaemonSetListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetDaemonSetListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetDaemonSetDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetDaemonSetDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetDaemonSetStatsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetDaemonSetStatsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployDaemonSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployDaemonSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteDaemonSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteDaemonSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//deployment
	case k8s2.GetDeploymentListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetDeploymentListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetDeploymentDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetDeploymentDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployDeploymentInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployDeploymentInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteDeploymentInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteDeploymentInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetDeploymentStatsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetDeploymentStatsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetDeploymentPodListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetDeploymentPodListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//endpoints
	case k8s2.GetEndpointsListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetEndpointsListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetEndpointsDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetEndpointsDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployEndpointsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployEndpointsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteEndpointsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteEndpointsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//endpointSlice
	case k8s2.GetEndpointSliceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetEndpointSliceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetEndpointSliceDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetEndpointSliceDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployEndpointSliceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployEndpointSliceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteEndpointSliceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteEndpointSliceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//event
	case k8s2.GetEventListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetEventListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetEventDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetEventDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//hpa
	case k8s2.GetHpaListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetHpaListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetHpaDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetHpaDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ingress
	case k8s2.GetIngressListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetIngressListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetIngressDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetIngressDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployIngressInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployIngressInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteIngressInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteIngressInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//istioGateway
	case k8s2.GetIstioGatewayListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetIstioGatewayListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetIstioGatewayDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetIstioGatewayDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployIstioGatewayInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployIstioGatewayInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteIstioGatewayInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteIstioGatewayInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//job
	case k8s2.GetJobListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetJobListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteJobInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteJobInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//loadBalancer
	case k8s2.GetLoadBalancerListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetLoadBalancerListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetLoadBalancerDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetLoadBalancerDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//manifest
	case k8s2.DeployManifestInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployManifestInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//network
	case k8s2.GetNetworkPolicyListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetNetworkPolicyListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetNetworkPolicyDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetNetworkPolicyDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//node
	case k8s2.GetNodeListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetNodeListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetNodeInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetNodeInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.NodeCordonInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.NodeCordonInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.NodeTaintInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.NodeTaintInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.NodeUnTaintInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.NodeUnTaintInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//pod
	case k8s2.GetPodListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetPodDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetPodStatsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodStatsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetPodLogsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodLogsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployPodInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployPodInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeletePodInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeletePodInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//podDisruptionBudget
	case k8s2.GetPodDisruptionBudgetsListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodDisruptionBudgetsListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetPodDisruptionBudgetsDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodDisruptionBudgetsDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployPodDisruptionBudgetsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployPodDisruptionBudgetsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeletePodDisruptionBudgetsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeletePodDisruptionBudgetsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//podMetrics
	case k8s2.GetPodMetricsListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodMetricsListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetPodMetricsDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPodMetricsDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//pv
	case k8s2.GetPvListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPvListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetPvDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPvDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployPvInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployPvInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeletePvInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeletePvInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//pvc
	case k8s2.GetPvcListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPvcListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetPvcDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetPvcDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployPvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployPvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeletePvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeletePvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//replicaset
	case k8s2.GetReplicaSetListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetReplicaSetListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetReplicaSetDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetReplicaSetDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployReplicaSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployReplicaSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteReplicaSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteReplicaSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//replicationController
	case k8s2.GetReplicationControllerListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetReplicationControllerListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetReplicationControllerDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetReplicationControllerDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployReplicationControllerInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployReplicationControllerInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteReplicationControllerInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteReplicationControllerInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//resourceQuota
	case k8s2.GetResourceQuotaListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetResourceQuotaListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetResourceQuotaDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetResourceQuotaDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployResourceQuotaInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployResourceQuotaInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteResourceQuotaInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteResourceQuotaInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//role
	case k8s2.GetRoleListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetRoleListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetRoleDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetRoleDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//roleBinding
	case k8s2.GetRoleBindingListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetRoleBindingListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetRoleBindingDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetRoleBindingDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//sa
	case k8s2.GetServiceAccountListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetServiceAccountListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetServiceAccountDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetServiceAccountDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployServiceAccountInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployServiceAccountInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteServiceAccountInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteServiceAccountInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//secret
	case k8s2.GetSecretListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetSecretListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetSecretDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetSecretDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeploySecretInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeploySecretInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteSecretInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteSecretInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//statefulSet
	case k8s2.GetStatefulSetListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetStatefulSetListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetStatefulSetDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetStatefulSetDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetStatefulSetPodListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetStatefulSetPodListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetStatefulSetStatsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetStatefulSetStatsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployStatefulSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployStatefulSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteStatefulSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteStatefulSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//storageClass
	case k8s2.GetStorageClassListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetStorageClassListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetStorageClassDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetStorageClassDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployStorageClassInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployStorageClassInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteStorageClassInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteStorageClassInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//svc
	case k8s2.GetSvcListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetSvcListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetSvcDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetSvcDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeploySvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeploySvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteSvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteSvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//virtualService
	case k8s2.GetVirtualServiceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVirtualServiceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetVirtualServiceDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVirtualServiceDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployVirtualServiceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployVirtualServiceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteVirtualServiceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteVirtualServiceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//volumeSnapshot
	case k8s2.GetVolumeSnapshotListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVolumeSnapshotListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetVolumeSnapshotDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVolumeSnapshotDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeployVolumeSnapshotInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeployVolumeSnapshotInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.DeleteVolumeSnapshotInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.DeleteVolumeSnapshotInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//volumeSnapshotClass
	case k8s2.GetVolumeSnapshotClassListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVolumeSnapshotClassListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetVolumeSnapshotClassDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVolumeSnapshotClassDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//volumeSnapshotContent
	case k8s2.GetVolumeSnapshotContentListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVolumeSnapshotContentListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s2.GetVolumeSnapshotContentDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s2.GetVolumeSnapshotContentDetailsInputParams) (interface{}, error))
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
