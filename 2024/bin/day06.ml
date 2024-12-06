open Utils

let find_guard grid =
  let rec loop y lines =
    match lines with
    | line :: tail -> (
        match String.index line '^' with
        | x -> (x, y)
        | exception Not_found -> loop (y + 1) tail)
    | [] -> raise Not_found
  in
  loop 0 grid

let find_obstructions grid =
  let rec find_obstructions x y lines obstructions =
    match lines with
    | line :: tail -> (
        match String.index_from line x '#' with
        | x' -> find_obstructions (x' + 1) y lines (IntPairSet.add (x', y) obstructions)
        | exception Not_found -> find_obstructions 0 (y + 1) tail obstructions)
    | [] -> obstructions
  in
  find_obstructions 0 0 grid IntPairSet.empty

let clockwise = function
  | (0, -1) -> (1, 0)
  | (1, 0) -> (0, 1)
  | (0, 1) -> (-1, 0)
  | (-1, 0) -> (0, -1)
  | _ -> raise (Invalid_argument "bad delta")

let get_visited_cells height width guard obstructions =
  let out_of_bounds (x, y) = x < 0 || y < 0 || x >= width || y >= height in
  let rec loop (gx, gy) (dx, dy) visited =
    let visited' = IntPairSet.add (gx, gy) visited
    and next = (gx + dx, gy + dy) in
    if out_of_bounds next then
      visited'
    else
      match IntPairSet.find next obstructions with
      | _ -> loop (gx, gy) (clockwise (dx, dy)) visited'
      | exception Not_found -> loop next (dx, dy) visited'
  in
  loop guard (0, -1) IntPairSet.empty

let _ =
  let grid = read_lines "input/day06.txt" (fun line -> line) in
  let height = List.length grid
  and width = String.length (List.hd grid)
  and guard = find_guard grid
  and obstructions = find_obstructions grid in
  let part1 = get_visited_cells height width guard obstructions |> IntPairSet.cardinal in
  Printf.printf "Part 1: %d\n" part1
