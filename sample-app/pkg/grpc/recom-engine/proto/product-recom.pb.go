// Code generated by protoc-gen-go. DO NOT EDIT.
// source: product-recom.proto

package recomengine

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetProductRecomRequest struct {
	UserID               int64    `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProductRecomRequest) Reset()         { *m = GetProductRecomRequest{} }
func (m *GetProductRecomRequest) String() string { return proto.CompactTextString(m) }
func (*GetProductRecomRequest) ProtoMessage()    {}
func (*GetProductRecomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_recom_b4e3faaf20c90db0, []int{0}
}
func (m *GetProductRecomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRecomRequest.Unmarshal(m, b)
}
func (m *GetProductRecomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRecomRequest.Marshal(b, m, deterministic)
}
func (dst *GetProductRecomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRecomRequest.Merge(dst, src)
}
func (m *GetProductRecomRequest) XXX_Size() int {
	return xxx_messageInfo_GetProductRecomRequest.Size(m)
}
func (m *GetProductRecomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRecomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRecomRequest proto.InternalMessageInfo

func (m *GetProductRecomRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *GetProductRecomRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetProductRecomResponse struct {
	Data                 []*ProductRecom   `protobuf:"bytes,1,rep,name=data" json:"data,omitempty"`
	Meta                 *ProductRecomMeta `protobuf:"bytes,2,opt,name=meta" json:"meta,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetProductRecomResponse) Reset()         { *m = GetProductRecomResponse{} }
func (m *GetProductRecomResponse) String() string { return proto.CompactTextString(m) }
func (*GetProductRecomResponse) ProtoMessage()    {}
func (*GetProductRecomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_recom_b4e3faaf20c90db0, []int{1}
}
func (m *GetProductRecomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProductRecomResponse.Unmarshal(m, b)
}
func (m *GetProductRecomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProductRecomResponse.Marshal(b, m, deterministic)
}
func (dst *GetProductRecomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProductRecomResponse.Merge(dst, src)
}
func (m *GetProductRecomResponse) XXX_Size() int {
	return xxx_messageInfo_GetProductRecomResponse.Size(m)
}
func (m *GetProductRecomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProductRecomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProductRecomResponse proto.InternalMessageInfo

func (m *GetProductRecomResponse) GetData() []*ProductRecom {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *GetProductRecomResponse) GetMeta() *ProductRecomMeta {
	if m != nil {
		return m.Meta
	}
	return nil
}

type ProductRecomMeta struct {
	OverallAffinities    float32  `protobuf:"fixed32,1,opt,name=overallAffinities" json:"overallAffinities,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRecomMeta) Reset()         { *m = ProductRecomMeta{} }
func (m *ProductRecomMeta) String() string { return proto.CompactTextString(m) }
func (*ProductRecomMeta) ProtoMessage()    {}
func (*ProductRecomMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_recom_b4e3faaf20c90db0, []int{2}
}
func (m *ProductRecomMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRecomMeta.Unmarshal(m, b)
}
func (m *ProductRecomMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRecomMeta.Marshal(b, m, deterministic)
}
func (dst *ProductRecomMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRecomMeta.Merge(dst, src)
}
func (m *ProductRecomMeta) XXX_Size() int {
	return xxx_messageInfo_ProductRecomMeta.Size(m)
}
func (m *ProductRecomMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRecomMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRecomMeta proto.InternalMessageInfo

func (m *ProductRecomMeta) GetOverallAffinities() float32 {
	if m != nil {
		return m.OverallAffinities
	}
	return 0
}

type ProductRecom struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Price                int64    `protobuf:"varint,2,opt,name=price" json:"price,omitempty"`
	ShopId               int64    `protobuf:"varint,3,opt,name=shop_id,json=shopId" json:"shop_id,omitempty"`
	Created              int64    `protobuf:"varint,4,opt,name=created" json:"created,omitempty"`
	Url                  string   `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	Name                 string   `protobuf:"bytes,6,opt,name=name" json:"name,omitempty"`
	Images               []string `protobuf:"bytes,7,rep,name=images" json:"images,omitempty"`
	Affinity             float32  `protobuf:"fixed32,8,opt,name=affinity" json:"affinity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRecom) Reset()         { *m = ProductRecom{} }
func (m *ProductRecom) String() string { return proto.CompactTextString(m) }
func (*ProductRecom) ProtoMessage()    {}
func (*ProductRecom) Descriptor() ([]byte, []int) {
	return fileDescriptor_product_recom_b4e3faaf20c90db0, []int{3}
}
func (m *ProductRecom) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRecom.Unmarshal(m, b)
}
func (m *ProductRecom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRecom.Marshal(b, m, deterministic)
}
func (dst *ProductRecom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRecom.Merge(dst, src)
}
func (m *ProductRecom) XXX_Size() int {
	return xxx_messageInfo_ProductRecom.Size(m)
}
func (m *ProductRecom) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRecom.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRecom proto.InternalMessageInfo

func (m *ProductRecom) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ProductRecom) GetPrice() int64 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *ProductRecom) GetShopId() int64 {
	if m != nil {
		return m.ShopId
	}
	return 0
}

func (m *ProductRecom) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ProductRecom) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *ProductRecom) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProductRecom) GetImages() []string {
	if m != nil {
		return m.Images
	}
	return nil
}

func (m *ProductRecom) GetAffinity() float32 {
	if m != nil {
		return m.Affinity
	}
	return 0
}

func init() {
	proto.RegisterType((*GetProductRecomRequest)(nil), "recomengine.GetProductRecomRequest")
	proto.RegisterType((*GetProductRecomResponse)(nil), "recomengine.GetProductRecomResponse")
	proto.RegisterType((*ProductRecomMeta)(nil), "recomengine.productRecomMeta")
	proto.RegisterType((*ProductRecom)(nil), "recomengine.productRecom")
}

func init() { proto.RegisterFile("product-recom.proto", fileDescriptor_product_recom_b4e3faaf20c90db0) }

var fileDescriptor_product_recom_b4e3faaf20c90db0 = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xdd, 0x4a, 0x3b, 0x31,
	0x10, 0xc5, 0xd9, 0x8f, 0x7e, 0x4d, 0xff, 0xfc, 0xa9, 0xa3, 0xb4, 0x51, 0x10, 0x96, 0xbd, 0xda,
	0x0b, 0x5b, 0xb0, 0xbe, 0x80, 0x82, 0x28, 0xbd, 0x10, 0x24, 0x2f, 0x20, 0x71, 0x33, 0xad, 0x81,
	0xdd, 0xcd, 0x9a, 0x64, 0x05, 0xf1, 0xe9, 0x7c, 0x33, 0xd9, 0x6c, 0x95, 0xc5, 0xe2, 0xdd, 0xfc,
	0x72, 0x4e, 0x26, 0x73, 0x26, 0x70, 0x5c, 0x1b, 0x2d, 0x9b, 0xdc, 0x2d, 0x0d, 0xe5, 0xba, 0x5c,
	0xd5, 0x46, 0x3b, 0x8d, 0x53, 0x0f, 0x54, 0xed, 0x54, 0x45, 0xe9, 0x1d, 0xcc, 0xef, 0xc9, 0x3d,
	0x76, 0x36, 0xde, 0x0a, 0x9c, 0x5e, 0x1b, 0xb2, 0x0e, 0xe7, 0x30, 0x6c, 0x2c, 0x99, 0xcd, 0x2d,
	0x0b, 0x92, 0x20, 0x8b, 0xf8, 0x9e, 0xf0, 0x04, 0x06, 0x85, 0x2a, 0x95, 0x63, 0x61, 0x12, 0x64,
	0x03, 0xde, 0x41, 0xfa, 0x01, 0x8b, 0x83, 0x3e, 0xb6, 0xd6, 0x95, 0x25, 0x5c, 0x42, 0x2c, 0x85,
	0x13, 0x2c, 0x48, 0xa2, 0x6c, 0xba, 0x3e, 0x5d, 0xf5, 0x9e, 0x5f, 0xd5, 0xfd, 0x0b, 0xde, 0x86,
	0x97, 0x10, 0x97, 0xe4, 0x84, 0x6f, 0x3f, 0x5d, 0x9f, 0xff, 0x69, 0x7f, 0x20, 0x27, 0xb8, 0xb7,
	0xa6, 0xd7, 0x30, 0xfb, 0xad, 0xe0, 0x05, 0x1c, 0xe9, 0x37, 0x32, 0xa2, 0x28, 0x6e, 0xb6, 0x5b,
	0x55, 0x29, 0xa7, 0xc8, 0xfa, 0x24, 0x21, 0x3f, 0x14, 0xd2, 0xcf, 0x00, 0xfe, 0xf5, 0x5b, 0xe0,
	0x7f, 0x08, 0x95, 0xdc, 0x27, 0x0f, 0x95, 0x6c, 0x53, 0xd7, 0x46, 0xe5, 0xe4, 0xc7, 0x8a, 0x78,
	0x07, 0xb8, 0x80, 0x91, 0x7d, 0xd1, 0xf5, 0x93, 0x92, 0x2c, 0xea, 0x96, 0xd4, 0xe2, 0x46, 0x22,
	0x83, 0x51, 0x6e, 0x48, 0x38, 0x92, 0x2c, 0xf6, 0xc2, 0x37, 0xe2, 0x0c, 0xa2, 0xc6, 0x14, 0x6c,
	0x90, 0x04, 0xd9, 0x84, 0xb7, 0x25, 0x22, 0xc4, 0x95, 0x28, 0x89, 0x0d, 0xfd, 0x91, 0xaf, 0xdb,
	0xe5, 0xab, 0x52, 0xec, 0xc8, 0xb2, 0x51, 0x12, 0x65, 0x13, 0xbe, 0x27, 0x3c, 0x83, 0xb1, 0xe8,
	0xa6, 0x7e, 0x67, 0x63, 0x1f, 0xe6, 0x87, 0x9f, 0x87, 0xfe, 0x7b, 0xaf, 0xbe, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x6c, 0x22, 0xd0, 0xd3, 0xf5, 0x01, 0x00, 0x00,
}
