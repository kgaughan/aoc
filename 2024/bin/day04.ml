let read_file path =
  In_channel.with_open_text path In_channel.input_lines

let count_occurrences str grid =
  let height = Array.length grid in
  let width = String.length grid.(0) in
  let max_offset = String.length str in
  let rec match_at (x, y) offset next =
    if x < 0 || x = width || y = height || grid.(y).[x] <> str.[offset] then
      0
    else if offset < max_offset - 1 then
      match_at (next x y) (offset + 1) next
    else
      1
  in
  let count_at x y =
    match_at (x, y) 0 (fun x y -> (x + 1, y)) +
    match_at (x, y) 0 (fun x y -> (x - 1, y + 1)) +
    match_at (x, y) 0 (fun x y -> (x + 1, y + 1)) +
    match_at (x, y) 0 (fun x y -> (x, y  + 1))
  and total = ref 0 in
  for y = 0 to height - 1 do
    for x = 0 to width - 1 do
      total := !total + (count_at x y)
    done
  done;
  !total

let count_crosses grid =
  let height = Array.length grid in
  let width = String.length grid.(0) in
  let is_match x y =
    grid.(y + 1).[x + 1] = 'A' &&
    (grid.(y).[x] = 'M' && grid.(y + 2).[x + 2] = 'S' || grid.(y).[x] = 'S' && grid.(y + 2).[x + 2] = 'M') &&
    (grid.(y + 2).[x] = 'M' && grid.(y).[x + 2] = 'S' || grid.(y + 2).[x] = 'S' && grid.(y).[x + 2] = 'M')
  in
  let total = ref 0 in
  for y = 0 to height - 3 do
    for x = 0 to width - 3 do
      if is_match x y then
        total := !total + 1
    done
  done;
  !total

let _ =
  let grid = read_file "input/day04.txt" |> Array.of_list in
  let part1 = (count_occurrences "XMAS" grid) + (count_occurrences "SAMX" grid) in
  let part2 = count_crosses grid in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
