%{
package day6

import (
	"log"
)

// turn (on|off) x1,y1 to x2,y2
// toggle x1,y1 to x2,y2
%}

%union {
	n int
	coords Coord[2]
}

%token NUM
%token TOGGLE
%token TURN
%token OFF
%token OFF
%token THROUGH

%%

list	: /* empty */
		| list stmt
		; 

stmt	: TURN ON coord THROUGH coord { }
		| TURN OFF coord THROUGH coord { }
		| TOGGLE coord THROUGH coord { }
		;

coord	: NUM ',' NUM { $$.coord[$$.n].X =  }
		;

%%
