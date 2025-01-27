let read path = Assembunny.Asmb_parser.main Assembunny.Asmb_lexer.tokens (Lexing.from_channel (open_in path))

let part_one input =
  let state = Assembunny.Asmb.execute input in
  Printf.printf "Part 1: %d\n%!" state.a

let part_two input = ignore input
