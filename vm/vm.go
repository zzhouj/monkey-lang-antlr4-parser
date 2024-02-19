package vm

import (
	"fmt"
	"monkey/code"
	"monkey/compiler"
	"monkey/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object
	sp    int
}

func New(bc *compiler.ByteCode) *VM {
	return &VM{
		constants:    bc.Constants,
		instructions: bc.Instructions,
		stack:        make([]object.Object, StackSize),
		sp:           0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp > 0 {
		return vm.stack[vm.sp-1]
	}
	return nil
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		op := code.OpCode(vm.instructions[ip])

		switch op {
		case code.OpConstant:
			constIdx := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			if int(constIdx) >= len(vm.constants) {
				return fmt.Errorf("invalid constant index: %d", constIdx)
			}

			if err := vm.push(vm.constants[constIdx]); err != nil {
				return err
			}

		case code.OpAdd:
			right := vm.pop()
			left := vm.pop()
			if right == nil || left == nil {
				return fmt.Errorf("OpAdd missing operands")
			}

			lv := left.(*object.Integer).Value
			rv := right.(*object.Integer).Value

			result := lv + rv
			vm.push(&object.Integer{Value: result})
		}
	}

	return nil
}

func (vm *VM) push(obj object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("stack overflow")
	}

	vm.stack[vm.sp] = obj
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	if vm.sp > 0 {
		obj := vm.stack[vm.sp-1]
		vm.sp--
		return obj
	}

	return nil
}
