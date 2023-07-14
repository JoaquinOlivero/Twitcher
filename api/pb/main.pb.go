// Code generated by protoc-gen-go. DO NOT EDIT.
// source: main.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type Song struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Page                 string   `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	Author               string   `protobuf:"bytes,3,opt,name=author,proto3" json:"author,omitempty"`
	Audio                string   `protobuf:"bytes,4,opt,name=audio,proto3" json:"audio,omitempty"`
	Cover                string   `protobuf:"bytes,5,opt,name=cover,proto3" json:"cover,omitempty"`
	Bitrate              int32    `protobuf:"varint,6,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Song) Reset()         { *m = Song{} }
func (m *Song) String() string { return proto.CompactTextString(m) }
func (*Song) ProtoMessage()    {}
func (*Song) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{0}
}

func (m *Song) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Song.Unmarshal(m, b)
}
func (m *Song) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Song.Marshal(b, m, deterministic)
}
func (m *Song) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Song.Merge(m, src)
}
func (m *Song) XXX_Size() int {
	return xxx_messageInfo_Song.Size(m)
}
func (m *Song) XXX_DiscardUnknown() {
	xxx_messageInfo_Song.DiscardUnknown(m)
}

var xxx_messageInfo_Song proto.InternalMessageInfo

func (m *Song) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Song) GetPage() string {
	if m != nil {
		return m.Page
	}
	return ""
}

func (m *Song) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Song) GetAudio() string {
	if m != nil {
		return m.Audio
	}
	return ""
}

func (m *Song) GetCover() string {
	if m != nil {
		return m.Cover
	}
	return ""
}

func (m *Song) GetBitrate() int32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

type SongPlaylist struct {
	Songs                []*Song  `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SongPlaylist) Reset()         { *m = SongPlaylist{} }
func (m *SongPlaylist) String() string { return proto.CompactTextString(m) }
func (*SongPlaylist) ProtoMessage()    {}
func (*SongPlaylist) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{1}
}

func (m *SongPlaylist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SongPlaylist.Unmarshal(m, b)
}
func (m *SongPlaylist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SongPlaylist.Marshal(b, m, deterministic)
}
func (m *SongPlaylist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SongPlaylist.Merge(m, src)
}
func (m *SongPlaylist) XXX_Size() int {
	return xxx_messageInfo_SongPlaylist.Size(m)
}
func (m *SongPlaylist) XXX_DiscardUnknown() {
	xxx_messageInfo_SongPlaylist.DiscardUnknown(m)
}

var xxx_messageInfo_SongPlaylist proto.InternalMessageInfo

func (m *SongPlaylist) GetSongs() []*Song {
	if m != nil {
		return m.Songs
	}
	return nil
}

type AudioResponse struct {
	Ready                bool     `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AudioResponse) Reset()         { *m = AudioResponse{} }
func (m *AudioResponse) String() string { return proto.CompactTextString(m) }
func (*AudioResponse) ProtoMessage()    {}
func (*AudioResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{2}
}

func (m *AudioResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioResponse.Unmarshal(m, b)
}
func (m *AudioResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioResponse.Marshal(b, m, deterministic)
}
func (m *AudioResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioResponse.Merge(m, src)
}
func (m *AudioResponse) XXX_Size() int {
	return xxx_messageInfo_AudioResponse.Size(m)
}
func (m *AudioResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AudioResponse proto.InternalMessageInfo

func (m *AudioResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

type OutputRequest struct {
	Mode                 string   `protobuf:"bytes,1,opt,name=mode,proto3" json:"mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputRequest) Reset()         { *m = OutputRequest{} }
func (m *OutputRequest) String() string { return proto.CompactTextString(m) }
func (*OutputRequest) ProtoMessage()    {}
func (*OutputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{3}
}

func (m *OutputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputRequest.Unmarshal(m, b)
}
func (m *OutputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputRequest.Marshal(b, m, deterministic)
}
func (m *OutputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputRequest.Merge(m, src)
}
func (m *OutputRequest) XXX_Size() int {
	return xxx_messageInfo_OutputRequest.Size(m)
}
func (m *OutputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OutputRequest proto.InternalMessageInfo

func (m *OutputRequest) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

type OutputResponse struct {
	Ready                bool     `protobuf:"varint,1,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputResponse) Reset()         { *m = OutputResponse{} }
func (m *OutputResponse) String() string { return proto.CompactTextString(m) }
func (*OutputResponse) ProtoMessage()    {}
func (*OutputResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{4}
}

func (m *OutputResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputResponse.Unmarshal(m, b)
}
func (m *OutputResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputResponse.Marshal(b, m, deterministic)
}
func (m *OutputResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputResponse.Merge(m, src)
}
func (m *OutputResponse) XXX_Size() int {
	return xxx_messageInfo_OutputResponse.Size(m)
}
func (m *OutputResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OutputResponse proto.InternalMessageInfo

func (m *OutputResponse) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

type StatusResponse struct {
	Output               bool     `protobuf:"varint,1,opt,name=output,proto3" json:"output,omitempty"`
	Audio                bool     `protobuf:"varint,2,opt,name=audio,proto3" json:"audio,omitempty"`
	Stream               bool     `protobuf:"varint,3,opt,name=stream,proto3" json:"stream,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{5}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetOutput() bool {
	if m != nil {
		return m.Output
	}
	return false
}

func (m *StatusResponse) GetAudio() bool {
	if m != nil {
		return m.Audio
	}
	return false
}

func (m *StatusResponse) GetStream() bool {
	if m != nil {
		return m.Stream
	}
	return false
}

type SDP struct {
	Sdp                  string   `protobuf:"bytes,1,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SDP) Reset()         { *m = SDP{} }
func (m *SDP) String() string { return proto.CompactTextString(m) }
func (*SDP) ProtoMessage()    {}
func (*SDP) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{6}
}

func (m *SDP) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SDP.Unmarshal(m, b)
}
func (m *SDP) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SDP.Marshal(b, m, deterministic)
}
func (m *SDP) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SDP.Merge(m, src)
}
func (m *SDP) XXX_Size() int {
	return xxx_messageInfo_SDP.Size(m)
}
func (m *SDP) XXX_DiscardUnknown() {
	xxx_messageInfo_SDP.DiscardUnknown(m)
}

var xxx_messageInfo_SDP proto.InternalMessageInfo

func (m *SDP) GetSdp() string {
	if m != nil {
		return m.Sdp
	}
	return ""
}

type StatusNCSResponse struct {
	Active               bool     `protobuf:"varint,1,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusNCSResponse) Reset()         { *m = StatusNCSResponse{} }
func (m *StatusNCSResponse) String() string { return proto.CompactTextString(m) }
func (*StatusNCSResponse) ProtoMessage()    {}
func (*StatusNCSResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{7}
}

func (m *StatusNCSResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusNCSResponse.Unmarshal(m, b)
}
func (m *StatusNCSResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusNCSResponse.Marshal(b, m, deterministic)
}
func (m *StatusNCSResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusNCSResponse.Merge(m, src)
}
func (m *StatusNCSResponse) XXX_Size() int {
	return xxx_messageInfo_StatusNCSResponse.Size(m)
}
func (m *StatusNCSResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusNCSResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusNCSResponse proto.InternalMessageInfo

func (m *StatusNCSResponse) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type TwitchStreamKey struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Active               bool     `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TwitchStreamKey) Reset()         { *m = TwitchStreamKey{} }
func (m *TwitchStreamKey) String() string { return proto.CompactTextString(m) }
func (*TwitchStreamKey) ProtoMessage()    {}
func (*TwitchStreamKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{8}
}

func (m *TwitchStreamKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TwitchStreamKey.Unmarshal(m, b)
}
func (m *TwitchStreamKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TwitchStreamKey.Marshal(b, m, deterministic)
}
func (m *TwitchStreamKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TwitchStreamKey.Merge(m, src)
}
func (m *TwitchStreamKey) XXX_Size() int {
	return xxx_messageInfo_TwitchStreamKey.Size(m)
}
func (m *TwitchStreamKey) XXX_DiscardUnknown() {
	xxx_messageInfo_TwitchStreamKey.DiscardUnknown(m)
}

var xxx_messageInfo_TwitchStreamKey proto.InternalMessageInfo

func (m *TwitchStreamKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TwitchStreamKey) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type DevCredentials struct {
	ClientId             string   `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Secret               string   `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	Active               bool     `protobuf:"varint,3,opt,name=active,proto3" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DevCredentials) Reset()         { *m = DevCredentials{} }
func (m *DevCredentials) String() string { return proto.CompactTextString(m) }
func (*DevCredentials) ProtoMessage()    {}
func (*DevCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{9}
}

func (m *DevCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevCredentials.Unmarshal(m, b)
}
func (m *DevCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevCredentials.Marshal(b, m, deterministic)
}
func (m *DevCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevCredentials.Merge(m, src)
}
func (m *DevCredentials) XXX_Size() int {
	return xxx_messageInfo_DevCredentials.Size(m)
}
func (m *DevCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_DevCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_DevCredentials proto.InternalMessageInfo

func (m *DevCredentials) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *DevCredentials) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *DevCredentials) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type UserAuth struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAuth) Reset()         { *m = UserAuth{} }
func (m *UserAuth) String() string { return proto.CompactTextString(m) }
func (*UserAuth) ProtoMessage()    {}
func (*UserAuth) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ed94b0a22d11796, []int{10}
}

func (m *UserAuth) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuth.Unmarshal(m, b)
}
func (m *UserAuth) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuth.Marshal(b, m, deterministic)
}
func (m *UserAuth) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuth.Merge(m, src)
}
func (m *UserAuth) XXX_Size() int {
	return xxx_messageInfo_UserAuth.Size(m)
}
func (m *UserAuth) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuth.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuth proto.InternalMessageInfo

func (m *UserAuth) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*Song)(nil), "service.Song")
	proto.RegisterType((*SongPlaylist)(nil), "service.SongPlaylist")
	proto.RegisterType((*AudioResponse)(nil), "service.AudioResponse")
	proto.RegisterType((*OutputRequest)(nil), "service.OutputRequest")
	proto.RegisterType((*OutputResponse)(nil), "service.OutputResponse")
	proto.RegisterType((*StatusResponse)(nil), "service.StatusResponse")
	proto.RegisterType((*SDP)(nil), "service.SDP")
	proto.RegisterType((*StatusNCSResponse)(nil), "service.StatusNCSResponse")
	proto.RegisterType((*TwitchStreamKey)(nil), "service.TwitchStreamKey")
	proto.RegisterType((*DevCredentials)(nil), "service.DevCredentials")
	proto.RegisterType((*UserAuth)(nil), "service.UserAuth")
}

func init() {
	proto.RegisterFile("main.proto", fileDescriptor_7ed94b0a22d11796)
}

var fileDescriptor_7ed94b0a22d11796 = []byte{
	// 714 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x5d, 0x4f, 0xdb, 0x4a,
	0x10, 0x4d, 0xc8, 0x07, 0x61, 0xf8, 0xba, 0x2c, 0x10, 0x7c, 0x7d, 0xa5, 0x2b, 0xb4, 0xe8, 0xde,
	0x22, 0x55, 0x0a, 0x12, 0x3c, 0x22, 0xb5, 0xa4, 0x09, 0x48, 0xb4, 0x82, 0x46, 0x09, 0xf4, 0xa1,
	0xea, 0xcb, 0xc6, 0x9e, 0x26, 0x56, 0x12, 0xaf, 0xbb, 0xbb, 0x0e, 0xca, 0x7f, 0xe8, 0x8f, 0xe8,
	0x4f, 0xad, 0xd6, 0x6b, 0x3b, 0x71, 0x5a, 0x53, 0x5a, 0xde, 0x76, 0xc6, 0x33, 0x67, 0xce, 0xf1,
	0xce, 0x59, 0x80, 0x09, 0xf3, 0xfc, 0x46, 0x20, 0xb8, 0xe2, 0x64, 0x55, 0xa2, 0x98, 0x7a, 0x0e,
	0xda, 0xff, 0x0c, 0x38, 0x1f, 0x8c, 0xf1, 0x24, 0x4a, 0xf7, 0xc3, 0xcf, 0x27, 0x38, 0x09, 0xd4,
	0xcc, 0x54, 0xd1, 0xaf, 0x45, 0x28, 0xf7, 0xb8, 0x3f, 0x20, 0x04, 0xca, 0x3e, 0x9b, 0xa0, 0x55,
	0x3c, 0x2c, 0x1e, 0xaf, 0x75, 0xa3, 0xb3, 0xce, 0x05, 0x6c, 0x80, 0xd6, 0x8a, 0xc9, 0xe9, 0x33,
	0xa9, 0x43, 0x95, 0x85, 0x6a, 0xc8, 0x85, 0x55, 0x8a, 0xb2, 0x71, 0x44, 0xf6, 0xa0, 0xc2, 0x42,
	0xd7, 0xe3, 0x56, 0x39, 0x4a, 0x9b, 0x40, 0x67, 0x1d, 0x3e, 0x45, 0x61, 0x55, 0x4c, 0x36, 0x0a,
	0x88, 0x05, 0xab, 0x7d, 0x4f, 0x09, 0xa6, 0xd0, 0xaa, 0x1e, 0x16, 0x8f, 0x2b, 0xdd, 0x24, 0xa4,
	0x67, 0xb0, 0xa1, 0xd9, 0x74, 0xc6, 0x6c, 0x36, 0xf6, 0xa4, 0x22, 0x47, 0x50, 0x91, 0xdc, 0x1f,
	0x48, 0xab, 0x78, 0x58, 0x3a, 0x5e, 0x3f, 0xdd, 0x6c, 0xc4, 0xa2, 0x1a, 0xba, 0xaa, 0x6b, 0xbe,
	0xd1, 0xff, 0x60, 0xb3, 0xa9, 0xa7, 0x75, 0x51, 0x06, 0xdc, 0x97, 0xa8, 0xa7, 0x0a, 0x64, 0xee,
	0x2c, 0x12, 0x53, 0xeb, 0x9a, 0x80, 0x1e, 0xc1, 0xe6, 0xfb, 0x50, 0x05, 0xa1, 0xea, 0xe2, 0x97,
	0x10, 0xa5, 0xd2, 0xf2, 0x26, 0xdc, 0x4d, 0x25, 0xeb, 0x33, 0xfd, 0x1f, 0xb6, 0x92, 0xa2, 0x47,
	0xc1, 0x3e, 0xc0, 0x56, 0x4f, 0x31, 0x15, 0xca, 0xb4, 0xae, 0x0e, 0x55, 0x1e, 0x75, 0xc6, 0x85,
	0x71, 0x34, 0xff, 0x31, 0x2b, 0xa6, 0xdf, 0xfc, 0x98, 0x3a, 0x54, 0xa5, 0x12, 0xc8, 0x26, 0xd1,
	0x6f, 0xac, 0x75, 0xe3, 0x88, 0x1e, 0x40, 0xa9, 0xd7, 0xee, 0x90, 0xbf, 0xa0, 0x24, 0xdd, 0x20,
	0x66, 0xa6, 0x8f, 0xf4, 0x25, 0xec, 0x98, 0x81, 0xb7, 0xad, 0xde, 0xe2, 0x4c, 0xe6, 0x28, 0x6f,
	0x8a, 0xc9, 0x4c, 0x13, 0xd1, 0x73, 0xd8, 0xbe, 0x7b, 0xf0, 0x94, 0x33, 0xec, 0x45, 0xa8, 0xef,
	0x70, 0xa6, 0x11, 0x47, 0x38, 0x4b, 0x10, 0x47, 0x38, 0x5b, 0x68, 0x5e, 0xc9, 0x34, 0x7f, 0x82,
	0xad, 0x36, 0x4e, 0x5b, 0x02, 0x5d, 0xf4, 0x95, 0xc7, 0xc6, 0x92, 0xd8, 0x50, 0x73, 0xc6, 0x1e,
	0xfa, 0xea, 0xda, 0x8d, 0x01, 0xd2, 0x38, 0x12, 0x82, 0x8e, 0x40, 0x15, 0x6f, 0x49, 0x1c, 0x2d,
	0xa0, 0x97, 0x32, 0xe8, 0xff, 0x42, 0xed, 0x5e, 0xa2, 0x68, 0x86, 0x6a, 0xa8, 0x2f, 0xc0, 0x59,
	0xb8, 0x00, 0x7d, 0x3e, 0xfd, 0x06, 0x50, 0xbe, 0x61, 0x9e, 0x4f, 0x2e, 0x81, 0xb4, 0x04, 0x32,
	0x85, 0x99, 0x85, 0xa8, 0x37, 0xcc, 0x36, 0x37, 0x92, 0x6d, 0x6e, 0x5c, 0xea, 0x6d, 0xb6, 0xf7,
	0x33, 0x9b, 0x91, 0x94, 0xd3, 0x02, 0xb9, 0x82, 0xdd, 0x56, 0x28, 0x04, 0xfa, 0xea, 0x79, 0x38,
	0x97, 0x40, 0xee, 0x03, 0x77, 0x99, 0xce, 0xcf, 0xcb, 0xed, 0x1c, 0x74, 0x5a, 0x20, 0x2f, 0x60,
	0xb5, 0x23, 0x70, 0xea, 0xe1, 0x03, 0xd9, 0x98, 0xf7, 0xb6, 0x3b, 0x76, 0x26, 0xa2, 0x05, 0xf2,
	0x1a, 0xd6, 0x7b, 0x8a, 0x09, 0x65, 0x6e, 0x30, 0x97, 0x6f, 0xfe, 0xa4, 0x57, 0x00, 0x3d, 0xc5,
	0x83, 0xe7, 0xf5, 0x33, 0xa1, 0x9a, 0x66, 0x5f, 0x73, 0xfb, 0x13, 0xda, 0x19, 0x0b, 0xd2, 0x02,
	0xb9, 0x88, 0x05, 0x18, 0x3b, 0x91, 0x79, 0x61, 0xc6, 0x84, 0xf6, 0xc1, 0x0f, 0xf9, 0x14, 0x21,
	0x56, 0x90, 0x02, 0xfc, 0xae, 0x82, 0x73, 0xa8, 0x1a, 0xcb, 0xe4, 0xf6, 0xce, 0x87, 0x67, 0xcd,
	0x4c, 0x0b, 0xa4, 0x05, 0xdb, 0x57, 0x9e, 0xef, 0xde, 0xe2, 0x83, 0xbe, 0x59, 0xed, 0xba, 0x3f,
	0x60, 0xd0, 0x84, 0xb5, 0xd4, 0xb4, 0xb9, 0xed, 0xf6, 0x12, 0x89, 0x05, 0x83, 0xd3, 0x02, 0xb9,
	0x86, 0xdd, 0xd8, 0xca, 0x6c, 0x8a, 0x73, 0x3b, 0x5b, 0x69, 0xd3, 0x92, 0xd1, 0x1f, 0x61, 0xf3,
	0x16, 0xf6, 0x5a, 0x43, 0x74, 0x46, 0xcb, 0x4f, 0x43, 0x1e, 0xb1, 0xdc, 0x19, 0x11, 0xad, 0xfd,
	0x36, 0x8e, 0x51, 0xe1, 0x53, 0xc1, 0xf2, 0x69, 0xdd, 0x80, 0xa5, 0xb5, 0x19, 0xa0, 0xa5, 0x97,
	0x67, 0x7e, 0x41, 0xd9, 0x0f, 0x8f, 0xc0, 0xdd, 0xc2, 0xdf, 0x0b, 0x2a, 0x97, 0xf0, 0x7e, 0xbd,
	0x08, 0xd9, 0x86, 0x08, 0xcf, 0x5e, 0x54, 0xfa, 0x44, 0xc0, 0x7c, 0x7e, 0x17, 0xb0, 0x63, 0x90,
	0x9a, 0x8e, 0x83, 0x52, 0xde, 0xf1, 0x11, 0xfa, 0x64, 0x27, 0x9d, 0x9f, 0x3c, 0x8e, 0xf9, 0x08,
	0x6f, 0xaa, 0x1f, 0xcb, 0x8d, 0x93, 0xa0, 0xdf, 0xaf, 0x46, 0x5f, 0xce, 0xbe, 0x07, 0x00, 0x00,
	0xff, 0xff, 0x7a, 0xb1, 0x44, 0x5c, 0xf6, 0x07, 0x00, 0x00,
}
