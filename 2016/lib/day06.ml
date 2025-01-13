let read = Io.read_lines

module Histogram = Map.Make (Char)

let build_repetition_code_histograms lines =
  let cols = Array.make (String.length (List.hd lines)) Histogram.empty in
  List.iter
    (fun line ->
      String.iteri
        (fun i ch ->
          match Histogram.find_opt ch cols.(i) with
          | Some v -> cols.(i) <- Histogram.add ch (v + 1) cols.(i)
          | None -> cols.(i) <- Histogram.add ch 1 cols.(i))
        line)
    lines;
  cols

let get_codes cmp hs =
  Array.map (fun h -> Histogram.to_seq h |> List.of_seq |> List.sort cmp |> List.hd |> fst) hs
  |> Array.to_seq |> String.of_seq

let part_one lines =
  let cmp (_, n1) (_, n2) = Int.compare n2 n1 in
  build_repetition_code_histograms lines |> get_codes cmp |> Printf.printf "Part 1: %s\n%!"

let part_two lines =
  let cmp (_, n1) (_, n2) = Int.compare n1 n2 in
  build_repetition_code_histograms lines |> get_codes cmp |> Printf.printf "Part 2: %s\n%!"
