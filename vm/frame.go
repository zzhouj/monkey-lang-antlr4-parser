package vm

import (
	"monkey/code"
	"monkey/object"
)

type Frame struct {
	cl   *object.Closure
	ip   int
	base int
}

func NewFrame(cl *object.Closure, base int) *Frame {
	return &Frame{
		cl:   cl,
		ip:   -1,
		base: base,
	}
}

func (f *Frame) Instructions() code.Instructions {
	return f.cl.Fn.Instructions
}
