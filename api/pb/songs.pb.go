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

type OutputResponse struct {
	Bitrate              string   `protobuf:"bytes,1,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	Time                 string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
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

func (m *OutputResponse) GetBitrate() string {
	if m != nil {
		return m.Bitrate
	}
	return ""
}

func (m *OutputResponse) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
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

func init() {
	proto.RegisterType((*Song)(nil), "service.Song")
	proto.RegisterType((*SongPlaylist)(nil), "service.SongPlaylist")
	proto.RegisterType((*AudioStream)(nil), "service.AudioStream")
	proto.RegisterType((*OutputResponse)(nil), "service.OutputResponse")
	proto.RegisterType((*Empty)(nil), "service.Empty")
}

func init() { proto.RegisterFile("proto/songs.proto", fileDescriptor_23af63290edd64ad) }

var fileDescriptor_23af63290edd64ad = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x6f, 0xda, 0x40,
	0x10, 0xc5, 0xe0, 0x8f, 0x76, 0x68, 0x51, 0xbb, 0xa5, 0xed, 0x8a, 0x13, 0xda, 0x5e, 0x90, 0x2a,
	0x99, 0x02, 0xea, 0x11, 0xd4, 0x06, 0xe5, 0x18, 0x05, 0x19, 0xe5, 0x92, 0xdb, 0x02, 0x23, 0xc7,
	0x12, 0xf6, 0xae, 0xd6, 0x6b, 0x22, 0xfe, 0x43, 0xfe, 0x67, 0xfe, 0x46, 0xb4, 0x6b, 0xb0, 0x00,
	0x73, 0xca, 0x6d, 0xde, 0xf3, 0x7b, 0x3b, 0x6f, 0x66, 0x0c, 0x5f, 0xa5, 0x12, 0x5a, 0x0c, 0x73,
	0x91, 0xc5, 0x79, 0x68, 0x6b, 0x12, 0xe4, 0xa8, 0x76, 0xc9, 0x1a, 0xd9, 0x8b, 0x03, 0xee, 0x52,
	0x64, 0x31, 0x21, 0xe0, 0x66, 0x3c, 0x45, 0xea, 0xf4, 0x9d, 0xc1, 0xc7, 0xc8, 0xd6, 0x86, 0x93,
	0x3c, 0x46, 0xda, 0x2c, 0x39, 0x53, 0x93, 0x1f, 0xe0, 0xf3, 0x42, 0x3f, 0x09, 0x45, 0x5b, 0x96,
	0x3d, 0x20, 0xd2, 0x05, 0x8f, 0x17, 0x9b, 0x44, 0x50, 0xd7, 0xd2, 0x25, 0x30, 0xec, 0x5a, 0xec,
	0x50, 0x51, 0xaf, 0x64, 0x2d, 0x20, 0x14, 0x82, 0x55, 0xa2, 0x15, 0xd7, 0x48, 0xfd, 0xbe, 0x33,
	0xf0, 0xa2, 0x23, 0x64, 0x13, 0xf8, 0x64, 0xd2, 0x2c, 0xb6, 0x7c, 0xbf, 0x4d, 0x72, 0x4d, 0x7e,
	0x81, 0x67, 0x63, 0x53, 0xa7, 0xdf, 0x1a, 0xb4, 0xc7, 0x9f, 0xc3, 0x43, 0xee, 0xd0, 0xa8, 0xa2,
	0xf2, 0x1b, 0xfb, 0x07, 0xed, 0xff, 0xa6, 0xdb, 0x52, 0x2b, 0xe4, 0x29, 0x19, 0xc1, 0x07, 0x79,
	0xf0, 0xdb, 0x69, 0xda, 0xe3, 0xef, 0x67, 0xb6, 0xe3, 0xe3, 0x51, 0x25, 0x63, 0x33, 0xe8, 0xdc,
	0x17, 0x5a, 0x16, 0x3a, 0xc2, 0x5c, 0x8a, 0x2c, 0xc7, 0xd3, 0x88, 0xe5, 0x46, 0x8e, 0xd0, 0x2c,
	0x45, 0x27, 0x69, 0xb5, 0x14, 0x53, 0xb3, 0x00, 0xbc, 0xdb, 0x54, 0xea, 0xfd, 0xf8, 0xb5, 0x09,
	0x5f, 0xca, 0x18, 0x77, 0x3c, 0xe3, 0x31, 0xa6, 0x98, 0x69, 0x32, 0x05, 0x32, 0x57, 0xc8, 0x35,
	0x9e, 0x8d, 0xd6, 0xa9, 0x42, 0x59, 0x6b, 0xef, 0x7a, 0x48, 0xd6, 0x20, 0x33, 0xf8, 0x36, 0x2f,
	0x94, 0xc2, 0x4c, 0xbf, 0xcf, 0x3f, 0x05, 0xf2, 0x20, 0x37, 0x97, 0xed, 0xaf, 0xcb, 0x7b, 0x17,
	0xaf, 0xb2, 0x06, 0xf9, 0x0d, 0xc1, 0x42, 0xe1, 0x2e, 0xc1, 0xe7, 0x5a, 0xcb, 0xba, 0x78, 0x04,
	0x9e, 0x3d, 0x45, 0x4d, 0xda, 0xad, 0xf0, 0xc9, 0xa9, 0x58, 0xe3, 0x8f, 0x43, 0xfe, 0x82, 0x5f,
	0xee, 0xbe, 0xe6, 0xf9, 0x59, 0xe1, 0xf3, 0xe3, 0x18, 0xdb, 0x8d, 0xff, 0xe8, 0x86, 0x43, 0xb9,
	0x5a, 0xf9, 0xf6, 0x87, 0x9e, 0xbc, 0x05, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x62, 0x5e, 0xbd, 0xe5,
	0x02, 0x00, 0x00,
}
