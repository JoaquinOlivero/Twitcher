// Original file: ../proto/main.proto

import type { VideoInfo as _service_VideoInfo, VideoInfo__Output as _service_VideoInfo__Output } from '../service/VideoInfo';

export interface UploadVideoRequest {
  'info'?: (_service_VideoInfo | null);
  'chunk'?: (Buffer | Uint8Array | string);
  'data'?: "info"|"chunk";
}

export interface UploadVideoRequest__Output {
  'info'?: (_service_VideoInfo__Output);
  'chunk'?: (Buffer);
}
