let read_input path =
  let parse lines =
    let (valid, desired) = Utils.split_sections lines in
    let patterns = List.hd valid |> String.split_on_char ',' |> List.map String.trim in
    (patterns, desired)
  in
  In_channel.with_open_text path Utils.input_lines |> parse

let count_possible pattern trie =
  let cache = Hashtbl.create (String.length pattern) in
  let rec loop pattern trie =
    let count acc m = acc + loop (Utils.rest_of pattern (String.length m)) trie in
    if pattern = "" then
      1
    else
      match Hashtbl.find_opt cache pattern with
      | Some s -> s
      | None ->
          let result = Trie.find_prefixes pattern trie |> List.fold_left count 0 in
          Hashtbl.add cache pattern result;
          result
  in
  loop pattern trie

let _ =
  let (patterns, desired) = read_input "input/day19.txt" in
  let trie = Trie.create () in
  List.iter (fun s -> Trie.add s trie) patterns;
  let part1 = List.fold_left (fun acc pattern -> acc + if count_possible pattern trie > 0 then 1 else 0) 0 desired in
  Printf.printf "Part 1: %d\n" part1;
  let part2 = List.fold_left (fun acc pattern -> acc + count_possible pattern trie) 0 desired in
  Printf.printf "Part 2: %d\n" part2
