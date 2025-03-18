// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: pkg/common/pb/controller.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A worker can send two kinds of messages in the stream:
//  1. WorkerIdentification: "I belong to group X, here is my auth token"
//  2. TaskResult: "I finished task <task_id>, here is the result"
type TaskStreamRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*TaskStreamRequest_WorkerInfo
	//	*TaskStreamRequest_TaskResult
	//	*TaskStreamRequest_LogsResult
	Payload       isTaskStreamRequest_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskStreamRequest) Reset() {
	*x = TaskStreamRequest{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskStreamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStreamRequest) ProtoMessage() {}

func (x *TaskStreamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStreamRequest.ProtoReflect.Descriptor instead.
func (*TaskStreamRequest) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{0}
}

func (x *TaskStreamRequest) GetPayload() isTaskStreamRequest_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TaskStreamRequest) GetWorkerInfo() *WorkerIdentification {
	if x != nil {
		if x, ok := x.Payload.(*TaskStreamRequest_WorkerInfo); ok {
			return x.WorkerInfo
		}
	}
	return nil
}

func (x *TaskStreamRequest) GetTaskResult() *TaskResult {
	if x != nil {
		if x, ok := x.Payload.(*TaskStreamRequest_TaskResult); ok {
			return x.TaskResult
		}
	}
	return nil
}

func (x *TaskStreamRequest) GetLogsResult() *LogsResult {
	if x != nil {
		if x, ok := x.Payload.(*TaskStreamRequest_LogsResult); ok {
			return x.LogsResult
		}
	}
	return nil
}

type isTaskStreamRequest_Payload interface {
	isTaskStreamRequest_Payload()
}

type TaskStreamRequest_WorkerInfo struct {
	WorkerInfo *WorkerIdentification `protobuf:"bytes,1,opt,name=worker_info,json=workerInfo,proto3,oneof"`
}

type TaskStreamRequest_TaskResult struct {
	TaskResult *TaskResult `protobuf:"bytes,2,opt,name=task_result,json=taskResult,proto3,oneof"`
}

type TaskStreamRequest_LogsResult struct {
	LogsResult *LogsResult `protobuf:"bytes,3,opt,name=logs_result,json=logsResult,proto3,oneof"`
}

func (*TaskStreamRequest_WorkerInfo) isTaskStreamRequest_Payload() {}

func (*TaskStreamRequest_TaskResult) isTaskStreamRequest_Payload() {}

func (*TaskStreamRequest_LogsResult) isTaskStreamRequest_Payload() {}

type WorkerIdentification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	GroupName     string                 `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
	AuthToken     string                 `protobuf:"bytes,2,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"` // could be a token or JWT, etc.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkerIdentification) Reset() {
	*x = WorkerIdentification{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerIdentification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerIdentification) ProtoMessage() {}

func (x *WorkerIdentification) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerIdentification.ProtoReflect.Descriptor instead.
func (*WorkerIdentification) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerIdentification) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *WorkerIdentification) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

type TaskResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Success       bool                   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	Output        string                 `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskResult) Reset() {
	*x = TaskResult{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResult) ProtoMessage() {}

func (x *TaskResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResult.ProtoReflect.Descriptor instead.
func (*TaskResult) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{2}
}

func (x *TaskResult) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskResult) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TaskResult) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type LogsResult struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Output        []byte                 `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	Cancel        bool                   `protobuf:"varint,3,opt,name=cancel,proto3" json:"cancel,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogsResult) Reset() {
	*x = LogsResult{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogsResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsResult) ProtoMessage() {}

func (x *LogsResult) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsResult.ProtoReflect.Descriptor instead.
func (*LogsResult) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{3}
}

func (x *LogsResult) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *LogsResult) GetOutput() []byte {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *LogsResult) GetCancel() bool {
	if x != nil {
		return x.Cancel
	}
	return false
}

// The controller can send two kinds of messages back to the worker:
//  1. Task: "Please execute this task with ID <task_id> and some payload"
//  2. Ack:  "Acknowledgement for something, or simple keepalive"
type TaskStreamResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*TaskStreamResponse_NewTask
	//	*TaskStreamResponse_Ack
	//	*TaskStreamResponse_NewPodLogsStream
	Payload       isTaskStreamResponse_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TaskStreamResponse) Reset() {
	*x = TaskStreamResponse{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TaskStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskStreamResponse) ProtoMessage() {}

func (x *TaskStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskStreamResponse.ProtoReflect.Descriptor instead.
func (*TaskStreamResponse) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{4}
}

func (x *TaskStreamResponse) GetPayload() isTaskStreamResponse_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *TaskStreamResponse) GetNewTask() *Task {
	if x != nil {
		if x, ok := x.Payload.(*TaskStreamResponse_NewTask); ok {
			return x.NewTask
		}
	}
	return nil
}

func (x *TaskStreamResponse) GetAck() *Ack {
	if x != nil {
		if x, ok := x.Payload.(*TaskStreamResponse_Ack); ok {
			return x.Ack
		}
	}
	return nil
}

func (x *TaskStreamResponse) GetNewPodLogsStream() *PodLogsStream {
	if x != nil {
		if x, ok := x.Payload.(*TaskStreamResponse_NewPodLogsStream); ok {
			return x.NewPodLogsStream
		}
	}
	return nil
}

type isTaskStreamResponse_Payload interface {
	isTaskStreamResponse_Payload()
}

type TaskStreamResponse_NewTask struct {
	NewTask *Task `protobuf:"bytes,1,opt,name=new_task,json=newTask,proto3,oneof"`
}

type TaskStreamResponse_Ack struct {
	Ack *Ack `protobuf:"bytes,2,opt,name=ack,proto3,oneof"`
}

type TaskStreamResponse_NewPodLogsStream struct {
	NewPodLogsStream *PodLogsStream `protobuf:"bytes,3,opt,name=new_pod_logs_stream,json=newPodLogsStream,proto3,oneof"`
}

func (*TaskStreamResponse_NewTask) isTaskStreamResponse_Payload() {}

func (*TaskStreamResponse_Ack) isTaskStreamResponse_Payload() {}

func (*TaskStreamResponse_NewPodLogsStream) isTaskStreamResponse_Payload() {}

type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload       string                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Input         string                 `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{5}
}

func (x *Task) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *Task) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type PodLogsStream struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Payload       string                 `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Input         string                 `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PodLogsStream) Reset() {
	*x = PodLogsStream{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PodLogsStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodLogsStream) ProtoMessage() {}

func (x *PodLogsStream) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodLogsStream.ProtoReflect.Descriptor instead.
func (*PodLogsStream) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{6}
}

func (x *PodLogsStream) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PodLogsStream) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodLogsStream) GetPayload() string {
	if x != nil {
		return x.Payload
	}
	return ""
}

func (x *PodLogsStream) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ack) Reset() {
	*x = Ack{}
	mi := &file_pkg_common_pb_controller_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_common_pb_controller_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_pkg_common_pb_controller_proto_rawDescGZIP(), []int{7}
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_common_pb_controller_proto protoreflect.FileDescriptor

var file_pkg_common_pb_controller_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x22, 0xc1, 0x01, 0x0a, 0x11, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x77, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x00, 0x52, 0x0a,
	0x74, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x31, 0x0a, 0x0b, 0x6c, 0x6f,
	0x67, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48,
	0x00, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x54, 0x0a, 0x14, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x57,
	0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x55, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x73, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06,
	0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x22, 0xa7,
	0x01, 0x0a, 0x12, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x5f, 0x74, 0x61, 0x73,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1b, 0x0a, 0x03,
	0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x41,
	0x63, 0x6b, 0x48, 0x00, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x42, 0x0a, 0x13, 0x6e, 0x65, 0x77,
	0x5f, 0x70, 0x6f, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x64, 0x4c,
	0x6f, 0x67, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x48, 0x00, 0x52, 0x10, 0x6e, 0x65, 0x77,
	0x50, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5a, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x63, 0x0a, 0x0d, 0x50, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x73, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x1f, 0x0a, 0x03, 0x41, 0x63, 0x6b,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4d, 0x0a, 0x0a, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0a, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x70, 0x62, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x6b, 0x67,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_common_pb_controller_proto_rawDescOnce sync.Once
	file_pkg_common_pb_controller_proto_rawDescData = file_pkg_common_pb_controller_proto_rawDesc
)

func file_pkg_common_pb_controller_proto_rawDescGZIP() []byte {
	file_pkg_common_pb_controller_proto_rawDescOnce.Do(func() {
		file_pkg_common_pb_controller_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_common_pb_controller_proto_rawDescData)
	})
	return file_pkg_common_pb_controller_proto_rawDescData
}

var file_pkg_common_pb_controller_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_common_pb_controller_proto_goTypes = []any{
	(*TaskStreamRequest)(nil),    // 0: pb.TaskStreamRequest
	(*WorkerIdentification)(nil), // 1: pb.WorkerIdentification
	(*TaskResult)(nil),           // 2: pb.TaskResult
	(*LogsResult)(nil),           // 3: pb.LogsResult
	(*TaskStreamResponse)(nil),   // 4: pb.TaskStreamResponse
	(*Task)(nil),                 // 5: pb.Task
	(*PodLogsStream)(nil),        // 6: pb.PodLogsStream
	(*Ack)(nil),                  // 7: pb.Ack
}
var file_pkg_common_pb_controller_proto_depIdxs = []int32{
	1, // 0: pb.TaskStreamRequest.worker_info:type_name -> pb.WorkerIdentification
	2, // 1: pb.TaskStreamRequest.task_result:type_name -> pb.TaskResult
	3, // 2: pb.TaskStreamRequest.logs_result:type_name -> pb.LogsResult
	5, // 3: pb.TaskStreamResponse.new_task:type_name -> pb.Task
	7, // 4: pb.TaskStreamResponse.ack:type_name -> pb.Ack
	6, // 5: pb.TaskStreamResponse.new_pod_logs_stream:type_name -> pb.PodLogsStream
	0, // 6: pb.Controller.TaskStream:input_type -> pb.TaskStreamRequest
	4, // 7: pb.Controller.TaskStream:output_type -> pb.TaskStreamResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_pkg_common_pb_controller_proto_init() }
func file_pkg_common_pb_controller_proto_init() {
	if File_pkg_common_pb_controller_proto != nil {
		return
	}
	file_pkg_common_pb_controller_proto_msgTypes[0].OneofWrappers = []any{
		(*TaskStreamRequest_WorkerInfo)(nil),
		(*TaskStreamRequest_TaskResult)(nil),
		(*TaskStreamRequest_LogsResult)(nil),
	}
	file_pkg_common_pb_controller_proto_msgTypes[4].OneofWrappers = []any{
		(*TaskStreamResponse_NewTask)(nil),
		(*TaskStreamResponse_Ack)(nil),
		(*TaskStreamResponse_NewPodLogsStream)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_common_pb_controller_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_common_pb_controller_proto_goTypes,
		DependencyIndexes: file_pkg_common_pb_controller_proto_depIdxs,
		MessageInfos:      file_pkg_common_pb_controller_proto_msgTypes,
	}.Build()
	File_pkg_common_pb_controller_proto = out.File
	file_pkg_common_pb_controller_proto_rawDesc = nil
	file_pkg_common_pb_controller_proto_goTypes = nil
	file_pkg_common_pb_controller_proto_depIdxs = nil
}
