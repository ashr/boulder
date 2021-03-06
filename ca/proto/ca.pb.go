// Code generated by protoc-gen-go.
// source: ca/proto/ca.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	ca/proto/ca.proto

It has these top-level messages:
	IssueCertificateRequest
	IssuePrecertificateResponse
	IssueCertificateForPrecertificateRequest
	GenerateOCSPRequest
	OCSPResponse
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/letsencrypt/boulder/core/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type IssueCertificateRequest struct {
	Csr              []byte `protobuf:"bytes,1,opt,name=csr" json:"csr,omitempty"`
	RegistrationID   *int64 `protobuf:"varint,2,opt,name=registrationID" json:"registrationID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IssueCertificateRequest) Reset()                    { *m = IssueCertificateRequest{} }
func (m *IssueCertificateRequest) String() string            { return proto1.CompactTextString(m) }
func (*IssueCertificateRequest) ProtoMessage()               {}
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IssueCertificateRequest) GetCsr() []byte {
	if m != nil {
		return m.Csr
	}
	return nil
}

func (m *IssueCertificateRequest) GetRegistrationID() int64 {
	if m != nil && m.RegistrationID != nil {
		return *m.RegistrationID
	}
	return 0
}

type IssuePrecertificateResponse struct {
	Precert           *core.Precertificate    `protobuf:"bytes,1,opt,name=precert" json:"precert,omitempty"`
	SctFetchingConfig *core.SCTFetchingConfig `protobuf:"bytes,2,opt,name=sctFetchingConfig" json:"sctFetchingConfig,omitempty"`
	XXX_unrecognized  []byte                  `json:"-"`
}

func (m *IssuePrecertificateResponse) Reset()                    { *m = IssuePrecertificateResponse{} }
func (m *IssuePrecertificateResponse) String() string            { return proto1.CompactTextString(m) }
func (*IssuePrecertificateResponse) ProtoMessage()               {}
func (*IssuePrecertificateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IssuePrecertificateResponse) GetPrecert() *core.Precertificate {
	if m != nil {
		return m.Precert
	}
	return nil
}

func (m *IssuePrecertificateResponse) GetSctFetchingConfig() *core.SCTFetchingConfig {
	if m != nil {
		return m.SctFetchingConfig
	}
	return nil
}

type IssueCertificateForPrecertificateRequest struct {
	IssueReq         *IssueCertificateRequest `protobuf:"bytes,1,opt,name=issueReq" json:"issueReq,omitempty"`
	PrecertDER       []byte                   `protobuf:"bytes,2,opt,name=precertDER" json:"precertDER,omitempty"`
	SCTs             [][]byte                 `protobuf:"bytes,3,rep,name=SCTs,json=sCTs" json:"SCTs,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *IssueCertificateForPrecertificateRequest) Reset() {
	*m = IssueCertificateForPrecertificateRequest{}
}
func (m *IssueCertificateForPrecertificateRequest) String() string { return proto1.CompactTextString(m) }
func (*IssueCertificateForPrecertificateRequest) ProtoMessage()    {}
func (*IssueCertificateForPrecertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{2}
}

func (m *IssueCertificateForPrecertificateRequest) GetIssueReq() *IssueCertificateRequest {
	if m != nil {
		return m.IssueReq
	}
	return nil
}

func (m *IssueCertificateForPrecertificateRequest) GetPrecertDER() []byte {
	if m != nil {
		return m.PrecertDER
	}
	return nil
}

func (m *IssueCertificateForPrecertificateRequest) GetSCTs() [][]byte {
	if m != nil {
		return m.SCTs
	}
	return nil
}

type GenerateOCSPRequest struct {
	CertDER          []byte  `protobuf:"bytes,1,opt,name=certDER" json:"certDER,omitempty"`
	Status           *string `protobuf:"bytes,2,opt,name=status" json:"status,omitempty"`
	Reason           *int32  `protobuf:"varint,3,opt,name=reason" json:"reason,omitempty"`
	RevokedAt        *int64  `protobuf:"varint,4,opt,name=revokedAt" json:"revokedAt,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GenerateOCSPRequest) Reset()                    { *m = GenerateOCSPRequest{} }
func (m *GenerateOCSPRequest) String() string            { return proto1.CompactTextString(m) }
func (*GenerateOCSPRequest) ProtoMessage()               {}
func (*GenerateOCSPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GenerateOCSPRequest) GetCertDER() []byte {
	if m != nil {
		return m.CertDER
	}
	return nil
}

func (m *GenerateOCSPRequest) GetStatus() string {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return ""
}

func (m *GenerateOCSPRequest) GetReason() int32 {
	if m != nil && m.Reason != nil {
		return *m.Reason
	}
	return 0
}

func (m *GenerateOCSPRequest) GetRevokedAt() int64 {
	if m != nil && m.RevokedAt != nil {
		return *m.RevokedAt
	}
	return 0
}

type OCSPResponse struct {
	Response         []byte `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *OCSPResponse) Reset()                    { *m = OCSPResponse{} }
func (m *OCSPResponse) String() string            { return proto1.CompactTextString(m) }
func (*OCSPResponse) ProtoMessage()               {}
func (*OCSPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OCSPResponse) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto1.RegisterType((*IssueCertificateRequest)(nil), "ca.IssueCertificateRequest")
	proto1.RegisterType((*IssuePrecertificateResponse)(nil), "ca.IssuePrecertificateResponse")
	proto1.RegisterType((*IssueCertificateForPrecertificateRequest)(nil), "ca.IssueCertificateForPrecertificateRequest")
	proto1.RegisterType((*GenerateOCSPRequest)(nil), "ca.GenerateOCSPRequest")
	proto1.RegisterType((*OCSPResponse)(nil), "ca.OCSPResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CertificateAuthority service

type CertificateAuthorityClient interface {
	IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error)
	IssuePrecertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssuePrecertificateResponse, error)
	IssueCertificateForPrecertificate(ctx context.Context, in *IssueCertificateForPrecertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error)
}

type certificateAuthorityClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAuthorityClient(cc *grpc.ClientConn) CertificateAuthorityClient {
	return &certificateAuthorityClient{cc}
}

func (c *certificateAuthorityClient) IssueCertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error) {
	out := new(core.Certificate)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/IssueCertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) IssuePrecertificate(ctx context.Context, in *IssueCertificateRequest, opts ...grpc.CallOption) (*IssuePrecertificateResponse, error) {
	out := new(IssuePrecertificateResponse)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/IssuePrecertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) IssueCertificateForPrecertificate(ctx context.Context, in *IssueCertificateForPrecertificateRequest, opts ...grpc.CallOption) (*core.Certificate, error) {
	out := new(core.Certificate)
	err := grpc.Invoke(ctx, "/ca.CertificateAuthority/IssueCertificateForPrecertificate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CertificateAuthority service

type CertificateAuthorityServer interface {
	IssueCertificate(context.Context, *IssueCertificateRequest) (*core.Certificate, error)
	IssuePrecertificate(context.Context, *IssueCertificateRequest) (*IssuePrecertificateResponse, error)
	IssueCertificateForPrecertificate(context.Context, *IssueCertificateForPrecertificateRequest) (*core.Certificate, error)
}

func RegisterCertificateAuthorityServer(s *grpc.Server, srv CertificateAuthorityServer) {
	s.RegisterService(&_CertificateAuthority_serviceDesc, srv)
}

func _CertificateAuthority_IssueCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssueCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssueCertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_IssuePrecertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssuePrecertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssuePrecertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssuePrecertificate(ctx, req.(*IssueCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_IssueCertificateForPrecertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IssueCertificateForPrecertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).IssueCertificateForPrecertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.CertificateAuthority/IssueCertificateForPrecertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).IssueCertificateForPrecertificate(ctx, req.(*IssueCertificateForPrecertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.CertificateAuthority",
	HandlerType: (*CertificateAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IssueCertificate",
			Handler:    _CertificateAuthority_IssueCertificate_Handler,
		},
		{
			MethodName: "IssuePrecertificate",
			Handler:    _CertificateAuthority_IssuePrecertificate_Handler,
		},
		{
			MethodName: "IssueCertificateForPrecertificate",
			Handler:    _CertificateAuthority_IssueCertificateForPrecertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca/proto/ca.proto",
}

// Client API for OCSPGenerator service

type OCSPGeneratorClient interface {
	GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error)
}

type oCSPGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewOCSPGeneratorClient(cc *grpc.ClientConn) OCSPGeneratorClient {
	return &oCSPGeneratorClient{cc}
}

func (c *oCSPGeneratorClient) GenerateOCSP(ctx context.Context, in *GenerateOCSPRequest, opts ...grpc.CallOption) (*OCSPResponse, error) {
	out := new(OCSPResponse)
	err := grpc.Invoke(ctx, "/ca.OCSPGenerator/GenerateOCSP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OCSPGenerator service

type OCSPGeneratorServer interface {
	GenerateOCSP(context.Context, *GenerateOCSPRequest) (*OCSPResponse, error)
}

func RegisterOCSPGeneratorServer(s *grpc.Server, srv OCSPGeneratorServer) {
	s.RegisterService(&_OCSPGenerator_serviceDesc, srv)
}

func _OCSPGenerator_GenerateOCSP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateOCSPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OCSPGeneratorServer).GenerateOCSP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ca.OCSPGenerator/GenerateOCSP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OCSPGeneratorServer).GenerateOCSP(ctx, req.(*GenerateOCSPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OCSPGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ca.OCSPGenerator",
	HandlerType: (*OCSPGeneratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateOCSP",
			Handler:    _OCSPGenerator_GenerateOCSP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca/proto/ca.proto",
}

func init() { proto1.RegisterFile("ca/proto/ca.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0xe3, 0x94, 0xb4, 0x53, 0x83, 0x92, 0x69, 0x21, 0x56, 0x8a, 0x20, 0xf8, 0x80, 0x2c,
	0x84, 0x5c, 0x29, 0x17, 0x0e, 0x9c, 0x8a, 0xd3, 0xa2, 0x4a, 0x48, 0x54, 0x9b, 0x70, 0xe1, 0xb6,
	0x5a, 0x26, 0xa9, 0x85, 0xe4, 0x4d, 0x77, 0xd7, 0x48, 0x1c, 0xf8, 0x0b, 0x24, 0x0e, 0xfc, 0x2c,
	0xf2, 0x7a, 0x1d, 0x8c, 0x71, 0xdb, 0xdb, 0xce, 0xbc, 0xf1, 0xbc, 0x37, 0xf3, 0xc6, 0x30, 0x16,
	0xfc, 0x74, 0xab, 0xa4, 0x91, 0xa7, 0x82, 0x27, 0xf6, 0x81, 0x7d, 0xc1, 0xa7, 0x8f, 0x85, 0x54,
	0x54, 0x03, 0x52, 0x51, 0x05, 0x45, 0x4b, 0x98, 0x5c, 0x6a, 0x5d, 0x50, 0x4a, 0xca, 0x64, 0xeb,
	0x4c, 0x70, 0x43, 0x8c, 0x6e, 0x0a, 0xd2, 0x06, 0x47, 0xe0, 0x0b, 0xad, 0x42, 0x6f, 0xe6, 0xc5,
	0x01, 0x2b, 0x9f, 0xf8, 0x12, 0x1e, 0x29, 0xda, 0x64, 0xda, 0x28, 0x6e, 0x32, 0x99, 0x5f, 0x2e,
	0xc2, 0xfe, 0xcc, 0x8b, 0x7d, 0xd6, 0xca, 0x46, 0x3f, 0x3d, 0x38, 0xb1, 0x5d, 0xaf, 0x14, 0x89,
	0x66, 0x63, 0xbd, 0x95, 0xb9, 0x26, 0x4c, 0x60, 0xb8, 0xad, 0x10, 0xdb, 0xfd, 0x70, 0x7e, 0x9c,
	0x58, 0x49, 0xad, 0xf2, 0xba, 0x08, 0xcf, 0x61, 0xac, 0x85, 0xb9, 0x20, 0x23, 0xae, 0xb3, 0x7c,
	0x93, 0xca, 0x7c, 0x9d, 0x6d, 0x2c, 0xf5, 0xe1, 0x7c, 0x52, 0x7d, 0xb9, 0x4c, 0x57, 0xff, 0xc2,
	0xec, 0xff, 0x2f, 0xa2, 0x5f, 0x1e, 0xc4, 0xed, 0x61, 0x2f, 0xa4, 0x6a, 0x8b, 0xac, 0xa6, 0x7f,
	0x03, 0xfb, 0x59, 0x59, 0xcb, 0xe8, 0xc6, 0x89, 0x3c, 0x49, 0x04, 0x4f, 0x6e, 0x59, 0x16, 0xdb,
	0x15, 0xe3, 0x33, 0x00, 0xa7, 0x7b, 0x71, 0xce, 0xac, 0xca, 0x80, 0x35, 0x32, 0x88, 0x30, 0x58,
	0xa6, 0x2b, 0x1d, 0xfa, 0x33, 0x3f, 0x0e, 0xd8, 0x40, 0xa7, 0x2b, 0x1d, 0xfd, 0x80, 0xa3, 0xf7,
	0x94, 0x93, 0xe2, 0x86, 0x3e, 0xa6, 0xcb, 0xab, 0x5a, 0x43, 0x08, 0xc3, 0xba, 0x4f, 0xe5, 0x42,
	0x1d, 0xe2, 0x13, 0x78, 0xa0, 0x0d, 0x37, 0x85, 0xb6, 0x04, 0x07, 0xcc, 0x45, 0x65, 0x5e, 0x11,
	0xd7, 0x32, 0x0f, 0xfd, 0x99, 0x17, 0xef, 0x31, 0x17, 0xe1, 0x53, 0x38, 0x50, 0xf4, 0x4d, 0x7e,
	0xa5, 0x2f, 0x67, 0x26, 0x1c, 0x58, 0xd3, 0xfe, 0x26, 0xa2, 0x57, 0x10, 0x54, 0xb4, 0xce, 0x9f,
	0x29, 0xec, 0x2b, 0xf7, 0x76, 0xc4, 0xbb, 0x78, 0xfe, 0xbb, 0x0f, 0xc7, 0x8d, 0xf9, 0xcf, 0x0a,
	0x73, 0x2d, 0x55, 0x66, 0xbe, 0xe3, 0x02, 0x46, 0xed, 0xe5, 0xe0, 0x5d, 0x2b, 0x9b, 0x8e, 0x2b,
	0xeb, 0x1a, 0x48, 0xd4, 0xc3, 0x4f, 0x70, 0xd4, 0x71, 0x39, 0x77, 0x37, 0x7a, 0xbe, 0x03, 0xbb,
	0xef, 0x2d, 0xea, 0xe1, 0x1a, 0x5e, 0xdc, 0xeb, 0x3c, 0xbe, 0xee, 0x22, 0xb9, 0xed, 0x40, 0x3a,
	0xe5, 0xcf, 0x3f, 0xc0, 0xc3, 0x72, 0x93, 0xce, 0x4c, 0xa9, 0xf0, 0x2d, 0x04, 0x4d, 0x67, 0x71,
	0x52, 0x72, 0x74, 0x78, 0x3d, 0x1d, 0x95, 0x40, 0xd3, 0x85, 0xa8, 0xf7, 0x6e, 0xf8, 0x79, 0xcf,
	0xfe, 0xa5, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xc4, 0xc9, 0xb3, 0xd4, 0x03, 0x00, 0x00,
}
