type robot = {
  px : int;
  py : int;
  vx : int;
  vy : int;
}

let read_input path =
  Utils.read_lines path (fun line -> Scanf.sscanf line "p=%d,%d v=%d,%d" (fun x y vx vy -> { px = x; py = y; vx; vy }))

let ( % ) x y =
  let result = x mod y in
  if result >= 0 then
    result
  else
    result + y

(* As a rough estimate of the amount of entropy in a grid, this does a
   simplified run-length encode of the data by counting how often we flip
   between zero and non-zero in the grid. The result is the number of tokens
   and RLE encoder would spit out. *)
let rle_entropy grid =
  let height = Array.length grid
  and width = Array.length grid.(0)
  and previous = ref (grid.(0).(0) > 0)
  and flips = ref 0 in
  for y = 0 to height - 1 do
    for x = 0 to width - 1 do
      if grid.(y).(x) > 0 <> !previous then (
        previous := not !previous;
        flips := !flips + 1)
    done
  done;
  float_of_int !flips /. float_of_int (width * height)

let as_grid width height robots =
  let grid = Array.make_matrix height width 0 in
  List.iter (fun r -> grid.(r.py).(r.px) <- grid.(r.py).(r.px) + 1) robots;
  grid

let render ?(skip = false) grid =
  let height = Array.length grid
  and width = Array.length grid.(0) in
  for y = 0 to height - 1 do
    if (not skip) || y <> height / 2 then
      for x = 0 to width - 1 do
        if skip && x = width / 2 then
          print_char ' '
        else
          match grid.(y).(x) with
          | 0 -> print_char '.'
          | n -> print_int n
      done;
    print_newline ()
  done

let simulate_once width height =
  List.map (fun r -> { px = (r.px + r.vx) % width; py = (r.py + r.vy) % height; vx = r.vx; vy = r.vy })

let simulate fn seconds robots =
  let rec loop seconds robots =
    if seconds = 0 then
      robots
    else
      loop (seconds - 1) (fn robots)
  in
  loop seconds robots

let get_safety_factor width height robots =
  let get_quadrant x y =
    let xq = if x < width / 2 then Some 0 else if x > width / 2 then Some 1 else None
    and yq = if y < height / 2 then Some 0 else if y > height / 2 then Some 1 else None in
    match (xq, yq) with
    | (Some h, Some w) -> Some (h + (w * 2))
    | _ -> None
  in
  let quadrants = Array.make 4 0 in
  List.iter
    (fun r ->
      match get_quadrant r.px r.py with
      | Some q -> quadrants.(q) <- quadrants.(q) + 1
      | None -> ())
    robots;
  Array.fold_left ( * ) 1 quadrants

let _ =
  let width = 101
  and height = 103 in
  let robots = read_input "input/day14.txt" in
  let result = simulate (simulate_once width height) 100 robots in
  let part1 = get_safety_factor width height result in
  Printf.printf "Part 1: %d\n%!" part1

let _ =
  let width = 101
  and height = 103 in
  let robots = read_input "input/day14.txt" in
  let rec loop robots attempts =
    let result = simulate_once width height robots in
    let grid = as_grid width height result in
    let entropy = rle_entropy grid in
    if entropy < 0.05 then (
      Printf.printf "Part 2: %d -> %f\n" (attempts + 1) entropy;
      render grid);
    if attempts < 10000 then
      loop result (attempts + 1)
  in
  loop robots 0
