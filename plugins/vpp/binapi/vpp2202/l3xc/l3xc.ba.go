// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package l3xc contains generated bindings for API file l3xc.api.
//
// Contents:
// -  1 struct
// -  8 messages
package l3xc

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	fib_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2202/fib_types"
	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2202/interface_types"
	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2202/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "l3xc"
	APIVersion = "1.0.1"
	VersionCrc = 0x520bfc6e
)

// L3xc defines type 'l3xc'.
type L3xc struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIP6     bool                           `binapi:"bool,name=is_ip6" json:"is_ip6,omitempty"`
	NPaths    uint8                          `binapi:"u8,name=n_paths" json:"-"`
	Paths     []fib_types.FibPath            `binapi:"fib_path[n_paths],name=paths" json:"paths,omitempty"`
}

// L3xcDel defines message 'l3xc_del'.
type L3xcDel struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	IsIP6     bool                           `binapi:"bool,name=is_ip6" json:"is_ip6,omitempty"`
}

func (m *L3xcDel) Reset()               { *m = L3xcDel{} }
func (*L3xcDel) GetMessageName() string { return "l3xc_del" }
func (*L3xcDel) GetCrcString() string   { return "e7dbef91" }
func (*L3xcDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L3xcDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.IsIP6
	return size
}
func (m *L3xcDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBool(m.IsIP6)
	return buf.Bytes(), nil
}
func (m *L3xcDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.IsIP6 = buf.DecodeBool()
	return nil
}

// L3xcDelReply defines message 'l3xc_del_reply'.
type L3xcDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *L3xcDelReply) Reset()               { *m = L3xcDelReply{} }
func (*L3xcDelReply) GetMessageName() string { return "l3xc_del_reply" }
func (*L3xcDelReply) GetCrcString() string   { return "e8d4e804" }
func (*L3xcDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L3xcDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *L3xcDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *L3xcDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// L3xcDetails defines message 'l3xc_details'.
type L3xcDetails struct {
	L3xc L3xc `binapi:"l3xc,name=l3xc" json:"l3xc,omitempty"`
}

func (m *L3xcDetails) Reset()               { *m = L3xcDetails{} }
func (*L3xcDetails) GetMessageName() string { return "l3xc_details" }
func (*L3xcDetails) GetCrcString() string   { return "bc5bf852" }
func (*L3xcDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L3xcDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.L3xc.SwIfIndex
	size += 1 // m.L3xc.IsIP6
	size += 1 // m.L3xc.NPaths
	for j2 := 0; j2 < len(m.L3xc.Paths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.L3xc.Paths) {
			s2 = m.L3xc.Paths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *L3xcDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.L3xc.SwIfIndex))
	buf.EncodeBool(m.L3xc.IsIP6)
	buf.EncodeUint8(uint8(len(m.L3xc.Paths)))
	for j1 := 0; j1 < len(m.L3xc.Paths); j1++ {
		var v1 fib_types.FibPath // Paths
		if j1 < len(m.L3xc.Paths) {
			v1 = m.L3xc.Paths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *L3xcDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.L3xc.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.L3xc.IsIP6 = buf.DecodeBool()
	m.L3xc.NPaths = buf.DecodeUint8()
	m.L3xc.Paths = make([]fib_types.FibPath, m.L3xc.NPaths)
	for j1 := 0; j1 < len(m.L3xc.Paths); j1++ {
		m.L3xc.Paths[j1].SwIfIndex = buf.DecodeUint32()
		m.L3xc.Paths[j1].TableID = buf.DecodeUint32()
		m.L3xc.Paths[j1].RpfID = buf.DecodeUint32()
		m.L3xc.Paths[j1].Weight = buf.DecodeUint8()
		m.L3xc.Paths[j1].Preference = buf.DecodeUint8()
		m.L3xc.Paths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.L3xc.Paths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.L3xc.Paths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.L3xc.Paths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.L3xc.Paths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.L3xc.Paths[j1].Nh.ObjID = buf.DecodeUint32()
		m.L3xc.Paths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.L3xc.Paths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.L3xc.Paths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.L3xc.Paths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.L3xc.Paths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.L3xc.Paths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// L3xcDump defines message 'l3xc_dump'.
type L3xcDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *L3xcDump) Reset()               { *m = L3xcDump{} }
func (*L3xcDump) GetMessageName() string { return "l3xc_dump" }
func (*L3xcDump) GetCrcString() string   { return "f9e6675e" }
func (*L3xcDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L3xcDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *L3xcDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *L3xcDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// L3xcPluginGetVersion defines message 'l3xc_plugin_get_version'.
type L3xcPluginGetVersion struct{}

func (m *L3xcPluginGetVersion) Reset()               { *m = L3xcPluginGetVersion{} }
func (*L3xcPluginGetVersion) GetMessageName() string { return "l3xc_plugin_get_version" }
func (*L3xcPluginGetVersion) GetCrcString() string   { return "51077d14" }
func (*L3xcPluginGetVersion) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L3xcPluginGetVersion) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *L3xcPluginGetVersion) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *L3xcPluginGetVersion) Unmarshal(b []byte) error {
	return nil
}

// L3xcPluginGetVersionReply defines message 'l3xc_plugin_get_version_reply'.
type L3xcPluginGetVersionReply struct {
	Major uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
}

func (m *L3xcPluginGetVersionReply) Reset()               { *m = L3xcPluginGetVersionReply{} }
func (*L3xcPluginGetVersionReply) GetMessageName() string { return "l3xc_plugin_get_version_reply" }
func (*L3xcPluginGetVersionReply) GetCrcString() string   { return "9b32cf86" }
func (*L3xcPluginGetVersionReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L3xcPluginGetVersionReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Major
	size += 4 // m.Minor
	return size
}
func (m *L3xcPluginGetVersionReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Major)
	buf.EncodeUint32(m.Minor)
	return buf.Bytes(), nil
}
func (m *L3xcPluginGetVersionReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Major = buf.DecodeUint32()
	m.Minor = buf.DecodeUint32()
	return nil
}

// L3xcUpdate defines message 'l3xc_update'.
type L3xcUpdate struct {
	L3xc L3xc `binapi:"l3xc,name=l3xc" json:"l3xc,omitempty"`
}

func (m *L3xcUpdate) Reset()               { *m = L3xcUpdate{} }
func (*L3xcUpdate) GetMessageName() string { return "l3xc_update" }
func (*L3xcUpdate) GetCrcString() string   { return "e96aabdf" }
func (*L3xcUpdate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *L3xcUpdate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.L3xc.SwIfIndex
	size += 1 // m.L3xc.IsIP6
	size += 1 // m.L3xc.NPaths
	for j2 := 0; j2 < len(m.L3xc.Paths); j2++ {
		var s2 fib_types.FibPath
		_ = s2
		if j2 < len(m.L3xc.Paths) {
			s2 = m.L3xc.Paths[j2]
		}
		size += 4      // s2.SwIfIndex
		size += 4      // s2.TableID
		size += 4      // s2.RpfID
		size += 1      // s2.Weight
		size += 1      // s2.Preference
		size += 4      // s2.Type
		size += 4      // s2.Flags
		size += 4      // s2.Proto
		size += 1 * 16 // s2.Nh.Address
		size += 4      // s2.Nh.ViaLabel
		size += 4      // s2.Nh.ObjID
		size += 4      // s2.Nh.ClassifyTableIndex
		size += 1      // s2.NLabels
		for j3 := 0; j3 < 16; j3++ {
			size += 1 // s2.LabelStack[j3].IsUniform
			size += 4 // s2.LabelStack[j3].Label
			size += 1 // s2.LabelStack[j3].TTL
			size += 1 // s2.LabelStack[j3].Exp
		}
	}
	return size
}
func (m *L3xcUpdate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.L3xc.SwIfIndex))
	buf.EncodeBool(m.L3xc.IsIP6)
	buf.EncodeUint8(uint8(len(m.L3xc.Paths)))
	for j1 := 0; j1 < len(m.L3xc.Paths); j1++ {
		var v1 fib_types.FibPath // Paths
		if j1 < len(m.L3xc.Paths) {
			v1 = m.L3xc.Paths[j1]
		}
		buf.EncodeUint32(v1.SwIfIndex)
		buf.EncodeUint32(v1.TableID)
		buf.EncodeUint32(v1.RpfID)
		buf.EncodeUint8(v1.Weight)
		buf.EncodeUint8(v1.Preference)
		buf.EncodeUint32(uint32(v1.Type))
		buf.EncodeUint32(uint32(v1.Flags))
		buf.EncodeUint32(uint32(v1.Proto))
		buf.EncodeBytes(v1.Nh.Address.XXX_UnionData[:], 16)
		buf.EncodeUint32(v1.Nh.ViaLabel)
		buf.EncodeUint32(v1.Nh.ObjID)
		buf.EncodeUint32(v1.Nh.ClassifyTableIndex)
		buf.EncodeUint8(v1.NLabels)
		for j2 := 0; j2 < 16; j2++ {
			buf.EncodeUint8(v1.LabelStack[j2].IsUniform)
			buf.EncodeUint32(v1.LabelStack[j2].Label)
			buf.EncodeUint8(v1.LabelStack[j2].TTL)
			buf.EncodeUint8(v1.LabelStack[j2].Exp)
		}
	}
	return buf.Bytes(), nil
}
func (m *L3xcUpdate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.L3xc.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.L3xc.IsIP6 = buf.DecodeBool()
	m.L3xc.NPaths = buf.DecodeUint8()
	m.L3xc.Paths = make([]fib_types.FibPath, m.L3xc.NPaths)
	for j1 := 0; j1 < len(m.L3xc.Paths); j1++ {
		m.L3xc.Paths[j1].SwIfIndex = buf.DecodeUint32()
		m.L3xc.Paths[j1].TableID = buf.DecodeUint32()
		m.L3xc.Paths[j1].RpfID = buf.DecodeUint32()
		m.L3xc.Paths[j1].Weight = buf.DecodeUint8()
		m.L3xc.Paths[j1].Preference = buf.DecodeUint8()
		m.L3xc.Paths[j1].Type = fib_types.FibPathType(buf.DecodeUint32())
		m.L3xc.Paths[j1].Flags = fib_types.FibPathFlags(buf.DecodeUint32())
		m.L3xc.Paths[j1].Proto = fib_types.FibPathNhProto(buf.DecodeUint32())
		copy(m.L3xc.Paths[j1].Nh.Address.XXX_UnionData[:], buf.DecodeBytes(16))
		m.L3xc.Paths[j1].Nh.ViaLabel = buf.DecodeUint32()
		m.L3xc.Paths[j1].Nh.ObjID = buf.DecodeUint32()
		m.L3xc.Paths[j1].Nh.ClassifyTableIndex = buf.DecodeUint32()
		m.L3xc.Paths[j1].NLabels = buf.DecodeUint8()
		for j2 := 0; j2 < 16; j2++ {
			m.L3xc.Paths[j1].LabelStack[j2].IsUniform = buf.DecodeUint8()
			m.L3xc.Paths[j1].LabelStack[j2].Label = buf.DecodeUint32()
			m.L3xc.Paths[j1].LabelStack[j2].TTL = buf.DecodeUint8()
			m.L3xc.Paths[j1].LabelStack[j2].Exp = buf.DecodeUint8()
		}
	}
	return nil
}

// L3xcUpdateReply defines message 'l3xc_update_reply'.
type L3xcUpdateReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	StatsIndex uint32 `binapi:"u32,name=stats_index" json:"stats_index,omitempty"`
}

func (m *L3xcUpdateReply) Reset()               { *m = L3xcUpdateReply{} }
func (*L3xcUpdateReply) GetMessageName() string { return "l3xc_update_reply" }
func (*L3xcUpdateReply) GetCrcString() string   { return "1992deab" }
func (*L3xcUpdateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *L3xcUpdateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.StatsIndex
	return size
}
func (m *L3xcUpdateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.StatsIndex)
	return buf.Bytes(), nil
}
func (m *L3xcUpdateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.StatsIndex = buf.DecodeUint32()
	return nil
}

func init() { file_l3xc_binapi_init() }
func file_l3xc_binapi_init() {
	api.RegisterMessage((*L3xcDel)(nil), "l3xc_del_e7dbef91")
	api.RegisterMessage((*L3xcDelReply)(nil), "l3xc_del_reply_e8d4e804")
	api.RegisterMessage((*L3xcDetails)(nil), "l3xc_details_bc5bf852")
	api.RegisterMessage((*L3xcDump)(nil), "l3xc_dump_f9e6675e")
	api.RegisterMessage((*L3xcPluginGetVersion)(nil), "l3xc_plugin_get_version_51077d14")
	api.RegisterMessage((*L3xcPluginGetVersionReply)(nil), "l3xc_plugin_get_version_reply_9b32cf86")
	api.RegisterMessage((*L3xcUpdate)(nil), "l3xc_update_e96aabdf")
	api.RegisterMessage((*L3xcUpdateReply)(nil), "l3xc_update_reply_1992deab")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*L3xcDel)(nil),
		(*L3xcDelReply)(nil),
		(*L3xcDetails)(nil),
		(*L3xcDump)(nil),
		(*L3xcPluginGetVersion)(nil),
		(*L3xcPluginGetVersionReply)(nil),
		(*L3xcUpdate)(nil),
		(*L3xcUpdateReply)(nil),
	}
}
