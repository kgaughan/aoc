#!/usr/bin/env ocaml

exception Unexpected_length

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

let count_triangles lst =
  List.length (filter_valid_triangles lst)

let rejig_triples lst =
  let rec loop acc = function
    | [] -> acc
    | (a1, a2, a3) :: (b1, b2, b3) :: (c1, c2, c3) :: rst ->
        loop ((a1, b1, c1) :: (a2, b2, c2) :: (a3, b3, c3) :: acc) rst
    | _ -> raise Unexpected_length
  in
  loop [] lst

let () =
  let src = ref "" in
  let part2 = ref false in
  let speclist = [
    ("-f", Arg.Set_string src, "File to read from");
    ("-2", Arg.Set part2, "Run part 2");
  ] in
  Arg.parse speclist print_endline "Usage:";
  if !src = ""
  then exit 1
  else let lst = read_file !src in
       let jigged = if !part2 then rejig_triples lst else lst in
       Printf.printf "Valid triangles: %d\n" (count_triangles jigged)
