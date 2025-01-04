let read_entries ic =
  let rec read acc =
    Scanf.kscanf ic (fun _ _ -> List.rev acc) "%[a-z-]-%d[%[a-z]]" (fun name num chk -> read ((name, num, chk) :: acc))
  in
  read []

let read_file filename = read_entries (Scanf.Scanning.open_in filename)
let chars_of_string str = Array.to_list (Array.init (String.length str) (String.get str))

let string_of_chars lst =
  let buf = Buffer.create 16 in
  List.iter (Buffer.add_char buf) lst;
  Buffer.contents buf

(*
let make_checksum name =
  let as_chars = chars_of_string name in
  let tbl = Hashtbl.create (String.length name) in
  let add ch =
    if Hashtbl.mem tbl ch then
      Hashtbl.replace tbl ch (1 + Hashtbl.find tbl ch)
    else
      Hashtbl.add tbl ch 1
  in
  let comparator (ch1, n1) (ch2, n1) =
    match compare ch1 ch2 with
    | 0 -> compare n1 n2
    | c -> c
  in
  let _ = List.iter (fun ch -> if ch <> '-' then add ch else ()) as_chars in
  let unsorted = Hashtbl.fold (fun k v acc -> (k, v) :: acc) tbl [] in
  let sorted = List.sort comparator unsorted in
  let ordered_chars = List.map sorted fst in
  string_of_chars ordered_chars
*)

let () = ()
