%token <int> NUMBER
%token <int * int> DIM
%token ROTATE_ROW ROTATE_COLUMN RECT BY END

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
| ROTATE_ROW NUMBER BY NUMBER { Day08_aux.RotateRow ($2, $4) }
| ROTATE_COLUMN NUMBER BY NUMBER { Day08_aux.RotateColumn ($2, $4) }
;
