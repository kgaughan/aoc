let () =
  let day = ref 0
  and src = ref "" in
  let spec = [("-day", Arg.Set_int day, "Day to run"); ("-file", Arg.Set_string src, "File to read input from")] in
  Arg.parse spec (fun _ -> ()) "AoC 2016 runner\n\nUsage:";
  if !src = "" then src := Printf.sprintf "input/day%02d.txt" !day;
  match !day with
  | 1 ->
      let input = Day01.read !src in
      Day01.part_one input;
      Day01.part_two input
  | 2 ->
      let input = Day02.read !src in
      Day02.part_one input;
      Day02.part_two input
  | 3 ->
      let input = Day03.read !src in
      Day03.part_one input;
      Day03.part_two input
  | 4 ->
      let input = Day04.read !src in
      Day04.part_one input;
      Day04.part_two input
  | 5 ->
      let input = Day05.read !src in
      Day05.part_one input;
      Day05.part_two input
  | _ -> print_endline "No such solution"
