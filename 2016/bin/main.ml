let () =
  let day = ref 0
  and src = ref "" in
  let spec = [("-day", Arg.Set_int day, "Day to run"); ("-file", Arg.Set_string src, "File to read input from")] in
  Arg.parse spec (fun _ -> ()) "AoC 2016 runner\n\nUsage:";
  let run reader part_one part_two =
    let input = reader (if !src = "" then Printf.sprintf "input/day%02d.txt" !day else !src) in
    part_one input;
    part_two input
  in
  match !day with
  | 1 -> run Day01.read Day01.part_one Day01.part_two
  | 2 -> run Day02.read Day02.part_one Day02.part_two
  | 3 -> run Day03.read Day03.part_one Day03.part_two
  | 4 -> run Day04.read Day04.part_one Day04.part_two
  | 5 -> run Day05.read Day05.part_one Day05.part_two
  | 6 -> run Day06.read Day06.part_one Day06.part_two
  | 7 -> run Day07.read Day07.part_one Day07.part_two
  | 8 -> run Day08.read Day08.part_one Day08.part_two
  | _ -> print_endline "No such solution"
