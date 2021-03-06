// Code generated by protoc-gen-gogo.
// source: ipam.proto
// DO NOT EDIT!

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IPAMRequest struct {
	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Subnet      string `protobuf:"bytes,2,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (m *IPAMRequest) Reset()                    { *m = IPAMRequest{} }
func (m *IPAMRequest) String() string            { return proto.CompactTextString(m) }
func (*IPAMRequest) ProtoMessage()               {}
func (*IPAMRequest) Descriptor() ([]byte, []int) { return fileDescriptorIpam, []int{0} }

func (m *IPAMRequest) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *IPAMRequest) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

type IPConfig struct {
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
}

func (m *IPConfig) Reset()                    { *m = IPConfig{} }
func (m *IPConfig) String() string            { return proto.CompactTextString(m) }
func (*IPConfig) ProtoMessage()               {}
func (*IPConfig) Descriptor() ([]byte, []int) { return fileDescriptorIpam, []int{1} }

func (m *IPConfig) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *IPConfig) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *IPConfig) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

type DNS struct {
	Nameservers []string `protobuf:"bytes,1,rep,name=nameservers" json:"nameservers,omitempty"`
	Domain      string   `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Search      []string `protobuf:"bytes,3,rep,name=search" json:"search,omitempty"`
	Options     []string `protobuf:"bytes,4,rep,name=options" json:"options,omitempty"`
}

func (m *DNS) Reset()                    { *m = DNS{} }
func (m *DNS) String() string            { return proto.CompactTextString(m) }
func (*DNS) ProtoMessage()               {}
func (*DNS) Descriptor() ([]byte, []int) { return fileDescriptorIpam, []int{2} }

func (m *DNS) GetNameservers() []string {
	if m != nil {
		return m.Nameservers
	}
	return nil
}

func (m *DNS) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *DNS) GetSearch() []string {
	if m != nil {
		return m.Search
	}
	return nil
}

func (m *DNS) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

type IPAMResponse struct {
	IP  *IPConfig `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	DNS *DNS      `protobuf:"bytes,2,opt,name=dns" json:"dns,omitempty"`
}

func (m *IPAMResponse) Reset()                    { *m = IPAMResponse{} }
func (m *IPAMResponse) String() string            { return proto.CompactTextString(m) }
func (*IPAMResponse) ProtoMessage()               {}
func (*IPAMResponse) Descriptor() ([]byte, []int) { return fileDescriptorIpam, []int{3} }

func (m *IPAMResponse) GetIP() *IPConfig {
	if m != nil {
		return m.IP
	}
	return nil
}

func (m *IPAMResponse) GetDNS() *DNS {
	if m != nil {
		return m.DNS
	}
	return nil
}

type IPReleaseResponse struct {
}

func (m *IPReleaseResponse) Reset()                    { *m = IPReleaseResponse{} }
func (m *IPReleaseResponse) String() string            { return proto.CompactTextString(m) }
func (*IPReleaseResponse) ProtoMessage()               {}
func (*IPReleaseResponse) Descriptor() ([]byte, []int) { return fileDescriptorIpam, []int{4} }

type IPReleaseRequest struct {
	ContainerID string `protobuf:"bytes,1,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *IPReleaseRequest) Reset()                    { *m = IPReleaseRequest{} }
func (m *IPReleaseRequest) String() string            { return proto.CompactTextString(m) }
func (*IPReleaseRequest) ProtoMessage()               {}
func (*IPReleaseRequest) Descriptor() ([]byte, []int) { return fileDescriptorIpam, []int{5} }

func (m *IPReleaseRequest) GetContainerID() string {
	if m != nil {
		return m.ContainerID
	}
	return ""
}

func (m *IPReleaseRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*IPAMRequest)(nil), "api.IPAMRequest")
	proto.RegisterType((*IPConfig)(nil), "api.IPConfig")
	proto.RegisterType((*DNS)(nil), "api.DNS")
	proto.RegisterType((*IPAMResponse)(nil), "api.IPAMResponse")
	proto.RegisterType((*IPReleaseResponse)(nil), "api.IPReleaseResponse")
	proto.RegisterType((*IPReleaseRequest)(nil), "api.IPReleaseRequest")
}

func init() { proto.RegisterFile("ipam.proto", fileDescriptorIpam) }

var fileDescriptorIpam = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0xc7, 0xd5, 0x04, 0xf5, 0xe3, 0x52, 0x04, 0x18, 0x84, 0x2c, 0x96, 0x56, 0x41, 0x48, 0x9d,
	0x8a, 0x54, 0x9e, 0x80, 0xb6, 0x4b, 0x06, 0xaa, 0xc8, 0x5d, 0x80, 0x05, 0xdc, 0xe6, 0x08, 0x96,
	0xa8, 0xed, 0xc6, 0x29, 0x88, 0x97, 0xed, 0x90, 0x27, 0x41, 0x76, 0x62, 0xd4, 0x85, 0x85, 0x2d,
	0xbf, 0xfb, 0xdf, 0xd7, 0xff, 0x62, 0x00, 0xa1, 0xf9, 0x66, 0xac, 0x0b, 0x55, 0x2a, 0x12, 0x72,
	0x2d, 0xae, 0x2e, 0x72, 0x95, 0x2b, 0xc7, 0xb7, 0xf6, 0xab, 0x96, 0xe2, 0x27, 0x88, 0x92, 0xf4,
	0xfe, 0x81, 0xe1, 0x76, 0x87, 0xa6, 0x24, 0x13, 0xe8, 0xaf, 0x95, 0x2c, 0xb9, 0x90, 0x58, 0xbc,
	0x88, 0x8c, 0xb6, 0x86, 0xad, 0x51, 0x6f, 0x7a, 0x52, 0xed, 0x07, 0xd1, 0xcc, 0xc7, 0x93, 0x39,
	0x8b, 0x7e, 0x93, 0x92, 0x8c, 0x5c, 0x42, 0xdb, 0xec, 0x56, 0x12, 0x4b, 0x1a, 0xd8, 0x6c, 0xd6,
	0x50, 0xfc, 0x08, 0xdd, 0x24, 0x9d, 0x29, 0xf9, 0x26, 0x72, 0x42, 0xa1, 0xf3, 0x89, 0x85, 0x11,
	0x4a, 0xd6, 0x2d, 0x99, 0x47, 0xab, 0xf0, 0x2c, 0x2b, 0xd0, 0x98, 0xa6, 0xdc, 0xa3, 0x55, 0x72,
	0x5e, 0xe2, 0x17, 0xff, 0xa6, 0x61, 0xad, 0x34, 0x18, 0x6f, 0x21, 0x9c, 0x2f, 0x96, 0x64, 0x08,
	0x91, 0xe4, 0x1b, 0x34, 0x58, 0xd8, 0x66, 0xb4, 0x35, 0x0c, 0x47, 0x3d, 0x76, 0x18, 0xb2, 0xab,
	0x65, 0x6a, 0xc3, 0x85, 0xf4, 0xab, 0xd5, 0xe4, 0x56, 0x46, 0x5e, 0xac, 0xdf, 0x69, 0xe8, 0x8a,
	0x1a, 0xb2, 0x23, 0x95, 0x2e, 0x85, 0x92, 0x86, 0x1e, 0x39, 0xc1, 0x63, 0xfc, 0x0c, 0xfd, 0xfa,
	0x4e, 0x46, 0x2b, 0x69, 0x90, 0xdc, 0x40, 0x20, 0xb4, 0xf3, 0x12, 0x4d, 0x8e, 0xc7, 0x5c, 0x8b,
	0xb1, 0xf7, 0x3a, 0x6d, 0x57, 0xfb, 0x41, 0x90, 0xa4, 0x2c, 0x10, 0x9a, 0x5c, 0x43, 0x98, 0xc9,
	0xda, 0x59, 0x34, 0xe9, 0xba, 0xbc, 0xf9, 0x62, 0x39, 0xed, 0x54, 0xfb, 0x81, 0xb5, 0xc0, 0xac,
	0x1a, 0x9f, 0xc3, 0x59, 0x92, 0x32, 0xfc, 0x40, 0x6e, 0xd0, 0x0f, 0x88, 0x5f, 0xe1, 0xf4, 0x20,
	0xf8, 0xff, 0xbf, 0xf3, 0xe7, 0x7d, 0x57, 0x6d, 0xf7, 0x02, 0xee, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x27, 0x39, 0xce, 0xb9, 0x2a, 0x02, 0x00, 0x00,
}
