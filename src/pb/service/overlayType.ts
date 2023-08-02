// Original file: ../proto/main.proto

export const overlayType = {
  img: 0,
  textbox: 1,
} as const;

export type overlayType =
  | 'img'
  | 0
  | 'textbox'
  | 1

export type overlayType__Output = typeof overlayType[keyof typeof overlayType]
