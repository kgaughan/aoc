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
      let this = next_secret prev in
      let prev_price = prev mod 10
      and this_price = this mod 10 in
      let diff = this_price - prev_price in
      loop (n - 1) this ((this_price, diff) :: acc)
  in
  loop n s []

let _ =
  let secret_numbers = Utils.read_lines "input/day22.txt" int_of_string in
  let part1 = List.map (fun secret -> repeat next_secret 2000 secret) secret_numbers |> Utils.sum in
  Printf.printf "Part 1: %d\n" part1
