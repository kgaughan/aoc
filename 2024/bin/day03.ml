type instruction = 
  | Mul of int * int
  | Do
  | Dont

let read_file path =
  In_channel.with_open_text path In_channel.input_all

let get_all_matches pattern contents fn =
  let rec next_match i acc =
    match Str.search_forward pattern contents i with
    | _ -> next_match (Str.match_end ()) (fn acc)
    | exception Not_found -> acc
  in List.rev (next_match 0 [])

let extract_instructions contents =
  let mul_pat = Str.regexp {|mul(\([0-9]+\),\([0-9]+\))\|do()\|don't()|} in
  let get_group i = int_of_string (Str.matched_group i contents) in
  let make_instruction () =
    let full_match = Str.matched_string contents in
    if String.starts_with ~prefix:"mul(" full_match then
      Mul ((get_group 1), (get_group 2))
    else if String.starts_with ~prefix:"do(" full_match then
      Do
    else
      Dont
  in get_all_matches mul_pat contents (fun acc -> make_instruction () :: acc)

let part1 instructions =
  let eval acc inst =
    match inst with
    | Mul (a, b) -> acc + (a * b)
    | Do
    | Dont -> acc
  in List.fold_left eval 0 instructions

let part2 instructions =
  let rec eval instructions acc active =
    match instructions with
    | Mul (a, b) :: tl -> eval tl (if active then acc + (a * b) else acc) active
    | Do :: tl -> eval tl acc true
    | Dont :: tl -> eval tl acc false
    | [] -> acc
  in eval instructions 0 true

let _ =
  let instructions = extract_instructions (read_file "input/day03.txt") in
  Printf.printf "Part 1: %d; Part 2: %d\n" (part1 instructions) (part2 instructions)
