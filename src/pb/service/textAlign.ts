// Original file: ../proto/main.proto

export const textAlign = {
  left: 0,
  center: 1,
  right: 2,
} as const;

export type textAlign =
  | 'left'
  | 0
  | 'center'
  | 1
  | 'right'
  | 2

export type textAlign__Output = typeof textAlign[keyof typeof textAlign]
