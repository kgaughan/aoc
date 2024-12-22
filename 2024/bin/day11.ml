let has_even_number_of_digits n = int_of_float (log10 (float_of_int n)) mod 2 = 1

let split_number n =
  let as_string = string_of_int n in
  let half = String.length as_string / 2 in
  (int_of_string (String.sub as_string 0 half), int_of_string (String.sub as_string half half))

let process n tiles =
  let cache = Hashtbl.create 100000 in
  let rec blink tile n =
    if n = 0 then
      1
    else
      match Hashtbl.find_opt cache (tile, n) with
      | Some result -> result
      | None ->
          let result =
            match tile with
            | 0 -> blink 1 (n - 1)
            | _ when has_even_number_of_digits tile ->
                let (a, b) = split_number tile in
                blink a (n - 1) + blink b (n - 1)
            | _ -> blink (tile * 2024) (n - 1)
          in
          Hashtbl.add cache (tile, n) result;
          result
  in
  List.map (fun tile -> blink tile n) tiles |> Utils.sum

module IntMap = Map.Make (Int)

(* Something I've seen other people do other than memoisation is to keep a
   map of the frequencies of each number. I would expect the frequency method
   is faster for smaller numbers of iterations, but that the memoisation method
   comes to dominate later as it can short-circuit paths. Now, part of this
   slowness could be down to how Map (as opposed to Hashtbl) is implemented in
   OCaml, but I haven't benchmarked that. May use of the IntMap.update function
   might also impact its speed.

   If nothing else, this is a demonstration of a significant chunk of the Map
   module's interface. *)
let process_freq n tiles =
  let incr k v map =
    match IntMap.find_opt k map with
    | Some _ -> IntMap.update k (Option.map (( + ) v)) map
    | None -> IntMap.add k v map
  in
  let initial = List.fold_left (fun map tile -> incr tile 1 map) IntMap.empty tiles in
  let rec blink n map =
    if n = 0 then
      map
    else
      blink (n - 1)
        (IntMap.fold
           (fun tile v map ->
             (match tile with
             | 0 -> map |> incr 1 v
             | _ when has_even_number_of_digits tile ->
                 let (a, b) = split_number tile in
                 map |> incr a v |> incr b v
             | _ -> map |> incr (tile * 2024) v)
             |> incr tile ~-v)
           map map)
  in
  IntMap.fold (fun _ v total -> v + total) (blink n initial) 0

let _ =
  let tiles = Utils.read_line "input/day11.txt" |> Utils.parse_ints in
  let part1 = Utils.time "part 1" (fun () -> process 25 tiles) in
  Printf.printf "Part 1: %d\n" part1;
  let part2 = Utils.time "part 2 " (fun () -> process 75 tiles) in
  Printf.printf "Part 2: %d\n" part2;
  let part2_freq = Utils.time "part 2 (bonus)" (fun () -> process_freq 75 tiles) in
  Printf.printf "Bonus part 2 (frequency): %d\n" part2_freq
