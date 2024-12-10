let read_input path = In_channel.with_open_text path In_channel.input_line |> Option.value ~default:""

type block =
  | Block of int
  | Empty

let parse_digit ch =
  let zero = int_of_char '0' in
  int_of_char ch - zero

let count_blocks = String.fold_left (fun acc ch -> acc + parse_digit ch) 0

let populate_blocks disc_map =
  let block_count = count_blocks disc_map in
  let blocks = Array.make block_count Empty in
  let offset = ref 0 in
  let collect i ch =
    let length = parse_digit ch in
    if i mod 2 = 0 then
      for j = !offset to !offset + length - 1 do
        blocks.(j) <- Block (i / 2)
      done;
    offset := !offset + length
  in
  String.iteri collect disc_map;
  blocks

let pack_aggressively blocks =
  let duplicate = Array.make (Array.length blocks) Empty in
  let i = ref 0
  and j = ref (Array.length blocks - 1) in
  while !i <= !j do
    if blocks.(!i) <> Empty then (
      duplicate.(!i) <- blocks.(!i);
      i := !i + 1)
    else (
      if blocks.(!j) <> Empty then (
        duplicate.(!i) <- blocks.(!j);
        i := !i + 1);
      j := !j - 1)
  done;
  duplicate

(*
 * This is awful. A better approach I only thought of after I'd done this, and
 * closer to the original way I was thinking of structuring things, would be to
 * scan for all the windows first, then scan backwards through the files
 * comparing them to each of the windows until a suitable window was found. The
 * file would be adjusted so its offset is at the start of the window, and the
 * window would be shrank suitably and its offset also adjusted.
 *
 * This would be far less fussy than the procedural method below, which is also
 * hamstrung by the fact that there's no way to break out of loops in OCaml. I
 * only implemented it the way I did because I got lazy with part 1. Even then,
 * the file could be chopped up similarly to the windows, so.. *shrug*
 *
 * Anyway, witness my shame.
 *)
let pack_carefully blocks =
  let dup = Array.copy blocks in
  let i = ref 0
  and j = ref (Array.length blocks - 1)
  and gap = ref 0
  and suitable = ref false
  and file_size = ref 0 in
  let last_id = ref Int.max_int in
  while !j > 0 do
    (* get the size of the file at the top *)
    file_size := 1;
    while !j > 0 && dup.(!j) = dup.(!j - 1) do
      file_size := !file_size + 1;
      j := !j - 1
    done;
    if
      !j > 0
      &&
      match dup.(!j) with
      | Block id -> id < !last_id
      | Empty -> false
    then (
      (* scan from the bottom for a suitable gap *)
      i := 0;
      suitable := false;
      while !i < !j && not !suitable do
        while dup.(!i) <> Empty do
          i := !i + 1
        done;
        (* find the size of the gap *)
        gap := 0;
        while !i + !gap < !j && dup.(!i + !gap) = Empty do
          gap := !gap + 1
        done;
        if !gap >= !file_size then
          suitable := true
        else
          i := !i + !gap
      done;
      if !suitable then
        for k = 0 to !file_size - 1 do
          dup.(!i + k) <- blocks.(!j + k);
          dup.(!j + k) <- Empty
        done;
      (match dup.(!j) with
      | Block id -> last_id := id
      | Empty -> last_id := Int.max_int);
      (* Skip to the next file *)
      j := !j - 1;
      while dup.(!j) = Empty do
        j := !j - 1
      done)
    else
      j := !j - 1
  done;
  dup

let checksum blocks =
  let sum_block (i, acc) block =
    match block with
    | Block id -> (i + 1, acc + (i * id))
    | Empty -> (i + 1, acc)
  in
  snd (Array.fold_left sum_block (0, 0) blocks)

let _ =
  let disc_map = read_input "input/day09.txt" in
  let blocks = populate_blocks disc_map in
  let part1 = pack_aggressively blocks |> checksum in
  let part2 = pack_carefully blocks |> checksum in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
