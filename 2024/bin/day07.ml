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

let find_solvable equations =
  let solve (sum, numbers) =
    let rec check numbers acc =
      match numbers with
      | hd :: tl -> check tl (hd + acc) || check tl (hd * acc)
      | [] -> acc = sum
    in
    check numbers 0
  in
  List.filter solve equations

let _ =
  let equations = read_input "input/day07.txt" in
  let part1 = find_solvable equations |> List.map fst |> Utils.sum in
  Printf.printf "Part 1: %d\n" part1
