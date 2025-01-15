%token <int> NUMBER
%token <char> VAR
%token <int * int> DIM
%token ROTATE ROW COLUMN RECT BY EQUALS END

%start main
%type <Day08_aux.operation list> main

%%

main:
| exprs END { $1 |> List.rev }
;

exprs:
| expr { [$1] }
| exprs expr { $2 :: $1 }
;

expr:
| RECT DIM { Day08_aux.Rect ($2) }
| ROTATE ROW VAR EQUALS NUMBER BY NUMBER { Day08_aux.RotateRow ($5, $7) }
| ROTATE COLUMN VAR EQUALS NUMBER BY NUMBER { Day08_aux.RotateColumn ($5, $7) }
;
