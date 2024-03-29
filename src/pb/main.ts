import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { MainClient as _service_MainClient, MainDefinition as _service_MainDefinition } from './service/Main';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  google: {
    protobuf: {
      Empty: MessageTypeDefinition
    }
  }
  service: {
    AudioResponse: MessageTypeDefinition
    BackgroundVideo: MessageTypeDefinition
    BackgroundVideosResponse: MessageTypeDefinition
    DevCredentials: MessageTypeDefinition
    Main: SubtypeConstructor<typeof grpc.Client, _service_MainClient> & { service: _service_MainDefinition }
    Overlay: MessageTypeDefinition
    Overlays: MessageTypeDefinition
    SDP: MessageTypeDefinition
    SaveStreamParametersRequest: MessageTypeDefinition
    Song: MessageTypeDefinition
    SongPlaylist: MessageTypeDefinition
    StatusNCSResponse: MessageTypeDefinition
    StatusResponse: MessageTypeDefinition
    StreamParametersResponse: MessageTypeDefinition
    StreamResponse: MessageTypeDefinition
    TwitchStreamKey: MessageTypeDefinition
    UploadVideoRequest: MessageTypeDefinition
    UploadVideoResponse: MessageTypeDefinition
    UserAuth: MessageTypeDefinition
    VideoInfo: MessageTypeDefinition
  }
}

