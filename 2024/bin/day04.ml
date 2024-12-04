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

let _ =
  let grid = read_file "input/day04.txt" |> Array.of_list in
  let count = (count_occurrences "XMAS" grid) + (count_occurrences "SAMX" grid) in
  Printf.printf "Part 1: %d\n" count
