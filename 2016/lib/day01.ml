type direction =
  | Left of int
  | Right of int

exception Unexpected_direction

let parse path =
  let buf = Scanf.Scanning.from_string path in
  let construct direction distance =
    match direction with
    | "L" -> Left distance
    | "R" -> Right distance
    | _ -> raise Unexpected_direction
  in
  let rec parse_elem acc =
    if Scanf.Scanning.end_of_input buf then
      acc
    else
      parse_elem (Scanf.bscanf buf "%1[LR]%d%_[, ]" construct :: acc)
  in
  List.rev (parse_elem [])

let rotate (n, e, s, w) = function
  | Left _ -> (e, s, w, n)
  | Right _ -> (w, n, e, s)

let distance = function
  | Left d -> d
  | Right d -> d

let rec follow directions (x, y) compass =
  let move_to d (n, e, s, w) = (x + (e * d) - (w * d), y + (n * d) - (s * d)) in
  let move = function
    | [] -> (x, y)
    | d :: ds' ->
        let compass' = rotate compass d in
        let (x', y') = move_to (distance d) compass' in
        follow ds' (x', y') compass'
  in
  move directions

let test path =
  let parsed = parse path in
  let (x, y) = follow parsed (0, 0) (1, 0, 0, 0) in
  let distance = abs x + abs y in
  Printf.printf "x: %4d, y: %4d, distance: %4d\n" x y distance

let part_one () = List.iter test ["R2, L3"; "R2, R2, R2"; "R5, L5, R5, R3"]
let part_two () = ()
