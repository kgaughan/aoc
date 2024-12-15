let read_input path =
  let split_sections lines =
    let rec loop acc = function
      | "" :: tl -> (acc, tl)
      | hd :: tl -> loop (hd :: acc) tl
      | [] -> (acc, [])
    in
    loop [] lines
  in
  let parse lines =
    let (orderings, updates) = split_sections lines in
    (List.map (Utils.parse_pair "%d|%d") orderings, List.map (Utils.parse_ints ~sep:',') updates)
  in
  In_channel.with_open_text path Utils.input_lines |> parse

let must_precede first second = List.exists (fun entry -> entry = (second, first))

let rec is_in_order orderings = function
  | page :: tl -> List.for_all (fun following -> must_precede following page orderings) tl && is_in_order orderings tl
  | [] -> true

let fix_order orderings = List.sort (fun a b -> if must_precede a b orderings then -1 else 1)
let get_middle_entry lst = List.nth lst (List.length lst / 2)

let _ =
  let (orderings, updates) = read_input "input/day05.txt" in
  let (valid_updates, invalid_updates) = List.partition (is_in_order orderings) updates in
  let sum_middle_entries updates = List.map get_middle_entry updates |> Utils.sum in
  let part1 = sum_middle_entries valid_updates
  and part2 = List.map (fix_order orderings) invalid_updates |> sum_middle_entries in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
