// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	api.proto
	chaincode.proto
	chaincodeevent.proto
	devops.proto
	events.proto
	fabric.proto
	server_admin.proto

It has these top-level messages:
	BlockNumber
	BlockCount
	ChaincodeID
	ChaincodeInput
	ChaincodeSpec
	ChaincodeDeploymentSpec
	ChaincodeInvocationSpec
	ChaincodeSecurityContext
	ChaincodeMessage
	PutStateInfo
	RangeQueryState
	RangeQueryStateNext
	RangeQueryStateClose
	RangeQueryStateKeyValue
	RangeQueryStateResponse
	ChaincodeEvent
	Secret
	SigmaInput
	ExecuteWithBinding
	SigmaOutput
	BuildResult
	TransactionRequest
	ChaincodeReg
	Interest
	Register
	Rejection
	Unregister
	Event
	Transaction
	TransactionBlock
	TransactionResult
	Block
	BlockchainInfo
	NonHashData
	PeerAddress
	PeerID
	UniquePeerID
	PeerEndpoint
	PeersMessage
	PeersAddresses
	HelloMessage
	Message
	Response
	BlockState
	SyncBlockRange
	SyncBlocks
	SyncStateSnapshotRequest
	SyncStateSnapshot
	SyncStateDeltasRequest
	SyncStateDeltas
	ServerStatus
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

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

// Specifies the block number to be returned from the blockchain.
type BlockNumber struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *BlockNumber) Reset()                    { *m = BlockNumber{} }
func (m *BlockNumber) String() string            { return proto.CompactTextString(m) }
func (*BlockNumber) ProtoMessage()               {}
func (*BlockNumber) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BlockNumber) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

// Specifies the current number of blocks in the blockchain.
type BlockCount struct {
	Count uint64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *BlockCount) Reset()                    { *m = BlockCount{} }
func (m *BlockCount) String() string            { return proto.CompactTextString(m) }
func (*BlockCount) ProtoMessage()               {}
func (*BlockCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BlockCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*BlockNumber)(nil), "protos.BlockNumber")
	proto.RegisterType((*BlockCount)(nil), "protos.BlockCount")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Openchain service

type OpenchainClient interface {
	// GetBlockchainInfo returns information about the blockchain ledger such as
	// height, current block hash, and previous block hash.
	GetBlockchainInfo(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockchainInfo, error)
	// GetBlockByNumber returns the data contained within a specific block in the
	// blockchain. The genesis block is block zero.
	GetBlockByNumber(ctx context.Context, in *BlockNumber, opts ...grpc.CallOption) (*Block, error)
	// GetBlockCount returns the current number of blocks in the blockchain data
	// structure.
	GetBlockCount(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockCount, error)
	// GetPeers returns a list of all peer nodes currently connected to the target
	// peer.
	GetPeers(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*PeersMessage, error)
}

type openchainClient struct {
	cc *grpc.ClientConn
}

func NewOpenchainClient(cc *grpc.ClientConn) OpenchainClient {
	return &openchainClient{cc}
}

func (c *openchainClient) GetBlockchainInfo(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockchainInfo, error) {
	out := new(BlockchainInfo)
	err := grpc.Invoke(ctx, "/protos.Openchain/GetBlockchainInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openchainClient) GetBlockByNumber(ctx context.Context, in *BlockNumber, opts ...grpc.CallOption) (*Block, error) {
	out := new(Block)
	err := grpc.Invoke(ctx, "/protos.Openchain/GetBlockByNumber", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openchainClient) GetBlockCount(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*BlockCount, error) {
	out := new(BlockCount)
	err := grpc.Invoke(ctx, "/protos.Openchain/GetBlockCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *openchainClient) GetPeers(ctx context.Context, in *google_protobuf1.Empty, opts ...grpc.CallOption) (*PeersMessage, error) {
	out := new(PeersMessage)
	err := grpc.Invoke(ctx, "/protos.Openchain/GetPeers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Openchain service

type OpenchainServer interface {
	// GetBlockchainInfo returns information about the blockchain ledger such as
	// height, current block hash, and previous block hash.
	GetBlockchainInfo(context.Context, *google_protobuf1.Empty) (*BlockchainInfo, error)
	// GetBlockByNumber returns the data contained within a specific block in the
	// blockchain. The genesis block is block zero.
	GetBlockByNumber(context.Context, *BlockNumber) (*Block, error)
	// GetBlockCount returns the current number of blocks in the blockchain data
	// structure.
	GetBlockCount(context.Context, *google_protobuf1.Empty) (*BlockCount, error)
	// GetPeers returns a list of all peer nodes currently connected to the target
	// peer.
	GetPeers(context.Context, *google_protobuf1.Empty) (*PeersMessage, error)
}

func RegisterOpenchainServer(s *grpc.Server, srv OpenchainServer) {
	s.RegisterService(&_Openchain_serviceDesc, srv)
}

func _Openchain_GetBlockchainInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenchainServer).GetBlockchainInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Openchain/GetBlockchainInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenchainServer).GetBlockchainInfo(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openchain_GetBlockByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockNumber)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenchainServer).GetBlockByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Openchain/GetBlockByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenchainServer).GetBlockByNumber(ctx, req.(*BlockNumber))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openchain_GetBlockCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenchainServer).GetBlockCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Openchain/GetBlockCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenchainServer).GetBlockCount(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Openchain_GetPeers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf1.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpenchainServer).GetPeers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Openchain/GetPeers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpenchainServer).GetPeers(ctx, req.(*google_protobuf1.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Openchain_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Openchain",
	HandlerType: (*OpenchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockchainInfo",
			Handler:    _Openchain_GetBlockchainInfo_Handler,
		},
		{
			MethodName: "GetBlockByNumber",
			Handler:    _Openchain_GetBlockByNumber_Handler,
		},
		{
			MethodName: "GetBlockCount",
			Handler:    _Openchain_GetBlockCount_Handler,
		},
		{
			MethodName: "GetPeers",
			Handler:    _Openchain_GetPeers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0xc8, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x52, 0x3c, 0x69, 0x89, 0x49, 0x45,
	0x99, 0xc9, 0x10, 0x51, 0x29, 0xe9, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0x7d, 0x30, 0x2f, 0xa9,
	0x34, 0x4d, 0x3f, 0x35, 0xb7, 0xa0, 0xa4, 0x12, 0x22, 0xa9, 0xa4, 0xca, 0xc5, 0xed, 0x94, 0x93,
	0x9f, 0x9c, 0xed, 0x57, 0x9a, 0x9b, 0x94, 0x5a, 0x24, 0x24, 0xc6, 0xc5, 0x96, 0x07, 0x66, 0x49,
	0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0x41, 0x79, 0x4a, 0x4a, 0x5c, 0x5c, 0x60, 0x65, 0xce, 0xf9,
	0xa5, 0x79, 0x25, 0x42, 0x22, 0x5c, 0xac, 0xc9, 0x20, 0x06, 0x54, 0x11, 0x84, 0x63, 0xd4, 0xce,
	0xc4, 0xc5, 0xe9, 0x5f, 0x90, 0x9a, 0x97, 0x9c, 0x91, 0x98, 0x99, 0x27, 0xe4, 0xca, 0x25, 0xe8,
	0x9e, 0x5a, 0x02, 0xd6, 0x04, 0x16, 0xf0, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xd3, 0x83, 0xb8, 0x45,
	0x0f, 0xe6, 0x16, 0x3d, 0x57, 0x90, 0x5b, 0xa4, 0xc4, 0x20, 0x02, 0xc5, 0x7a, 0xa8, 0xea, 0x95,
	0x18, 0x84, 0x2c, 0xb8, 0x04, 0x60, 0xc6, 0x38, 0x55, 0x42, 0x1d, 0x29, 0x8c, 0xa2, 0x1a, 0x22,
	0x28, 0xc5, 0x8b, 0x22, 0xa8, 0xc4, 0x20, 0x64, 0xcb, 0xc5, 0x0b, 0xd3, 0x09, 0x71, 0x35, 0x2e,
	0xcb, 0x85, 0x50, 0x74, 0x82, 0xd5, 0x2a, 0x31, 0x08, 0x59, 0x71, 0x71, 0xb8, 0xa7, 0x96, 0x04,
	0xa4, 0xa6, 0x16, 0x15, 0xe3, 0xd4, 0x29, 0x02, 0xd3, 0x09, 0x56, 0xe6, 0x9b, 0x5a, 0x5c, 0x9c,
	0x98, 0x9e, 0xaa, 0xc4, 0x90, 0x04, 0x89, 0x07, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41,
	0xce, 0xa8, 0xc9, 0x9b, 0x01, 0x00, 0x00,
}
