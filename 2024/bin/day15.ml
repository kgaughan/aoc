type cell =
  | Empty
  | Box
  | BoxLeft
  | BoxRight
  | Wall
  | Robot

let read_input path =
  let to_cell = function
    | '#' -> Wall
    | '@' -> Robot
    | 'O' -> Box
    | _ -> Empty (* note: '[' and ']' don't appear in the input *)
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

let expand_map original =
  let height = Array.length original
  and width = Array.length original.(0) in
  let expanded = Array.make_matrix height (2 * width) Empty in
  let replace x y l r =
    expanded.(y).(x * 2) <- l;
    expanded.(y).((x * 2) + 1) <- r
  in
  for y = 0 to height - 1 do
    for x = 0 to width - 1 do
      match original.(y).(x) with
      | Box -> replace x y BoxLeft BoxRight
      | Wall -> replace x y Wall Wall
      | Robot -> replace x y Robot Empty
      | _ -> ()
    done
  done;
  expanded

let find_robot map = Utils.find_cell (fun cell -> cell = Robot) map |> Option.get

let sum_coordinates =
  let to_coordinate x y cell = if cell = Box || cell = BoxLeft then x + (100 * y) else 0 in
  Utils.fold_matrix (fun x y acc cell -> acc + to_coordinate x y cell) 0

let direction = function
  | '^' -> (0, -1)
  | 'v' -> (0, 1)
  | '<' -> (-1, 0)
  | '>' -> (1, 0)
  | _ -> raise (Invalid_argument "Bad direction")

let simulate map movements =
  let (initial_x, initial_y) = find_robot map in
  let x = ref initial_x
  and y = ref initial_y in
  let rec can_push x y dx dy =
    match map.(y + dy).(x + dx) with
    | Empty -> true
    | BoxLeft | BoxRight ->
        let simple_case = can_push (x + dx) (y + dy) dx dy in
        if dy = 0 || map.(y + dy).(x) = map.(y).(x) then
          simple_case
        else
          let offset = if map.(y + dy).(x) = BoxLeft then 1 else -1 in
          simple_case && can_push (x + dx + offset) (y + dy) dx dy
    | Box -> can_push (x + dx) (y + dy) dx dy
    | _ -> false
  in
  let push x y dx dy =
    let rec loop x y =
      (match map.(y + dy).(x + dx) with
      | Box -> loop (x + dx) (y + dy)
      | BoxLeft ->
          loop (x + dx) (y + dy);
          if dy <> 0 then loop (x + dx + 1) (y + dy)
      | BoxRight ->
          loop (x + dx) (y + dy);
          if dy <> 0 then loop (x + dx - 1) (y + dy)
      | _ -> ());
      map.(y + dy).(x + dx) <- map.(y).(x);
      map.(y).(x) <- Empty
    in
    loop x y
  in
  let move ch =
    let (dx, dy) = direction ch in
    match map.(!y + dy).(!x + dx) with
    | Empty ->
        x := !x + dx;
        y := !y + dy
    | BoxLeft | BoxRight | Box ->
        if can_push !x !y dx dy then (
          push !x !y dx dy;
          x := !x + dx;
          y := !y + dy)
    | _ -> ()
  in
  map.(initial_y).(initial_x) <- Empty;
  String.iter move movements

let _ =
  let (map, movements) = read_input "input/day15.txt" in
  simulate map movements;
  Printf.printf "Part 1: %d\n" (sum_coordinates map)

let _ =
  let (map, movements) = read_input "input/day15.txt" in
  let expanded_map = expand_map map in
  simulate expanded_map movements;
  Printf.printf "Part 2: %d\n" (sum_coordinates expanded_map)
