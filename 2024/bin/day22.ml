let next_secret s =
  let mask = (1 lsl 24) - 1 in
  let s1 = (s lsl 6) lxor s land mask in
  let s2 = (s1 lsr 5) lxor s1 land mask in
  (s2 lsl 11) lxor s2 land mask

let repeat fn n v =
  let rec loop n v = if n = 0 then v else loop (n - 1) (fn v) in
  loop n v

let generate_prices_and_diffs n s =
  let rec loop n prev acc =
    if n = 0 then
      List.rev acc
    else
      let current = next_secret prev in
      let prev_price = prev mod 10
      and price = current mod 10 in
      let diff = price - prev_price in
      loop (n - 1) current ((price, diff) :: acc)
  in
  loop n s []

let recognise_sequences lst =
  let rec loop lst acc =
    match lst with
    | e1 :: e2 :: e3 :: e4 :: tl ->
        if fst e4 > 0 then
          let key = List.fold_left (fun acc (_, diff) -> diff + 10 + (acc * 20)) 0 [e1; e2; e3; e4] in
          loop (e2 :: e3 :: e4 :: tl) ((key, fst e4) :: acc)
        else
          loop (e2 :: e3 :: e4 :: tl) acc
    | _ -> acc
  in
  loop lst []

module IntMap = Map.Make (Int)

let find_sequence_prices secret =
  generate_prices_and_diffs 2000 secret |> recognise_sequences |> List.to_seq |> IntMap.of_seq

let seq_sums secret_numbers =
  List.fold_left
    (fun map secret ->
      let seqs = find_sequence_prices secret in
      IntMap.fold
        (fun key price map ->
          match IntMap.find_opt key map with
          | Some total -> IntMap.add key (total + price) map
          | None -> IntMap.add key price map)
        seqs map)
    IntMap.empty secret_numbers

let max_value map = IntMap.fold (fun _ v max -> Int.max v max) map 0

let _ =
  let secret_numbers = Utils.read_lines "input/day22.txt" int_of_string in
  let part1 = List.map (fun secret -> repeat next_secret 2000 secret) secret_numbers |> Utils.sum in
  Printf.printf "Part 1: %d\n" part1;
  let part2 = seq_sums secret_numbers |> max_value in
  Printf.printf "Part 2: %d\n" part2
