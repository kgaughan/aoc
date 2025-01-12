open Io
open Movement

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
      parse_elem (Scanf.bscanf buf "%1[LR]%d%_[, \n]" construct :: acc)
  in
  parse_elem [] |> List.rev

let rec follow directions (x, y) compass =
  let move_to d (ns, we) = (x + (we * d), y + (ns * d)) in
  let move = function
    | [] -> (x, y)
    | d :: ds ->
        let compass' = rotate compass d in
        let (x', y') = move_to (distance d) compass' in
        follow ds (x', y') compass'
  in
  move directions

let shortest_manhattan_path directions =
  let (x, y) = follow directions (0, 0) (1, 0) in
  abs x + abs y

let find_duplicate_visit directions =
  let visited = Hashtbl.create 10 in
  let rec move_to (x, y) d (ns, we) =
    if d = 0 then
      Either.Right (x, y)
    else
      let (x', y') = (x + we, y + ns) in
      if Hashtbl.mem visited (x', y') then
        Either.Left (x', y')
      else (
        Hashtbl.add visited (x', y') true;
        move_to (x', y') (d - 1) (ns, we))
  in
  let rec move directions (x, y) compass =
    match directions with
    | [] -> 0
    | d :: ds -> (
        let compass' = rotate compass d in
        match move_to (x, y) (distance d) compass' with
        | Either.Left (x', y') -> abs x' + abs y'
        | Either.Right (x', y') -> move ds (x', y') compass')
  in
  move directions (0, 0) (1, 0)

let read filename = read_all filename |> parse
let part_one input = shortest_manhattan_path input |> Printf.printf "Part 1: %d\n%!"
let part_two input = find_duplicate_visit input |> Printf.printf "Part 2: %d\n%!"
