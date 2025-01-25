{
open Day11_parser

exception Error of string
}

rule tokens = parse
| [' ' '\n' '-' ',' '.'] { tokens lexbuf }
| "and" { tokens lexbuf }
| "The" { THE }
| ['a'-'z']+ as id { ID id }
| "floor" ' '+ "contains" { FLOOR_CONTAINS }
| "nothing" ' '+ "relevant" { NOTHING_RELEVANT }
| 'a' ' '+ (['a'-'z']+ as element) ' '+ "generator" { GENERATOR element }
| 'a' ' '+ (['a'-'z']+ as element) "-compatible" ' '+ "microchip" { GENERATOR element }
| eof { END }
| _ { raise (Error (Printf.sprintf "At offset %d: unexpected character.\n" (Lexing.lexeme_start lexbuf))) }
