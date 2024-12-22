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

let general_strategy buf dx dy =
  if dx < 0 then (
    gen_move '<' '>' dx buf;
    gen_move '^' 'v' dy buf)
  else (
    gen_move '^' 'v' dy buf;
    gen_move '<' '>' dx buf)

let convolute_numpad =
  let strategy buf ((x, y), (dx, dy)) =
    (match ((x, y), (x + dx, y + dy)) with
    | ((0, y'), (x', 3)) when y' < 3 && x' > 0 ->
        gen_move '<' '>' dx buf;
        gen_move '^' 'v' dy buf
    | ((x', 3), (0, y')) when y' < 3 && x' > 0 ->
        gen_move '^' 'v' dy buf;
        gen_move '<' '>' dx buf
    | _ -> general_strategy buf dx dy);
    Buffer.add_char buf 'A'
  in
  convolute numpad (2, 3) strategy

let convolute_dpad =
  let strategy buf ((x, y), (dx, dy)) =
    (match ((x, y), (x + dx, y + dy)) with
    | ((0, _), (_, _)) ->
        gen_move '<' '>' dx buf;
        gen_move '^' 'v' dy buf
    | ((_, _), (0, _)) ->
        gen_move '^' 'v' dy buf;
        gen_move '<' '>' dx buf
    | _ -> general_strategy buf dx dy);
    Buffer.add_char buf 'A'
  in
  convolute dpad (2, 0) strategy

let trim_last ?(n = 1) s = String.sub s 0 (String.length s - n)

let recursive_convolute_dpad n s =
  (* Breaks a set of moves into a list we can process *)
  let chunks s = trim_last s |> String.split_on_char 'A' |> List.map (fun s -> String.cat s "A") in
  let cache = Hashtbl.create 100000 in
  let rec loop n s =
    if n = 0 then
      String.length s
    else
      let parts = chunks s in
      List.fold_left
        (fun acc chunk ->
          match Hashtbl.find_opt cache (chunk, n) with
          | Some result -> acc + result
          | None ->
              let result = loop (n - 1) (convolute_dpad chunk) in
              Hashtbl.add cache (chunk, n) result;
              acc + result)
        0 parts
  in
  loop n s

let get_complexity codes results = List.map2 (fun code result -> code * String.length result) codes results |> Utils.sum
let get_complexity2 codes results = List.map2 ( * ) codes results |> Utils.sum

let _ =
  let codes = Utils.read_lines "input/day21.txt" Fun.id in
  let convoluted_codes = List.map convolute_numpad codes in
  let numeric_codes = List.map (fun code -> int_of_string (trim_last code)) codes in
  let part1 =
    List.map (fun code -> convolute_dpad code |> convolute_dpad) convoluted_codes |> get_complexity numeric_codes
  in
  let part2 =
    List.map (fun code -> recursive_convolute_dpad 25 code) convoluted_codes |> get_complexity2 numeric_codes
  in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
