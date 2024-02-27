package vm

import (
	"fmt"
	"monkey/code"
	"monkey/compiler"
	"monkey/object"
)

var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

const (
	StackSize   = 2048
	GlobalsSize = 65536
)

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack   []object.Object
	sp      int
	globals []object.Object
}

func New(bc *compiler.ByteCode) *VM {
	return &VM{
		constants:    bc.Constants,
		instructions: bc.Instructions,
		stack:        make([]object.Object, StackSize),
		sp:           0,
		globals:      make([]object.Object, GlobalsSize),
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

			err := vm.push(vm.constants[constIdx])
			if err != nil {
				return err
			}

		case code.OpPop:
			vm.pop()

		case code.OpAdd, code.OpSub, code.OpMul, code.OpDiv:
			err := vm.executeBinaryOperation(op)
			if err != nil {
				return err
			}

		case code.OpTrue:
			err := vm.push(TRUE)
			if err != nil {
				return err
			}

		case code.OpFalse:
			err := vm.push(FALSE)
			if err != nil {
				return err
			}

		case code.OpEQ, code.OpNE, code.OpGT:
			err := vm.executeComparison(op)
			if err != nil {
				return err
			}

		case code.OpBang:
			err := vm.executeBangOperator()
			if err != nil {
				return err
			}

		case code.OpMinus:
			err := vm.executeMinusOperator()
			if err != nil {
				return err
			}

		case code.OpJump:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip = pos - 1

		case code.OpJumpNotTruthy:
			pos := int(code.ReadUint16(vm.instructions[ip+1:]))
			ip += 2

			condition := vm.pop()
			if !isTruthy(condition) {
				ip = pos - 1
			}

		case code.OpNull:
			err := vm.push(NULL)
			if err != nil {
				return err
			}

		case code.OpSetGlobal:
			idx := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			vm.globals[idx] = vm.pop()

		case code.OpGetGlobal:
			idx := code.ReadUint16(vm.instructions[ip+1:])
			ip += 2

			err := vm.push(vm.globals[idx])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj.Value
	case *object.Null:
		return false
	default:
		return true
	}
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

func (vm *VM) executeComparison(op code.OpCode) error {
	right := vm.pop()
	left := vm.pop()
	if right == nil || left == nil {
		return fmt.Errorf("comparison missing operands")
	}

	if left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ {
		return vm.executeIntegerComparison(op, left, right)
	}

	switch op {
	case code.OpEQ:
		return vm.push(nativeBoolToBooleanObject(left == right))
	case code.OpNE:
		return vm.push(nativeBoolToBooleanObject(left != right))
	default:
		return fmt.Errorf("unsupported types for comparison: %s %s", left.Type(), right.Type())
	}
}

func (vm *VM) executeIntegerComparison(op code.OpCode, left, right object.Object) error {
	lv := left.(*object.Integer).Value
	rv := right.(*object.Integer).Value

	var result bool

	switch op {
	case code.OpEQ:
		result = lv == rv

	case code.OpNE:
		result = lv != rv

	case code.OpGT:
		result = lv > rv

	default:
		return fmt.Errorf("unknown integer comparison: op code %d", op)
	}

	return vm.push(nativeBoolToBooleanObject(result))
}

func nativeBoolToBooleanObject(nb bool) object.Object {
	if nb {
		return TRUE
	} else {
		return FALSE
	}
}

func (vm *VM) executeBangOperator() error {
	operand := vm.pop()
	if operand == nil {
		return fmt.Errorf("bang operator missing operand")
	}

	switch operand {
	case FALSE, NULL:
		return vm.push(TRUE)
	default:
		return vm.push(FALSE)
	}
}

func (vm *VM) executeMinusOperator() error {
	operand := vm.pop()
	if operand == nil {
		return fmt.Errorf("minus operator missing operand")
	}

	if operand.Type() != object.INTEGER_OBJ {
		return fmt.Errorf("unsupported types for minus operator: %s", operand.Type())
	}

	value := operand.(*object.Integer).Value
	return vm.push(&object.Integer{Value: -value})
}
