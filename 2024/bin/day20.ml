let trace start endpoint track =
  (* Assumption: S will be surrounded by walls on all but one side *)
  let find_opening (x, y) directions =
    List.filter_map (fun (dx, dy) -> if track.(y + dy).(x + dx) = '#' then None else Some (dx, dy)) directions
    |> List.hd
  in
  let initial_direction = find_opening start [(1, 0); (-1, 0); (0, 1); (0, -1)] in
  let rec loop here direction acc distance =
    if here = endpoint then
      (here, distance) :: acc
    else
      let d = find_opening here (Utils.turns_90 direction) in
      let next = (fst d + fst here, snd d + snd here) in
      loop next d ((here, distance) :: acc) (distance + 1)
  in
  loop start initial_direction [] 0

let race jump_radius threshold trail =
  let jumps = Utils.manhattan_circle 1 jump_radius in
  let find_adjacents_past_threshold (x, y) distance acc =
    List.fold_left
      (fun acc (dx, dy) ->
        match Utils.IntPairMap.find_opt (x + dx, y + dy) trail with
        | Some d -> acc + if d - abs dx - abs dy - distance >= threshold then 1 else 0
        | None -> acc)
      acc jumps
  in
  Utils.IntPairMap.fold find_adjacents_past_threshold trail 0

let _ =
  let track = Utils.read_lines "input/day20.txt" (fun line -> String.to_seq line |> Array.of_seq) |> Array.of_list in
  let start = Utils.find_cell (fun ch -> ch = 'S') track |> Option.get in
  let endpoint = Utils.find_cell (fun ch -> ch = 'E') track |> Option.get in
  let points = trace start endpoint track |> List.to_seq |> Utils.IntPairMap.of_seq in
  let part1 = Utils.time "part 1" (fun () -> race 2 100 points) in
  let part2 = Utils.time "part 2" (fun () -> race 20 100 points) in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
