// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: rpc/builderrpc/builder.proto

package builderrpc

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

type BWhoAmIResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State string `protobuf:"bytes,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *BWhoAmIResponse) Reset() {
	*x = BWhoAmIResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_builderrpc_builder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BWhoAmIResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BWhoAmIResponse) ProtoMessage() {}

func (x *BWhoAmIResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_builderrpc_builder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BWhoAmIResponse.ProtoReflect.Descriptor instead.
func (*BWhoAmIResponse) Descriptor() ([]byte, []int) {
	return file_rpc_builderrpc_builder_proto_rawDescGZIP(), []int{0}
}

func (x *BWhoAmIResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type BuildRawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Gitlink     string            `protobuf:"bytes,2,opt,name=Gitlink,proto3" json:"Gitlink,omitempty"`
	Branch      string            `protobuf:"bytes,3,opt,name=Branch,proto3" json:"Branch,omitempty"`
	BuildCmd    string            `protobuf:"bytes,4,opt,name=BuildCmd,proto3" json:"BuildCmd,omitempty"`
	StartCmd    string            `protobuf:"bytes,5,opt,name=StartCmd,proto3" json:"StartCmd,omitempty"`
	RuntimeEnv  string            `protobuf:"bytes,6,opt,name=RuntimeEnv,proto3" json:"RuntimeEnv,omitempty"`
	EnvVars     map[string]string `protobuf:"bytes,7,rep,name=EnvVars,proto3" json:"EnvVars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RunningPort string            `protobuf:"bytes,8,opt,name=runningPort,proto3" json:"runningPort,omitempty"`
}

func (x *BuildRawRequest) Reset() {
	*x = BuildRawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_builderrpc_builder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRawRequest) ProtoMessage() {}

func (x *BuildRawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_builderrpc_builder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRawRequest.ProtoReflect.Descriptor instead.
func (*BuildRawRequest) Descriptor() ([]byte, []int) {
	return file_rpc_builderrpc_builder_proto_rawDescGZIP(), []int{1}
}

func (x *BuildRawRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildRawRequest) GetGitlink() string {
	if x != nil {
		return x.Gitlink
	}
	return ""
}

func (x *BuildRawRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *BuildRawRequest) GetBuildCmd() string {
	if x != nil {
		return x.BuildCmd
	}
	return ""
}

func (x *BuildRawRequest) GetStartCmd() string {
	if x != nil {
		return x.StartCmd
	}
	return ""
}

func (x *BuildRawRequest) GetRuntimeEnv() string {
	if x != nil {
		return x.RuntimeEnv
	}
	return ""
}

func (x *BuildRawRequest) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

func (x *BuildRawRequest) GetRunningPort() string {
	if x != nil {
		return x.RunningPort
	}
	return ""
}

type BuildRawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageName   string `protobuf:"bytes,3,opt,name=ImageName,proto3" json:"ImageName,omitempty"`
	RunningPort string `protobuf:"bytes,4,opt,name=RunningPort,proto3" json:"RunningPort,omitempty"`
}

func (x *BuildRawResponse) Reset() {
	*x = BuildRawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_builderrpc_builder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildRawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildRawResponse) ProtoMessage() {}

func (x *BuildRawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_builderrpc_builder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildRawResponse.ProtoReflect.Descriptor instead.
func (*BuildRawResponse) Descriptor() ([]byte, []int) {
	return file_rpc_builderrpc_builder_proto_rawDescGZIP(), []int{2}
}

func (x *BuildRawResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BuildRawResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildRawResponse) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *BuildRawResponse) GetRunningPort() string {
	if x != nil {
		return x.RunningPort
	}
	return ""
}

type BuildSpecRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string            `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Gitlink string            `protobuf:"bytes,2,opt,name=Gitlink,proto3" json:"Gitlink,omitempty"`
	Branch  string            `protobuf:"bytes,3,opt,name=Branch,proto3" json:"Branch,omitempty"`
	EnvVars map[string]string `protobuf:"bytes,4,rep,name=EnvVars,proto3" json:"EnvVars,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BuildSpecRequest) Reset() {
	*x = BuildSpecRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_builderrpc_builder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSpecRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSpecRequest) ProtoMessage() {}

func (x *BuildSpecRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_builderrpc_builder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSpecRequest.ProtoReflect.Descriptor instead.
func (*BuildSpecRequest) Descriptor() ([]byte, []int) {
	return file_rpc_builderrpc_builder_proto_rawDescGZIP(), []int{3}
}

func (x *BuildSpecRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildSpecRequest) GetGitlink() string {
	if x != nil {
		return x.Gitlink
	}
	return ""
}

func (x *BuildSpecRequest) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *BuildSpecRequest) GetEnvVars() map[string]string {
	if x != nil {
		return x.EnvVars
	}
	return nil
}

type BuildSpecResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success     bool   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	ImageName   string `protobuf:"bytes,3,opt,name=ImageName,proto3" json:"ImageName,omitempty"`
	RunningPort string `protobuf:"bytes,4,opt,name=RunningPort,proto3" json:"RunningPort,omitempty"`
}

func (x *BuildSpecResponse) Reset() {
	*x = BuildSpecResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_builderrpc_builder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildSpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildSpecResponse) ProtoMessage() {}

func (x *BuildSpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_builderrpc_builder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildSpecResponse.ProtoReflect.Descriptor instead.
func (*BuildSpecResponse) Descriptor() ([]byte, []int) {
	return file_rpc_builderrpc_builder_proto_rawDescGZIP(), []int{4}
}

func (x *BuildSpecResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *BuildSpecResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BuildSpecResponse) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *BuildSpecResponse) GetRunningPort() string {
	if x != nil {
		return x.RunningPort
	}
	return ""
}

type BuildHealthResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CpuUsage  float32 `protobuf:"fixed32,1,opt,name=CpuUsage,proto3" json:"CpuUsage,omitempty"`
	MemUsage  float32 `protobuf:"fixed32,2,opt,name=MemUsage,proto3" json:"MemUsage,omitempty"`
	DiskUsage float32 `protobuf:"fixed32,3,opt,name=DiskUsage,proto3" json:"DiskUsage,omitempty"`
}

func (x *BuildHealthResponse) Reset() {
	*x = BuildHealthResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_builderrpc_builder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuildHealthResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuildHealthResponse) ProtoMessage() {}

func (x *BuildHealthResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_builderrpc_builder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuildHealthResponse.ProtoReflect.Descriptor instead.
func (*BuildHealthResponse) Descriptor() ([]byte, []int) {
	return file_rpc_builderrpc_builder_proto_rawDescGZIP(), []int{5}
}

func (x *BuildHealthResponse) GetCpuUsage() float32 {
	if x != nil {
		return x.CpuUsage
	}
	return 0
}

func (x *BuildHealthResponse) GetMemUsage() float32 {
	if x != nil {
		return x.MemUsage
	}
	return 0
}

func (x *BuildHealthResponse) GetDiskUsage() float32 {
	if x != nil {
		return x.DiskUsage
	}
	return 0
}

var File_rpc_builderrpc_builder_proto protoreflect.FileDescriptor

var file_rpc_builderrpc_builder_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x72, 0x70, 0x63,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x27, 0x0a, 0x0f, 0x42,
	0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x22, 0xc6, 0x02, 0x0a, 0x0f, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x61,
	0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x47, 0x69, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47,
	0x69, 0x74, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x1a,
	0x0a, 0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6d, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x43, 0x6d, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x6d, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x43, 0x6d, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x45, 0x6e, 0x76, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x52, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x45, 0x6e, 0x76, 0x12, 0x37, 0x0a, 0x07, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72,
	0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80, 0x01,
	0x0a, 0x10, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x72, 0x74,
	0x22, 0xce, 0x01, 0x0a, 0x10, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x47, 0x69, 0x74,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x47, 0x69, 0x74, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x38, 0x0a, 0x07, 0x45,
	0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x45, 0x6e,
	0x76, 0x56, 0x61, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x45, 0x6e, 0x76, 0x56, 0x61, 0x72, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x52, 0x75, 0x6e, 0x6e, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x6b, 0x0a, 0x13, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x43, 0x70, 0x75, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x4d, 0x65, 0x6d, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x44, 0x69, 0x73, 0x6b, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x32, 0xf5, 0x01, 0x0a, 0x0e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x42, 0x57, 0x68, 0x6f, 0x41, 0x6d,
	0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x08, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x52, 0x61, 0x77, 0x12, 0x10, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x52,
	0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34,
	0x0a, 0x09, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x12, 0x11, 0x2e, 0x42, 0x75,
	0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x12, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x1a, 0x14, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x79, 0x75, 0x73, 0x68, 0x31, 0x38,
	0x30, 0x32, 0x33, 0x2f, 0x4c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x5f, 0x46, 0x79, 0x70, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65,
	0x72, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_builderrpc_builder_proto_rawDescOnce sync.Once
	file_rpc_builderrpc_builder_proto_rawDescData = file_rpc_builderrpc_builder_proto_rawDesc
)

func file_rpc_builderrpc_builder_proto_rawDescGZIP() []byte {
	file_rpc_builderrpc_builder_proto_rawDescOnce.Do(func() {
		file_rpc_builderrpc_builder_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_builderrpc_builder_proto_rawDescData)
	})
	return file_rpc_builderrpc_builder_proto_rawDescData
}

var file_rpc_builderrpc_builder_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_rpc_builderrpc_builder_proto_goTypes = []interface{}{
	(*BWhoAmIResponse)(nil),     // 0: BWhoAmIResponse
	(*BuildRawRequest)(nil),     // 1: BuildRawRequest
	(*BuildRawResponse)(nil),    // 2: BuildRawResponse
	(*BuildSpecRequest)(nil),    // 3: BuildSpecRequest
	(*BuildSpecResponse)(nil),   // 4: BuildSpecResponse
	(*BuildHealthResponse)(nil), // 5: BuildHealthResponse
	nil,                         // 6: BuildRawRequest.EnvVarsEntry
	nil,                         // 7: BuildSpecRequest.EnvVarsEntry
	(*emptypb.Empty)(nil),       // 8: google.protobuf.Empty
}
var file_rpc_builderrpc_builder_proto_depIdxs = []int32{
	6, // 0: BuildRawRequest.EnvVars:type_name -> BuildRawRequest.EnvVarsEntry
	7, // 1: BuildSpecRequest.EnvVars:type_name -> BuildSpecRequest.EnvVarsEntry
	8, // 2: BuilderService.WhoAmI:input_type -> google.protobuf.Empty
	1, // 3: BuilderService.BuildRaw:input_type -> BuildRawRequest
	3, // 4: BuilderService.BuildSpec:input_type -> BuildSpecRequest
	8, // 5: BuilderService.BuildHealthMetrics:input_type -> google.protobuf.Empty
	0, // 6: BuilderService.WhoAmI:output_type -> BWhoAmIResponse
	2, // 7: BuilderService.BuildRaw:output_type -> BuildRawResponse
	4, // 8: BuilderService.BuildSpec:output_type -> BuildSpecResponse
	5, // 9: BuilderService.BuildHealthMetrics:output_type -> BuildHealthResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_builderrpc_builder_proto_init() }
func file_rpc_builderrpc_builder_proto_init() {
	if File_rpc_builderrpc_builder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_builderrpc_builder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BWhoAmIResponse); i {
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
		file_rpc_builderrpc_builder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRawRequest); i {
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
		file_rpc_builderrpc_builder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildRawResponse); i {
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
		file_rpc_builderrpc_builder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSpecRequest); i {
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
		file_rpc_builderrpc_builder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildSpecResponse); i {
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
		file_rpc_builderrpc_builder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuildHealthResponse); i {
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
			RawDescriptor: file_rpc_builderrpc_builder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_builderrpc_builder_proto_goTypes,
		DependencyIndexes: file_rpc_builderrpc_builder_proto_depIdxs,
		MessageInfos:      file_rpc_builderrpc_builder_proto_msgTypes,
	}.Build()
	File_rpc_builderrpc_builder_proto = out.File
	file_rpc_builderrpc_builder_proto_rawDesc = nil
	file_rpc_builderrpc_builder_proto_goTypes = nil
	file_rpc_builderrpc_builder_proto_depIdxs = nil
}
