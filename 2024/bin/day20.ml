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

let find_adjacents (x, y) width height trail jumps =
  let visited = Hashtbl.create 1000 in
  List.filter_map
    (fun (dx, dy) ->
      let (x', y') = (x + dx, y + dy) in
      if Hashtbl.mem visited (x', y') then
        None
      else if x' >= 0 && x < width && y >= 0 && y < height then
        match Utils.IntPairMap.find_opt (x', y') trail with
        | Some d ->
            Hashtbl.add visited (x', y') true;
            Some (d - abs dx - abs dy)
        | None -> None
      else
        None)
    jumps

let manhattan_circle i r =
  let rec loop r offset acc =
    if offset = 0 then
      acc
    else
      loop r (offset - 1)
        ((offset, r - offset) :: (r - offset, -offset) :: (-offset, offset - r) :: (offset - r, offset) :: acc)
  in
  let rec fill width acc =
    if width = i then
      acc
    else
      fill (width - 1) (loop width width acc)
  in
  fill r []

let race dist threshold width height trail =
  let jumps = manhattan_circle 1 dist in
  Utils.IntPairMap.fold
    (fun pos distance acc ->
      find_adjacents pos width height trail jumps
      |> List.fold_left (fun acc v -> acc + if v - distance >= threshold then 1 else 0) acc)
    trail 0

let _ =
  let track = Utils.read_lines "input/day20.txt" (fun line -> String.to_seq line |> Array.of_seq) |> Array.of_list in
  let start = Utils.find_cell (fun ch -> ch = 'S') track |> Option.get in
  let endpoint = Utils.find_cell (fun ch -> ch = 'E') track |> Option.get in
  let points = trace start endpoint track |> List.to_seq |> Utils.IntPairMap.of_seq in
  let part1 = race 2 100 (Array.length track.(0)) (Array.length track) points in
  let part2 = race 20 100 (Array.length track.(0)) (Array.length track) points in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
