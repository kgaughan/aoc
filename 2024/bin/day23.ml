let read_edges path = Utils.read_lines path (Utils.parse_pair "%s@-%s")

let to_adjacency_list lst =
  let al = Hashtbl.create (List.length lst) in
  let add a b =
    match Hashtbl.find_opt al a with
    | Some neighbours -> Hashtbl.replace al a (b :: neighbours)
    | None -> Hashtbl.add al a [b]
  in
  List.iter
    (fun (a, b) ->
      (* the graph is undirected *)
      add a b;
      add b a)
    lst;
  al

module StringSet = Set.Make (String)

let bron_kerbosch graph =
  let intersection lst set = List.filter (fun n -> StringSet.mem n set) lst |> StringSet.of_list in
  let rec aux acc potential remaining skip =
    if StringSet.is_empty remaining && StringSet.is_empty skip then
      potential :: acc
    else
      let (acc', _, _) =
        StringSet.fold
          (fun node (acc, remaining, skip) ->
            let neighbours = Hashtbl.find graph node in
            let potential' = StringSet.add node potential
            and remaining' = intersection neighbours remaining
            and skip' = intersection neighbours skip in
            let acc' = aux acc potential' remaining' skip' in
            (acc', StringSet.remove node remaining, StringSet.add node skip))
          remaining (acc, remaining, skip)
      in
      acc'
  in
  aux [] StringSet.empty (Hashtbl.to_seq_keys graph |> StringSet.of_seq) StringSet.empty

let find_3_cliques graph =
  Hashtbl.fold
    (fun node neighbours acc ->
      Utils.pairwise_combinations neighbours
      |> List.fold_left
           (fun acc (a, b) ->
             if List.mem a (Hashtbl.find graph b) then
               List.sort compare [a; b; node] :: acc
             else
               acc)
           acc)
    graph []
  |> List.sort_uniq compare

let part1 graph =
  find_3_cliques graph |> List.filter (fun l -> List.exists (String.starts_with ~prefix:"t") l) |> List.length

let part2 graph =
  bron_kerbosch graph
  |> List.fold_left
       (fun acc clique -> if StringSet.cardinal clique > StringSet.cardinal acc then clique else acc)
       StringSet.empty
  |> StringSet.to_seq |> List.of_seq |> List.sort compare |> String.concat ","

let _ =
  let graph = read_edges "input/day23.txt" |> to_adjacency_list in
  Printf.printf "Part 1: %d\n%!" (Utils.time "part 1" (fun () -> part1 graph));
  Printf.printf "Part 2: %s\n%!" (Utils.time "part 2" (fun () -> part2 graph))
