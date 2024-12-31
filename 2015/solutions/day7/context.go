package day7

type Context struct {
	symbols map[string]Node
}

func NewContext() Context {
	return Context{
		symbols: make(map[string]Node),
	}
}

func (ctx Context) EvaluateID(id string) Value {
	node := ctx.symbols[id]
	v := node.Evaluate(ctx)
	switch node.(type) {
	default:
		ctx.symbols[id] = Literal{v}
		return v
	case Literal:
		return v
	}
}

func (ctx Context) Add(id string, node Node) {
	ctx.symbols[id] = node
}

func (ctx Context) String() string {
	return "<script>"
}

func (Context) Evaluate(_ Context) Value {
	return 0
}
