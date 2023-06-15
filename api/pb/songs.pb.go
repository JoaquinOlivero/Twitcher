// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/songs.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	return fileDescriptor_23af63290edd64ad, []int{0}
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
	return fileDescriptor_23af63290edd64ad, []int{1}
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

type AudioStream struct {
	Playlist             *SongPlaylist `protobuf:"bytes,1,opt,name=playlist,proto3" json:"playlist,omitempty"`
	Ready                bool          `protobuf:"varint,2,opt,name=ready,proto3" json:"ready,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AudioStream) Reset()         { *m = AudioStream{} }
func (m *AudioStream) String() string { return proto.CompactTextString(m) }
func (*AudioStream) ProtoMessage()    {}
func (*AudioStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_23af63290edd64ad, []int{2}
}

func (m *AudioStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioStream.Unmarshal(m, b)
}
func (m *AudioStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioStream.Marshal(b, m, deterministic)
}
func (m *AudioStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioStream.Merge(m, src)
}
func (m *AudioStream) XXX_Size() int {
	return xxx_messageInfo_AudioStream.Size(m)
}
func (m *AudioStream) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioStream.DiscardUnknown(m)
}

var xxx_messageInfo_AudioStream proto.InternalMessageInfo

func (m *AudioStream) GetPlaylist() *SongPlaylist {
	if m != nil {
		return m.Playlist
	}
	return nil
}

func (m *AudioStream) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
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
	return fileDescriptor_23af63290edd64ad, []int{3}
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_23af63290edd64ad, []int{4}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

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
	return fileDescriptor_23af63290edd64ad, []int{5}
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

func init() {
	proto.RegisterType((*Song)(nil), "service.Song")
	proto.RegisterType((*SongPlaylist)(nil), "service.SongPlaylist")
	proto.RegisterType((*AudioStream)(nil), "service.AudioStream")
	proto.RegisterType((*OutputResponse)(nil), "service.OutputResponse")
	proto.RegisterType((*Empty)(nil), "service.Empty")
	proto.RegisterType((*SDP)(nil), "service.SDP")
}

func init() { proto.RegisterFile("proto/songs.proto", fileDescriptor_23af63290edd64ad) }

var fileDescriptor_23af63290edd64ad = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x89, 0xed, 0x94, 0x49, 0xa8, 0xca, 0x50, 0xa8, 0x95, 0x53, 0xb4, 0x48, 0x28, 0x5c,
	0x9c, 0x36, 0x15, 0x07, 0x0e, 0x45, 0x82, 0x96, 0x23, 0x22, 0xb2, 0x81, 0x03, 0xb7, 0x4d, 0x32,
	0x4a, 0x2d, 0x35, 0xde, 0xd5, 0x7a, 0x9c, 0x2a, 0x47, 0xee, 0x7c, 0x34, 0xda, 0xdd, 0x34, 0x6d,
	0xea, 0x70, 0x80, 0xdb, 0xbc, 0xe7, 0x37, 0xf3, 0xde, 0x8c, 0xd6, 0xf0, 0x5c, 0x1b, 0xc5, 0x6a,
	0x54, 0xa9, 0x72, 0x51, 0xa5, 0xae, 0xc6, 0x4e, 0x45, 0x66, 0x55, 0xcc, 0x48, 0xfc, 0x0e, 0x20,
	0xcc, 0x55, 0xb9, 0x40, 0x84, 0xb0, 0x94, 0x4b, 0x4a, 0x82, 0x41, 0x30, 0x7c, 0x9a, 0xb9, 0xda,
	0x72, 0x5a, 0x2e, 0x28, 0x79, 0xe2, 0x39, 0x5b, 0xe3, 0x2b, 0x88, 0x65, 0xcd, 0xd7, 0xca, 0x24,
	0x6d, 0xc7, 0x6e, 0x10, 0x1e, 0x43, 0x24, 0xeb, 0x79, 0xa1, 0x92, 0xd0, 0xd1, 0x1e, 0x58, 0x76,
	0xa6, 0x56, 0x64, 0x92, 0xc8, 0xb3, 0x0e, 0x60, 0x02, 0x9d, 0x69, 0xc1, 0x46, 0x32, 0x25, 0xf1,
	0x20, 0x18, 0x46, 0xd9, 0x1d, 0x14, 0xe7, 0xd0, 0xb3, 0x69, 0x26, 0x37, 0x72, 0x7d, 0x53, 0x54,
	0x8c, 0xaf, 0x21, 0x72, 0xb1, 0x93, 0x60, 0xd0, 0x1e, 0x76, 0xc7, 0xcf, 0xd2, 0x4d, 0xee, 0xd4,
	0xaa, 0x32, 0xff, 0x4d, 0xfc, 0x80, 0xee, 0x47, 0xeb, 0x96, 0xb3, 0x21, 0xb9, 0xc4, 0x33, 0x38,
	0xd0, 0x9b, 0x7e, 0xb7, 0x4d, 0x77, 0xfc, 0x72, 0xa7, 0xed, 0x6e, 0x78, 0xb6, 0x95, 0xd9, 0x98,
	0x86, 0xe4, 0x7c, 0xed, 0x36, 0x3d, 0xc8, 0x3c, 0x10, 0x6f, 0xe0, 0xf0, 0x6b, 0xcd, 0xba, 0xe6,
	0x8c, 0x2a, 0xad, 0xca, 0x8a, 0xee, 0x75, 0xc1, 0x43, 0x5d, 0x07, 0xa2, 0xcf, 0x4b, 0xcd, 0x6b,
	0x71, 0x02, 0xed, 0xfc, 0x6a, 0x82, 0x47, 0xd0, 0xae, 0xe6, 0x7a, 0x73, 0x49, 0x5b, 0x8e, 0x7f,
	0x85, 0x70, 0xe4, 0xd3, 0x7d, 0x91, 0xa5, 0x5c, 0xd0, 0x92, 0x4a, 0xc6, 0x0b, 0xc0, 0x4b, 0x43,
	0x92, 0x69, 0x67, 0xe3, 0xc3, 0x6d, 0x56, 0x37, 0xb3, 0xbf, 0x3f, 0xbb, 0x68, 0xe1, 0x07, 0x78,
	0x71, 0x59, 0x1b, 0x43, 0x25, 0xff, 0x5f, 0xff, 0x05, 0xe0, 0x77, 0x3d, 0x7f, 0x6c, 0xbf, 0x5f,
	0xde, 0x7f, 0x34, 0x55, 0xb4, 0xf0, 0x3d, 0xf4, 0xfc, 0x71, 0x72, 0x96, 0x5c, 0x57, 0x0d, 0xdf,
	0x93, 0x2d, 0xde, 0xbd, 0xa1, 0x68, 0xe1, 0x5b, 0xe8, 0x4c, 0x0c, 0xad, 0x0a, 0xba, 0xc5, 0xde,
	0xbd, 0xdd, 0xd5, 0xa4, 0xbf, 0x83, 0x44, 0xeb, 0x34, 0xc0, 0x11, 0x74, 0x73, 0x96, 0x86, 0xbf,
	0xdd, 0x16, 0x3c, 0xbb, 0x6e, 0x98, 0x34, 0x63, 0x9d, 0x41, 0xe4, 0xde, 0x42, 0x43, 0x7a, 0xbc,
	0xc5, 0x0f, 0xde, 0x8a, 0xf3, 0x78, 0x07, 0xb1, 0x8f, 0xf8, 0x0f, 0x3b, 0x9c, 0x06, 0x98, 0x02,
	0xe4, 0xac, 0xf4, 0x5f, 0x5a, 0x1b, 0xc9, 0x3e, 0xc5, 0x3f, 0xc3, 0x74, 0xa4, 0xa7, 0xd3, 0xd8,
	0xfd, 0x81, 0xe7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x1e, 0x95, 0x11, 0x96, 0x03, 0x00,
	0x00,
}
