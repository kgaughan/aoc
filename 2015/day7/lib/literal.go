package lib

import (
	"fmt"
)

type Literal struct {
	v Value
}

func (e Literal) String() string {
	return fmt.Sprintf("%v", e.v)
}

func (e Literal) Evaluate(ctx Context) Value {
	return e.v
}
