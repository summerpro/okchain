// Code generated by protoc-gen-go. DO NOT EDIT.
// source: explorer.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	explorer.proto
	inner.proto
	message.proto
	okchain.proto

It has these top-level messages:
	AccountRequest
	AccountResponse
	TransactionResponse
	DsBlockRequest
	TxBlockRequest
	TransactionRequest
	EmptyRequest
	RecentTransactionsRequest
	RecentTransactionsResponse
	RegisterAccountRequest
	RegisterAccountResponse
	GetTransactionsByAccountRequest
	GetTransactionsByAccountResponse
	NetworkInfoRequest
	NetworkInfoResponse
	PeerStatusResponse
	Envelope
	SecretEnvelope
	Secret
	GossipMessage
	StateInfo
	Properties
	StateInfoSnapshot
	StateInfoPullRequest
	ConnEstablish
	PeerIdentity
	DataRequest
	GossipHello
	DataUpdate
	DataDigest
	DataMessage
	PrivateDataMessage
	Payload
	PrivatePayload
	AliveMessage
	LeadershipMessage
	PeerTime
	MembershipRequest
	MembershipResponse
	Member
	Empty
	RemoteStateRequest
	RemoteStateResponse
	RemotePvtDataRequest
	PvtDataDigest
	RemotePvtDataResponse
	PvtDataElement
	PvtDataPayload
	Acknowledgement
	Chaincode
	TransactionMessage
	Message
	PeerID
	Timestamp
	PeerEndpoint
	StartPoW
	PoWSubmission
	MiningResult
	ConsensusPayload
	Announce
	Commit
	Challenge
	Response
	CollectiveSig
	FinalCommit
	FinalChallenge
	FinalResponse
	FinalCollectiveSig
	VCBlockHeader
	VCBlock
	MicroBlock
	StakeWeight
	Account
	Transaction
	DSBlock
	DSBlockHeader
	DSBlockBody
	TxBlock
	TxBlockHeader
	TxBlockBody
	ConfigRpcResponse
	InformDs
	InformSharding
*/
package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// *********Account relative*****
type AccountRequest struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *AccountRequest) Reset()                    { *m = AccountRequest{} }
func (m *AccountRequest) String() string            { return proto.CompactTextString(m) }
func (*AccountRequest) ProtoMessage()               {}
func (*AccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AccountRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type AccountResponse struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *AccountResponse) Reset()                    { *m = AccountResponse{} }
func (m *AccountResponse) String() string            { return proto.CompactTextString(m) }
func (*AccountResponse) ProtoMessage()               {}
func (*AccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccountResponse) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type TransactionResponse struct {
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *TransactionResponse) Reset()                    { *m = TransactionResponse{} }
func (m *TransactionResponse) String() string            { return proto.CompactTextString(m) }
func (*TransactionResponse) ProtoMessage()               {}
func (*TransactionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TransactionResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type DsBlockRequest struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *DsBlockRequest) Reset()                    { *m = DsBlockRequest{} }
func (m *DsBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*DsBlockRequest) ProtoMessage()               {}
func (*DsBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DsBlockRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type TxBlockRequest struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *TxBlockRequest) Reset()                    { *m = TxBlockRequest{} }
func (m *TxBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*TxBlockRequest) ProtoMessage()               {}
func (*TxBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *TxBlockRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type TransactionRequest struct {
	Addr []byte `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
}

func (m *TransactionRequest) Reset()                    { *m = TransactionRequest{} }
func (m *TransactionRequest) String() string            { return proto.CompactTextString(m) }
func (*TransactionRequest) ProtoMessage()               {}
func (*TransactionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TransactionRequest) GetAddr() []byte {
	if m != nil {
		return m.Addr
	}
	return nil
}

type EmptyRequest struct {
}

func (m *EmptyRequest) Reset()                    { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string            { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()               {}
func (*EmptyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type RecentTransactionsRequest struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *RecentTransactionsRequest) Reset()                    { *m = RecentTransactionsRequest{} }
func (m *RecentTransactionsRequest) String() string            { return proto.CompactTextString(m) }
func (*RecentTransactionsRequest) ProtoMessage()               {}
func (*RecentTransactionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RecentTransactionsRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type RecentTransactionsResponse struct {
	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *RecentTransactionsResponse) Reset()                    { *m = RecentTransactionsResponse{} }
func (m *RecentTransactionsResponse) String() string            { return proto.CompactTextString(m) }
func (*RecentTransactionsResponse) ProtoMessage()               {}
func (*RecentTransactionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RecentTransactionsResponse) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type RegisterAccountRequest struct {
	Address []byte `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *RegisterAccountRequest) Reset()                    { *m = RegisterAccountRequest{} }
func (m *RegisterAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterAccountRequest) ProtoMessage()               {}
func (*RegisterAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RegisterAccountRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type RegisterAccountResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
}

func (m *RegisterAccountResponse) Reset()                    { *m = RegisterAccountResponse{} }
func (m *RegisterAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterAccountResponse) ProtoMessage()               {}
func (*RegisterAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RegisterAccountResponse) GetCode() uint32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type GetTransactionsByAccountRequest struct {
	Number  uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
	Address []byte `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *GetTransactionsByAccountRequest) Reset()         { *m = GetTransactionsByAccountRequest{} }
func (m *GetTransactionsByAccountRequest) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsByAccountRequest) ProtoMessage()    {}
func (*GetTransactionsByAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11}
}

func (m *GetTransactionsByAccountRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *GetTransactionsByAccountRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

type GetTransactionsByAccountResponse struct {
	Transactions []*Transaction `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *GetTransactionsByAccountResponse) Reset()         { *m = GetTransactionsByAccountResponse{} }
func (m *GetTransactionsByAccountResponse) String() string { return proto.CompactTextString(m) }
func (*GetTransactionsByAccountResponse) ProtoMessage()    {}
func (*GetTransactionsByAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{12}
}

func (m *GetTransactionsByAccountResponse) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

type NetworkInfoRequest struct {
	Number uint64 `protobuf:"varint,1,opt,name=number" json:"number,omitempty"`
}

func (m *NetworkInfoRequest) Reset()                    { *m = NetworkInfoRequest{} }
func (m *NetworkInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*NetworkInfoRequest) ProtoMessage()               {}
func (*NetworkInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *NetworkInfoRequest) GetNumber() uint64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type NetworkInfoResponse struct {
	DsList         []*PeerEndpoint `protobuf:"bytes,1,rep,name=DsList" json:"DsList,omitempty"`
	ShardingList   []*PeerEndpoint `protobuf:"bytes,2,rep,name=ShardingList" json:"ShardingList,omitempty"`
	LookupList     []*PeerEndpoint `protobuf:"bytes,3,rep,name=LookupList" json:"LookupList,omitempty"`
	RemotePeerList []*PeerEndpoint `protobuf:"bytes,4,rep,name=RemotePeerList" json:"RemotePeerList,omitempty"`
}

func (m *NetworkInfoResponse) Reset()                    { *m = NetworkInfoResponse{} }
func (m *NetworkInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*NetworkInfoResponse) ProtoMessage()               {}
func (*NetworkInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *NetworkInfoResponse) GetDsList() []*PeerEndpoint {
	if m != nil {
		return m.DsList
	}
	return nil
}

func (m *NetworkInfoResponse) GetShardingList() []*PeerEndpoint {
	if m != nil {
		return m.ShardingList
	}
	return nil
}

func (m *NetworkInfoResponse) GetLookupList() []*PeerEndpoint {
	if m != nil {
		return m.LookupList
	}
	return nil
}

func (m *NetworkInfoResponse) GetRemotePeerList() []*PeerEndpoint {
	if m != nil {
		return m.RemotePeerList
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountRequest)(nil), "protos.AccountRequest")
	proto.RegisterType((*AccountResponse)(nil), "protos.AccountResponse")
	proto.RegisterType((*TransactionResponse)(nil), "protos.TransactionResponse")
	proto.RegisterType((*DsBlockRequest)(nil), "protos.DsBlockRequest")
	proto.RegisterType((*TxBlockRequest)(nil), "protos.TxBlockRequest")
	proto.RegisterType((*TransactionRequest)(nil), "protos.TransactionRequest")
	proto.RegisterType((*EmptyRequest)(nil), "protos.EmptyRequest")
	proto.RegisterType((*RecentTransactionsRequest)(nil), "protos.RecentTransactionsRequest")
	proto.RegisterType((*RecentTransactionsResponse)(nil), "protos.RecentTransactionsResponse")
	proto.RegisterType((*RegisterAccountRequest)(nil), "protos.RegisterAccountRequest")
	proto.RegisterType((*RegisterAccountResponse)(nil), "protos.RegisterAccountResponse")
	proto.RegisterType((*GetTransactionsByAccountRequest)(nil), "protos.GetTransactionsByAccountRequest")
	proto.RegisterType((*GetTransactionsByAccountResponse)(nil), "protos.GetTransactionsByAccountResponse")
	proto.RegisterType((*NetworkInfoRequest)(nil), "protos.NetworkInfoRequest")
	proto.RegisterType((*NetworkInfoResponse)(nil), "protos.NetworkInfoResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Backend service

type BackendClient interface {
	GetNetworkInfo(ctx context.Context, in *NetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error)
	GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error)
	GetDsBlock(ctx context.Context, in *DsBlockRequest, opts ...grpc.CallOption) (*DSBlock, error)
	GetTxBlock(ctx context.Context, in *TxBlockRequest, opts ...grpc.CallOption) (*TxBlock, error)
	GetTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Transaction, error)
	GetRecentTransactions(ctx context.Context, in *RecentTransactionsRequest, opts ...grpc.CallOption) (*RecentTransactionsResponse, error)
	GetTransactionsByAccount(ctx context.Context, in *GetTransactionsByAccountRequest, opts ...grpc.CallOption) (*GetTransactionsByAccountResponse, error)
	GetLatestDsBlock(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DSBlock, error)
	GetLatestTxBlock(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TxBlock, error)
	RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error)
	SendTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error)
	Notify(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ConfigRpcResponse, error)
}

type backendClient struct {
	cc *grpc.ClientConn
}

func NewBackendClient(cc *grpc.ClientConn) BackendClient {
	return &backendClient{cc}
}

func (c *backendClient) GetNetworkInfo(ctx context.Context, in *NetworkInfoRequest, opts ...grpc.CallOption) (*NetworkInfoResponse, error) {
	out := new(NetworkInfoResponse)
	err := grpc.Invoke(ctx, "/protos.Backend/GetNetworkInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetAccount(ctx context.Context, in *AccountRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := grpc.Invoke(ctx, "/protos.Backend/GetAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetDsBlock(ctx context.Context, in *DsBlockRequest, opts ...grpc.CallOption) (*DSBlock, error) {
	out := new(DSBlock)
	err := grpc.Invoke(ctx, "/protos.Backend/GetDsBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetTxBlock(ctx context.Context, in *TxBlockRequest, opts ...grpc.CallOption) (*TxBlock, error) {
	out := new(TxBlock)
	err := grpc.Invoke(ctx, "/protos.Backend/GetTxBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetTransaction(ctx context.Context, in *TransactionRequest, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := grpc.Invoke(ctx, "/protos.Backend/GetTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetRecentTransactions(ctx context.Context, in *RecentTransactionsRequest, opts ...grpc.CallOption) (*RecentTransactionsResponse, error) {
	out := new(RecentTransactionsResponse)
	err := grpc.Invoke(ctx, "/protos.Backend/GetRecentTransactions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetTransactionsByAccount(ctx context.Context, in *GetTransactionsByAccountRequest, opts ...grpc.CallOption) (*GetTransactionsByAccountResponse, error) {
	out := new(GetTransactionsByAccountResponse)
	err := grpc.Invoke(ctx, "/protos.Backend/GetTransactionsByAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetLatestDsBlock(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*DSBlock, error) {
	out := new(DSBlock)
	err := grpc.Invoke(ctx, "/protos.Backend/GetLatestDsBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) GetLatestTxBlock(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*TxBlock, error) {
	out := new(TxBlock)
	err := grpc.Invoke(ctx, "/protos.Backend/GetLatestTxBlock", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) RegisterAccount(ctx context.Context, in *RegisterAccountRequest, opts ...grpc.CallOption) (*RegisterAccountResponse, error) {
	out := new(RegisterAccountResponse)
	err := grpc.Invoke(ctx, "/protos.Backend/RegisterAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) SendTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*TransactionResponse, error) {
	out := new(TransactionResponse)
	err := grpc.Invoke(ctx, "/protos.Backend/SendTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backendClient) Notify(ctx context.Context, in *Message, opts ...grpc.CallOption) (*ConfigRpcResponse, error) {
	out := new(ConfigRpcResponse)
	err := grpc.Invoke(ctx, "/protos.Backend/Notify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Backend service

type BackendServer interface {
	GetNetworkInfo(context.Context, *NetworkInfoRequest) (*NetworkInfoResponse, error)
	GetAccount(context.Context, *AccountRequest) (*AccountResponse, error)
	GetDsBlock(context.Context, *DsBlockRequest) (*DSBlock, error)
	GetTxBlock(context.Context, *TxBlockRequest) (*TxBlock, error)
	GetTransaction(context.Context, *TransactionRequest) (*Transaction, error)
	GetRecentTransactions(context.Context, *RecentTransactionsRequest) (*RecentTransactionsResponse, error)
	GetTransactionsByAccount(context.Context, *GetTransactionsByAccountRequest) (*GetTransactionsByAccountResponse, error)
	GetLatestDsBlock(context.Context, *EmptyRequest) (*DSBlock, error)
	GetLatestTxBlock(context.Context, *EmptyRequest) (*TxBlock, error)
	RegisterAccount(context.Context, *RegisterAccountRequest) (*RegisterAccountResponse, error)
	SendTransaction(context.Context, *Transaction) (*TransactionResponse, error)
	Notify(context.Context, *Message) (*ConfigRpcResponse, error)
}

func RegisterBackendServer(s *grpc.Server, srv BackendServer) {
	s.RegisterService(&_Backend_serviceDesc, srv)
}

func _Backend_GetNetworkInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetNetworkInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetNetworkInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetNetworkInfo(ctx, req.(*NetworkInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetAccount(ctx, req.(*AccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetDsBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DsBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetDsBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetDsBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetDsBlock(ctx, req.(*DsBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetTxBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetTxBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetTxBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetTxBlock(ctx, req.(*TxBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetTransaction(ctx, req.(*TransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetRecentTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecentTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetRecentTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetRecentTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetRecentTransactions(ctx, req.(*RecentTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetTransactionsByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionsByAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetTransactionsByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetTransactionsByAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetTransactionsByAccount(ctx, req.(*GetTransactionsByAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetLatestDsBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetLatestDsBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetLatestDsBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetLatestDsBlock(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_GetLatestTxBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).GetLatestTxBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/GetLatestTxBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).GetLatestTxBlock(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_RegisterAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).RegisterAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/RegisterAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).RegisterAccount(ctx, req.(*RegisterAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).SendTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Backend_Notify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackendServer).Notify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Backend/Notify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackendServer).Notify(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Backend_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Backend",
	HandlerType: (*BackendServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkInfo",
			Handler:    _Backend_GetNetworkInfo_Handler,
		},
		{
			MethodName: "GetAccount",
			Handler:    _Backend_GetAccount_Handler,
		},
		{
			MethodName: "GetDsBlock",
			Handler:    _Backend_GetDsBlock_Handler,
		},
		{
			MethodName: "GetTxBlock",
			Handler:    _Backend_GetTxBlock_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Backend_GetTransaction_Handler,
		},
		{
			MethodName: "GetRecentTransactions",
			Handler:    _Backend_GetRecentTransactions_Handler,
		},
		{
			MethodName: "GetTransactionsByAccount",
			Handler:    _Backend_GetTransactionsByAccount_Handler,
		},
		{
			MethodName: "GetLatestDsBlock",
			Handler:    _Backend_GetLatestDsBlock_Handler,
		},
		{
			MethodName: "GetLatestTxBlock",
			Handler:    _Backend_GetLatestTxBlock_Handler,
		},
		{
			MethodName: "RegisterAccount",
			Handler:    _Backend_RegisterAccount_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _Backend_SendTransaction_Handler,
		},
		{
			MethodName: "Notify",
			Handler:    _Backend_Notify_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "explorer.proto",
}

func init() { proto.RegisterFile("explorer.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x5d, 0x6f, 0x12, 0x4d,
	0x14, 0x86, 0x96, 0xd0, 0xbc, 0xa7, 0x74, 0x79, 0x33, 0xd5, 0x96, 0xae, 0x89, 0xc5, 0xb9, 0x91,
	0x9a, 0xda, 0x0b, 0x6a, 0xac, 0x89, 0x4d, 0x8c, 0x14, 0x42, 0x8c, 0xd8, 0x98, 0x05, 0xaf, 0xbc,
	0x71, 0xbb, 0x7b, 0x80, 0x0d, 0x65, 0x66, 0xdd, 0x19, 0x62, 0xf9, 0xd7, 0xfa, 0x0f, 0x0c, 0xb3,
	0xb3, 0x5f, 0xb0, 0x40, 0x13, 0xaf, 0x18, 0xce, 0x3c, 0xcf, 0xf9, 0x7a, 0xe6, 0x9c, 0x05, 0x03,
	0x1f, 0xfc, 0x7b, 0x1e, 0x60, 0x70, 0xe1, 0x07, 0x5c, 0x72, 0x52, 0x56, 0x3f, 0xc2, 0x3c, 0xe0,
	0x13, 0x67, 0x6c, 0x7b, 0x2c, 0x34, 0xd3, 0x57, 0x60, 0x7c, 0x74, 0x1c, 0x3e, 0x63, 0xd2, 0xc2,
	0x9f, 0x33, 0x14, 0x92, 0xd4, 0x60, 0xcf, 0x76, 0xdd, 0x00, 0x85, 0xa8, 0x15, 0xeb, 0xc5, 0x46,
	0xc5, 0x8a, 0xfe, 0xd2, 0x6b, 0xa8, 0xc6, 0x58, 0xe1, 0x73, 0x26, 0x90, 0x9c, 0xc1, 0x9e, 0x1d,
	0x9a, 0x14, 0x78, 0xbf, 0x59, 0x0d, 0xfd, 0x8a, 0x8b, 0x08, 0x19, 0xdd, 0xd3, 0x33, 0x38, 0x1c,
	0x04, 0x36, 0x13, 0xb6, 0x23, 0x3d, 0xce, 0x62, 0x0f, 0x04, 0x4a, 0x63, 0x5b, 0x8c, 0x15, 0xfd,
	0x3f, 0x4b, 0x9d, 0x69, 0x03, 0x8c, 0xb6, 0x68, 0xdd, 0x73, 0x67, 0x12, 0x25, 0x75, 0x04, 0x65,
	0x36, 0x9b, 0xde, 0x61, 0xa0, 0x70, 0x25, 0x4b, 0xff, 0x5b, 0x20, 0x07, 0x0f, 0x8f, 0x44, 0x92,
	0x4c, 0xf8, 0x10, 0x4d, 0xa0, 0xb4, 0xa8, 0x4e, 0x57, 0xaa, 0xce, 0xd4, 0x80, 0x4a, 0x67, 0xea,
	0xcb, 0xb9, 0xc6, 0xd0, 0x4b, 0x38, 0xb1, 0xd0, 0x41, 0x26, 0x53, 0x7c, 0xb1, 0x2d, 0xdc, 0x37,
	0x30, 0xf3, 0x48, 0xba, 0xe8, 0x2b, 0xa8, 0xc8, 0x94, 0xbd, 0x56, 0xac, 0xef, 0x36, 0xf6, 0x9b,
	0x87, 0x51, 0xef, 0xd2, 0x89, 0x66, 0x80, 0xb4, 0x09, 0x47, 0x16, 0x8e, 0x3c, 0x21, 0x31, 0x78,
	0xb4, 0x6c, 0xaf, 0xe1, 0x78, 0x85, 0x93, 0x34, 0xdf, 0xe1, 0x2e, 0x2a, 0xc6, 0x81, 0xa5, 0xce,
	0xb4, 0x0f, 0xa7, 0x5d, 0xcc, 0xa4, 0xdd, 0x9a, 0x2f, 0xc5, 0x5a, 0x53, 0x74, 0x3a, 0x87, 0x9d,
	0x6c, 0x0e, 0xdf, 0xa1, 0xbe, 0xde, 0xe9, 0xbf, 0x36, 0xe5, 0x1c, 0xc8, 0x2d, 0xca, 0x5f, 0x3c,
	0x98, 0x7c, 0x62, 0x43, 0xbe, 0x4d, 0x99, 0x3f, 0x45, 0x38, 0xcc, 0xc0, 0x75, 0xf8, 0x73, 0x28,
	0xb7, 0x45, 0xcf, 0x13, 0x52, 0x07, 0x7e, 0x12, 0x05, 0xfe, 0x8a, 0x18, 0x74, 0x98, 0xeb, 0x73,
	0x8f, 0x49, 0x4b, 0x63, 0xc8, 0x3b, 0xa8, 0xf4, 0xc7, 0x76, 0xe0, 0x7a, 0x6c, 0xa4, 0x38, 0x3b,
	0x1b, 0x38, 0x19, 0x24, 0x79, 0x03, 0xd0, 0xe3, 0x7c, 0x32, 0xf3, 0x15, 0x6f, 0x77, 0x03, 0x2f,
	0x85, 0x23, 0xd7, 0x60, 0x58, 0x38, 0xe5, 0x12, 0x17, 0x08, 0xc5, 0x2c, 0x6d, 0x60, 0x2e, 0x61,
	0x9b, 0xbf, 0xcb, 0xb0, 0xd7, 0xb2, 0x9d, 0x09, 0x32, 0x97, 0x7c, 0x06, 0xa3, 0x8b, 0x32, 0xd5,
	0x01, 0x62, 0x46, 0x3e, 0x56, 0xbb, 0x68, 0x3e, 0xcb, 0xbd, 0x0b, 0x5b, 0x46, 0x0b, 0xe4, 0x03,
	0x40, 0x17, 0xa5, 0x56, 0x92, 0x1c, 0x2d, 0x0f, 0xbf, 0x76, 0x72, 0xbc, 0x62, 0x8f, 0x1d, 0x5c,
	0x29, 0x07, 0x7a, 0xda, 0x13, 0x07, 0xd9, 0xf1, 0x37, 0xe3, 0xad, 0xd2, 0xee, 0x2b, 0x7b, 0x4c,
	0xd4, 0xc3, 0x9f, 0x10, 0xb3, 0xdb, 0x20, 0x21, 0x6a, 0x3b, 0x2d, 0x90, 0x1b, 0x55, 0x7f, 0xea,
	0x35, 0x25, 0xf5, 0xaf, 0x2e, 0x08, 0x33, 0xef, 0xf9, 0xd1, 0x02, 0xf9, 0x01, 0x4f, 0xbb, 0x28,
	0x57, 0x27, 0x9c, 0xbc, 0x88, 0xf0, 0x6b, 0x57, 0x86, 0x49, 0x37, 0x41, 0xe2, 0xc6, 0x70, 0xa8,
	0xad, 0x9b, 0x18, 0xf2, 0x32, 0xf2, 0xb0, 0x65, 0x50, 0xcd, 0xc6, 0x76, 0x60, 0x1c, 0xf0, 0x3d,
	0xfc, 0xdf, 0x45, 0xd9, 0xb3, 0x25, 0x8a, 0x58, 0x8f, 0xf8, 0x75, 0xa5, 0x17, 0x62, 0x9e, 0x1a,
	0x69, 0x72, 0xa4, 0xc9, 0x16, 0x72, 0xa2, 0xc8, 0x00, 0xaa, 0x4b, 0x0b, 0x8a, 0x3c, 0x4f, 0x7a,
	0x94, 0xb7, 0xed, 0xcc, 0xd3, 0xb5, 0xf7, 0x71, 0x3d, 0x1d, 0xa8, 0xf6, 0x91, 0xb9, 0x69, 0xa1,
	0xf3, 0xc4, 0x4c, 0x5e, 0x78, 0xce, 0xd7, 0x89, 0x16, 0xc8, 0x5b, 0x28, 0xdf, 0x72, 0xe9, 0x0d,
	0xe7, 0x24, 0xce, 0xfc, 0x0b, 0x0a, 0x61, 0x8f, 0xd0, 0x3c, 0x89, 0x0c, 0x37, 0x9c, 0x0d, 0xbd,
	0x91, 0xe5, 0x3b, 0x09, 0xef, 0x2e, 0xfc, 0xde, 0x5e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe6,
	0x4a, 0xae, 0x05, 0x88, 0x07, 0x00, 0x00,
}
