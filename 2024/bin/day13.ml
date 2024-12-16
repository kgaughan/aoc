type line =
  | Button of int * int
  | Prize of int * int
  | Empty

let group lst =
  let rec loop acc lst =
    match lst with
    | Button (xa, ya) :: Button (xb, yb) :: Prize (xp, yp) :: tl -> loop ((xa, ya, xb, yb, xp, yp) :: acc) tl
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

(* I know the proper solution here is linear programming, but it's been so long
   since I've done that... *)
let play (xa, ya, xb, yb, xp, yp) =
  (* the upper limits for how often we can use a given button *)
  let max_b = Int.min (xp / xb) (yp / yb) in
  let rec loop nb cost =
    if nb = 0 then
      cost
    else
      let xpr = xp - (nb * xb)
      and ypr = yp - (nb * yb) in
      if xpr mod xa = 0 && ypr mod ya = 0 && xpr / xa = ypr / ya then
        loop (nb - 1) (Int.min (nb + (3 * (xpr / xa))) cost)
      else
        loop (nb - 1) cost
  in
  let result = loop max_b Int.max_int in
  if result = Int.max_int then
    None
  else
    Some result

let _ =
  let machines = read_input "input/day13.txt" in
  let part1 = List.filter_map play machines |> Utils.sum in
  Printf.printf "%d\n" part1
