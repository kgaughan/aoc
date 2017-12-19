%{
package day7

import (
	"fmt"
)
%}

%union {
	node Node
}

%token <node> ASSIGN
%token <node> NOT AND OR
%token <node> LSHIFT RSHIFT
%token <node> IDENTIFIER LITERAL

%type <node> stmts expr val

%%

top:	stmts {
			fmt.Println($1.(Context).EvaluateID("a"))
		}

stmts:	/* empty */ {
			$$ = NewContext()
		}
|		stmts expr ASSIGN IDENTIFIER {
			$$.(Context).Add($4.String(), $2)
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
