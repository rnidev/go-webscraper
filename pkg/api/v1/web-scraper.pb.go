// Code generated by protoc-gen-go. DO NOT EDIT.
// source: web-scraper.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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

//Project Object
type Product struct {
	Asin                 string               `protobuf:"bytes,1,opt,name=asin,proto3" json:"asin,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Categories           []*ProductCategory   `protobuf:"bytes,3,rep,name=categories,proto3" json:"categories,omitempty"`
	Ranks                []*ProductRank       `protobuf:"bytes,4,rep,name=ranks,proto3" json:"ranks,omitempty"`
	Dimensions           []string             `protobuf:"bytes,5,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Product) Reset()         { *m = Product{} }
func (m *Product) String() string { return proto.CompactTextString(m) }
func (*Product) ProtoMessage()    {}
func (*Product) Descriptor() ([]byte, []int) {
	return fileDescriptor_c181f2f37fbaedca, []int{0}
}

func (m *Product) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Product.Unmarshal(m, b)
}
func (m *Product) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Product.Marshal(b, m, deterministic)
}
func (m *Product) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Product.Merge(m, src)
}
func (m *Product) XXX_Size() int {
	return xxx_messageInfo_Product.Size(m)
}
func (m *Product) XXX_DiscardUnknown() {
	xxx_messageInfo_Product.DiscardUnknown(m)
}

var xxx_messageInfo_Product proto.InternalMessageInfo

func (m *Product) GetAsin() string {
	if m != nil {
		return m.Asin
	}
	return ""
}

func (m *Product) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Product) GetCategories() []*ProductCategory {
	if m != nil {
		return m.Categories
	}
	return nil
}

func (m *Product) GetRanks() []*ProductRank {
	if m != nil {
		return m.Ranks
	}
	return nil
}

func (m *Product) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Product) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

//ProjectCategoryObject
type ProductCategory struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level                int64    `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductCategory) Reset()         { *m = ProductCategory{} }
func (m *ProductCategory) String() string { return proto.CompactTextString(m) }
func (*ProductCategory) ProtoMessage()    {}
func (*ProductCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_c181f2f37fbaedca, []int{1}
}

func (m *ProductCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductCategory.Unmarshal(m, b)
}
func (m *ProductCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductCategory.Marshal(b, m, deterministic)
}
func (m *ProductCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductCategory.Merge(m, src)
}
func (m *ProductCategory) XXX_Size() int {
	return xxx_messageInfo_ProductCategory.Size(m)
}
func (m *ProductCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductCategory.DiscardUnknown(m)
}

var xxx_messageInfo_ProductCategory proto.InternalMessageInfo

func (m *ProductCategory) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductCategory) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

//ProjectRankObject
type ProductRank struct {
	RankInfo             string   `protobuf:"bytes,1,opt,name=rank_info,json=rankInfo,proto3" json:"rank_info,omitempty"`
	Level                int64    `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRank) Reset()         { *m = ProductRank{} }
func (m *ProductRank) String() string { return proto.CompactTextString(m) }
func (*ProductRank) ProtoMessage()    {}
func (*ProductRank) Descriptor() ([]byte, []int) {
	return fileDescriptor_c181f2f37fbaedca, []int{2}
}

func (m *ProductRank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRank.Unmarshal(m, b)
}
func (m *ProductRank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRank.Marshal(b, m, deterministic)
}
func (m *ProductRank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRank.Merge(m, src)
}
func (m *ProductRank) XXX_Size() int {
	return xxx_messageInfo_ProductRank.Size(m)
}
func (m *ProductRank) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRank.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRank proto.InternalMessageInfo

func (m *ProductRank) GetRankInfo() string {
	if m != nil {
		return m.RankInfo
	}
	return ""
}

func (m *ProductRank) GetLevel() int64 {
	if m != nil {
		return m.Level
	}
	return 0
}

//Expected Request For GetProduct
type GetProductRequest struct {
	Asin                 string   `protobuf:"bytes,1,opt,name=asin,proto3" json:"asin,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRequest) Reset()         { *m = GetProductRequest{} }
func (m *GetProductRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRequest) ProtoMessage()    {}
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c181f2f37fbaedca, []int{3}
}

func (m *GetProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRequest.Unmarshal(m, b)
}
func (m *GetProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRequest.Marshal(b, m, deterministic)
}
func (m *GetProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRequest.Merge(m, src)
}
func (m *GetProductRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductRequest.Size(m)
}
func (m *GetProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRequest proto.InternalMessageInfo

func (m *GetProductRequest) GetAsin() string {
	if m != nil {
		return m.Asin
	}
	return ""
}

//Expected Response From GetProduct
type GetProductResponse struct {
	Product              *Product `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductResponse) Reset()         { *m = GetProductResponse{} }
func (m *GetProductResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductResponse) ProtoMessage()    {}
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c181f2f37fbaedca, []int{4}
}

func (m *GetProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductResponse.Unmarshal(m, b)
}
func (m *GetProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductResponse.Marshal(b, m, deterministic)
}
func (m *GetProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductResponse.Merge(m, src)
}
func (m *GetProductResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductResponse.Size(m)
}
func (m *GetProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductResponse proto.InternalMessageInfo

func (m *GetProductResponse) GetProduct() *Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func init() {
	proto.RegisterType((*Product)(nil), "v1.Product")
	proto.RegisterType((*ProductCategory)(nil), "v1.ProductCategory")
	proto.RegisterType((*ProductRank)(nil), "v1.ProductRank")
	proto.RegisterType((*GetProductRequest)(nil), "v1.GetProductRequest")
	proto.RegisterType((*GetProductResponse)(nil), "v1.GetProductResponse")
}

func init() { proto.RegisterFile("web-scraper.proto", fileDescriptor_c181f2f37fbaedca) }

var fileDescriptor_c181f2f37fbaedca = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x71, 0xd2, 0xa6, 0x64, 0x22, 0x51, 0xba, 0x7c, 0x28, 0x0a, 0xa8, 0xac, 0x22, 0x95,
	0x46, 0x15, 0xf1, 0x36, 0x6e, 0x2f, 0xb4, 0x97, 0x16, 0x0e, 0x88, 0x0b, 0x42, 0x06, 0x54, 0x89,
	0x4b, 0xb5, 0xb6, 0xa7, 0xce, 0xd2, 0x78, 0xd7, 0xdd, 0x5d, 0x3b, 0x7c, 0x88, 0x0b, 0x8f, 0x00,
	0x37, 0xde, 0x85, 0x03, 0xcf, 0xc0, 0x2b, 0x70, 0xe0, 0x31, 0x90, 0x3f, 0x02, 0x56, 0x03, 0x17,
	0x7b, 0xfd, 0x9f, 0xdf, 0xce, 0xfc, 0x67, 0x34, 0x86, 0x8d, 0x39, 0x06, 0x63, 0x13, 0x6a, 0x9e,
	0xa2, 0x76, 0x53, 0xad, 0xac, 0x22, 0xad, 0x7c, 0x32, 0xb8, 0x17, 0x2b, 0x15, 0xcf, 0x90, 0x95,
	0x4a, 0x90, 0x9d, 0x31, 0x2b, 0x12, 0x34, 0x96, 0x27, 0x69, 0x05, 0x0d, 0xee, 0xd6, 0x00, 0x4f,
	0x05, 0xe3, 0x52, 0x2a, 0xcb, 0xad, 0x50, 0xd2, 0xd4, 0xd1, 0x07, 0xe5, 0x2b, 0x1c, 0xc7, 0x28,
	0xc7, 0x66, 0xce, 0xe3, 0x18, 0x35, 0x53, 0x69, 0x49, 0x2c, 0xd3, 0xc3, 0x5f, 0x0e, 0xac, 0x3d,
	0xd7, 0x2a, 0xca, 0x42, 0x4b, 0x08, 0xac, 0x70, 0x23, 0x64, 0xdf, 0xa1, 0xce, 0xa8, 0xeb, 0x97,
	0xe7, 0x42, 0x93, 0x3c, 0xc1, 0x7e, 0xab, 0xd2, 0x8a, 0x33, 0xd9, 0x03, 0x08, 0xb9, 0xc5, 0x58,
	0x69, 0x81, 0xa6, 0xdf, 0xa6, 0xed, 0x51, 0xcf, 0xbb, 0xe1, 0xe6, 0x13, 0xb7, 0x4e, 0xf4, 0xb8,
	0x0a, 0xbe, 0xf3, 0x1b, 0x18, 0xd9, 0x82, 0x55, 0xcd, 0xe5, 0xb9, 0xe9, 0xaf, 0x94, 0xfc, 0x7a,
	0x83, 0xf7, 0xb9, 0x3c, 0xf7, 0xab, 0x28, 0xd9, 0x04, 0x88, 0x44, 0x82, 0xd2, 0x14, 0x1e, 0xfb,
	0xab, 0xb4, 0x3d, 0xea, 0xfa, 0x0d, 0x85, 0x3c, 0x04, 0x08, 0x35, 0x72, 0x8b, 0xd1, 0x29, 0xb7,
	0xfd, 0x0e, 0x75, 0x46, 0x3d, 0x6f, 0xe0, 0x56, 0x03, 0x71, 0x17, 0x13, 0x73, 0x5f, 0x2e, 0x26,
	0xe6, 0x77, 0x6b, 0xfa, 0xd8, 0x0e, 0x0f, 0x61, 0xfd, 0x92, 0xc1, 0x3f, 0xdd, 0x39, 0x8d, 0xee,
	0x6e, 0xc2, 0xea, 0x0c, 0x73, 0x9c, 0x95, 0x2d, 0xb7, 0xfd, 0xea, 0x63, 0x78, 0x04, 0xbd, 0x86,
	0x5b, 0x72, 0x07, 0xba, 0x85, 0xdf, 0x53, 0x21, 0xcf, 0x54, 0x7d, 0xfb, 0x6a, 0x21, 0x3c, 0x95,
	0x67, 0xea, 0x3f, 0x19, 0xb6, 0x61, 0xe3, 0x09, 0xda, 0x45, 0x12, 0xbc, 0xc8, 0xd0, 0xfc, 0x73,
	0xe4, 0xc3, 0x43, 0x20, 0x4d, 0xd0, 0xa4, 0x4a, 0x1a, 0x24, 0x5b, 0xb0, 0x96, 0x56, 0x52, 0x09,
	0xf7, 0xbc, 0x5e, 0x73, 0x82, 0x8b, 0x98, 0x77, 0x01, 0x70, 0x82, 0xc1, 0x8b, 0x6a, 0xa9, 0x48,
	0x08, 0xf0, 0x37, 0x15, 0xb9, 0x55, 0xdc, 0x58, 0xf2, 0x30, 0xb8, 0x7d, 0x59, 0xae, 0x2a, 0x0e,
	0xef, 0x7f, 0xfa, 0xf1, 0xf3, 0x4b, 0x8b, 0x92, 0x4d, 0x96, 0x4f, 0x18, 0x4f, 0xf8, 0x7b, 0x25,
	0x59, 0x5d, 0x86, 0x15, 0x46, 0xd9, 0x87, 0xe2, 0xf9, 0xf1, 0xd1, 0x77, 0xe7, 0xf3, 0xf1, 0x37,
	0x87, 0xbc, 0x82, 0xde, 0x09, 0x06, 0xb4, 0x2e, 0x3d, 0x3c, 0x86, 0x8e, 0x9f, 0x09, 0xfa, 0x4c,
	0x90, 0xed, 0xa9, 0xb5, 0xa9, 0x39, 0x60, 0x2c, 0x16, 0x76, 0x9a, 0x05, 0x6e, 0xa8, 0x12, 0xa6,
	0xa5, 0x88, 0x30, 0x67, 0xb1, 0x1a, 0xcf, 0x31, 0xa8, 0x7f, 0x81, 0xc1, 0x35, 0x9d, 0x89, 0xa3,
	0x08, 0x73, 0x2d, 0x45, 0x01, 0x79, 0xed, 0x89, 0xbb, 0xbb, 0xe3, 0x38, 0xde, 0x75, 0x9e, 0xa6,
	0x33, 0x11, 0x96, 0x7b, 0xcb, 0xde, 0x18, 0x25, 0x0f, 0x96, 0x14, 0xff, 0x10, 0xda, 0xfb, 0xbb,
	0xfb, 0x64, 0x1f, 0x76, 0x7c, 0xb4, 0x99, 0x96, 0x18, 0xd1, 0xf9, 0x14, 0x25, 0xb5, 0x53, 0xa4,
	0x1a, 0x8d, 0xca, 0x74, 0x88, 0x34, 0x52, 0x68, 0xa8, 0x54, 0x96, 0xe2, 0x5b, 0x61, 0xac, 0x4b,
	0x3a, 0xb0, 0xf2, 0xb5, 0xe5, 0xac, 0xbd, 0xbe, 0x12, 0x74, 0xca, 0xd5, 0xd9, 0xfb, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x02, 0xb3, 0x31, 0x85, 0x93, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WebScraperClient is the client API for WebScraper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WebScraperClient interface {
	//This end point takes Amazon Product ASIN, and returns name, categories, ranks and dimensions
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
}

type webScraperClient struct {
	cc *grpc.ClientConn
}

func NewWebScraperClient(cc *grpc.ClientConn) WebScraperClient {
	return &webScraperClient{cc}
}

func (c *webScraperClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/v1.WebScraper/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebScraperServer is the server API for WebScraper service.
type WebScraperServer interface {
	//This end point takes Amazon Product ASIN, and returns name, categories, ranks and dimensions
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
}

func RegisterWebScraperServer(s *grpc.Server, srv WebScraperServer) {
	s.RegisterService(&_WebScraper_serviceDesc, srv)
}

func _WebScraper_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebScraperServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.WebScraper/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebScraperServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebScraper_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.WebScraper",
	HandlerType: (*WebScraperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduct",
			Handler:    _WebScraper_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "web-scraper.proto",
}