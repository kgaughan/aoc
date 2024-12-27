package day7

type Identifier struct {
	id string
}

func (e Identifier) String() string {
	return e.id
}

func (e Identifier) Evaluate(ctx Context) Value {
	return ctx.EvaluateID(e.id)
}
