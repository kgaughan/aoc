type maze_element =
  | Wall
  | Empty
  | Visited of int
  | Start
  | End

type direction =
  | North
  | South
  | East
  | West

type reindeer = {
  direction : direction;
  score : int;
  x : int;
  y : int;
}

let read_input path =
  let to_make_element = function
    | '#' -> Wall
    | 'S' -> Start
    | 'E' -> End
    | _ -> Empty
  in
  Utils.read_lines path (Utils.array_of_string to_make_element Empty) |> Array.of_list

let forward = function
  | North -> (0, -1)
  | South -> (0, 1)
  | East -> (1, 0)
  | West -> (-1, 0)

let clockwise = function
  | North -> East
  | South -> West
  | East -> South
  | West -> North

let anticlockwise = function
  | North -> West
  | South -> East
  | East -> North
  | West -> South

let travel start_cell end_cell maze =
  let turn movement r = { r with direction = movement r.direction; score = r.score + 1000 }
  and try_move r =
    let (dx, dy) = forward r.direction in
    let next_pos = { r with score = r.score + 1; x = r.x + dx; y = r.y + dy } in
    match maze.(next_pos.y).(next_pos.x) with
    | Wall -> None
    | Visited score when score <= next_pos.score -> None
    | _ -> Some next_pos
  and paths = Queue.create () in
  let push r =
    if (r.x, r.y) <> end_cell then
      Queue.add r paths;
    maze.(r.y).(r.x) <- Visited r.score
  in
  push { direction = East; score = 0; x = fst start_cell; y = snd start_cell };
  let rec loop () =
    if Queue.length paths = 0 then
      match maze.(snd end_cell).(fst end_cell) with
      | Visited n -> n
      | _ -> 0
    else
      let head = Queue.take paths in
      head |> try_move |> Option.iter push;
      head |> turn anticlockwise |> try_move |> Option.iter push;
      head |> turn clockwise |> try_move |> Option.iter push;
      loop ()
  in
  loop ()

let backtrace start_cell end_cell initial_score maze =
  let turn movement r = { r with direction = movement r.direction; score = r.score - 1000 }
  and try_move r =
    let (dx, dy) = forward r.direction in
    let next_pos = { r with score = r.score - 1; x = r.x + dx; y = r.y + dy } in
    (* The 1000 here looks like a fudge, but it's not: it's to deal with tiles
       with turns *)
    match maze.(next_pos.y).(next_pos.x) with
    | Visited score when next_pos.score = score || next_pos.score - 1000 = score -> Some next_pos
    | _ -> None
  and visited = Hashtbl.create 64
  and paths = Queue.create () in
  let push r =
    if (r.x, r.y) <> start_cell then
      Queue.add r paths;
    match Hashtbl.find_opt visited (r.x, r.y) with
    | Some _ -> ()
    | None -> Hashtbl.add visited (r.x, r.y) true
  in
  push { direction = West; score = initial_score; x = fst end_cell; y = snd end_cell };
  push { direction = South; score = initial_score; x = fst end_cell; y = snd end_cell };
  let rec loop () =
    if Queue.length paths = 0 then
      Hashtbl.length visited
    else
      let head = Queue.take paths in
      head |> try_move |> Option.iter push;
      head |> turn anticlockwise |> try_move |> Option.iter push;
      head |> turn clockwise |> try_move |> Option.iter push;
      loop ()
  in
  loop ()

let _ =
  let maze = read_input "input/day16.txt" in
  let start_cell = Utils.find_cell (fun cell -> cell = Start) maze |> Option.get
  and end_cell = Utils.find_cell (fun cell -> cell = End) maze |> Option.get in
  maze.(snd start_cell).(fst start_cell) <- Empty;
  maze.(snd end_cell).(fst end_cell) <- Empty;
  let score = travel start_cell end_cell maze in
  Printf.printf "Part 1: %d\n" score;
  let in_best_path = backtrace start_cell end_cell score maze in
  Printf.printf "Part 2: %d\n" in_best_path
