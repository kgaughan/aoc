let has_even_number_of_digits n = int_of_float (log10 (float_of_int n)) mod 2 = 1

let split_number n =
  let as_string = string_of_int n in
  let half = String.length as_string / 2 in
  (int_of_string (String.sub as_string 0 half), int_of_string (String.sub as_string half half))

let rec process n tiles =
  let rec transform tiles acc =
    match tiles with
    | 0 :: tl -> transform tl (1 :: acc)
    | n :: tl when has_even_number_of_digits n ->
        let (a, b) = split_number n in
        transform tl (a :: b :: acc)
    | n :: tl -> transform tl ((n * 2024) :: acc)
    | [] -> acc
  in
  if n = 0 then
    tiles
  else
    process (n - 1) (transform tiles [])

let _ =
  let tiles = Utils.read_line "input/day11.txt" |> String.split_on_char ' ' |> List.map int_of_string in
  let part1 = process 25 tiles in
  let cache = Hashtbl.create 100000 in
  let part2 =
    List.map
      (fun n ->
        match Hashtbl.find_opt cache n with
        | Some r -> r
        | None ->
            let r = process 50 [n] |> List.length in
            Hashtbl.add cache n r;
            r)
      part1
    |> Utils.sum
  in
  Printf.printf "Part 1: %d; Part 2: %d\n" (List.length part1) part2
