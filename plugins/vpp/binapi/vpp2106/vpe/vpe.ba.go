// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package vpe contains generated bindings for API file vpe.api.
//
// Contents:
// -  1 struct
// - 26 messages
package vpe

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	vpe_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2106/vpe_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "vpe"
	APIVersion = "1.6.1"
	VersionCrc = 0x40364e00
)

// ThreadData defines type 'thread_data'.
type ThreadData struct {
	ID        uint32 `binapi:"u32,name=id" json:"id,omitempty"`
	Name      string `binapi:"string[64],name=name" json:"name,omitempty"`
	Type      string `binapi:"string[64],name=type" json:"type,omitempty"`
	PID       uint32 `binapi:"u32,name=pid" json:"pid,omitempty"`
	CPUID     uint32 `binapi:"u32,name=cpu_id" json:"cpu_id,omitempty"`
	Core      uint32 `binapi:"u32,name=core" json:"core,omitempty"`
	CPUSocket uint32 `binapi:"u32,name=cpu_socket" json:"cpu_socket,omitempty"`
}

// AddNodeNext defines message 'add_node_next'.
type AddNodeNext struct {
	NodeName string `binapi:"string[64],name=node_name" json:"node_name,omitempty"`
	NextName string `binapi:"string[64],name=next_name" json:"next_name,omitempty"`
}

func (m *AddNodeNext) Reset()               { *m = AddNodeNext{} }
func (*AddNodeNext) GetMessageName() string { return "add_node_next" }
func (*AddNodeNext) GetCrcString() string   { return "2457116d" }
func (*AddNodeNext) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *AddNodeNext) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.NodeName
	size += 64 // m.NextName
	return size
}
func (m *AddNodeNext) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.NodeName, 64)
	buf.EncodeString(m.NextName, 64)
	return buf.Bytes(), nil
}
func (m *AddNodeNext) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeName = buf.DecodeString(64)
	m.NextName = buf.DecodeString(64)
	return nil
}

// AddNodeNextReply defines message 'add_node_next_reply'.
type AddNodeNextReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	NextIndex uint32 `binapi:"u32,name=next_index" json:"next_index,omitempty"`
}

func (m *AddNodeNextReply) Reset()               { *m = AddNodeNextReply{} }
func (*AddNodeNextReply) GetMessageName() string { return "add_node_next_reply" }
func (*AddNodeNextReply) GetCrcString() string   { return "2ed75f32" }
func (*AddNodeNextReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *AddNodeNextReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.NextIndex
	return size
}
func (m *AddNodeNextReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.NextIndex)
	return buf.Bytes(), nil
}
func (m *AddNodeNextReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.NextIndex = buf.DecodeUint32()
	return nil
}

// Cli defines message 'cli'.
type Cli struct {
	CmdInShmem uint64 `binapi:"u64,name=cmd_in_shmem" json:"cmd_in_shmem,omitempty"`
}

func (m *Cli) Reset()               { *m = Cli{} }
func (*Cli) GetMessageName() string { return "cli" }
func (*Cli) GetCrcString() string   { return "23bfbfff" }
func (*Cli) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *Cli) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.CmdInShmem
	return size
}
func (m *Cli) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint64(m.CmdInShmem)
	return buf.Bytes(), nil
}
func (m *Cli) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.CmdInShmem = buf.DecodeUint64()
	return nil
}

// CliInband defines message 'cli_inband'.
type CliInband struct {
	Cmd string `binapi:"string[],name=cmd" json:"cmd,omitempty"`
}

func (m *CliInband) Reset()               { *m = CliInband{} }
func (*CliInband) GetMessageName() string { return "cli_inband" }
func (*CliInband) GetCrcString() string   { return "f8377302" }
func (*CliInband) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *CliInband) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 + len(m.Cmd) // m.Cmd
	return size
}
func (m *CliInband) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Cmd, 0)
	return buf.Bytes(), nil
}
func (m *CliInband) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cmd = buf.DecodeString(0)
	return nil
}

// CliInbandReply defines message 'cli_inband_reply'.
type CliInbandReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Reply  string `binapi:"string[],name=reply" json:"reply,omitempty"`
}

func (m *CliInbandReply) Reset()               { *m = CliInbandReply{} }
func (*CliInbandReply) GetMessageName() string { return "cli_inband_reply" }
func (*CliInbandReply) GetCrcString() string   { return "05879051" }
func (*CliInbandReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CliInbandReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4                // m.Retval
	size += 4 + len(m.Reply) // m.Reply
	return size
}
func (m *CliInbandReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.Reply, 0)
	return buf.Bytes(), nil
}
func (m *CliInbandReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Reply = buf.DecodeString(0)
	return nil
}

// CliReply defines message 'cli_reply'.
type CliReply struct {
	Retval       int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ReplyInShmem uint64 `binapi:"u64,name=reply_in_shmem" json:"reply_in_shmem,omitempty"`
}

func (m *CliReply) Reset()               { *m = CliReply{} }
func (*CliReply) GetMessageName() string { return "cli_reply" }
func (*CliReply) GetCrcString() string   { return "06d68297" }
func (*CliReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *CliReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.ReplyInShmem
	return size
}
func (m *CliReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint64(m.ReplyInShmem)
	return buf.Bytes(), nil
}
func (m *CliReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ReplyInShmem = buf.DecodeUint64()
	return nil
}

// ControlPing defines message 'control_ping'.
type ControlPing struct{}

func (m *ControlPing) Reset()               { *m = ControlPing{} }
func (*ControlPing) GetMessageName() string { return "control_ping" }
func (*ControlPing) GetCrcString() string   { return "51077d14" }
func (*ControlPing) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ControlPing) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ControlPing) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ControlPing) Unmarshal(b []byte) error {
	return nil
}

// ControlPingReply defines message 'control_ping_reply'.
type ControlPingReply struct {
	Retval      int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ClientIndex uint32 `binapi:"u32,name=client_index" json:"client_index,omitempty"`
	VpePID      uint32 `binapi:"u32,name=vpe_pid" json:"vpe_pid,omitempty"`
}

func (m *ControlPingReply) Reset()               { *m = ControlPingReply{} }
func (*ControlPingReply) GetMessageName() string { return "control_ping_reply" }
func (*ControlPingReply) GetCrcString() string   { return "f6b0b8ca" }
func (*ControlPingReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ControlPingReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.ClientIndex
	size += 4 // m.VpePID
	return size
}
func (m *ControlPingReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.ClientIndex)
	buf.EncodeUint32(m.VpePID)
	return buf.Bytes(), nil
}
func (m *ControlPingReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ClientIndex = buf.DecodeUint32()
	m.VpePID = buf.DecodeUint32()
	return nil
}

// GetF64EndianValue defines message 'get_f64_endian_value'.
type GetF64EndianValue struct {
	F64One float64 `binapi:"f64,name=f64_one,default=1" json:"f64_one,omitempty"`
}

func (m *GetF64EndianValue) Reset()               { *m = GetF64EndianValue{} }
func (*GetF64EndianValue) GetMessageName() string { return "get_f64_endian_value" }
func (*GetF64EndianValue) GetCrcString() string   { return "809fcd44" }
func (*GetF64EndianValue) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetF64EndianValue) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.F64One
	return size
}
func (m *GetF64EndianValue) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeFloat64(m.F64One)
	return buf.Bytes(), nil
}
func (m *GetF64EndianValue) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.F64One = buf.DecodeFloat64()
	return nil
}

// GetF64EndianValueReply defines message 'get_f64_endian_value_reply'.
type GetF64EndianValueReply struct {
	Retval       uint32  `binapi:"u32,name=retval" json:"retval,omitempty"`
	F64OneResult float64 `binapi:"f64,name=f64_one_result" json:"f64_one_result,omitempty"`
}

func (m *GetF64EndianValueReply) Reset()               { *m = GetF64EndianValueReply{} }
func (*GetF64EndianValueReply) GetMessageName() string { return "get_f64_endian_value_reply" }
func (*GetF64EndianValueReply) GetCrcString() string   { return "7e02e404" }
func (*GetF64EndianValueReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetF64EndianValueReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.F64OneResult
	return size
}
func (m *GetF64EndianValueReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Retval)
	buf.EncodeFloat64(m.F64OneResult)
	return buf.Bytes(), nil
}
func (m *GetF64EndianValueReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeUint32()
	m.F64OneResult = buf.DecodeFloat64()
	return nil
}

// GetF64IncrementByOne defines message 'get_f64_increment_by_one'.
type GetF64IncrementByOne struct {
	F64Value float64 `binapi:"f64,name=f64_value,default=1" json:"f64_value,omitempty"`
}

func (m *GetF64IncrementByOne) Reset()               { *m = GetF64IncrementByOne{} }
func (*GetF64IncrementByOne) GetMessageName() string { return "get_f64_increment_by_one" }
func (*GetF64IncrementByOne) GetCrcString() string   { return "b64f027e" }
func (*GetF64IncrementByOne) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetF64IncrementByOne) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.F64Value
	return size
}
func (m *GetF64IncrementByOne) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeFloat64(m.F64Value)
	return buf.Bytes(), nil
}
func (m *GetF64IncrementByOne) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.F64Value = buf.DecodeFloat64()
	return nil
}

// GetF64IncrementByOneReply defines message 'get_f64_increment_by_one_reply'.
type GetF64IncrementByOneReply struct {
	Retval   uint32  `binapi:"u32,name=retval" json:"retval,omitempty"`
	F64Value float64 `binapi:"f64,name=f64_value" json:"f64_value,omitempty"`
}

func (m *GetF64IncrementByOneReply) Reset()               { *m = GetF64IncrementByOneReply{} }
func (*GetF64IncrementByOneReply) GetMessageName() string { return "get_f64_increment_by_one_reply" }
func (*GetF64IncrementByOneReply) GetCrcString() string   { return "d25dbaa3" }
func (*GetF64IncrementByOneReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetF64IncrementByOneReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.F64Value
	return size
}
func (m *GetF64IncrementByOneReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Retval)
	buf.EncodeFloat64(m.F64Value)
	return buf.Bytes(), nil
}
func (m *GetF64IncrementByOneReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeUint32()
	m.F64Value = buf.DecodeFloat64()
	return nil
}

// GetNextIndex defines message 'get_next_index'.
type GetNextIndex struct {
	NodeName string `binapi:"string[64],name=node_name" json:"node_name,omitempty"`
	NextName string `binapi:"string[64],name=next_name" json:"next_name,omitempty"`
}

func (m *GetNextIndex) Reset()               { *m = GetNextIndex{} }
func (*GetNextIndex) GetMessageName() string { return "get_next_index" }
func (*GetNextIndex) GetCrcString() string   { return "2457116d" }
func (*GetNextIndex) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetNextIndex) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.NodeName
	size += 64 // m.NextName
	return size
}
func (m *GetNextIndex) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.NodeName, 64)
	buf.EncodeString(m.NextName, 64)
	return buf.Bytes(), nil
}
func (m *GetNextIndex) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeName = buf.DecodeString(64)
	m.NextName = buf.DecodeString(64)
	return nil
}

// GetNextIndexReply defines message 'get_next_index_reply'.
type GetNextIndexReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	NextIndex uint32 `binapi:"u32,name=next_index" json:"next_index,omitempty"`
}

func (m *GetNextIndexReply) Reset()               { *m = GetNextIndexReply{} }
func (*GetNextIndexReply) GetMessageName() string { return "get_next_index_reply" }
func (*GetNextIndexReply) GetCrcString() string   { return "2ed75f32" }
func (*GetNextIndexReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetNextIndexReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.NextIndex
	return size
}
func (m *GetNextIndexReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.NextIndex)
	return buf.Bytes(), nil
}
func (m *GetNextIndexReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.NextIndex = buf.DecodeUint32()
	return nil
}

// GetNodeGraph defines message 'get_node_graph'.
type GetNodeGraph struct{}

func (m *GetNodeGraph) Reset()               { *m = GetNodeGraph{} }
func (*GetNodeGraph) GetMessageName() string { return "get_node_graph" }
func (*GetNodeGraph) GetCrcString() string   { return "51077d14" }
func (*GetNodeGraph) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetNodeGraph) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *GetNodeGraph) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *GetNodeGraph) Unmarshal(b []byte) error {
	return nil
}

// GetNodeGraphReply defines message 'get_node_graph_reply'.
type GetNodeGraphReply struct {
	Retval       int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	ReplyInShmem uint64 `binapi:"u64,name=reply_in_shmem" json:"reply_in_shmem,omitempty"`
}

func (m *GetNodeGraphReply) Reset()               { *m = GetNodeGraphReply{} }
func (*GetNodeGraphReply) GetMessageName() string { return "get_node_graph_reply" }
func (*GetNodeGraphReply) GetCrcString() string   { return "06d68297" }
func (*GetNodeGraphReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetNodeGraphReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.ReplyInShmem
	return size
}
func (m *GetNodeGraphReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint64(m.ReplyInShmem)
	return buf.Bytes(), nil
}
func (m *GetNodeGraphReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.ReplyInShmem = buf.DecodeUint64()
	return nil
}

// GetNodeIndex defines message 'get_node_index'.
type GetNodeIndex struct {
	NodeName string `binapi:"string[64],name=node_name" json:"node_name,omitempty"`
}

func (m *GetNodeIndex) Reset()               { *m = GetNodeIndex{} }
func (*GetNodeIndex) GetMessageName() string { return "get_node_index" }
func (*GetNodeIndex) GetCrcString() string   { return "f1984c64" }
func (*GetNodeIndex) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GetNodeIndex) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 64 // m.NodeName
	return size
}
func (m *GetNodeIndex) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.NodeName, 64)
	return buf.Bytes(), nil
}
func (m *GetNodeIndex) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.NodeName = buf.DecodeString(64)
	return nil
}

// GetNodeIndexReply defines message 'get_node_index_reply'.
type GetNodeIndexReply struct {
	Retval    int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	NodeIndex uint32 `binapi:"u32,name=node_index" json:"node_index,omitempty"`
}

func (m *GetNodeIndexReply) Reset()               { *m = GetNodeIndexReply{} }
func (*GetNodeIndexReply) GetMessageName() string { return "get_node_index_reply" }
func (*GetNodeIndexReply) GetCrcString() string   { return "a8600b89" }
func (*GetNodeIndexReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GetNodeIndexReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.NodeIndex
	return size
}
func (m *GetNodeIndexReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.NodeIndex)
	return buf.Bytes(), nil
}
func (m *GetNodeIndexReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.NodeIndex = buf.DecodeUint32()
	return nil
}

// LogDetails defines message 'log_details'.
type LogDetails struct {
	Timestamp vpe_types.Timestamp `binapi:"timestamp,name=timestamp" json:"timestamp,omitempty"`
	Level     vpe_types.LogLevel  `binapi:"log_level,name=level" json:"level,omitempty"`
	MsgClass  string              `binapi:"string[32],name=msg_class" json:"msg_class,omitempty"`
	Message   string              `binapi:"string[256],name=message" json:"message,omitempty"`
}

func (m *LogDetails) Reset()               { *m = LogDetails{} }
func (*LogDetails) GetMessageName() string { return "log_details" }
func (*LogDetails) GetCrcString() string   { return "03d61cc0" }
func (*LogDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LogDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8   // m.Timestamp
	size += 4   // m.Level
	size += 32  // m.MsgClass
	size += 256 // m.Message
	return size
}
func (m *LogDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeFloat64(float64(m.Timestamp))
	buf.EncodeUint32(uint32(m.Level))
	buf.EncodeString(m.MsgClass, 32)
	buf.EncodeString(m.Message, 256)
	return buf.Bytes(), nil
}
func (m *LogDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Timestamp = vpe_types.Timestamp(buf.DecodeFloat64())
	m.Level = vpe_types.LogLevel(buf.DecodeUint32())
	m.MsgClass = buf.DecodeString(32)
	m.Message = buf.DecodeString(256)
	return nil
}

// LogDump defines message 'log_dump'.
type LogDump struct {
	StartTimestamp vpe_types.Timestamp `binapi:"timestamp,name=start_timestamp" json:"start_timestamp,omitempty"`
}

func (m *LogDump) Reset()               { *m = LogDump{} }
func (*LogDump) GetMessageName() string { return "log_dump" }
func (*LogDump) GetCrcString() string   { return "6ab31753" }
func (*LogDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LogDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8 // m.StartTimestamp
	return size
}
func (m *LogDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeFloat64(float64(m.StartTimestamp))
	return buf.Bytes(), nil
}
func (m *LogDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.StartTimestamp = vpe_types.Timestamp(buf.DecodeFloat64())
	return nil
}

// ShowThreads defines message 'show_threads'.
type ShowThreads struct{}

func (m *ShowThreads) Reset()               { *m = ShowThreads{} }
func (*ShowThreads) GetMessageName() string { return "show_threads" }
func (*ShowThreads) GetCrcString() string   { return "51077d14" }
func (*ShowThreads) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ShowThreads) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ShowThreads) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ShowThreads) Unmarshal(b []byte) error {
	return nil
}

// ShowThreadsReply defines message 'show_threads_reply'.
type ShowThreadsReply struct {
	Retval     int32        `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count      uint32       `binapi:"u32,name=count" json:"-"`
	ThreadData []ThreadData `binapi:"thread_data[count],name=thread_data" json:"thread_data,omitempty"`
}

func (m *ShowThreadsReply) Reset()               { *m = ShowThreadsReply{} }
func (*ShowThreadsReply) GetMessageName() string { return "show_threads_reply" }
func (*ShowThreadsReply) GetCrcString() string   { return "efd78e83" }
func (*ShowThreadsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ShowThreadsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Count
	for j1 := 0; j1 < len(m.ThreadData); j1++ {
		var s1 ThreadData
		_ = s1
		if j1 < len(m.ThreadData) {
			s1 = m.ThreadData[j1]
		}
		size += 4  // s1.ID
		size += 64 // s1.Name
		size += 64 // s1.Type
		size += 4  // s1.PID
		size += 4  // s1.CPUID
		size += 4  // s1.Core
		size += 4  // s1.CPUSocket
	}
	return size
}
func (m *ShowThreadsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(len(m.ThreadData)))
	for j0 := 0; j0 < len(m.ThreadData); j0++ {
		var v0 ThreadData // ThreadData
		if j0 < len(m.ThreadData) {
			v0 = m.ThreadData[j0]
		}
		buf.EncodeUint32(v0.ID)
		buf.EncodeString(v0.Name, 64)
		buf.EncodeString(v0.Type, 64)
		buf.EncodeUint32(v0.PID)
		buf.EncodeUint32(v0.CPUID)
		buf.EncodeUint32(v0.Core)
		buf.EncodeUint32(v0.CPUSocket)
	}
	return buf.Bytes(), nil
}
func (m *ShowThreadsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint32()
	m.ThreadData = make([]ThreadData, m.Count)
	for j0 := 0; j0 < len(m.ThreadData); j0++ {
		m.ThreadData[j0].ID = buf.DecodeUint32()
		m.ThreadData[j0].Name = buf.DecodeString(64)
		m.ThreadData[j0].Type = buf.DecodeString(64)
		m.ThreadData[j0].PID = buf.DecodeUint32()
		m.ThreadData[j0].CPUID = buf.DecodeUint32()
		m.ThreadData[j0].Core = buf.DecodeUint32()
		m.ThreadData[j0].CPUSocket = buf.DecodeUint32()
	}
	return nil
}

// ShowVersion defines message 'show_version'.
type ShowVersion struct{}

func (m *ShowVersion) Reset()               { *m = ShowVersion{} }
func (*ShowVersion) GetMessageName() string { return "show_version" }
func (*ShowVersion) GetCrcString() string   { return "51077d14" }
func (*ShowVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ShowVersion) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ShowVersion) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ShowVersion) Unmarshal(b []byte) error {
	return nil
}

// ShowVersionReply defines message 'show_version_reply'.
type ShowVersionReply struct {
	Retval         int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Program        string `binapi:"string[32],name=program" json:"program,omitempty"`
	Version        string `binapi:"string[32],name=version" json:"version,omitempty"`
	BuildDate      string `binapi:"string[32],name=build_date" json:"build_date,omitempty"`
	BuildDirectory string `binapi:"string[256],name=build_directory" json:"build_directory,omitempty"`
}

func (m *ShowVersionReply) Reset()               { *m = ShowVersionReply{} }
func (*ShowVersionReply) GetMessageName() string { return "show_version_reply" }
func (*ShowVersionReply) GetCrcString() string   { return "c919bde1" }
func (*ShowVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ShowVersionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.Retval
	size += 32  // m.Program
	size += 32  // m.Version
	size += 32  // m.BuildDate
	size += 256 // m.BuildDirectory
	return size
}
func (m *ShowVersionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeString(m.Program, 32)
	buf.EncodeString(m.Version, 32)
	buf.EncodeString(m.BuildDate, 32)
	buf.EncodeString(m.BuildDirectory, 256)
	return buf.Bytes(), nil
}
func (m *ShowVersionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Program = buf.DecodeString(32)
	m.Version = buf.DecodeString(32)
	m.BuildDate = buf.DecodeString(32)
	m.BuildDirectory = buf.DecodeString(256)
	return nil
}

// ShowVpeSystemTime defines message 'show_vpe_system_time'.
type ShowVpeSystemTime struct{}

func (m *ShowVpeSystemTime) Reset()               { *m = ShowVpeSystemTime{} }
func (*ShowVpeSystemTime) GetMessageName() string { return "show_vpe_system_time" }
func (*ShowVpeSystemTime) GetCrcString() string   { return "51077d14" }
func (*ShowVpeSystemTime) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *ShowVpeSystemTime) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *ShowVpeSystemTime) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *ShowVpeSystemTime) Unmarshal(b []byte) error {
	return nil
}

// ShowVpeSystemTimeReply defines message 'show_vpe_system_time_reply'.
type ShowVpeSystemTimeReply struct {
	Retval        int32               `binapi:"i32,name=retval" json:"retval,omitempty"`
	VpeSystemTime vpe_types.Timestamp `binapi:"timestamp,name=vpe_system_time" json:"vpe_system_time,omitempty"`
}

func (m *ShowVpeSystemTimeReply) Reset()               { *m = ShowVpeSystemTimeReply{} }
func (*ShowVpeSystemTimeReply) GetMessageName() string { return "show_vpe_system_time_reply" }
func (*ShowVpeSystemTimeReply) GetCrcString() string   { return "7ffd8193" }
func (*ShowVpeSystemTimeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *ShowVpeSystemTimeReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 8 // m.VpeSystemTime
	return size
}
func (m *ShowVpeSystemTimeReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeFloat64(float64(m.VpeSystemTime))
	return buf.Bytes(), nil
}
func (m *ShowVpeSystemTimeReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.VpeSystemTime = vpe_types.Timestamp(buf.DecodeFloat64())
	return nil
}

func init() { file_vpe_binapi_init() }
func file_vpe_binapi_init() {
	api.RegisterMessage((*AddNodeNext)(nil), "add_node_next_2457116d")
	api.RegisterMessage((*AddNodeNextReply)(nil), "add_node_next_reply_2ed75f32")
	api.RegisterMessage((*Cli)(nil), "cli_23bfbfff")
	api.RegisterMessage((*CliInband)(nil), "cli_inband_f8377302")
	api.RegisterMessage((*CliInbandReply)(nil), "cli_inband_reply_05879051")
	api.RegisterMessage((*CliReply)(nil), "cli_reply_06d68297")
	api.RegisterMessage((*ControlPing)(nil), "control_ping_51077d14")
	api.RegisterMessage((*ControlPingReply)(nil), "control_ping_reply_f6b0b8ca")
	api.RegisterMessage((*GetF64EndianValue)(nil), "get_f64_endian_value_809fcd44")
	api.RegisterMessage((*GetF64EndianValueReply)(nil), "get_f64_endian_value_reply_7e02e404")
	api.RegisterMessage((*GetF64IncrementByOne)(nil), "get_f64_increment_by_one_b64f027e")
	api.RegisterMessage((*GetF64IncrementByOneReply)(nil), "get_f64_increment_by_one_reply_d25dbaa3")
	api.RegisterMessage((*GetNextIndex)(nil), "get_next_index_2457116d")
	api.RegisterMessage((*GetNextIndexReply)(nil), "get_next_index_reply_2ed75f32")
	api.RegisterMessage((*GetNodeGraph)(nil), "get_node_graph_51077d14")
	api.RegisterMessage((*GetNodeGraphReply)(nil), "get_node_graph_reply_06d68297")
	api.RegisterMessage((*GetNodeIndex)(nil), "get_node_index_f1984c64")
	api.RegisterMessage((*GetNodeIndexReply)(nil), "get_node_index_reply_a8600b89")
	api.RegisterMessage((*LogDetails)(nil), "log_details_03d61cc0")
	api.RegisterMessage((*LogDump)(nil), "log_dump_6ab31753")
	api.RegisterMessage((*ShowThreads)(nil), "show_threads_51077d14")
	api.RegisterMessage((*ShowThreadsReply)(nil), "show_threads_reply_efd78e83")
	api.RegisterMessage((*ShowVersion)(nil), "show_version_51077d14")
	api.RegisterMessage((*ShowVersionReply)(nil), "show_version_reply_c919bde1")
	api.RegisterMessage((*ShowVpeSystemTime)(nil), "show_vpe_system_time_51077d14")
	api.RegisterMessage((*ShowVpeSystemTimeReply)(nil), "show_vpe_system_time_reply_7ffd8193")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*AddNodeNext)(nil),
		(*AddNodeNextReply)(nil),
		(*Cli)(nil),
		(*CliInband)(nil),
		(*CliInbandReply)(nil),
		(*CliReply)(nil),
		(*ControlPing)(nil),
		(*ControlPingReply)(nil),
		(*GetF64EndianValue)(nil),
		(*GetF64EndianValueReply)(nil),
		(*GetF64IncrementByOne)(nil),
		(*GetF64IncrementByOneReply)(nil),
		(*GetNextIndex)(nil),
		(*GetNextIndexReply)(nil),
		(*GetNodeGraph)(nil),
		(*GetNodeGraphReply)(nil),
		(*GetNodeIndex)(nil),
		(*GetNodeIndexReply)(nil),
		(*LogDetails)(nil),
		(*LogDump)(nil),
		(*ShowThreads)(nil),
		(*ShowThreadsReply)(nil),
		(*ShowVersion)(nil),
		(*ShowVersionReply)(nil),
		(*ShowVpeSystemTime)(nil),
		(*ShowVpeSystemTimeReply)(nil),
	}
}
