#use "topfind"
#require "str"

type instruction = 
  | Mul of int * int
  | Do
  | Dont

let read_all stream =
  really_input_string stream (Int64.to_int (In_channel.length stream))

let read_file path = 
  let stream = open_in path in
  let contents = read_all stream in
  close_in stream;
  contents

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
  in
  let rec get_all_matches i acc =
    match Str.search_forward mul_pat contents i with
    | _ -> get_all_matches (Str.match_end ()) (make_instruction () :: acc)
    | exception Not_found -> acc
  in
  List.rev (get_all_matches 0 [])

let _ =
  let contents = read_file "input" in
  let part1_impl acc inst =
    match inst with
    | Mul (a, b) -> acc + (a * b)
    | Do
    | Dont -> acc
  in
  let part2_impl instructions =
    let rec eval instructions acc active =
      match instructions with
      | hd :: tl -> (
          match hd with
          | Mul (a, b) -> eval tl (if active then acc + (a * b) else acc) active
          | Do -> eval tl acc true
          | Dont -> eval tl acc false
        )
      | [] -> acc
    in eval instructions 0 true
  in
  let instructions = extract_instructions contents in
  let part1 = List.fold_left part1_impl 0 instructions in
  let part2 = part2_impl instructions in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
