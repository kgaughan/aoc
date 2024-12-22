let numpad =
  [
    ('A', (2, 3));
    ('0', (1, 3));
    ('1', (0, 2));
    ('2', (1, 2));
    ('3', (2, 2));
    ('4', (0, 1));
    ('5', (1, 1));
    ('6', (2, 1));
    ('7', (0, 0));
    ('8', (1, 0));
    ('9', (2, 0));
  ]

let dpad = [('^', (1, 0)); ('A', (2, 0)); ('<', (0, 1)); ('v', (1, 1)); ('>', (2, 1))]

let gen_move back forth n buf =
  if n > 0 then
    for _ = 0 to n - 1 do
      Buffer.add_char buf forth
    done
  else if n < 0 then
    for _ = 0 downto n + 1 do
      Buffer.add_char buf back
    done

let convolute mapping initial_pos strategy code =
  let to_dpad moves =
    let buf = Buffer.create (List.length moves + 1) in
    List.iter (strategy buf) moves;
    Buffer.contents buf
  in
  let rec loop i (x, y) acc =
    if i = String.length code then
      acc
    else
      let (x', y') = List.assoc code.[i] mapping in
      let (dx, dy) = (x' - x, y' - y) in
      loop (i + 1) (x', y') (((x, y), (dx, dy)) :: acc)
  in
  loop 0 initial_pos [] |> List.rev |> to_dpad

let convolute_numpad =
  let strategy buf ((x, y), (dx, dy)) =
    (match ((x, y), (x + dx, y + dy)) with
    | ((0, y'), (x', 3)) when y' < 3 && x' > 0 ->
        gen_move '<' '>' dx buf;
        gen_move '^' 'v' dy buf
    | ((x', 3), (0, y')) when y' < 3 && x' > 0 ->
        gen_move '^' 'v' dy buf;
        gen_move '<' '>' dx buf
    | _ ->
        if dx < 0 then (
          gen_move '<' '>' dx buf;
          gen_move '^' 'v' dy buf)
        else (
          gen_move '^' 'v' dy buf;
          gen_move '<' '>' dx buf));
    Buffer.add_char buf 'A'
  in
  convolute numpad (2, 3) strategy

let convolute_dpad =
  let strategy buf ((x, y), (dx, dy)) =
    if x + dx > 0 && dx > 1 then (
      gen_move '^' 'v' dy buf;
      gen_move '<' '>' dx buf)
    else if x + dx > 0 then (
      gen_move '<' '>' dx buf;
      gen_move '^' 'v' dy buf)
    else (
      gen_move '^' 'v' dy buf;
      gen_move '<' '>' dx buf);
    Buffer.add_char buf 'A'
  in
  convolute dpad (2, 0) strategy

let dump codes results = List.iter2 (fun code result -> Printf.printf "%s: %s\n" code result) codes results

let get_complexity codes results =
  List.map2
    (fun code result -> int_of_string (String.sub code 0 (String.length code - 1)) * String.length result)
    codes results

let _ =
  let codes = Utils.read_lines "input/day21.txt" Fun.id in
  let part1 =
    List.map (fun code -> convolute_numpad code |> convolute_dpad |> convolute_dpad) codes
    |> get_complexity codes |> Utils.sum
  in
  Printf.printf "Part 1: %d\n" part1
