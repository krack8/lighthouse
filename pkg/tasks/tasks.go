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
	RegisterTask(k8s.NamespaceService().GetNamespaceList, k8s.GetNamespaceListInputParams{})
	RegisterTask(k8s.NamespaceService().GetNamespaceNameList, k8s.GetNamespaceNamesInputParams{})
	RegisterTask(k8s.NamespaceService().GetNamespaceDetails, k8s.GetNamespaceInputParams{})
	RegisterTask(k8s.NamespaceService().DeployNamespace, k8s.DeployNamespaceInputParams{})
	RegisterTask(k8s.NamespaceService().DeleteNamespace, k8s.DeleteNamespaceInputParams{})
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
