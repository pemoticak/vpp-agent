// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/vpp.proto

package vpp

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	abf "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/abf"
	acl "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/acl"
	interfaces "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/interfaces"
	ipsec "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/ipsec"
	l2 "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/l2"
	l3 "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/l3"
	nat "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/nat"
	punt "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/punt"
	srv6 "go.ligato.io/vpp-agent/v2/proto/ligato/vpp/srv6"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ConfigData struct {
	Interfaces           []*interfaces.Interface         `protobuf:"bytes,10,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	Spans                []*interfaces.Span              `protobuf:"bytes,11,rep,name=spans,proto3" json:"spans,omitempty"`
	Acls                 []*acl.ACL                      `protobuf:"bytes,20,rep,name=acls,proto3" json:"acls,omitempty"`
	Abfs                 []*abf.ABF                      `protobuf:"bytes,21,rep,name=abfs,proto3" json:"abfs,omitempty"`
	BridgeDomains        []*l2.BridgeDomain              `protobuf:"bytes,30,rep,name=bridge_domains,json=bridgeDomains,proto3" json:"bridge_domains,omitempty"`
	Fibs                 []*l2.FIBEntry                  `protobuf:"bytes,31,rep,name=fibs,proto3" json:"fibs,omitempty"`
	XconnectPairs        []*l2.XConnectPair              `protobuf:"bytes,32,rep,name=xconnect_pairs,json=xconnectPairs,proto3" json:"xconnect_pairs,omitempty"`
	Routes               []*l3.Route                     `protobuf:"bytes,40,rep,name=routes,proto3" json:"routes,omitempty"`
	Arps                 []*l3.ARPEntry                  `protobuf:"bytes,41,rep,name=arps,proto3" json:"arps,omitempty"`
	ProxyArp             *l3.ProxyARP                    `protobuf:"bytes,42,opt,name=proxy_arp,json=proxyArp,proto3" json:"proxy_arp,omitempty"`
	IpscanNeighbor       *l3.IPScanNeighbor              `protobuf:"bytes,43,opt,name=ipscan_neighbor,json=ipscanNeighbor,proto3" json:"ipscan_neighbor,omitempty"`
	Vrfs                 []*l3.VrfTable                  `protobuf:"bytes,44,rep,name=vrfs,proto3" json:"vrfs,omitempty"`
	Nat44Global          *nat.Nat44Global                `protobuf:"bytes,50,opt,name=nat44_global,json=nat44Global,proto3" json:"nat44_global,omitempty"`
	Dnat44S              []*nat.DNat44                   `protobuf:"bytes,51,rep,name=dnat44s,proto3" json:"dnat44s,omitempty"`
	Nat44Interfaces      []*nat.Nat44Interface           `protobuf:"bytes,52,rep,name=nat44_interfaces,json=nat44Interfaces,proto3" json:"nat44_interfaces,omitempty"`
	Nat44Pools           []*nat.Nat44AddressPool         `protobuf:"bytes,53,rep,name=nat44_pools,json=nat44Pools,proto3" json:"nat44_pools,omitempty"`
	IpsecSpds            []*ipsec.SecurityPolicyDatabase `protobuf:"bytes,60,rep,name=ipsec_spds,json=ipsecSpds,proto3" json:"ipsec_spds,omitempty"`
	IpsecSas             []*ipsec.SecurityAssociation    `protobuf:"bytes,61,rep,name=ipsec_sas,json=ipsecSas,proto3" json:"ipsec_sas,omitempty"`
	PuntIpredirects      []*punt.IPRedirect              `protobuf:"bytes,70,rep,name=punt_ipredirects,json=puntIpredirects,proto3" json:"punt_ipredirects,omitempty"`
	PuntTohosts          []*punt.ToHost                  `protobuf:"bytes,71,rep,name=punt_tohosts,json=puntTohosts,proto3" json:"punt_tohosts,omitempty"`
	PuntExceptions       []*punt.Exception               `protobuf:"bytes,72,rep,name=punt_exceptions,json=puntExceptions,proto3" json:"punt_exceptions,omitempty"`
	Srv6Localsids        []*srv6.LocalSID                `protobuf:"bytes,80,rep,name=srv6_localsids,json=srv6Localsids,proto3" json:"srv6_localsids,omitempty"`
	Srv6Policies         []*srv6.Policy                  `protobuf:"bytes,81,rep,name=srv6_policies,json=srv6Policies,proto3" json:"srv6_policies,omitempty"`
	Srv6Steerings        []*srv6.Steering                `protobuf:"bytes,82,rep,name=srv6_steerings,json=srv6Steerings,proto3" json:"srv6_steerings,omitempty"`
	Srv6Global           *srv6.SRv6Global                `protobuf:"bytes,83,opt,name=srv6_global,json=srv6Global,proto3" json:"srv6_global,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ConfigData) Reset()         { *m = ConfigData{} }
func (m *ConfigData) String() string { return proto.CompactTextString(m) }
func (*ConfigData) ProtoMessage()    {}
func (*ConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_0138a1608d5d59f2, []int{0}
}

func (m *ConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigData.Unmarshal(m, b)
}
func (m *ConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigData.Marshal(b, m, deterministic)
}
func (m *ConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigData.Merge(m, src)
}
func (m *ConfigData) XXX_Size() int {
	return xxx_messageInfo_ConfigData.Size(m)
}
func (m *ConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigData proto.InternalMessageInfo

func (m *ConfigData) GetInterfaces() []*interfaces.Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *ConfigData) GetSpans() []*interfaces.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

func (m *ConfigData) GetAcls() []*acl.ACL {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *ConfigData) GetAbfs() []*abf.ABF {
	if m != nil {
		return m.Abfs
	}
	return nil
}

func (m *ConfigData) GetBridgeDomains() []*l2.BridgeDomain {
	if m != nil {
		return m.BridgeDomains
	}
	return nil
}

func (m *ConfigData) GetFibs() []*l2.FIBEntry {
	if m != nil {
		return m.Fibs
	}
	return nil
}

func (m *ConfigData) GetXconnectPairs() []*l2.XConnectPair {
	if m != nil {
		return m.XconnectPairs
	}
	return nil
}

func (m *ConfigData) GetRoutes() []*l3.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *ConfigData) GetArps() []*l3.ARPEntry {
	if m != nil {
		return m.Arps
	}
	return nil
}

func (m *ConfigData) GetProxyArp() *l3.ProxyARP {
	if m != nil {
		return m.ProxyArp
	}
	return nil
}

func (m *ConfigData) GetIpscanNeighbor() *l3.IPScanNeighbor {
	if m != nil {
		return m.IpscanNeighbor
	}
	return nil
}

func (m *ConfigData) GetVrfs() []*l3.VrfTable {
	if m != nil {
		return m.Vrfs
	}
	return nil
}

func (m *ConfigData) GetNat44Global() *nat.Nat44Global {
	if m != nil {
		return m.Nat44Global
	}
	return nil
}

func (m *ConfigData) GetDnat44S() []*nat.DNat44 {
	if m != nil {
		return m.Dnat44S
	}
	return nil
}

func (m *ConfigData) GetNat44Interfaces() []*nat.Nat44Interface {
	if m != nil {
		return m.Nat44Interfaces
	}
	return nil
}

func (m *ConfigData) GetNat44Pools() []*nat.Nat44AddressPool {
	if m != nil {
		return m.Nat44Pools
	}
	return nil
}

func (m *ConfigData) GetIpsecSpds() []*ipsec.SecurityPolicyDatabase {
	if m != nil {
		return m.IpsecSpds
	}
	return nil
}

func (m *ConfigData) GetIpsecSas() []*ipsec.SecurityAssociation {
	if m != nil {
		return m.IpsecSas
	}
	return nil
}

func (m *ConfigData) GetPuntIpredirects() []*punt.IPRedirect {
	if m != nil {
		return m.PuntIpredirects
	}
	return nil
}

func (m *ConfigData) GetPuntTohosts() []*punt.ToHost {
	if m != nil {
		return m.PuntTohosts
	}
	return nil
}

func (m *ConfigData) GetPuntExceptions() []*punt.Exception {
	if m != nil {
		return m.PuntExceptions
	}
	return nil
}

func (m *ConfigData) GetSrv6Localsids() []*srv6.LocalSID {
	if m != nil {
		return m.Srv6Localsids
	}
	return nil
}

func (m *ConfigData) GetSrv6Policies() []*srv6.Policy {
	if m != nil {
		return m.Srv6Policies
	}
	return nil
}

func (m *ConfigData) GetSrv6Steerings() []*srv6.Steering {
	if m != nil {
		return m.Srv6Steerings
	}
	return nil
}

func (m *ConfigData) GetSrv6Global() *srv6.SRv6Global {
	if m != nil {
		return m.Srv6Global
	}
	return nil
}

type Notification struct {
	Interface            *interfaces.InterfaceNotification `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_0138a1608d5d59f2, []int{1}
}

func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetInterface() *interfaces.InterfaceNotification {
	if m != nil {
		return m.Interface
	}
	return nil
}

type Stats struct {
	Interface            *interfaces.InterfaceStats `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Stats) Reset()         { *m = Stats{} }
func (m *Stats) String() string { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()    {}
func (*Stats) Descriptor() ([]byte, []int) {
	return fileDescriptor_0138a1608d5d59f2, []int{2}
}

func (m *Stats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stats.Unmarshal(m, b)
}
func (m *Stats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stats.Marshal(b, m, deterministic)
}
func (m *Stats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stats.Merge(m, src)
}
func (m *Stats) XXX_Size() int {
	return xxx_messageInfo_Stats.Size(m)
}
func (m *Stats) XXX_DiscardUnknown() {
	xxx_messageInfo_Stats.DiscardUnknown(m)
}

var xxx_messageInfo_Stats proto.InternalMessageInfo

func (m *Stats) GetInterface() *interfaces.InterfaceStats {
	if m != nil {
		return m.Interface
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigData)(nil), "ligato.vpp.ConfigData")
	proto.RegisterType((*Notification)(nil), "ligato.vpp.Notification")
	proto.RegisterType((*Stats)(nil), "ligato.vpp.Stats")
}

func init() { proto.RegisterFile("ligato/vpp/vpp.proto", fileDescriptor_0138a1608d5d59f2) }

var fileDescriptor_0138a1608d5d59f2 = []byte{
	// 902 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0x6f, 0x6f, 0xdb, 0x36,
	0x10, 0xc6, 0x51, 0xac, 0xed, 0x1a, 0x3a, 0x7f, 0x0a, 0x2e, 0x6b, 0xd9, 0x6c, 0xeb, 0xbc, 0x00,
	0xc1, 0xb2, 0x36, 0xb3, 0x37, 0x3b, 0xeb, 0x8b, 0xa1, 0x1b, 0x6a, 0x27, 0x75, 0xea, 0x21, 0x08,
	0x34, 0x3a, 0x18, 0x86, 0xbe, 0x11, 0x28, 0x99, 0x72, 0x09, 0x68, 0x24, 0xc1, 0x63, 0x8c, 0xe4,
	0xfb, 0xee, 0x83, 0x14, 0x3c, 0x49, 0xb6, 0x1c, 0x2b, 0xc8, 0x0b, 0x19, 0x3a, 0xde, 0xef, 0x79,
	0x8e, 0x26, 0x8f, 0x14, 0xd9, 0xcd, 0xd5, 0x4c, 0x78, 0xd3, 0x9d, 0x5b, 0x1b, 0x9e, 0x8e, 0x75,
	0xc6, 0x1b, 0x4a, 0x8a, 0xd1, 0xce, 0xdc, 0xda, 0x3d, 0x56, 0x23, 0x44, 0x92, 0x85, 0xa7, 0xa0,
	0x56, 0x33, 0x69, 0x1e, 0x9e, 0x32, 0x73, 0x50, 0xcb, 0x28, 0xed, 0xa5, 0xcb, 0x44, 0x2a, 0x61,
	0xf9, 0x5a, 0x62, 0xed, 0x66, 0x0c, 0xac, 0xd0, 0x25, 0xf1, 0xc3, 0x1d, 0x84, 0x17, 0xbe, 0x32,
	0xf9, 0xb6, 0x8e, 0x58, 0x90, 0x69, 0xf1, 0xdb, 0x60, 0x90, 0xf7, 0xba, 0x89, 0x53, 0xd3, 0x99,
	0x8c, 0xa7, 0xe6, 0x3f, 0xa1, 0xaa, 0x1a, 0xcf, 0x57, 0x91, 0x4c, 0x25, 0x0d, 0xce, 0x79, 0xaf,
	0x7b, 0x9d, 0x1a, 0xad, 0x65, 0xea, 0x9b, 0x64, 0xfd, 0xae, 0x70, 0xe5, 0xe2, 0xed, 0x3d, 0x5b,
	0x4d, 0xe4, 0xfd, 0x72, 0xfc, 0xc5, 0xea, 0xb8, 0x33, 0x57, 0x8b, 0xff, 0x70, 0xcb, 0x6b, 0xee,
	0x9a, 0x96, 0x58, 0x0b, 0x1f, 0x9e, 0x32, 0xb3, 0x57, 0xcb, 0xd8, 0x2b, 0xed, 0xf1, 0xa7, 0x21,
	0x07, 0x6e, 0xfe, 0x06, 0x7f, 0x8a, 0xdc, 0xfe, 0xff, 0x2d, 0x42, 0x4e, 0x8c, 0xce, 0xd4, 0xec,
	0x54, 0x78, 0x41, 0xdf, 0x11, 0xb2, 0x5c, 0x57, 0x46, 0xda, 0x5f, 0x1c, 0xb6, 0x7a, 0xed, 0xce,
	0x72, 0xfb, 0x3b, 0xcb, 0x6c, 0x67, 0x5c, 0xbd, 0xf2, 0x9a, 0x86, 0xfe, 0x4a, 0x1e, 0x85, 0x0d,
	0x03, 0xd6, 0x42, 0xf1, 0x37, 0x77, 0x88, 0x27, 0x56, 0x68, 0x5e, 0x90, 0xf4, 0x47, 0xf2, 0x50,
	0xa4, 0x39, 0xb0, 0x5d, 0x54, 0x7c, 0x55, 0x57, 0x84, 0x1e, 0x1a, 0x9c, 0x9c, 0x73, 0x04, 0x10,
	0x4c, 0x32, 0x60, 0x5f, 0x37, 0x80, 0x49, 0xd6, 0x19, 0x0c, 0x47, 0x1c, 0x01, 0x3a, 0x24, 0xdb,
	0x2b, 0x5b, 0x0b, 0xec, 0xe5, 0xfa, 0x6c, 0xf2, 0x5e, 0x67, 0x88, 0xd0, 0x29, 0x32, 0x7c, 0x2b,
	0xa9, 0x45, 0x40, 0x5f, 0x93, 0x87, 0x99, 0x4a, 0x80, 0x7d, 0x8f, 0xca, 0xe7, 0xb7, 0x94, 0xa3,
	0xf1, 0xf0, 0xbd, 0xf6, 0xee, 0x86, 0x23, 0x14, 0x0a, 0x56, 0xfd, 0x10, 0x5b, 0xa1, 0x1c, 0xb0,
	0x76, 0x63, 0xc1, 0x7f, 0x4f, 0x0a, 0x28, 0x12, 0xca, 0xf1, 0xad, 0x4a, 0x12, 0x22, 0xa0, 0x47,
	0xe4, 0x31, 0x36, 0x01, 0xb0, 0x43, 0xd4, 0xee, 0xae, 0x68, 0xfb, 0x1d, 0x1e, 0x92, 0xbc, 0x64,
	0xc2, 0xf4, 0x84, 0xb3, 0xc0, 0x7e, 0x6a, 0x98, 0x5e, 0xbf, 0x33, 0xe0, 0x51, 0x39, 0xbd, 0x00,
	0xd1, 0x63, 0xb2, 0x61, 0x9d, 0xb9, 0xbe, 0x89, 0x85, 0xb3, 0xec, 0x55, 0xfb, 0x41, 0x83, 0x22,
	0x0a, 0xf9, 0x01, 0x8f, 0xf8, 0x13, 0x24, 0x07, 0xce, 0xd2, 0x11, 0xd9, 0x51, 0x16, 0x52, 0xa1,
	0x63, 0x2d, 0xd5, 0xec, 0x53, 0x62, 0x1c, 0x7b, 0x8d, 0xda, 0xef, 0x6e, 0x69, 0xc7, 0xd1, 0x24,
	0x15, 0xfa, 0xa2, 0x84, 0xf8, 0x76, 0xa1, 0xaa, 0xe2, 0x30, 0xd5, 0xb9, 0xcb, 0x80, 0x1d, 0x35,
	0x4e, 0xf5, 0x1f, 0x97, 0x5d, 0x8a, 0x24, 0x97, 0x1c, 0x21, 0xfa, 0x27, 0xd9, 0xd4, 0xc2, 0x1f,
	0x1f, 0xc7, 0xb3, 0xdc, 0x24, 0x22, 0x67, 0x3d, 0xac, 0xb8, 0xb2, 0x8e, 0xa1, 0xeb, 0x2f, 0x02,
	0x73, 0x86, 0x08, 0x6f, 0xe9, 0x65, 0x40, 0x7f, 0x21, 0x5f, 0x4e, 0x31, 0x06, 0xd6, 0xc7, 0x7a,
	0xcf, 0x6e, 0x4b, 0x4f, 0x51, 0xcb, 0x2b, 0x8c, 0x8e, 0xc9, 0xd3, 0xa2, 0x62, 0xad, 0xf3, 0x8f,
	0x51, 0xfa, 0xb2, 0xb1, 0xea, 0xb2, 0xef, 0x77, 0xf4, 0x4a, 0x0c, 0x74, 0x40, 0x8a, 0xb9, 0xc4,
	0xd6, 0x98, 0x1c, 0xd8, 0x6f, 0xeb, 0xe7, 0x67, 0xe1, 0x32, 0x98, 0x4e, 0x9d, 0x04, 0x88, 0x8c,
	0xc9, 0x39, 0x41, 0x51, 0x78, 0x05, 0x7a, 0x46, 0x08, 0x5e, 0x58, 0x31, 0xd8, 0x29, 0xb0, 0xb7,
	0xe8, 0x70, 0xb8, 0x72, 0x88, 0xf0, 0x3a, 0x9b, 0xc8, 0xf4, 0xca, 0x29, 0x7f, 0x13, 0x99, 0x5c,
	0xa5, 0x37, 0xe1, 0xec, 0x26, 0x02, 0x24, 0xdf, 0xc0, 0xec, 0xc4, 0x4e, 0x43, 0x4b, 0x6e, 0x94,
	0x46, 0x02, 0xd8, 0x1f, 0xe8, 0x73, 0x70, 0xb7, 0xcf, 0x00, 0xc0, 0xa4, 0x4a, 0x78, 0x65, 0x34,
	0x7f, 0x52, 0x98, 0x08, 0xa0, 0x23, 0xf2, 0x34, 0xdc, 0x23, 0xb1, 0xb2, 0x4e, 0x4e, 0x95, 0x93,
	0xa9, 0x07, 0x36, 0x5a, 0x6f, 0x6c, 0xbc, 0x6b, 0xc6, 0x11, 0x2f, 0x19, 0xbe, 0x13, 0x06, 0xc6,
	0x4b, 0x0d, 0xfd, 0x9d, 0x6c, 0xa2, 0x8f, 0x37, 0x9f, 0x0c, 0x78, 0x60, 0x67, 0xeb, 0x9d, 0x80,
	0x1e, 0x97, 0xe6, 0x83, 0x01, 0xcf, 0x5b, 0x21, 0xb8, 0x2c, 0x58, 0x7a, 0x42, 0xd0, 0x2e, 0x96,
	0xd7, 0xa9, 0xb4, 0x61, 0x7e, 0xc0, 0x3e, 0xa0, 0x7c, 0x6f, 0x4d, 0xfe, 0xbe, 0x42, 0xf8, 0x76,
	0x88, 0x17, 0x21, 0xd0, 0x77, 0x64, 0x3b, 0x5c, 0x7a, 0x71, 0x6e, 0x52, 0x91, 0x83, 0x9a, 0x02,
	0x8b, 0xd0, 0xe3, 0x45, 0xdd, 0x03, 0xaf, 0xc5, 0xf3, 0x40, 0x4c, 0xc6, 0xa7, 0x7c, 0x2b, 0x84,
	0xe7, 0x15, 0x4f, 0xdf, 0x12, 0x1c, 0x88, 0x6d, 0x58, 0x70, 0x25, 0x81, 0xfd, 0xbd, 0xfe, 0x1f,
	0xd0, 0xa0, 0xd8, 0x11, 0xbe, 0x19, 0x82, 0xa8, 0x84, 0x17, 0xf5, 0xc1, 0x4b, 0xe9, 0x94, 0x9e,
	0x01, 0xe3, 0x77, 0xd4, 0x9f, 0x94, 0x44, 0x51, 0xbf, 0x8a, 0x42, 0xfd, 0x16, 0x3a, 0x94, 0xc7,
	0x62, 0xb2, 0x7e, 0x2c, 0x0a, 0x39, 0x9f, 0xbf, 0x29, 0x8f, 0x05, 0x09, 0x03, 0xc5, 0xfb, 0xfe,
	0x47, 0xb2, 0x79, 0x61, 0xbc, 0xca, 0x54, 0x8a, 0x5b, 0x4c, 0xff, 0x22, 0x1b, 0x8b, 0x6e, 0x67,
	0x0f, 0xd0, 0xeb, 0xe8, 0xbe, 0x6b, 0xbe, 0x6e, 0xc0, 0x97, 0xf2, 0xfd, 0x73, 0xf2, 0x68, 0xe2,
	0x05, 0xee, 0xd4, 0x9a, 0xe9, 0xc1, 0x7d, 0xa6, 0xa8, 0xac, 0xb9, 0x0d, 0x8f, 0x3e, 0xbe, 0x9a,
	0x99, 0x4a, 0xa5, 0xf0, 0xa3, 0xf5, 0xb3, 0x98, 0x49, 0xed, 0xbb, 0xf3, 0x5e, 0x17, 0xbf, 0x59,
	0xdd, 0xe5, 0xe7, 0x2c, 0x79, 0x8c, 0x23, 0xfd, 0xcf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x36,
	0x23, 0xb2, 0xb8, 0x08, 0x00, 0x00,
}
