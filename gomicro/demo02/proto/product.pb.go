// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product.proto

package proto

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

// 商品模块 - 请求参数
type GetProductByIDRequest struct {
	Pid                  int64    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductByIDRequest) Reset()         { *m = GetProductByIDRequest{} }
func (m *GetProductByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductByIDRequest) ProtoMessage()    {}
func (*GetProductByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{0}
}

func (m *GetProductByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductByIDRequest.Unmarshal(m, b)
}
func (m *GetProductByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetProductByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductByIDRequest.Merge(m, src)
}
func (m *GetProductByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductByIDRequest.Size(m)
}
func (m *GetProductByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductByIDRequest proto.InternalMessageInfo

func (m *GetProductByIDRequest) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

// 商品模块 - 返回参数
type GetProductByIDResponse struct {
	Code                 int64                       `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string                      `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data                 *GetProductByIDResponseData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *GetProductByIDResponse) Reset()         { *m = GetProductByIDResponse{} }
func (m *GetProductByIDResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductByIDResponse) ProtoMessage()    {}
func (*GetProductByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{1}
}

func (m *GetProductByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductByIDResponse.Unmarshal(m, b)
}
func (m *GetProductByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductByIDResponse.Marshal(b, m, deterministic)
}
func (m *GetProductByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductByIDResponse.Merge(m, src)
}
func (m *GetProductByIDResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductByIDResponse.Size(m)
}
func (m *GetProductByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductByIDResponse proto.InternalMessageInfo

func (m *GetProductByIDResponse) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GetProductByIDResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *GetProductByIDResponse) GetData() *GetProductByIDResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetProductByIDResponseData struct {
	Pid                  int64    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductByIDResponseData) Reset()         { *m = GetProductByIDResponseData{} }
func (m *GetProductByIDResponseData) String() string { return proto.CompactTextString(m) }
func (*GetProductByIDResponseData) ProtoMessage()    {}
func (*GetProductByIDResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f0fd8b59378f44a5, []int{2}
}

func (m *GetProductByIDResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductByIDResponseData.Unmarshal(m, b)
}
func (m *GetProductByIDResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductByIDResponseData.Marshal(b, m, deterministic)
}
func (m *GetProductByIDResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductByIDResponseData.Merge(m, src)
}
func (m *GetProductByIDResponseData) XXX_Size() int {
	return xxx_messageInfo_GetProductByIDResponseData.Size(m)
}
func (m *GetProductByIDResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductByIDResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductByIDResponseData proto.InternalMessageInfo

func (m *GetProductByIDResponseData) GetPid() int64 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *GetProductByIDResponseData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*GetProductByIDRequest)(nil), "proto.GetProductByIDRequest")
	proto.RegisterType((*GetProductByIDResponse)(nil), "proto.GetProductByIDResponse")
	proto.RegisterType((*GetProductByIDResponseData)(nil), "proto.GetProductByIDResponseData")
}

func init() { proto.RegisterFile("product.proto", fileDescriptor_f0fd8b59378f44a5) }

var fileDescriptor_f0fd8b59378f44a5 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x28, 0xca, 0x4f,
	0x29, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x9a, 0x5c,
	0xa2, 0xee, 0xa9, 0x25, 0x01, 0x10, 0x29, 0xa7, 0x4a, 0x4f, 0x97, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x01, 0x2e, 0xe6, 0x82, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xe6, 0x20,
	0x10, 0x53, 0xa9, 0x94, 0x4b, 0x0c, 0x5d, 0x69, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x10,
	0x17, 0x4b, 0x72, 0x7e, 0x4a, 0x2a, 0x54, 0x31, 0x98, 0x0d, 0xd2, 0x9f, 0x5b, 0x9c, 0x2e, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x62, 0x0a, 0x99, 0x72, 0xb1, 0xa4, 0x24, 0x96, 0x24, 0x4a,
	0x30, 0x2b, 0x30, 0x6a, 0x70, 0x1b, 0x29, 0x42, 0xdc, 0xa1, 0x87, 0xdd, 0x48, 0x97, 0xc4, 0x92,
	0xc4, 0x20, 0xb0, 0x72, 0x25, 0x27, 0x2e, 0x29, 0xdc, 0x6a, 0x30, 0x9d, 0x09, 0x72, 0x4c, 0x5e,
	0x62, 0x6e, 0x2a, 0xd4, 0x66, 0x30, 0xdb, 0x28, 0x8a, 0x8b, 0x1d, 0x6a, 0x80, 0x90, 0x3f, 0x17,
	0x1f, 0xaa, 0x71, 0x42, 0x32, 0x38, 0x5c, 0x02, 0x0e, 0x07, 0x29, 0x59, 0xbc, 0xee, 0x54, 0x62,
	0x48, 0x62, 0x03, 0xcb, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x16, 0xfb, 0xa3, 0x67, 0x60,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Product service

type ProductClient interface {
	// 根据商品ID获取商品信息
	GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...client.CallOption) (*GetProductByIDResponse, error)
}

type productClient struct {
	c           client.Client
	serviceName string
}

func NewProductClient(serviceName string, c client.Client) ProductClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proto"
	}
	return &productClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *productClient) GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...client.CallOption) (*GetProductByIDResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Product.GetProductByID", in)
	out := new(GetProductByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Product service

type ProductHandler interface {
	// 根据商品ID获取商品信息
	GetProductByID(context.Context, *GetProductByIDRequest, *GetProductByIDResponse) error
}

func RegisterProductHandler(s server.Server, hdlr ProductHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Product{hdlr}, opts...))
}

type Product struct {
	ProductHandler
}

func (h *Product) GetProductByID(ctx context.Context, in *GetProductByIDRequest, out *GetProductByIDResponse) error {
	return h.ProductHandler.GetProductByID(ctx, in, out)
}