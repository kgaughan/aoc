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

(* Cramer's rule for solving a series of linear equations, as cogged from
   https://www.youtube.com/watch?v=vXqlIOX2itM - TBH, I don't entirely
   understand this anymore, and it's pretty clear that I need to go back and
   revise how systems of linear equations work.

   Also:
    * https://en.wikipedia.org/wiki/System_of_linear_equations
    * https://en.wikipedia.org/wiki/Cramer%27s_rule
    * https://en.wikipedia.org/wiki/Gaussian_elimination
*)
let play (a1, a2, b1, b2, c1, c2) =
  let determiner = (a1 * b2) - (b1 * a2) in
  let x = ((c1 * b2) - (b1 * c2)) / determiner
  and y = ((a1 * c2) - (c1 * a2)) / determiner in
  if (a1 * x) + (b1 * y) = c1 && (a2 * x) + (b2 * y) = c2 then
    Some ((3 * x) + y)
  else
    None

let _ =
  let machines = read_input "input/day13.txt" in
  let part1 = List.filter_map play machines |> Utils.sum in
  Printf.printf "Part 1: %d\n" part1;
  let part2 = List.map make_monsterous machines |> List.filter_map play |> Utils.sum in
  Printf.printf "Part 2: %d\n" part2
