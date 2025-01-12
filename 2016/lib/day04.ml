let read_entries ic =
  let rec read acc =
    if Scanf.Scanning.end_of_input ic then
      acc
    else
      Scanf.bscanf ic "%[a-z-]%d[%[a-z]]\n" (fun name num chk ->
          read ((String.sub name 0 (String.length name - 1), num, chk) :: acc))
  in
  read []

let read filename = Scanf.Scanning.open_in filename |> read_entries
let chars_of_string str = String.to_seq str |> List.of_seq

let make_checksum name =
  let tbl = Hashtbl.create (String.length name) in
  let add ch =
    match Hashtbl.find_opt tbl ch with
    | Some n -> Hashtbl.replace tbl ch (1 + n)
    | None -> Hashtbl.add tbl ch 1
  in
  let comparator (ch1, n1) (ch2, n2) =
    match compare n1 n2 with
    | 0 -> compare ch1 ch2
    | x -> -x
  in
  List.iter (fun ch -> if ch <> '-' then add ch else ()) (chars_of_string name);
  let histogram = Hashtbl.fold (fun k v acc -> (k, v) :: acc) tbl [] |> List.sort comparator in
  let untruncated = histogram |> List.map fst |> List.to_seq |> String.of_seq in
  String.sub untruncated 0 5

let caesar n input =
  String.map
    (fun ch ->
      if ch = '-' then
        ' '
      else
        char_of_int (((int_of_char ch - 0x61 + n) mod 26) + 0x61))
    input

let contains s1 s2 =
  let re = Str.regexp_string s2 in
  try
    ignore (Str.search_forward re s1 0);
    true
  with Not_found -> false

let valid_sectors input =
  List.fold_left (fun acc (n, id, c) -> if c = make_checksum n then (n, id) :: acc else acc) [] input

let part_one input =
  valid_sectors input |> List.fold_left (fun acc (_, id) -> acc + id) 0 |> Printf.printf "Part 1: %d\n%!"

let part_two input =
  valid_sectors input
  |> List.fold_left (fun result (n, id) -> if contains (caesar id n) "north" then id else result) 0
  |> Printf.printf "Part 2: %d\n%!"
