open Utils

type tile =
  | Empty
  | Corrupted

let read_input path = read_lines path (parse_pair "%d,%d")

let populate_grid width height limit points =
  let grid = Array.make_matrix width height Empty in
  let rec loop n lst =
    if n > 0 then
      match lst with
      | (x, y) :: tl ->
          grid.(y).(x) <- Corrupted;
          loop (n - 1) tl
      | [] -> (grid, [])
    else
      (grid, lst)
  in
  loop limit points

let a_star start goal heuristic grid =
  let sentinel = (-1, -1)
  and width = Array.length grid.(0)
  and height = Array.length grid
  and open_set = IntPairSet.add start IntPairSet.empty
  and closed_set = IntPairSet.empty
  and came_from = IntPairMap.empty
  and g_scores = IntPairMap.add start 0 IntPairMap.empty
  and f_scores = IntPairMap.add start (heuristic start goal) IntPairMap.empty in
  let in_bounds (x, y) = x >= 0 && y >= 0 && x < width && y < height in
  let directions = [(0, -1); (0, 1); (-1, 0); (1, 0)] in
  let find_neighbours (x, y) =
    List.map (fun (dx, dy) -> (x + dx, y + dy)) directions
    |> List.filter (fun (x, y) -> in_bounds (x, y) && grid.(y).(x) <> Corrupted)
  in
  let reconstruct_path came_from current =
    let rec loop acc current =
      let from = IntPairMap.find current came_from in
      if from = start then from :: acc else loop (from :: acc) from
    in
    loop [current] current
  in
  let get_g_score p g_scores =
    match IntPairMap.find_opt p g_scores with
    | Some v -> v
    | None -> max_int
  in
  let score_neighbour current get_weight ((open_set, closed_set, came_from, g_scores, f_scores) as ctx) neighbour =
    if IntPairSet.mem neighbour closed_set then
      ctx
    else
      let tentative_g_score = get_g_score current g_scores + get_weight neighbour in
      if tentative_g_score < get_g_score neighbour g_scores then
        let came_from = IntPairMap.add neighbour current came_from
        and g_scores = IntPairMap.add neighbour tentative_g_score g_scores in
        let f_score = tentative_g_score + heuristic neighbour goal in
        let f_scores = IntPairMap.add neighbour f_score f_scores
        and open_set =
          if not (IntPairSet.mem neighbour open_set) then
            IntPairSet.add neighbour open_set
          else
            open_set
        in
        (open_set, closed_set, came_from, g_scores, f_scores)
      else
        ctx
  in
  let prioritise f_scores p1 p2 =
    if p2 = sentinel then
      p1
    else
      let s1 = IntPairMap.find p1 f_scores
      and s2 = IntPairMap.find p2 f_scores in
      if s1 < s2 then p1 else p2
  in
  let rec loop (open_set, closed_set, came_from, g_scores, f_scores) =
    if IntPairSet.is_empty open_set then
      None
    else
      let current = IntPairSet.fold (prioritise f_scores) open_set sentinel in
      if current = goal then
        Some (reconstruct_path came_from current)
      else
        let open_set = IntPairSet.remove current open_set in
        find_neighbours current
        |> List.fold_left (score_neighbour current (fun _ -> 1)) (open_set, closed_set, came_from, g_scores, f_scores)
        |> loop
  in
  loop (open_set, closed_set, came_from, g_scores, f_scores)

let walk grid =
  let manhattan_distance (x1, y1) (x2, y2) = abs (x1 - x2) + abs (y1 - y2) in
  let goal = (Array.length grid.(0) - 1, Array.length grid - 1) in
  a_star (0, 0) goal manhattan_distance grid

let try_remaining path grid remaining =
  let corrupt (x, y) = grid.(y).(x) <- Corrupted in
  let rec loop visited remaining =
    match remaining with
    | p :: tl ->
        corrupt p;
        if IntPairSet.mem p visited then
          match walk grid with
          | Some path -> loop (IntPairSet.of_list path) tl
          | None -> Some p
        else
          loop visited tl
    | [] -> None
  in
  loop (IntPairSet.of_list path) remaining

let _ =
  let points = read_input "input/day18.txt" in
  let (grid, remaining) = populate_grid 71 71 1024 points in
  let path = walk grid |> Option.get in
  Printf.printf "Part 1: %d\n" (List.length path - 1);
  let coordinate = try_remaining path grid remaining |> Option.get in
  Printf.printf "Part 2: %d,%d\n" (fst coordinate) (snd coordinate)
