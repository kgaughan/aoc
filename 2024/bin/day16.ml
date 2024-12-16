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
  let make_maze_line line =
    let result = Array.make (String.length line) Empty in
    String.iteri
      (fun i ch ->
        result.(i) <-
          (match ch with
          | '#' -> Wall
          | 'S' -> Start
          | 'E' -> End
          | _ -> Empty))
      line;
    result
  in
  let lines = Utils.read_lines path make_maze_line |> Array.of_list in
  lines

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

let dump maze =
  let height = Array.length maze
  and width = Array.length maze.(0) in
  for y = 0 to height - 1 do
    for x = 0 to width - 1 do
      print_char
        (match maze.(y).(x) with
        | Wall -> '#'
        | Empty -> '.'
        | Visited _ -> 'v'
        | Start -> 'S'
        | End -> 'E')
    done;
    print_newline ()
  done

let travel start_cell end_cell maze =
  let turn movement r = { direction = movement r.direction; score = r.score + 1000; x = r.x; y = r.y }
  and try_move r =
    let (dx, dy) = forward r.direction in
    let next_pos = { direction = r.direction; score = r.score + 1; x = r.x + dx; y = r.y + dy } in
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

let _ =
  let maze = read_input "input/day16.txt" in
  let start_cell = Utils.find_cell (fun cell -> cell = Start) maze |> Option.get
  and end_cell = Utils.find_cell (fun cell -> cell = End) maze |> Option.get in
  maze.(snd start_cell).(fst start_cell) <- Empty;
  maze.(snd end_cell).(fst end_cell) <- Empty;
  let score = travel start_cell end_cell maze in
  Printf.printf "Part 1: %d\n" score
