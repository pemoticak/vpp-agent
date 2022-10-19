// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package vpe contains generated bindings for API file vpe.api.
//
// Contents:
//   6 messages
//
package vpe

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	vpe_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2210/vpe_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "vpe"
	APIVersion = "1.7.0"
	VersionCrc = 0xbbfa7484
)

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
	api.RegisterMessage((*LogDetails)(nil), "log_details_03d61cc0")
	api.RegisterMessage((*LogDump)(nil), "log_dump_6ab31753")
	api.RegisterMessage((*ShowVersion)(nil), "show_version_51077d14")
	api.RegisterMessage((*ShowVersionReply)(nil), "show_version_reply_c919bde1")
	api.RegisterMessage((*ShowVpeSystemTime)(nil), "show_vpe_system_time_51077d14")
	api.RegisterMessage((*ShowVpeSystemTimeReply)(nil), "show_vpe_system_time_reply_7ffd8193")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LogDetails)(nil),
		(*LogDump)(nil),
		(*ShowVersion)(nil),
		(*ShowVersionReply)(nil),
		(*ShowVpeSystemTime)(nil),
		(*ShowVpeSystemTimeReply)(nil),
	}
}