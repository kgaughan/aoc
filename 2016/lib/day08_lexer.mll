{
open Day08_parser

exception Error of string
}

rule tokens = parse
| [' ' '\n'] { tokens lexbuf }
| "rotate" ' '+ "row" ' '+ "y=" { ROTATE_ROW }
| "rotate" ' '+ "column" ' '+ "x=" { ROTATE_COLUMN }
| "rect" { RECT }
| "by" { BY }
| (['0'-'9']+ as x) 'x' (['0'-'9']+ as y) { DIM (int_of_string x, int_of_string y) }
| ['0'-'9']+ as i { NUMBER (int_of_string i) }
| eof { END }
| _ { raise (Error (Printf.sprintf "At offset %d: unexpected character.\n" (Lexing.lexeme_start lexbuf))) }
