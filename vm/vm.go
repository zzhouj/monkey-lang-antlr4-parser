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
	MaxFrames   = 1024
)

type VM struct {
	constants []object.Object

	stack []object.Object
	sp    int

	globals []object.Object

	frames     []*Frame
	frameIndex int
}

func New(bc *compiler.ByteCode) *VM {
	mainFn := &object.CompiledFunction{Instructions: bc.Instructions}
	mainClosure := &object.Closure{Fn: mainFn}
	mainFrame := NewFrame(mainClosure, 0)

	frames := make([]*Frame, MaxFrames)
	frames[0] = mainFrame

	return &VM{
		constants: bc.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,

		globals: make([]object.Object, GlobalsSize),

		frames:     frames,
		frameIndex: 0,
	}
}

func NewWithState(bc *compiler.ByteCode, globals []object.Object) *VM {
	vm := New(bc)
	vm.globals = globals
	return vm
}

func (vm *VM) curFrame() *Frame {
	return vm.frames[vm.frameIndex]
}

func (vm *VM) pushFrame(f *Frame) {
	vm.frameIndex++
	vm.frames[vm.frameIndex] = f
	vm.sp = f.base + f.cl.Fn.NumLocals
}

func (vm *VM) popFrame() *Frame {
	f := vm.frames[vm.frameIndex]
	vm.frameIndex--
	vm.sp = f.base - 1
	return f
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
	var (
		ip  int
		ins code.Instructions
		op  code.OpCode
	)

	for vm.curFrame().ip < len(vm.curFrame().Instructions())-1 {
		vm.curFrame().ip++

		ip = vm.curFrame().ip
		ins = vm.curFrame().Instructions()
		op = code.OpCode(ins[ip])

		switch op {
		case code.OpConstant:
			constIdx := code.ReadUint16(ins[ip+1:])
			vm.curFrame().ip += 2

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
			pos := int(code.ReadUint16(ins[ip+1:]))
			vm.curFrame().ip = pos - 1

		case code.OpJumpNotTruthy:
			pos := int(code.ReadUint16(ins[ip+1:]))
			vm.curFrame().ip += 2

			condition := vm.pop()
			if !isTruthy(condition) {
				vm.curFrame().ip = pos - 1
			}

		case code.OpNull:
			err := vm.push(NULL)
			if err != nil {
				return err
			}

		case code.OpSetGlobal:
			idx := code.ReadUint16(ins[ip+1:])
			vm.curFrame().ip += 2

			vm.globals[idx] = vm.pop()

		case code.OpGetGlobal:
			idx := code.ReadUint16(ins[ip+1:])
			vm.curFrame().ip += 2

			err := vm.push(vm.globals[idx])
			if err != nil {
				return err
			}

		case code.OpArray:
			numElem := int(code.ReadUint16(ins[ip+1:]))
			vm.curFrame().ip += 2

			arr := vm.buildArray(vm.sp-numElem, vm.sp)
			vm.sp -= numElem

			err := vm.push(arr)
			if err != nil {
				return err
			}

		case code.OpHash:
			numElem := int(code.ReadUint16(ins[ip+1:]))
			vm.curFrame().ip += 2

			hash, err := vm.buildHash(vm.sp-numElem, vm.sp)
			if err != nil {
				return err
			}
			vm.sp -= numElem

			err = vm.push(hash)
			if err != nil {
				return err
			}

		case code.OpIndex:
			index := vm.pop()
			left := vm.pop()

			err := vm.executeIndexExpression(left, index)
			if err != nil {
				return err
			}

		case code.OpCall:
			numArgs := int(code.ReadUint8(ins[ip+1:]))
			vm.curFrame().ip += 1

			err := vm.executeCall(numArgs)
			if err != nil {
				return err
			}

		case code.OpReturnValue, code.OpReturn:
			var retVal object.Object
			if op == code.OpReturnValue {
				retVal = vm.pop()
			} else {
				retVal = NULL
			}

			vm.popFrame()

			err := vm.push(retVal)
			if err != nil {
				return err
			}

		case code.OpSetLocal:
			idx := int(code.ReadUint8(ins[ip+1:]))
			vm.curFrame().ip += 1

			vm.stack[vm.curFrame().base+idx] = vm.pop()

		case code.OpGetLocal:
			idx := int(code.ReadUint8(ins[ip+1:]))
			vm.curFrame().ip += 1

			err := vm.push(vm.stack[vm.curFrame().base+idx])
			if err != nil {
				return err
			}

		case code.OpGetBuiltin:
			idx := int(code.ReadUint8(ins[ip+1:]))
			vm.curFrame().ip += 1

			builtin := object.Builtins[idx]
			err := vm.push(builtin.Builtin)
			if err != nil {
				return err
			}

		case code.OpClosure:
			constIdx := int(code.ReadUint16(ins[ip+1:]))
			_ = int(code.ReadUint8(ins[ip+3:]))
			vm.curFrame().ip += 3

			err := vm.pushClosure(constIdx)
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

	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return vm.executeBinaryStringOperation(op, left, right)

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

func (vm *VM) executeBinaryStringOperation(op code.OpCode, left, right object.Object) error {
	if op != code.OpAdd {
		return fmt.Errorf("unknown string operation: op code %d", op)
	}

	lv := left.(*object.String).Value
	rv := right.(*object.String).Value

	return vm.push(&object.String{Value: lv + rv})
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

func (vm *VM) buildArray(start, end int) object.Object {
	elements := make([]object.Object, end-start)

	for i := start; i < end; i++ {
		elements[i-start] = vm.stack[i]
	}

	return &object.Array{Elements: elements}
}

func (vm *VM) buildHash(start, end int) (object.Object, error) {
	pairs := make(map[object.HashKey]object.HashPair)

	for i := start; i < end; i += 2 {
		k := vm.stack[i]
		v := vm.stack[i+1]

		hk, ok := k.(object.Hashable)
		if !ok {
			return nil, fmt.Errorf("unusable as hash key: %s", k.Type())
		}

		pairs[hk.HashKey()] = object.HashPair{Key: k, Value: v}
	}

	return &object.Hash{Pairs: pairs}, nil
}

func (vm *VM) executeIndexExpression(left, index object.Object) error {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return vm.executeArrayIndex(left, index)
	case left.Type() == object.HASH_OBJ:
		return vm.executeHashIndex(left, index)
	default:
		return fmt.Errorf("index operator not supported: %s[%s]", left.Type(), index.Type())
	}
}

func (vm *VM) executeArrayIndex(left, index object.Object) error {
	arr := left.(*object.Array)
	i := index.(*object.Integer).Value

	if i < 0 || i >= int64(len(arr.Elements)) {
		return vm.push(NULL)
	}

	return vm.push(arr.Elements[i])
}

func (vm *VM) executeHashIndex(left, index object.Object) error {
	hash := left.(*object.Hash)
	key, ok := index.(object.Hashable)
	if !ok {
		return fmt.Errorf("unusable as hash key: %s", index.Type())
	}

	pair, ok := hash.Pairs[key.HashKey()]
	if !ok {
		return vm.push(NULL)
	}

	return vm.push(pair.Value)
}

func (vm *VM) executeCall(numArgs int) error {
	callee := vm.stack[vm.sp-1-numArgs]
	switch callee := callee.(type) {
	case *object.Closure:
		return vm.callClosure(callee, numArgs)
	case *object.Builtin:
		return vm.callBuiltin(callee, numArgs)
	default:
		return fmt.Errorf("calling non-function: %T", callee)
	}
}

func (vm *VM) callClosure(cl *object.Closure, numArgs int) error {
	if numArgs != cl.Fn.NumParameters {
		return fmt.Errorf("wrong number of arguments: want=%d, got=%d", cl.Fn.NumParameters, numArgs)
	}

	frame := NewFrame(cl, vm.sp-numArgs)
	vm.pushFrame(frame)

	return nil
}

func (vm *VM) callBuiltin(builtin *object.Builtin, numArgs int) error {
	args := vm.stack[vm.sp-numArgs : vm.sp]

	result := builtin.Fn(args...)
	vm.sp -= numArgs + 1

	if result != nil {
		vm.push(result)
	} else {
		vm.push(NULL)
	}

	return nil
}

func (vm *VM) pushClosure(constIdx int) error {
	constant := vm.constants[constIdx]
	fn, ok := constant.(*object.CompiledFunction)
	if !ok {
		return fmt.Errorf("not a function: %T", constant)
	}

	closure := &object.Closure{Fn: fn}
	return vm.push(closure)
}
