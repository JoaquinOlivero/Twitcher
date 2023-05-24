import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { SongsManagementClient as _service_SongsManagementClient, SongsManagementDefinition as _service_SongsManagementDefinition } from './service/SongsManagement';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  service: {
    Empty: MessageTypeDefinition
    Playlist: MessageTypeDefinition
    Song: MessageTypeDefinition
    SongsManagement: SubtypeConstructor<typeof grpc.Client, _service_SongsManagementClient> & { service: _service_SongsManagementDefinition }
  }
}

