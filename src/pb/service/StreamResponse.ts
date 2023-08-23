// Original file: ../proto/main.proto

import type { StatusResponse as _service_StatusResponse, StatusResponse__Output as _service_StatusResponse__Output } from '../service/StatusResponse';

export interface StreamResponse {
  'volume'?: (number | string);
  'status'?: (_service_StatusResponse | null);
}

export interface StreamResponse__Output {
  'volume'?: (number);
  'status'?: (_service_StatusResponse__Output);
}
