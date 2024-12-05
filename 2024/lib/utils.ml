let read_lines path line_parser =
  let read_lines ic = In_channel.input_lines ic |> List.map line_parser in
  In_channel.with_open_text path read_lines

let read_all path = 
  In_channel.with_open_text path In_channel.input_all

let make_counter_table size xs =
  let tbl = Hashtbl.create size in
  let add k =
    match Hashtbl.find_opt tbl k with
    | Some v -> Hashtbl.replace tbl k (v + 1)
    | None -> Hashtbl.add tbl k 1
  in
  List.iter add xs;
  tbl

let get_all_matches pattern contents fn =
  let rec next_match i acc =
    match Str.search_forward pattern contents i with
    | _ -> next_match (Str.match_end ()) (fn acc)
    | exception Not_found -> acc
  in List.rev (next_match 0 [])

let sum =
  List.fold_left (+) 0
