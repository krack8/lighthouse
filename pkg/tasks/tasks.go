package tasks

import (
	"github.com/krack8/lighthouse/pkg/k8s"
	"github.com/krack8/lighthouse/pkg/log"
	"reflect"
	"runtime"
	"strings"
	"time"
)

//backlog
//Implement hash key later

type RetryOptions struct {
	InitialInterval time.Duration
	Interval        time.Duration
	RetryAttempts   int
}

type Options struct {
	Timeout         time.Duration
	InitialWaitTime time.Duration
}
type Task struct {
	TaskId       string
	TaskName     string
	TaskGroup    interface{}
	TaskFunc     interface{}
	TaskInput    interface{}
	Options      Options
	RetryOptions RetryOptions
}

var TaskRegistry = make(map[string]*Task)

func RegisterTask(funcTask interface{}, input interface{}) {
	task := &Task{TaskFunc: funcTask}
	task.TaskName = GetFuncName(funcTask)
	task.TaskInput = input
	TaskRegistry[task.TaskName] = task
}

func GetTask(taskName string) *Task {
	task, ok := TaskRegistry[taskName]
	if !ok {
		log.Logger.Errorw("Task %s not found")
		return nil
	}
	return task
}

func InitTaskRegistry() {
	//namespace
	RegisterTask(k8s.NamespaceService().GetNamespaceList, k8s.GetNamespaceListInputParams{})
	RegisterTask(k8s.NamespaceService().GetNamespaceNameList, k8s.GetNamespaceNamesInputParams{})
	RegisterTask(k8s.NamespaceService().GetNamespaceDetails, k8s.GetNamespaceInputParams{})
	RegisterTask(k8s.NamespaceService().DeployNamespace, k8s.DeployNamespaceInputParams{})
	RegisterTask(k8s.NamespaceService().DeleteNamespace, k8s.DeleteNamespaceInputParams{})

	//certficate
	RegisterTask(k8s.CertificateService().GetCertificateList, k8s.GetCertificateListInputParams{})
	RegisterTask(k8s.CertificateService().GetCertificateDetails, k8s.GetCertificateDetailsInputParams{})
	RegisterTask(k8s.CertificateService().DeployCertificate, k8s.DeployCertificateInputParams{})
	RegisterTask(k8s.CertificateService().DeleteCertificate, k8s.DeleteCertificateInputParams{})

	//clusterRole
	RegisterTask(k8s.ClusterRoleService().GetClusterRoleList, k8s.GetClusterRoleListInputParams{})
	RegisterTask(k8s.ClusterRoleService().GetClusterRoleDetails, k8s.GetClusterRoleDetailsInputParams{})
	RegisterTask(k8s.ClusterRoleService().DeployClusterRole, k8s.DeployClusterRoleInputParams{})
	RegisterTask(k8s.ClusterRoleService().DeleteClusterRole, k8s.DeleteClusterRoleInputParams{})

	//clusterRoleBinding
	RegisterTask(k8s.ClusterRoleBindingService().GetClusterRoleBindingList, k8s.GetClusterRoleBindingListInputParams{})
	RegisterTask(k8s.ClusterRoleBindingService().GetClusterRoleBindingDetails, k8s.GetClusterRoleBindingDetailsInputParams{})
	RegisterTask(k8s.ClusterRoleBindingService().DeployClusterRoleBinding, k8s.DeployClusterRoleBindingInputParams{})
	RegisterTask(k8s.ClusterRoleBindingService().DeleteClusterRoleBinding, k8s.DeleteClusterRoleBindingInputParams{})

	//configMap
	RegisterTask(k8s.ConfigMapService().GetConfigMapList, k8s.GetConfigMapListInputParams{})
	RegisterTask(k8s.ConfigMapService().GetConfigMapDetails, k8s.GetConfigMapDetailsInputParams{})
	RegisterTask(k8s.ConfigMapService().DeployConfigMap, k8s.DeployConfigMapInputParams{})
	RegisterTask(k8s.ConfigMapService().DeleteConfigMap, k8s.DeleteConfigMapInputParams{})

	//controllerRevision
	RegisterTask(k8s.ControllerRevisionService().GetControllerRevisionList, k8s.GetControllerRevisionListInputParams{})
	RegisterTask(k8s.ControllerRevisionService().GetControllerRevisionDetails, k8s.GetControllerRevisionDetailsInputParams{})
	RegisterTask(k8s.ControllerRevisionService().DeployControllerRevision, k8s.DeployControllerRevisionInputParams{})
	RegisterTask(k8s.ControllerRevisionService().DeleteControllerRevision, k8s.DeleteControllerRevisionInputParams{})
}

func GetFuncName(funcTask interface{}) string {
	functionName := runtime.FuncForPC(reflect.ValueOf(funcTask).Pointer()).Name()
	lastDotIndex := strings.LastIndex(functionName, ".")
	if lastDotIndex != -1 {
		functionName = functionName[lastDotIndex+1:]
	}
	lastSubsIndex := strings.Index(functionName, "-")
	if lastSubsIndex != -1 {
		functionName = functionName[:lastSubsIndex]
	}
	return functionName
}

func GetCurrentTaskName() string {
	pc, _, _, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()
	lastDotIndex := strings.LastIndex(functionName, ".")
	if lastDotIndex != -1 {
		functionName = functionName[lastDotIndex+1:]
	}
	lastSubsIndex := strings.Index(functionName, "-")
	if lastSubsIndex != -1 {
		functionName = functionName[:lastSubsIndex]
	}
	return functionName
}

func GetTaskName(funcTask interface{}) string {
	return GetFuncName(funcTask)
}
