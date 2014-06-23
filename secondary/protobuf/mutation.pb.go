// Code generated by protoc-gen-go.
// source: mutation.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// List of possible mutation commands.
type Command int32

const (
	Command_Upsert         Command = 1
	Command_Deletion       Command = 2
	Command_UpsertDeletion Command = 3
	Command_Sync           Command = 4
	Command_DropData       Command = 5
	Command_StreamBegin    Command = 6
	Command_StreamEnd      Command = 7
)

var Command_name = map[int32]string{
	1: "Upsert",
	2: "Deletion",
	3: "UpsertDeletion",
	4: "Sync",
	5: "DropData",
	6: "StreamBegin",
	7: "StreamEnd",
}
var Command_value = map[string]int32{
	"Upsert":         1,
	"Deletion":       2,
	"UpsertDeletion": 3,
	"Sync":           4,
	"DropData":       5,
	"StreamBegin":    6,
	"StreamEnd":      7,
}

func (x Command) Enum() *Command {
	p := new(Command)
	*p = x
	return p
}
func (x Command) String() string {
	return proto.EnumName(Command_name, int32(x))
}
func (x *Command) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Command_value, data, "Command")
	if err != nil {
		return err
	}
	*x = Command(value)
	return nil
}

// A single mutation message that will framed and transported by router.
// For efficiency mutations from mutiple vbuckets (bounded to same connection)
// can be packed into the same message.
type Payload struct {
	Version *uint32 `protobuf:"varint,1,req,name=version" json:"version,omitempty"`
	// -- Following fields are mutually exclusive --
	Vbkeys           []*VbKeyVersions `protobuf:"bytes,2,rep,name=vbkeys" json:"vbkeys,omitempty"`
	Vbmap            *VbConnectionMap `protobuf:"bytes,3,opt,name=vbmap" json:"vbmap,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}

func (m *Payload) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Payload) GetVbkeys() []*VbKeyVersions {
	if m != nil {
		return m.Vbkeys
	}
	return nil
}

func (m *Payload) GetVbmap() *VbConnectionMap {
	if m != nil {
		return m.Vbmap
	}
	return nil
}

// List of vbuckets that will be streamed via a newly opened connection.
type VbConnectionMap struct {
	Bucket           *string  `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Vbuckets         []uint32 `protobuf:"varint,3,rep,name=vbuckets" json:"vbuckets,omitempty"`
	Vbuuids          []uint64 `protobuf:"varint,4,rep,name=vbuuids" json:"vbuuids,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *VbConnectionMap) Reset()         { *m = VbConnectionMap{} }
func (m *VbConnectionMap) String() string { return proto.CompactTextString(m) }
func (*VbConnectionMap) ProtoMessage()    {}

func (m *VbConnectionMap) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *VbConnectionMap) GetVbuckets() []uint32 {
	if m != nil {
		return m.Vbuckets
	}
	return nil
}

func (m *VbConnectionMap) GetVbuuids() []uint64 {
	if m != nil {
		return m.Vbuuids
	}
	return nil
}

type VbKeyVersions struct {
	Vbucket          *uint32        `protobuf:"varint,2,req,name=vbucket" json:"vbucket,omitempty"`
	Vbuuid           *uint64        `protobuf:"varint,3,req,name=vbuuid" json:"vbuuid,omitempty"`
	Bucketname       *string        `protobuf:"bytes,4,opt,name=bucketname" json:"bucketname,omitempty"`
	Kvs              []*KeyVersions `protobuf:"bytes,5,rep,name=kvs" json:"kvs,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *VbKeyVersions) Reset()         { *m = VbKeyVersions{} }
func (m *VbKeyVersions) String() string { return proto.CompactTextString(m) }
func (*VbKeyVersions) ProtoMessage()    {}

func (m *VbKeyVersions) GetVbucket() uint32 {
	if m != nil && m.Vbucket != nil {
		return *m.Vbucket
	}
	return 0
}

func (m *VbKeyVersions) GetVbuuid() uint64 {
	if m != nil && m.Vbuuid != nil {
		return *m.Vbuuid
	}
	return 0
}

func (m *VbKeyVersions) GetBucketname() string {
	if m != nil && m.Bucketname != nil {
		return *m.Bucketname
	}
	return ""
}

func (m *VbKeyVersions) GetKvs() []*KeyVersions {
	if m != nil {
		return m.Kvs
	}
	return nil
}

// mutations are broadly divided into data and control messages. The division
// is based on the commands.
//
// Interpreting seq.no:
// 1. For Upsert, Deletion, UpsertDeletion messages, sequence number corresponds
//    to kv mutation.
// 2. For Sync message, it is the latest kv mutation sequence-no. received for
//   a vbucket.
// 3. For DropData message, it is the first kv mutation that was dropped due
//    to buffer overflow.
// 4. For StreamBegin, it is the first kv mutation received after opening a
//    vbucket stream with kv.
// 5. For StreamEnd, it is the last kv mutation received before ending a vbucket
//    stream with kv.
//
// fields `docid`, `uuids`, `keys`, `oldkeys` are valid only for
// Upsert, Deletion, UpsertDeletion messages.
type KeyVersions struct {
	Seqno            *uint64  `protobuf:"varint,1,req,name=seqno" json:"seqno,omitempty"`
	Docid            []byte   `protobuf:"bytes,2,opt,name=docid" json:"docid,omitempty"`
	Uuids            []uint64 `protobuf:"varint,3,rep,name=uuids" json:"uuids,omitempty"`
	Commands         []uint32 `protobuf:"varint,4,rep,name=commands" json:"commands,omitempty"`
	Keys             [][]byte `protobuf:"bytes,5,rep,name=keys" json:"keys,omitempty"`
	Oldkeys          [][]byte `protobuf:"bytes,6,rep,name=oldkeys" json:"oldkeys,omitempty"`
	Partnkeys        [][]byte `protobuf:"bytes,7,rep,name=partnkeys" json:"partnkeys,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *KeyVersions) Reset()         { *m = KeyVersions{} }
func (m *KeyVersions) String() string { return proto.CompactTextString(m) }
func (*KeyVersions) ProtoMessage()    {}

func (m *KeyVersions) GetSeqno() uint64 {
	if m != nil && m.Seqno != nil {
		return *m.Seqno
	}
	return 0
}

func (m *KeyVersions) GetDocid() []byte {
	if m != nil {
		return m.Docid
	}
	return nil
}

func (m *KeyVersions) GetUuids() []uint64 {
	if m != nil {
		return m.Uuids
	}
	return nil
}

func (m *KeyVersions) GetCommands() []uint32 {
	if m != nil {
		return m.Commands
	}
	return nil
}

func (m *KeyVersions) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

func (m *KeyVersions) GetOldkeys() [][]byte {
	if m != nil {
		return m.Oldkeys
	}
	return nil
}

func (m *KeyVersions) GetPartnkeys() [][]byte {
	if m != nil {
		return m.Partnkeys
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.Command", Command_name, Command_value)
}
