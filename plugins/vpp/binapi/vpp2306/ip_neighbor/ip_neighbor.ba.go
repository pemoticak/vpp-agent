// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package ip_neighbor contains generated bindings for API file ip_neighbor.api.
//
// Contents:
// -  2 enums
// -  1 struct
// - 18 messages
package ip_neighbor

import (
	"strconv"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	ethernet_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/ethernet_types"
	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/interface_types"
	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/ip_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ip_neighbor"
	APIVersion = "1.0.0"
	VersionCrc = 0xfde4a69a
)

// IPNeighborEventFlags defines enum 'ip_neighbor_event_flags'.
type IPNeighborEventFlags uint32

const (
	IP_NEIGHBOR_API_EVENT_FLAG_ADDED   IPNeighborEventFlags = 1
	IP_NEIGHBOR_API_EVENT_FLAG_REMOVED IPNeighborEventFlags = 2
)

var (
	IPNeighborEventFlags_name = map[uint32]string{
		1: "IP_NEIGHBOR_API_EVENT_FLAG_ADDED",
		2: "IP_NEIGHBOR_API_EVENT_FLAG_REMOVED",
	}
	IPNeighborEventFlags_value = map[string]uint32{
		"IP_NEIGHBOR_API_EVENT_FLAG_ADDED":   1,
		"IP_NEIGHBOR_API_EVENT_FLAG_REMOVED": 2,
	}
)

func (x IPNeighborEventFlags) String() string {
	s, ok := IPNeighborEventFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := IPNeighborEventFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "IPNeighborEventFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// IPNeighborFlags defines enum 'ip_neighbor_flags'.
type IPNeighborFlags uint8

const (
	IP_API_NEIGHBOR_FLAG_NONE         IPNeighborFlags = 0
	IP_API_NEIGHBOR_FLAG_STATIC       IPNeighborFlags = 1
	IP_API_NEIGHBOR_FLAG_NO_FIB_ENTRY IPNeighborFlags = 2
)

var (
	IPNeighborFlags_name = map[uint8]string{
		0: "IP_API_NEIGHBOR_FLAG_NONE",
		1: "IP_API_NEIGHBOR_FLAG_STATIC",
		2: "IP_API_NEIGHBOR_FLAG_NO_FIB_ENTRY",
	}
	IPNeighborFlags_value = map[string]uint8{
		"IP_API_NEIGHBOR_FLAG_NONE":         0,
		"IP_API_NEIGHBOR_FLAG_STATIC":       1,
		"IP_API_NEIGHBOR_FLAG_NO_FIB_ENTRY": 2,
	}
)

func (x IPNeighborFlags) String() string {
	s, ok := IPNeighborFlags_name[uint8(x)]
	if ok {
		return s
	}
	str := func(n uint8) string {
		s, ok := IPNeighborFlags_name[uint8(n)]
		if ok {
			return s
		}
		return "IPNeighborFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint8(0); i <= 8; i++ {
		val := uint8(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint8(x))
	}
	return s
}

// IPNeighbor defines type 'ip_neighbor'.
type IPNeighbor struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	Flags      IPNeighborFlags                `binapi:"ip_neighbor_flags,name=flags" json:"flags,omitempty"`
	MacAddress ethernet_types.MacAddress      `binapi:"mac_address,name=mac_address" json:"mac_address,omitempty"`
	IPAddress  ip_types.Address               `binapi:"address,name=ip_address" json:"ip_address,omitempty"`
}

// IP neighbor add / del request
//   - is_add - 1 to add neighbor, 0 to delete
//   - neighbor - the neighbor to add/remove
//
// IPNeighborAddDel defines message 'ip_neighbor_add_del'.
type IPNeighborAddDel struct {
	IsAdd    bool       `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	Neighbor IPNeighbor `binapi:"ip_neighbor,name=neighbor" json:"neighbor,omitempty"`
}

func (m *IPNeighborAddDel) Reset()               { *m = IPNeighborAddDel{} }
func (*IPNeighborAddDel) GetMessageName() string { return "ip_neighbor_add_del" }
func (*IPNeighborAddDel) GetCrcString() string   { return "0607c257" }
func (*IPNeighborAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPNeighborAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 4      // m.Neighbor.SwIfIndex
	size += 1      // m.Neighbor.Flags
	size += 1 * 6  // m.Neighbor.MacAddress
	size += 1      // m.Neighbor.IPAddress.Af
	size += 1 * 16 // m.Neighbor.IPAddress.Un
	return size
}
func (m *IPNeighborAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(uint32(m.Neighbor.SwIfIndex))
	buf.EncodeUint8(uint8(m.Neighbor.Flags))
	buf.EncodeBytes(m.Neighbor.MacAddress[:], 6)
	buf.EncodeUint8(uint8(m.Neighbor.IPAddress.Af))
	buf.EncodeBytes(m.Neighbor.IPAddress.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *IPNeighborAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.Neighbor.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Neighbor.Flags = IPNeighborFlags(buf.DecodeUint8())
	copy(m.Neighbor.MacAddress[:], buf.DecodeBytes(6))
	m.Neighbor.IPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Neighbor.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// IP neighbor add / del reply
//   - retval - return value
//   - stats_index - the index to use for this neighbor in the stats segment
//
// IPNeighborAddDelReply defines message 'ip_neighbor_add_del_reply'.
type IPNeighborAddDelReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	StatsIndex uint32 `binapi:"u32,name=stats_index" json:"stats_index,omitempty"`
}

func (m *IPNeighborAddDelReply) Reset()               { *m = IPNeighborAddDelReply{} }
func (*IPNeighborAddDelReply) GetMessageName() string { return "ip_neighbor_add_del_reply" }
func (*IPNeighborAddDelReply) GetCrcString() string   { return "1992deab" }
func (*IPNeighborAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPNeighborAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.StatsIndex
	return size
}
func (m *IPNeighborAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.StatsIndex)
	return buf.Bytes(), nil
}
func (m *IPNeighborAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.StatsIndex = buf.DecodeUint32()
	return nil
}

// Enable/disable periodic IP neighbor scan
//   - af - Address family v4/v6
//   - max_number - The maximum number of neighbours that will be created.
//     default 50k
//   - max_age - The maximum age (in seconds) before an inactive neighbour
//     is flushed
//     default 0 => never
//   - recycle - If max_number of neighbours is reached and new ones need
//     to be created should the oldest neighbour be 'recycled'.
//
// IPNeighborConfig defines message 'ip_neighbor_config'.
type IPNeighborConfig struct {
	Af        ip_types.AddressFamily `binapi:"address_family,name=af" json:"af,omitempty"`
	MaxNumber uint32                 `binapi:"u32,name=max_number" json:"max_number,omitempty"`
	MaxAge    uint32                 `binapi:"u32,name=max_age" json:"max_age,omitempty"`
	Recycle   bool                   `binapi:"bool,name=recycle" json:"recycle,omitempty"`
}

func (m *IPNeighborConfig) Reset()               { *m = IPNeighborConfig{} }
func (*IPNeighborConfig) GetMessageName() string { return "ip_neighbor_config" }
func (*IPNeighborConfig) GetCrcString() string   { return "f4a5cf44" }
func (*IPNeighborConfig) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPNeighborConfig) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Af
	size += 4 // m.MaxNumber
	size += 4 // m.MaxAge
	size += 1 // m.Recycle
	return size
}
func (m *IPNeighborConfig) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Af))
	buf.EncodeUint32(m.MaxNumber)
	buf.EncodeUint32(m.MaxAge)
	buf.EncodeBool(m.Recycle)
	return buf.Bytes(), nil
}
func (m *IPNeighborConfig) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Af = ip_types.AddressFamily(buf.DecodeUint8())
	m.MaxNumber = buf.DecodeUint32()
	m.MaxAge = buf.DecodeUint32()
	m.Recycle = buf.DecodeBool()
	return nil
}

// IPNeighborConfigReply defines message 'ip_neighbor_config_reply'.
type IPNeighborConfigReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IPNeighborConfigReply) Reset()               { *m = IPNeighborConfigReply{} }
func (*IPNeighborConfigReply) GetMessageName() string { return "ip_neighbor_config_reply" }
func (*IPNeighborConfigReply) GetCrcString() string   { return "e8d4e804" }
func (*IPNeighborConfigReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPNeighborConfigReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IPNeighborConfigReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IPNeighborConfigReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IP neighbors dump response
//   - age - time between last update and sending this message, in seconds
//   - neighbour - the neighbor
//
// IPNeighborDetails defines message 'ip_neighbor_details'.
type IPNeighborDetails struct {
	Age      float64    `binapi:"f64,name=age" json:"age,omitempty"`
	Neighbor IPNeighbor `binapi:"ip_neighbor,name=neighbor" json:"neighbor,omitempty"`
}

func (m *IPNeighborDetails) Reset()               { *m = IPNeighborDetails{} }
func (*IPNeighborDetails) GetMessageName() string { return "ip_neighbor_details" }
func (*IPNeighborDetails) GetCrcString() string   { return "e29d79f0" }
func (*IPNeighborDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPNeighborDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 8      // m.Age
	size += 4      // m.Neighbor.SwIfIndex
	size += 1      // m.Neighbor.Flags
	size += 1 * 6  // m.Neighbor.MacAddress
	size += 1      // m.Neighbor.IPAddress.Af
	size += 1 * 16 // m.Neighbor.IPAddress.Un
	return size
}
func (m *IPNeighborDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeFloat64(m.Age)
	buf.EncodeUint32(uint32(m.Neighbor.SwIfIndex))
	buf.EncodeUint8(uint8(m.Neighbor.Flags))
	buf.EncodeBytes(m.Neighbor.MacAddress[:], 6)
	buf.EncodeUint8(uint8(m.Neighbor.IPAddress.Af))
	buf.EncodeBytes(m.Neighbor.IPAddress.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *IPNeighborDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Age = buf.DecodeFloat64()
	m.Neighbor.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Neighbor.Flags = IPNeighborFlags(buf.DecodeUint8())
	copy(m.Neighbor.MacAddress[:], buf.DecodeBytes(6))
	m.Neighbor.IPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Neighbor.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// Dump IP neighbors
//   - sw_if_index - the interface to dump neighbors, ~0 == all
//   - af - address family is ipv[6|4]
//
// IPNeighborDump defines message 'ip_neighbor_dump'.
type IPNeighborDump struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
	Af        ip_types.AddressFamily         `binapi:"address_family,name=af" json:"af,omitempty"`
}

func (m *IPNeighborDump) Reset()               { *m = IPNeighborDump{} }
func (*IPNeighborDump) GetMessageName() string { return "ip_neighbor_dump" }
func (*IPNeighborDump) GetCrcString() string   { return "d817a484" }
func (*IPNeighborDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPNeighborDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	size += 1 // m.Af
	return size
}
func (m *IPNeighborDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeUint8(uint8(m.Af))
	return buf.Bytes(), nil
}
func (m *IPNeighborDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Af = ip_types.AddressFamily(buf.DecodeUint8())
	return nil
}

// Tell client about an IP4 ARP resolution event or
//
//	       MAC/IP info from ARP requests in L2 BDs
//	- pid - client pid registered to receive notification
//	- neighbor - new neighbor created
//
// IPNeighborEvent defines message 'ip_neighbor_event'.
// Deprecated: the message will be removed in the future versions
type IPNeighborEvent struct {
	PID      uint32     `binapi:"u32,name=pid" json:"pid,omitempty"`
	Neighbor IPNeighbor `binapi:"ip_neighbor,name=neighbor" json:"neighbor,omitempty"`
}

func (m *IPNeighborEvent) Reset()               { *m = IPNeighborEvent{} }
func (*IPNeighborEvent) GetMessageName() string { return "ip_neighbor_event" }
func (*IPNeighborEvent) GetCrcString() string   { return "bdb092b2" }
func (*IPNeighborEvent) GetMessageType() api.MessageType {
	return api.EventMessage
}

func (m *IPNeighborEvent) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.PID
	size += 4      // m.Neighbor.SwIfIndex
	size += 1      // m.Neighbor.Flags
	size += 1 * 6  // m.Neighbor.MacAddress
	size += 1      // m.Neighbor.IPAddress.Af
	size += 1 * 16 // m.Neighbor.IPAddress.Un
	return size
}
func (m *IPNeighborEvent) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PID)
	buf.EncodeUint32(uint32(m.Neighbor.SwIfIndex))
	buf.EncodeUint8(uint8(m.Neighbor.Flags))
	buf.EncodeBytes(m.Neighbor.MacAddress[:], 6)
	buf.EncodeUint8(uint8(m.Neighbor.IPAddress.Af))
	buf.EncodeBytes(m.Neighbor.IPAddress.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *IPNeighborEvent) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PID = buf.DecodeUint32()
	m.Neighbor.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Neighbor.Flags = IPNeighborFlags(buf.DecodeUint8())
	copy(m.Neighbor.MacAddress[:], buf.DecodeBytes(6))
	m.Neighbor.IPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Neighbor.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// Tell client about an IP4 ARP resolution event or
//
//	       MAC/IP info from ARP requests in L2 BDs
//	- pid - client pid registered to receive notification
//	- flags - Flags
//	- neighbor -  neighbor
//
// IPNeighborEventV2 defines message 'ip_neighbor_event_v2'.
type IPNeighborEventV2 struct {
	PID      uint32               `binapi:"u32,name=pid" json:"pid,omitempty"`
	Flags    IPNeighborEventFlags `binapi:"ip_neighbor_event_flags,name=flags" json:"flags,omitempty"`
	Neighbor IPNeighbor           `binapi:"ip_neighbor,name=neighbor" json:"neighbor,omitempty"`
}

func (m *IPNeighborEventV2) Reset()               { *m = IPNeighborEventV2{} }
func (*IPNeighborEventV2) GetMessageName() string { return "ip_neighbor_event_v2" }
func (*IPNeighborEventV2) GetCrcString() string   { return "c1d53dc0" }
func (*IPNeighborEventV2) GetMessageType() api.MessageType {
	return api.EventMessage
}

func (m *IPNeighborEventV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4      // m.PID
	size += 4      // m.Flags
	size += 4      // m.Neighbor.SwIfIndex
	size += 1      // m.Neighbor.Flags
	size += 1 * 6  // m.Neighbor.MacAddress
	size += 1      // m.Neighbor.IPAddress.Af
	size += 1 * 16 // m.Neighbor.IPAddress.Un
	return size
}
func (m *IPNeighborEventV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.PID)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeUint32(uint32(m.Neighbor.SwIfIndex))
	buf.EncodeUint8(uint8(m.Neighbor.Flags))
	buf.EncodeBytes(m.Neighbor.MacAddress[:], 6)
	buf.EncodeUint8(uint8(m.Neighbor.IPAddress.Af))
	buf.EncodeBytes(m.Neighbor.IPAddress.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *IPNeighborEventV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PID = buf.DecodeUint32()
	m.Flags = IPNeighborEventFlags(buf.DecodeUint32())
	m.Neighbor.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.Neighbor.Flags = IPNeighborFlags(buf.DecodeUint8())
	copy(m.Neighbor.MacAddress[:], buf.DecodeBytes(6))
	m.Neighbor.IPAddress.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.Neighbor.IPAddress.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// IP neighbor flush request - removes *all* neighbours.
//
//	 dynamic and static from API/CLI and dynamic from data-plane.
//	- af - Flush neighbours of this address family
//	- sw_if_index - Flush on this interface (~0 => all interfaces)
//
// IPNeighborFlush defines message 'ip_neighbor_flush'.
type IPNeighborFlush struct {
	Af        ip_types.AddressFamily         `binapi:"address_family,name=af" json:"af,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *IPNeighborFlush) Reset()               { *m = IPNeighborFlush{} }
func (*IPNeighborFlush) GetMessageName() string { return "ip_neighbor_flush" }
func (*IPNeighborFlush) GetCrcString() string   { return "16aa35d2" }
func (*IPNeighborFlush) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPNeighborFlush) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.Af
	size += 4 // m.SwIfIndex
	return size
}
func (m *IPNeighborFlush) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(uint8(m.Af))
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *IPNeighborFlush) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Af = ip_types.AddressFamily(buf.DecodeUint8())
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// IPNeighborFlushReply defines message 'ip_neighbor_flush_reply'.
type IPNeighborFlushReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IPNeighborFlushReply) Reset()               { *m = IPNeighborFlushReply{} }
func (*IPNeighborFlushReply) GetMessageName() string { return "ip_neighbor_flush_reply" }
func (*IPNeighborFlushReply) GetCrcString() string   { return "e8d4e804" }
func (*IPNeighborFlushReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPNeighborFlushReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IPNeighborFlushReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IPNeighborFlushReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IP neighbour replace begin
//
//	The use-case is that, for some unspecified reason, the control plane
//	has a different set of neighbours it than VPP
//	currently has. The CP would thus like to 'replace' VPP's set
//	only by specifying what the new set shall be, i.e. it is not
//	going to delete anything that already exists, rather, it wants any
//	unspecified neighbors deleted implicitly.
//	The CP declares the start of this procedure with this replace_begin
//	API Call, and when it has populated all neighbours it wants, it calls
//	the below replace_end API. From this point on it is of course free
//	to add and delete neighbours as usual.
//	The underlying mechanism by which VPP implements this replace is
//	intentionally left unspecified.
//
// IPNeighborReplaceBegin defines message 'ip_neighbor_replace_begin'.
type IPNeighborReplaceBegin struct{}

func (m *IPNeighborReplaceBegin) Reset()               { *m = IPNeighborReplaceBegin{} }
func (*IPNeighborReplaceBegin) GetMessageName() string { return "ip_neighbor_replace_begin" }
func (*IPNeighborReplaceBegin) GetCrcString() string   { return "51077d14" }
func (*IPNeighborReplaceBegin) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPNeighborReplaceBegin) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IPNeighborReplaceBegin) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IPNeighborReplaceBegin) Unmarshal(b []byte) error {
	return nil
}

// IPNeighborReplaceBeginReply defines message 'ip_neighbor_replace_begin_reply'.
type IPNeighborReplaceBeginReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IPNeighborReplaceBeginReply) Reset()               { *m = IPNeighborReplaceBeginReply{} }
func (*IPNeighborReplaceBeginReply) GetMessageName() string { return "ip_neighbor_replace_begin_reply" }
func (*IPNeighborReplaceBeginReply) GetCrcString() string   { return "e8d4e804" }
func (*IPNeighborReplaceBeginReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPNeighborReplaceBeginReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IPNeighborReplaceBeginReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IPNeighborReplaceBeginReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// IP neighbour replace end
//
//	see ip_neighbor_replace_begin description.
//
// IPNeighborReplaceEnd defines message 'ip_neighbor_replace_end'.
type IPNeighborReplaceEnd struct{}

func (m *IPNeighborReplaceEnd) Reset()               { *m = IPNeighborReplaceEnd{} }
func (*IPNeighborReplaceEnd) GetMessageName() string { return "ip_neighbor_replace_end" }
func (*IPNeighborReplaceEnd) GetCrcString() string   { return "51077d14" }
func (*IPNeighborReplaceEnd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IPNeighborReplaceEnd) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *IPNeighborReplaceEnd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *IPNeighborReplaceEnd) Unmarshal(b []byte) error {
	return nil
}

// IPNeighborReplaceEndReply defines message 'ip_neighbor_replace_end_reply'.
type IPNeighborReplaceEndReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IPNeighborReplaceEndReply) Reset()               { *m = IPNeighborReplaceEndReply{} }
func (*IPNeighborReplaceEndReply) GetMessageName() string { return "ip_neighbor_replace_end_reply" }
func (*IPNeighborReplaceEndReply) GetCrcString() string   { return "e8d4e804" }
func (*IPNeighborReplaceEndReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IPNeighborReplaceEndReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IPNeighborReplaceEndReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IPNeighborReplaceEndReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Register for IP neighbour events creation
//   - enable - 1 => register for events, 0 => cancel registration
//   - pid - sender's pid
//   - ip - exact IP address of interested neighbor resolution event
//   - sw_if_index - interface on which the IP address is present.
//
// WantIPNeighborEvents defines message 'want_ip_neighbor_events'.
// Deprecated: the message will be removed in the future versions
type WantIPNeighborEvents struct {
	Enable    bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
	PID       uint32                         `binapi:"u32,name=pid" json:"pid,omitempty"`
	IP        ip_types.Address               `binapi:"address,name=ip" json:"ip,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *WantIPNeighborEvents) Reset()               { *m = WantIPNeighborEvents{} }
func (*WantIPNeighborEvents) GetMessageName() string { return "want_ip_neighbor_events" }
func (*WantIPNeighborEvents) GetCrcString() string   { return "73e70a86" }
func (*WantIPNeighborEvents) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WantIPNeighborEvents) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Enable
	size += 4      // m.PID
	size += 1      // m.IP.Af
	size += 1 * 16 // m.IP.Un
	size += 4      // m.SwIfIndex
	return size
}
func (m *WantIPNeighborEvents) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Enable)
	buf.EncodeUint32(m.PID)
	buf.EncodeUint8(uint8(m.IP.Af))
	buf.EncodeBytes(m.IP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *WantIPNeighborEvents) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	m.PID = buf.DecodeUint32()
	m.IP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.IP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// WantIPNeighborEventsReply defines message 'want_ip_neighbor_events_reply'.
// Deprecated: the message will be removed in the future versions
type WantIPNeighborEventsReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WantIPNeighborEventsReply) Reset()               { *m = WantIPNeighborEventsReply{} }
func (*WantIPNeighborEventsReply) GetMessageName() string { return "want_ip_neighbor_events_reply" }
func (*WantIPNeighborEventsReply) GetCrcString() string   { return "e8d4e804" }
func (*WantIPNeighborEventsReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WantIPNeighborEventsReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WantIPNeighborEventsReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WantIPNeighborEventsReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Register for IP neighbour events (creation or deletion)
//   - enable - 1 => register for events, 0 => cancel registration
//   - pid - sender's pid
//   - ip - exact IP address of interested neighbor resolution event
//   - sw_if_index - interface on which the IP address is present.
//
// WantIPNeighborEventsV2 defines message 'want_ip_neighbor_events_v2'.
type WantIPNeighborEventsV2 struct {
	Enable    bool                           `binapi:"bool,name=enable" json:"enable,omitempty"`
	PID       uint32                         `binapi:"u32,name=pid" json:"pid,omitempty"`
	IP        ip_types.Address               `binapi:"address,name=ip" json:"ip,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index,default=4294967295" json:"sw_if_index,omitempty"`
}

func (m *WantIPNeighborEventsV2) Reset()               { *m = WantIPNeighborEventsV2{} }
func (*WantIPNeighborEventsV2) GetMessageName() string { return "want_ip_neighbor_events_v2" }
func (*WantIPNeighborEventsV2) GetCrcString() string   { return "73e70a86" }
func (*WantIPNeighborEventsV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *WantIPNeighborEventsV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.Enable
	size += 4      // m.PID
	size += 1      // m.IP.Af
	size += 1 * 16 // m.IP.Un
	size += 4      // m.SwIfIndex
	return size
}
func (m *WantIPNeighborEventsV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.Enable)
	buf.EncodeUint32(m.PID)
	buf.EncodeUint8(uint8(m.IP.Af))
	buf.EncodeBytes(m.IP.Un.XXX_UnionData[:], 16)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *WantIPNeighborEventsV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Enable = buf.DecodeBool()
	m.PID = buf.DecodeUint32()
	m.IP.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.IP.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// WantIPNeighborEventsV2Reply defines message 'want_ip_neighbor_events_v2_reply'.
type WantIPNeighborEventsV2Reply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *WantIPNeighborEventsV2Reply) Reset() { *m = WantIPNeighborEventsV2Reply{} }
func (*WantIPNeighborEventsV2Reply) GetMessageName() string {
	return "want_ip_neighbor_events_v2_reply"
}
func (*WantIPNeighborEventsV2Reply) GetCrcString() string { return "e8d4e804" }
func (*WantIPNeighborEventsV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *WantIPNeighborEventsV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *WantIPNeighborEventsV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *WantIPNeighborEventsV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ip_neighbor_binapi_init() }
func file_ip_neighbor_binapi_init() {
	api.RegisterMessage((*IPNeighborAddDel)(nil), "ip_neighbor_add_del_0607c257")
	api.RegisterMessage((*IPNeighborAddDelReply)(nil), "ip_neighbor_add_del_reply_1992deab")
	api.RegisterMessage((*IPNeighborConfig)(nil), "ip_neighbor_config_f4a5cf44")
	api.RegisterMessage((*IPNeighborConfigReply)(nil), "ip_neighbor_config_reply_e8d4e804")
	api.RegisterMessage((*IPNeighborDetails)(nil), "ip_neighbor_details_e29d79f0")
	api.RegisterMessage((*IPNeighborDump)(nil), "ip_neighbor_dump_d817a484")
	api.RegisterMessage((*IPNeighborEvent)(nil), "ip_neighbor_event_bdb092b2")
	api.RegisterMessage((*IPNeighborEventV2)(nil), "ip_neighbor_event_v2_c1d53dc0")
	api.RegisterMessage((*IPNeighborFlush)(nil), "ip_neighbor_flush_16aa35d2")
	api.RegisterMessage((*IPNeighborFlushReply)(nil), "ip_neighbor_flush_reply_e8d4e804")
	api.RegisterMessage((*IPNeighborReplaceBegin)(nil), "ip_neighbor_replace_begin_51077d14")
	api.RegisterMessage((*IPNeighborReplaceBeginReply)(nil), "ip_neighbor_replace_begin_reply_e8d4e804")
	api.RegisterMessage((*IPNeighborReplaceEnd)(nil), "ip_neighbor_replace_end_51077d14")
	api.RegisterMessage((*IPNeighborReplaceEndReply)(nil), "ip_neighbor_replace_end_reply_e8d4e804")
	api.RegisterMessage((*WantIPNeighborEvents)(nil), "want_ip_neighbor_events_73e70a86")
	api.RegisterMessage((*WantIPNeighborEventsReply)(nil), "want_ip_neighbor_events_reply_e8d4e804")
	api.RegisterMessage((*WantIPNeighborEventsV2)(nil), "want_ip_neighbor_events_v2_73e70a86")
	api.RegisterMessage((*WantIPNeighborEventsV2Reply)(nil), "want_ip_neighbor_events_v2_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IPNeighborAddDel)(nil),
		(*IPNeighborAddDelReply)(nil),
		(*IPNeighborConfig)(nil),
		(*IPNeighborConfigReply)(nil),
		(*IPNeighborDetails)(nil),
		(*IPNeighborDump)(nil),
		(*IPNeighborEvent)(nil),
		(*IPNeighborEventV2)(nil),
		(*IPNeighborFlush)(nil),
		(*IPNeighborFlushReply)(nil),
		(*IPNeighborReplaceBegin)(nil),
		(*IPNeighborReplaceBeginReply)(nil),
		(*IPNeighborReplaceEnd)(nil),
		(*IPNeighborReplaceEndReply)(nil),
		(*WantIPNeighborEvents)(nil),
		(*WantIPNeighborEventsReply)(nil),
		(*WantIPNeighborEventsV2)(nil),
		(*WantIPNeighborEventsV2Reply)(nil),
	}
}
