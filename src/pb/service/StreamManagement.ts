// Original file: ../proto/songs.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { AudioResponse as _service_AudioResponse, AudioResponse__Output as _service_AudioResponse__Output } from '../service/AudioResponse';
import type { Empty as _service_Empty, Empty__Output as _service_Empty__Output } from '../service/Empty';
import type { OutputRequest as _service_OutputRequest, OutputRequest__Output as _service_OutputRequest__Output } from '../service/OutputRequest';
import type { OutputResponse as _service_OutputResponse, OutputResponse__Output as _service_OutputResponse__Output } from '../service/OutputResponse';
import type { SDP as _service_SDP, SDP__Output as _service_SDP__Output } from '../service/SDP';
import type { SongPlaylist as _service_SongPlaylist, SongPlaylist__Output as _service_SongPlaylist__Output } from '../service/SongPlaylist';
import type { StatusResponse as _service_StatusResponse, StatusResponse__Output as _service_StatusResponse__Output } from '../service/StatusResponse';

export interface StreamManagementClient extends grpc.Client {
  CreateSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CreateSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CreateSongPlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CreateSongPlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  createSongPlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  
  CurrentSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CurrentSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CurrentSongPlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  CurrentSongPlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  currentSongPlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_SongPlaylist__Output>): grpc.ClientUnaryCall;
  
  Preview(argument: _service_SDP, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  Preview(argument: _service_SDP, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  Preview(argument: _service_SDP, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  Preview(argument: _service_SDP, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, options: grpc.CallOptions, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  preview(argument: _service_SDP, callback: grpc.requestCallback<_service_SDP__Output>): grpc.ClientUnaryCall;
  
  StartAudio(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  StartAudio(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  StartAudio(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  StartAudio(argument: _service_Empty, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  startAudio(argument: _service_Empty, callback: grpc.requestCallback<_service_AudioResponse__Output>): grpc.ClientUnaryCall;
  
  StartOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  StartOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  StartOutput(argument: _service_OutputRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  StartOutput(argument: _service_OutputRequest, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  startOutput(argument: _service_OutputRequest, callback: grpc.requestCallback<_service_OutputResponse__Output>): grpc.ClientUnaryCall;
  
  StartStream(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StartStream(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StartStream(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StartStream(argument: _service_Empty, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  startStream(argument: _service_Empty, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  
  Status(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  Status(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  Status(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  Status(argument: _service_Empty, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  status(argument: _service_Empty, callback: grpc.requestCallback<_service_StatusResponse__Output>): grpc.ClientUnaryCall;
  
  StopOutput(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StopOutput(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StopOutput(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StopOutput(argument: _service_Empty, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopOutput(argument: _service_Empty, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  
  StopStream(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StopStream(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StopStream(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  StopStream(argument: _service_Empty, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  stopStream(argument: _service_Empty, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  
  UpdateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  UpdateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  UpdateSongPlaylist(argument: _service_SongPlaylist, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  UpdateSongPlaylist(argument: _service_SongPlaylist, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  updateSongPlaylist(argument: _service_SongPlaylist, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientUnaryCall;
  
}

export interface StreamManagementHandlers extends grpc.UntypedServiceImplementation {
  CreateSongPlaylist: grpc.handleUnaryCall<_service_Empty__Output, _service_SongPlaylist>;
  
  CurrentSongPlaylist: grpc.handleUnaryCall<_service_Empty__Output, _service_SongPlaylist>;
  
  Preview: grpc.handleUnaryCall<_service_SDP__Output, _service_SDP>;
  
  StartAudio: grpc.handleUnaryCall<_service_Empty__Output, _service_AudioResponse>;
  
  StartOutput: grpc.handleUnaryCall<_service_OutputRequest__Output, _service_OutputResponse>;
  
  StartStream: grpc.handleUnaryCall<_service_Empty__Output, _service_Empty>;
  
  Status: grpc.handleUnaryCall<_service_Empty__Output, _service_StatusResponse>;
  
  StopOutput: grpc.handleUnaryCall<_service_Empty__Output, _service_Empty>;
  
  StopStream: grpc.handleUnaryCall<_service_Empty__Output, _service_Empty>;
  
  UpdateSongPlaylist: grpc.handleUnaryCall<_service_SongPlaylist__Output, _service_Empty>;
  
}

export interface StreamManagementDefinition extends grpc.ServiceDefinition {
  CreateSongPlaylist: MethodDefinition<_service_Empty, _service_SongPlaylist, _service_Empty__Output, _service_SongPlaylist__Output>
  CurrentSongPlaylist: MethodDefinition<_service_Empty, _service_SongPlaylist, _service_Empty__Output, _service_SongPlaylist__Output>
  Preview: MethodDefinition<_service_SDP, _service_SDP, _service_SDP__Output, _service_SDP__Output>
  StartAudio: MethodDefinition<_service_Empty, _service_AudioResponse, _service_Empty__Output, _service_AudioResponse__Output>
  StartOutput: MethodDefinition<_service_OutputRequest, _service_OutputResponse, _service_OutputRequest__Output, _service_OutputResponse__Output>
  StartStream: MethodDefinition<_service_Empty, _service_Empty, _service_Empty__Output, _service_Empty__Output>
  Status: MethodDefinition<_service_Empty, _service_StatusResponse, _service_Empty__Output, _service_StatusResponse__Output>
  StopOutput: MethodDefinition<_service_Empty, _service_Empty, _service_Empty__Output, _service_Empty__Output>
  StopStream: MethodDefinition<_service_Empty, _service_Empty, _service_Empty__Output, _service_Empty__Output>
  UpdateSongPlaylist: MethodDefinition<_service_SongPlaylist, _service_Empty, _service_SongPlaylist__Output, _service_Empty__Output>
}
