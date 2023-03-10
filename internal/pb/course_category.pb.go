// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/course_category.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	proto/course_category.proto

It has these top-level messages:
	Blank
	Category
	CategoryResponse
	CreateCategoryRequest
	CategoryList
	GetCategoryRequest
*/
package pb

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

type Blank struct {
}

func (m *Blank) Reset()                    { *m = Blank{} }
func (m *Blank) String() string            { return proto.CompactTextString(m) }
func (*Blank) ProtoMessage()               {}
func (*Blank) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Category struct {
	Id          string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
}

func (m *Category) Reset()                    { *m = Category{} }
func (m *Category) String() string            { return proto.CompactTextString(m) }
func (*Category) ProtoMessage()               {}
func (*Category) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Category) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Category) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Category) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CategoryResponse struct {
	Category *Category `protobuf:"bytes,1,opt,name=category" json:"category,omitempty"`
}

func (m *CategoryResponse) Reset()                    { *m = CategoryResponse{} }
func (m *CategoryResponse) String() string            { return proto.CompactTextString(m) }
func (*CategoryResponse) ProtoMessage()               {}
func (*CategoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CategoryResponse) GetCategory() *Category {
	if m != nil {
		return m.Category
	}
	return nil
}

type CreateCategoryRequest struct {
	Name        string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *CreateCategoryRequest) Reset()                    { *m = CreateCategoryRequest{} }
func (m *CreateCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateCategoryRequest) ProtoMessage()               {}
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CreateCategoryRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateCategoryRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CategoryList struct {
	Categories []*Category `protobuf:"bytes,1,rep,name=categories" json:"categories,omitempty"`
}

func (m *CategoryList) Reset()                    { *m = CategoryList{} }
func (m *CategoryList) String() string            { return proto.CompactTextString(m) }
func (*CategoryList) ProtoMessage()               {}
func (*CategoryList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CategoryList) GetCategories() []*Category {
	if m != nil {
		return m.Categories
	}
	return nil
}

type GetCategoryRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetCategoryRequest) Reset()                    { *m = GetCategoryRequest{} }
func (m *GetCategoryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetCategoryRequest) ProtoMessage()               {}
func (*GetCategoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetCategoryRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Blank)(nil), "pb.Blank")
	proto.RegisterType((*Category)(nil), "pb.Category")
	proto.RegisterType((*CategoryResponse)(nil), "pb.CategoryResponse")
	proto.RegisterType((*CreateCategoryRequest)(nil), "pb.CreateCategoryRequest")
	proto.RegisterType((*CategoryList)(nil), "pb.CategoryList")
	proto.RegisterType((*GetCategoryRequest)(nil), "pb.GetCategoryRequest")
}

func init() { proto.RegisterFile("proto/course_category.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xbb, 0xa9, 0x7f, 0xda, 0x49, 0xad, 0x65, 0x50, 0x89, 0x7a, 0xa9, 0x8b, 0x87, 0x1c,
	0xa4, 0x95, 0x8a, 0x27, 0x3d, 0xb5, 0x87, 0x5e, 0x14, 0xa4, 0xbd, 0x79, 0x91, 0x4d, 0x32, 0xc8,
	0x62, 0x9b, 0xc4, 0xdd, 0xad, 0xe0, 0xe7, 0xf4, 0x0b, 0x49, 0x96, 0x26, 0xa6, 0x69, 0x14, 0xbc,
	0x2d, 0xf3, 0x1e, 0x6f, 0x7e, 0xf3, 0x58, 0x38, 0x4f, 0x55, 0x62, 0x92, 0x61, 0x98, 0xac, 0x94,
	0xa6, 0x97, 0x50, 0x18, 0x7a, 0x4d, 0xd4, 0xe7, 0xc0, 0x4e, 0xd1, 0x49, 0x03, 0xbe, 0x0f, 0xbb,
	0xe3, 0x85, 0x88, 0xdf, 0xf8, 0x13, 0xb4, 0x26, 0x6b, 0x19, 0xbb, 0xe0, 0xc8, 0xc8, 0x63, 0x7d,
	0xe6, 0xb7, 0x67, 0x8e, 0x8c, 0x10, 0x61, 0x27, 0x16, 0x4b, 0xf2, 0x1c, 0x3b, 0xb1, 0x6f, 0xec,
	0x83, 0x1b, 0x91, 0x0e, 0x95, 0x4c, 0x8d, 0x4c, 0x62, 0xaf, 0x69, 0xa5, 0xf2, 0x88, 0xdf, 0x43,
	0x2f, 0x4f, 0x9c, 0x91, 0x4e, 0x93, 0x58, 0x13, 0xfa, 0xd0, 0xca, 0x21, 0x6c, 0xbe, 0x3b, 0xea,
	0x0c, 0xd2, 0x60, 0x50, 0xf8, 0x0a, 0x95, 0x3f, 0xc2, 0xf1, 0x44, 0x91, 0x30, 0xf4, 0x93, 0xf1,
	0xbe, 0x22, 0x6d, 0x0a, 0x18, 0xf6, 0x3b, 0x8c, 0x53, 0x07, 0xd3, 0xc9, 0x83, 0x1e, 0xa4, 0x36,
	0x78, 0x05, 0xb0, 0x5e, 0x25, 0x49, 0x7b, 0xac, 0xdf, 0xdc, 0x42, 0x29, 0xe9, 0xfc, 0x12, 0x70,
	0x4a, 0xa6, 0x4a, 0x52, 0xa9, 0x69, 0xf4, 0xe5, 0xc0, 0x61, 0xee, 0x99, 0x93, 0xfa, 0x90, 0x21,
	0xe1, 0x1d, 0x74, 0x37, 0xcf, 0xc0, 0x53, 0xbb, 0xa5, 0xee, 0xb4, 0xb3, 0x0d, 0x00, 0xde, 0xc0,
	0x29, 0x1c, 0x6d, 0x1a, 0xe7, 0x46, 0x91, 0x58, 0xfe, 0x15, 0xd1, 0x2b, 0x47, 0x64, 0x97, 0xf2,
	0x86, 0xcf, 0x70, 0x06, 0x17, 0x75, 0x41, 0x63, 0x19, 0x49, 0x45, 0x61, 0x56, 0x90, 0x58, 0xfc,
	0x03, 0xcc, 0x67, 0xd7, 0x0c, 0x87, 0xd0, 0xcd, 0xf2, 0x27, 0x45, 0x4b, 0xd8, 0xce, 0x5c, 0xf6,
	0x37, 0xd5, 0x61, 0xe0, 0x2d, 0xb8, 0xa5, 0x12, 0xf1, 0x24, 0xb3, 0x6c, 0xb7, 0x5a, 0xdd, 0x35,
	0x3e, 0x78, 0x76, 0x65, 0x6c, 0x48, 0xc5, 0x62, 0x31, 0x4c, 0x83, 0x60, 0xcf, 0xfe, 0xdd, 0x9b,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0xe4, 0xf1, 0xb9, 0xda, 0x02, 0x00, 0x00,
}
