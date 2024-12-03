let read_file path =
  let stream = open_in path in
  let to_ints line = String.split_on_char ' ' line |> List.map int_of_string in
  let rec read_line acc =
    match input_line stream with
    | line -> read_line ((to_ints line) :: acc)
    | exception End_of_file ->
      close_in stream;
      acc
  in List.rev (read_line [])

let rec delta = function
  | a :: b :: tl -> (b - a) :: delta (b :: tl)
  | _ -> []

let is_safe readings =
  let deltas = delta readings in
  List.for_all (fun n -> n < 0 && n >= -3) deltas ||
  List.for_all (fun n -> n > 0 && n <= 3) deltas

let is_safe_dampened readings =
  if is_safe readings then
    true
  else
    (* Filter one of the readings at each position and retry *)
    List.mapi (fun i _ -> List.filteri (fun j _ -> i <> j) readings) readings |>
    List.exists is_safe

let check_all check =
  List.fold_left (fun acc n -> if check(n) then acc + 1 else acc) 0

let _ =
  let reports = read_file "input/day02.txt" in
  Printf.printf "Part 1: %d; Part 2: %d\n"
    (check_all is_safe reports)
    (check_all is_safe_dampened reports)
