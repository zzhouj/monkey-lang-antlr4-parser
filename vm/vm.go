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

func (vm *VM) LastPopped() object.Object {
	return vm.stack[vm.sp]
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

		case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}

		case code.OpPop:
			vm.pop()
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

func (vm *VM) executeBinaryOperation(op code.OpCode) error {
	right := vm.pop()
	left := vm.pop()
	if right == nil || left == nil {
		return fmt.Errorf("binary operation missing operands")
	}

	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return vm.executeBinaryIntegerOperation(op, left, right)

	default:
		return fmt.Errorf("unsupported types for binary operation: %s %s", left.Type(), right.Type())
	}

}

func (vm *VM) executeBinaryIntegerOperation(op code.OpCode, left, right object.Object) error {
	lv := left.(*object.Integer).Value
	rv := right.(*object.Integer).Value

	var result int64

	switch op {
	case code.OpAdd:
		result = lv + rv

	case code.OpSub:
		result = lv - rv

	case code.OpMul:
		result = lv * rv

	case code.OpDiv:
		result = lv / rv

	default:
		return fmt.Errorf("unknown integer operation: op code %d", op)
	}

	return vm.push(&object.Integer{Value: result})
}
