let read_input path =
  Utils.read_lines path (fun line -> Scanf.sscanf line "%d %d" (fun x y -> (x, y)))

let _ =
  let lhs, rhs = read_input "input/day01.txt" |> List.split in
  let part1 = List.fold_left2
      (fun acc a b -> acc + Int.abs(a - b))
      0
      (List.sort compare lhs)
      (List.sort compare rhs)
  in
  let rhs_counters = Utils.make_counter_table 1000 rhs in
  let similarity n =
    match Hashtbl.find_opt rhs_counters n with
    | None -> 0
    | Some v -> n * v
  in
  let part2 = List.fold_left (fun acc n -> acc + similarity n) 0 lhs in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
