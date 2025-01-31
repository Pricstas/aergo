// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.21.4
// source: aergo_raft.proto

package types

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

// cluster member for raft consensus
type MembershipChangeType int32

const (
	MembershipChangeType_ADD_MEMBER    MembershipChangeType = 0
	MembershipChangeType_REMOVE_MEMBER MembershipChangeType = 1
)

// Enum value maps for MembershipChangeType.
var (
	MembershipChangeType_name = map[int32]string{
		0: "ADD_MEMBER",
		1: "REMOVE_MEMBER",
	}
	MembershipChangeType_value = map[string]int32{
		"ADD_MEMBER":    0,
		"REMOVE_MEMBER": 1,
	}
)

func (x MembershipChangeType) Enum() *MembershipChangeType {
	p := new(MembershipChangeType)
	*p = x
	return p
}

func (x MembershipChangeType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MembershipChangeType) Descriptor() protoreflect.EnumDescriptor {
	return file_aergo_raft_proto_enumTypes[0].Descriptor()
}

func (MembershipChangeType) Type() protoreflect.EnumType {
	return &file_aergo_raft_proto_enumTypes[0]
}

func (x MembershipChangeType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MembershipChangeType.Descriptor instead.
func (MembershipChangeType) EnumDescriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{0}
}

type ConfChangeState int32

const (
	ConfChangeState_CONF_CHANGE_STATE_PROPOSED ConfChangeState = 0
	ConfChangeState_CONF_CHANGE_STATE_SAVED    ConfChangeState = 1
	ConfChangeState_CONF_CHANGE_STATE_APPLIED  ConfChangeState = 2
)

// Enum value maps for ConfChangeState.
var (
	ConfChangeState_name = map[int32]string{
		0: "CONF_CHANGE_STATE_PROPOSED",
		1: "CONF_CHANGE_STATE_SAVED",
		2: "CONF_CHANGE_STATE_APPLIED",
	}
	ConfChangeState_value = map[string]int32{
		"CONF_CHANGE_STATE_PROPOSED": 0,
		"CONF_CHANGE_STATE_SAVED":    1,
		"CONF_CHANGE_STATE_APPLIED":  2,
	}
)

func (x ConfChangeState) Enum() *ConfChangeState {
	p := new(ConfChangeState)
	*p = x
	return p
}

func (x ConfChangeState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfChangeState) Descriptor() protoreflect.EnumDescriptor {
	return file_aergo_raft_proto_enumTypes[1].Descriptor()
}

func (ConfChangeState) Type() protoreflect.EnumType {
	return &file_aergo_raft_proto_enumTypes[1]
}

func (x ConfChangeState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfChangeState.Descriptor instead.
func (ConfChangeState) EnumDescriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{1}
}

type MemberAttr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID      uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PeerID  []byte `protobuf:"bytes,4,opt,name=peerID,proto3" json:"peerID,omitempty"`
}

func (x *MemberAttr) Reset() {
	*x = MemberAttr{}
	mi := &file_aergo_raft_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberAttr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberAttr) ProtoMessage() {}

func (x *MemberAttr) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberAttr.ProtoReflect.Descriptor instead.
func (*MemberAttr) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{0}
}

func (x *MemberAttr) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *MemberAttr) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MemberAttr) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *MemberAttr) GetPeerID() []byte {
	if x != nil {
		return x.PeerID
	}
	return nil
}

type MembershipChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      MembershipChangeType `protobuf:"varint,1,opt,name=type,proto3,enum=types.MembershipChangeType" json:"type,omitempty"`
	RequestID uint64               `protobuf:"varint,2,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Attr      *MemberAttr          `protobuf:"bytes,3,opt,name=attr,proto3" json:"attr,omitempty"`
}

func (x *MembershipChange) Reset() {
	*x = MembershipChange{}
	mi := &file_aergo_raft_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MembershipChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipChange) ProtoMessage() {}

func (x *MembershipChange) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipChange.ProtoReflect.Descriptor instead.
func (*MembershipChange) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{1}
}

func (x *MembershipChange) GetType() MembershipChangeType {
	if x != nil {
		return x.Type
	}
	return MembershipChangeType_ADD_MEMBER
}

func (x *MembershipChange) GetRequestID() uint64 {
	if x != nil {
		return x.RequestID
	}
	return 0
}

func (x *MembershipChange) GetAttr() *MemberAttr {
	if x != nil {
		return x.Attr
	}
	return nil
}

type MembershipChangeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Attr *MemberAttr `protobuf:"bytes,1,opt,name=attr,proto3" json:"attr,omitempty"`
}

func (x *MembershipChangeReply) Reset() {
	*x = MembershipChangeReply{}
	mi := &file_aergo_raft_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MembershipChangeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipChangeReply) ProtoMessage() {}

func (x *MembershipChangeReply) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipChangeReply.ProtoReflect.Descriptor instead.
func (*MembershipChangeReply) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{2}
}

func (x *MembershipChangeReply) GetAttr() *MemberAttr {
	if x != nil {
		return x.Attr
	}
	return nil
}

type HardStateInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Term   uint64 `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Commit uint64 `protobuf:"varint,2,opt,name=commit,proto3" json:"commit,omitempty"`
}

func (x *HardStateInfo) Reset() {
	*x = HardStateInfo{}
	mi := &file_aergo_raft_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HardStateInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HardStateInfo) ProtoMessage() {}

func (x *HardStateInfo) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HardStateInfo.ProtoReflect.Descriptor instead.
func (*HardStateInfo) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{3}
}

func (x *HardStateInfo) GetTerm() uint64 {
	if x != nil {
		return x.Term
	}
	return 0
}

func (x *HardStateInfo) GetCommit() uint64 {
	if x != nil {
		return x.Commit
	}
	return 0
}

// data types for raft support
// GetClusterInfoRequest
type GetClusterInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BestBlockHash []byte `protobuf:"bytes,1,opt,name=bestBlockHash,proto3" json:"bestBlockHash,omitempty"`
}

func (x *GetClusterInfoRequest) Reset() {
	*x = GetClusterInfoRequest{}
	mi := &file_aergo_raft_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClusterInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterInfoRequest) ProtoMessage() {}

func (x *GetClusterInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterInfoRequest.ProtoReflect.Descriptor instead.
func (*GetClusterInfoRequest) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{4}
}

func (x *GetClusterInfoRequest) GetBestBlockHash() []byte {
	if x != nil {
		return x.BestBlockHash
	}
	return nil
}

type GetClusterInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChainID       []byte         `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	ClusterID     uint64         `protobuf:"varint,2,opt,name=clusterID,proto3" json:"clusterID,omitempty"`
	Error         string         `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	MbrAttrs      []*MemberAttr  `protobuf:"bytes,4,rep,name=mbrAttrs,proto3" json:"mbrAttrs,omitempty"`
	BestBlockNo   uint64         `protobuf:"varint,5,opt,name=bestBlockNo,proto3" json:"bestBlockNo,omitempty"`    // best block number of this node
	HardStateInfo *HardStateInfo `protobuf:"bytes,6,opt,name=hardStateInfo,proto3" json:"hardStateInfo,omitempty"` // hardstate corresponding to bestblockhash of requester
}

func (x *GetClusterInfoResponse) Reset() {
	*x = GetClusterInfoResponse{}
	mi := &file_aergo_raft_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClusterInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClusterInfoResponse) ProtoMessage() {}

func (x *GetClusterInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClusterInfoResponse.ProtoReflect.Descriptor instead.
func (*GetClusterInfoResponse) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{5}
}

func (x *GetClusterInfoResponse) GetChainID() []byte {
	if x != nil {
		return x.ChainID
	}
	return nil
}

func (x *GetClusterInfoResponse) GetClusterID() uint64 {
	if x != nil {
		return x.ClusterID
	}
	return 0
}

func (x *GetClusterInfoResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *GetClusterInfoResponse) GetMbrAttrs() []*MemberAttr {
	if x != nil {
		return x.MbrAttrs
	}
	return nil
}

func (x *GetClusterInfoResponse) GetBestBlockNo() uint64 {
	if x != nil {
		return x.BestBlockNo
	}
	return 0
}

func (x *GetClusterInfoResponse) GetHardStateInfo() *HardStateInfo {
	if x != nil {
		return x.HardStateInfo
	}
	return nil
}

type ConfChangeProgress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State   ConfChangeState `protobuf:"varint,1,opt,name=State,proto3,enum=types.ConfChangeState" json:"State,omitempty"`
	Err     string          `protobuf:"bytes,2,opt,name=Err,proto3" json:"Err,omitempty"`
	Members []*MemberAttr   `protobuf:"bytes,3,rep,name=Members,proto3" json:"Members,omitempty"`
}

func (x *ConfChangeProgress) Reset() {
	*x = ConfChangeProgress{}
	mi := &file_aergo_raft_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConfChangeProgress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfChangeProgress) ProtoMessage() {}

func (x *ConfChangeProgress) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfChangeProgress.ProtoReflect.Descriptor instead.
func (*ConfChangeProgress) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{6}
}

func (x *ConfChangeProgress) GetState() ConfChangeState {
	if x != nil {
		return x.State
	}
	return ConfChangeState_CONF_CHANGE_STATE_PROPOSED
}

func (x *ConfChangeProgress) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

func (x *ConfChangeProgress) GetMembers() []*MemberAttr {
	if x != nil {
		return x.Members
	}
	return nil
}

// SnapshotResponse is response message of receiving peer
type SnapshotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  ResultStatus `protobuf:"varint,1,opt,name=status,proto3,enum=types.ResultStatus" json:"status,omitempty"`
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SnapshotResponse) Reset() {
	*x = SnapshotResponse{}
	mi := &file_aergo_raft_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnapshotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotResponse) ProtoMessage() {}

func (x *SnapshotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aergo_raft_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotResponse.ProtoReflect.Descriptor instead.
func (*SnapshotResponse) Descriptor() ([]byte, []int) {
	return file_aergo_raft_proto_rawDescGZIP(), []int{7}
}

func (x *SnapshotResponse) GetStatus() ResultStatus {
	if x != nil {
		return x.Status
	}
	return ResultStatus_OK
}

func (x *SnapshotResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_aergo_raft_proto protoreflect.FileDescriptor

var file_aergo_raft_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x65, 0x72, 0x67, 0x6f, 0x5f, 0x72, 0x61, 0x66, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x09, 0x70, 0x32, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x74,
	0x74, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2f, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x52, 0x04, 0x61,
	0x74, 0x74, 0x72, 0x22, 0x3e, 0x0a, 0x15, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69,
	0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x25, 0x0a, 0x04,
	0x61, 0x74, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x52, 0x04, 0x61,
	0x74, 0x74, 0x72, 0x22, 0x3b, 0x0a, 0x0d, 0x48, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x04, 0x74, 0x65, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x22, 0x3d, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x65, 0x73,
	0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x62, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22,
	0xf3, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x62, 0x72, 0x41,
	0x74, 0x74, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72, 0x52, 0x08, 0x6d,
	0x62, 0x72, 0x41, 0x74, 0x74, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x65, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x62, 0x65,
	0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x6f, 0x12, 0x3a, 0x0a, 0x0d, 0x68, 0x61, 0x72,
	0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2c, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x45, 0x72, 0x72, 0x12, 0x2b, 0x0a, 0x07,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x74, 0x74, 0x72,
	0x52, 0x07, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x59, 0x0a, 0x10, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2a, 0x39, 0x0a, 0x14, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a,
	0x41, 0x44, 0x44, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d,
	0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x01, 0x2a,
	0x6d, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4e, 0x46, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x4f, 0x53, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4e, 0x46, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47,
	0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x41, 0x56, 0x45, 0x44, 0x10, 0x01, 0x12,
	0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4e, 0x46, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x50, 0x50, 0x4c, 0x49, 0x45, 0x44, 0x10, 0x02, 0x42, 0x08,
	0x5a, 0x06, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aergo_raft_proto_rawDescOnce sync.Once
	file_aergo_raft_proto_rawDescData = file_aergo_raft_proto_rawDesc
)

func file_aergo_raft_proto_rawDescGZIP() []byte {
	file_aergo_raft_proto_rawDescOnce.Do(func() {
		file_aergo_raft_proto_rawDescData = protoimpl.X.CompressGZIP(file_aergo_raft_proto_rawDescData)
	})
	return file_aergo_raft_proto_rawDescData
}

var file_aergo_raft_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_aergo_raft_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_aergo_raft_proto_goTypes = []any{
	(MembershipChangeType)(0),      // 0: types.MembershipChangeType
	(ConfChangeState)(0),           // 1: types.ConfChangeState
	(*MemberAttr)(nil),             // 2: types.MemberAttr
	(*MembershipChange)(nil),       // 3: types.MembershipChange
	(*MembershipChangeReply)(nil),  // 4: types.MembershipChangeReply
	(*HardStateInfo)(nil),          // 5: types.HardStateInfo
	(*GetClusterInfoRequest)(nil),  // 6: types.GetClusterInfoRequest
	(*GetClusterInfoResponse)(nil), // 7: types.GetClusterInfoResponse
	(*ConfChangeProgress)(nil),     // 8: types.ConfChangeProgress
	(*SnapshotResponse)(nil),       // 9: types.SnapshotResponse
	(ResultStatus)(0),              // 10: types.ResultStatus
}
var file_aergo_raft_proto_depIdxs = []int32{
	0,  // 0: types.MembershipChange.type:type_name -> types.MembershipChangeType
	2,  // 1: types.MembershipChange.attr:type_name -> types.MemberAttr
	2,  // 2: types.MembershipChangeReply.attr:type_name -> types.MemberAttr
	2,  // 3: types.GetClusterInfoResponse.mbrAttrs:type_name -> types.MemberAttr
	5,  // 4: types.GetClusterInfoResponse.hardStateInfo:type_name -> types.HardStateInfo
	1,  // 5: types.ConfChangeProgress.State:type_name -> types.ConfChangeState
	2,  // 6: types.ConfChangeProgress.Members:type_name -> types.MemberAttr
	10, // 7: types.SnapshotResponse.status:type_name -> types.ResultStatus
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_aergo_raft_proto_init() }
func file_aergo_raft_proto_init() {
	if File_aergo_raft_proto != nil {
		return
	}
	file_p2p_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aergo_raft_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aergo_raft_proto_goTypes,
		DependencyIndexes: file_aergo_raft_proto_depIdxs,
		EnumInfos:         file_aergo_raft_proto_enumTypes,
		MessageInfos:      file_aergo_raft_proto_msgTypes,
	}.Build()
	File_aergo_raft_proto = out.File
	file_aergo_raft_proto_rawDesc = nil
	file_aergo_raft_proto_goTypes = nil
	file_aergo_raft_proto_depIdxs = nil
}
