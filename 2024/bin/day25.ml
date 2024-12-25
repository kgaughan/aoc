type component =
  | Lock of int array
  | Key of int array

let to_component lines =
  let pins = Array.make (String.length (List.hd lines)) 0 in
  let rec loop = function
    | hd :: tl ->
        String.iteri (fun i ch -> if ch = '#' then pins.(i) <- pins.(i) + 1) hd;
        loop tl
    | [] -> ()
  in
  loop lines;
  if List.hd lines = "#####" then Lock pins else Key pins

let read_input path = In_channel.with_open_text path Utils.input_lines |> Utils.split_blocks to_component

let count_matching components =
  let (locks, keys) =
    List.partition_map
      (fun c ->
        match c with
        | Lock a -> Left a
        | Key a -> Right a)
      components
  in
  Utils.combinations (fun acc a b -> acc + if Array.for_all2 (fun lh kh -> lh + kh <= 7) a b then 1 else 0) 0 locks keys

let _ =
  let components = read_input "input/day25.txt" in
  let matching = count_matching components in
  Printf.printf "Part 1: %d\n" matching
