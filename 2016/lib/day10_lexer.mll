{
open Day10_parser

exception Error of string
}

rule tokens = parse
| [' ' '\n'] { tokens lexbuf }
| "value" { VALUE }
| "goes" ' '+ "to" { GOES_TO }
| "bot" { BOT }
| "output" { OUTPUT }
| ['0'-'9']+ as i { NUM (int_of_string i) }
| "gives" ' '+ "low" ' '+ "to" { GIVES_LOW }
| "and" ' '+ "high" ' '+ "to" { GIVES_HIGH }
| eof { END }
| _ { raise (Error (Printf.sprintf "At offset %d: unexpected character.\n" (Lexing.lexeme_start lexbuf))) }
