open Utils

let find_guard grid =
  let rec loop y lines =
    match lines with
    | line :: tail -> (
        match String.index_opt line '^' with
        | Some x -> (x, y)
        | None -> loop (y + 1) tail)
    | [] -> raise Not_found
  in
  loop 0 grid

let find_obstructions grid =
  let rec find_obstructions x y lines obstructions =
    match lines with
    | line :: tail -> (
        match String.index_from_opt line x '#' with
        | Some x' -> find_obstructions (x' + 1) y lines (IntPairSet.add (x', y) obstructions)
        | None -> find_obstructions 0 (y + 1) tail obstructions)
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
      match IntPairSet.find_opt next obstructions with
      | Some _ -> loop (gx, gy) (clockwise (dx, dy)) visited'
      | None -> loop next (dx, dy) visited'
  in
  loop guard (0, -1) IntPairSet.empty

module LoopVisit = struct
  type t = {
    guard : IntPair.t;
    direction : IntPair.t;
  }

  let compare v1 v2 =
    match IntPair.compare v1.guard v2.guard with
    | 0 -> IntPair.compare v1.direction v2.direction
    | c -> c
end

module LoopVisitSet = Set.Make (LoopVisit)

let find_loop height width guard obstructions =
  let out_of_bounds (x, y) = x < 0 || y < 0 || x >= width || y >= height in
  let rec loop (gx, gy) (dx, dy) visited =
    let loop_rec = { LoopVisit.guard = (gx, gy); LoopVisit.direction = (dx, dy) } in
    match LoopVisitSet.find_opt loop_rec visited with
    | Some _ -> true
    | None -> (
        let next_pos = (gx + dx, gy + dy) in
        if out_of_bounds next_pos then
          false
        else
          match IntPairSet.find_opt next_pos obstructions with
          | Some _ -> loop (gx, gy) (clockwise (dx, dy)) (LoopVisitSet.add loop_rec visited)
          | None -> loop next_pos (dx, dy) visited)
  in
  loop guard (0, -1) LoopVisitSet.empty

let count_loops height width guard obstructions candidates =
  let check (x, y) acc =
    if (x, y) = guard then
      acc
    else
      match IntPairSet.find (x, y) obstructions with
      | _ -> acc
      | exception Not_found ->
          let obstructions' = IntPairSet.add (x, y) obstructions in
          if find_loop height width guard obstructions' then
            acc + 1
          else
            acc
  in
  IntPairSet.fold check candidates 0

let _ =
  let grid = read_lines "input/day06.txt" (fun line -> line) in
  let height = List.length grid
  and width = String.length (List.hd grid)
  and guard = find_guard grid
  and obstructions = find_obstructions grid in
  let visited = get_visited_cells height width guard obstructions in
  let part1 = IntPairSet.cardinal visited in
  Printf.printf "Part 1: %d\n" part1;
  print_endline "This may take a while...";
  let part2 = count_loops height width guard obstructions visited in
  Printf.printf "Part 2: %d\n" part2
