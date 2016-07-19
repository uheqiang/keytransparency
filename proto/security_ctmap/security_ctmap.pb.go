// Code generated by protoc-gen-go.
// source: proto/security_ctmap/security_ctmap.proto
// DO NOT EDIT!

/*
Package security_ctmap is a generated protocol buffer package.

It is generated from these files:
	proto/security_ctmap/security_ctmap.proto

It has these top-level messages:
	MapHead
	SignedMapHead
	DigitallySigned
	GetLeafRequest
	GetLeafResponse
	UpdateLeafRequest
	UpdateLeafResponse
	MutationEntry
*/
package security_ctmap

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import security_protobuf "github.com/google/key-transparency/proto/security_protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// HashAlgorithm defines the approved ways to hash the object.
type DigitallySigned_HashAlgorithm int32

const (
	DigitallySigned_NONE   DigitallySigned_HashAlgorithm = 0
	DigitallySigned_SHA256 DigitallySigned_HashAlgorithm = 4
	DigitallySigned_SHA512 DigitallySigned_HashAlgorithm = 6
)

var DigitallySigned_HashAlgorithm_name = map[int32]string{
	0: "NONE",
	4: "SHA256",
	6: "SHA512",
}
var DigitallySigned_HashAlgorithm_value = map[string]int32{
	"NONE":   0,
	"SHA256": 4,
	"SHA512": 6,
}

func (x DigitallySigned_HashAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_HashAlgorithm_name, int32(x))
}
func (DigitallySigned_HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 0}
}

// SignatureAlgorithm defines the way to sign the object.
type DigitallySigned_SignatureAlgorithm int32

const (
	DigitallySigned_ANONYMOUS DigitallySigned_SignatureAlgorithm = 0
	DigitallySigned_ECDSA     DigitallySigned_SignatureAlgorithm = 3
)

var DigitallySigned_SignatureAlgorithm_name = map[int32]string{
	0: "ANONYMOUS",
	3: "ECDSA",
}
var DigitallySigned_SignatureAlgorithm_value = map[string]int32{
	"ANONYMOUS": 0,
	"ECDSA":     3,
}

func (x DigitallySigned_SignatureAlgorithm) String() string {
	return proto.EnumName(DigitallySigned_SignatureAlgorithm_name, int32(x))
}
func (DigitallySigned_SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2, 1}
}

// MapHead is the head node of the Merkle Tree as well as additional metadata
// for the tree.
type MapHead struct {
	// realm is the domain identifier for the transparent map.
	Realm string `protobuf:"bytes,1,opt,name=realm" json:"realm,omitempty"`
	// epoch number
	Epoch int64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
	// root is the value of the root node of the merkle tree.
	Root []byte `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	// issue_time is the time when this epoch was released. Monotonically increasing.
	IssueTime *security_protobuf.Timestamp `protobuf:"bytes,4,opt,name=issue_time,json=issueTime" json:"issue_time,omitempty"`
}

func (m *MapHead) Reset()                    { *m = MapHead{} }
func (m *MapHead) String() string            { return proto.CompactTextString(m) }
func (*MapHead) ProtoMessage()               {}
func (*MapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MapHead) GetIssueTime() *security_protobuf.Timestamp {
	if m != nil {
		return m.IssueTime
	}
	return nil
}

// SignedMapHead represents a signed state of the Merkel tree.
type SignedMapHead struct {
	MapHead *MapHead `protobuf:"bytes,1,opt,name=map_head,json=mapHead" json:"map_head,omitempty"`
	// Signature of head, using the signature type of the key.
	// keyed by the first 64 bits bytes of the hash of the key.
	// TODO: Limit 1. Servers should only sign with one key at a time.
	// TODO: Create separate data structure for aggregating signatures from monitors.
	Signatures map[string]*DigitallySigned `protobuf:"bytes,2,rep,name=signatures" json:"signatures,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SignedMapHead) Reset()                    { *m = SignedMapHead{} }
func (m *SignedMapHead) String() string            { return proto.CompactTextString(m) }
func (*SignedMapHead) ProtoMessage()               {}
func (*SignedMapHead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SignedMapHead) GetMapHead() *MapHead {
	if m != nil {
		return m.MapHead
	}
	return nil
}

func (m *SignedMapHead) GetSignatures() map[string]*DigitallySigned {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// DigitallySigned defines a way to sign digital objects.
type DigitallySigned struct {
	HashAlgorithm DigitallySigned_HashAlgorithm      `protobuf:"varint,1,opt,name=hash_algorithm,json=hashAlgorithm,enum=security_ctmap.DigitallySigned_HashAlgorithm" json:"hash_algorithm,omitempty"`
	SigAlgorithm  DigitallySigned_SignatureAlgorithm `protobuf:"varint,2,opt,name=sig_algorithm,json=sigAlgorithm,enum=security_ctmap.DigitallySigned_SignatureAlgorithm" json:"sig_algorithm,omitempty"`
	Signature     []byte                             `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *DigitallySigned) Reset()                    { *m = DigitallySigned{} }
func (m *DigitallySigned) String() string            { return proto.CompactTextString(m) }
func (*DigitallySigned) ProtoMessage()               {}
func (*DigitallySigned) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// GetLeafRequest for a verifiable map leaf.
type GetLeafRequest struct {
	Index []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Epoch uint64 `protobuf:"varint,2,opt,name=epoch" json:"epoch,omitempty"`
}

func (m *GetLeafRequest) Reset()                    { *m = GetLeafRequest{} }
func (m *GetLeafRequest) String() string            { return proto.CompactTextString(m) }
func (*GetLeafRequest) ProtoMessage()               {}
func (*GetLeafRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// GetLeafResponse for a verifiable map leaf.
type GetLeafResponse struct {
	LeafData []byte `protobuf:"bytes,1,opt,name=leaf_data,json=leafData,proto3" json:"leaf_data,omitempty"`
	// neighbors is a list of all the adjacent nodes along the path
	// from the bottommost node to the head.
	Neighbors [][]byte `protobuf:"bytes,2,rep,name=neighbors,proto3" json:"neighbors,omitempty"`
}

func (m *GetLeafResponse) Reset()                    { *m = GetLeafResponse{} }
func (m *GetLeafResponse) String() string            { return proto.CompactTextString(m) }
func (*GetLeafResponse) ProtoMessage()               {}
func (*GetLeafResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

// UpdateLeafRequest submits a change for the value at index.
type UpdateLeafRequest struct {
	Index    []byte `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Mutation []byte `protobuf:"bytes,2,opt,name=mutation,proto3" json:"mutation,omitempty"`
}

func (m *UpdateLeafRequest) Reset()                    { *m = UpdateLeafRequest{} }
func (m *UpdateLeafRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateLeafRequest) ProtoMessage()               {}
func (*UpdateLeafRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// UpdateLeafResponse returns the current value of index.
type UpdateLeafResponse struct {
	Proof *GetLeafResponse `protobuf:"bytes,1,opt,name=proof" json:"proof,omitempty"`
}

func (m *UpdateLeafResponse) Reset()                    { *m = UpdateLeafResponse{} }
func (m *UpdateLeafResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateLeafResponse) ProtoMessage()               {}
func (*UpdateLeafResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateLeafResponse) GetProof() *GetLeafResponse {
	if m != nil {
		return m.Proof
	}
	return nil
}

// MutationEntry is either a mutation or an epoch advancement.
type MutationEntry struct {
	// Types that are valid to be assigned to Type:
	//	*MutationEntry_Update
	//	*MutationEntry_AdvanceEpoch
	Type isMutationEntry_Type `protobuf_oneof:"type"`
}

func (m *MutationEntry) Reset()                    { *m = MutationEntry{} }
func (m *MutationEntry) String() string            { return proto.CompactTextString(m) }
func (*MutationEntry) ProtoMessage()               {}
func (*MutationEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isMutationEntry_Type interface {
	isMutationEntry_Type()
}

type MutationEntry_Update struct {
	Update *UpdateLeafRequest `protobuf:"bytes,1,opt,name=update,oneof"`
}
type MutationEntry_AdvanceEpoch struct {
	AdvanceEpoch bool `protobuf:"varint,2,opt,name=advance_epoch,json=advanceEpoch,oneof"`
}

func (*MutationEntry_Update) isMutationEntry_Type()       {}
func (*MutationEntry_AdvanceEpoch) isMutationEntry_Type() {}

func (m *MutationEntry) GetType() isMutationEntry_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *MutationEntry) GetUpdate() *UpdateLeafRequest {
	if x, ok := m.GetType().(*MutationEntry_Update); ok {
		return x.Update
	}
	return nil
}

func (m *MutationEntry) GetAdvanceEpoch() bool {
	if x, ok := m.GetType().(*MutationEntry_AdvanceEpoch); ok {
		return x.AdvanceEpoch
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MutationEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MutationEntry_OneofMarshaler, _MutationEntry_OneofUnmarshaler, _MutationEntry_OneofSizer, []interface{}{
		(*MutationEntry_Update)(nil),
		(*MutationEntry_AdvanceEpoch)(nil),
	}
}

func _MutationEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MutationEntry)
	// type
	switch x := m.Type.(type) {
	case *MutationEntry_Update:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Update); err != nil {
			return err
		}
	case *MutationEntry_AdvanceEpoch:
		t := uint64(0)
		if x.AdvanceEpoch {
			t = 1
		}
		b.EncodeVarint(2<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case nil:
	default:
		return fmt.Errorf("MutationEntry.Type has unexpected type %T", x)
	}
	return nil
}

func _MutationEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MutationEntry)
	switch tag {
	case 1: // type.update
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(UpdateLeafRequest)
		err := b.DecodeMessage(msg)
		m.Type = &MutationEntry_Update{msg}
		return true, err
	case 2: // type.advance_epoch
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Type = &MutationEntry_AdvanceEpoch{x != 0}
		return true, err
	default:
		return false, nil
	}
}

func _MutationEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MutationEntry)
	// type
	switch x := m.Type.(type) {
	case *MutationEntry_Update:
		s := proto.Size(x.Update)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MutationEntry_AdvanceEpoch:
		n += proto.SizeVarint(2<<3 | proto.WireVarint)
		n += 1
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*MapHead)(nil), "security_ctmap.MapHead")
	proto.RegisterType((*SignedMapHead)(nil), "security_ctmap.SignedMapHead")
	proto.RegisterType((*DigitallySigned)(nil), "security_ctmap.DigitallySigned")
	proto.RegisterType((*GetLeafRequest)(nil), "security_ctmap.GetLeafRequest")
	proto.RegisterType((*GetLeafResponse)(nil), "security_ctmap.GetLeafResponse")
	proto.RegisterType((*UpdateLeafRequest)(nil), "security_ctmap.UpdateLeafRequest")
	proto.RegisterType((*UpdateLeafResponse)(nil), "security_ctmap.UpdateLeafResponse")
	proto.RegisterType((*MutationEntry)(nil), "security_ctmap.MutationEntry")
	proto.RegisterEnum("security_ctmap.DigitallySigned_HashAlgorithm", DigitallySigned_HashAlgorithm_name, DigitallySigned_HashAlgorithm_value)
	proto.RegisterEnum("security_ctmap.DigitallySigned_SignatureAlgorithm", DigitallySigned_SignatureAlgorithm_name, DigitallySigned_SignatureAlgorithm_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for VerifiableMapService service

type VerifiableMapServiceClient interface {
	// GetLeaf retrieves the value stored at a particular index.
	GetLeaf(ctx context.Context, in *GetLeafRequest, opts ...grpc.CallOption) (*GetLeafResponse, error)
	// UpdateLeaf submits a change to the value at index. Clients retry until
	// change is visble in GetLeafResponse.
	UpdateLeaf(ctx context.Context, in *UpdateLeafRequest, opts ...grpc.CallOption) (*GetLeafResponse, error)
}

type verifiableMapServiceClient struct {
	cc *grpc.ClientConn
}

func NewVerifiableMapServiceClient(cc *grpc.ClientConn) VerifiableMapServiceClient {
	return &verifiableMapServiceClient{cc}
}

func (c *verifiableMapServiceClient) GetLeaf(ctx context.Context, in *GetLeafRequest, opts ...grpc.CallOption) (*GetLeafResponse, error) {
	out := new(GetLeafResponse)
	err := grpc.Invoke(ctx, "/security_ctmap.VerifiableMapService/GetLeaf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verifiableMapServiceClient) UpdateLeaf(ctx context.Context, in *UpdateLeafRequest, opts ...grpc.CallOption) (*GetLeafResponse, error) {
	out := new(GetLeafResponse)
	err := grpc.Invoke(ctx, "/security_ctmap.VerifiableMapService/UpdateLeaf", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VerifiableMapService service

type VerifiableMapServiceServer interface {
	// GetLeaf retrieves the value stored at a particular index.
	GetLeaf(context.Context, *GetLeafRequest) (*GetLeafResponse, error)
	// UpdateLeaf submits a change to the value at index. Clients retry until
	// change is visble in GetLeafResponse.
	UpdateLeaf(context.Context, *UpdateLeafRequest) (*GetLeafResponse, error)
}

func RegisterVerifiableMapServiceServer(s *grpc.Server, srv VerifiableMapServiceServer) {
	s.RegisterService(&_VerifiableMapService_serviceDesc, srv)
}

func _VerifiableMapService_GetLeaf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeafRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiableMapServiceServer).GetLeaf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security_ctmap.VerifiableMapService/GetLeaf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiableMapServiceServer).GetLeaf(ctx, req.(*GetLeafRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerifiableMapService_UpdateLeaf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateLeafRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerifiableMapServiceServer).UpdateLeaf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/security_ctmap.VerifiableMapService/UpdateLeaf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerifiableMapServiceServer).UpdateLeaf(ctx, req.(*UpdateLeafRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VerifiableMapService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "security_ctmap.VerifiableMapService",
	HandlerType: (*VerifiableMapServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLeaf",
			Handler:    _VerifiableMapService_GetLeaf_Handler,
		},
		{
			MethodName: "UpdateLeaf",
			Handler:    _VerifiableMapService_UpdateLeaf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("proto/security_ctmap/security_ctmap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0x6e, 0x7e, 0x9a, 0x26, 0xd3, 0x38, 0xcd, 0x59, 0x55, 0x3a, 0x51, 0x4e, 0x75, 0xce, 0xc1,
	0x12, 0x12, 0x5c, 0xd4, 0x51, 0x8d, 0x8a, 0x10, 0xe5, 0x26, 0xd0, 0x88, 0x08, 0x9a, 0x16, 0x39,
	0x2d, 0x88, 0x1b, 0xa2, 0x4d, 0x32, 0x89, 0x2d, 0x1c, 0xdb, 0xd8, 0xeb, 0x8a, 0x88, 0x7b, 0x9e,
	0x88, 0x17, 0xe0, 0x95, 0x78, 0x02, 0xc6, 0xeb, 0xcd, 0x9f, 0x23, 0x94, 0xbb, 0x9d, 0xd9, 0xf9,
	0xbe, 0xf9, 0xbe, 0xf1, 0xac, 0xe1, 0x71, 0x10, 0xfa, 0xc2, 0x6f, 0x45, 0x38, 0x8a, 0x43, 0x47,
	0xcc, 0x07, 0x23, 0x31, 0xe3, 0x41, 0x26, 0x34, 0x64, 0x0d, 0xab, 0x6d, 0x66, 0x9b, 0xdd, 0xa9,
	0x23, 0xec, 0x78, 0x68, 0x8c, 0xfc, 0x59, 0x6b, 0xea, 0xfb, 0x53, 0x17, 0x5b, 0x68, 0xe2, 0xe9,
	0x67, 0x9c, 0x9f, 0x46, 0x18, 0xde, 0x63, 0xd8, 0xca, 0x70, 0xcb, 0x70, 0x18, 0x4f, 0x5a, 0xc2,
	0x99, 0x61, 0x24, 0xf8, 0x4c, 0x31, 0xeb, 0xdf, 0x73, 0x70, 0xd0, 0xe3, 0x41, 0x17, 0xf9, 0x98,
	0x1d, 0xc3, 0x7e, 0x88, 0xdc, 0x9d, 0x35, 0x72, 0xff, 0xe7, 0x1e, 0x55, 0xac, 0x34, 0x48, 0xb2,
	0x18, 0xf8, 0x23, 0xbb, 0x91, 0xa7, 0x6c, 0xc1, 0x4a, 0x03, 0xc6, 0xa0, 0x18, 0xfa, 0xbe, 0x68,
	0x14, 0x28, 0x59, 0xb5, 0xe4, 0x99, 0x5d, 0x00, 0x38, 0x51, 0x14, 0xe3, 0x20, 0x69, 0xd2, 0x28,
	0xd2, 0xcd, 0xa1, 0x79, 0x62, 0x2c, 0x34, 0x18, 0x0b, 0x0d, 0xc6, 0xed, 0x42, 0x83, 0x55, 0x91,
	0xf5, 0x49, 0xac, 0xff, 0xca, 0x81, 0xd6, 0x77, 0xa6, 0x1e, 0x8e, 0x17, 0x72, 0x4c, 0x28, 0x93,
	0xd7, 0x81, 0x4d, 0x67, 0xa9, 0xe8, 0xd0, 0xfc, 0xdb, 0xc8, 0x4c, 0x47, 0x95, 0x5a, 0x07, 0x33,
	0x85, 0xe9, 0x01, 0x44, 0x44, 0xc2, 0x45, 0x1c, 0x62, 0x44, 0x8a, 0x0b, 0x84, 0x3a, 0xcd, 0xa2,
	0x36, 0xda, 0xc8, 0x28, 0xad, 0xef, 0x78, 0x22, 0x9c, 0x5b, 0x6b, 0x04, 0xcd, 0x4f, 0x70, 0x94,
	0xb9, 0x66, 0x75, 0x28, 0xd0, 0x98, 0xd5, 0x88, 0x92, 0x23, 0x3b, 0x87, 0xfd, 0x7b, 0xee, 0xc6,
	0x28, 0x07, 0x74, 0x68, 0xfe, 0x97, 0x6d, 0x77, 0xe9, 0xd0, 0xd7, 0xe2, 0xae, 0x3b, 0x4f, 0xfb,
	0x5a, 0x69, 0xf5, 0xf3, 0xfc, 0xb3, 0x9c, 0xfe, 0x33, 0x0f, 0x47, 0x99, 0x6b, 0x76, 0x0b, 0x35,
	0x9b, 0x47, 0xf6, 0x80, 0xbb, 0x53, 0x9f, 0x58, 0xec, 0xf4, 0x73, 0xd4, 0xb6, 0x6d, 0x64, 0x80,
	0x46, 0x97, 0x50, 0xed, 0x05, 0xc8, 0xd2, 0xec, 0xf5, 0x90, 0x7d, 0x00, 0x8d, 0x7c, 0xad, 0x91,
	0xe6, 0x25, 0xa9, 0xb9, 0x8b, 0x74, 0x69, 0x7f, 0xc5, 0x5c, 0x25, 0xa2, 0x15, 0xf1, 0x09, 0x54,
	0x96, 0x03, 0x53, 0xdb, 0xb0, 0x4a, 0xe8, 0x67, 0xa0, 0x6d, 0xc8, 0x62, 0x65, 0x28, 0x5e, 0xdf,
	0x5c, 0x77, 0xea, 0x7b, 0x0c, 0xa0, 0xd4, 0xef, 0xb6, 0xcd, 0xf3, 0xa7, 0xf5, 0xa2, 0x3a, 0x9f,
	0x9f, 0x99, 0xf5, 0x92, 0x6e, 0x00, 0xdb, 0x6e, 0xca, 0x34, 0xa8, 0xb4, 0x09, 0xf8, 0xb1, 0x77,
	0x73, 0xd7, 0x27, 0x70, 0x05, 0xf6, 0x3b, 0xaf, 0x2e, 0xfb, 0xed, 0x7a, 0x41, 0x7f, 0x01, 0xb5,
	0xd7, 0x28, 0xae, 0x90, 0x4f, 0x2c, 0xfc, 0x12, 0xd3, 0x62, 0x25, 0x1b, 0xeb, 0x78, 0x63, 0xfc,
	0x2a, 0x07, 0x57, 0xb5, 0xd2, 0x60, 0x73, 0x8f, 0x8b, 0x6a, 0x8f, 0xf5, 0x2b, 0x38, 0x5a, 0xa2,
	0xa3, 0xc0, 0xf7, 0x22, 0x64, 0xff, 0x40, 0xc5, 0xa5, 0x78, 0x30, 0xe6, 0x82, 0x2b, 0x8a, 0x72,
	0x92, 0xb8, 0xa4, 0x38, 0xb1, 0xeb, 0xa1, 0x33, 0xb5, 0x87, 0x7e, 0x98, 0xee, 0x17, 0xd9, 0x5d,
	0x26, 0xf4, 0x0e, 0xfc, 0x75, 0x17, 0x10, 0x0e, 0x77, 0xcb, 0x69, 0xd2, 0x76, 0xc7, 0x82, 0x0b,
	0xc7, 0xf7, 0xa4, 0x22, 0x6a, 0xb2, 0x88, 0xf5, 0xb7, 0xc0, 0xd6, 0x69, 0x94, 0x2e, 0xda, 0x33,
	0x7a, 0x42, 0xfe, 0x44, 0x3d, 0x86, 0xad, 0x3d, 0xcb, 0xf8, 0xb0, 0xd2, 0x6a, 0xfd, 0x1b, 0x68,
	0x3d, 0x45, 0x9c, 0x6e, 0xf0, 0x05, 0x94, 0x62, 0xc9, 0xae, 0x88, 0x1e, 0x64, 0x89, 0xb6, 0x2c,
	0x74, 0xf7, 0x2c, 0x05, 0x61, 0x0f, 0x41, 0xe3, 0xe3, 0x7b, 0xee, 0x8d, 0x70, 0xb0, 0x9a, 0x66,
	0x99, 0x0a, 0xaa, 0x2a, 0xdd, 0x49, 0xb2, 0x2f, 0x4b, 0x50, 0x14, 0xf3, 0x00, 0xcd, 0x1f, 0x39,
	0x38, 0x7e, 0x8f, 0xa1, 0x33, 0x71, 0xf8, 0xd0, 0x45, 0x7a, 0x72, 0x7d, 0xfa, 0x3d, 0x39, 0x23,
	0x64, 0x6f, 0xe0, 0x40, 0xe9, 0x65, 0xff, 0xfe, 0xd1, 0x88, 0x6c, 0xde, 0xdc, 0x65, 0x94, 0xbd,
	0x03, 0x58, 0x49, 0x66, 0xbb, 0xed, 0xec, 0x64, 0x1c, 0x96, 0xe4, 0xbf, 0xea, 0xc9, 0xef, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xe7, 0x26, 0x2a, 0x6f, 0xa3, 0x05, 0x00, 0x00,
}
