package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type OpCode byte

const (
	OpConstant OpCode = iota
)

type Definition struct {
	Name          string
	OperandWidths []int
}

var difinitions = map[OpCode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
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
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(o))
		}

		offset += w
	}

	return instruction
}
