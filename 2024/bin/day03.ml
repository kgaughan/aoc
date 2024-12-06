type instruction = 
  | Mul of int * int
  | Do
  | Dont

let extract_instructions contents =
  let get_group i = int_of_string (Str.matched_group i contents) in
  let make_instruction () =
    match Str.matched_string contents with
    | "do()" -> Do
    | "don't()" -> Dont
    | _ -> Mul (get_group 1, get_group 2)
  and mul_pat = Str.regexp {|mul(\([0-9]+\),\([0-9]+\))\|do()\|don't()|} in
  Utils.get_all_matches mul_pat contents (fun acc -> make_instruction () :: acc)

let part2 instructions =
  let rec eval instructions acc active =
    match instructions with
    | Mul (a, b) :: tl -> eval tl (if active then acc + (a * b) else acc) active
    | Do :: tl -> eval tl acc true
    | Dont :: tl -> eval tl acc false
    | [] -> acc
  in
  eval instructions 0 true

let part1 instructions =
  List.filter (function Mul _ -> true | _ -> false) instructions |> part2

let _ =
  let instructions = extract_instructions (Utils.read_all "input/day03.txt") in
  Printf.printf "Part 1: %d; Part 2: %d\n" (part1 instructions) (part2 instructions)
