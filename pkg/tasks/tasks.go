package tasks

import (
	"fmt"
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
	fmt.Println("Printing tasks info :" + task.TaskName)
	TaskRegistry[task.TaskName] = task
	fmt.Println("Printing tasks ...")
	fmt.Println(task)
	fmt.Println("Printing tasks registry ...")
	fmt.Println(TaskRegistry)
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
	fmt.Println("Printing tasks Actual Name : " + functionName)
	return functionName
}

func GeneratePayloadTask() {

}
