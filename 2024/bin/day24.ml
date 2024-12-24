type expr =
  | Literal of int
  | Xor of string * string
  | And of string * string
  | Or of string * string

let read_input path =
  let to_expr line =
    let (a1, op, a2, dest) = Scanf.sscanf line "%s %s %s -> %s" (fun a1 op a2 dest -> (a1, op, a2, dest)) in
    match op with
    | "XOR" -> (dest, Xor (a1, a2))
    | "AND" -> (dest, And (a1, a2))
    | "OR" -> (dest, Or (a1, a2))
    | _ -> raise (Invalid_argument (Printf.sprintf "Bad operator: %s" op))
  in
  let to_literal line =
    let (var, value) = Utils.parse_pair "%s@: %d" line in
    (var, Literal value)
  in
  let parse lines =
    let (inputs, rules) = Utils.split_sections lines in
    List.append (List.rev_map to_literal inputs) (List.rev_map to_expr rules)
  in
  In_channel.with_open_text path Utils.input_lines |> parse

let to_index key = int_of_string (String.sub key 1 (String.length key - 1))

let rec eval key exprs cache =
  match Hashtbl.find_opt cache key with
  | Some v -> v
  | None ->
      let result =
        match Hashtbl.find exprs key with
        | Literal v -> v
        | Xor (k1, k2) -> eval k1 exprs cache lxor eval k2 exprs cache
        | And (k1, k2) -> eval k1 exprs cache land eval k2 exprs cache
        | Or (k1, k2) -> eval k1 exprs cache lor eval k2 exprs cache
      in
      Hashtbl.add cache key result;
      result

let extract_wire_value prefix exprs cache =
  let keys = Hashtbl.to_seq_keys exprs |> Seq.filter (fun key -> String.starts_with ~prefix key) |> List.of_seq in
  let bits = List.rev_map (fun key -> (key, eval key exprs cache)) keys in
  List.fold_left (fun acc (key, value) -> acc lor (value lsl to_index key)) 0 bits

let _ =
  let exprs = Hashtbl.create 400
  and cache = Hashtbl.create 400 in
  read_input "input/day24.txt" |> List.to_seq |> Hashtbl.add_seq exprs;
  let z = extract_wire_value "z" exprs cache in
  Printf.printf "Part 1: %d\n" z
