// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iot.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DeviceStatus struct {
	Date                 string   `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Unixtime             uint32   `protobuf:"varint,2,opt,name=unixtime,proto3" json:"unixtime,omitempty"`
	Satusgps             uint32   `protobuf:"varint,3,opt,name=satusgps,proto3" json:"satusgps,omitempty"`
	Latitude             float32  `protobuf:"fixed32,4,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float32  `protobuf:"fixed32,5,opt,name=longitude,proto3" json:"longitude,omitempty"`
	GPShdop              float32  `protobuf:"fixed32,6,opt,name=GPShdop,proto3" json:"GPShdop,omitempty"`
	Velocity             uint32   `protobuf:"varint,7,opt,name=velocity,proto3" json:"velocity,omitempty"`
	GPSsat               uint32   `protobuf:"varint,8,opt,name=GPSsat,proto3" json:"GPSsat,omitempty"`
	FixQuality           uint32   `protobuf:"varint,9,opt,name=fixQuality,proto3" json:"fixQuality,omitempty"`
	Temp                 float32  `protobuf:"fixed32,10,opt,name=temp,proto3" json:"temp,omitempty"`
	LuckyNumber          int32    `protobuf:"varint,11,opt,name=lucky_number,json=luckyNumber,proto3" json:"lucky_number,omitempty"`
	Ms                   string   `protobuf:"bytes,12,opt,name=ms,proto3" json:"ms,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceStatus) Reset()         { *m = DeviceStatus{} }
func (m *DeviceStatus) String() string { return proto.CompactTextString(m) }
func (*DeviceStatus) ProtoMessage()    {}
func (*DeviceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{0}
}

func (m *DeviceStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceStatus.Unmarshal(m, b)
}
func (m *DeviceStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceStatus.Marshal(b, m, deterministic)
}
func (m *DeviceStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceStatus.Merge(m, src)
}
func (m *DeviceStatus) XXX_Size() int {
	return xxx_messageInfo_DeviceStatus.Size(m)
}
func (m *DeviceStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceStatus proto.InternalMessageInfo

func (m *DeviceStatus) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *DeviceStatus) GetUnixtime() uint32 {
	if m != nil {
		return m.Unixtime
	}
	return 0
}

func (m *DeviceStatus) GetSatusgps() uint32 {
	if m != nil {
		return m.Satusgps
	}
	return 0
}

func (m *DeviceStatus) GetLatitude() float32 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *DeviceStatus) GetLongitude() float32 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *DeviceStatus) GetGPShdop() float32 {
	if m != nil {
		return m.GPShdop
	}
	return 0
}

func (m *DeviceStatus) GetVelocity() uint32 {
	if m != nil {
		return m.Velocity
	}
	return 0
}

func (m *DeviceStatus) GetGPSsat() uint32 {
	if m != nil {
		return m.GPSsat
	}
	return 0
}

func (m *DeviceStatus) GetFixQuality() uint32 {
	if m != nil {
		return m.FixQuality
	}
	return 0
}

func (m *DeviceStatus) GetTemp() float32 {
	if m != nil {
		return m.Temp
	}
	return 0
}

func (m *DeviceStatus) GetLuckyNumber() int32 {
	if m != nil {
		return m.LuckyNumber
	}
	return 0
}

func (m *DeviceStatus) GetMs() string {
	if m != nil {
		return m.Ms
	}
	return ""
}

// The response message containing the greetings
type DeviceStatusResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Time                 string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceStatusResponse) Reset()         { *m = DeviceStatusResponse{} }
func (m *DeviceStatusResponse) String() string { return proto.CompactTextString(m) }
func (*DeviceStatusResponse) ProtoMessage()    {}
func (*DeviceStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1804728d01c3c43d, []int{1}
}

func (m *DeviceStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceStatusResponse.Unmarshal(m, b)
}
func (m *DeviceStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceStatusResponse.Marshal(b, m, deterministic)
}
func (m *DeviceStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceStatusResponse.Merge(m, src)
}
func (m *DeviceStatusResponse) XXX_Size() int {
	return xxx_messageInfo_DeviceStatusResponse.Size(m)
}
func (m *DeviceStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceStatusResponse proto.InternalMessageInfo

func (m *DeviceStatusResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *DeviceStatusResponse) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceStatus)(nil), "proto.DeviceStatus")
	proto.RegisterType((*DeviceStatusResponse)(nil), "proto.DeviceStatusResponse")
}

func init() { proto.RegisterFile("iot.proto", fileDescriptor_1804728d01c3c43d) }

var fileDescriptor_1804728d01c3c43d = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcf, 0x6e, 0xf2, 0x30,
	0x10, 0xc4, 0xbf, 0xe4, 0xe3, 0x5f, 0x16, 0xe8, 0xc1, 0xad, 0x2a, 0x8b, 0x56, 0x55, 0xca, 0x29,
	0x27, 0x0e, 0xed, 0xb9, 0x97, 0x0a, 0x89, 0x1b, 0xa2, 0xe6, 0x01, 0x2a, 0x03, 0x5b, 0x6a, 0x35,
	0x89, 0x23, 0xbc, 0x46, 0xf0, 0x88, 0x7d, 0xab, 0xca, 0x1b, 0x42, 0x39, 0xf4, 0xe4, 0x9d, 0xf9,
	0xc9, 0x3b, 0xc9, 0x18, 0x12, 0x63, 0x69, 0x52, 0xed, 0x2c, 0x59, 0xd1, 0xe6, 0x63, 0xfc, 0x1d,
	0xc3, 0x60, 0x8a, 0x7b, 0xb3, 0xc6, 0x25, 0x69, 0xf2, 0x4e, 0x08, 0x68, 0x6d, 0x34, 0xa1, 0x8c,
	0xd2, 0x28, 0x4b, 0x14, 0xcf, 0x62, 0x04, 0x3d, 0x5f, 0x9a, 0x03, 0x99, 0x02, 0x65, 0x9c, 0x46,
	0xd9, 0x50, 0x9d, 0x75, 0x60, 0x2e, 0x5c, 0xdc, 0x56, 0x4e, 0xfe, 0xaf, 0x59, 0xa3, 0x03, 0xcb,
	0x35, 0x19, 0xf2, 0x1b, 0x94, 0xad, 0x34, 0xca, 0x62, 0x75, 0xd6, 0xe2, 0x1e, 0x92, 0xdc, 0x96,
	0xdb, 0x1a, 0xb6, 0x19, 0xfe, 0x1a, 0x42, 0x42, 0x77, 0xb6, 0x58, 0x7e, 0x6e, 0x6c, 0x25, 0x3b,
	0xcc, 0x1a, 0x19, 0x76, 0xee, 0x31, 0xb7, 0x6b, 0x43, 0x47, 0xd9, 0xad, 0xf3, 0x1a, 0x2d, 0x6e,
	0xa1, 0x33, 0x5b, 0x2c, 0x9d, 0x26, 0xd9, 0x63, 0x72, 0x52, 0xe2, 0x01, 0xe0, 0xc3, 0x1c, 0xde,
	0xbc, 0xce, 0xc3, 0xad, 0x84, 0xd9, 0x85, 0x13, 0xfe, 0x99, 0xb0, 0xa8, 0x24, 0x70, 0x14, 0xcf,
	0xe2, 0x11, 0x06, 0xb9, 0x5f, 0x7f, 0x1d, 0xdf, 0x4b, 0x5f, 0xac, 0x70, 0x27, 0xfb, 0x69, 0x94,
	0xb5, 0x55, 0x9f, 0xbd, 0x39, 0x5b, 0xe2, 0x0a, 0xe2, 0xc2, 0xc9, 0x01, 0x17, 0x15, 0x17, 0x6e,
	0xfc, 0x0a, 0x37, 0x97, 0x55, 0x2a, 0x74, 0x95, 0x2d, 0x1d, 0x86, 0xcf, 0x72, 0xec, 0x70, 0xa9,
	0x3d, 0x75, 0x52, 0x1c, 0xdb, 0x54, 0x9a, 0x28, 0x9e, 0x9f, 0xe6, 0x30, 0x3c, 0xed, 0xc0, 0x5d,
	0x38, 0xc4, 0x0b, 0x74, 0x67, 0x48, 0x53, 0x4d, 0x5a, 0x5c, 0xd7, 0x4f, 0x37, 0xb9, 0x0c, 0x19,
	0xdd, 0xfd, 0x61, 0x36, 0xc9, 0xe3, 0x7f, 0xab, 0x0e, 0xd3, 0xe7, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x59, 0x1d, 0x3a, 0x68, 0xfa, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceServiceClient interface {
	// Sends data devices
	GetData(ctx context.Context, in *DeviceStatus, opts ...grpc.CallOption) (*DeviceStatusResponse, error)
}

type deviceServiceClient struct {
	cc *grpc.ClientConn
}

func NewDeviceServiceClient(cc *grpc.ClientConn) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) GetData(ctx context.Context, in *DeviceStatus, opts ...grpc.CallOption) (*DeviceStatusResponse, error) {
	out := new(DeviceStatusResponse)
	err := c.cc.Invoke(ctx, "/proto.DeviceService/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
type DeviceServiceServer interface {
	// Sends data devices
	GetData(context.Context, *DeviceStatus) (*DeviceStatusResponse, error)
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DeviceService/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetData(ctx, req.(*DeviceStatus))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetData",
			Handler:    _DeviceService_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iot.proto",
}
