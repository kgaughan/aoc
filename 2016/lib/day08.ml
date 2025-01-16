open Day08_aux

let read path = Day08_parser.main Day08_lexer.tokens (Lexing.from_channel (open_in path))

let plot ops =
  let width = 50
  and height = 6 in
  let screen = Array.make_matrix width height false
  and temp = Array.make width false in
  let rec loop = function
    | Rect (w, h) :: xs ->
        for x = 0 to w - 1 do
          for y = 0 to h - 1 do
            screen.(x).(y) <- true
          done
        done;
        loop xs
    | RotateRow (row, amount) :: xs ->
        for x = 0 to width - 1 do
          temp.(x) <- screen.(x).(row)
        done;
        for x = 0 to width - 1 do
          screen.((x + amount) mod width).(row) <- temp.(x)
        done;
        loop xs
    | RotateColumn (col, amount) :: xs ->
        for y = 0 to height - 1 do
          temp.(y) <- screen.(col).(y)
        done;
        for y = 0 to height - 1 do
          screen.(col).((y + amount) mod height) <- temp.(y)
        done;
        loop xs
    | [] -> screen
  in
  loop ops

let count_lit screen =
  let lit = ref 0 in
  for y = 0 to Array.length screen.(0) - 1 do
    for x = 0 to Array.length screen - 1 do
      if screen.(x).(y) then lit := !lit + 1
    done
  done;
  !lit

let dump screen =
  for y = 0 to Array.length screen.(0) - 1 do
    for x = 0 to Array.length screen - 1 do
      print_char (if screen.(x).(y) then '#' else '.')
    done;
    print_newline ()
  done

let part_one input = plot input |> count_lit |> Printf.printf "Part 1: %d\n%!"
let part_two input = plot input |> dump
