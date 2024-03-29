// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: main.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Main_CreateSongPlaylist_FullMethodName         = "/service.Main/CreateSongPlaylist"
	Main_CurrentSongPlaylist_FullMethodName        = "/service.Main/CurrentSongPlaylist"
	Main_UpdateSongPlaylist_FullMethodName         = "/service.Main/UpdateSongPlaylist"
	Main_Preview_FullMethodName                    = "/service.Main/Preview"
	Main_StartPreview_FullMethodName               = "/service.Main/StartPreview"
	Main_StopPreview_FullMethodName                = "/service.Main/StopPreview"
	Main_StartStream_FullMethodName                = "/service.Main/StartStream"
	Main_StopStream_FullMethodName                 = "/service.Main/StopStream"
	Main_Status_FullMethodName                     = "/service.Main/Status"
	Main_SwapBackgroundVideo_FullMethodName        = "/service.Main/SwapBackgroundVideo"
	Main_FindNewSongsNCS_FullMethodName            = "/service.Main/FindNewSongsNCS"
	Main_StatusNCS_FullMethodName                  = "/service.Main/StatusNCS"
	Main_TwitchSaveStreamKey_FullMethodName        = "/service.Main/TwitchSaveStreamKey"
	Main_CheckTwitchStreamKey_FullMethodName       = "/service.Main/CheckTwitchStreamKey"
	Main_DeleteTwitchStreamKey_FullMethodName      = "/service.Main/DeleteTwitchStreamKey"
	Main_SaveTwitchDevCredentials_FullMethodName   = "/service.Main/SaveTwitchDevCredentials"
	Main_CheckTwitchDevCredentials_FullMethodName  = "/service.Main/CheckTwitchDevCredentials"
	Main_DeleteTwitchDevCredentials_FullMethodName = "/service.Main/DeleteTwitchDevCredentials"
	Main_TwitchAccessToken_FullMethodName          = "/service.Main/TwitchAccessToken"
	Main_StreamParameters_FullMethodName           = "/service.Main/StreamParameters"
	Main_SaveStreamParameters_FullMethodName       = "/service.Main/SaveStreamParameters"
	Main_GetOverlays_FullMethodName                = "/service.Main/GetOverlays"
	Main_BackgroundVideos_FullMethodName           = "/service.Main/BackgroundVideos"
	Main_UploadVideo_FullMethodName                = "/service.Main/UploadVideo"
	Main_DeleteBackgroundVideo_FullMethodName      = "/service.Main/DeleteBackgroundVideo"
)

// MainClient is the client API for Main service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainClient interface {
	CreateSongPlaylist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SongPlaylist, error)
	CurrentSongPlaylist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SongPlaylist, error)
	UpdateSongPlaylist(ctx context.Context, in *SongPlaylist, opts ...grpc.CallOption) (*empty.Empty, error)
	Preview(ctx context.Context, in *SDP, opts ...grpc.CallOption) (*SDP, error)
	StartPreview(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	StopPreview(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	StartStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StreamResponse, error)
	StopStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error)
	SwapBackgroundVideo(ctx context.Context, in *BackgroundVideo, opts ...grpc.CallOption) (*empty.Empty, error)
	FindNewSongsNCS(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	StatusNCS(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusNCSResponse, error)
	TwitchSaveStreamKey(ctx context.Context, in *TwitchStreamKey, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckTwitchStreamKey(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TwitchStreamKey, error)
	DeleteTwitchStreamKey(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	SaveTwitchDevCredentials(ctx context.Context, in *DevCredentials, opts ...grpc.CallOption) (*empty.Empty, error)
	CheckTwitchDevCredentials(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DevCredentials, error)
	DeleteTwitchDevCredentials(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
	TwitchAccessToken(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*empty.Empty, error)
	StreamParameters(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StreamParametersResponse, error)
	SaveStreamParameters(ctx context.Context, in *SaveStreamParametersRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetOverlays(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Overlays, error)
	BackgroundVideos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BackgroundVideosResponse, error)
	UploadVideo(ctx context.Context, opts ...grpc.CallOption) (Main_UploadVideoClient, error)
	DeleteBackgroundVideo(ctx context.Context, in *BackgroundVideo, opts ...grpc.CallOption) (*empty.Empty, error)
}

type mainClient struct {
	cc grpc.ClientConnInterface
}

func NewMainClient(cc grpc.ClientConnInterface) MainClient {
	return &mainClient{cc}
}

func (c *mainClient) CreateSongPlaylist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SongPlaylist, error) {
	out := new(SongPlaylist)
	err := c.cc.Invoke(ctx, Main_CreateSongPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) CurrentSongPlaylist(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*SongPlaylist, error) {
	out := new(SongPlaylist)
	err := c.cc.Invoke(ctx, Main_CurrentSongPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) UpdateSongPlaylist(ctx context.Context, in *SongPlaylist, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_UpdateSongPlaylist_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) Preview(ctx context.Context, in *SDP, opts ...grpc.CallOption) (*SDP, error) {
	out := new(SDP)
	err := c.cc.Invoke(ctx, Main_Preview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) StartPreview(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Main_StartPreview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) StopPreview(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Main_StopPreview_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) StartStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StreamResponse, error) {
	out := new(StreamResponse)
	err := c.cc.Invoke(ctx, Main_StartStream_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) StopStream(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Main_StopStream_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) Status(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, Main_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) SwapBackgroundVideo(ctx context.Context, in *BackgroundVideo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_SwapBackgroundVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) FindNewSongsNCS(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_FindNewSongsNCS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) StatusNCS(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StatusNCSResponse, error) {
	out := new(StatusNCSResponse)
	err := c.cc.Invoke(ctx, Main_StatusNCS_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) TwitchSaveStreamKey(ctx context.Context, in *TwitchStreamKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_TwitchSaveStreamKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) CheckTwitchStreamKey(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*TwitchStreamKey, error) {
	out := new(TwitchStreamKey)
	err := c.cc.Invoke(ctx, Main_CheckTwitchStreamKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) DeleteTwitchStreamKey(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_DeleteTwitchStreamKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) SaveTwitchDevCredentials(ctx context.Context, in *DevCredentials, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_SaveTwitchDevCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) CheckTwitchDevCredentials(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DevCredentials, error) {
	out := new(DevCredentials)
	err := c.cc.Invoke(ctx, Main_CheckTwitchDevCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) DeleteTwitchDevCredentials(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_DeleteTwitchDevCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) TwitchAccessToken(ctx context.Context, in *UserAuth, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_TwitchAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) StreamParameters(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StreamParametersResponse, error) {
	out := new(StreamParametersResponse)
	err := c.cc.Invoke(ctx, Main_StreamParameters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) SaveStreamParameters(ctx context.Context, in *SaveStreamParametersRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_SaveStreamParameters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) GetOverlays(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Overlays, error) {
	out := new(Overlays)
	err := c.cc.Invoke(ctx, Main_GetOverlays_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) BackgroundVideos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*BackgroundVideosResponse, error) {
	out := new(BackgroundVideosResponse)
	err := c.cc.Invoke(ctx, Main_BackgroundVideos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mainClient) UploadVideo(ctx context.Context, opts ...grpc.CallOption) (Main_UploadVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Main_ServiceDesc.Streams[0], Main_UploadVideo_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &mainUploadVideoClient{stream}
	return x, nil
}

type Main_UploadVideoClient interface {
	Send(*UploadVideoRequest) error
	CloseAndRecv() (*UploadVideoResponse, error)
	grpc.ClientStream
}

type mainUploadVideoClient struct {
	grpc.ClientStream
}

func (x *mainUploadVideoClient) Send(m *UploadVideoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mainUploadVideoClient) CloseAndRecv() (*UploadVideoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadVideoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mainClient) DeleteBackgroundVideo(ctx context.Context, in *BackgroundVideo, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, Main_DeleteBackgroundVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainServer is the server API for Main service.
// All implementations must embed UnimplementedMainServer
// for forward compatibility
type MainServer interface {
	CreateSongPlaylist(context.Context, *empty.Empty) (*SongPlaylist, error)
	CurrentSongPlaylist(context.Context, *empty.Empty) (*SongPlaylist, error)
	UpdateSongPlaylist(context.Context, *SongPlaylist) (*empty.Empty, error)
	Preview(context.Context, *SDP) (*SDP, error)
	StartPreview(context.Context, *empty.Empty) (*StatusResponse, error)
	StopPreview(context.Context, *empty.Empty) (*StatusResponse, error)
	StartStream(context.Context, *empty.Empty) (*StreamResponse, error)
	StopStream(context.Context, *empty.Empty) (*StatusResponse, error)
	Status(context.Context, *empty.Empty) (*StatusResponse, error)
	SwapBackgroundVideo(context.Context, *BackgroundVideo) (*empty.Empty, error)
	FindNewSongsNCS(context.Context, *empty.Empty) (*empty.Empty, error)
	StatusNCS(context.Context, *empty.Empty) (*StatusNCSResponse, error)
	TwitchSaveStreamKey(context.Context, *TwitchStreamKey) (*empty.Empty, error)
	CheckTwitchStreamKey(context.Context, *empty.Empty) (*TwitchStreamKey, error)
	DeleteTwitchStreamKey(context.Context, *empty.Empty) (*empty.Empty, error)
	SaveTwitchDevCredentials(context.Context, *DevCredentials) (*empty.Empty, error)
	CheckTwitchDevCredentials(context.Context, *empty.Empty) (*DevCredentials, error)
	DeleteTwitchDevCredentials(context.Context, *empty.Empty) (*empty.Empty, error)
	TwitchAccessToken(context.Context, *UserAuth) (*empty.Empty, error)
	StreamParameters(context.Context, *empty.Empty) (*StreamParametersResponse, error)
	SaveStreamParameters(context.Context, *SaveStreamParametersRequest) (*empty.Empty, error)
	GetOverlays(context.Context, *empty.Empty) (*Overlays, error)
	BackgroundVideos(context.Context, *empty.Empty) (*BackgroundVideosResponse, error)
	UploadVideo(Main_UploadVideoServer) error
	DeleteBackgroundVideo(context.Context, *BackgroundVideo) (*empty.Empty, error)
	mustEmbedUnimplementedMainServer()
}

// UnimplementedMainServer must be embedded to have forward compatible implementations.
type UnimplementedMainServer struct {
}

func (UnimplementedMainServer) CreateSongPlaylist(context.Context, *empty.Empty) (*SongPlaylist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSongPlaylist not implemented")
}
func (UnimplementedMainServer) CurrentSongPlaylist(context.Context, *empty.Empty) (*SongPlaylist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentSongPlaylist not implemented")
}
func (UnimplementedMainServer) UpdateSongPlaylist(context.Context, *SongPlaylist) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSongPlaylist not implemented")
}
func (UnimplementedMainServer) Preview(context.Context, *SDP) (*SDP, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Preview not implemented")
}
func (UnimplementedMainServer) StartPreview(context.Context, *empty.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPreview not implemented")
}
func (UnimplementedMainServer) StopPreview(context.Context, *empty.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopPreview not implemented")
}
func (UnimplementedMainServer) StartStream(context.Context, *empty.Empty) (*StreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartStream not implemented")
}
func (UnimplementedMainServer) StopStream(context.Context, *empty.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopStream not implemented")
}
func (UnimplementedMainServer) Status(context.Context, *empty.Empty) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedMainServer) SwapBackgroundVideo(context.Context, *BackgroundVideo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapBackgroundVideo not implemented")
}
func (UnimplementedMainServer) FindNewSongsNCS(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindNewSongsNCS not implemented")
}
func (UnimplementedMainServer) StatusNCS(context.Context, *empty.Empty) (*StatusNCSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusNCS not implemented")
}
func (UnimplementedMainServer) TwitchSaveStreamKey(context.Context, *TwitchStreamKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TwitchSaveStreamKey not implemented")
}
func (UnimplementedMainServer) CheckTwitchStreamKey(context.Context, *empty.Empty) (*TwitchStreamKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTwitchStreamKey not implemented")
}
func (UnimplementedMainServer) DeleteTwitchStreamKey(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTwitchStreamKey not implemented")
}
func (UnimplementedMainServer) SaveTwitchDevCredentials(context.Context, *DevCredentials) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveTwitchDevCredentials not implemented")
}
func (UnimplementedMainServer) CheckTwitchDevCredentials(context.Context, *empty.Empty) (*DevCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckTwitchDevCredentials not implemented")
}
func (UnimplementedMainServer) DeleteTwitchDevCredentials(context.Context, *empty.Empty) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTwitchDevCredentials not implemented")
}
func (UnimplementedMainServer) TwitchAccessToken(context.Context, *UserAuth) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TwitchAccessToken not implemented")
}
func (UnimplementedMainServer) StreamParameters(context.Context, *empty.Empty) (*StreamParametersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StreamParameters not implemented")
}
func (UnimplementedMainServer) SaveStreamParameters(context.Context, *SaveStreamParametersRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveStreamParameters not implemented")
}
func (UnimplementedMainServer) GetOverlays(context.Context, *empty.Empty) (*Overlays, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOverlays not implemented")
}
func (UnimplementedMainServer) BackgroundVideos(context.Context, *empty.Empty) (*BackgroundVideosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BackgroundVideos not implemented")
}
func (UnimplementedMainServer) UploadVideo(Main_UploadVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedMainServer) DeleteBackgroundVideo(context.Context, *BackgroundVideo) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBackgroundVideo not implemented")
}
func (UnimplementedMainServer) mustEmbedUnimplementedMainServer() {}

// UnsafeMainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainServer will
// result in compilation errors.
type UnsafeMainServer interface {
	mustEmbedUnimplementedMainServer()
}

func RegisterMainServer(s grpc.ServiceRegistrar, srv MainServer) {
	s.RegisterService(&Main_ServiceDesc, srv)
}

func _Main_CreateSongPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).CreateSongPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_CreateSongPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).CreateSongPlaylist(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_CurrentSongPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).CurrentSongPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_CurrentSongPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).CurrentSongPlaylist(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_UpdateSongPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongPlaylist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).UpdateSongPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_UpdateSongPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).UpdateSongPlaylist(ctx, req.(*SongPlaylist))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_Preview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SDP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).Preview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_Preview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).Preview(ctx, req.(*SDP))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_StartPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).StartPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_StartPreview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).StartPreview(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_StopPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).StopPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_StopPreview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).StopPreview(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_StartStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).StartStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_StartStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).StartStream(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_StopStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).StopStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_StopStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).StopStream(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).Status(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_SwapBackgroundVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackgroundVideo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).SwapBackgroundVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_SwapBackgroundVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).SwapBackgroundVideo(ctx, req.(*BackgroundVideo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_FindNewSongsNCS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).FindNewSongsNCS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_FindNewSongsNCS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).FindNewSongsNCS(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_StatusNCS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).StatusNCS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_StatusNCS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).StatusNCS(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_TwitchSaveStreamKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TwitchStreamKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).TwitchSaveStreamKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_TwitchSaveStreamKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).TwitchSaveStreamKey(ctx, req.(*TwitchStreamKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_CheckTwitchStreamKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).CheckTwitchStreamKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_CheckTwitchStreamKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).CheckTwitchStreamKey(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_DeleteTwitchStreamKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).DeleteTwitchStreamKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_DeleteTwitchStreamKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).DeleteTwitchStreamKey(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_SaveTwitchDevCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DevCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).SaveTwitchDevCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_SaveTwitchDevCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).SaveTwitchDevCredentials(ctx, req.(*DevCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_CheckTwitchDevCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).CheckTwitchDevCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_CheckTwitchDevCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).CheckTwitchDevCredentials(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_DeleteTwitchDevCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).DeleteTwitchDevCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_DeleteTwitchDevCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).DeleteTwitchDevCredentials(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_TwitchAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAuth)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).TwitchAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_TwitchAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).TwitchAccessToken(ctx, req.(*UserAuth))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_StreamParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).StreamParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_StreamParameters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).StreamParameters(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_SaveStreamParameters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveStreamParametersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).SaveStreamParameters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_SaveStreamParameters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).SaveStreamParameters(ctx, req.(*SaveStreamParametersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_GetOverlays_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).GetOverlays(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_GetOverlays_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).GetOverlays(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_BackgroundVideos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).BackgroundVideos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_BackgroundVideos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).BackgroundVideos(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Main_UploadVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MainServer).UploadVideo(&mainUploadVideoServer{stream})
}

type Main_UploadVideoServer interface {
	SendAndClose(*UploadVideoResponse) error
	Recv() (*UploadVideoRequest, error)
	grpc.ServerStream
}

type mainUploadVideoServer struct {
	grpc.ServerStream
}

func (x *mainUploadVideoServer) SendAndClose(m *UploadVideoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mainUploadVideoServer) Recv() (*UploadVideoRequest, error) {
	m := new(UploadVideoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Main_DeleteBackgroundVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackgroundVideo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainServer).DeleteBackgroundVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Main_DeleteBackgroundVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainServer).DeleteBackgroundVideo(ctx, req.(*BackgroundVideo))
	}
	return interceptor(ctx, in, info, handler)
}

// Main_ServiceDesc is the grpc.ServiceDesc for Main service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Main_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Main",
	HandlerType: (*MainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSongPlaylist",
			Handler:    _Main_CreateSongPlaylist_Handler,
		},
		{
			MethodName: "CurrentSongPlaylist",
			Handler:    _Main_CurrentSongPlaylist_Handler,
		},
		{
			MethodName: "UpdateSongPlaylist",
			Handler:    _Main_UpdateSongPlaylist_Handler,
		},
		{
			MethodName: "Preview",
			Handler:    _Main_Preview_Handler,
		},
		{
			MethodName: "StartPreview",
			Handler:    _Main_StartPreview_Handler,
		},
		{
			MethodName: "StopPreview",
			Handler:    _Main_StopPreview_Handler,
		},
		{
			MethodName: "StartStream",
			Handler:    _Main_StartStream_Handler,
		},
		{
			MethodName: "StopStream",
			Handler:    _Main_StopStream_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Main_Status_Handler,
		},
		{
			MethodName: "SwapBackgroundVideo",
			Handler:    _Main_SwapBackgroundVideo_Handler,
		},
		{
			MethodName: "FindNewSongsNCS",
			Handler:    _Main_FindNewSongsNCS_Handler,
		},
		{
			MethodName: "StatusNCS",
			Handler:    _Main_StatusNCS_Handler,
		},
		{
			MethodName: "TwitchSaveStreamKey",
			Handler:    _Main_TwitchSaveStreamKey_Handler,
		},
		{
			MethodName: "CheckTwitchStreamKey",
			Handler:    _Main_CheckTwitchStreamKey_Handler,
		},
		{
			MethodName: "DeleteTwitchStreamKey",
			Handler:    _Main_DeleteTwitchStreamKey_Handler,
		},
		{
			MethodName: "SaveTwitchDevCredentials",
			Handler:    _Main_SaveTwitchDevCredentials_Handler,
		},
		{
			MethodName: "CheckTwitchDevCredentials",
			Handler:    _Main_CheckTwitchDevCredentials_Handler,
		},
		{
			MethodName: "DeleteTwitchDevCredentials",
			Handler:    _Main_DeleteTwitchDevCredentials_Handler,
		},
		{
			MethodName: "TwitchAccessToken",
			Handler:    _Main_TwitchAccessToken_Handler,
		},
		{
			MethodName: "StreamParameters",
			Handler:    _Main_StreamParameters_Handler,
		},
		{
			MethodName: "SaveStreamParameters",
			Handler:    _Main_SaveStreamParameters_Handler,
		},
		{
			MethodName: "GetOverlays",
			Handler:    _Main_GetOverlays_Handler,
		},
		{
			MethodName: "BackgroundVideos",
			Handler:    _Main_BackgroundVideos_Handler,
		},
		{
			MethodName: "DeleteBackgroundVideo",
			Handler:    _Main_DeleteBackgroundVideo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadVideo",
			Handler:       _Main_UploadVideo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "main.proto",
}
