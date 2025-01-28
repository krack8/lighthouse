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
		//daemonSet
	case k8s.GetDaemonSetListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetDaemonSetListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetDaemonSetDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetDaemonSetDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployDaemonSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployDaemonSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteDaemonSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteDaemonSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//deployment
	case k8s.GetDeploymentListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetDeploymentListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetDeploymentDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetDeploymentDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployDeploymentInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployDeploymentInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteDeploymentInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteDeploymentInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetDeploymentStatsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetDeploymentStatsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetDeploymentPodListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetDeploymentPodListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//endpoints
	case k8s.GetEndpointsListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetEndpointsListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetEndpointsDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetEndpointsDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployEndpointsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployEndpointsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteEndpointsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteEndpointsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//endpointSlice
	case k8s.GetEndpointSliceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetEndpointSliceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetEndpointSliceDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetEndpointSliceDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployEndpointSliceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployEndpointSliceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteEndpointSliceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteEndpointSliceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//event
	case k8s.GetEventListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetEventListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetEventDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetEventDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//hpa
	case k8s.GetHpaListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetHpaListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetHpaDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetHpaDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//ingress
	case k8s.GetIngressListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetIngressListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetIngressDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetIngressDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployIngressInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployIngressInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteIngressInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteIngressInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//istioGateway
	case k8s.GetIstioGatewayListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetIstioGatewayListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetIstioGatewayDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetIstioGatewayDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployIstioGatewayInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployIstioGatewayInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteIstioGatewayInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteIstioGatewayInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//loadBalancer
	case k8s.GetLoadBalancerListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetLoadBalancerListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetLoadBalancerDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetLoadBalancerDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//manifest
	case k8s.DeployManifestInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployManifestInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//network
	case k8s.GetNetworkPolicyListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNetworkPolicyListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetNetworkPolicyDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNetworkPolicyDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//node
	case k8s.GetNodeListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNodeListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetNodeInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetNodeInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.NodeCordonInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.NodeCordonInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.NodeTaintInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.NodeTaintInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.NodeUnTaintInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.NodeUnTaintInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//pod
	case k8s.GetPodListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetPodDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetPodStatsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodStatsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetPodLogsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodLogsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployPodInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployPodInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeletePodInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeletePodInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//podDisruptionBudget
	case k8s.GetPodDisruptionBudgetsListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodDisruptionBudgetsListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetPodDisruptionBudgetsDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodDisruptionBudgetsDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployPodDisruptionBudgetsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployPodDisruptionBudgetsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeletePodDisruptionBudgetsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeletePodDisruptionBudgetsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//podMetrics
	case k8s.GetPodMetricsListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodMetricsListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetPodMetricsDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPodMetricsDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//pv
	case k8s.GetPvListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPvListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetPvDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPvDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployPvInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployPvInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeletePvInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeletePvInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//pvc
	case k8s.GetPvcListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPvcListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetPvcDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetPvcDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployPvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployPvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeletePvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeletePvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//replicaset
	case k8s.GetReplicaSetListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetReplicaSetListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetReplicaSetDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetReplicaSetDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployReplicaSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployReplicaSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteReplicaSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteReplicaSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//replicationController
	case k8s.GetReplicationControllerListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetReplicationControllerListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetReplicationControllerDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetReplicationControllerDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployReplicationControllerInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployReplicationControllerInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteReplicationControllerInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteReplicationControllerInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//resourceQuota
	case k8s.GetResourceQuotaListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetResourceQuotaListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetResourceQuotaDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetResourceQuotaDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployResourceQuotaInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployResourceQuotaInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteResourceQuotaInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteResourceQuotaInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//role
	case k8s.GetRoleListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetRoleListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetRoleDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetRoleDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteRoleInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteRoleInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//roleBinding
	case k8s.GetRoleBindingListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetRoleBindingListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetRoleBindingDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetRoleBindingDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteRoleBindingInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteRoleBindingInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//sa
	case k8s.GetServiceAccountListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetServiceAccountListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetServiceAccountDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetServiceAccountDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployServiceAccountInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployServiceAccountInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteServiceAccountInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteServiceAccountInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//secret
	case k8s.GetSecretListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetSecretListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetSecretDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetSecretDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeploySecretInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeploySecretInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteSecretInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteSecretInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//statefulSet
	case k8s.GetStatefulSetListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetStatefulSetListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetStatefulSetDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetStatefulSetDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetStatefulSetPodListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetStatefulSetPodListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetStatefulSetStatsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetStatefulSetStatsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployStatefulSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployStatefulSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteStatefulSetInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteStatefulSetInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//storageClass
	case k8s.GetStorageClassListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetStorageClassListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetStorageClassDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetStorageClassDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployStorageClassInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployStorageClassInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteStorageClassInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteStorageClassInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//svc
	case k8s.GetSvcListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetSvcListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetSvcDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetSvcDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeploySvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeploySvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteSvcInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteSvcInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//virtualService
	case k8s.GetVirtualServiceListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVirtualServiceListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetVirtualServiceDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVirtualServiceDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployVirtualServiceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployVirtualServiceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteVirtualServiceInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteVirtualServiceInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//volumeSnapshot
	case k8s.GetVolumeSnapshotListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVolumeSnapshotListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetVolumeSnapshotDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVolumeSnapshotDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeployVolumeSnapshotInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeployVolumeSnapshotInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.DeleteVolumeSnapshotInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.DeleteVolumeSnapshotInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//volumeSnapshotClass
	case k8s.GetVolumeSnapshotClassListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVolumeSnapshotClassListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetVolumeSnapshotClassDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVolumeSnapshotClassDetailsInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
		//volumeSnapshotContent
	case k8s.GetVolumeSnapshotContentListInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVolumeSnapshotContentListInputParams) (interface{}, error))
		if !exists {
			return nil, ErrTaskNotFound
		}
		res, err = execute(context.Background(), input)
		if err != nil {
			return nil, err
		}
		return res, nil
	case k8s.GetVolumeSnapshotContentDetailsInputParams:
		logTaskStarted(task)
		err = json.Unmarshal([]byte(task.Input), &input)
		if err != nil {
			return nil, err
		}
		execute, exists := newTask.TaskFunc.(func(context.Context, k8s.GetVolumeSnapshotContentDetailsInputParams) (interface{}, error))
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
