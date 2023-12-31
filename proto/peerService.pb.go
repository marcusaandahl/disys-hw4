// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.4
// source: proto/peerService.proto

package proto

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

type VoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterId int32 `protobuf:"varint,1,opt,name=requesterId,proto3" json:"requesterId,omitempty"`
	DeniedBy    int32 `protobuf:"varint,2,opt,name=deniedBy,proto3" json:"deniedBy,omitempty"`
	Denied      bool  `protobuf:"varint,3,opt,name=denied,proto3" json:"denied,omitempty"`
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peerService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peerService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_proto_peerService_proto_rawDescGZIP(), []int{0}
}

func (x *VoteRequest) GetRequesterId() int32 {
	if x != nil {
		return x.RequesterId
	}
	return 0
}

func (x *VoteRequest) GetDeniedBy() int32 {
	if x != nil {
		return x.DeniedBy
	}
	return 0
}

func (x *VoteRequest) GetDenied() bool {
	if x != nil {
		return x.Denied
	}
	return false
}

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPeerPort int32 `protobuf:"varint,2,opt,name=newPeerPort,proto3" json:"newPeerPort,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peerService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peerService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_proto_peerService_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectRequest) GetNewPeerPort() int32 {
	if x != nil {
		return x.NewPeerPort
	}
	return 0
}

type ConnectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewPeerNextPeerId   int32 `protobuf:"varint,1,opt,name=newPeerNextPeerId,proto3" json:"newPeerNextPeerId,omitempty"`
	NewPeerNextPeerPort int32 `protobuf:"varint,2,opt,name=newPeerNextPeerPort,proto3" json:"newPeerNextPeerPort,omitempty"`
	NewPeerId           int32 `protobuf:"varint,3,opt,name=newPeerId,proto3" json:"newPeerId,omitempty"`
}

func (x *ConnectResponse) Reset() {
	*x = ConnectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peerService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectResponse) ProtoMessage() {}

func (x *ConnectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peerService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectResponse.ProtoReflect.Descriptor instead.
func (*ConnectResponse) Descriptor() ([]byte, []int) {
	return file_proto_peerService_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectResponse) GetNewPeerNextPeerId() int32 {
	if x != nil {
		return x.NewPeerNextPeerId
	}
	return 0
}

func (x *ConnectResponse) GetNewPeerNextPeerPort() int32 {
	if x != nil {
		return x.NewPeerNextPeerPort
	}
	return 0
}

func (x *ConnectResponse) GetNewPeerId() int32 {
	if x != nil {
		return x.NewPeerId
	}
	return 0
}

type LeaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraitorId       int32 `protobuf:"varint,1,opt,name=traitorId,proto3" json:"traitorId,omitempty"`
	TraitorNextId   int32 `protobuf:"varint,2,opt,name=traitorNextId,proto3" json:"traitorNextId,omitempty"`
	TraitorNextPort int32 `protobuf:"varint,3,opt,name=traitorNextPort,proto3" json:"traitorNextPort,omitempty"`
}

func (x *LeaveRequest) Reset() {
	*x = LeaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_peerService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeaveRequest) ProtoMessage() {}

func (x *LeaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_peerService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeaveRequest.ProtoReflect.Descriptor instead.
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return file_proto_peerService_proto_rawDescGZIP(), []int{3}
}

func (x *LeaveRequest) GetTraitorId() int32 {
	if x != nil {
		return x.TraitorId
	}
	return 0
}

func (x *LeaveRequest) GetTraitorNextId() int32 {
	if x != nil {
		return x.TraitorNextId
	}
	return 0
}

func (x *LeaveRequest) GetTraitorNextPort() int32 {
	if x != nil {
		return x.TraitorNextPort
	}
	return 0
}

var File_proto_peerService_proto protoreflect.FileDescriptor

var file_proto_peerService_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a,
	0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x42, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x64, 0x65, 0x6e, 0x69, 0x65, 0x64, 0x42, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65,
	0x6e, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x64, 0x65, 0x6e, 0x69,
	0x65, 0x64, 0x22, 0x32, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x65, 0x72, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x50, 0x65,
	0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x6e, 0x65,
	0x77, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x6e, 0x65, 0x77, 0x50,
	0x65, 0x65, 0x72, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x65, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x13, 0x6e, 0x65, 0x77, 0x50, 0x65, 0x65, 0x72, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x65, 0x65, 0x72, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6e,
	0x65, 0x77, 0x50, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7c, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x69,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x72, 0x61,
	0x69, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x69, 0x74, 0x6f,
	0x72, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x74,
	0x72, 0x61, 0x69, 0x74, 0x6f, 0x72, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f,
	0x74, 0x72, 0x61, 0x69, 0x74, 0x6f, 0x72, 0x4e, 0x65, 0x78, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x69, 0x74, 0x6f, 0x72, 0x4e, 0x65,
	0x78, 0x74, 0x50, 0x6f, 0x72, 0x74, 0x32, 0xc4, 0x01, 0x0a, 0x09, 0x50, 0x65, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x42, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x4c, 0x65, 0x61, 0x76,
	0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4c, 0x65, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x72, 0x63,
	0x75, 0x73, 0x61, 0x61, 0x6e, 0x64, 0x61, 0x68, 0x6c, 0x2f, 0x64, 0x69, 0x73, 0x79, 0x73, 0x2d,
	0x68, 0x77, 0x34, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_peerService_proto_rawDescOnce sync.Once
	file_proto_peerService_proto_rawDescData = file_proto_peerService_proto_rawDesc
)

func file_proto_peerService_proto_rawDescGZIP() []byte {
	file_proto_peerService_proto_rawDescOnce.Do(func() {
		file_proto_peerService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_peerService_proto_rawDescData)
	})
	return file_proto_peerService_proto_rawDescData
}

var file_proto_peerService_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_peerService_proto_goTypes = []interface{}{
	(*VoteRequest)(nil),     // 0: proto.VoteRequest
	(*ConnectRequest)(nil),  // 1: proto.ConnectRequest
	(*ConnectResponse)(nil), // 2: proto.ConnectResponse
	(*LeaveRequest)(nil),    // 3: proto.LeaveRequest
	(*emptypb.Empty)(nil),   // 4: google.protobuf.Empty
}
var file_proto_peerService_proto_depIdxs = []int32{
	1, // 0: proto.PeerProto.CheckConnection:input_type -> proto.ConnectRequest
	3, // 1: proto.PeerProto.LeaveNetwork:input_type -> proto.LeaveRequest
	0, // 2: proto.PeerProto.Vote:input_type -> proto.VoteRequest
	2, // 3: proto.PeerProto.CheckConnection:output_type -> proto.ConnectResponse
	4, // 4: proto.PeerProto.LeaveNetwork:output_type -> google.protobuf.Empty
	4, // 5: proto.PeerProto.Vote:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_peerService_proto_init() }
func file_proto_peerService_proto_init() {
	if File_proto_peerService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_peerService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoteRequest); i {
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
		file_proto_peerService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
		file_proto_peerService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectResponse); i {
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
		file_proto_peerService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeaveRequest); i {
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
			RawDescriptor: file_proto_peerService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_peerService_proto_goTypes,
		DependencyIndexes: file_proto_peerService_proto_depIdxs,
		MessageInfos:      file_proto_peerService_proto_msgTypes,
	}.Build()
	File_proto_peerService_proto = out.File
	file_proto_peerService_proto_rawDesc = nil
	file_proto_peerService_proto_goTypes = nil
	file_proto_peerService_proto_depIdxs = nil
}
