type cell =
  | Empty
  | Box
  | Wall
  | Robot

let read_input path =
  let to_cell = function
    | '#' -> Wall
    | '@' -> Robot
    | 'O' -> Box
    | _ -> Empty
  in
  let parse_map grid =
    let result = Array.make_matrix (List.length grid) (List.hd grid |> String.length) Empty in
    List.iteri (fun y line -> String.iteri (fun x ch -> result.(y).(x) <- to_cell ch) line) grid;
    result
  in
  let parse lines =
    let (map, movements) = Utils.split_sections lines in
    (parse_map map, String.concat "" movements)
  in
  In_channel.with_open_text path Utils.input_lines |> parse

let find_robot map = Utils.find_cell (fun cell -> cell = Robot) map |> Option.get

let sum_coordinates =
  let to_coordinate x y cell = if cell = Box then x + (100 * y) else 0 in
  Utils.fold_matrix (fun x y acc cell -> acc + to_coordinate x y cell) 0

let simulate map movements =
  let (initial_x, initial_y) = find_robot map in
  let x = ref initial_x
  and y = ref initial_y
  and direction = function
    | '^' -> (0, -1)
    | 'v' -> (0, 1)
    | '<' -> (-1, 0)
    | '>' -> (1, 0)
    | _ -> raise (Invalid_argument "Bad direction")
  in
  let rec try_push x y dx dy =
    match map.(y + dy).(x + dx) with
    | Empty -> Some (x + dx, y + dy)
    | Box -> try_push (x + dx) (y + dy) dx dy
    | _ -> None
  in
  let move ch =
    let (dx, dy) = direction ch in
    match map.(!y + dy).(!x + dx) with
    | Empty ->
        x := !x + dx;
        y := !y + dy
    | Box -> (
        match try_push (!x + dx) (!y + dy) dx dy with
        | Some (empty_x, empty_y) ->
            map.(empty_y).(empty_x) <- map.(!y + dy).(!x + dx);
            map.(!y + dy).(!x + dx) <- Empty;
            x := !x + dx;
            y := !y + dy
        | None -> ())
    | _ -> ()
  in
  map.(initial_x).(initial_y) <- Empty;
  String.iter move movements

let _ =
  let (map, movements) = read_input "input/day15-sample.txt" in
  simulate map movements;
  Printf.printf "Part 1: %d\n" (sum_coordinates map)
