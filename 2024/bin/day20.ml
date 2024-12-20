let nsew = [(1, 0); (-1, 0); (0, 1); (0, -1)]

let trace start endpoint track =
  (* Assumption: S will be surrounded by walls on all but one side *)
  let find_opening (x, y) directions =
    List.filter_map (fun (dx, dy) -> if track.(y + dy).(x + dx) = '#' then None else Some (dx, dy)) directions
    |> List.hd
  in
  let initial_direction = find_opening start nsew in
  let rec loop here direction acc distance =
    if here = endpoint then
      (here, distance) :: acc
    else
      let d = find_opening here (Utils.turns_90 direction) in
      let next = (fst d + fst here, snd d + snd here) in
      loop next d ((here, distance) :: acc) (distance + 1)
  in
  loop start initial_direction [] 0

let find_adjacents (x, y) width height trail =
  List.filter_map
    (fun (dx, dy) ->
      let (x', y') = (x + (dx * 2), y + (dy * 2)) in
      if x' >= 0 && x < width && y >= 0 && y < height then
        Utils.IntPairMap.find_opt (x', y') trail
      else
        None)
    nsew

let do_part1 width height trail =
  Utils.IntPairMap.fold
    (fun pos distance acc ->
      find_adjacents pos width height trail
      |> List.fold_left (fun acc v -> acc + if v - distance - 2 >= 100 then 1 else 0) acc)
    trail 0

let _ =
  let track = Utils.read_lines "input/day20.txt" (fun line -> String.to_seq line |> Array.of_seq) |> Array.of_list in
  let start = Utils.find_cell (fun ch -> ch = 'S') track |> Option.get in
  let endpoint = Utils.find_cell (fun ch -> ch = 'E') track |> Option.get in
  let points = trace start endpoint track |> List.to_seq |> Utils.IntPairMap.of_seq in
  let part1 = do_part1 (Array.length track.(0)) (Array.length track) points in
  Printf.printf "Part 1: %d\n" part1
