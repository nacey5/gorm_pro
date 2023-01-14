// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tag.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetTagListRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                uint32   `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTagListRequest) Reset()         { *m = GetTagListRequest{} }
func (m *GetTagListRequest) String() string { return proto.CompactTextString(m) }
func (*GetTagListRequest) ProtoMessage()    {}
func (*GetTagListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f545bcde37ecb5, []int{0}
}

func (m *GetTagListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTagListRequest.Unmarshal(m, b)
}
func (m *GetTagListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTagListRequest.Marshal(b, m, deterministic)
}
func (m *GetTagListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTagListRequest.Merge(m, src)
}
func (m *GetTagListRequest) XXX_Size() int {
	return xxx_messageInfo_GetTagListRequest.Size(m)
}
func (m *GetTagListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTagListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTagListRequest proto.InternalMessageInfo

func (m *GetTagListRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetTagListRequest) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type Tag struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	State                uint32   `protobuf:"varint,3,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f545bcde37ecb5, []int{1}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetState() uint32 {
	if m != nil {
		return m.State
	}
	return 0
}

type GetTagListReply struct {
	List                 []*Tag   `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Pager                *Pager   `protobuf:"bytes,2,opt,name=pager,proto3" json:"pager,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTagListReply) Reset()         { *m = GetTagListReply{} }
func (m *GetTagListReply) String() string { return proto.CompactTextString(m) }
func (*GetTagListReply) ProtoMessage()    {}
func (*GetTagListReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_27f545bcde37ecb5, []int{2}
}

func (m *GetTagListReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTagListReply.Unmarshal(m, b)
}
func (m *GetTagListReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTagListReply.Marshal(b, m, deterministic)
}
func (m *GetTagListReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTagListReply.Merge(m, src)
}
func (m *GetTagListReply) XXX_Size() int {
	return xxx_messageInfo_GetTagListReply.Size(m)
}
func (m *GetTagListReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTagListReply.DiscardUnknown(m)
}

var xxx_messageInfo_GetTagListReply proto.InternalMessageInfo

func (m *GetTagListReply) GetList() []*Tag {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *GetTagListReply) GetPager() *Pager {
	if m != nil {
		return m.Pager
	}
	return nil
}

func init() {
	proto.RegisterType((*GetTagListRequest)(nil), "proto.GetTagListRequest")
	proto.RegisterType((*Tag)(nil), "proto.Tag")
	proto.RegisterType((*GetTagListReply)(nil), "proto.GetTagListReply")
}

func init() { proto.RegisterFile("tag.proto", fileDescriptor_27f545bcde37ecb5) }

var fileDescriptor_27f545bcde37ecb5 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x6d, 0xbb, 0x15, 0x76, 0x76, 0x55, 0x1c, 0x44, 0xca, 0x1e, 0xa4, 0xf4, 0xd4, 0xd3,
	0x1e, 0xea, 0x59, 0xf4, 0xe6, 0x45, 0x44, 0x62, 0xfc, 0x01, 0xe3, 0xee, 0x10, 0x02, 0xed, 0x26,
	0x36, 0xa3, 0xb0, 0xff, 0x5e, 0x4c, 0x16, 0x5d, 0xb0, 0xa7, 0x24, 0xef, 0xbd, 0x7c, 0x6f, 0x06,
	0xe6, 0x42, 0x66, 0xed, 0x47, 0x27, 0x0e, 0xcb, 0x78, 0xac, 0x96, 0x1b, 0x37, 0x0c, 0x6e, 0x97,
	0xc4, 0xe6, 0x0e, 0x2e, 0x1f, 0x59, 0x34, 0x99, 0x27, 0x1b, 0x44, 0xf1, 0xc7, 0x27, 0x07, 0x41,
	0x84, 0xd9, 0x8e, 0x06, 0xae, 0xb2, 0x3a, 0x6b, 0xe7, 0x2a, 0xde, 0xf1, 0x0a, 0xca, 0x20, 0x24,
	0x5c, 0xe5, 0x75, 0xd6, 0x9e, 0xa9, 0xf4, 0x68, 0xee, 0xa1, 0xd0, 0x64, 0xf0, 0x1c, 0x72, 0xbb,
	0x8d, 0xf1, 0x42, 0xe5, 0x76, 0xfb, 0x0b, 0xc8, 0xa7, 0x00, 0xc5, 0x31, 0xe0, 0x0d, 0x2e, 0x8e,
	0xfb, 0x7d, 0xbf, 0xc7, 0x1b, 0x98, 0xf5, 0x36, 0x48, 0x95, 0xd5, 0x45, 0xbb, 0xe8, 0x20, 0x0d,
	0xba, 0xd6, 0x64, 0x54, 0xd4, 0xb1, 0x81, 0xd2, 0x93, 0xe1, 0x31, 0xd2, 0x17, 0xdd, 0xf2, 0x10,
	0x78, 0xf9, 0xd1, 0x54, 0xb2, 0xba, 0x67, 0x00, 0x4d, 0xe6, 0x95, 0xc7, 0x2f, 0xbb, 0x61, 0x7c,
	0x00, 0xf8, 0x2b, 0xc1, 0xea, 0xf0, 0xe1, 0xdf, 0xde, 0xab, 0xeb, 0x09, 0xc7, 0xf7, 0xfb, 0xe6,
	0xe4, 0xfd, 0x34, 0x1a, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x34, 0x34, 0xe2, 0x4f,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TagServiceClient is the client API for TagService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TagServiceClient interface {
	GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error)
}

type tagServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTagServiceClient(cc grpc.ClientConnInterface) TagServiceClient {
	return &tagServiceClient{cc}
}

func (c *tagServiceClient) GetTagList(ctx context.Context, in *GetTagListRequest, opts ...grpc.CallOption) (*GetTagListReply, error) {
	out := new(GetTagListReply)
	err := c.cc.Invoke(ctx, "/proto.TagService/GetTagList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TagServiceServer is the server API for TagService service.
type TagServiceServer interface {
	GetTagList(context.Context, *GetTagListRequest) (*GetTagListReply, error)
}

// UnimplementedTagServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTagServiceServer struct {
}

func (*UnimplementedTagServiceServer) GetTagList(ctx context.Context, req *GetTagListRequest) (*GetTagListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTagList not implemented")
}

func RegisterTagServiceServer(s *grpc.Server, srv TagServiceServer) {
	s.RegisterService(&_TagService_serviceDesc, srv)
}

func _TagService_GetTagList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTagListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TagServiceServer).GetTagList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TagService/GetTagList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TagServiceServer).GetTagList(ctx, req.(*GetTagListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TagService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TagService",
	HandlerType: (*TagServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTagList",
			Handler:    _TagService_GetTagList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tag.proto",
}
