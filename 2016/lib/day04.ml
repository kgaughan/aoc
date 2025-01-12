let read_entries ic =
  let rec read acc =
    if Scanf.Scanning.end_of_input ic then
      acc
    else
      Scanf.bscanf ic "%[a-z-]%d[%[a-z]]\n" (fun name num chk -> read ((name, num, chk) :: acc))
  in
  read []

let read filename = Scanf.Scanning.open_in filename |> read_entries

let make_histogram s =
  let add tbl ch =
    (match Hashtbl.find_opt tbl ch with
    | Some n -> Hashtbl.replace tbl ch (1 + n)
    | None -> if ch <> '-' then Hashtbl.add tbl ch 1);
    tbl
  in
  String.fold_left add (Hashtbl.create (String.length s)) s

let make_checksum name =
  let tbl = make_histogram name in
  let comparator ch1 ch2 =
    match compare (Hashtbl.find tbl ch1) (Hashtbl.find tbl ch2) with
    | 0 -> compare ch1 ch2
    | x -> -x
  in
  Hashtbl.to_seq_keys tbl |> List.of_seq |> List.sort comparator
  |> List.filteri (fun i _ -> i < 5)
  |> List.to_seq |> String.of_seq

let caesar n input =
  let rotate ch = ((int_of_char ch - 0x61 + n) mod 26) + 0x61 |> char_of_int in
  String.map rotate input

let contains s1 s2 =
  try
    ignore (Str.search_forward (Str.regexp_string s2) s1 0);
    true
  with Not_found -> false

let valid_sectors input = List.filter (fun (n, _, c) -> c = make_checksum n) input

let part_one input =
  valid_sectors input |> List.fold_left (fun acc (_, id, _) -> acc + id) 0 |> Printf.printf "Part 1: %d\n%!"

let part_two input =
  let (_, id, _) = valid_sectors input |> List.find (fun (n, id, _) -> contains (caesar id n) "north") in
  Printf.printf "Part 2: %d\n%!" id
