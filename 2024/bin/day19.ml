let read_input path =
  let parse lines =
    let (valid, desired) = Utils.split_sections lines in
    let patterns = List.hd valid |> Utils.split_line in
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
  let trie = Trie.of_list patterns in
  let possible_counts = List.map (fun pattern -> count_possible pattern trie) desired in
  let part1 = List.fold_left (fun acc n -> acc + if n > 0 then 1 else 0) 0 possible_counts
  and part2 = Utils.sum possible_counts in
  Printf.printf "Part 1: %d; part 2: %d\n" part1 part2
