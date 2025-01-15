type operation =
  | Rect of (int * int)
  | RotateRow of (int * int)
  | RotateColumn of (int * int)
