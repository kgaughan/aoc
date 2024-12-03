let read_input path =
  let parse_line line = Scanf.sscanf line "%d %d" (fun x y -> (x, y)) in
  let read_lines ic = In_channel.input_lines ic |> List.map parse_line in
  In_channel.with_open_text path read_lines |> List.split

let make_counter_table xs =
  let tbl = Hashtbl.create 1000 in
  let add k =
    match Hashtbl.find_opt tbl k with
    | Some v -> Hashtbl.replace tbl k (v + 1)
    | None -> Hashtbl.add tbl k 1
  in
  List.iter add xs;
  tbl

let _ =
  let lhs, rhs = read_input "input/day01.txt" in
  let part1 = List.fold_left2
      (fun acc a b -> acc + Int.abs(a - b))
      0
      (List.sort compare lhs)
      (List.sort compare rhs)
  in
  let rhs_counters = make_counter_table rhs in
  let similarity n =
    match Hashtbl.find_opt rhs_counters n with
    | None -> 0
    | Some v -> n * v
  in
  let part2 = List.fold_left (fun acc n -> acc + (similarity n)) 0 lhs in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
