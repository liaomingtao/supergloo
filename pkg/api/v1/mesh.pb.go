// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mesh.proto

package v1 // import "github.com/solo-io/supergloo/pkg/api/v1"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import bytes "bytes"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// TODO: Eventually we want to plug in new meshes easier, but for now it's useful to enumerate in the config
type MeshType int32

const (
	MeshType_ISTIO    MeshType = 0
	MeshType_LINKERD1 MeshType = 1
)

var MeshType_name = map[int32]string{
	0: "ISTIO",
	1: "LINKERD1",
}
var MeshType_value = map[string]int32{
	"ISTIO":    0,
	"LINKERD1": 1,
}

func (x MeshType) String() string {
	return proto.EnumName(MeshType_name, int32(x))
}
func (MeshType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{0}
}

// Any user-configurable settings for a service mesh
// This isn't meant to cover install or initial configuration of the mesh
type MeshConfig struct {
	Metadata             *MeshMetadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Encryption           *Encryption   `protobuf:"bytes,2,opt,name=encryption" json:"encryption,omitempty"`
	Ingress              *Ingress      `protobuf:"bytes,3,opt,name=ingress" json:"ingress,omitempty"`
	Routes               []*Route      `protobuf:"bytes,4,rep,name=routes" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MeshConfig) Reset()         { *m = MeshConfig{} }
func (m *MeshConfig) String() string { return proto.CompactTextString(m) }
func (*MeshConfig) ProtoMessage()    {}
func (*MeshConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{0}
}
func (m *MeshConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshConfig.Unmarshal(m, b)
}
func (m *MeshConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshConfig.Marshal(b, m, deterministic)
}
func (dst *MeshConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshConfig.Merge(dst, src)
}
func (m *MeshConfig) XXX_Size() int {
	return xxx_messageInfo_MeshConfig.Size(m)
}
func (m *MeshConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MeshConfig proto.InternalMessageInfo

func (m *MeshConfig) GetMetadata() *MeshMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MeshConfig) GetEncryption() *Encryption {
	if m != nil {
		return m.Encryption
	}
	return nil
}

func (m *MeshConfig) GetIngress() *Ingress {
	if m != nil {
		return m.Ingress
	}
	return nil
}

func (m *MeshConfig) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

type MeshMetadata struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MeshType             MeshType `protobuf:"varint,2,opt,name=meshType,proto3,enum=supergloo.solo.io.MeshType" json:"meshType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeshMetadata) Reset()         { *m = MeshMetadata{} }
func (m *MeshMetadata) String() string { return proto.CompactTextString(m) }
func (*MeshMetadata) ProtoMessage()    {}
func (*MeshMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{1}
}
func (m *MeshMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeshMetadata.Unmarshal(m, b)
}
func (m *MeshMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeshMetadata.Marshal(b, m, deterministic)
}
func (dst *MeshMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeshMetadata.Merge(dst, src)
}
func (m *MeshMetadata) XXX_Size() int {
	return xxx_messageInfo_MeshMetadata.Size(m)
}
func (m *MeshMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_MeshMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_MeshMetadata proto.InternalMessageInfo

func (m *MeshMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MeshMetadata) GetMeshType() MeshType {
	if m != nil {
		return m.MeshType
	}
	return MeshType_ISTIO
}

// Defines mesh-level configuration for encryption. Supports communication within a mesh and through ingress.
// Communicating between TLS and non-TLS enabled services is not supported. TODO: Is this ok?
// TODO: What do we need to support communication across mesh?
type Encryption struct {
	// If set to true, TLS is enabled across the entire mesh.
	TlsEnabled bool `protobuf:"varint,1,opt,name=tlsEnabled,proto3" json:"tlsEnabled,omitempty"`
	// If TLS is enabled, this is the name of the secret containing the certs.
	// When using Istio, this should either be "istio.default", meaning Istio is using the default Citadel cert
	// generation, or "cacert", which is a custom-uploaded Kubernetes secret containing all the cert files.
	// When using Linkerd, this is the name of a secret that will be mounted into the linkerd Kubernetes DaemonSet.
	CertSecret           string   `protobuf:"bytes,2,opt,name=certSecret,proto3" json:"certSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Encryption) Reset()         { *m = Encryption{} }
func (m *Encryption) String() string { return proto.CompactTextString(m) }
func (*Encryption) ProtoMessage()    {}
func (*Encryption) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{2}
}
func (m *Encryption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Encryption.Unmarshal(m, b)
}
func (m *Encryption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Encryption.Marshal(b, m, deterministic)
}
func (dst *Encryption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Encryption.Merge(dst, src)
}
func (m *Encryption) XXX_Size() int {
	return xxx_messageInfo_Encryption.Size(m)
}
func (m *Encryption) XXX_DiscardUnknown() {
	xxx_messageInfo_Encryption.DiscardUnknown(m)
}

var xxx_messageInfo_Encryption proto.InternalMessageInfo

func (m *Encryption) GetTlsEnabled() bool {
	if m != nil {
		return m.TlsEnabled
	}
	return false
}

func (m *Encryption) GetCertSecret() string {
	if m != nil {
		return m.CertSecret
	}
	return ""
}

// Defines Ingress for the service mesh. The Ingress port is 80 for HTTP with no TLS, and 443 for HTTP with TLS.
type Ingress struct {
	// The name of the gateway. In Istio, a Gateway is configured with this name, and a VirtualService references
	// this gateway name. In Linkerd1, an Ingress is configured with this name.
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host                 []*IngressHost `protobuf:"bytes,2,rep,name=host" json:"host,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Ingress) Reset()         { *m = Ingress{} }
func (m *Ingress) String() string { return proto.CompactTextString(m) }
func (*Ingress) ProtoMessage()    {}
func (*Ingress) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{3}
}
func (m *Ingress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ingress.Unmarshal(m, b)
}
func (m *Ingress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ingress.Marshal(b, m, deterministic)
}
func (dst *Ingress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ingress.Merge(dst, src)
}
func (m *Ingress) XXX_Size() int {
	return xxx_messageInfo_Ingress.Size(m)
}
func (m *Ingress) XXX_DiscardUnknown() {
	xxx_messageInfo_Ingress.DiscardUnknown(m)
}

var xxx_messageInfo_Ingress proto.InternalMessageInfo

func (m *Ingress) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ingress) GetHost() []*IngressHost {
	if m != nil {
		return m.Host
	}
	return nil
}

type IngressHost struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Paths                []*Path  `protobuf:"bytes,2,rep,name=paths" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IngressHost) Reset()         { *m = IngressHost{} }
func (m *IngressHost) String() string { return proto.CompactTextString(m) }
func (*IngressHost) ProtoMessage()    {}
func (*IngressHost) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{4}
}
func (m *IngressHost) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IngressHost.Unmarshal(m, b)
}
func (m *IngressHost) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IngressHost.Marshal(b, m, deterministic)
}
func (dst *IngressHost) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IngressHost.Merge(dst, src)
}
func (m *IngressHost) XXX_Size() int {
	return xxx_messageInfo_IngressHost.Size(m)
}
func (m *IngressHost) XXX_DiscardUnknown() {
	xxx_messageInfo_IngressHost.DiscardUnknown(m)
}

var xxx_messageInfo_IngressHost proto.InternalMessageInfo

func (m *IngressHost) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *IngressHost) GetPaths() []*Path {
	if m != nil {
		return m.Paths
	}
	return nil
}

type Path struct {
	PathPrefix           string       `protobuf:"bytes,1,opt,name=pathPrefix,proto3" json:"pathPrefix,omitempty"`
	Destination          *Destination `protobuf:"bytes,2,opt,name=destination" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{5}
}
func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (dst *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(dst, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *Path) GetDestination() *Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

type Destination struct {
	// In Itsio, this is the name of the Kubernetes VirtualService. In Linkerd1, this is the name of the Kubernetes Service.
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Port                 int32             `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Labels               map[string]string `protobuf:"bytes,3,rep,name=labels" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Destination) Reset()         { *m = Destination{} }
func (m *Destination) String() string { return proto.CompactTextString(m) }
func (*Destination) ProtoMessage()    {}
func (*Destination) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{6}
}
func (m *Destination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Destination.Unmarshal(m, b)
}
func (m *Destination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Destination.Marshal(b, m, deterministic)
}
func (dst *Destination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Destination.Merge(dst, src)
}
func (m *Destination) XXX_Size() int {
	return xxx_messageInfo_Destination.Size(m)
}
func (m *Destination) XXX_DiscardUnknown() {
	xxx_messageInfo_Destination.DiscardUnknown(m)
}

var xxx_messageInfo_Destination proto.InternalMessageInfo

func (m *Destination) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Destination) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Destination) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// This is a single route, of which there may be many for a VirtualService or Service. This is a nice abstraction because
// it is easier to reason about how to apply a single route on different meshes. It's not a nice abstraction because it is
// a much more limited set of route configuration options than what itsio supports, and it may not be obvious how to apply
// all the routes for a particular name. Just a straw man to start discussing routing use cases and glooshot.
type Route struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Host                 string         `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	PathPrefix           string         `protobuf:"bytes,3,opt,name=pathPrefix,proto3" json:"pathPrefix,omitempty"`
	Rewrite              string         `protobuf:"bytes,4,opt,name=rewrite,proto3" json:"rewrite,omitempty"`
	Destination          *Destination   `protobuf:"bytes,5,opt,name=destination" json:"destination,omitempty"`
	RoutePlugins         []*RoutePlugin `protobuf:"bytes,6,rep,name=routePlugins" json:"routePlugins,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{7}
}
func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (dst *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(dst, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Route) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Route) GetPathPrefix() string {
	if m != nil {
		return m.PathPrefix
	}
	return ""
}

func (m *Route) GetRewrite() string {
	if m != nil {
		return m.Rewrite
	}
	return ""
}

func (m *Route) GetDestination() *Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (m *Route) GetRoutePlugins() []*RoutePlugin {
	if m != nil {
		return m.RoutePlugins
	}
	return nil
}

type RoutePlugin struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoutePlugin) Reset()         { *m = RoutePlugin{} }
func (m *RoutePlugin) String() string { return proto.CompactTextString(m) }
func (*RoutePlugin) ProtoMessage()    {}
func (*RoutePlugin) Descriptor() ([]byte, []int) {
	return fileDescriptor_mesh_0d0ef610050ba3ff, []int{8}
}
func (m *RoutePlugin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoutePlugin.Unmarshal(m, b)
}
func (m *RoutePlugin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoutePlugin.Marshal(b, m, deterministic)
}
func (dst *RoutePlugin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoutePlugin.Merge(dst, src)
}
func (m *RoutePlugin) XXX_Size() int {
	return xxx_messageInfo_RoutePlugin.Size(m)
}
func (m *RoutePlugin) XXX_DiscardUnknown() {
	xxx_messageInfo_RoutePlugin.DiscardUnknown(m)
}

var xxx_messageInfo_RoutePlugin proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MeshConfig)(nil), "supergloo.solo.io.MeshConfig")
	proto.RegisterType((*MeshMetadata)(nil), "supergloo.solo.io.MeshMetadata")
	proto.RegisterType((*Encryption)(nil), "supergloo.solo.io.Encryption")
	proto.RegisterType((*Ingress)(nil), "supergloo.solo.io.Ingress")
	proto.RegisterType((*IngressHost)(nil), "supergloo.solo.io.IngressHost")
	proto.RegisterType((*Path)(nil), "supergloo.solo.io.Path")
	proto.RegisterType((*Destination)(nil), "supergloo.solo.io.Destination")
	proto.RegisterMapType((map[string]string)(nil), "supergloo.solo.io.Destination.LabelsEntry")
	proto.RegisterType((*Route)(nil), "supergloo.solo.io.Route")
	proto.RegisterType((*RoutePlugin)(nil), "supergloo.solo.io.RoutePlugin")
	proto.RegisterEnum("supergloo.solo.io.MeshType", MeshType_name, MeshType_value)
}
func (this *MeshConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshConfig)
	if !ok {
		that2, ok := that.(MeshConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Metadata.Equal(that1.Metadata) {
		return false
	}
	if !this.Encryption.Equal(that1.Encryption) {
		return false
	}
	if !this.Ingress.Equal(that1.Ingress) {
		return false
	}
	if len(this.Routes) != len(that1.Routes) {
		return false
	}
	for i := range this.Routes {
		if !this.Routes[i].Equal(that1.Routes[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MeshMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MeshMetadata)
	if !ok {
		that2, ok := that.(MeshMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.MeshType != that1.MeshType {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Encryption) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Encryption)
	if !ok {
		that2, ok := that.(Encryption)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TlsEnabled != that1.TlsEnabled {
		return false
	}
	if this.CertSecret != that1.CertSecret {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Ingress) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Ingress)
	if !ok {
		that2, ok := that.(Ingress)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if len(this.Host) != len(that1.Host) {
		return false
	}
	for i := range this.Host {
		if !this.Host[i].Equal(that1.Host[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *IngressHost) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IngressHost)
	if !ok {
		that2, ok := that.(IngressHost)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Host != that1.Host {
		return false
	}
	if len(this.Paths) != len(that1.Paths) {
		return false
	}
	for i := range this.Paths {
		if !this.Paths[i].Equal(that1.Paths[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Path) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Path)
	if !ok {
		that2, ok := that.(Path)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.PathPrefix != that1.PathPrefix {
		return false
	}
	if !this.Destination.Equal(that1.Destination) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Destination) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Destination)
	if !ok {
		that2, ok := that.(Destination)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Port != that1.Port {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Route) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Host != that1.Host {
		return false
	}
	if this.PathPrefix != that1.PathPrefix {
		return false
	}
	if this.Rewrite != that1.Rewrite {
		return false
	}
	if !this.Destination.Equal(that1.Destination) {
		return false
	}
	if len(this.RoutePlugins) != len(that1.RoutePlugins) {
		return false
	}
	for i := range this.RoutePlugins {
		if !this.RoutePlugins[i].Equal(that1.RoutePlugins[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *RoutePlugin) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RoutePlugin)
	if !ok {
		that2, ok := that.(RoutePlugin)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("mesh.proto", fileDescriptor_mesh_0d0ef610050ba3ff) }

var fileDescriptor_mesh_0d0ef610050ba3ff = []byte{
	// 573 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0xfd, 0xdc, 0xd8, 0xf9, 0x19, 0xe7, 0x43, 0x61, 0x55, 0x09, 0x2b, 0x88, 0x50, 0x99, 0x0b,
	0xaa, 0x4a, 0xb1, 0x69, 0x40, 0xe2, 0x4f, 0x48, 0x28, 0x34, 0x12, 0x11, 0x29, 0x84, 0x6d, 0xaf,
	0xe0, 0xca, 0x49, 0xb6, 0xf6, 0xaa, 0x8e, 0xd7, 0xda, 0xdd, 0x14, 0xf2, 0x46, 0x3c, 0x00, 0x4f,
	0xc4, 0x0b, 0xc0, 0x23, 0xa0, 0xdd, 0x38, 0xc9, 0x96, 0x3a, 0x48, 0xdc, 0xcd, 0xce, 0x9c, 0x33,
	0xb3, 0xe7, 0xec, 0xd8, 0x00, 0x73, 0x22, 0x92, 0x20, 0xe7, 0x4c, 0x32, 0x74, 0x5b, 0x2c, 0x72,
	0xc2, 0xe3, 0x94, 0xb1, 0x40, 0xb0, 0x94, 0x05, 0x94, 0xb5, 0xf7, 0x63, 0x16, 0x33, 0x5d, 0x0d,
	0x55, 0xb4, 0x02, 0xfa, 0xbf, 0x2c, 0x80, 0x53, 0x22, 0x92, 0x37, 0x2c, 0xbb, 0xa0, 0x31, 0x7a,
	0x09, 0xf5, 0x39, 0x91, 0xd1, 0x2c, 0x92, 0x91, 0x67, 0x1d, 0x58, 0x87, 0x6e, 0xef, 0x7e, 0x70,
	0xa3, 0x55, 0xa0, 0x08, 0xa7, 0x05, 0x0c, 0x6f, 0x08, 0xe8, 0x15, 0x00, 0xc9, 0xa6, 0x7c, 0x99,
	0x4b, 0xca, 0x32, 0x6f, 0x4f, 0xd3, 0xef, 0x95, 0xd0, 0x07, 0x1b, 0x10, 0x36, 0x08, 0xe8, 0x09,
	0xd4, 0x68, 0x16, 0x73, 0x22, 0x84, 0x57, 0xd1, 0xdc, 0x76, 0x09, 0x77, 0xb8, 0x42, 0xe0, 0x35,
	0x14, 0x3d, 0x82, 0x2a, 0x67, 0x0b, 0x49, 0x84, 0x67, 0x1f, 0x54, 0x0e, 0xdd, 0x9e, 0x57, 0x42,
	0xc2, 0x0a, 0x80, 0x0b, 0x9c, 0xff, 0x19, 0x9a, 0xa6, 0x00, 0x84, 0xc0, 0xce, 0xa2, 0x39, 0xd1,
	0x7a, 0x1b, 0x58, 0xc7, 0xe8, 0xa9, 0xf2, 0x41, 0x24, 0xe7, 0xcb, 0x9c, 0x68, 0x21, 0xb7, 0x7a,
	0x77, 0x77, 0xf8, 0xa0, 0x20, 0x78, 0x03, 0xf6, 0x47, 0x00, 0x5b, 0x79, 0xa8, 0x03, 0x20, 0x53,
	0x31, 0xc8, 0xa2, 0x49, 0x4a, 0x66, 0x7a, 0x40, 0x1d, 0x1b, 0x19, 0x55, 0x9f, 0x12, 0x2e, 0xcf,
	0xc8, 0x94, 0x13, 0xa9, 0x07, 0x35, 0xb0, 0x91, 0xf1, 0x3f, 0x42, 0xad, 0x10, 0x5c, 0x7a, 0xcb,
	0x1e, 0xd8, 0x09, 0x13, 0x8a, 0xa8, 0x94, 0x77, 0x76, 0xdb, 0xf5, 0x96, 0x09, 0x89, 0x35, 0xd6,
	0x1f, 0x83, 0x6b, 0x24, 0x55, 0x5b, 0xdd, 0xa2, 0x68, 0xab, 0x62, 0xd4, 0x05, 0x27, 0x8f, 0x64,
	0x22, 0x8a, 0xbe, 0x77, 0x4a, 0xfa, 0x8e, 0x23, 0x99, 0xe0, 0x15, 0xca, 0x4f, 0xc0, 0x56, 0x47,
	0x25, 0x46, 0x25, 0xc6, 0x9c, 0x5c, 0xd0, 0xaf, 0x45, 0x43, 0x23, 0x83, 0x5e, 0x83, 0x3b, 0x23,
	0x42, 0xd2, 0x2c, 0x32, 0xf6, 0xa3, 0xec, 0xd2, 0x27, 0x5b, 0x14, 0x36, 0x29, 0xfe, 0x77, 0x0b,
	0x5c, 0xa3, 0x58, 0xea, 0x09, 0x02, 0x3b, 0x67, 0x7c, 0x65, 0xa6, 0x83, 0x75, 0x8c, 0xfa, 0x50,
	0x4d, 0xa3, 0x09, 0x49, 0xd5, 0x62, 0x29, 0x45, 0x47, 0x7f, 0x1f, 0x1a, 0x8c, 0x34, 0x78, 0x90,
	0x49, 0xbe, 0xc4, 0x05, 0xb3, 0xfd, 0x1c, 0x5c, 0x23, 0x8d, 0x5a, 0x50, 0xb9, 0x24, 0xcb, 0x62,
	0xb2, 0x0a, 0xd1, 0x3e, 0x38, 0x57, 0x51, 0xba, 0x20, 0xc5, 0x33, 0xae, 0x0e, 0x2f, 0xf6, 0x9e,
	0x59, 0xfe, 0x4f, 0x0b, 0x1c, 0xbd, 0x82, 0xbb, 0x2e, 0x5c, 0x3c, 0xe2, 0xf6, 0x05, 0xae, 0x5b,
	0x59, 0xb9, 0x61, 0xa5, 0x07, 0x35, 0x4e, 0xbe, 0x70, 0x2a, 0x89, 0x67, 0xeb, 0xe2, 0xfa, 0xf8,
	0xa7, 0xc9, 0xce, 0x3f, 0x9b, 0x8c, 0xfa, 0xd0, 0xd4, 0x1f, 0xca, 0x38, 0x5d, 0xc4, 0x34, 0x13,
	0x5e, 0x75, 0xe7, 0x72, 0xe1, 0x2d, 0x0c, 0x5f, 0xe3, 0xf8, 0xff, 0x83, 0x6b, 0x14, 0x8f, 0x1e,
	0x40, 0x7d, 0xfd, 0xa9, 0xa0, 0x06, 0x38, 0xc3, 0xb3, 0xf3, 0xe1, 0x87, 0xd6, 0x7f, 0xa8, 0x09,
	0xf5, 0xd1, 0xf0, 0xfd, 0xbb, 0x01, 0x3e, 0x39, 0x6e, 0x59, 0xfd, 0xee, 0xb7, 0x1f, 0x1d, 0xeb,
	0xd3, 0xc3, 0x98, 0xca, 0x64, 0x31, 0x09, 0xa6, 0x6c, 0x1e, 0xaa, 0x39, 0x5d, 0xca, 0xc2, 0xcd,
	0xe4, 0x30, 0xbf, 0x8c, 0xc3, 0x28, 0xa7, 0xe1, 0xd5, 0xf1, 0xa4, 0xaa, 0xff, 0x5f, 0x8f, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x51, 0xeb, 0xb7, 0xf6, 0x04, 0x00, 0x00,
}