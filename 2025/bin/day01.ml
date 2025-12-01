let read_input path = Utils.read_lines path (Utils.parse_pair "%c%d")

let _ =
  let to_offsets =
    List.map (function
      | ('R', n) -> n
      | ('L', n) -> -n
      | _ -> 0)
  in
  let rotations = read_input "input/day01.txt" |> to_offsets in
  let ( % ) n clamp =
    let result = n mod clamp in
    if result < 0 then result + clamp else result
  in
  let count_zeroes =
    List.fold_left
      (fun (pos, zeroes) rotation ->
        let pos' = (pos + rotation) % 100 in
        (pos', zeroes + if pos' = 0 then 1 else 0))
      (50, 0)
  in
  let count_zero_clicks =
    List.fold_left
      (fun (pos, zeros) rotation ->
        let pos' = (pos + rotation) % 100
        and modulus = Int.abs ((pos + rotation) / 100) in
        (* there has to be a better ay to phrase this than these two expressions... *)
        let rotations_through_zero = modulus + if pos + rotation < 0 && pos != 0 then 1 else 0 in
        let at_zero = if pos' = 0 && rotations_through_zero = 0 then 1 else 0 in
        (pos', zeros + rotations_through_zero + at_zero))
      (50, 0)
  in
  let part1 = rotations |> count_zeroes |> snd in
  let part2 = rotations |> count_zero_clicks |> snd in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
