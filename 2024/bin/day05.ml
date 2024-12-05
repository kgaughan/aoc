let read_input path =
  let parse_order_line line = Scanf.sscanf line "%d|%d" (fun x y -> (x, y))
  and parse_page_list line = String.split_on_char ',' line |> List.map int_of_string
  and split_sections lines =
    let rec loop acc = function
      | "" :: tl -> (acc, tl)
      | hd :: tl -> loop (hd :: acc) tl
      | [] -> (acc, [])
    in loop [] lines
  in
  let parse lines =
    let orderings, updates = split_sections lines in
    (List.map parse_order_line orderings, List.map parse_page_list updates)
  in
  In_channel.with_open_text path In_channel.input_lines |> parse

let intersecting l1 l2 =
  List.exists (fun x -> List.exists (fun y -> y = x) l2) l1

let is_valid_update orderings =
  let precedes_in_order page following =
    let expected_predecessors = List.filter (fun (_, y) -> y = page) orderings |> List.map fst in
    not (intersecting following expected_predecessors)
  in
  let rec in_order = function
    | page :: tl -> 
      if precedes_in_order page tl then
        in_order tl
      else
        false
    | [] -> true
  in in_order

let get_middle_entry lst =
  List.nth lst ((List.length lst) / 2)

let _ =
  let orderings, updates = read_input "input/day05.txt" in
  let valid_updates = List.filter (is_valid_update orderings) updates in
  let part1 = List.map get_middle_entry valid_updates |> List.fold_left (+) 0 in
  Printf.printf "Part 1: %d\n" part1
