%token <string> GENERATOR
%token <string> MICROCHIP
%token <string> ID
%token THE FLOOR_CONTAINS NOTHING_RELEVANT AND END

%start main
%type <Day11_aux.device list list> main

%%

main:
| floors END { $1 |> List.rev }
;

floors:
| floor { [$1] }
| floors floor { $2 :: $1 }
;

floor:
| THE ID FLOOR_CONTAINS devices { $4 }
;

devices:
| NOTHING_RELEVANT { [] }
| device { [$1] }
| devices device { $2 :: $1 }
;

device:
| GENERATOR { Day11_aux.Generator $1 }
| MICROCHIP { Day11_aux.Microchip $1 }
;
