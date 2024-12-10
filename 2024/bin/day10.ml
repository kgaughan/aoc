let find_trailheads grid =
  let seek_y (y, acc) line =
    let seek_x (x, acc) ch = (x + 1, if ch = '0' then (x, y) :: acc else acc) in
    (y + 1, snd (String.fold_left seek_x (0, acc) line))
  in
  snd (Array.fold_left seek_y (0, []) grid)

let follow_trail grid (x, y) =
  let width = String.length grid.(0)
  and height = Array.length grid in
  let get_altitude x y = Utils.parse_digit grid.(y).[x] in
  let directions = [(1, 0); (0, 1); (-1, 0); (0, -1)] in
  let rec walk (x, y) =
    let altitude = get_altitude x y in
    let is_valid (dx, dy) =
      let (x', y') = (x + dx, y + dy) in
      x' >= 0 && y' >= 0 && x' < width && y' < height && get_altitude x' y' = altitude + 1
    in
    if altitude = 9 then
      [(x, y)]
    else
      List.filter is_valid directions |> List.map (fun (dx, dy) -> walk (x + dx, y + dy)) |> List.flatten
  in
  walk (x, y)

let _ =
  let topology = Utils.read_lines "input/day10.txt" (fun x -> x) |> Array.of_list in
  let trailheads = find_trailheads topology in
  let paths = List.map (follow_trail topology) trailheads in
  let part1 = List.map (List.sort_uniq compare) paths |> List.map List.length |> Utils.sum in
  let part2 = List.map List.length paths |> Utils.sum in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
