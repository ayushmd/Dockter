// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: rpc/workerrpc/worker.proto

package workerrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WWhoAmIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *WWhoAmIResponse) Reset() {
	*x = WWhoAmIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WWhoAmIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WWhoAmIResponse) ProtoMessage() {}

func (x *WWhoAmIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WWhoAmIResponse.ProtoReflect.Descriptor instead.
func (*WWhoAmIResponse) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{0}
}

func (x *WWhoAmIResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type HealthScoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Score float32 `protobuf:"fixed32,1,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *HealthScoreResponse) Reset() {
	*x = HealthScoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthScoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthScoreResponse) ProtoMessage() {}

func (x *HealthScoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthScoreResponse.ProtoReflect.Descriptor instead.
func (*HealthScoreResponse) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{1}
}

func (x *HealthScoreResponse) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

type HealthMetricResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuUsage  float32 `protobuf:"fixed32,1,opt,name=CpuUsage,proto3" json:"CpuUsage,omitempty"`
	MemUsage  float32 `protobuf:"fixed32,2,opt,name=MemUsage,proto3" json:"MemUsage,omitempty"`
	DiskUsage float32 `protobuf:"fixed32,3,opt,name=DiskUsage,proto3" json:"DiskUsage,omitempty"`
}

func (x *HealthMetricResponse) Reset() {
	*x = HealthMetricResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthMetricResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthMetricResponse) ProtoMessage() {}

func (x *HealthMetricResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthMetricResponse.ProtoReflect.Descriptor instead.
func (*HealthMetricResponse) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{2}
}

func (x *HealthMetricResponse) GetCpuUsage() float32 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *HealthMetricResponse) GetMemUsage() float32 {
	if x != nil {
		return x.MemUsage
	}
	return 0
}

func (x *HealthMetricResponse) GetDiskUsage() float32 {
	if x != nil {
		return x.DiskUsage
	}
	return 0
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageName   string `protobuf:"bytes,2,opt,name=ImageName,proto3" json:"ImageName,omitempty"`
	RunningPort string `protobuf:"bytes,3,opt,name=RunningPort,proto3" json:"RunningPort,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{3}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *Task) GetRunningPort() string {
	if x != nil {
		return x.RunningPort
	}
	return ""
}

type Port struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IP          string `protobuf:"bytes,1,opt,name=IP,proto3" json:"IP,omitempty"`
	RunningPort int32  `protobuf:"varint,2,opt,name=RunningPort,proto3" json:"RunningPort,omitempty"`
	HostPort    int32  `protobuf:"varint,3,opt,name=HostPort,proto3" json:"HostPort,omitempty"`
	Type        string `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Port) Reset() {
	*x = Port{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Port) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Port) ProtoMessage() {}

func (x *Port) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Port.ProtoReflect.Descriptor instead.
func (*Port) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{4}
}

func (x *Port) GetIP() string {
	if x != nil {
		return x.IP
	}
	return ""
}

func (x *Port) GetRunningPort() int32 {
	if x != nil {
		return x.RunningPort
	}
	return 0
}

func (x *Port) GetHostPort() int32 {
	if x != nil {
		return x.HostPort
	}
	return 0
}

func (x *Port) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type RunningTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string  `protobuf:"bytes,1,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	ImageName   string  `protobuf:"bytes,2,opt,name=ImageName,proto3" json:"ImageName,omitempty"`
	Ports       []*Port `protobuf:"bytes,3,rep,name=Ports,proto3" json:"Ports,omitempty"`
	State       string  `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *RunningTask) Reset() {
	*x = RunningTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunningTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunningTask) ProtoMessage() {}

func (x *RunningTask) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunningTask.ProtoReflect.Descriptor instead.
func (*RunningTask) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{5}
}

func (x *RunningTask) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *RunningTask) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *RunningTask) GetPorts() []*Port {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *RunningTask) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type RunningTasks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Containers []*RunningTask `protobuf:"bytes,1,rep,name=containers,proto3" json:"containers,omitempty"`
}

func (x *RunningTasks) Reset() {
	*x = RunningTasks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunningTasks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunningTasks) ProtoMessage() {}

func (x *RunningTasks) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunningTasks.ProtoReflect.Descriptor instead.
func (*RunningTasks) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{6}
}

func (x *RunningTasks) GetContainers() []*RunningTask {
	if x != nil {
		return x.Containers
	}
	return nil
}

type AddTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContainerID string `protobuf:"bytes,1,opt,name=ContainerID,proto3" json:"ContainerID,omitempty"`
	HostPort    string `protobuf:"bytes,2,opt,name=HostPort,proto3" json:"HostPort,omitempty"`
}

func (x *AddTaskResponse) Reset() {
	*x = AddTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddTaskResponse) ProtoMessage() {}

func (x *AddTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddTaskResponse.ProtoReflect.Descriptor instead.
func (*AddTaskResponse) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{7}
}

func (x *AddTaskResponse) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *AddTaskResponse) GetHostPort() string {
	if x != nil {
		return x.HostPort
	}
	return ""
}

type TerminateTaskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *TerminateTaskRequest) Reset() {
	*x = TerminateTaskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateTaskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateTaskRequest) ProtoMessage() {}

func (x *TerminateTaskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateTaskRequest.ProtoReflect.Descriptor instead.
func (*TerminateTaskRequest) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{8}
}

func (x *TerminateTaskRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TerminateTaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success string `protobuf:"bytes,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *TerminateTaskResponse) Reset() {
	*x = TerminateTaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminateTaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminateTaskResponse) ProtoMessage() {}

func (x *TerminateTaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminateTaskResponse.ProtoReflect.Descriptor instead.
func (*TerminateTaskResponse) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{9}
}

func (x *TerminateTaskResponse) GetSuccess() string {
	if x != nil {
		return x.Success
	}
	return ""
}

type Tasks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *Tasks) Reset() {
	*x = Tasks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_workerrpc_worker_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tasks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tasks) ProtoMessage() {}

func (x *Tasks) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_workerrpc_worker_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tasks.ProtoReflect.Descriptor instead.
func (*Tasks) Descriptor() ([]byte, []int) {
	return file_rpc_workerrpc_worker_proto_rawDescGZIP(), []int{10}
}

func (x *Tasks) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_rpc_workerrpc_worker_proto protoreflect.FileDescriptor

var file_rpc_workerrpc_worker_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x72, 0x70, 0x63, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x72, 0x70, 0x63, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0f, 0x57, 0x57, 0x68,
	0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x2b, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22,
	0x6c, 0x0a, 0x14, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x70, 0x75, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x43, 0x70, 0x75, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x4d, 0x65, 0x6d, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5a, 0x0a,
	0x04, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69,
	0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x75,
	0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x68, 0x0a, 0x04, 0x50, 0x6f, 0x72,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49,
	0x50, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x73, 0x6b, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x05, 0x50, 0x6f, 0x72, 0x74, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x2c, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x52, 0x75, 0x6e,
	0x6e, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x73, 0x22, 0x4f, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x48, 0x6f, 0x73,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x48, 0x6f, 0x73,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x2a, 0x0a, 0x14, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x31, 0x0a, 0x15, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x61,
	0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x24, 0x0a, 0x05, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x1b, 0x0a,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x32, 0xe3, 0x02, 0x0a, 0x0d, 0x57,
	0x6f, 0x72, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06,
	0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x57, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x72,
	0x65, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x0d, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x15, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x24, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x05,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x1a, 0x10, 0x2e, 0x41, 0x64, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0d, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x15, 0x2e, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x0d, 0x2e, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x22, 0x00,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x79, 0x75, 0x73, 0x68, 0x31, 0x38, 0x30, 0x32, 0x33, 0x2f, 0x4c, 0x6f, 0x61, 0x64, 0x5f, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x46, 0x79, 0x70, 0x2f, 0x72, 0x70, 0x63, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_workerrpc_worker_proto_rawDescOnce sync.Once
	file_rpc_workerrpc_worker_proto_rawDescData = file_rpc_workerrpc_worker_proto_rawDesc
)

func file_rpc_workerrpc_worker_proto_rawDescGZIP() []byte {
	file_rpc_workerrpc_worker_proto_rawDescOnce.Do(func() {
		file_rpc_workerrpc_worker_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_workerrpc_worker_proto_rawDescData)
	})
	return file_rpc_workerrpc_worker_proto_rawDescData
}

var file_rpc_workerrpc_worker_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rpc_workerrpc_worker_proto_goTypes = []interface{}{
	(*WWhoAmIResponse)(nil),       // 0: WWhoAmIResponse
	(*HealthScoreResponse)(nil),   // 1: HealthScoreResponse
	(*HealthMetricResponse)(nil),  // 2: HealthMetricResponse
	(*Task)(nil),                  // 3: Task
	(*Port)(nil),                  // 4: Port
	(*RunningTask)(nil),           // 5: RunningTask
	(*RunningTasks)(nil),          // 6: RunningTasks
	(*AddTaskResponse)(nil),       // 7: AddTaskResponse
	(*TerminateTaskRequest)(nil),  // 8: TerminateTaskRequest
	(*TerminateTaskResponse)(nil), // 9: TerminateTaskResponse
	(*Tasks)(nil),                 // 10: Tasks
	(*emptypb.Empty)(nil),         // 11: google.protobuf.Empty
}
var file_rpc_workerrpc_worker_proto_depIdxs = []int32{
	4,  // 0: RunningTask.Ports:type_name -> Port
	5,  // 1: RunningTasks.containers:type_name -> RunningTask
	3,  // 2: Tasks.tasks:type_name -> Task
	11, // 3: WorkerService.WhoAmI:input_type -> google.protobuf.Empty
	11, // 4: WorkerService.HealthScore:input_type -> google.protobuf.Empty
	11, // 5: WorkerService.HealthMetrics:input_type -> google.protobuf.Empty
	3,  // 6: WorkerService.AddTask:input_type -> Task
	8,  // 7: WorkerService.TerminateTask:input_type -> TerminateTaskRequest
	11, // 8: WorkerService.GetTasks:input_type -> google.protobuf.Empty
	0,  // 9: WorkerService.WhoAmI:output_type -> WWhoAmIResponse
	1,  // 10: WorkerService.HealthScore:output_type -> HealthScoreResponse
	2,  // 11: WorkerService.HealthMetrics:output_type -> HealthMetricResponse
	7,  // 12: WorkerService.AddTask:output_type -> AddTaskResponse
	9,  // 13: WorkerService.TerminateTask:output_type -> TerminateTaskResponse
	6,  // 14: WorkerService.GetTasks:output_type -> RunningTasks
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_rpc_workerrpc_worker_proto_init() }
func file_rpc_workerrpc_worker_proto_init() {
	if File_rpc_workerrpc_worker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_workerrpc_worker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WWhoAmIResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthScoreResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthMetricResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Port); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunningTask); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunningTasks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddTaskResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateTaskRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminateTaskResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_rpc_workerrpc_worker_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tasks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_workerrpc_worker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_workerrpc_worker_proto_goTypes,
		DependencyIndexes: file_rpc_workerrpc_worker_proto_depIdxs,
		MessageInfos:      file_rpc_workerrpc_worker_proto_msgTypes,
	}.Build()
	File_rpc_workerrpc_worker_proto = out.File
	file_rpc_workerrpc_worker_proto_rawDesc = nil
	file_rpc_workerrpc_worker_proto_goTypes = nil
	file_rpc_workerrpc_worker_proto_depIdxs = nil
}
