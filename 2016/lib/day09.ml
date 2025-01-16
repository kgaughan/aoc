let read path = Io.read_all path |> String.trim

let rec decompress s version =
  let rec consume size offset =
    if offset = String.length s then
      size
    else
      match s.[offset] with
      | '(' ->
          let x_offset = String.index_from s offset 'x'
          and close_offset = String.index_from s offset ')' in
          let to_consume = int_of_string (String.sub s (offset + 1) (x_offset - offset - 1))
          and times = int_of_string (String.sub s (x_offset + 1) (close_offset - x_offset - 1)) in
          let chunk_size =
            if version = 1 then
              to_consume
            else
              decompress (String.sub s (close_offset + 1) to_consume) version
          in
          consume (size + (times * chunk_size)) (close_offset + to_consume + 1)
      | _ -> consume (size + 1) (offset + 1)
  in
  consume 0 0

let part_one input = decompress input 1 |> Printf.printf "Part 1: %d\n%!"
let part_two input = decompress input 2 |> Printf.printf "Part 2: %d\n%!"
