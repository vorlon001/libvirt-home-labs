// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v4.23.4
// source: virsh.proto

package server

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type VirshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmname string `protobuf:"bytes,1,opt,name=vmname,proto3" json:"vmname,omitempty"`
}

func (x *VirshRequest) Reset() {
	*x = VirshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virsh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirshRequest) ProtoMessage() {}

func (x *VirshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_virsh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirshRequest.ProtoReflect.Descriptor instead.
func (*VirshRequest) Descriptor() ([]byte, []int) {
	return file_virsh_proto_rawDescGZIP(), []int{0}
}

func (x *VirshRequest) GetVmname() string {
	if x != nil {
		return x.Vmname
	}
	return ""
}

type VirshMachineMigrate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vmname string `protobuf:"bytes,1,opt,name=vmname,proto3" json:"vmname,omitempty"`
	Tomove string `protobuf:"bytes,2,opt,name=tomove,proto3" json:"tomove,omitempty"`
}

func (x *VirshMachineMigrate) Reset() {
	*x = VirshMachineMigrate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virsh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirshMachineMigrate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirshMachineMigrate) ProtoMessage() {}

func (x *VirshMachineMigrate) ProtoReflect() protoreflect.Message {
	mi := &file_virsh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirshMachineMigrate.ProtoReflect.Descriptor instead.
func (*VirshMachineMigrate) Descriptor() ([]byte, []int) {
	return file_virsh_proto_rawDescGZIP(), []int{1}
}

func (x *VirshMachineMigrate) GetVmname() string {
	if x != nil {
		return x.Vmname
	}
	return ""
}

func (x *VirshMachineMigrate) GetTomove() string {
	if x != nil {
		return x.Tomove
	}
	return ""
}

type VirshCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xml string `protobuf:"bytes,1,opt,name=xml,proto3" json:"xml,omitempty"`
}

func (x *VirshCreateRequest) Reset() {
	*x = VirshCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virsh_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirshCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirshCreateRequest) ProtoMessage() {}

func (x *VirshCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_virsh_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirshCreateRequest.ProtoReflect.Descriptor instead.
func (*VirshCreateRequest) Descriptor() ([]byte, []int) {
	return file_virsh_proto_rawDescGZIP(), []int{2}
}

func (x *VirshCreateRequest) GetXml() string {
	if x != nil {
		return x.Xml
	}
	return ""
}

type VirshReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VirshReply) Reset() {
	*x = VirshReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_virsh_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirshReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirshReply) ProtoMessage() {}

func (x *VirshReply) ProtoReflect() protoreflect.Message {
	mi := &file_virsh_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirshReply.ProtoReflect.Descriptor instead.
func (*VirshReply) Descriptor() ([]byte, []int) {
	return file_virsh_proto_rawDescGZIP(), []int{3}
}

func (x *VirshReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VirshReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_virsh_proto protoreflect.FileDescriptor

var file_virsh_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76,
	0x69, 0x72, 0x73, 0x68, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0c, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x56, 0x69,
	0x72, 0x73, 0x68, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x6d, 0x6f, 0x76,
	0x65, 0x22, 0x26, 0x0a, 0x12, 0x56, 0x69, 0x72, 0x73, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x78, 0x6d, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x6d, 0x6c, 0x22, 0x3a, 0x0a, 0x0a, 0x56, 0x69, 0x72,
	0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x32, 0xdb, 0x07, 0x0a, 0x05, 0x56, 0x69, 0x72, 0x73, 0x68, 0x12,
	0x5c, 0x0a, 0x0c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x13, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72,
	0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x12,
	0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x5b, 0x0a,
	0x0d, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19,
	0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73,
	0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69,
	0x72, 0x73, 0x68, 0x2f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x5e, 0x0a, 0x0d, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x76, 0x69,
	0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x66, 0x0a, 0x11, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x6f, 0x66, 0x74, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x12,
	0x13, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72,
	0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22,
	0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x53,
	0x6f, 0x66, 0x74, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d,
	0x65, 0x7d, 0x12, 0x66, 0x0a, 0x11, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x48, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x12, 0x13, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e,
	0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76,
	0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x22, 0x21, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x48, 0x61, 0x72, 0x64, 0x52, 0x65, 0x62, 0x6f, 0x6f,
	0x74, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x62, 0x0a, 0x0f, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x13, 0x2e,
	0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x22, 0x1f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x53, 0x68, 0x75,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x5c,
	0x0a, 0x0c, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x13,
	0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73,
	0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x1c,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x5c, 0x0a, 0x0c,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x50, 0x61, 0x75, 0x73, 0x65, 0x12, 0x13, 0x2e, 0x76,
	0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x22, 0x1c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x50, 0x61, 0x75, 0x73,
	0x65, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x5e, 0x0a, 0x0d, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x76, 0x69,
	0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72, 0x73, 0x68, 0x2f, 0x52, 0x65, 0x73, 0x75, 0x6d,
	0x65, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x12, 0x67, 0x0a, 0x0e, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x76,
	0x69, 0x72, 0x73, 0x68, 0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x65, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x1a, 0x11, 0x2e, 0x76, 0x69, 0x72, 0x73, 0x68,
	0x2e, 0x56, 0x69, 0x72, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x20, 0x22, 0x1e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x56, 0x69, 0x72,
	0x73, 0x68, 0x2f, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x2f, 0x7b, 0x76, 0x6d, 0x6e, 0x61,
	0x6d, 0x65, 0x7d, 0x42, 0x1e, 0x5a, 0x16, 0x69, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x2f, 0x63, 0x6f, 0x62, 0x72, 0x61, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0xa2, 0x02, 0x03,
	0x48, 0x4c, 0x57, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_virsh_proto_rawDescOnce sync.Once
	file_virsh_proto_rawDescData = file_virsh_proto_rawDesc
)

func file_virsh_proto_rawDescGZIP() []byte {
	file_virsh_proto_rawDescOnce.Do(func() {
		file_virsh_proto_rawDescData = protoimpl.X.CompressGZIP(file_virsh_proto_rawDescData)
	})
	return file_virsh_proto_rawDescData
}

var file_virsh_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_virsh_proto_goTypes = []interface{}{
	(*VirshRequest)(nil),        // 0: virsh.VirshRequest
	(*VirshMachineMigrate)(nil), // 1: virsh.VirshMachineMigrate
	(*VirshCreateRequest)(nil),  // 2: virsh.VirshCreateRequest
	(*VirshReply)(nil),          // 3: virsh.VirshReply
}
var file_virsh_proto_depIdxs = []int32{
	0,  // 0: virsh.Virsh.MachineState:input_type -> virsh.VirshRequest
	2,  // 1: virsh.Virsh.MachineCreate:input_type -> virsh.VirshCreateRequest
	0,  // 2: virsh.Virsh.MachineDelete:input_type -> virsh.VirshRequest
	0,  // 3: virsh.Virsh.MachineSoftReboot:input_type -> virsh.VirshRequest
	0,  // 4: virsh.Virsh.MachineHardReboot:input_type -> virsh.VirshRequest
	0,  // 5: virsh.Virsh.MachineShutdown:input_type -> virsh.VirshRequest
	0,  // 6: virsh.Virsh.MachineStart:input_type -> virsh.VirshRequest
	0,  // 7: virsh.Virsh.MachinePause:input_type -> virsh.VirshRequest
	0,  // 8: virsh.Virsh.MachineResume:input_type -> virsh.VirshRequest
	1,  // 9: virsh.Virsh.MachineMigrate:input_type -> virsh.VirshMachineMigrate
	3,  // 10: virsh.Virsh.MachineState:output_type -> virsh.VirshReply
	3,  // 11: virsh.Virsh.MachineCreate:output_type -> virsh.VirshReply
	3,  // 12: virsh.Virsh.MachineDelete:output_type -> virsh.VirshReply
	3,  // 13: virsh.Virsh.MachineSoftReboot:output_type -> virsh.VirshReply
	3,  // 14: virsh.Virsh.MachineHardReboot:output_type -> virsh.VirshReply
	3,  // 15: virsh.Virsh.MachineShutdown:output_type -> virsh.VirshReply
	3,  // 16: virsh.Virsh.MachineStart:output_type -> virsh.VirshReply
	3,  // 17: virsh.Virsh.MachinePause:output_type -> virsh.VirshReply
	3,  // 18: virsh.Virsh.MachineResume:output_type -> virsh.VirshReply
	3,  // 19: virsh.Virsh.MachineMigrate:output_type -> virsh.VirshReply
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_virsh_proto_init() }
func file_virsh_proto_init() {
	if File_virsh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_virsh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirshRequest); i {
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
		file_virsh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirshMachineMigrate); i {
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
		file_virsh_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirshCreateRequest); i {
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
		file_virsh_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirshReply); i {
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
			RawDescriptor: file_virsh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_virsh_proto_goTypes,
		DependencyIndexes: file_virsh_proto_depIdxs,
		MessageInfos:      file_virsh_proto_msgTypes,
	}.Build()
	File_virsh_proto = out.File
	file_virsh_proto_rawDesc = nil
	file_virsh_proto_goTypes = nil
	file_virsh_proto_depIdxs = nil
}