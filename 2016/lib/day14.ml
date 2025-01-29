let read filename = Io.read_all filename |> String.trim

let find_triplet s =
  let finish = String.length s - 3 in
  let rec loop i =
    if s.[i] = s.[i + 1] && s.[i] = s.[i + 2] then
      Some s.[i]
    else if i = finish then
      None
    else
      loop (i + 1)
  in
  loop 0

let contains_quintuple s ch =
  let rec loop i =
    if s.[i] = ch && s.[i + 1] = ch && s.[i + 2] = ch && s.[i + 3] = ch && s.[i + 4] = ch then
      true
    else if i = String.length s - 5 then
      false
    else
      loop (i + 1)
  in
  loop 0

let generate_otp salt stretches =
  let rec stretch n s = if n = 0 then s else Digest.string s |> Digest.to_hex |> stretch (n - 1) in
  let cache = Hashtbl.create 50000 in
  let fetch_digest n =
    match Hashtbl.find_opt cache n with
    | Some h -> h
    | None ->
        let h = Printf.sprintf "%s%d" salt n |> stretch (stretches + 1) in
        Hashtbl.add cache n h;
        h
  in
  let rec check_following ch i n =
    if n = 0 then
      None
    else
      let digest = fetch_digest i in
      if contains_quintuple digest ch then
        Some digest
      else
        check_following ch (i + 1) (n - 1)
  in
  let rec find_key last_key to_match i =
    if to_match = 0 then
      last_key
    else
      let digest = fetch_digest i in
      match find_triplet digest with
      | Some ch -> (
          match check_following ch (i + 1) 1000 with
          | Some _ -> find_key i (to_match - 1) (i + 1)
          | None -> find_key last_key to_match (i + 1))
      | None -> find_key last_key to_match (i + 1)
  in
  find_key 0 64 0

let part_one input = generate_otp input 0 |> Printf.printf "Part 1: %d\n%!"
let part_two input = generate_otp input 2016 |> Printf.printf "Part 2: %d\n%!"
