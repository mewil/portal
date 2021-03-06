// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: file.proto

package pb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}

var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}

func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}

type FilePart struct {
	Content              []byte   `protobuf:"bytes,1,opt,name=Content,proto3" json:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilePart) Reset()         { *m = FilePart{} }
func (m *FilePart) String() string { return proto.CompactTextString(m) }
func (*FilePart) ProtoMessage()    {}
func (*FilePart) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{0}
}
func (m *FilePart) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilePart) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilePart.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilePart) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilePart.Merge(m, src)
}
func (m *FilePart) XXX_Size() int {
	return m.Size()
}
func (m *FilePart) XXX_DiscardUnknown() {
	xxx_messageInfo_FilePart.DiscardUnknown(m)
}

var xxx_messageInfo_FilePart proto.InternalMessageInfo

func (m *FilePart) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type FileRequest struct {
	FileID               string   `protobuf:"bytes,1,opt,name=FileID,proto3" json:"file_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileRequest) Reset()         { *m = FileRequest{} }
func (m *FileRequest) String() string { return proto.CompactTextString(m) }
func (*FileRequest) ProtoMessage()    {}
func (*FileRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{1}
}
func (m *FileRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileRequest.Merge(m, src)
}
func (m *FileRequest) XXX_Size() int {
	return m.Size()
}
func (m *FileRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FileRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FileRequest proto.InternalMessageInfo

func (m *FileRequest) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

type FileStats struct {
	ETag                 string   `protobuf:"bytes,1,opt,name=ETag,proto3" json:"etag"`
	FileSize             int64    `protobuf:"varint,2,opt,name=FileSize,proto3" json:"size"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileStats) Reset()         { *m = FileStats{} }
func (m *FileStats) String() string { return proto.CompactTextString(m) }
func (*FileStats) ProtoMessage()    {}
func (*FileStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{2}
}
func (m *FileStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FileStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FileStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FileStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileStats.Merge(m, src)
}
func (m *FileStats) XXX_Size() int {
	return m.Size()
}
func (m *FileStats) XXX_DiscardUnknown() {
	xxx_messageInfo_FileStats.DiscardUnknown(m)
}

var xxx_messageInfo_FileStats proto.InternalMessageInfo

func (m *FileStats) GetETag() string {
	if m != nil {
		return m.ETag
	}
	return ""
}

func (m *FileStats) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

type UploadStatus struct {
	FileID               string           `protobuf:"bytes,1,opt,name=FileID,proto3" json:"file_id"`
	Status               UploadStatusCode `protobuf:"varint,2,opt,name=Status,proto3,enum=pb.UploadStatusCode" json:"status"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadStatus) Reset()         { *m = UploadStatus{} }
func (m *UploadStatus) String() string { return proto.CompactTextString(m) }
func (*UploadStatus) ProtoMessage()    {}
func (*UploadStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_9188e3b7e55e1162, []int{3}
}
func (m *UploadStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UploadStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UploadStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UploadStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadStatus.Merge(m, src)
}
func (m *UploadStatus) XXX_Size() int {
	return m.Size()
}
func (m *UploadStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadStatus.DiscardUnknown(m)
}

var xxx_messageInfo_UploadStatus proto.InternalMessageInfo

func (m *UploadStatus) GetFileID() string {
	if m != nil {
		return m.FileID
	}
	return ""
}

func (m *UploadStatus) GetStatus() UploadStatusCode {
	if m != nil {
		return m.Status
	}
	return UploadStatusCode_Unknown
}

func init() {
	proto.RegisterEnum("pb.UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
	proto.RegisterType((*FilePart)(nil), "pb.FilePart")
	proto.RegisterType((*FileRequest)(nil), "pb.FileRequest")
	proto.RegisterType((*FileStats)(nil), "pb.FileStats")
	proto.RegisterType((*UploadStatus)(nil), "pb.UploadStatus")
}

func init() { proto.RegisterFile("file.proto", fileDescriptor_9188e3b7e55e1162) }

var fileDescriptor_9188e3b7e55e1162 = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x4f, 0xe2, 0x40,
	0x18, 0x86, 0x3b, 0x5d, 0xd2, 0xc2, 0x47, 0x77, 0xb7, 0x99, 0xec, 0x81, 0x90, 0x4d, 0x21, 0xdd,
	0x35, 0x21, 0x06, 0x0b, 0xc2, 0xc5, 0x33, 0x28, 0xc6, 0x13, 0xa6, 0xc8, 0xd9, 0xb4, 0x74, 0xa8,
	0x0d, 0xa5, 0x53, 0xe9, 0x54, 0x13, 0xfe, 0x88, 0xfe, 0x24, 0x8f, 0xfe, 0x02, 0x62, 0xf0, 0xc6,
	0xaf, 0x30, 0x9d, 0x69, 0x8d, 0xe2, 0xc5, 0xdb, 0x7c, 0xdf, 0xfb, 0xbc, 0x6f, 0xdf, 0x7c, 0x29,
	0xc0, 0x3c, 0x08, 0x89, 0x15, 0xaf, 0x28, 0xa3, 0x58, 0x8e, 0xdd, 0xfa, 0x91, 0x1f, 0xb0, 0x9b,
	0xd4, 0xb5, 0x66, 0x74, 0xd9, 0xf1, 0xa9, 0x4f, 0x3b, 0x5c, 0x72, 0xd3, 0x39, 0x9f, 0xf8, 0xc0,
	0x5f, 0xc2, 0x62, 0x1e, 0x43, 0x79, 0x14, 0x84, 0xe4, 0xd2, 0x59, 0x31, 0x7c, 0x00, 0xea, 0x90,
	0x46, 0x8c, 0x44, 0xac, 0x86, 0x9a, 0xa8, 0xa5, 0x0d, 0xaa, 0xbb, 0x4d, 0x43, 0x9d, 0x89, 0x95,
	0x5d, 0x68, 0x66, 0x0f, 0xaa, 0x99, 0xc5, 0x26, 0xb7, 0x29, 0x49, 0x18, 0xfe, 0x07, 0x4a, 0x36,
	0x5e, 0x9c, 0x72, 0x53, 0x45, 0x98, 0xb2, 0x52, 0xd7, 0x81, 0x67, 0xe7, 0x92, 0x39, 0x86, 0x4a,
	0xf6, 0x9a, 0x30, 0x87, 0x25, 0xf8, 0x2f, 0x94, 0xce, 0xae, 0x1c, 0x3f, 0xe7, 0xcb, 0xbb, 0x4d,
	0xa3, 0x44, 0x98, 0xe3, 0xdb, 0x7c, 0x8b, 0xff, 0x8b, 0x46, 0x93, 0x60, 0x4d, 0x6a, 0x72, 0x13,
	0xb5, 0x7e, 0x08, 0x22, 0x09, 0xd6, 0xc4, 0x7e, 0x57, 0xcc, 0x25, 0x68, 0xd3, 0x38, 0xa4, 0x8e,
	0x97, 0x45, 0xa6, 0xc9, 0xb7, 0x5a, 0xe0, 0x13, 0x50, 0x04, 0xce, 0x83, 0x7f, 0xf5, 0xfe, 0x58,
	0xb1, 0x6b, 0x7d, 0x8c, 0x19, 0x52, 0x8f, 0x0c, 0x60, 0xb7, 0x69, 0x28, 0x09, 0x9f, 0xed, 0x9c,
	0x3f, 0xec, 0x83, 0xbe, 0xcf, 0xe1, 0x2a, 0xa8, 0xd3, 0x68, 0x11, 0xd1, 0xfb, 0x48, 0x97, 0xb0,
	0x02, 0xf2, 0x78, 0xa1, 0x23, 0x0c, 0xa0, 0x8c, 0x9c, 0x20, 0x24, 0x9e, 0x2e, 0xf7, 0x1e, 0x10,
	0xa8, 0xbc, 0xf0, 0xdd, 0x0c, 0xb7, 0x41, 0x11, 0x01, 0x58, 0xcb, 0x3e, 0x5a, 0xdc, 0xbc, 0xae,
	0xef, 0x57, 0x30, 0xa5, 0x16, 0xc2, 0x6d, 0x50, 0xcf, 0x09, 0xcb, 0x20, 0xfc, 0xbb, 0xc0, 0xf3,
	0x7b, 0xd7, 0x3f, 0xf9, 0x4d, 0xa9, 0x8b, 0x70, 0x17, 0xb4, 0x9c, 0x16, 0xf7, 0xfd, 0x62, 0xf9,
	0x59, 0x2c, 0xb8, 0x6e, 0x4a, 0x03, 0xfd, 0x69, 0x6b, 0xa0, 0xe7, 0xad, 0x81, 0x5e, 0xb6, 0x06,
	0x7a, 0x7c, 0x35, 0x24, 0x57, 0xe1, 0xbf, 0x43, 0xff, 0x2d, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x91,
	0x23, 0x3a, 0x4f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FileSvcClient is the client API for FileSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FileSvcClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (FileSvc_UploadClient, error)
	GetFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (FileSvc_GetFileClient, error)
	GetFileStats(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileStats, error)
}

type fileSvcClient struct {
	cc *grpc.ClientConn
}

func NewFileSvcClient(cc *grpc.ClientConn) FileSvcClient {
	return &fileSvcClient{cc}
}

func (c *fileSvcClient) Upload(ctx context.Context, opts ...grpc.CallOption) (FileSvc_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileSvc_serviceDesc.Streams[0], "/pb.FileSvc/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSvcUploadClient{stream}
	return x, nil
}

type FileSvc_UploadClient interface {
	Send(*FilePart) error
	CloseAndRecv() (*UploadStatus, error)
	grpc.ClientStream
}

type fileSvcUploadClient struct {
	grpc.ClientStream
}

func (x *fileSvcUploadClient) Send(m *FilePart) error {
	return x.ClientStream.SendMsg(m)
}

func (x *fileSvcUploadClient) CloseAndRecv() (*UploadStatus, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadStatus)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSvcClient) GetFile(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (FileSvc_GetFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &_FileSvc_serviceDesc.Streams[1], "/pb.FileSvc/GetFile", opts...)
	if err != nil {
		return nil, err
	}
	x := &fileSvcGetFileClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FileSvc_GetFileClient interface {
	Recv() (*FilePart, error)
	grpc.ClientStream
}

type fileSvcGetFileClient struct {
	grpc.ClientStream
}

func (x *fileSvcGetFileClient) Recv() (*FilePart, error) {
	m := new(FilePart)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fileSvcClient) GetFileStats(ctx context.Context, in *FileRequest, opts ...grpc.CallOption) (*FileStats, error) {
	out := new(FileStats)
	err := c.cc.Invoke(ctx, "/pb.FileSvc/GetFileStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileSvcServer is the server API for FileSvc service.
type FileSvcServer interface {
	Upload(FileSvc_UploadServer) error
	GetFile(*FileRequest, FileSvc_GetFileServer) error
	GetFileStats(context.Context, *FileRequest) (*FileStats, error)
}

// UnimplementedFileSvcServer can be embedded to have forward compatible implementations.
type UnimplementedFileSvcServer struct {
}

func (*UnimplementedFileSvcServer) Upload(srv FileSvc_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (*UnimplementedFileSvcServer) GetFile(req *FileRequest, srv FileSvc_GetFileServer) error {
	return status.Errorf(codes.Unimplemented, "method GetFile not implemented")
}
func (*UnimplementedFileSvcServer) GetFileStats(ctx context.Context, req *FileRequest) (*FileStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileStats not implemented")
}

func RegisterFileSvcServer(s *grpc.Server, srv FileSvcServer) {
	s.RegisterService(&_FileSvc_serviceDesc, srv)
}

func _FileSvc_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FileSvcServer).Upload(&fileSvcUploadServer{stream})
}

type FileSvc_UploadServer interface {
	SendAndClose(*UploadStatus) error
	Recv() (*FilePart, error)
	grpc.ServerStream
}

type fileSvcUploadServer struct {
	grpc.ServerStream
}

func (x *fileSvcUploadServer) SendAndClose(m *UploadStatus) error {
	return x.ServerStream.SendMsg(m)
}

func (x *fileSvcUploadServer) Recv() (*FilePart, error) {
	m := new(FilePart)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FileSvc_GetFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FileRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FileSvcServer).GetFile(m, &fileSvcGetFileServer{stream})
}

type FileSvc_GetFileServer interface {
	Send(*FilePart) error
	grpc.ServerStream
}

type fileSvcGetFileServer struct {
	grpc.ServerStream
}

func (x *fileSvcGetFileServer) Send(m *FilePart) error {
	return x.ServerStream.SendMsg(m)
}

func _FileSvc_GetFileStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileSvcServer).GetFileStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.FileSvc/GetFileStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileSvcServer).GetFileStats(ctx, req.(*FileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FileSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.FileSvc",
	HandlerType: (*FileSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFileStats",
			Handler:    _FileSvc_GetFileStats_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _FileSvc_Upload_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetFile",
			Handler:       _FileSvc_GetFile_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "file.proto",
}

func (m *FilePart) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilePart) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FilePart) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintFile(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FileRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FileID) > 0 {
		i -= len(m.FileID)
		copy(dAtA[i:], m.FileID)
		i = encodeVarintFile(dAtA, i, uint64(len(m.FileID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FileStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FileStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FileStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.FileSize != 0 {
		i = encodeVarintFile(dAtA, i, uint64(m.FileSize))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ETag) > 0 {
		i -= len(m.ETag)
		copy(dAtA[i:], m.ETag)
		i = encodeVarintFile(dAtA, i, uint64(len(m.ETag)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UploadStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UploadStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UploadStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintFile(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x10
	}
	if len(m.FileID) > 0 {
		i -= len(m.FileID)
		copy(dAtA[i:], m.FileID)
		i = encodeVarintFile(dAtA, i, uint64(len(m.FileID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFile(dAtA []byte, offset int, v uint64) int {
	offset -= sovFile(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FilePart) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FileRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FileID)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *FileStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ETag)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	if m.FileSize != 0 {
		n += 1 + sovFile(uint64(m.FileSize))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *UploadStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FileID)
	if l > 0 {
		n += 1 + l + sovFile(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovFile(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFile(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFile(x uint64) (n int) {
	return sovFile(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilePart) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFile
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
			return fmt.Errorf("proto: FilePart: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilePart: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = append(m.Content[:0], dAtA[iNdEx:postIndex]...)
			if m.Content == nil {
				m.Content = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFile
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
			return fmt.Errorf("proto: FileRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FileStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFile
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
			return fmt.Errorf("proto: FileStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FileStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ETag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ETag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileSize", wireType)
			}
			m.FileSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FileSize |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *UploadStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFile
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
			return fmt.Errorf("proto: UploadStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UploadStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
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
				return ErrInvalidLengthFile
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFile
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFile
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= UploadStatusCode(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipFile(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFile
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFile(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFile
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
					return 0, ErrIntOverflowFile
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
					return 0, ErrIntOverflowFile
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
				return 0, ErrInvalidLengthFile
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFile
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFile
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFile        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFile          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFile = fmt.Errorf("proto: unexpected end of group")
)
