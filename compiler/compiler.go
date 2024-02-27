package compiler

import (
	"fmt"
	"monkey/ast"
	"monkey/code"
	"monkey/object"
)

type EmittedInstruction struct {
	OpCode   code.OpCode
	Position int
}

type Compiler struct {
	instructions    code.Instructions
	constants       []object.Object
	lastInstruction EmittedInstruction
	prevInstruction EmittedInstruction
	symbolTable     *SymbolTable
}

type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

func New() *Compiler {
	return &Compiler{
		instructions: code.Instructions{},
		constants:    []object.Object{},
		symbolTable:  NewSymbolTable(),
	}
}

func NewWithState(constants []object.Object, symbolTable *SymbolTable) *Compiler {
	c := New()
	c.constants = constants
	c.symbolTable = symbolTable
	return c
}

func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, s := range node.Statements {
			err := c.Compile(s)
			if err != nil {
				return err
			}
		}

	case *ast.ExpressionStatement:
		err := c.Compile(node.Expression)
		if err != nil {
			return err
		}
		c.emit(code.OpPop)

	case *ast.InfixExpression:
		if node.Operator == "<" {
			err := c.Compile(node.Right)
			if err != nil {
				return err
			}

			err = c.Compile(node.Left)
			if err != nil {
				return err
			}

			c.emit(code.OpGT)
			return nil
		}

		err := c.Compile(node.Left)
		if err != nil {
			return err
		}

		err = c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "+":
			c.emit(code.OpAdd)
		case "-":
			c.emit(code.OpSub)
		case "*":
			c.emit(code.OpMul)
		case "/":
			c.emit(code.OpDiv)
		case "==":
			c.emit(code.OpEQ)
		case "!=":
			c.emit(code.OpNE)
		case ">":
			c.emit(code.OpGT)
		default:
			return fmt.Errorf("unknown operator %q", node.Operator)
		}

	case *ast.PrefixExpression:
		err := c.Compile(node.Right)
		if err != nil {
			return err
		}

		switch node.Operator {
		case "!":
			c.emit(code.OpBang)
		case "-":
			c.emit(code.OpMinus)
		default:
			return fmt.Errorf("unknown operator %q", node.Operator)
		}

	case *ast.IntegerLiteral:
		intObj := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(intObj))

	case *ast.StringLiteral:
		strObj := &object.String{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(strObj))

	case *ast.BooleanLiteral:
		if node.Value {
			c.emit(code.OpTrue)
		} else {
			c.emit(code.OpFalse)
		}

	case *ast.IfExpression:
		err := c.Compile(node.Condition)
		if err != nil {
			return err
		}

		jntPos := c.emit(code.OpJumpNotTruthy, 9999)

		err = c.Compile(node.Consequence)
		if err != nil {
			return err
		}
		c.removeLastPopIf()

		jPos := c.emit(code.OpJump, 9999)

		c.changeOperands(jntPos, len(c.instructions))

		if node.Alternative == nil {
			c.emit(code.OpNull)
		} else {
			err := c.Compile(node.Alternative)
			if err != nil {
				return err
			}
			c.removeLastPopIf()
		}

		c.changeOperands(jPos, len(c.instructions))

	case *ast.BlockStatement:
		for _, stmt := range node.Statements {
			err := c.Compile(stmt)
			if err != nil {
				return err
			}
		}

	case *ast.LetStatement:
		err := c.Compile(node.Value)
		if err != nil {
			return err
		}
		symbol := c.symbolTable.Define(node.Name.Value)
		c.emit(code.OpSetGlobal, symbol.Index)

	case *ast.Identifier:
		symbol, ok := c.symbolTable.Resolve(node.Value)
		if !ok {
			return fmt.Errorf("undefined variable %q", node.Value)
		}
		c.emit(code.OpGetGlobal, symbol.Index)

	case *ast.ArrayLiteral:
		for _, el := range node.Elements {
			err := c.Compile(el)
			if err != nil {
				return err
			}
		}
		c.emit(code.OpArray, len(node.Elements))
	}

	return nil
}

func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}

func (c *Compiler) emit(op code.OpCode, operands ...int) int {
	ins := code.Make(op, operands...)
	pos := c.addInstruction(ins)

	c.setLastInstruction(op, pos)

	return pos
}

func (c *Compiler) addInstruction(ins []byte) int {
	pos := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return pos
}

func (c *Compiler) setLastInstruction(op code.OpCode, pos int) {
	c.prevInstruction = c.lastInstruction
	c.lastInstruction = EmittedInstruction{OpCode: op, Position: pos}
}

func (c *Compiler) removeLastPopIf() {
	if c.lastInstruction.OpCode == code.OpPop {
		c.instructions = c.instructions[:c.lastInstruction.Position]
		c.lastInstruction = c.prevInstruction
	}
}

func (c *Compiler) replaceInstruction(pos int, instruction []byte) {
	for i := 0; i < len(instruction); i++ {
		c.instructions[pos+i] = instruction[i]
	}
}

func (c *Compiler) changeOperands(pos int, operands ...int) {
	op := code.OpCode(c.instructions[pos])
	ins := code.Make(op, operands...)
	c.replaceInstruction(pos, ins)
}
