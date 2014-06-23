// Code generated by protoc-gen-go.
// source: common.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	common.proto
	coordinator.proto
	ddl.proto
	index.proto
	indexer.proto
	mutation.proto
	projector.proto
	rollback.proto

It has these top-level messages:
	Error
	Timestamp
	BranchTimestamp
	Actor
	FailoverLog
	HeartBeat
	StateContextVersion
*/
package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// An actor is a process in secondary index cluster and can be one of the following.
type ActorRole int32

const (
	ActorRole_Projector        ActorRole = 1
	ActorRole_Indexer          ActorRole = 2
	ActorRole_IndexCoordinator ActorRole = 3
)

var ActorRole_name = map[int32]string{
	1: "Projector",
	2: "Indexer",
	3: "IndexCoordinator",
}
var ActorRole_value = map[string]int32{
	"Projector":        1,
	"Indexer":          2,
	"IndexCoordinator": 3,
}

func (x ActorRole) Enum() *ActorRole {
	p := new(ActorRole)
	*p = x
	return p
}
func (x ActorRole) String() string {
	return proto.EnumName(ActorRole_name, int32(x))
}
func (x *ActorRole) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ActorRole_value, data, "ActorRole")
	if err != nil {
		return err
	}
	*x = ActorRole(value)
	return nil
}

type Error struct {
	Error            *string `protobuf:"bytes,1,req,name=error" json:"error,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}

func (m *Error) GetError() string {
	if m != nil && m.Error != nil {
		return *m.Error
	}
	return ""
}

// logical clock
type Timestamp struct {
	Bucket           *string  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Vbnos            []uint32 `protobuf:"varint,2,rep,name=vbnos" json:"vbnos,omitempty"`
	Seqnos           []uint64 `protobuf:"varint,3,rep,name=seqnos" json:"seqnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Timestamp) Reset()         { *m = Timestamp{} }
func (m *Timestamp) String() string { return proto.CompactTextString(m) }
func (*Timestamp) ProtoMessage()    {}

func (m *Timestamp) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *Timestamp) GetVbnos() []uint32 {
	if m != nil {
		return m.Vbnos
	}
	return nil
}

func (m *Timestamp) GetSeqnos() []uint64 {
	if m != nil {
		return m.Seqnos
	}
	return nil
}

// logical clock that also associate a vbucket branch for the specified sequence number
type BranchTimestamp struct {
	Bucket           *string  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Vbnos            []uint32 `protobuf:"varint,2,rep,name=vbnos" json:"vbnos,omitempty"`
	Seqnos           []uint64 `protobuf:"varint,3,rep,name=seqnos" json:"seqnos,omitempty"`
	Vbuuids          []uint64 `protobuf:"varint,4,rep,name=vbuuids" json:"vbuuids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *BranchTimestamp) Reset()         { *m = BranchTimestamp{} }
func (m *BranchTimestamp) String() string { return proto.CompactTextString(m) }
func (*BranchTimestamp) ProtoMessage()    {}

func (m *BranchTimestamp) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *BranchTimestamp) GetVbnos() []uint32 {
	if m != nil {
		return m.Vbnos
	}
	return nil
}

func (m *BranchTimestamp) GetSeqnos() []uint64 {
	if m != nil {
		return m.Seqnos
	}
	return nil
}

func (m *BranchTimestamp) GetVbuuids() []uint64 {
	if m != nil {
		return m.Vbuuids
	}
	return nil
}

// Actors joining or leaving the topology. An actor is identified using its
// admin-port's connection-address <host:port>.
type Actor struct {
	Active           *bool       `protobuf:"varint,1,req,name=active" json:"active,omitempty"`
	ConnectionAddr   *string     `protobuf:"bytes,2,req,name=connectionAddr" json:"connectionAddr,omitempty"`
	Roles            []ActorRole `protobuf:"varint,3,rep,name=roles,enum=protobuf.ActorRole" json:"roles,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Actor) Reset()         { *m = Actor{} }
func (m *Actor) String() string { return proto.CompactTextString(m) }
func (*Actor) ProtoMessage()    {}

func (m *Actor) GetActive() bool {
	if m != nil && m.Active != nil {
		return *m.Active
	}
	return false
}

func (m *Actor) GetConnectionAddr() string {
	if m != nil && m.ConnectionAddr != nil {
		return *m.ConnectionAddr
	}
	return ""
}

func (m *Actor) GetRoles() []ActorRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

// failover log for a vbucket.
type FailoverLog struct {
	Vbno             *uint32  `protobuf:"varint,1,req,name=vbno" json:"vbno,omitempty"`
	Vbuuids          []uint64 `protobuf:"varint,2,rep,name=vbuuids" json:"vbuuids,omitempty"`
	Seqnos           []uint64 `protobuf:"varint,3,rep,name=seqnos" json:"seqnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FailoverLog) Reset()         { *m = FailoverLog{} }
func (m *FailoverLog) String() string { return proto.CompactTextString(m) }
func (*FailoverLog) ProtoMessage()    {}

func (m *FailoverLog) GetVbno() uint32 {
	if m != nil && m.Vbno != nil {
		return *m.Vbno
	}
	return 0
}

func (m *FailoverLog) GetVbuuids() []uint64 {
	if m != nil {
		return m.Vbuuids
	}
	return nil
}

func (m *FailoverLog) GetSeqnos() []uint64 {
	if m != nil {
		return m.Seqnos
	}
	return nil
}

// Periodic Heartbeat from indexer to Coordinator, for each bucket.
type HeartBeat struct {
	HwTimestamp         *Timestamp `protobuf:"bytes,1,req,name=hwTimestamp" json:"hwTimestamp,omitempty"`
	LastPersistenceHash *uint32    `protobuf:"varint,2,req,name=lastPersistenceHash" json:"lastPersistenceHash,omitempty"`
	LastStabilityHash   *uint32    `protobuf:"varint,3,req,name=lastStabilityHash" json:"lastStabilityHash,omitempty"`
	FreeQueue           *uint32    `protobuf:"varint,4,req,name=freeQueue" json:"freeQueue,omitempty"`
	XXX_unrecognized    []byte     `json:"-"`
}

func (m *HeartBeat) Reset()         { *m = HeartBeat{} }
func (m *HeartBeat) String() string { return proto.CompactTextString(m) }
func (*HeartBeat) ProtoMessage()    {}

func (m *HeartBeat) GetHwTimestamp() *Timestamp {
	if m != nil {
		return m.HwTimestamp
	}
	return nil
}

func (m *HeartBeat) GetLastPersistenceHash() uint32 {
	if m != nil && m.LastPersistenceHash != nil {
		return *m.LastPersistenceHash
	}
	return 0
}

func (m *HeartBeat) GetLastStabilityHash() uint32 {
	if m != nil && m.LastStabilityHash != nil {
		return *m.LastStabilityHash
	}
	return 0
}

func (m *HeartBeat) GetFreeQueue() uint32 {
	if m != nil && m.FreeQueue != nil {
		return *m.FreeQueue
	}
	return 0
}

// Version information for Coordinator's StateContext
type StateContextVersion struct {
	Cas              *uint64 `protobuf:"varint,1,req,name=cas" json:"cas,omitempty"`
	Checksum         *uint32 `protobuf:"varint,2,req,name=checksum" json:"checksum,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StateContextVersion) Reset()         { *m = StateContextVersion{} }
func (m *StateContextVersion) String() string { return proto.CompactTextString(m) }
func (*StateContextVersion) ProtoMessage()    {}

func (m *StateContextVersion) GetCas() uint64 {
	if m != nil && m.Cas != nil {
		return *m.Cas
	}
	return 0
}

func (m *StateContextVersion) GetChecksum() uint32 {
	if m != nil && m.Checksum != nil {
		return *m.Checksum
	}
	return 0
}

func init() {
	proto.RegisterEnum("protobuf.ActorRole", ActorRole_name, ActorRole_value)
}
