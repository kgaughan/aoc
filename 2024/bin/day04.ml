let count_occurrences str grid =
  let height = Array.length grid
  and width = String.length grid.(0)
  and max_offset = String.length str in
  let rec match_at (x, y) offset (dx, dy) =
    if x < 0 || y < 0 || x >= width || y >= height || grid.(y).[x] <> str.[offset] then
      0
    else if offset < max_offset - 1 then
      match_at (x + dx, y + dy) (offset + 1) (dx, dy)
    else
      1
  in
  let count_at x y = List.map (match_at (x, y) 0) deltas |> Utils.sum
  and total = ref 0 in
  for y = 0 to height - 1 do
    for x = 0 to width - 1 do
      total := !total + count_at x y
    done
  done;
  !total

let count_crosses grid =
  let is_match x y =
    grid.(y + 1).[x + 1] = 'A' &&
    (grid.(y).[x] = 'M' && grid.(y + 2).[x + 2] = 'S' || grid.(y).[x] = 'S' && grid.(y + 2).[x + 2] = 'M') &&
    (grid.(y + 2).[x] = 'M' && grid.(y).[x + 2] = 'S' || grid.(y + 2).[x] = 'S' && grid.(y).[x + 2] = 'M')
  and height = Array.length grid
  and width = String.length grid.(0)
  and total = ref 0 in
  for y = 0 to height - 3 do
    for x = 0 to width - 3 do
      if is_match x y then
        total := !total + 1
    done
  done;
  !total

let _ =
  let grid = Utils.read_lines "input/day04.txt" (fun line -> line) |> Array.of_list in
  let part1 = (count_occurrences "XMAS" grid)
  and part2 = count_crosses grid in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
