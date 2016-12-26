#!/usr/bin/env ocaml

type direction = Left of int
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
    if Scanf.Scanning.end_of_input buf
    then acc
    else parse_elem ((Scanf.bscanf buf "%1[LR]%d%_[, ]" construct) :: acc)
  in List.rev (parse_elem [])

let rotate (n, e, s, w) = function
  | Left _  -> (e, s, w, n)
  | Right _ -> (w, n, e, s)

let distance = function
  | Left d  -> d
  | Right d -> d

let rec follow directions (x, y) compass =
  let move_to d (n, e, s, w) =
    (x + (e * d) - (w * d),
     y + (n * d) - (s * d)) in
  let move = function
    | []       -> (x, y)
    | d :: ds' ->
        let compass' = rotate compass d in
        let x', y' = move_to (distance d) compass' in
        follow ds' (x', y') compass'
  in move directions

let test path =
  let parsed = parse path in
  let (x, y) = follow parsed (0, 0) (1, 0, 0, 0) in
  let distance = (abs x) + (abs y) in
  Printf.printf "x: %4d, y: %4d, distance: %4d\n" x y distance

let part_one () =
  List.iter test [
      "R2, L3"
    ; "R2, R2, R2"
    ; "R5, L5, R5, R3"
    ; "R1, R3, L2, L5, L2, L1, R3, L4, R2, L2, L4, R2, L1, R1, L2, R3, L1, L4, R2, L5, R3, R4, L1, R2, L1, R3, L4, R5, L4, L5, R5, L3, R2, L3, L3, R1, R3, L4, R2, R5, L4, R1, L1, L1, R5, L2, R1, L2, R188, L5, L3, R5, R1, L2, L4, R3, R5, L3, R3, R45, L4, R4, R72, R2, R3, L1, R1, L1, L1, R192, L1, L1, L1, L4, R1, L2, L5, L3, R5, L3, R3, L4, L3, R1, R4, L2, R2, R3, L5, R3, L1, R1, R4, L2, L3, R1, R3, L4, L3, L4, L2, L2, R1, R3, L5, L1, R4, R2, L4, L1, R3, R3, R1, L5, L2, R4, R4, R2, R1, R5, R5, L4, L1, R5, R3, R4, R5, R3, L1, L2, L4, R1, R4, R5, L2, L3, R4, L4, R2, L2, L4, L2, R5, R1, R4, R3, R5, L4, L4, L5, L5, R3, R4, L1, L3, R2, L2, R1, L3, L5, R5, R5, R3, L4, L2, R4, R5, R1, R4, L3"
    ]

let part_two () =
  ()

let () =
  part_one ();
  part_two ()
