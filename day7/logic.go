package day7

import (
	"fmt"
)

// An alias so if we need to change this, we can.
type Value uint16

type Node interface {
	Evaluate(ctx Context) Value
	String() string
}

type Literal struct {
	v Value
}

func (e Literal) String() string {
	return fmt.Sprintf("%v", e.v)
}

func (e Literal) Evaluate(ctx Context) Value {
	return e.v
}

type Identifier struct {
	id string
}

func (e Identifier) String() string {
	return e.id
}

func (e Identifier) Evaluate(ctx Context) Value {
	return ctx.EvaluateID(e.id)
}

type opAnd struct {
	left, right Node
}

func (op opAnd) String() string {
	return fmt.Sprintf("(and %v %v)", op.left, op.right)
}

func (op opAnd) Evaluate(ctx Context) Value {
	return op.left.Evaluate(ctx) & op.right.Evaluate(ctx)
}

type opOr struct {
	left, right Node
}

func (op opOr) String() string {
	return fmt.Sprintf("(or %v %v)", op.left, op.right)
}

func (op opOr) Evaluate(ctx Context) Value {
	return op.left.Evaluate(ctx) | op.right.Evaluate(ctx)
}

type opNot struct {
	e Node
}

func (op opNot) String() string {
	return fmt.Sprintf("(not %v)", op.e)
}

func (op opNot) Evaluate(ctx Context) Value {
	return ^op.e.Evaluate(ctx)
}

type opLShift struct {
	left, right Node
}

func (op opLShift) String() string {
	return fmt.Sprintf("(<< %v %v)", op.left, op.right)
}

func (op opLShift) Evaluate(ctx Context) Value {
	return op.left.Evaluate(ctx) << op.right.Evaluate(ctx)
}

type opRShift struct {
	left, right Node
}

func (op opRShift) String() string {
	return fmt.Sprintf("(>> %v %v)", op.left, op.right)
}

func (op opRShift) Evaluate(ctx Context) Value {
	return op.left.Evaluate(ctx) >> op.right.Evaluate(ctx)
}

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

func (_ Context) Evaluate(_ Context) Value {
	return 0
}
