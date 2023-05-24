// Original file: ../proto/songs.proto

import type { Song as _service_Song, Song__Output as _service_Song__Output } from '../service/Song';

export interface Playlist {
  'songs'?: (_service_Song)[];
}

export interface Playlist__Output {
  'songs'?: (_service_Song__Output)[];
}
