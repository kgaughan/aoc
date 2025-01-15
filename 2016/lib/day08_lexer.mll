{
open Day08_parser

exception Error of string
}

rule tokens = parse
| [' ' '\n'] { tokens lexbuf }
| "rotate" { ROTATE }
| "row" { ROW }
| "column" { COLUMN }
| "rect" { RECT }
| "by" { BY }
| (['0'-'9']+ as x) 'x' (['0'-'9']+ as y) { DIM (int_of_string x, int_of_string y) }
| '=' { EQUALS }
| ['0'-'9']+ as i { NUMBER (int_of_string i) }
| ['a'-'z'] as v { VAR v }
| eof { END }
| _ { raise (Error (Printf.sprintf "At offset %d: unexpected character.\n" (Lexing.lexeme_start lexbuf))) }
