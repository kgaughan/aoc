{
open Asmb_parser

exception Error of string

let to_reg = function
  | 'a' -> Asmb.A
  | 'b' -> Asmb.B
  | 'c' -> Asmb.C
  | 'd' -> Asmb.D
  | x -> raise (Error (Printf.sprintf "Bad register: %c\n" x))
}

rule tokens = parse
| [' ' '\n'] { tokens lexbuf }
| "cpy" { CPY }
| "inc" { INC }
| "dec" { DEC }
| "jnz" { JNZ }
| ['a' 'b' 'c' 'd'] as reg { REG (to_reg reg) }
| ('-'? ['0'-'9']+) as number { LITERAL (int_of_string number) }
| eof { END }
| _ { raise (Error (Printf.sprintf "At offset %d: unexpected character.\n" (Lexing.lexeme_start lexbuf))) }
