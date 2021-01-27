// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/cloud/dnsmgr/dnsmgrpb/service.proto

package dnsmgr

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	uuidpb "pixielabs.ai/pixielabs/src/api/public/uuidpb"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetDNSAddressRequest struct {
	ClusterID *uuidpb.UUID `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	IPAddress string       `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (m *GetDNSAddressRequest) Reset()      { *m = GetDNSAddressRequest{} }
func (*GetDNSAddressRequest) ProtoMessage() {}
func (*GetDNSAddressRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4b129026db367e, []int{0}
}
func (m *GetDNSAddressRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDNSAddressRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDNSAddressRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDNSAddressRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDNSAddressRequest.Merge(m, src)
}
func (m *GetDNSAddressRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetDNSAddressRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDNSAddressRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDNSAddressRequest proto.InternalMessageInfo

func (m *GetDNSAddressRequest) GetClusterID() *uuidpb.UUID {
	if m != nil {
		return m.ClusterID
	}
	return nil
}

func (m *GetDNSAddressRequest) GetIPAddress() string {
	if m != nil {
		return m.IPAddress
	}
	return ""
}

type GetDNSAddressResponse struct {
	DNSAddress string `protobuf:"bytes,1,opt,name=dns_address,json=dnsAddress,proto3" json:"dns_address,omitempty"`
}

func (m *GetDNSAddressResponse) Reset()      { *m = GetDNSAddressResponse{} }
func (*GetDNSAddressResponse) ProtoMessage() {}
func (*GetDNSAddressResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4b129026db367e, []int{1}
}
func (m *GetDNSAddressResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetDNSAddressResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetDNSAddressResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetDNSAddressResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDNSAddressResponse.Merge(m, src)
}
func (m *GetDNSAddressResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetDNSAddressResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDNSAddressResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDNSAddressResponse proto.InternalMessageInfo

func (m *GetDNSAddressResponse) GetDNSAddress() string {
	if m != nil {
		return m.DNSAddress
	}
	return ""
}

type GetSSLCertsRequest struct {
	ClusterID *uuidpb.UUID `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
}

func (m *GetSSLCertsRequest) Reset()      { *m = GetSSLCertsRequest{} }
func (*GetSSLCertsRequest) ProtoMessage() {}
func (*GetSSLCertsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4b129026db367e, []int{2}
}
func (m *GetSSLCertsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSSLCertsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSSLCertsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSSLCertsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSSLCertsRequest.Merge(m, src)
}
func (m *GetSSLCertsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetSSLCertsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSSLCertsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSSLCertsRequest proto.InternalMessageInfo

func (m *GetSSLCertsRequest) GetClusterID() *uuidpb.UUID {
	if m != nil {
		return m.ClusterID
	}
	return nil
}

type GetSSLCertsResponse struct {
	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Cert string `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (m *GetSSLCertsResponse) Reset()      { *m = GetSSLCertsResponse{} }
func (*GetSSLCertsResponse) ProtoMessage() {}
func (*GetSSLCertsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c4b129026db367e, []int{3}
}
func (m *GetSSLCertsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetSSLCertsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetSSLCertsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetSSLCertsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSSLCertsResponse.Merge(m, src)
}
func (m *GetSSLCertsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetSSLCertsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSSLCertsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetSSLCertsResponse proto.InternalMessageInfo

func (m *GetSSLCertsResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *GetSSLCertsResponse) GetCert() string {
	if m != nil {
		return m.Cert
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDNSAddressRequest)(nil), "pl.services.GetDNSAddressRequest")
	proto.RegisterType((*GetDNSAddressResponse)(nil), "pl.services.GetDNSAddressResponse")
	proto.RegisterType((*GetSSLCertsRequest)(nil), "pl.services.GetSSLCertsRequest")
	proto.RegisterType((*GetSSLCertsResponse)(nil), "pl.services.GetSSLCertsResponse")
}

func init() {
	proto.RegisterFile("src/cloud/dnsmgr/dnsmgrpb/service.proto", fileDescriptor_9c4b129026db367e)
}

var fileDescriptor_9c4b129026db367e = []byte{
	// 443 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x6f, 0x13, 0x41,
	0x10, 0xbd, 0x05, 0x84, 0xe4, 0xb1, 0x0c, 0x68, 0x01, 0x29, 0x72, 0xb1, 0x36, 0xd7, 0x90, 0x02,
	0x76, 0xa5, 0x40, 0x17, 0x51, 0x90, 0x58, 0x0a, 0x96, 0x20, 0x42, 0x77, 0x8a, 0x84, 0x68, 0x22,
	0xdf, 0xed, 0x72, 0x2c, 0x5c, 0x7c, 0xcb, 0x7e, 0x20, 0xe8, 0x90, 0xf8, 0x03, 0xfc, 0x0c, 0xfe,
	0x04, 0x3d, 0xa5, 0xcb, 0x54, 0x11, 0x5e, 0x37, 0x94, 0xf9, 0x09, 0x28, 0xbb, 0xe7, 0xc8, 0x31,
	0x1f, 0x15, 0xd5, 0xbe, 0x77, 0x7a, 0xf3, 0xe6, 0xcd, 0xcd, 0xc0, 0x5d, 0xa3, 0x4b, 0x56, 0xd6,
	0x8d, 0xe3, 0x8c, 0x4f, 0xcd, 0x51, 0xa5, 0xdb, 0x47, 0x15, 0xcc, 0x08, 0xfd, 0x5e, 0x96, 0x82,
	0x2a, 0xdd, 0xd8, 0x06, 0x77, 0x55, 0x4d, 0xdb, 0x2f, 0xa6, 0x7f, 0xbf, 0x92, 0xf6, 0xb5, 0x2b,
	0x68, 0xd9, 0x1c, 0xb1, 0xaa, 0xa9, 0x1a, 0x16, 0x34, 0x85, 0x7b, 0x15, 0x58, 0x20, 0x01, 0xc5,
	0xda, 0xfe, 0xf0, 0xac, 0xc9, 0x44, 0x49, 0xa6, 0x5c, 0x51, 0xcb, 0x92, 0x39, 0x27, 0xb9, 0x2a,
	0xc2, 0x13, 0x15, 0xe9, 0x67, 0x04, 0xb7, 0xf6, 0x84, 0x1d, 0xed, 0xe7, 0x8f, 0x39, 0xd7, 0xc2,
	0x98, 0x4c, 0xbc, 0x73, 0xc2, 0x58, 0xfc, 0x08, 0xa0, 0xac, 0x9d, 0xb1, 0x42, 0x1f, 0x4a, 0xbe,
	0x81, 0x86, 0x68, 0xb3, 0xbb, 0x75, 0x9d, 0xaa, 0x9a, 0x46, 0x0f, 0x7a, 0x70, 0x30, 0x1e, 0xed,
	0xf4, 0xfc, 0xc9, 0xa0, 0xb3, 0x1b, 0x65, 0xe3, 0x51, 0xd6, 0x69, 0x2b, 0xc6, 0x1c, 0xdf, 0x03,
	0x90, 0xea, 0x70, 0x12, 0x3d, 0x37, 0x2e, 0x0d, 0xd1, 0x66, 0x27, 0xaa, 0xc7, 0xcf, 0x97, 0x8d,
	0x3a, 0x52, 0xb5, 0x30, 0x7d, 0x02, 0xb7, 0xd7, 0x42, 0x18, 0xd5, 0x4c, 0x8d, 0xc0, 0x0c, 0xba,
	0x7c, 0x6a, 0xce, 0x7d, 0x50, 0xf0, 0xb9, 0xe6, 0x4f, 0x06, 0xb0, 0x22, 0x06, 0x3e, 0x35, 0x4b,
	0xa7, 0x1c, 0xf0, 0x9e, 0xb0, 0x79, 0xfe, 0x74, 0x57, 0x68, 0xfb, 0x9f, 0x86, 0x49, 0xb7, 0xe1,
	0xe6, 0x05, 0xd3, 0x36, 0xdc, 0x0d, 0xb8, 0xfc, 0x56, 0x7c, 0x8c, 0xa1, 0xb2, 0x33, 0x88, 0x31,
	0x5c, 0x29, 0x85, 0xb6, 0x71, 0xde, 0x2c, 0xe0, 0xad, 0x6f, 0x08, 0x7a, 0xa3, 0xfd, 0xfc, 0x59,
	0xa5, 0xf3, 0xb8, 0x45, 0xfc, 0x02, 0x7a, 0x17, 0xa6, 0xc5, 0x77, 0xe8, 0xca, 0x8e, 0xe9, 0x9f,
	0xd6, 0xd1, 0x4f, 0xff, 0x25, 0x89, 0x79, 0xd2, 0x04, 0x67, 0xd0, 0x5d, 0x09, 0x8a, 0x07, 0xeb,
	0x45, 0x6b, 0xff, 0xa5, 0x3f, 0xfc, 0xbb, 0x60, 0xe9, 0xb9, 0xf3, 0x66, 0x36, 0x27, 0xc9, 0xf1,
	0x9c, 0x24, 0xa7, 0x73, 0x82, 0x3e, 0x79, 0x82, 0xbe, 0x7a, 0x82, 0xbe, 0x7b, 0x82, 0x66, 0x9e,
	0xa0, 0x1f, 0x9e, 0xa0, 0x9f, 0x9e, 0x24, 0xa7, 0x9e, 0xa0, 0x2f, 0x0b, 0x92, 0xcc, 0x16, 0x24,
	0x39, 0x5e, 0x90, 0xe4, 0xe5, 0x43, 0x25, 0x3f, 0x48, 0x51, 0x4f, 0x0a, 0x43, 0x27, 0x92, 0x9d,
	0x13, 0xf6, 0xdb, 0xd5, 0x87, 0x33, 0xdc, 0x8e, 0xa4, 0xb8, 0x1a, 0xd8, 0x83, 0x5f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf8, 0x7c, 0x47, 0x77, 0x1d, 0x03, 0x00, 0x00,
}

func (this *GetDNSAddressRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDNSAddressRequest)
	if !ok {
		that2, ok := that.(GetDNSAddressRequest)
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
	if !this.ClusterID.Equal(that1.ClusterID) {
		return false
	}
	if this.IPAddress != that1.IPAddress {
		return false
	}
	return true
}
func (this *GetDNSAddressResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetDNSAddressResponse)
	if !ok {
		that2, ok := that.(GetDNSAddressResponse)
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
	if this.DNSAddress != that1.DNSAddress {
		return false
	}
	return true
}
func (this *GetSSLCertsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSSLCertsRequest)
	if !ok {
		that2, ok := that.(GetSSLCertsRequest)
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
	if !this.ClusterID.Equal(that1.ClusterID) {
		return false
	}
	return true
}
func (this *GetSSLCertsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetSSLCertsResponse)
	if !ok {
		that2, ok := that.(GetSSLCertsResponse)
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
	if this.Key != that1.Key {
		return false
	}
	if this.Cert != that1.Cert {
		return false
	}
	return true
}
func (this *GetDNSAddressRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dnsmgr.GetDNSAddressRequest{")
	if this.ClusterID != nil {
		s = append(s, "ClusterID: "+fmt.Sprintf("%#v", this.ClusterID)+",\n")
	}
	s = append(s, "IPAddress: "+fmt.Sprintf("%#v", this.IPAddress)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetDNSAddressResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dnsmgr.GetDNSAddressResponse{")
	s = append(s, "DNSAddress: "+fmt.Sprintf("%#v", this.DNSAddress)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSSLCertsRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&dnsmgr.GetSSLCertsRequest{")
	if this.ClusterID != nil {
		s = append(s, "ClusterID: "+fmt.Sprintf("%#v", this.ClusterID)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *GetSSLCertsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&dnsmgr.GetSSLCertsResponse{")
	s = append(s, "Key: "+fmt.Sprintf("%#v", this.Key)+",\n")
	s = append(s, "Cert: "+fmt.Sprintf("%#v", this.Cert)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringService(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DNSMgrServiceClient is the client API for DNSMgrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DNSMgrServiceClient interface {
	GetDNSAddress(ctx context.Context, in *GetDNSAddressRequest, opts ...grpc.CallOption) (*GetDNSAddressResponse, error)
	GetSSLCerts(ctx context.Context, in *GetSSLCertsRequest, opts ...grpc.CallOption) (*GetSSLCertsResponse, error)
}

type dNSMgrServiceClient struct {
	cc *grpc.ClientConn
}

func NewDNSMgrServiceClient(cc *grpc.ClientConn) DNSMgrServiceClient {
	return &dNSMgrServiceClient{cc}
}

func (c *dNSMgrServiceClient) GetDNSAddress(ctx context.Context, in *GetDNSAddressRequest, opts ...grpc.CallOption) (*GetDNSAddressResponse, error) {
	out := new(GetDNSAddressResponse)
	err := c.cc.Invoke(ctx, "/pl.services.DNSMgrService/GetDNSAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dNSMgrServiceClient) GetSSLCerts(ctx context.Context, in *GetSSLCertsRequest, opts ...grpc.CallOption) (*GetSSLCertsResponse, error) {
	out := new(GetSSLCertsResponse)
	err := c.cc.Invoke(ctx, "/pl.services.DNSMgrService/GetSSLCerts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DNSMgrServiceServer is the server API for DNSMgrService service.
type DNSMgrServiceServer interface {
	GetDNSAddress(context.Context, *GetDNSAddressRequest) (*GetDNSAddressResponse, error)
	GetSSLCerts(context.Context, *GetSSLCertsRequest) (*GetSSLCertsResponse, error)
}

// UnimplementedDNSMgrServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDNSMgrServiceServer struct {
}

func (*UnimplementedDNSMgrServiceServer) GetDNSAddress(ctx context.Context, req *GetDNSAddressRequest) (*GetDNSAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDNSAddress not implemented")
}
func (*UnimplementedDNSMgrServiceServer) GetSSLCerts(ctx context.Context, req *GetSSLCertsRequest) (*GetSSLCertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSSLCerts not implemented")
}

func RegisterDNSMgrServiceServer(s *grpc.Server, srv DNSMgrServiceServer) {
	s.RegisterService(&_DNSMgrService_serviceDesc, srv)
}

func _DNSMgrService_GetDNSAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDNSAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSMgrServiceServer).GetDNSAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pl.services.DNSMgrService/GetDNSAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSMgrServiceServer).GetDNSAddress(ctx, req.(*GetDNSAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DNSMgrService_GetSSLCerts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSSLCertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DNSMgrServiceServer).GetSSLCerts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pl.services.DNSMgrService/GetSSLCerts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DNSMgrServiceServer).GetSSLCerts(ctx, req.(*GetSSLCertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DNSMgrService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pl.services.DNSMgrService",
	HandlerType: (*DNSMgrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDNSAddress",
			Handler:    _DNSMgrService_GetDNSAddress_Handler,
		},
		{
			MethodName: "GetSSLCerts",
			Handler:    _DNSMgrService_GetSSLCerts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/cloud/dnsmgr/dnsmgrpb/service.proto",
}

func (m *GetDNSAddressRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDNSAddressRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDNSAddressRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IPAddress) > 0 {
		i -= len(m.IPAddress)
		copy(dAtA[i:], m.IPAddress)
		i = encodeVarintService(dAtA, i, uint64(len(m.IPAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.ClusterID != nil {
		{
			size, err := m.ClusterID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetDNSAddressResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetDNSAddressResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetDNSAddressResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DNSAddress) > 0 {
		i -= len(m.DNSAddress)
		copy(dAtA[i:], m.DNSAddress)
		i = encodeVarintService(dAtA, i, uint64(len(m.DNSAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetSSLCertsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSSLCertsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSSLCertsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClusterID != nil {
		{
			size, err := m.ClusterID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetSSLCertsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetSSLCertsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetSSLCertsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Cert) > 0 {
		i -= len(m.Cert)
		copy(dAtA[i:], m.Cert)
		i = encodeVarintService(dAtA, i, uint64(len(m.Cert)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintService(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetDNSAddressRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterID != nil {
		l = m.ClusterID.Size()
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.IPAddress)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetDNSAddressResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DNSAddress)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetSSLCertsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterID != nil {
		l = m.ClusterID.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetSSLCertsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.Cert)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetDNSAddressRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDNSAddressRequest{`,
		`ClusterID:` + strings.Replace(fmt.Sprintf("%v", this.ClusterID), "UUID", "uuidpb.UUID", 1) + `,`,
		`IPAddress:` + fmt.Sprintf("%v", this.IPAddress) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetDNSAddressResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetDNSAddressResponse{`,
		`DNSAddress:` + fmt.Sprintf("%v", this.DNSAddress) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSSLCertsRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSSLCertsRequest{`,
		`ClusterID:` + strings.Replace(fmt.Sprintf("%v", this.ClusterID), "UUID", "uuidpb.UUID", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *GetSSLCertsResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetSSLCertsResponse{`,
		`Key:` + fmt.Sprintf("%v", this.Key) + `,`,
		`Cert:` + fmt.Sprintf("%v", this.Cert) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetDNSAddressRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDNSAddressRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDNSAddressRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterID == nil {
				m.ClusterID = &uuidpb.UUID{}
			}
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetDNSAddressResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetDNSAddressResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetDNSAddressResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DNSAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DNSAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetSSLCertsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetSSLCertsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSSLCertsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterID == nil {
				m.ClusterID = &uuidpb.UUID{}
			}
			if err := m.ClusterID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GetSSLCertsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetSSLCertsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetSSLCertsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
