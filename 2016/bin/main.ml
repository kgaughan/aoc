let time name fn =
  let start = Unix.gettimeofday () in
  let result = fn () in
  let finish = Unix.gettimeofday () in
  Printf.printf "Time for %s: %fms\n%!" name ((finish -. start) *. 1000.0);
  result

let () =
  let day = ref 0
  and src = ref "" in
  let spec = [("-day", Arg.Set_int day, "Day to run"); ("-file", Arg.Set_string src, "File to read input from")] in
  Arg.parse spec (fun _ -> ()) "AoC 2016 runner\n\nUsage:";
  let run reader part_one part_two =
    let input = reader (if !src = "" then Printf.sprintf "input/day%02d.txt" !day else !src) in
    time "part one" (fun () -> part_one input);
    time "part two" (fun () -> part_two input)
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
  | 9 -> run Day09.read Day09.part_one Day09.part_two
  | 10 -> run Day10.read Day10.part_one Day10.part_two
  | 11 -> run Day11.read Day11.part_one Day11.part_two
  | 12 -> run Day12.read Day12.part_one Day12.part_two
  | 13 -> run Day13.read Day13.part_one Day13.part_two
  | 14 -> run Day14.read Day14.part_one Day14.part_two
  | 15 -> run Day15.read Day15.part_one Day15.part_two
  | _ -> print_endline "No such solution"
