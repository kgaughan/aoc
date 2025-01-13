let read filename = Io.read_all filename |> String.trim

let find_digit prefix n =
  let rec loop n =
    let hex_digest = Digest.to_hex (Digest.string (Printf.sprintf "%s%d" prefix n)) in
    if String.starts_with ~prefix:"00000" hex_digest then
      (hex_digest.[5], hex_digest.[6], n)
    else
      loop (n + 1)
  in
  loop n

let find_password prefix =
  let buf = Buffer.create 8 in
  let rec loop i n =
    if i = 0 then
      Buffer.contents buf
    else
      let (digit, _, n') = find_digit prefix n in
      Buffer.add_char buf digit;
      loop (i - 1) (n' + 1)
  in
  loop 8 0

let find_cinematic_password prefix =
  let password = Array.make 8 ' ' in
  let rec loop i n =
    if i = 0 then
      Array.to_seq password |> String.of_seq
    else
      let (offset_digit, digit, n') = find_digit prefix n in
      let offset = int_of_char offset_digit - 0x30 in
      if offset >= 0 && offset < 8 && password.(offset) = ' ' then (
        password.(offset) <- digit;
        loop (i - 1) (n' + 1))
      else
        loop i (n' + 1)
  in
  loop 8 0

let part_one input = find_password input |> Printf.printf "Part 1: %s\n%!"
let part_two input = find_cinematic_password input |> Printf.printf "Part 2: %s\n%!"
