let read filename =
  let mangle line =
    String.map (fun ch -> if ch = '[' || ch = ']' then '|' else ch) line |> String.split_on_char '|' |> Array.of_list
  in
  Io.read_lines filename |> List.map mangle

let match_abba s =
  let matched = ref false in
  for i = 0 to String.length s - 4 do
    if s.[i] <> s.[i + 1] && s.[i] = s.[i + 3] && s.[i + 1] = s.[i + 2] then matched := true
  done;
  !matched

let supports_tls parts =
  let supernet_matches = ref false
  and hypernet_matches = ref false
  and checked = Array.map match_abba parts in
  for i = 0 to Array.length checked - 1 do
    if i mod 2 = 0 then
      supernet_matches := !supernet_matches || checked.(i)
    else
      hypernet_matches := !hypernet_matches || checked.(i)
  done;
  !supernet_matches && not !hypernet_matches

let part_one input = List.filter supports_tls input |> List.length |> Printf.printf "Part 1: %d\n%!"

module StringSet = Set.Make (String)

let find_abas s =
  let collected = Hashtbl.create 50 in
  for i = 0 to String.length s - 3 do
    if s.[i] <> s.[i + 1] && s.[i] = s.[i + 2] then Hashtbl.add collected (String.sub s i 3) true
  done;
  Hashtbl.to_seq_keys collected |> StringSet.of_seq

let aba_to_bab s = [s.[1]; s.[0]; s.[1]] |> List.to_seq |> String.of_seq

let partition_nets addr =
  let concat = List.fold_left StringSet.union StringSet.empty in
  let rec loop supernets hypernets is_hypernet = function
    | x :: xs ->
        if is_hypernet then
          loop supernets (x :: hypernets) (not is_hypernet) xs
        else
          loop (x :: supernets) hypernets (not is_hypernet) xs
    | [] -> (concat supernets, concat hypernets)
  in
  loop [] [] false addr

let supports_ssl parts =
  let (supernets, hypernets) = Array.map find_abas parts |> Array.to_list |> partition_nets in
  let expected_babs = StringSet.map aba_to_bab supernets in
  StringSet.inter hypernets expected_babs |> StringSet.cardinal > 0

let part_two input =
  List.map supports_ssl input
  |> List.fold_left (fun total supports -> total + if supports then 1 else 0) 0
  |> Printf.printf "Part 2: %d\n%!"
