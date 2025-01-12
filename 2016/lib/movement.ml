type direction =
  | Left of int
  | Right of int

let rotate (ns, we) = function
  | Left _ -> (we, -ns)
  | Right _ -> (-we, ns)

let distance = function
  | Left d -> d
  | Right d -> d
