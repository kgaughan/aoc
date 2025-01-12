type line =
  | Button of int * int
  | Prize of int * int
  | Empty

let group lst =
  let rec loop acc lst =
    match lst with
    | Button (a1, a2) :: Button (b1, b2) :: Prize (c1, c2) :: tl -> loop ((a1, a2, b1, b2, c1, c2) :: acc) tl
    | Empty :: tl -> loop acc tl
    | [] -> acc
    | _ -> raise (Invalid_argument "WAT")
  in
  loop [] lst |> List.rev

let read_input path =
  Utils.read_lines path (fun line ->
      if String.starts_with ~prefix:"Button" line then
        Scanf.sscanf line "Button %c: X+%d, Y+%d" (fun _ x y -> Button (x, y))
      else if String.starts_with ~prefix:"Prize" line then
        Scanf.sscanf line "Prize: X=%d, Y=%d" (fun x y -> Prize (x, y))
      else
        Empty)
  |> group

let make_monsterous (a1, a2, b1, b2, c1, c2) = (a1, a2, b1, b2, 10000000000000 + c1, 10000000000000 + c2)

let play (a1, a2, b1, b2, c1, c2) =
  match Utils.cramer (a1, a2) (b1, b2) (c1, c2) with
  | Some (x, y) -> Some ((3 * x) + y)
  | None -> None

let _ =
  let machines = read_input "input/day13.txt" in
  let part1 = List.filter_map play machines |> Utils.sum in
  Printf.printf "Part 1: %d\n%!" part1;
  let part2 = List.map make_monsterous machines |> List.filter_map play |> Utils.sum in
  Printf.printf "Part 2: %d\n%!" part2
