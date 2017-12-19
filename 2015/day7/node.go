package day7

// An alias so if we need to change this, we can.
type Value uint16

type Node interface {
	Evaluate(ctx Context) Value
	String() string
}
