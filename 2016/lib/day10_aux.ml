type recipient =
  | Output of int
  | Bot of int

type action =
  | Receive of int * int
  | Give of int * recipient * recipient
