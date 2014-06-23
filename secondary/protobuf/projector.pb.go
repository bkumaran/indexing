// Code generated by protoc-gen-go.
// source: projector.proto
// DO NOT EDIT!

package protobuf

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Requested by Coordinator during system-start, re-connect, rollback
type FailoverLogRequest struct {
	Pool             *string  `protobuf:"bytes,1,req,name=pool" json:"pool,omitempty"`
	Bucket           *string  `protobuf:"bytes,2,req,name=bucket" json:"bucket,omitempty"`
	Vbnos            []uint32 `protobuf:"varint,3,rep,name=vbnos" json:"vbnos,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *FailoverLogRequest) Reset()         { *m = FailoverLogRequest{} }
func (m *FailoverLogRequest) String() string { return proto.CompactTextString(m) }
func (*FailoverLogRequest) ProtoMessage()    {}

func (m *FailoverLogRequest) GetPool() string {
	if m != nil && m.Pool != nil {
		return *m.Pool
	}
	return ""
}

func (m *FailoverLogRequest) GetBucket() string {
	if m != nil && m.Bucket != nil {
		return *m.Bucket
	}
	return ""
}

func (m *FailoverLogRequest) GetVbnos() []uint32 {
	if m != nil {
		return m.Vbnos
	}
	return nil
}

type FailoverLogResponse struct {
	Logs             []*FailoverLog `protobuf:"bytes,1,rep,name=logs" json:"logs,omitempty"`
	Err              *Error         `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *FailoverLogResponse) Reset()         { *m = FailoverLogResponse{} }
func (m *FailoverLogResponse) String() string { return proto.CompactTextString(m) }
func (*FailoverLogResponse) ProtoMessage()    {}

func (m *FailoverLogResponse) GetLogs() []*FailoverLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *FailoverLogResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

// Requested by Coordinator or indexer to start a new mutation stream.
// BranchTimestamp.Vbnos should be in sort order
type MutationStreamRequest struct {
	Topic             *string            `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Flag              *uint32            `protobuf:"varint,2,req,name=flag" json:"flag,omitempty"`
	Pools             []string           `protobuf:"bytes,3,rep,name=pools" json:"pools,omitempty"`
	Buckets           []string           `protobuf:"bytes,4,rep,name=buckets" json:"buckets,omitempty"`
	RestartTimestamps []*BranchTimestamp `protobuf:"bytes,5,rep,name=restartTimestamps" json:"restartTimestamps,omitempty"`
	// list of index applicable for this stream, optional as well
	Instances        []*IndexInst `protobuf:"bytes,6,rep,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *MutationStreamRequest) Reset()         { *m = MutationStreamRequest{} }
func (m *MutationStreamRequest) String() string { return proto.CompactTextString(m) }
func (*MutationStreamRequest) ProtoMessage()    {}

func (m *MutationStreamRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *MutationStreamRequest) GetFlag() uint32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *MutationStreamRequest) GetPools() []string {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *MutationStreamRequest) GetBuckets() []string {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *MutationStreamRequest) GetRestartTimestamps() []*BranchTimestamp {
	if m != nil {
		return m.RestartTimestamps
	}
	return nil
}

func (m *MutationStreamRequest) GetInstances() []*IndexInst {
	if m != nil {
		return m.Instances
	}
	return nil
}

type MutationStreamResponse struct {
	Topic   *string  `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Flag    *uint32  `protobuf:"varint,2,req,name=flag" json:"flag,omitempty"`
	Pools   []string `protobuf:"bytes,3,rep,name=pools" json:"pools,omitempty"`
	Buckets []string `protobuf:"bytes,4,rep,name=buckets" json:"buckets,omitempty"`
	// per bucket failover-timestamp, kv-timestamp for all active vbuckets,
	// for each bucket, after executing the request.
	FailoverTimestamps []*BranchTimestamp `protobuf:"bytes,5,rep,name=failoverTimestamps" json:"failoverTimestamps,omitempty"`
	KvTimestamps       []*BranchTimestamp `protobuf:"bytes,6,rep,name=kvTimestamps" json:"kvTimestamps,omitempty"`
	IndexUuids         []uint64           `protobuf:"varint,7,rep,name=indexUuids" json:"indexUuids,omitempty"`
	Err                *Error             `protobuf:"bytes,8,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized   []byte             `json:"-"`
}

func (m *MutationStreamResponse) Reset()         { *m = MutationStreamResponse{} }
func (m *MutationStreamResponse) String() string { return proto.CompactTextString(m) }
func (*MutationStreamResponse) ProtoMessage()    {}

func (m *MutationStreamResponse) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *MutationStreamResponse) GetFlag() uint32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *MutationStreamResponse) GetPools() []string {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *MutationStreamResponse) GetBuckets() []string {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *MutationStreamResponse) GetFailoverTimestamps() []*BranchTimestamp {
	if m != nil {
		return m.FailoverTimestamps
	}
	return nil
}

func (m *MutationStreamResponse) GetKvTimestamps() []*BranchTimestamp {
	if m != nil {
		return m.KvTimestamps
	}
	return nil
}

func (m *MutationStreamResponse) GetIndexUuids() []uint64 {
	if m != nil {
		return m.IndexUuids
	}
	return nil
}

func (m *MutationStreamResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

// Requested by Coordinator or indexer to restart or shutdown vbuckets from an
// active mutation stream. Returns back MutationStreamResponse.
type UpdateMutationStreamRequest struct {
	Topic             *string            `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Flag              *uint32            `protobuf:"varint,2,req,name=flag" json:"flag,omitempty"`
	Pools             []string           `protobuf:"bytes,3,rep,name=pools" json:"pools,omitempty"`
	Buckets           []string           `protobuf:"bytes,4,rep,name=buckets" json:"buckets,omitempty"`
	RestartTimestamps []*BranchTimestamp `protobuf:"bytes,5,rep,name=restartTimestamps" json:"restartTimestamps,omitempty"`
	Instances         []*IndexInst       `protobuf:"bytes,6,rep,name=instances" json:"instances,omitempty"`
	XXX_unrecognized  []byte             `json:"-"`
}

func (m *UpdateMutationStreamRequest) Reset()         { *m = UpdateMutationStreamRequest{} }
func (m *UpdateMutationStreamRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateMutationStreamRequest) ProtoMessage()    {}

func (m *UpdateMutationStreamRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *UpdateMutationStreamRequest) GetFlag() uint32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *UpdateMutationStreamRequest) GetPools() []string {
	if m != nil {
		return m.Pools
	}
	return nil
}

func (m *UpdateMutationStreamRequest) GetBuckets() []string {
	if m != nil {
		return m.Buckets
	}
	return nil
}

func (m *UpdateMutationStreamRequest) GetRestartTimestamps() []*BranchTimestamp {
	if m != nil {
		return m.RestartTimestamps
	}
	return nil
}

func (m *UpdateMutationStreamRequest) GetInstances() []*IndexInst {
	if m != nil {
		return m.Instances
	}
	return nil
}

// Requested by third party component that wants to subscribe to a topic-name.
// Error message will be sent as response
type SubscribeStreamRequest struct {
	Topic            *string      `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	Flag             *uint32      `protobuf:"varint,2,req,name=flag" json:"flag,omitempty"`
	Instances        []*IndexInst `protobuf:"bytes,3,rep,name=instances" json:"instances,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *SubscribeStreamRequest) Reset()         { *m = SubscribeStreamRequest{} }
func (m *SubscribeStreamRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeStreamRequest) ProtoMessage()    {}

func (m *SubscribeStreamRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

func (m *SubscribeStreamRequest) GetFlag() uint32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func (m *SubscribeStreamRequest) GetInstances() []*IndexInst {
	if m != nil {
		return m.Instances
	}
	return nil
}

// Requested by indexer / coordinator to inform router to re-connect with
// downstream endpoint. Error message will be sent as response.
type RepairDownstreamEndpoints struct {
	Topic            *string `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RepairDownstreamEndpoints) Reset()         { *m = RepairDownstreamEndpoints{} }
func (m *RepairDownstreamEndpoints) String() string { return proto.CompactTextString(m) }
func (*RepairDownstreamEndpoints) ProtoMessage()    {}

func (m *RepairDownstreamEndpoints) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

// Requested by coordinator to should down a mutation stream and all KV
// connections active for that stream. Error message will be sent as response.
type ShutdownStreamRequest struct {
	Topic            *string `protobuf:"bytes,1,req,name=topic" json:"topic,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ShutdownStreamRequest) Reset()         { *m = ShutdownStreamRequest{} }
func (m *ShutdownStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ShutdownStreamRequest) ProtoMessage()    {}

func (m *ShutdownStreamRequest) GetTopic() string {
	if m != nil && m.Topic != nil {
		return *m.Topic
	}
	return ""
}

// Requested by Coordinator during bootstrap handshake to get the current list
// of active streams from projector
type ActiveStreamRequest struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ActiveStreamRequest) Reset()         { *m = ActiveStreamRequest{} }
func (m *ActiveStreamRequest) String() string { return proto.CompactTextString(m) }
func (*ActiveStreamRequest) ProtoMessage()    {}

type ActiveStreamResponse struct {
	Streams          []*MutationStreamResponse `protobuf:"bytes,1,rep,name=streams" json:"streams,omitempty"`
	Err              *Error                    `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *ActiveStreamResponse) Reset()         { *m = ActiveStreamResponse{} }
func (m *ActiveStreamResponse) String() string { return proto.CompactTextString(m) }
func (*ActiveStreamResponse) ProtoMessage()    {}

func (m *ActiveStreamResponse) GetStreams() []*MutationStreamResponse {
	if m != nil {
		return m.Streams
	}
	return nil
}

func (m *ActiveStreamResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

// Requested by Coordinator during initial index build, to calculate
// initial-build-timestamp for each bucket.
type CurrentTimestampRequest struct {
	Buckets          []string `protobuf:"bytes,1,rep,name=buckets" json:"buckets,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *CurrentTimestampRequest) Reset()         { *m = CurrentTimestampRequest{} }
func (m *CurrentTimestampRequest) String() string { return proto.CompactTextString(m) }
func (*CurrentTimestampRequest) ProtoMessage()    {}

func (m *CurrentTimestampRequest) GetBuckets() []string {
	if m != nil {
		return m.Buckets
	}
	return nil
}

type CurrentTimestampResponse struct {
	CurrentTimestamps []*BranchTimestamp `protobuf:"bytes,1,rep,name=currentTimestamps" json:"currentTimestamps,omitempty"`
	XXX_unrecognized  []byte             `json:"-"`
}

func (m *CurrentTimestampResponse) Reset()         { *m = CurrentTimestampResponse{} }
func (m *CurrentTimestampResponse) String() string { return proto.CompactTextString(m) }
func (*CurrentTimestampResponse) ProtoMessage()    {}

func (m *CurrentTimestampResponse) GetCurrentTimestamps() []*BranchTimestamp {
	if m != nil {
		return m.CurrentTimestamps
	}
	return nil
}

func init() {
}
