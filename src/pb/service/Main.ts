// Original file: ../proto/main.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { AudioResponse as _service_AudioResponse, AudioResponse__Output as _service_AudioResponse__Output } from '../service/AudioResponse';
import type { BackgroundVideo as _service_BackgroundVideo, BackgroundVideo__Output as _service_BackgroundVideo__Output } from '../service/BackgroundVideo';
import type { BackgroundVideosResponse as _service_BackgroundVideosResponse, BackgroundVideosResponse__Output as _service_BackgroundVideosResponse__Output } from '../service/BackgroundVideosResponse';
import type { DevCredentials as _service_DevCredentials, DevCredentials__Output as _service_DevCredentials__Output } from '../service/DevCredentials';
import type { Empty as _google_protobuf_Empty, Empty__Output as _google_protobuf_Empty__Output } from '../google/protobuf/Empty';
import type { OutputRequest as _service_OutputRequest, OutputRequest__Output as _service_OutputRequest__Output } from '../service/OutputRequest';
import type { OutputResponse as _service_OutputResponse, OutputResponse__Output as _service_OutputResponse__Output } from '../service/OutputResponse';
import type { Overlays as _service_Overlays, Overlays__Output as _service_Overlays__Output } from '../service/Overlays';
import type { SDP as _service_SDP, SDP__Output as _service_SDP__Output } from '../service/SDP';
import type { SongPlaylist as _service_SongPlaylist, SongPlaylist__Output as _service_SongPlaylist__Output } from '../service/SongPlaylist';
import type { StatusNCSResponse as _service_StatusNCSResponse, StatusNCSResponse__Output as _service_StatusNCSResponse__Output } from '../service/StatusNCSResponse';
import type { StatusResponse as _service_StatusResponse, StatusResponse__Output as _service_StatusResponse__Output } from '../service/StatusResponse';
import type { TwitchStreamKey as _service_TwitchStreamKey, TwitchStreamKey__Output as _service_TwitchStreamKey__Output } from '../service/TwitchStreamKey';
import type { UploadVideoRequest as _service_UploadVideoRequest, UploadVideoRequest__Output as _service_UploadVideoRequest__Output } from '../service/UploadVideoRequest';
import type { UploadVideoResponse as _service_UploadVideoResponse, UploadVideoResponse__Output as _service_UploadVideoResponse__Output } from '../service/UploadVideoResponse';
import type { UserAuth as _service_UserAuth, UserAuth__Output as _service_UserAuth__Output } from '../service/UserAuth';

export interface MainClient extends grpc.Client {
  BackgroundVideos(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  BackgroundVideos(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  BackgroundVideos(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  BackgroundVideos(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  backgroundVideos(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  backgroundVideos(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  backgroundVideos(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  backgroundVideos(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_BackgroundVideosResponse__Output>): grpc.ClientUnaryCall;
  
  CheckTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  CheckTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  CheckTwitchDevCredentials(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  CheckTwitchDevCredentials(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  checkTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  checkTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  checkTwitchDevCredentials(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  checkTwitchDevCredentials(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_DevCredentials__Output>): grpc.ClientUnaryCall;
  
  CheckTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  CheckTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  CheckTwitchStreamKey(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  CheckTwitchStreamKey(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  checkTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  checkTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  checkTwitchStreamKey(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  checkTwitchStreamKey(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_TwitchStreamKey__Output>): grpc.ClientUnaryCall;
  
  CreateSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CreateSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CreateSongPlaylist(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CreateSongPlaylist(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  
  CurrentSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CurrentSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CurrentSongPlaylist(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CurrentSongPlaylist(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  
  DeleteBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteBackgroundVideo(argument: _service_BackgroundVideo, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteBackgroundVideo(argument: _service_BackgroundVideo, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteBackgroundVideo(argument: _service_BackgroundVideo, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteBackgroundVideo(argument: _service_BackgroundVideo, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  DeleteTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteTwitchDevCredentials(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteTwitchDevCredentials(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchDevCredentials(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchDevCredentials(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchDevCredentials(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  DeleteTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteTwitchStreamKey(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  DeleteTwitchStreamKey(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchStreamKey(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchStreamKey(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  deleteTwitchStreamKey(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  FindNewSongsNCS(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  FindNewSongsNCS(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  FindNewSongsNCS(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  FindNewSongsNCS(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  findNewSongsNcs(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  findNewSongsNcs(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  findNewSongsNcs(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  findNewSongsNcs(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  GetOverlays(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  GetOverlays(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  GetOverlays(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  GetOverlays(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  getOverlays(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  getOverlays(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  getOverlays(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  getOverlays(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_Overlays__Output>): grpc.ClientUnaryCall;
  
  Preview(argument: _service_SDP, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  Preview(argument: _service_SDP, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  Preview(argument: _service_SDP, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  Preview(argument: _service_SDP, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  
  SaveTwitchDevCredentials(argument: _service_DevCredentials, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  SaveTwitchDevCredentials(argument: _service_DevCredentials, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  SaveTwitchDevCredentials(argument: _service_DevCredentials, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  SaveTwitchDevCredentials(argument: _service_DevCredentials, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  saveTwitchDevCredentials(argument: _service_DevCredentials, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  saveTwitchDevCredentials(argument: _service_DevCredentials, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  saveTwitchDevCredentials(argument: _service_DevCredentials, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  saveTwitchDevCredentials(argument: _service_DevCredentials, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  StartAudio(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  StartAudio(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  StartAudio(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  StartAudio(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  
  StartOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  StartOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  StartOutput(argument: _service_OutputRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  StartOutput(argument: _service_OutputRequest, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  
  StartStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StartStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StartStream(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StartStream(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  Status(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  Status(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  Status(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  Status(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  
  StatusNCS(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  StatusNCS(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  StatusNCS(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  StatusNCS(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  statusNcs(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  statusNcs(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  statusNcs(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  statusNcs(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_service_StatusNCSResponse__Output>): grpc.ClientUnaryCall;
  
  StopOutput(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StopOutput(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StopOutput(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StopOutput(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  StopStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StopStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StopStream(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  StopStream(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _google_protobuf_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _google_protobuf_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _google_protobuf_Empty, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  SwapBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  SwapBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  SwapBackgroundVideo(argument: _service_BackgroundVideo, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  SwapBackgroundVideo(argument: _service_BackgroundVideo, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  swapBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  swapBackgroundVideo(argument: _service_BackgroundVideo, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  swapBackgroundVideo(argument: _service_BackgroundVideo, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  swapBackgroundVideo(argument: _service_BackgroundVideo, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  TwitchAccessToken(argument: _service_UserAuth, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  TwitchAccessToken(argument: _service_UserAuth, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  TwitchAccessToken(argument: _service_UserAuth, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  TwitchAccessToken(argument: _service_UserAuth, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchAccessToken(argument: _service_UserAuth, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchAccessToken(argument: _service_UserAuth, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchAccessToken(argument: _service_UserAuth, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchAccessToken(argument: _service_UserAuth, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  TwitchSaveStreamKey(argument: _service_TwitchStreamKey, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  TwitchSaveStreamKey(argument: _service_TwitchStreamKey, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  TwitchSaveStreamKey(argument: _service_TwitchStreamKey, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  TwitchSaveStreamKey(argument: _service_TwitchStreamKey, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchSaveStreamKey(argument: _service_TwitchStreamKey, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchSaveStreamKey(argument: _service_TwitchStreamKey, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchSaveStreamKey(argument: _service_TwitchStreamKey, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  twitchSaveStreamKey(argument: _service_TwitchStreamKey, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  UpdateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  UpdateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  UpdateSongPlaylist(argument: _service_SongPlaylist, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  UpdateSongPlaylist(argument: _service_SongPlaylist, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, options: grpc.CallOptions, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, callback: grpc.requestCallback<_google_protobuf_Empty__Output>): grpc.ClientUnaryCall;
  
  UploadVideo(metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  UploadVideo(metadata: grpc.Metadata, callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  UploadVideo(options: grpc.CallOptions, callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  UploadVideo(callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  uploadVideo(metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  uploadVideo(metadata: grpc.Metadata, callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  uploadVideo(options: grpc.CallOptions, callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  uploadVideo(callback: grpc.requestCallback<_service_UploadVideoResponse__Output>): grpc.ClientWritableStream<_service_UploadVideoRequest>;
  
}

export interface MainHandlers extends grpc.UntypedServiceImplementation {
  BackgroundVideos: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_BackgroundVideosResponse>;
  
  CheckTwitchDevCredentials: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_DevCredentials>;
  
  CheckTwitchStreamKey: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_TwitchStreamKey>;
  
  CreateSongPlaylist: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_SongPlaylist>;
  
  CurrentSongPlaylist: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_SongPlaylist>;
  
  DeleteBackgroundVideo: grpc.handleUnaryCall<_service_BackgroundVideo__Output, _google_protobuf_Empty>;
  
  DeleteTwitchDevCredentials: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _google_protobuf_Empty>;
  
  DeleteTwitchStreamKey: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _google_protobuf_Empty>;
  
  FindNewSongsNCS: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _google_protobuf_Empty>;
  
  GetOverlays: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_Overlays>;
  
  Preview: grpc.handleUnaryCall<_service_SDP__Output, _service_SDP>;
  
  SaveTwitchDevCredentials: grpc.handleUnaryCall<_service_DevCredentials__Output, _google_protobuf_Empty>;
  
  StartAudio: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_AudioResponse>;
  
  StartOutput: grpc.handleUnaryCall<_service_OutputRequest__Output, _service_OutputResponse>;
  
  StartStream: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _google_protobuf_Empty>;
  
  Status: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_StatusResponse>;
  
  StatusNCS: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _service_StatusNCSResponse>;
  
  StopOutput: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _google_protobuf_Empty>;
  
  StopStream: grpc.handleUnaryCall<_google_protobuf_Empty__Output, _google_protobuf_Empty>;
  
  SwapBackgroundVideo: grpc.handleUnaryCall<_service_BackgroundVideo__Output, _google_protobuf_Empty>;
  
  TwitchAccessToken: grpc.handleUnaryCall<_service_UserAuth__Output, _google_protobuf_Empty>;
  
  TwitchSaveStreamKey: grpc.handleUnaryCall<_service_TwitchStreamKey__Output, _google_protobuf_Empty>;
  
  UpdateSongPlaylist: grpc.handleUnaryCall<_service_SongPlaylist__Output, _google_protobuf_Empty>;
  
  UploadVideo: grpc.handleClientStreamingCall<_service_UploadVideoRequest__Output, _service_UploadVideoResponse>;
  
}

export interface MainDefinition extends grpc.ServiceDefinition {
  BackgroundVideos: MethodDefinition<_google_protobuf_Empty, _service_BackgroundVideosResponse, _google_protobuf_Empty__Output, _service_BackgroundVideosResponse__Output>
  CheckTwitchDevCredentials: MethodDefinition<_google_protobuf_Empty, _service_DevCredentials, _google_protobuf_Empty__Output, _service_DevCredentials__Output>
  CheckTwitchStreamKey: MethodDefinition<_google_protobuf_Empty, _service_TwitchStreamKey, _google_protobuf_Empty__Output, _service_TwitchStreamKey__Output>
  CreateSongPlaylist: MethodDefinition<_google_protobuf_Empty, _service_SongPlaylist, _google_protobuf_Empty__Output, _service_SongPlaylist__Output>
  CurrentSongPlaylist: MethodDefinition<_google_protobuf_Empty, _service_SongPlaylist, _google_protobuf_Empty__Output, _service_SongPlaylist__Output>
  DeleteBackgroundVideo: MethodDefinition<_service_BackgroundVideo, _google_protobuf_Empty, _service_BackgroundVideo__Output, _google_protobuf_Empty__Output>
  DeleteTwitchDevCredentials: MethodDefinition<_google_protobuf_Empty, _google_protobuf_Empty, _google_protobuf_Empty__Output, _google_protobuf_Empty__Output>
  DeleteTwitchStreamKey: MethodDefinition<_google_protobuf_Empty, _google_protobuf_Empty, _google_protobuf_Empty__Output, _google_protobuf_Empty__Output>
  FindNewSongsNCS: MethodDefinition<_google_protobuf_Empty, _google_protobuf_Empty, _google_protobuf_Empty__Output, _google_protobuf_Empty__Output>
  GetOverlays: MethodDefinition<_google_protobuf_Empty, _service_Overlays, _google_protobuf_Empty__Output, _service_Overlays__Output>
  Preview: MethodDefinition<_service_SDP, _service_SDP, _service_SDP__Output, _service_SDP__Output>
  SaveTwitchDevCredentials: MethodDefinition<_service_DevCredentials, _google_protobuf_Empty, _service_DevCredentials__Output, _google_protobuf_Empty__Output>
  StartAudio: MethodDefinition<_google_protobuf_Empty, _service_AudioResponse, _google_protobuf_Empty__Output, _service_AudioResponse__Output>
  StartOutput: MethodDefinition<_service_OutputRequest, _service_OutputResponse, _service_OutputRequest__Output, _service_OutputResponse__Output>
  StartStream: MethodDefinition<_google_protobuf_Empty, _google_protobuf_Empty, _google_protobuf_Empty__Output, _google_protobuf_Empty__Output>
  Status: MethodDefinition<_google_protobuf_Empty, _service_StatusResponse, _google_protobuf_Empty__Output, _service_StatusResponse__Output>
  StatusNCS: MethodDefinition<_google_protobuf_Empty, _service_StatusNCSResponse, _google_protobuf_Empty__Output, _service_StatusNCSResponse__Output>
  StopOutput: MethodDefinition<_google_protobuf_Empty, _google_protobuf_Empty, _google_protobuf_Empty__Output, _google_protobuf_Empty__Output>
  StopStream: MethodDefinition<_google_protobuf_Empty, _google_protobuf_Empty, _google_protobuf_Empty__Output, _google_protobuf_Empty__Output>
  SwapBackgroundVideo: MethodDefinition<_service_BackgroundVideo, _google_protobuf_Empty, _service_BackgroundVideo__Output, _google_protobuf_Empty__Output>
  TwitchAccessToken: MethodDefinition<_service_UserAuth, _google_protobuf_Empty, _service_UserAuth__Output, _google_protobuf_Empty__Output>
  TwitchSaveStreamKey: MethodDefinition<_service_TwitchStreamKey, _google_protobuf_Empty, _service_TwitchStreamKey__Output, _google_protobuf_Empty__Output>
  UpdateSongPlaylist: MethodDefinition<_service_SongPlaylist, _google_protobuf_Empty, _service_SongPlaylist__Output, _google_protobuf_Empty__Output>
  UploadVideo: MethodDefinition<_service_UploadVideoRequest, _service_UploadVideoResponse, _service_UploadVideoRequest__Output, _service_UploadVideoResponse__Output>
}
