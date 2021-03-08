// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.3
// source: relay_proto/x/relay/types/shared.proto

package types

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Origin int32

const (
	Origin_LOCAL  Origin = 0
	Origin_REMOTE Origin = 1
)

// Enum value maps for Origin.
var (
	Origin_name = map[int32]string{
		0: "LOCAL",
		1: "REMOTE",
	}
	Origin_value = map[string]int32{
		"LOCAL":  0,
		"REMOTE": 1,
	}
)

func (x Origin) Enum() *Origin {
	p := new(Origin)
	*p = x
	return p
}

func (x Origin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Origin) Descriptor() protoreflect.EnumDescriptor {
	return file_relay_proto_x_relay_types_shared_proto_enumTypes[0].Descriptor()
}

func (Origin) Type() protoreflect.EnumType {
	return &file_relay_proto_x_relay_types_shared_proto_enumTypes[0]
}

func (x Origin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Origin.Descriptor instead.
func (Origin) EnumDescriptor() ([]byte, []int) {
	return file_relay_proto_x_relay_types_shared_proto_rawDescGZIP(), []int{0}
}

type BitcoinHeader struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Raw        []byte `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	Hash       []byte `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Height     uint32 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	PrevHash   []byte `protobuf:"bytes,4,opt,name=prev_hash,json=prevHash,proto3" json:"prev_hash,omitempty"`
	MerkleRoot []byte `protobuf:"bytes,5,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
}

func (x *BitcoinHeader) Reset() {
	*x = BitcoinHeader{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BitcoinHeader) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BitcoinHeader) ProtoMessage() {}

func (x *BitcoinHeader) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BitcoinHeader.ProtoReflect.Descriptor instead.
func (*BitcoinHeader) Descriptor() ([]byte, []int) {
	return file_relay_proto_x_relay_types_shared_proto_rawDescGZIP(), []int{0}
}

func (x *BitcoinHeader) GetRaw() []byte {
	if x != nil {
		return x.Raw
	}
	return nil
}

func (x *BitcoinHeader) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *BitcoinHeader) GetHeight() uint32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *BitcoinHeader) GetPrevHash() []byte {
	if x != nil {
		return x.PrevHash
	}
	return nil
}

func (x *BitcoinHeader) GetMerkleRoot() []byte {
	if x != nil {
		return x.MerkleRoot
	}
	return nil
}

type SPVProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version           []byte         `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Vin               []byte         `protobuf:"bytes,2,opt,name=vin,proto3" json:"vin,omitempty"`
	Vout              []byte         `protobuf:"bytes,3,opt,name=vout,proto3" json:"vout,omitempty"`
	Locktime          []byte         `protobuf:"bytes,4,opt,name=locktime,proto3" json:"locktime,omitempty"`
	TxID              []byte         `protobuf:"bytes,5,opt,name=tx_i_d,json=txID,proto3" json:"tx_i_d,omitempty"`
	Index             uint32         `protobuf:"varint,6,opt,name=index,proto3" json:"index,omitempty"`
	ConfirmingHeader  *BitcoinHeader `protobuf:"bytes,7,opt,name=confirming_header,json=confirmingHeader,proto3" json:"confirming_header,omitempty"`
	IntermediateNodes []byte         `protobuf:"bytes,8,opt,name=intermediate_nodes,json=intermediateNodes,proto3" json:"intermediate_nodes,omitempty"`
}

func (x *SPVProof) Reset() {
	*x = SPVProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SPVProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SPVProof) ProtoMessage() {}

func (x *SPVProof) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SPVProof.ProtoReflect.Descriptor instead.
func (*SPVProof) Descriptor() ([]byte, []int) {
	return file_relay_proto_x_relay_types_shared_proto_rawDescGZIP(), []int{1}
}

func (x *SPVProof) GetVersion() []byte {
	if x != nil {
		return x.Version
	}
	return nil
}

func (x *SPVProof) GetVin() []byte {
	if x != nil {
		return x.Vin
	}
	return nil
}

func (x *SPVProof) GetVout() []byte {
	if x != nil {
		return x.Vout
	}
	return nil
}

func (x *SPVProof) GetLocktime() []byte {
	if x != nil {
		return x.Locktime
	}
	return nil
}

func (x *SPVProof) GetTxID() []byte {
	if x != nil {
		return x.TxID
	}
	return nil
}

func (x *SPVProof) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *SPVProof) GetConfirmingHeader() *BitcoinHeader {
	if x != nil {
		return x.ConfirmingHeader
	}
	return nil
}

func (x *SPVProof) GetIntermediateNodes() []byte {
	if x != nil {
		return x.IntermediateNodes
	}
	return nil
}

type FilledRequestInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InputIndex  uint32 `protobuf:"varint,1,opt,name=input_index,json=inputIndex,proto3" json:"input_index,omitempty"`
	OutputIndex uint32 `protobuf:"varint,2,opt,name=output_index,json=outputIndex,proto3" json:"output_index,omitempty"`
	ID          []byte `protobuf:"bytes,3,opt,name=i_d,json=iD,proto3" json:"i_d,omitempty"`
}

func (x *FilledRequestInfo) Reset() {
	*x = FilledRequestInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilledRequestInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilledRequestInfo) ProtoMessage() {}

func (x *FilledRequestInfo) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilledRequestInfo.ProtoReflect.Descriptor instead.
func (*FilledRequestInfo) Descriptor() ([]byte, []int) {
	return file_relay_proto_x_relay_types_shared_proto_rawDescGZIP(), []int{2}
}

func (x *FilledRequestInfo) GetInputIndex() uint32 {
	if x != nil {
		return x.InputIndex
	}
	return 0
}

func (x *FilledRequestInfo) GetOutputIndex() uint32 {
	if x != nil {
		return x.OutputIndex
	}
	return 0
}

func (x *FilledRequestInfo) GetID() []byte {
	if x != nil {
		return x.ID
	}
	return nil
}

type ProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spends     []byte `protobuf:"bytes,1,opt,name=spends,proto3" json:"spends,omitempty"`
	Pays       []byte `protobuf:"bytes,2,opt,name=pays,proto3" json:"pays,omitempty"`
	PaysValue  uint64 `protobuf:"varint,3,opt,name=pays_value,json=paysValue,proto3" json:"pays_value,omitempty"`
	AciveState bool   `protobuf:"varint,4,opt,name=acive_state,json=aciveState,proto3" json:"acive_state,omitempty"`
	NumConfs   uint32 `protobuf:"varint,5,opt,name=num_confs,json=numConfs,proto3" json:"num_confs,omitempty"`
	Origin     Origin `protobuf:"varint,6,opt,name=origin,proto3,enum=relay.Origin" json:"origin,omitempty"`
	Action     []byte `protobuf:"bytes,7,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *ProofRequest) Reset() {
	*x = ProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofRequest) ProtoMessage() {}

func (x *ProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofRequest.ProtoReflect.Descriptor instead.
func (*ProofRequest) Descriptor() ([]byte, []int) {
	return file_relay_proto_x_relay_types_shared_proto_rawDescGZIP(), []int{3}
}

func (x *ProofRequest) GetSpends() []byte {
	if x != nil {
		return x.Spends
	}
	return nil
}

func (x *ProofRequest) GetPays() []byte {
	if x != nil {
		return x.Pays
	}
	return nil
}

func (x *ProofRequest) GetPaysValue() uint64 {
	if x != nil {
		return x.PaysValue
	}
	return 0
}

func (x *ProofRequest) GetAciveState() bool {
	if x != nil {
		return x.AciveState
	}
	return false
}

func (x *ProofRequest) GetNumConfs() uint32 {
	if x != nil {
		return x.NumConfs
	}
	return 0
}

func (x *ProofRequest) GetOrigin() Origin {
	if x != nil {
		return x.Origin
	}
	return Origin_LOCAL
}

func (x *ProofRequest) GetAction() []byte {
	if x != nil {
		return x.Action
	}
	return nil
}

type FilledRequests struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proof  *SPVProof            `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
	Filled []*FilledRequestInfo `protobuf:"bytes,2,rep,name=filled,proto3" json:"filled,omitempty"`
}

func (x *FilledRequests) Reset() {
	*x = FilledRequests{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilledRequests) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilledRequests) ProtoMessage() {}

func (x *FilledRequests) ProtoReflect() protoreflect.Message {
	mi := &file_relay_proto_x_relay_types_shared_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilledRequests.ProtoReflect.Descriptor instead.
func (*FilledRequests) Descriptor() ([]byte, []int) {
	return file_relay_proto_x_relay_types_shared_proto_rawDescGZIP(), []int{4}
}

func (x *FilledRequests) GetProof() *SPVProof {
	if x != nil {
		return x.Proof
	}
	return nil
}

func (x *FilledRequests) GetFilled() []*FilledRequestInfo {
	if x != nil {
		return x.Filled
	}
	return nil
}

var File_relay_proto_x_relay_types_shared_proto protoreflect.FileDescriptor

var file_relay_proto_x_relay_types_shared_proto_rawDesc = []byte{
	0x0a, 0x26, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x78, 0x2f,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x22,
	0x8b, 0x01, 0x0a, 0x0d, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03,
	0x72, 0x61, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x76, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x08, 0x70, 0x72, 0x65, 0x76, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0a, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x22, 0x84, 0x02,
	0x0a, 0x08, 0x53, 0x50, 0x56, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x03, 0x76, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x76, 0x6f, 0x75, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6c, 0x6f,
	0x63, 0x6b, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x06, 0x74, 0x78, 0x5f, 0x69, 0x5f, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x78, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x41, 0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x69, 0x6e, 0x67,
	0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x42, 0x69, 0x74, 0x63, 0x6f, 0x69, 0x6e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x69, 0x6e, 0x67, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x12, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65,
	0x64, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x11, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x73, 0x22, 0x68, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x0f, 0x0a,
	0x03, 0x69, 0x5f, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x44, 0x22, 0xd6,
	0x01, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x79, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x61, 0x79, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x79, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x70, 0x61, 0x79, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63,
	0x69, 0x76, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x61, 0x63, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x75, 0x6d, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x73, 0x12, 0x25, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x0e, 0x46, 0x69, 0x6c, 0x6c, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x6f, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x53, 0x50, 0x56, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x6f, 0x66,
	0x12, 0x30, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x6c,
	0x65, 0x64, 0x2a, 0x1f, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x09, 0x0a, 0x05,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x4d, 0x4f, 0x54,
	0x45, 0x10, 0x01, 0x42, 0x0f, 0x5a, 0x0d, 0x78, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_proto_x_relay_types_shared_proto_rawDescOnce sync.Once
	file_relay_proto_x_relay_types_shared_proto_rawDescData = file_relay_proto_x_relay_types_shared_proto_rawDesc
)

func file_relay_proto_x_relay_types_shared_proto_rawDescGZIP() []byte {
	file_relay_proto_x_relay_types_shared_proto_rawDescOnce.Do(func() {
		file_relay_proto_x_relay_types_shared_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_proto_x_relay_types_shared_proto_rawDescData)
	})
	return file_relay_proto_x_relay_types_shared_proto_rawDescData
}

var file_relay_proto_x_relay_types_shared_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relay_proto_x_relay_types_shared_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_relay_proto_x_relay_types_shared_proto_goTypes = []interface{}{
	(Origin)(0),               // 0: relay.Origin
	(*BitcoinHeader)(nil),     // 1: relay.BitcoinHeader
	(*SPVProof)(nil),          // 2: relay.SPVProof
	(*FilledRequestInfo)(nil), // 3: relay.FilledRequestInfo
	(*ProofRequest)(nil),      // 4: relay.ProofRequest
	(*FilledRequests)(nil),    // 5: relay.FilledRequests
}
var file_relay_proto_x_relay_types_shared_proto_depIdxs = []int32{
	1, // 0: relay.SPVProof.confirming_header:type_name -> relay.BitcoinHeader
	0, // 1: relay.ProofRequest.origin:type_name -> relay.Origin
	2, // 2: relay.FilledRequests.proof:type_name -> relay.SPVProof
	3, // 3: relay.FilledRequests.filled:type_name -> relay.FilledRequestInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_relay_proto_x_relay_types_shared_proto_init() }
func file_relay_proto_x_relay_types_shared_proto_init() {
	if File_relay_proto_x_relay_types_shared_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_proto_x_relay_types_shared_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BitcoinHeader); i {
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
		file_relay_proto_x_relay_types_shared_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SPVProof); i {
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
		file_relay_proto_x_relay_types_shared_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilledRequestInfo); i {
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
		file_relay_proto_x_relay_types_shared_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofRequest); i {
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
		file_relay_proto_x_relay_types_shared_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilledRequests); i {
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
			RawDescriptor: file_relay_proto_x_relay_types_shared_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_relay_proto_x_relay_types_shared_proto_goTypes,
		DependencyIndexes: file_relay_proto_x_relay_types_shared_proto_depIdxs,
		EnumInfos:         file_relay_proto_x_relay_types_shared_proto_enumTypes,
		MessageInfos:      file_relay_proto_x_relay_types_shared_proto_msgTypes,
	}.Build()
	File_relay_proto_x_relay_types_shared_proto = out.File
	file_relay_proto_x_relay_types_shared_proto_rawDesc = nil
	file_relay_proto_x_relay_types_shared_proto_goTypes = nil
	file_relay_proto_x_relay_types_shared_proto_depIdxs = nil
}