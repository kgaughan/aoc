package lib

import (
	"fmt"
)

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
