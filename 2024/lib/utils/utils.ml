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

let split_line ?(delim = ',') line = String.split_on_char delim line |> List.map String.trim
let int_concat ?(delim = ",") ns = List.map string_of_int ns |> String.concat delim
let parse_pair fmt line = Scanf.sscanf line fmt (fun x y -> (x, y))
let parse_ints ?(sep = ' ') line = String.split_on_char sep line |> List.map int_of_string
let sum = List.fold_left ( + ) 0
let rest_of s i = String.sub s i (String.length s - i)
let in_bounds matrix (x, y) = x >= 0 && y >= 0 && x < Array.length matrix.(0) && y < Array.length matrix
let turns_90 (a, b) = [(a, b); (b, -a); (-b, a)]

let find_neighbours (x, y) directions check =
  List.filter_map
    (fun (dx, dy) ->
      let next = (x + dx, y + dy) in
      if check next then Some next else None)
    directions

(* Cramer's rule *)
let cramer (a1, a2) (b1, b2) (c1, c2) =
  let determiner = (a1 * b2) - (b1 * a2) in
  let x = ((c1 * b2) - (b1 * c2)) / determiner
  and y = ((a1 * c2) - (c1 * a2)) / determiner in
  if (a1 * x) + (b1 * y) = c1 && (a2 * x) + (b2 * y) = c2 then
    Some (x, y)
  else
    None

let manhattan_distance (x1, y1) (x2, y2) = abs (x1 - x2) + abs (y1 - y2)

let manhattan_circle i r =
  let rec loop r offset acc =
    if offset = 0 then
      acc
    else
      loop r (offset - 1)
        ((offset, r - offset) :: (r - offset, -offset) :: (-offset, offset - r) :: (offset - r, offset) :: acc)
  in
  let rec fill width acc =
    if width = i then
      acc
    else
      fill (width - 1) (loop width width acc)
  in
  fill r []

let render fn grid =
  let width = Array.length grid.(0)
  and height = Array.length grid in
  for y = 0 to height - 1 do
    for x = 0 to width - 1 do
      fn x y grid.(y).(x)
    done;
    print_newline ()
  done

let time name fn =
  let start = Unix.gettimeofday () in
  let result = fn () in
  let finish = Unix.gettimeofday () in
  Printf.printf "Time for %s: %fms\n" name ((finish -. start) *. 1000.0);
  flush stdout;
  result

module IntPair = struct
  type t = int * int

  let compare (x0, y0) (x1, y1) =
    match Int.compare x0 x1 with
    | 0 -> Int.compare y0 y1
    | c -> c
end

module IntPairSet = Set.Make (IntPair)
module IntPairMap = Map.Make (IntPair)

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
