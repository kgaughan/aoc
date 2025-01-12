exception Bad_direction

let direction = function
  | 'U' -> (0, -1)
  | 'D' -> (0, 1)
  | 'R' -> (1, 0)
  | 'L' -> (-1, 0)
  | _ -> raise Bad_direction

let rec follow_line to_digit (x, y) = function
  | [] -> (x, y)
  | d :: ds ->
      let (dx, dy) = direction d in
      let x' = x + dx in
      let y' = y + dy in
      let loc = if to_digit (x', y') = None then (x, y) else (x', y') in
      follow_line to_digit loc ds

(* A three by three grid *)
let to_digit_grid = function
  | (0, 0) -> Some '1'
  | (1, 0) -> Some '2'
  | (2, 0) -> Some '3'
  | (0, 1) -> Some '4'
  | (1, 1) -> Some '5'
  | (2, 1) -> Some '6'
  | (0, 2) -> Some '7'
  | (1, 2) -> Some '8'
  | (2, 2) -> Some '9'
  | _ -> None

(* A diamond-shaped grid *)
let to_digit_diamond = function
  | (2, 0) -> Some '1'
  | (1, 1) -> Some '2'
  | (2, 1) -> Some '3'
  | (3, 1) -> Some '4'
  | (0, 2) -> Some '5'
  | (1, 2) -> Some '6'
  | (2, 2) -> Some '7'
  | (3, 2) -> Some '8'
  | (4, 2) -> Some '9'
  | (1, 3) -> Some 'A'
  | (2, 3) -> Some 'B'
  | (3, 3) -> Some 'C'
  | (2, 4) -> Some 'D'
  | _ -> None

let charlist_of_string str = Array.to_list (Array.init (String.length str) (String.get str))

let follow_set initial to_digit lines =
  let rec follow_set' (x, y) lines acc =
    match lines with
    | [] -> List.rev acc
    | l :: ls ->
        let (x', y') = follow_line to_digit (x, y) (charlist_of_string l) in
        follow_set' (x', y') ls ((x', y') :: acc)
  in
  (* (1, 1) is the '5' position *)
  follow_set' initial lines []

let harness initial to_digit lines =
  let print_digit pos =
    match to_digit pos with
    | Some ch -> print_char ch
    | None -> ()
  in
  List.iter print_digit (follow_set initial to_digit lines);
  print_newline ()

let read = Io.read_lines

let part_one input =
  print_string "Part 1: ";
  harness (1, 1) to_digit_grid input

let part_two input =
  print_string "Part 2: ";
  harness (0, 2) to_digit_diamond input
