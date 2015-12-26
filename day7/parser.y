%{
package day7

import (
	"log"
)
%}

%union {
	node Node
}

%token <node> ASSIGN
%token <node> NOT AND OR
%token <node> LSHIFT RSHIFT
%token <node> IDENTIFIER LITERAL

%type <node> expr val

%%

top:	stmts {
			log.Println("top")
		}

stmts:	/* empty */ {
			// Initialise statement list?
			log.Println("empty statment")
		}
|		stmts stmt {
			// Append statement to statement list.
			log.Println("appending statement")
		}
;

stmt:	expr ASSIGN IDENTIFIER {
			log.Println("Assign")
		}
;

expr:	val {
			$$ = $1
		}
|		val AND val {
			$$ = opAnd{$1, $3}
		}
|		val OR val {
			$$ = opOr{$1, $3}
		}
|		val RSHIFT val {
			$$ = opRShift{$1, $3}
		}
|		val LSHIFT val {
			$$ = opLShift{$1, $3}
		}
|		NOT val {
			$$ = opNot{$2}
		}
;

val:	LITERAL {
			$$ = $1
		}
|		IDENTIFIER {
			$$ = $1
		}
;

%%
