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

let find_cliques graph =
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

let _ =
  let graph = read_edges "input/day23-sample.txt" |> to_adjacency_list in
  let cliques = find_cliques graph in
  Printf.printf "%d\n" (List.length cliques);
  List.iter (fun set -> StringSet.to_seq set |> List.of_seq |> String.concat ", " |> print_endline) cliques;
  let part1 =
    List.filter (fun s -> StringSet.cardinal s = 3 && StringSet.exists (String.starts_with ~prefix:"t") s) cliques
  in
  Printf.printf "%d\n" (List.length part1)
