let input_lines ic =
  let rec loop acc =
    match In_channel.input_line ic with
    | Some line -> loop (line :: acc)
    | None -> acc
  in
  loop [] |> List.rev

let read_lines path line_parser =
  let read_lines ic = input_lines ic |> List.map line_parser in
  In_channel.with_open_text path read_lines

let parse_pair fmt line = Scanf.sscanf line fmt (fun x y -> (x, y))
