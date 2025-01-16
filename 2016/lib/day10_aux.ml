type recipient =
  | Output of int
  | Bot of int

type action =
  | Receive of int * recipient
  | Give of int * recipient * recipient
