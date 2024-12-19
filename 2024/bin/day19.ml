let read_input path =
  let parse lines =
    let (valid, desired) = Utils.split_sections lines in
    let patterns = List.hd valid |> String.split_on_char ',' |> List.map String.trim in
    (patterns, desired)
  in
  In_channel.with_open_text path Utils.input_lines |> parse

type trie = {
  mutable is_leaf : bool;
  mutable branches : (char * trie) list;
}

let create_trie () = { branches = []; is_leaf = false }

let add_to_trie s t =
  let l = String.length s in
  let rec insert i t =
    let is_leaf = i = l - 1 in
    match List.assoc_opt s.[i] t.branches with
    | Some t' ->
        if is_leaf then
          t'.is_leaf <- true
        else
          insert (i + 1) t'
    | None ->
        let new_entry = { is_leaf; branches = [] } in
        if not is_leaf then insert (i + 1) new_entry;
        t.branches <- (s.[i], new_entry) :: t.branches
  in
  insert 0 t

let find_prefixes s t =
  let rec loop i acc t =
    let acc' = if t.is_leaf then String.sub s 0 i :: acc else acc in
    if i = String.length s then
      acc'
    else
      List.fold_left (fun acc (ch, t) -> if ch = s.[i] then loop (i + 1) acc t else acc) acc' t.branches
  in
  let result = loop 0 [] t in
  result

let rest_of s i = String.sub s i (String.length s - i)

let count_possible pattern trie =
  let cache = Hashtbl.create 128 in
  let rec loop pattern trie =
    if pattern = "" then
      1
    else
      match Hashtbl.find_opt cache pattern with
      | Some s -> s
      | None ->
          let matches = find_prefixes pattern trie in
          let result = List.fold_left (fun acc m -> acc + loop (rest_of pattern (String.length m)) trie) 0 matches in
          Hashtbl.add cache pattern result;
          result
  in
  loop pattern trie

let _ =
  let (patterns, desired) = read_input "input/day19.txt" in
  let trie = create_trie () in
  List.iter (fun s -> add_to_trie s trie) patterns;
  let part1 = List.fold_left (fun acc pattern -> acc + if count_possible pattern trie > 0 then 1 else 0) 0 desired in
  Printf.printf "Part 1: %d\n" part1;
  let part2 = List.fold_left (fun acc pattern -> acc + count_possible pattern trie) 0 desired in
  Printf.printf "Part 2: %d\n" part2
