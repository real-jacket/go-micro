// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package go_micro_srv_consignment

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight,proto3" json:"weight,omitempty"`
	Container            []*Container `protobuf:"bytes,4,rep,name=container,proto3" json:"container,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId,proto3" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{0}
}

func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (m *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(m, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainer() []*Container {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{1}
}

func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (m *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(m, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment,proto3" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments,proto3" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5e5ab05dfa973d5, []int{3}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_e5e5ab05dfa973d5)
}

var fileDescriptor_e5e5ab05dfa973d5 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x49, 0xff, 0xe7, 0x52, 0x51, 0xe1, 0x01, 0x2c, 0x18, 0x88, 0x52, 0x90, 0x3a, 0x05,
	0xa9, 0x3c, 0x01, 0xea, 0x50, 0x65, 0x75, 0x47, 0x06, 0x28, 0xf1, 0x29, 0xb5, 0x44, 0xec, 0x60,
	0xbb, 0xe5, 0xd1, 0xe0, 0x31, 0x78, 0x24, 0x54, 0xa7, 0xa1, 0x46, 0xa8, 0xa8, 0x9b, 0x7f, 0x77,
	0xfe, 0x2e, 0x5f, 0x4e, 0x86, 0x71, 0xa5, 0x95, 0x55, 0x77, 0xb9, 0x92, 0x46, 0x14, 0xb2, 0x44,
	0x69, 0xfd, 0x73, 0xea, 0xba, 0x84, 0x16, 0x2a, 0x2d, 0x45, 0xae, 0x55, 0x6a, 0xf4, 0x26, 0xf5,
	0xfa, 0xc9, 0x47, 0x00, 0xd1, 0x6c, 0x9f, 0xc9, 0x29, 0xb4, 0x04, 0xa7, 0x41, 0x1c, 0x4c, 0x42,
	0xd6, 0x12, 0x9c, 0xc4, 0x10, 0x71, 0x34, 0xb9, 0x16, 0x95, 0x15, 0x4a, 0xd2, 0x96, 0x6b, 0xf8,
	0x25, 0x72, 0x0e, 0xbd, 0x77, 0x14, 0xc5, 0xca, 0xd2, 0x76, 0x1c, 0x4c, 0xba, 0x6c, 0x97, 0xc8,
	0x03, 0x84, 0xb9, 0x92, 0x76, 0x29, 0x24, 0x6a, 0xda, 0x89, 0xdb, 0x93, 0x68, 0x3a, 0x4e, 0x0f,
	0x79, 0xa4, 0xb3, 0xe6, 0x2a, 0xdb, 0x53, 0xe4, 0x0a, 0xc2, 0x0d, 0x1a, 0x83, 0xaf, 0x4f, 0x82,
	0xd3, 0xae, 0xfb, 0xf4, 0xa0, 0x2e, 0x64, 0x3c, 0x29, 0x21, 0xfc, 0x81, 0xfe, 0x68, 0x5f, 0x43,
	0x94, 0xaf, 0x8d, 0x55, 0x25, 0xea, 0x2d, 0x5b, 0x6b, 0x43, 0x53, 0xca, 0xf8, 0xd6, 0x5a, 0x69,
	0x51, 0x08, 0xe9, 0xac, 0x43, 0xb6, 0x4b, 0xe4, 0x02, 0xfa, 0x6b, 0x53, 0x43, 0x9d, 0xba, 0xb1,
	0x8d, 0x19, 0x4f, 0x86, 0x00, 0x73, 0xb4, 0x0c, 0xdf, 0xd6, 0x68, 0x6c, 0xf2, 0x19, 0xc0, 0x80,
	0xa1, 0xa9, 0x94, 0x34, 0x48, 0x28, 0xf4, 0x73, 0x8d, 0x4b, 0x8b, 0xb5, 0xc1, 0x80, 0x35, 0x91,
	0xcc, 0x21, 0xf2, 0x7e, 0xd2, 0x69, 0x44, 0xd3, 0xdb, 0x7f, 0xb7, 0xd0, 0x9c, 0x99, 0x4f, 0x92,
	0x0c, 0x86, 0x5e, 0x34, 0xb4, 0xed, 0xf6, 0x79, 0xe4, 0xa4, 0x5f, 0xe8, 0xf4, 0x2b, 0x80, 0xd1,
	0x62, 0x25, 0xaa, 0x4a, 0xc8, 0x62, 0x81, 0x7a, 0x23, 0x72, 0x24, 0xcf, 0x70, 0x36, 0x73, 0xca,
	0xfe, 0x53, 0x38, 0x6e, 0xfa, 0x65, 0x72, 0xf8, 0x5a, 0xb3, 0xa1, 0xe4, 0x84, 0x3c, 0xc2, 0x68,
	0x8e, 0xd6, 0xe3, 0x0c, 0xb9, 0x39, 0x0c, 0xee, 0x37, 0x7d, 0xcc, 0xf8, 0x97, 0x9e, 0x7b, 0xe5,
	0xf7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xaa, 0xea, 0xfa, 0x0c, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}