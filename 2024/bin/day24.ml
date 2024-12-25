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

let get_keys_with_prefix prefix exprs =
  Hashtbl.to_seq_keys exprs |> Seq.filter (fun key -> String.starts_with ~prefix key)

let extract_wire_value prefix exprs cache =
  get_keys_with_prefix prefix exprs
  |> Seq.map (fun key -> (key, eval key exprs cache))
  |> Seq.fold_left (fun acc (key, value) -> acc lor (value lsl to_index key)) 0

let wire_to_expression key exprs =
  let result = Buffer.create 256 in
  let rec loop key =
    let do_op op k k1 k2 =
      Buffer.add_char result '(';
      Buffer.add_string result k;
      Buffer.add_char result '>';
      Buffer.add_string result op;
      loop k1;
      Buffer.add_char result ' ';
      loop k2;
      Buffer.add_char result ')'
    in
    match Hashtbl.find exprs key with
    | Literal _ -> Buffer.add_string result key
    | Xor (k1, k2) -> do_op "XOR " key k1 k2
    | And (k1, k2) -> do_op "AND " key k1 k2
    | Or (k1, k2) -> do_op "OR " key k1 k2
  in
  loop key;
  Buffer.contents result

(* This sucks, and there are almost certainly much better ways of implementing
   this, but it works. Well, mostly. There is a bug in how z06 in my input is
   handled, because it looks like it's just a series of carries, which means
   it's caught by the logic to handle z45, the final sign/carry bit. Submitted
   the answer with that fudge as I know it's otherwise correct.

   I think an alternative approach might be some stepwise reduction, filling
   in the various bits and reducing things down until backwards from the
   outputs, starting with z00 and progressing through the bits until a
   contradiction is found, though I'll need to be careful aobut the handline of
   carry bits in that case, and I'm not even sure if it'll really work.

   This is one I'll need to revisit. *)
let check_adder key exprs =
  let matches_terminal_wires kl kr i =
    let tk1 = Printf.sprintf "x%02d" i
    and tk2 = Printf.sprintf "y%02d" i in
    (kl, kr) = (tk1, tk2) || (kl, kr) = (tk2, tk1)
  in

  let rec recognise_carry l r parent i =
    let recognise_carry_side k i =
      match Hashtbl.find exprs k with
      | And (l, r) when matches_terminal_wires l r i -> `AndTerminal
      | And (l, r) -> recognise_half_adder l r i
      | _ -> `Bad k
    in
    match (recognise_carry_side l i, recognise_carry_side r i) with
    | (`AndTerminal, `HalfAdder) | (`HalfAdder, `AndTerminal) | (`AndTerminal, `Carry) | (`Carry, `AndTerminal) -> `Or
    | (`Bad k, _) | (_, `Bad k) -> `Bad k
    | _ -> `Bad parent
  and recognise_half_adder l r i =
    let recognise_half_adder_side k i =
      match Hashtbl.find exprs k with
      | Xor (l, r) when matches_terminal_wires l r i -> `Xor
      | And (l, r) when matches_terminal_wires l r (i - 1) -> `And
      | Or (l, r) -> recognise_carry l r k (i - 1)
      | _ -> `Bad k
    in
    match (recognise_half_adder_side l i, recognise_half_adder_side r i) with
    | (`Xor, `And) | (`And, `Xor) -> `HalfAdder
    | (`Xor, `Or) | (`Or, `Xor) -> `Carry
    | (`Bad k', _) | (_, `Bad k') -> `Bad k'
    | _ -> raise (Invalid_argument "Impossible condition recognising half adder!")
  in

  let recognise_root key i =
    match Hashtbl.find exprs key with
    | Xor (l, r) when matches_terminal_wires l r i -> None
    | Xor (l, r) -> (
        match recognise_half_adder l r i with
        | `Bad k -> Some k
        | _ -> None)
    | Or (l, r) -> (
        match recognise_carry l r key (i - 1) with
        | `Bad k -> Some k
        | _ -> None)
    | _ -> Some key
  in
  recognise_root key (to_index key)

let gather_to lst = Option.fold ~none:lst ~some:(fun x -> x :: lst)

let _ =
  let exprs = Hashtbl.create 400
  and cache = Hashtbl.create 400 in
  read_input "input/day24.txt" |> List.to_seq |> Hashtbl.add_seq exprs;
  let z = extract_wire_value "z" exprs cache in
  Printf.printf "Part 1: %d\n" z;
  let wires = get_keys_with_prefix "z" exprs in
  Seq.iter (fun key -> print_endline (wire_to_expression key exprs)) wires;
  Seq.fold_left (fun acc k -> check_adder k exprs |> gather_to acc) [] wires
  |> List.sort_uniq compare |> String.concat "," |> Printf.printf "Part 2: %s\n"
