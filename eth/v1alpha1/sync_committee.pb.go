// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.8
// source: eth/v1alpha1/sync_committee.proto

package eth

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	github_com_prysmaticlabs_eth2_types "github.com/prysmaticlabs/eth2-types"
	_ "github.com/prysmaticlabs/ethereumapis/eth/ext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type SyncCommitteeSignature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot           github_com_prysmaticlabs_eth2_types.Slot           `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	BlockRoot      []byte                                             `protobuf:"bytes,2,opt,name=block_root,json=blockRoot,proto3" json:"block_root,omitempty" ssz-size:"32"`
	ValidatorIndex github_com_prysmaticlabs_eth2_types.ValidatorIndex `protobuf:"varint,3,opt,name=validator_index,json=validatorIndex,proto3" json:"validator_index,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.ValidatorIndex"`
	Signature      []byte                                             `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty" ssz-size:"96"`
}

func (x *SyncCommitteeSignature) Reset() {
	*x = SyncCommitteeSignature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCommitteeSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCommitteeSignature) ProtoMessage() {}

func (x *SyncCommitteeSignature) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCommitteeSignature.ProtoReflect.Descriptor instead.
func (*SyncCommitteeSignature) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_sync_committee_proto_rawDescGZIP(), []int{0}
}

func (x *SyncCommitteeSignature) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *SyncCommitteeSignature) GetBlockRoot() []byte {
	if x != nil {
		return x.BlockRoot
	}
	return nil
}

func (x *SyncCommitteeSignature) GetValidatorIndex() github_com_prysmaticlabs_eth2_types.ValidatorIndex {
	if x != nil {
		return x.ValidatorIndex
	}
	return github_com_prysmaticlabs_eth2_types.ValidatorIndex(0)
}

func (x *SyncCommitteeSignature) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type SyncCommitteeContribution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slot              github_com_prysmaticlabs_eth2_types.Slot `protobuf:"varint,1,opt,name=slot,proto3" json:"slot,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.Slot"`
	BlockRoot         []byte                                   `protobuf:"bytes,2,opt,name=block_root,json=blockRoot,proto3" json:"block_root,omitempty" ssz-size:"32"`
	SubcommitteeIndex uint64                                   `protobuf:"varint,3,opt,name=subcommittee_index,json=subcommitteeIndex,proto3" json:"subcommittee_index,omitempty"`
	AggregationBits   []byte                                   `protobuf:"bytes,4,opt,name=aggregation_bits,json=aggregationBits,proto3" json:"aggregation_bits,omitempty" ssz-size:"32"`
	Signature         []byte                                   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty" ssz-size:"96"`
}

func (x *SyncCommitteeContribution) Reset() {
	*x = SyncCommitteeContribution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncCommitteeContribution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncCommitteeContribution) ProtoMessage() {}

func (x *SyncCommitteeContribution) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncCommitteeContribution.ProtoReflect.Descriptor instead.
func (*SyncCommitteeContribution) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_sync_committee_proto_rawDescGZIP(), []int{1}
}

func (x *SyncCommitteeContribution) GetSlot() github_com_prysmaticlabs_eth2_types.Slot {
	if x != nil {
		return x.Slot
	}
	return github_com_prysmaticlabs_eth2_types.Slot(0)
}

func (x *SyncCommitteeContribution) GetBlockRoot() []byte {
	if x != nil {
		return x.BlockRoot
	}
	return nil
}

func (x *SyncCommitteeContribution) GetSubcommitteeIndex() uint64 {
	if x != nil {
		return x.SubcommitteeIndex
	}
	return 0
}

func (x *SyncCommitteeContribution) GetAggregationBits() []byte {
	if x != nil {
		return x.AggregationBits
	}
	return nil
}

func (x *SyncCommitteeContribution) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type ContributionAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AggregatorIndex github_com_prysmaticlabs_eth2_types.ValidatorIndex `protobuf:"varint,1,opt,name=aggregator_index,json=aggregatorIndex,proto3" json:"aggregator_index,omitempty" cast-type:"github.com/prysmaticlabs/eth2-types.ValidatorIndex"`
	Contribution    *SyncCommitteeContribution                         `protobuf:"bytes,2,opt,name=contribution,proto3" json:"contribution,omitempty"`
	SelectionProof  []byte                                             `protobuf:"bytes,3,opt,name=selection_proof,json=selectionProof,proto3" json:"selection_proof,omitempty" ssz-size:"96"`
}

func (x *ContributionAndProof) Reset() {
	*x = ContributionAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContributionAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContributionAndProof) ProtoMessage() {}

func (x *ContributionAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContributionAndProof.ProtoReflect.Descriptor instead.
func (*ContributionAndProof) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_sync_committee_proto_rawDescGZIP(), []int{2}
}

func (x *ContributionAndProof) GetAggregatorIndex() github_com_prysmaticlabs_eth2_types.ValidatorIndex {
	if x != nil {
		return x.AggregatorIndex
	}
	return github_com_prysmaticlabs_eth2_types.ValidatorIndex(0)
}

func (x *ContributionAndProof) GetContribution() *SyncCommitteeContribution {
	if x != nil {
		return x.Contribution
	}
	return nil
}

func (x *ContributionAndProof) GetSelectionProof() []byte {
	if x != nil {
		return x.SelectionProof
	}
	return nil
}

type SignedContributionAndProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *ContributionAndProof `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Signature []byte                `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty" ssz-size:"96"`
}

func (x *SignedContributionAndProof) Reset() {
	*x = SignedContributionAndProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignedContributionAndProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedContributionAndProof) ProtoMessage() {}

func (x *SignedContributionAndProof) ProtoReflect() protoreflect.Message {
	mi := &file_eth_v1alpha1_sync_committee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedContributionAndProof.ProtoReflect.Descriptor instead.
func (*SignedContributionAndProof) Descriptor() ([]byte, []int) {
	return file_eth_v1alpha1_sync_committee_proto_rawDescGZIP(), []int{3}
}

func (x *SignedContributionAndProof) GetMessage() *ContributionAndProof {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SignedContributionAndProof) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_eth_v1alpha1_sync_committee_proto protoreflect.FileDescriptor

var file_eth_v1alpha1_sync_committee_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x65, 0x74, 0x68, 0x65, 0x72,
	0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x65, 0x78, 0x74, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02,
	0x0a, 0x16, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x53,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x53, 0x6c, 0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06,
	0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6f,
	0x74, 0x12, 0x5f, 0x0a, 0x0f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x36, 0x82, 0xb5, 0x18, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x0e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x8c, 0x02, 0x0a, 0x19, 0x53, 0x79, 0x6e,
	0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x2c, 0x82, 0xb5, 0x18, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x6c, 0x61,
	0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x04, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5,
	0x18, 0x02, 0x33, 0x32, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x74, 0x12,
	0x2d, 0x0a, 0x12, 0x73, 0x75, 0x62, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x73, 0x75, 0x62,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x31,
	0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x62, 0x69,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x33, 0x32,
	0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x69, 0x74,
	0x73, 0x12, 0x24, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x80, 0x02, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x12, 0x61, 0x0a, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x36, 0x82, 0xb5, 0x18, 0x32,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d,
	0x61, 0x74, 0x69, 0x63, 0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x32, 0x2d, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x52, 0x0f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x54, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x0f, 0x73, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x6f, 0x66, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x53,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x45, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x65, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x64, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x24, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x42, 0x06, 0x8a, 0xb5, 0x18, 0x02, 0x39, 0x36, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x99, 0x01, 0x0a, 0x19, 0x6f, 0x72, 0x67, 0x2e, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74,
	0x74, 0x65, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x61, 0x74, 0x69, 0x63,
	0x6c, 0x61, 0x62, 0x73, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x65, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x65,
	0x74, 0x68, 0xaa, 0x02, 0x15, 0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x45, 0x74,
	0x68, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x15, 0x45, 0x74, 0x68,
	0x65, 0x72, 0x65, 0x75, 0x6d, 0x5c, 0x45, 0x74, 0x68, 0x5c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_eth_v1alpha1_sync_committee_proto_rawDescOnce sync.Once
	file_eth_v1alpha1_sync_committee_proto_rawDescData = file_eth_v1alpha1_sync_committee_proto_rawDesc
)

func file_eth_v1alpha1_sync_committee_proto_rawDescGZIP() []byte {
	file_eth_v1alpha1_sync_committee_proto_rawDescOnce.Do(func() {
		file_eth_v1alpha1_sync_committee_proto_rawDescData = protoimpl.X.CompressGZIP(file_eth_v1alpha1_sync_committee_proto_rawDescData)
	})
	return file_eth_v1alpha1_sync_committee_proto_rawDescData
}

var file_eth_v1alpha1_sync_committee_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_eth_v1alpha1_sync_committee_proto_goTypes = []interface{}{
	(*SyncCommitteeSignature)(nil),     // 0: ethereum.eth.v1alpha1.SyncCommitteeSignature
	(*SyncCommitteeContribution)(nil),  // 1: ethereum.eth.v1alpha1.SyncCommitteeContribution
	(*ContributionAndProof)(nil),       // 2: ethereum.eth.v1alpha1.ContributionAndProof
	(*SignedContributionAndProof)(nil), // 3: ethereum.eth.v1alpha1.SignedContributionAndProof
}
var file_eth_v1alpha1_sync_committee_proto_depIdxs = []int32{
	1, // 0: ethereum.eth.v1alpha1.ContributionAndProof.contribution:type_name -> ethereum.eth.v1alpha1.SyncCommitteeContribution
	2, // 1: ethereum.eth.v1alpha1.SignedContributionAndProof.message:type_name -> ethereum.eth.v1alpha1.ContributionAndProof
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_eth_v1alpha1_sync_committee_proto_init() }
func file_eth_v1alpha1_sync_committee_proto_init() {
	if File_eth_v1alpha1_sync_committee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eth_v1alpha1_sync_committee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCommitteeSignature); i {
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
		file_eth_v1alpha1_sync_committee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncCommitteeContribution); i {
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
		file_eth_v1alpha1_sync_committee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContributionAndProof); i {
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
		file_eth_v1alpha1_sync_committee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignedContributionAndProof); i {
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
			RawDescriptor: file_eth_v1alpha1_sync_committee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_eth_v1alpha1_sync_committee_proto_goTypes,
		DependencyIndexes: file_eth_v1alpha1_sync_committee_proto_depIdxs,
		MessageInfos:      file_eth_v1alpha1_sync_committee_proto_msgTypes,
	}.Build()
	File_eth_v1alpha1_sync_committee_proto = out.File
	file_eth_v1alpha1_sync_committee_proto_rawDesc = nil
	file_eth_v1alpha1_sync_committee_proto_goTypes = nil
	file_eth_v1alpha1_sync_committee_proto_depIdxs = nil
}
