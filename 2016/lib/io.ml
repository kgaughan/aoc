let read_all path = In_channel.with_open_text path In_channel.input_all

let input_lines ic =
  let rec loop acc =
    match In_channel.input_line ic with
    | Some line -> loop (line :: acc)
    | None -> acc
  in
  loop []

let read_lines path = In_channel.with_open_text path input_lines |> List.rev
