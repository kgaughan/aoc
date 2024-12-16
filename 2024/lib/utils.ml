let input_lines ic =
  let rec loop acc =
    match In_channel.input_line ic with
    | Some line -> loop (line :: acc)
    | None -> acc
  in
  loop [] |> List.rev

let read_all path = In_channel.with_open_text path In_channel.input_all

let read_lines path line_parser =
  let read_lines ic = input_lines ic |> List.map line_parser in
  In_channel.with_open_text path read_lines

let read_line path = In_channel.with_open_text path In_channel.input_line |> Option.value ~default:""

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
  in
  List.rev (next_match 0 [])

let parse_digit ch =
  let zero = int_of_char '0' in
  int_of_char ch - zero

let array_of_string fn empty line =
  let result = Array.make (String.length line) empty in
  String.iteri (fun i ch -> result.(i) <- fn ch) line;
  result

let parse_pair fmt line = Scanf.sscanf line fmt (fun x y -> (x, y))
let parse_ints ?(sep = ' ') line = String.split_on_char sep line |> List.map int_of_string
let sum = List.fold_left ( + ) 0

module IntPair = struct
  type t = int * int

  let compare (x0, y0) (x1, y1) =
    match Int.compare x0 x1 with
    | 0 -> Int.compare y0 y1
    | c -> c
end

module IntPairSet = Set.Make (IntPair)

let split_sections lines =
  let rec loop acc = function
    | "" :: tl -> (acc |> List.rev, tl)
    | hd :: tl -> loop (hd :: acc) tl
    | [] -> (acc, [])
  in
  loop [] lines

let fold_matrix fn init matrix =
  Array.fold_left
    (fun (y, acc) row -> (y + 1, Array.fold_left (fun (x, acc) cell -> (x + 1, fn x y acc cell)) (0, acc) row |> snd))
    (0, init) matrix
  |> snd

let find_cell fn matrix =
  let height = Array.length matrix
  and width = Array.length matrix.(0) in
  let rec loop_rows y =
    let rec loop_cols x =
      if x = width then
        None
      else if fn matrix.(y).(x) then
        Some (x, y)
      else
        loop_cols (x + 1)
    in
    if y = height then
      None
    else
      match loop_cols 0 with
      | Some pos -> Some pos
      | None -> loop_rows (y + 1)
  in
  loop_rows 0
