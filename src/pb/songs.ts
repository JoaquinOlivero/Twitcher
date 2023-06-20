import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { StreamManagementClient as _service_StreamManagementClient, StreamManagementDefinition as _service_StreamManagementDefinition } from './service/StreamManagement';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  service: {
    AudioResponse: MessageTypeDefinition
    Empty: MessageTypeDefinition
    OutputResponse: MessageTypeDefinition
    SDP: MessageTypeDefinition
    Song: MessageTypeDefinition
    SongPlaylist: MessageTypeDefinition
    StreamManagement: SubtypeConstructor<typeof grpc.Client, _service_StreamManagementClient> & { service: _service_StreamManagementDefinition }
  }
}

