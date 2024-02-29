package compiler

import (
	"fmt"
	"monkey/ast"
	"monkey/code"
	"monkey/object"
	"sort"
)

type EmittedInstruction struct {
	OpCode   code.OpCode
	Position int
}

type CompilationScope struct {
	instructions    code.Instructions
	lastInstruction EmittedInstruction
	prevInstruction EmittedInstruction
}

type Compiler struct {
	constants []object.Object

	symbolTable *SymbolTable

	scopes     []CompilationScope
	scopeIndex int
}

type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

func New() *Compiler {
	mainScope := CompilationScope{
		instructions:    code.Instructions{},
		lastInstruction: EmittedInstruction{},
		prevInstruction: EmittedInstruction{},
	}

	return &Compiler{
		constants:   []object.Object{},
		symbolTable: NewSymbolTable(),
		scopes:      []CompilationScope{mainScope},
		scopeIndex:  0,
	}
}

func NewWithState(constants []object.Object, symbolTable *SymbolTable) *Compiler {
	c := New()
	c.constants = constants
	c.symbolTable = symbolTable
	return c
}

func (c *Compiler) curInstructions() code.Instructions {
	return c.scopes[c.scopeIndex].instructions
}

func (c *Compiler) enterScope() {
	scope := CompilationScope{
		instructions:    code.Instructions{},
		lastInstruction: EmittedInstruction{},
		prevInstruction: EmittedInstruction{},
	}

	c.scopes = append(c.scopes, scope)
	c.scopeIndex++
}

func (c *Compiler) leaveScope() code.Instructions {
	ins := c.curInstructions()

	c.scopes = c.scopes[:len(c.scopes)-1]
	c.scopeIndex--

	return ins
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
		c.removeLastPop()

		jPos := c.emit(code.OpJump, 9999)

		c.changeOperands(jntPos, len(c.curInstructions()))

		if node.Alternative == nil {
			c.emit(code.OpNull)
		} else {
			err := c.Compile(node.Alternative)
			if err != nil {
				return err
			}
			c.removeLastPop()
		}

		c.changeOperands(jPos, len(c.curInstructions()))

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

	case *ast.HashLiteral:
		keys := []ast.Expression{}
		for k := range node.Pairs {
			keys = append(keys, k)
		}

		sort.Slice(keys, func(i, j int) bool {
			return keys[i].String() < keys[j].String()
		})

		for _, k := range keys {
			err := c.Compile(k)
			if err != nil {
				return err
			}

			err = c.Compile(node.Pairs[k])
			if err != nil {
				return err
			}
		}

		c.emit(code.OpHash, len(node.Pairs)*2)

	case *ast.IndexExpression:
		err := c.Compile(node.Left)
		if err != nil {
			return err
		}

		err = c.Compile(node.Index)
		if err != nil {
			return err
		}

		c.emit(code.OpIndex)

	case *ast.FunctionLiteral:
		c.enterScope()

		err := c.Compile(node.Body)
		if err != nil {
			return err
		}

		c.replaceLastPopWithReturn()

		if !c.lastInstructionIs(code.OpReturnValue) {
			c.emit(code.OpReturn)
		}

		ins := c.leaveScope()

		cf := &object.CompiledFunction{Instructions: ins}
		c.emit(code.OpConstant, c.addConstant(cf))

	case *ast.ReturnStatement:
		err := c.Compile(node.ReturnValue)
		if err != nil {
			return err
		}

		c.emit(code.OpReturnValue)

	}

	return nil
}

func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.curInstructions(),
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

func (c *Compiler) addInstruction(newIns []byte) int {
	ins := c.curInstructions()
	pos := len(ins)
	c.scopes[c.scopeIndex].instructions = append(ins, newIns...)
	return pos
}

func (c *Compiler) setLastInstruction(op code.OpCode, pos int) {
	c.scopes[c.scopeIndex].prevInstruction = c.scopes[c.scopeIndex].lastInstruction
	c.scopes[c.scopeIndex].lastInstruction = EmittedInstruction{OpCode: op, Position: pos}
}

func (c *Compiler) lastInstructionIs(op code.OpCode) bool {
	if len(c.curInstructions()) == 0 {
		return false
	}

	return c.scopes[c.scopeIndex].lastInstruction.OpCode == op
}

func (c *Compiler) removeLastPop() {
	if c.scopes[c.scopeIndex].lastInstruction.OpCode == code.OpPop {
		c.scopes[c.scopeIndex].instructions = c.scopes[c.scopeIndex].instructions[:c.scopes[c.scopeIndex].lastInstruction.Position]
		c.scopes[c.scopeIndex].lastInstruction = c.scopes[c.scopeIndex].prevInstruction
	}
}

func (c *Compiler) replaceLastPopWithReturn() {
	if c.scopes[c.scopeIndex].lastInstruction.OpCode == code.OpPop {
		c.replaceInstruction(c.scopes[c.scopeIndex].lastInstruction.Position, code.Make(code.OpReturnValue))
		c.scopes[c.scopeIndex].lastInstruction.OpCode = code.OpReturnValue
	}
}

func (c *Compiler) replaceInstruction(pos int, newIns []byte) {
	ins := c.curInstructions()
	for i := 0; i < len(newIns); i++ {
		ins[pos+i] = newIns[i]
	}
}

func (c *Compiler) changeOperands(pos int, operands ...int) {
	op := code.OpCode(c.curInstructions()[pos])
	ins := code.Make(op, operands...)
	c.replaceInstruction(pos, ins)
}
