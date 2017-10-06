// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/server/server.proto

/*
Package server is a generated protocol buffer package.

It is generated from these files:
	proto/server/server.proto

It has these top-level messages:
	CommandRequest
	CommandResponse
	LogEntry
	RequestVoteRequest
	RequestVoteResponse
	AppendEntriesRequest
	AppendEntriesResponse
	Peer
	Node
	RaftGroup
	ServerConfig
	ApproveAppendResponse
	Client
	Nothing
*/
package server

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

type CommandRequest struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	ClientId  uint64 `protobuf:"varint,2,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	RequestId uint64 `protobuf:"varint,3,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	FuncId    uint64 `protobuf:"varint,4,opt,name=func_id,json=funcId" json:"func_id,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Arg       []byte `protobuf:"bytes,6,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (m *CommandRequest) Reset()                    { *m = CommandRequest{} }
func (m *CommandRequest) String() string            { return proto.CompactTextString(m) }
func (*CommandRequest) ProtoMessage()               {}
func (*CommandRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommandRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *CommandRequest) GetClientId() uint64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func (m *CommandRequest) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *CommandRequest) GetFuncId() uint64 {
	if m != nil {
		return m.FuncId
	}
	return 0
}

func (m *CommandRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CommandRequest) GetArg() []byte {
	if m != nil {
		return m.Arg
	}
	return nil
}

type CommandResponse struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	LeaderId  uint64 `protobuf:"varint,2,opt,name=leader_id,json=leaderId" json:"leader_id,omitempty"`
	NodeId    uint64 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	RequestId uint64 `protobuf:"varint,4,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	Signature []byte `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Result    []byte `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
}

func (m *CommandResponse) Reset()                    { *m = CommandResponse{} }
func (m *CommandResponse) String() string            { return proto.CompactTextString(m) }
func (*CommandResponse) ProtoMessage()               {}
func (*CommandResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommandResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *CommandResponse) GetLeaderId() uint64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *CommandResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *CommandResponse) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *CommandResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *CommandResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

type LogEntry struct {
	Term    uint64          `protobuf:"varint,1,opt,name=term" json:"term,omitempty"`
	Index   uint64          `protobuf:"varint,2,opt,name=index" json:"index,omitempty"`
	Hash    []byte          `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Command *CommandRequest `protobuf:"bytes,4,opt,name=command" json:"command,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LogEntry) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *LogEntry) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *LogEntry) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *LogEntry) GetCommand() *CommandRequest {
	if m != nil {
		return m.Command
	}
	return nil
}

type RequestVoteRequest struct {
	Group       uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	LogIndex    uint64 `protobuf:"varint,3,opt,name=log_index,json=logIndex" json:"log_index,omitempty"`
	LogTerm     uint64 `protobuf:"varint,4,opt,name=log_term,json=logTerm" json:"log_term,omitempty"`
	CandidateId uint64 `protobuf:"varint,5,opt,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	Signature   []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RequestVoteRequest) Reset()                    { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()               {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RequestVoteRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetLogIndex() uint64 {
	if m != nil {
		return m.LogIndex
	}
	return 0
}

func (m *RequestVoteRequest) GetLogTerm() uint64 {
	if m != nil {
		return m.LogTerm
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidateId() uint64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVoteRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type RequestVoteResponse struct {
	Group       uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term        uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	NodeId      uint64 `protobuf:"varint,3,opt,name=node_id,json=nodeId" json:"node_id,omitempty"`
	CandidateId uint64 `protobuf:"varint,4,opt,name=candidate_id,json=candidateId" json:"candidate_id,omitempty"`
	Granted     bool   `protobuf:"varint,5,opt,name=granted" json:"granted,omitempty"`
	Signature   []byte `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *RequestVoteResponse) Reset()                    { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()               {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestVoteResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *RequestVoteResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteResponse) GetNodeId() uint64 {
	if m != nil {
		return m.NodeId
	}
	return 0
}

func (m *RequestVoteResponse) GetCandidateId() uint64 {
	if m != nil {
		return m.CandidateId
	}
	return 0
}

func (m *RequestVoteResponse) GetGranted() bool {
	if m != nil {
		return m.Granted
	}
	return false
}

func (m *RequestVoteResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type AppendEntriesRequest struct {
	Group        uint64                 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term         uint64                 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	LeaderId     uint64                 `protobuf:"varint,3,opt,name=leader_id,json=leaderId" json:"leader_id,omitempty"`
	PrevLogIndex uint64                 `protobuf:"varint,4,opt,name=prev_log_index,json=prevLogIndex" json:"prev_log_index,omitempty"`
	PrevLogTerm  uint64                 `protobuf:"varint,5,opt,name=prev_log_term,json=prevLogTerm" json:"prev_log_term,omitempty"`
	Signature    []byte                 `protobuf:"bytes,6,opt,name=signature,proto3" json:"signature,omitempty"`
	QuorumVotes  []*RequestVoteResponse `protobuf:"bytes,7,rep,name=quorum_votes,json=quorumVotes" json:"quorum_votes,omitempty"`
	Entries      []*LogEntry            `protobuf:"bytes,8,rep,name=entries" json:"entries,omitempty"`
}

func (m *AppendEntriesRequest) Reset()                    { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()               {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *AppendEntriesRequest) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeaderId() uint64 {
	if m != nil {
		return m.LeaderId
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogIndex() uint64 {
	if m != nil {
		return m.PrevLogIndex
	}
	return 0
}

func (m *AppendEntriesRequest) GetPrevLogTerm() uint64 {
	if m != nil {
		return m.PrevLogTerm
	}
	return 0
}

func (m *AppendEntriesRequest) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *AppendEntriesRequest) GetQuorumVotes() []*RequestVoteResponse {
	if m != nil {
		return m.QuorumVotes
	}
	return nil
}

func (m *AppendEntriesRequest) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type AppendEntriesResponse struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Term      uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	Index     uint64 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	Peer      uint64 `protobuf:"varint,4,opt,name=peer" json:"peer,omitempty"`
	Successed bool   `protobuf:"varint,5,opt,name=successed" json:"successed,omitempty"`
	Convinced bool   `protobuf:"varint,6,opt,name=convinced" json:"convinced,omitempty"`
	Hash      []byte `protobuf:"bytes,7,opt,name=hash,proto3" json:"hash,omitempty"`
	Signature []byte `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *AppendEntriesResponse) Reset()                    { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()               {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AppendEntriesResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *AppendEntriesResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResponse) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *AppendEntriesResponse) GetPeer() uint64 {
	if m != nil {
		return m.Peer
	}
	return 0
}

func (m *AppendEntriesResponse) GetSuccessed() bool {
	if m != nil {
		return m.Successed
	}
	return false
}

func (m *AppendEntriesResponse) GetConvinced() bool {
	if m != nil {
		return m.Convinced
	}
	return false
}

func (m *AppendEntriesResponse) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *AppendEntriesResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Peer struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Group      uint64 `protobuf:"varint,2,opt,name=group" json:"group,omitempty"`
	Host       uint64 `protobuf:"varint,3,opt,name=host" json:"host,omitempty"`
	NextIndex  uint64 `protobuf:"varint,4,opt,name=next_index,json=nextIndex" json:"next_index,omitempty"`
	MatchIndex uint64 `protobuf:"varint,5,opt,name=match_index,json=matchIndex" json:"match_index,omitempty"`
	LastSeen   uint64 `protobuf:"varint,6,opt,name=last_seen,json=lastSeen" json:"last_seen,omitempty"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Peer) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Peer) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *Peer) GetHost() uint64 {
	if m != nil {
		return m.Host
	}
	return 0
}

func (m *Peer) GetNextIndex() uint64 {
	if m != nil {
		return m.NextIndex
	}
	return 0
}

func (m *Peer) GetMatchIndex() uint64 {
	if m != nil {
		return m.MatchIndex
	}
	return 0
}

func (m *Peer) GetLastSeen() uint64 {
	if m != nil {
		return m.LastSeen
	}
	return 0
}

type Node struct {
	Id         uint64   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	LastSeen   uint64   `protobuf:"varint,2,opt,name=last_seen,json=lastSeen" json:"last_seen,omitempty"`
	Online     bool     `protobuf:"varint,3,opt,name=online" json:"online,omitempty"`
	ServerAddr string   `protobuf:"bytes,4,opt,name=server_addr,json=serverAddr" json:"server_addr,omitempty"`
	PublicKey  []byte   `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	Peers      []uint64 `protobuf:"varint,6,rep,packed,name=peers" json:"peers,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Node) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Node) GetLastSeen() uint64 {
	if m != nil {
		return m.LastSeen
	}
	return 0
}

func (m *Node) GetOnline() bool {
	if m != nil {
		return m.Online
	}
	return false
}

func (m *Node) GetServerAddr() string {
	if m != nil {
		return m.ServerAddr
	}
	return ""
}

func (m *Node) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *Node) GetPeers() []uint64 {
	if m != nil {
		return m.Peers
	}
	return nil
}

type RaftGroup struct {
	Replications uint32 `protobuf:"varint,1,opt,name=replications" json:"replications,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	LeaderPeer   uint64 `protobuf:"varint,3,opt,name=leader_peer,json=leaderPeer" json:"leader_peer,omitempty"`
	Term         uint64 `protobuf:"varint,4,opt,name=term" json:"term,omitempty"`
}

func (m *RaftGroup) Reset()                    { *m = RaftGroup{} }
func (m *RaftGroup) String() string            { return proto.CompactTextString(m) }
func (*RaftGroup) ProtoMessage()               {}
func (*RaftGroup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RaftGroup) GetReplications() uint32 {
	if m != nil {
		return m.Replications
	}
	return 0
}

func (m *RaftGroup) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RaftGroup) GetLeaderPeer() uint64 {
	if m != nil {
		return m.LeaderPeer
	}
	return 0
}

func (m *RaftGroup) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

type ServerConfig struct {
	PublicKey []byte `protobuf:"bytes,1,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ServerConfig) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type ApproveAppendResponse struct {
	Group     uint64 `protobuf:"varint,1,opt,name=group" json:"group,omitempty"`
	Peer      uint64 `protobuf:"varint,2,opt,name=peer" json:"peer,omitempty"`
	Index     uint64 `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
	Appended  bool   `protobuf:"varint,4,opt,name=appended" json:"appended,omitempty"`
	Delayed   bool   `protobuf:"varint,5,opt,name=delayed" json:"delayed,omitempty"`
	Failed    bool   `protobuf:"varint,6,opt,name=failed" json:"failed,omitempty"`
	Signature []byte `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *ApproveAppendResponse) Reset()                    { *m = ApproveAppendResponse{} }
func (m *ApproveAppendResponse) String() string            { return proto.CompactTextString(m) }
func (*ApproveAppendResponse) ProtoMessage()               {}
func (*ApproveAppendResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ApproveAppendResponse) GetGroup() uint64 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *ApproveAppendResponse) GetPeer() uint64 {
	if m != nil {
		return m.Peer
	}
	return 0
}

func (m *ApproveAppendResponse) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *ApproveAppendResponse) GetAppended() bool {
	if m != nil {
		return m.Appended
	}
	return false
}

func (m *ApproveAppendResponse) GetDelayed() bool {
	if m != nil {
		return m.Delayed
	}
	return false
}

func (m *ApproveAppendResponse) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *ApproveAppendResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Client struct {
	Id         uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Address    string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	PrivateKey []byte `protobuf:"bytes,3,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
}

func (m *Client) Reset()                    { *m = Client{} }
func (m *Client) String() string            { return proto.CompactTextString(m) }
func (*Client) ProtoMessage()               {}
func (*Client) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Client) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Client) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Client) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

type Nothing struct {
}

func (m *Nothing) Reset()                    { *m = Nothing{} }
func (m *Nothing) String() string            { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()               {}
func (*Nothing) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func init() {
	proto.RegisterType((*CommandRequest)(nil), "server.CommandRequest")
	proto.RegisterType((*CommandResponse)(nil), "server.CommandResponse")
	proto.RegisterType((*LogEntry)(nil), "server.LogEntry")
	proto.RegisterType((*RequestVoteRequest)(nil), "server.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "server.RequestVoteResponse")
	proto.RegisterType((*AppendEntriesRequest)(nil), "server.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "server.AppendEntriesResponse")
	proto.RegisterType((*Peer)(nil), "server.Peer")
	proto.RegisterType((*Node)(nil), "server.Node")
	proto.RegisterType((*RaftGroup)(nil), "server.RaftGroup")
	proto.RegisterType((*ServerConfig)(nil), "server.ServerConfig")
	proto.RegisterType((*ApproveAppendResponse)(nil), "server.ApproveAppendResponse")
	proto.RegisterType((*Client)(nil), "server.Client")
	proto.RegisterType((*Nothing)(nil), "server.Nothing")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for BFTRaft service

type BFTRaftClient interface {
	ExecCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error)
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
	ApproveAppend(ctx context.Context, in *AppendEntriesResponse, opts ...grpc.CallOption) (*ApproveAppendResponse, error)
}

type bFTRaftClient struct {
	cc *grpc.ClientConn
}

func NewBFTRaftClient(cc *grpc.ClientConn) BFTRaftClient {
	return &bFTRaftClient{cc}
}

func (c *bFTRaftClient) ExecCommand(ctx context.Context, in *CommandRequest, opts ...grpc.CallOption) (*CommandResponse, error) {
	out := new(CommandResponse)
	err := grpc.Invoke(ctx, "/server.BFTRaft/ExecCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFTRaftClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := grpc.Invoke(ctx, "/server.BFTRaft/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFTRaftClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := grpc.Invoke(ctx, "/server.BFTRaft/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bFTRaftClient) ApproveAppend(ctx context.Context, in *AppendEntriesResponse, opts ...grpc.CallOption) (*ApproveAppendResponse, error) {
	out := new(ApproveAppendResponse)
	err := grpc.Invoke(ctx, "/server.BFTRaft/ApproveAppend", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BFTRaft service

type BFTRaftServer interface {
	ExecCommand(context.Context, *CommandRequest) (*CommandResponse, error)
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
	ApproveAppend(context.Context, *AppendEntriesResponse) (*ApproveAppendResponse, error)
}

func RegisterBFTRaftServer(s *grpc.Server, srv BFTRaftServer) {
	s.RegisterService(&_BFTRaft_serviceDesc, srv)
}

func _BFTRaft_ExecCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).ExecCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.BFTRaft/ExecCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).ExecCommand(ctx, req.(*CommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFTRaft_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.BFTRaft/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFTRaft_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.BFTRaft/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BFTRaft_ApproveAppend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BFTRaftServer).ApproveAppend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/server.BFTRaft/ApproveAppend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BFTRaftServer).ApproveAppend(ctx, req.(*AppendEntriesResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _BFTRaft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "server.BFTRaft",
	HandlerType: (*BFTRaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecCommand",
			Handler:    _BFTRaft_ExecCommand_Handler,
		},
		{
			MethodName: "RequestVote",
			Handler:    _BFTRaft_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _BFTRaft_AppendEntries_Handler,
		},
		{
			MethodName: "ApproveAppend",
			Handler:    _BFTRaft_ApproveAppend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/server/server.proto",
}

func init() { proto.RegisterFile("proto/server/server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 940 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdf, 0x6e, 0x1b, 0xc5,
	0x17, 0xce, 0xae, 0x1d, 0xaf, 0x7d, 0xec, 0xe4, 0x57, 0xcd, 0xaf, 0xb4, 0x6e, 0xda, 0xd2, 0xb0,
	0xe2, 0x22, 0x42, 0xa2, 0xa0, 0x70, 0x8f, 0x28, 0x51, 0x81, 0xa8, 0x55, 0x40, 0xdb, 0x8a, 0x5b,
	0x6b, 0xba, 0x73, 0xb2, 0x5e, 0x61, 0xcf, 0x6c, 0x67, 0x66, 0xad, 0x98, 0x77, 0xe1, 0x82, 0x6b,
	0xae, 0x00, 0x09, 0x1e, 0x80, 0xa7, 0xe0, 0x6d, 0xd0, 0xfc, 0xd9, 0x5d, 0xef, 0xda, 0x71, 0x05,
	0x57, 0x99, 0xf3, 0xcd, 0x9c, 0xc9, 0x77, 0xce, 0xf7, 0xcd, 0x59, 0xc3, 0x83, 0x42, 0x0a, 0x2d,
	0x3e, 0x51, 0x28, 0x57, 0x28, 0xfd, 0x9f, 0xa7, 0x16, 0x23, 0x03, 0x17, 0xc5, 0xbf, 0x04, 0x70,
	0x7c, 0x21, 0x96, 0x4b, 0xca, 0x59, 0x82, 0x6f, 0x4b, 0x54, 0x9a, 0xdc, 0x85, 0xc3, 0x4c, 0x8a,
	0xb2, 0x98, 0x06, 0xa7, 0xc1, 0x59, 0x3f, 0x71, 0x01, 0x79, 0x08, 0xa3, 0x74, 0x91, 0x23, 0xd7,
	0xb3, 0x9c, 0x4d, 0x43, 0xbb, 0x33, 0x74, 0xc0, 0x25, 0x23, 0x8f, 0x01, 0xa4, 0xcb, 0x36, 0xbb,
	0x3d, 0xbb, 0x3b, 0xf2, 0xc8, 0x25, 0x23, 0xf7, 0x21, 0xba, 0x2e, 0x79, 0x6a, 0xf6, 0xfa, 0x76,
	0x6f, 0x60, 0xc2, 0x4b, 0x46, 0x1e, 0xc1, 0x48, 0xe5, 0x19, 0xa7, 0xba, 0x94, 0x38, 0x3d, 0x3c,
	0x0d, 0xce, 0x26, 0x49, 0x03, 0x90, 0x3b, 0xd0, 0xa3, 0x32, 0x9b, 0x0e, 0x2c, 0x6e, 0x96, 0xf1,
	0x6f, 0x01, 0xfc, 0xaf, 0x66, 0xab, 0x0a, 0xc1, 0x15, 0xde, 0x4e, 0x77, 0x81, 0x94, 0xa1, 0xdc,
	0xa0, 0xeb, 0x00, 0xc7, 0x87, 0x0b, 0x86, 0x0d, 0xd7, 0x81, 0x09, 0xb7, 0xea, 0xe8, 0x77, 0xeb,
	0xd8, 0x4f, 0xf7, 0x1e, 0x0c, 0x24, 0xaa, 0x72, 0xa1, 0x3d, 0x63, 0x1f, 0xc5, 0x3f, 0xc2, 0xf0,
	0xa5, 0xc8, 0x9e, 0x73, 0x2d, 0xd7, 0x84, 0x40, 0x5f, 0xa3, 0x5c, 0x7a, 0xae, 0x76, 0x6d, 0x0a,
	0xc8, 0x39, 0xc3, 0x1b, 0x4f, 0xd3, 0x05, 0xe6, 0xe4, 0x9c, 0xaa, 0xb9, 0x25, 0x38, 0x49, 0xec,
	0x9a, 0x7c, 0x0a, 0x51, 0xea, 0xaa, 0xb7, 0xdc, 0xc6, 0xe7, 0xf7, 0x9e, 0x7a, 0x51, 0xdb, 0x12,
	0x26, 0xd5, 0xb1, 0xf8, 0xcf, 0x00, 0x88, 0x07, 0xbf, 0x17, 0x1a, 0xf7, 0x4b, 0x5c, 0x91, 0x0b,
	0x37, 0xc8, 0x99, 0x3e, 0x8a, 0x6c, 0xe6, 0x08, 0xf6, 0x7c, 0x1f, 0x45, 0x76, 0x69, 0x39, 0x3e,
	0x00, 0xb3, 0x9e, 0xd9, 0x24, 0xd7, 0xac, 0x68, 0x21, 0xb2, 0xd7, 0x26, 0xef, 0x03, 0x98, 0xa4,
	0x94, 0xb3, 0x9c, 0x51, 0x6d, 0xfb, 0x7c, 0x68, 0xb7, 0xc7, 0x35, 0xd6, 0xed, 0xe6, 0xa0, 0xd3,
	0xcd, 0xf8, 0xf7, 0x00, 0xfe, 0xdf, 0x62, 0xbe, 0x57, 0xee, 0x5d, 0xd4, 0x6f, 0x55, 0xb9, 0xcb,
	0xad, 0xbf, 0xcd, 0x6d, 0x0a, 0x51, 0x26, 0x29, 0xd7, 0xe8, 0x98, 0x0f, 0x93, 0x2a, 0x7c, 0x07,
	0xeb, 0x5f, 0x43, 0xb8, 0xfb, 0xac, 0x28, 0x90, 0x33, 0xa3, 0x77, 0x8e, 0xea, 0xbf, 0x75, 0xbc,
	0x76, 0x6e, 0xaf, 0xe3, 0xdc, 0x0f, 0xe1, 0xb8, 0x90, 0xb8, 0x9a, 0x35, 0x9a, 0x38, 0xf2, 0x13,
	0x83, 0xbe, 0xac, 0x74, 0x89, 0xe1, 0xa8, 0x3e, 0x65, 0xef, 0xf7, 0xdd, 0xf7, 0x87, 0xac, 0x40,
	0x7b, 0xeb, 0x20, 0x9f, 0xc3, 0xe4, 0x6d, 0x29, 0x64, 0xb9, 0x9c, 0xad, 0x84, 0x46, 0x35, 0x8d,
	0x4e, 0x7b, 0x67, 0xe3, 0xf3, 0x87, 0x95, 0xdd, 0x76, 0x08, 0x93, 0x8c, 0x5d, 0x82, 0xc1, 0x14,
	0xf9, 0x08, 0x22, 0x74, 0x0d, 0x98, 0x0e, 0x6d, 0xea, 0x9d, 0x2a, 0xb5, 0x7a, 0x0a, 0x49, 0x75,
	0x20, 0xfe, 0x3b, 0x80, 0xf7, 0x3a, 0x3d, 0xfb, 0xd7, 0x5a, 0xd7, 0x6f, 0xa8, 0xd7, 0x79, 0x43,
	0x05, 0xa2, 0xf4, 0x3d, 0xb2, 0x6b, 0x5b, 0x77, 0x99, 0xa6, 0xa8, 0x54, 0xad, 0x6d, 0x03, 0x98,
	0xdd, 0x54, 0xf0, 0x55, 0xce, 0x53, 0x64, 0xb6, 0x2b, 0xc3, 0xa4, 0x01, 0xea, 0x37, 0x19, 0x6d,
	0xbc, 0xc9, 0x56, 0x1f, 0x87, 0x5d, 0x3f, 0xfc, 0x14, 0x40, 0xff, 0x3b, 0xf3, 0x6f, 0x8f, 0x21,
	0xcc, 0x99, 0xaf, 0x23, 0xcc, 0x59, 0x53, 0x5a, 0xd8, 0x29, 0x6d, 0x2e, 0x94, 0xf6, 0x55, 0xd8,
	0xb5, 0x99, 0x49, 0x1c, 0x6f, 0x74, 0x4b, 0xee, 0x91, 0x41, 0x9c, 0xd6, 0x4f, 0x60, 0xbc, 0xa4,
	0x3a, 0x9d, 0xfb, 0x7d, 0xa7, 0x34, 0x58, 0xc8, 0x1d, 0x30, 0x7e, 0xa2, 0x4a, 0xcf, 0x14, 0x22,
	0xb7, 0x25, 0x19, 0x3f, 0x51, 0xa5, 0x5f, 0x21, 0xf2, 0xf8, 0xe7, 0x00, 0xfa, 0x57, 0x82, 0xe1,
	0x16, 0xbf, 0x56, 0x56, 0xd8, 0xce, 0x32, 0x93, 0x4e, 0xf0, 0x45, 0xce, 0xd1, 0x12, 0x1d, 0x26,
	0x3e, 0x32, 0x5c, 0x9c, 0xca, 0x33, 0xca, 0x98, 0x6b, 0xfb, 0x28, 0x01, 0x07, 0x3d, 0x63, 0x4c,
	0x9a, 0x5a, 0x8a, 0xf2, 0xcd, 0x22, 0x4f, 0x67, 0x3f, 0xe0, 0xba, 0x9a, 0xa0, 0x0e, 0x79, 0x81,
	0x6b, 0xd3, 0x14, 0xa3, 0x91, 0x9a, 0x0e, 0x4e, 0x7b, 0xa6, 0x29, 0x36, 0x88, 0x35, 0x8c, 0x12,
	0x7a, 0xad, 0xbf, 0xb6, 0x1d, 0x8a, 0x61, 0x22, 0xb1, 0x58, 0xe4, 0x29, 0xd5, 0xb9, 0xe0, 0xca,
	0x32, 0x3e, 0x4a, 0x5a, 0x98, 0xaf, 0x25, 0xac, 0x6b, 0x79, 0x02, 0x63, 0xff, 0xa2, 0xac, 0x1b,
	0x5c, 0x73, 0xc1, 0x41, 0x56, 0x9c, 0xca, 0x51, 0xfd, 0xc6, 0x51, 0xf1, 0xc7, 0x30, 0x79, 0x65,
	0x89, 0x5f, 0x08, 0x7e, 0x9d, 0x67, 0x1d, 0xea, 0x41, 0x87, 0x7a, 0xfc, 0x97, 0x33, 0xb1, 0x14,
	0x2b, 0x74, 0x5e, 0x7e, 0xb7, 0x89, 0x2d, 0x99, 0x70, 0xc3, 0x9a, 0xbb, 0x4d, 0x7c, 0x02, 0x43,
	0x6a, 0x6f, 0x44, 0x37, 0xa9, 0x86, 0x49, 0x1d, 0x9b, 0x31, 0xc5, 0x70, 0x41, 0xd7, 0xcd, 0x98,
	0xf2, 0xa1, 0x91, 0xe8, 0x9a, 0xe6, 0x8b, 0xda, 0xc5, 0x3e, 0x6a, 0xdb, 0x35, 0xea, 0xda, 0x35,
	0x81, 0xc1, 0x85, 0xfd, 0xa6, 0x6f, 0xf9, 0x61, 0x0a, 0x91, 0xd1, 0x14, 0x95, 0xb2, 0x94, 0x47,
	0x49, 0x15, 0x92, 0xf7, 0x01, 0x0a, 0x99, 0xaf, 0xa8, 0xc6, 0x17, 0xb8, 0xf6, 0x9f, 0xab, 0x0d,
	0x24, 0x1e, 0x41, 0x74, 0x25, 0xf4, 0x3c, 0xe7, 0xd9, 0xf9, 0x1f, 0x21, 0x44, 0x5f, 0x7e, 0xf5,
	0xda, 0xa8, 0x49, 0xbe, 0x80, 0xf1, 0xf3, 0x1b, 0x4c, 0xfd, 0x87, 0x8b, 0xdc, 0xf2, 0x25, 0x3b,
	0xb9, 0xbf, 0x85, 0xbb, 0xb6, 0xc6, 0x07, 0xe4, 0x1b, 0x18, 0x6f, 0xcc, 0x21, 0x72, 0xb2, 0x73,
	0x38, 0xb9, 0x5b, 0xf6, 0x0d, 0xae, 0xf8, 0x80, 0x5c, 0xc1, 0x51, 0x6b, 0x00, 0x91, 0x47, 0xd5,
	0xf9, 0x5d, 0xb3, 0xfc, 0xe4, 0xf1, 0x2d, 0xbb, 0xf5, 0x7d, 0xdf, 0xda, 0xfb, 0x1a, 0x2f, 0x90,
	0xfd, 0x19, 0xad, 0x0b, 0xb7, 0x1d, 0x14, 0x1f, 0xbc, 0x19, 0xd8, 0x1f, 0x6d, 0x9f, 0xfd, 0x13,
	0x00, 0x00, 0xff, 0xff, 0xb8, 0x3d, 0x01, 0xf3, 0xd1, 0x09, 0x00, 0x00,
}
