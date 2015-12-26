package day7

// An alias so if we need to change this, we can.
type Value uint16

type Node interface {
	Evaluate(idMap map[string]Node) Value
}

type Literal struct {
	v Value
}

func (e Literal) Evaluate(idMap map[string]Node) Value {
	return e.v
}

type Identifier struct {
	id string
}

func (e Identifier) Evaluate(idMap map[string]Node) Value {
	return idMap[e.id].Evaluate(idMap)
}

type opAnd struct {
	left, right Node
}

func (op opAnd) Evaluate(idMap map[string]Node) Value {
	return op.left.Evaluate(idMap) & op.right.Evaluate(idMap)
}

type opOr struct {
	left, right Node
}

func (op opOr) Evaluate(idMap map[string]Node) Value {
	return op.left.Evaluate(idMap) | op.right.Evaluate(idMap)
}

type opNot struct {
	e Node
}

func (op opNot) Evaluate(idMap map[string]Node) Value {
	return ^op.e.Evaluate(idMap)
}

type opLShift struct {
	left, right Node
}

func (op opLShift) Evaluate(idMap map[string]Node) Value {
	return op.left.Evaluate(idMap) << op.right.Evaluate(idMap)
}

type opRShift struct {
	left, right Node
}

func (op opRShift) Evaluate(idMap map[string]Node) Value {
	return op.left.Evaluate(idMap) >> op.right.Evaluate(idMap)
}
