%token <int> NUM
%token VALUE GOES_TO BOT OUTPUT GIVES_LOW GIVES_HIGH END

%start main
%type <Day10_aux.action list> main

%%

main:
| exprs END { $1 |> List.rev }
;

exprs:
| expr { [$1] }
| exprs expr { $2 :: $1 }
;

expr:
| VALUE NUM GOES_TO BOT NUM { Day10_aux.Receive ($2, $5) }
| BOT NUM GIVES_LOW recipient GIVES_HIGH recipient { Day10_aux.Give ($2, $4, $6) }
;

recipient:
| BOT NUM { Day10_aux.Bot $2 }
| OUTPUT NUM { Day10_aux.Output $2 }
;
