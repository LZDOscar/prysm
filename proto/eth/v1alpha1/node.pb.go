// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/eth/v1alpha1/node.proto

package eth

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	io "io"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SyncStatus struct {
	Syncing              bool     `protobuf:"varint,1,opt,name=syncing,proto3" json:"syncing,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncStatus) Reset()         { *m = SyncStatus{} }
func (m *SyncStatus) String() string { return proto.CompactTextString(m) }
func (*SyncStatus) ProtoMessage()    {}
func (*SyncStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_98054421e2cad574, []int{0}
}
func (m *SyncStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SyncStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SyncStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SyncStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncStatus.Merge(m, src)
}
func (m *SyncStatus) XXX_Size() int {
	return m.Size()
}
func (m *SyncStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SyncStatus proto.InternalMessageInfo

func (m *SyncStatus) GetSyncing() bool {
	if m != nil {
		return m.Syncing
	}
	return false
}

type Genesis struct {
	GenesisTime            *types.Timestamp `protobuf:"bytes,1,opt,name=genesis_time,json=genesisTime,proto3" json:"genesis_time,omitempty"`
	DepositContractAddress []byte           `protobuf:"bytes,2,opt,name=deposit_contract_address,json=depositContractAddress,proto3" json:"deposit_contract_address,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}         `json:"-"`
	XXX_unrecognized       []byte           `json:"-"`
	XXX_sizecache          int32            `json:"-"`
}

func (m *Genesis) Reset()         { *m = Genesis{} }
func (m *Genesis) String() string { return proto.CompactTextString(m) }
func (*Genesis) ProtoMessage()    {}
func (*Genesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_98054421e2cad574, []int{1}
}
func (m *Genesis) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Genesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Genesis.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Genesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Genesis.Merge(m, src)
}
func (m *Genesis) XXX_Size() int {
	return m.Size()
}
func (m *Genesis) XXX_DiscardUnknown() {
	xxx_messageInfo_Genesis.DiscardUnknown(m)
}

var xxx_messageInfo_Genesis proto.InternalMessageInfo

func (m *Genesis) GetGenesisTime() *types.Timestamp {
	if m != nil {
		return m.GenesisTime
	}
	return nil
}

func (m *Genesis) GetDepositContractAddress() []byte {
	if m != nil {
		return m.DepositContractAddress
	}
	return nil
}

type Version struct {
	Version              string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Metadata             string   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_98054421e2cad574, []int{2}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Version.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return m.Size()
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Version) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

type ImplementedServices struct {
	Services             []string `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImplementedServices) Reset()         { *m = ImplementedServices{} }
func (m *ImplementedServices) String() string { return proto.CompactTextString(m) }
func (*ImplementedServices) ProtoMessage()    {}
func (*ImplementedServices) Descriptor() ([]byte, []int) {
	return fileDescriptor_98054421e2cad574, []int{3}
}
func (m *ImplementedServices) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ImplementedServices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ImplementedServices.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ImplementedServices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImplementedServices.Merge(m, src)
}
func (m *ImplementedServices) XXX_Size() int {
	return m.Size()
}
func (m *ImplementedServices) XXX_DiscardUnknown() {
	xxx_messageInfo_ImplementedServices.DiscardUnknown(m)
}

var xxx_messageInfo_ImplementedServices proto.InternalMessageInfo

func (m *ImplementedServices) GetServices() []string {
	if m != nil {
		return m.Services
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncStatus)(nil), "ethereum.eth.v1alpha1.SyncStatus")
	proto.RegisterType((*Genesis)(nil), "ethereum.eth.v1alpha1.Genesis")
	proto.RegisterType((*Version)(nil), "ethereum.eth.v1alpha1.Version")
	proto.RegisterType((*ImplementedServices)(nil), "ethereum.eth.v1alpha1.ImplementedServices")
}

func init() { proto.RegisterFile("proto/eth/v1alpha1/node.proto", fileDescriptor_98054421e2cad574) }

var fileDescriptor_98054421e2cad574 = []byte{
	// 462 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x41, 0x8b, 0x13, 0x31,
	0x14, 0xc7, 0xc9, 0xae, 0xd8, 0xdd, 0xec, 0x7a, 0x89, 0xb8, 0x0e, 0xb3, 0xbb, 0xb5, 0x8e, 0x20,
	0x8b, 0x87, 0x84, 0xae, 0x08, 0x82, 0x88, 0xe8, 0x22, 0x45, 0x10, 0x0f, 0x53, 0xf1, 0xe0, 0xa5,
	0xa4, 0x33, 0xcf, 0x99, 0x40, 0x93, 0x0c, 0x93, 0xd7, 0x42, 0xaf, 0xfd, 0x0a, 0x7e, 0x29, 0x6f,
	0x0a, 0x7e, 0x01, 0x29, 0x7e, 0x10, 0x99, 0x99, 0x44, 0xc5, 0xb6, 0x0b, 0xde, 0xf2, 0xf2, 0x7f,
	0xff, 0xf7, 0x4b, 0xff, 0xaf, 0x43, 0xcf, 0xab, 0xda, 0xa2, 0x15, 0x80, 0xa5, 0x58, 0x0c, 0xe5,
	0xac, 0x2a, 0xe5, 0x50, 0x18, 0x9b, 0x03, 0x6f, 0xef, 0xd9, 0x1d, 0xc0, 0x12, 0x6a, 0x98, 0x6b,
	0x0e, 0x58, 0xf2, 0xd0, 0x11, 0x9f, 0x15, 0xd6, 0x16, 0x33, 0x10, 0xb2, 0x52, 0x42, 0x1a, 0x63,
	0x51, 0xa2, 0xb2, 0xc6, 0x75, 0xa6, 0xf8, 0xd4, 0xab, 0x6d, 0x35, 0x9d, 0x7f, 0x12, 0xa0, 0x2b,
	0x5c, 0x7a, 0xf1, 0xde, 0xbf, 0x22, 0x2a, 0x0d, 0x0e, 0xa5, 0xae, 0xba, 0x86, 0xe4, 0x21, 0xa5,
	0xe3, 0xa5, 0xc9, 0xc6, 0x28, 0x71, 0xee, 0x58, 0x44, 0x7b, 0x6e, 0x69, 0x32, 0x65, 0x8a, 0x88,
	0x0c, 0xc8, 0xc5, 0x41, 0x1a, 0xca, 0x64, 0x45, 0x68, 0x6f, 0x04, 0x06, 0x9c, 0x72, 0xec, 0x39,
	0x3d, 0x2e, 0xba, 0xe3, 0xa4, 0x19, 0xd7, 0xb6, 0x1e, 0x5d, 0xc6, 0xbc, 0x63, 0xf1, 0xc0, 0xe2,
	0xef, 0x03, 0x2b, 0x3d, 0xf2, 0xfd, 0xcd, 0x0d, 0x7b, 0x4a, 0xa3, 0x1c, 0x2a, 0xeb, 0x14, 0x4e,
	0x32, 0x6b, 0xb0, 0x96, 0x19, 0x4e, 0x64, 0x9e, 0xd7, 0xe0, 0x5c, 0xb4, 0x37, 0x20, 0x17, 0xc7,
	0xe9, 0x89, 0xd7, 0xaf, 0xbc, 0xfc, 0xb2, 0x53, 0x93, 0x17, 0xb4, 0xf7, 0x01, 0x6a, 0xa7, 0xac,
	0x69, 0x5e, 0xba, 0xe8, 0x8e, 0x2d, 0xfe, 0x30, 0x0d, 0x25, 0x8b, 0xe9, 0x81, 0x06, 0x94, 0xb9,
	0x44, 0xd9, 0x8e, 0x3b, 0x4c, 0x7f, 0xd7, 0xc9, 0x90, 0xde, 0x7e, 0xa3, 0xab, 0x19, 0x68, 0x30,
	0x08, 0xf9, 0x18, 0xea, 0x85, 0xca, 0xc0, 0x35, 0x16, 0xe7, 0xcf, 0x11, 0x19, 0xec, 0x37, 0x96,
	0x50, 0x5f, 0x7e, 0xdd, 0xa7, 0x37, 0xde, 0xd9, 0x1c, 0x98, 0xa1, 0xb7, 0x46, 0x80, 0x7f, 0x85,
	0x75, 0xb2, 0xf1, 0x83, 0x5f, 0x37, 0xc9, 0xc7, 0xf7, 0xf9, 0xd6, 0x35, 0xf2, 0x3f, 0xd6, 0x24,
	0x59, 0x7d, 0xff, 0xf9, 0x79, 0xef, 0x8c, 0xc5, 0x9b, 0x7f, 0x05, 0xe1, 0x13, 0x67, 0x25, 0xa5,
	0x23, 0xc0, 0x90, 0xf9, 0x2e, 0x58, 0x7f, 0x07, 0xcc, 0xfb, 0xae, 0x25, 0xf9, 0xa5, 0x78, 0x52,
	0x48, 0xf6, 0x7f, 0x49, 0xde, 0x77, 0x2d, 0x29, 0xec, 0x66, 0x45, 0xe8, 0xdd, 0xb7, 0xca, 0xe1,
	0xb6, 0x25, 0xec, 0xe2, 0x3e, 0xda, 0xc1, 0xdd, 0x32, 0x23, 0x79, 0xd0, 0xbe, 0xe1, 0x9c, 0x9d,
	0x6e, 0xcb, 0xd5, 0x37, 0xbd, 0xba, 0xfa, 0xb2, 0xee, 0x93, 0x6f, 0xeb, 0x3e, 0xf9, 0xb1, 0xee,
	0x93, 0x8f, 0x4f, 0x0a, 0x85, 0xe5, 0x7c, 0xca, 0x33, 0xab, 0x45, 0x55, 0x2f, 0x9d, 0x96, 0xa8,
	0xb2, 0x99, 0x9c, 0xba, 0xae, 0x12, 0x9b, 0x5f, 0xec, 0x33, 0xc0, 0x72, 0x7a, 0xb3, 0xbd, 0x7f,
	0xfc, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x53, 0xd1, 0xfd, 0x57, 0xd2, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeClient is the client API for Node service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeClient interface {
	GetSyncStatus(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*SyncStatus, error)
	GetGenesis(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*Genesis, error)
	GetVersion(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*Version, error)
	ListImplementedServices(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ImplementedServices, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) GetSyncStatus(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*SyncStatus, error) {
	out := new(SyncStatus)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/GetSyncStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetGenesis(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*Genesis, error) {
	out := new(Genesis)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/GetGenesis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) GetVersion(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) ListImplementedServices(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*ImplementedServices, error) {
	out := new(ImplementedServices)
	err := c.cc.Invoke(ctx, "/ethereum.eth.v1alpha1.Node/ListImplementedServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeServer is the server API for Node service.
type NodeServer interface {
	GetSyncStatus(context.Context, *types.Empty) (*SyncStatus, error)
	GetGenesis(context.Context, *types.Empty) (*Genesis, error)
	GetVersion(context.Context, *types.Empty) (*Version, error)
	ListImplementedServices(context.Context, *types.Empty) (*ImplementedServices, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_GetSyncStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetSyncStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/GetSyncStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetSyncStatus(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetGenesis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetGenesis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/GetGenesis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetGenesis(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).GetVersion(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_ListImplementedServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).ListImplementedServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ethereum.eth.v1alpha1.Node/ListImplementedServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).ListImplementedServices(ctx, req.(*types.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ethereum.eth.v1alpha1.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSyncStatus",
			Handler:    _Node_GetSyncStatus_Handler,
		},
		{
			MethodName: "GetGenesis",
			Handler:    _Node_GetGenesis_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _Node_GetVersion_Handler,
		},
		{
			MethodName: "ListImplementedServices",
			Handler:    _Node_ListImplementedServices_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/eth/v1alpha1/node.proto",
}

func (m *SyncStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SyncStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Syncing {
		dAtA[i] = 0x8
		i++
		if m.Syncing {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Genesis) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Genesis) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.GenesisTime != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNode(dAtA, i, uint64(m.GenesisTime.Size()))
		n1, err := m.GenesisTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.DepositContractAddress) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.DepositContractAddress)))
		i += copy(dAtA[i:], m.DepositContractAddress)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Version) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Version) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Version) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.Version)))
		i += copy(dAtA[i:], m.Version)
	}
	if len(m.Metadata) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNode(dAtA, i, uint64(len(m.Metadata)))
		i += copy(dAtA[i:], m.Metadata)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *ImplementedServices) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ImplementedServices) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintNode(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SyncStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Syncing {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Genesis) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.GenesisTime != nil {
		l = m.GenesisTime.Size()
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.DepositContractAddress)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Version) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovNode(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ImplementedServices) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Services) > 0 {
		for _, s := range m.Services {
			l = len(s)
			n += 1 + l + sovNode(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovNode(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNode(x uint64) (n int) {
	return sovNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SyncStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SyncStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SyncStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Syncing", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Syncing = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Genesis) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Genesis: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Genesis: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.GenesisTime == nil {
				m.GenesisTime = &types.Timestamp{}
			}
			if err := m.GenesisTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositContractAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DepositContractAddress = append(m.DepositContractAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.DepositContractAddress == nil {
				m.DepositContractAddress = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Version) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Version: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Version: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ImplementedServices) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNode
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ImplementedServices: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ImplementedServices: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNode
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthNode
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNode
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNode
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNode
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthNode
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNode
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipNode(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthNode
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthNode = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNode   = fmt.Errorf("proto: integer overflow")
)
