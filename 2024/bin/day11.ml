let has_even_number_of_digits n = int_of_float (log10 (float_of_int n)) mod 2 = 1

let split_number n =
  let as_string = string_of_int n in
  let half = String.length as_string / 2 in
  (int_of_string (String.sub as_string 0 half), int_of_string (String.sub as_string half half))

let process n tiles =
  let cache = Hashtbl.create 100000 in
  let rec transform tile n =
    if n = 0 then
      1
    else
      match Hashtbl.find_opt cache (tile, n) with
      | Some r -> r
      | None ->
          let result =
            if tile = 0 then
              transform 1 (n - 1)
            else if has_even_number_of_digits tile && tile > 9 then
              let (a, b) = split_number tile in
              transform a (n - 1) + transform b (n - 1)
            else
              transform (tile * 2024) (n - 1)
          in
          Hashtbl.add cache (tile, n) result;
          result
  in
  List.map (fun tile -> transform tile n) tiles |> Utils.sum

let _ =
  let tiles = Utils.read_line "input/day11.txt" |> String.split_on_char ' ' |> List.map int_of_string in
  let part1 = process 25 tiles in
  let part2 = process 75 tiles in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
