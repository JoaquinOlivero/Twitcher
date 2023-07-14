// Original file: ../proto/main.proto

import type { Song as _service_Song, Song__Output as _service_Song__Output } from '../service/Song';

export interface SongPlaylist {
  'songs'?: (_service_Song)[];
}

export interface SongPlaylist__Output {
  'songs'?: (_service_Song__Output)[];
}
