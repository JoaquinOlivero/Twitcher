// Original file: ../proto/songs.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { Empty as _service_Empty, Empty__Output as _service_Empty__Output } from '../service/Empty';
import type { Playlist as _service_Playlist, Playlist__Output as _service_Playlist__Output } from '../service/Playlist';

export interface SongsManagementClient extends grpc.Client {
  CreatePlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  CreatePlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  CreatePlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  CreatePlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, metadata: grpc.Metadata, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, options: grpc.CallOptions, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  createPlaylist(argument: _service_Empty, callback: grpc.requestCallback<_service_Playlist__Output>): grpc.ClientUnaryCall;
  
}

export interface SongsManagementHandlers extends grpc.UntypedServiceImplementation {
  CreatePlaylist: grpc.handleUnaryCall<_service_Empty__Output, _service_Playlist>;
  
}

export interface SongsManagementDefinition extends grpc.ServiceDefinition {
  CreatePlaylist: MethodDefinition<_service_Empty, _service_Playlist, _service_Empty__Output, _service_Playlist__Output>
}
