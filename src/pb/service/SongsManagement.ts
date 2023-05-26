// Original file: ../proto/songs.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Empty as _service_Empty, Empty__Output as _service_Empty__Output } from '../service/Empty';
import type { Playlist as _service_Playlist, Playlist__Output as _service_Playlist__Output } from '../service/Playlist';
import type { UpdatePlaylistRequest as _service_UpdatePlaylistRequest, UpdatePlaylistRequest__Output as _service_UpdatePlaylistRequest__Output } from '../service/UpdatePlaylistRequest';

export interface SongsManagementClient extends grpc.Client {
  CreatePlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  CreatePlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  CreatePlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  CreatePlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  
  UpdatePlaylst(metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  UpdatePlaylst(metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  UpdatePlaylst(options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  UpdatePlaylst(callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  updatePlaylst(metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  updatePlaylst(metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  updatePlaylst(options: grpc.CallOptions, callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  updatePlaylst(callback: grpc.requestCallback<_service_Empty__Output>): grpc.ClientWritableStream<_service_UpdatePlaylistRequest>;
  
}

export interface SongsManagementHandlers extends grpc.UntypedServiceImplementation {
  CreatePlaylist: grpc.handleUnaryCall<_service_Empty__Output, _service_Playlist>;
  
  UpdatePlaylst: grpc.handleClientStreamingCall<_service_UpdatePlaylistRequest__Output, _service_Empty>;
  
}

export interface SongsManagementDefinition extends grpc.ServiceDefinition {
  CreatePlaylist: MethodDefinition<_service_Empty, _service_Playlist, _service_Empty__Output, _service_Playlist__Output>
  UpdatePlaylst: MethodDefinition<_service_UpdatePlaylistRequest, _service_Empty, _service_UpdatePlaylistRequest__Output, _service_Empty__Output>
}
