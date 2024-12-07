(* I'm aware there's a paper describing how this can be solved
   (https://arxiv.org/abs/1502.05450v1), but I'm not reading it and I'm just
   bruteforcing it. *)

let read_input path =
  let parse line =
    let (sum, rest) = Scanf.sscanf line "%d: %s@\n" (fun sum rest -> (sum, rest)) in
    let numbers = String.split_on_char ' ' rest |> List.map int_of_string in
    (sum, numbers)
  in
  In_channel.with_open_text path Utils.input_lines |> List.map parse

let solve1 (sum, numbers) =
  let rec check numbers acc =
    if acc > sum then
      false
    else
      match numbers with
      | hd :: tl -> check tl (hd + acc) || (hd > 0 && check tl (hd * acc))
      | [] -> acc = sum
  in
  check numbers 0

let solve2 (sum, numbers) =
  let rec check numbers acc =
    if acc > sum then
      false
    else
      match numbers with
      | hd :: tl ->
          if check tl (hd + acc) || (hd > 0 && check tl (hd * acc)) then
            true
          else
            check tl (int_of_string (string_of_int acc ^ string_of_int hd))
      | [] -> acc = sum
  in
  check numbers 0

let sum_solvable solve equations = List.filter solve equations |> List.map fst |> Utils.sum

let _ =
  let equations = read_input "input/day07-sample.txt" in
  let part1 = sum_solvable solve1 equations in
  let part2 = sum_solvable solve2 equations in
  Printf.printf "Part 1: %d; Part 2: %d\n" part1 part2
