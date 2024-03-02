package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type OpCode byte

const (
	OpConstant OpCode = iota
	OpPop
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpTrue
	OpFalse
	OpEQ
	OpNE
	OpGT
	OpMinus
	OpBang
	OpJumpNotTruthy
	OpJump
	OpNull
	OpSetGlobal
	OpGetGlobal
	OpArray
	OpHash
	OpIndex
	OpCall
	OpReturnValue
	OpReturn
	OpSetLocal
	OpGetLocal
	OpGetBuiltin
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var difinitions = map[OpCode]*Definition{
	OpConstant:      {"OpConstant", []int{2}},
	OpPop:           {"OpPop", []int{}},
	OpAdd:           {"OpAdd", []int{}},
	OpSub:           {"OpSub", []int{}},
	OpMul:           {"OpAdd", []int{}},
	OpDiv:           {"OpAdd", []int{}},
	OpTrue:          {"OpTrue", []int{}},
	OpFalse:         {"OpFalse", []int{}},
	OpEQ:            {"OpEQ", []int{}},
	OpNE:            {"OpNE", []int{}},
	OpGT:            {"OpGT", []int{}},
	OpMinus:         {"OpMinus", []int{}},
	OpBang:          {"OpBang", []int{}},
	OpJumpNotTruthy: {"OpJumpNotTruthy", []int{2}},
	OpJump:          {"OpJump", []int{2}},
	OpNull:          {"OpNull", []int{}},
	OpSetGlobal:     {"OpSetGlobal", []int{2}},
	OpGetGlobal:     {"OpGetGlobal", []int{2}},
	OpArray:         {"OpArray", []int{2}},
	OpHash:          {"OpHash", []int{2}},
	OpIndex:         {"OpIndex", []int{}},
	OpCall:          {"OpCall", []int{1}},
	OpReturnValue:   {"OpReturnValue", []int{}},
	OpReturn:        {"OpReturn", []int{}},
	OpSetLocal:      {"OpSetLocal", []int{1}},
	OpGetLocal:      {"OpGetLocal", []int{1}},
	OpGetBuiltin:    {"OpGetBuiltin", []int{1}},
}

func Lookup(op byte) (*Definition, error) {
	def, ok := difinitions[OpCode(op)]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", op)
	}
	return def, nil
}

func Make(op OpCode, operands ...int) []byte {
	def, ok := difinitions[op]
	if !ok {
		return []byte{}
	}

	insLen := 1
	for _, w := range def.OperandWidths {
		insLen += w
	}

	instruction := make([]byte, insLen)
	instruction[0] = byte(op)

	offset := 1
	for i, w := range def.OperandWidths {
		var o int
		if i < len(operands) {
			o = operands[i]
		}

		switch w {
		case 1:
			instruction[offset] = byte(o)
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}

		offset += w
	}

	return instruction
}

func (ins Instructions) String() string {
	var out bytes.Buffer

	for i := 0; i < len(ins); {
		def, err := Lookup(ins[i])
		if err != nil {
			fmt.Fprintf(&out, "ERROR: %v\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[i+1:])
		fmt.Fprintf(&out, "%04d %s\n", i, fmtInstruction(def, operands))

		i += read + 1
	}

	return out.String()
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, w := range def.OperandWidths {
		switch w {
		case 1:
			operands[i] = int(ReadUint8(ins[offset:]))
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += w
	}

	return operands, offset
}

func ReadUint8(ins Instructions) uint8 { return uint8(ins[0]) }

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}

func fmtInstruction(def *Definition, operands []int) string {
	if len(def.OperandWidths) != len(operands) {
		return fmt.Sprintf("ERROR: wrong number of operands. want=%d, got=%d", len(def.OperandWidths), len(operands))
	}

	switch len(operands) {
	case 0:
		return def.Name
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: not handled number of operands %d for %s", len(operands), def.Name)
}
