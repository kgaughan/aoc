#!/usr/bin/env ocaml

let is_valid_triangle (a, b, c) =
  (* Partly sort the numbers so the biggest comes last. *)
  let a', x = (min a b), (max a b) in
  let b', c' = (min x c), (max x c) in
  a' + b' > c'

let read_file filename =
  let ic = Scanf.Scanning.open_in filename in
  let rec read acc =
    Scanf.kscanf ic (fun _ _ -> List.rev acc)
                 " %d %d %d" (fun a b c -> read ((a, b, c) :: acc))
  in
  read []

let filter_valid_triangles =
  List.filter is_valid_triangle

let count_valid_triangles lst =
  List.length (filter_valid_triangles lst)

let count_triangles_in filename =
  count_valid_triangles (read_file filename)

let () =
  let src = ref "" in
  let speclist = [
    ("-f", Arg.Set_string src, "File to read from");
  ] in
  Arg.parse speclist print_endline "Usage:";
  if !src = ""
  then exit 1
  else Printf.printf "Valid triangles: %d\n" (count_triangles_in !src)
