%token <int> LITERAL
%token <Asmb.reg> REG
%token CPY INC DEC JNZ END

%start main
%type <Asmb.bytecode array> main

%%

main:
| instructions END { $1 |> List.rev |> Array.of_list }
;

instructions:
| instruction { [$1] }
| instructions instruction { $2 :: $1 }
;

instruction:
| CPY LITERAL REG { Asmb.CpyI ($2, $3) }
| CPY REG REG { Asmb.CpyR ($2, $3) }
| DEC REG { Asmb.Dec $2 }
| INC REG { Asmb.Inc $2 }
| JNZ REG LITERAL { Asmb.Jnz ($2, $3) }
;
