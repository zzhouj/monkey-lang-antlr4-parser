package compiler

import (
	"fmt"
	"monkey/ast"
	"monkey/code"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

type compilerTestCase struct {
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpAdd),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1; 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpPop),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 - 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpSub),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 * 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpMul),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 / 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpDiv),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "-1",
			expectedConstants: []interface{}{1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpMinus),
				code.Make(code.OpPop),
			},
		},
	}

	runCompilerTests(t, tests)
}

func TestBooleanExpression(t *testing.T) {
	tests := []compilerTestCase{
		{
			input:             "true",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpTrue),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "false",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpFalse),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 > 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpGT),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 < 2",
			expectedConstants: []interface{}{2, 1},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpGT),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 == 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpEQ),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "1 != 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
				code.Make(code.OpNE),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "true == false",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpTrue),
				code.Make(code.OpFalse),
				code.Make(code.OpEQ),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "true != false",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpTrue),
				code.Make(code.OpFalse),
				code.Make(code.OpNE),
				code.Make(code.OpPop),
			},
		},
		{
			input:             "!true",
			expectedConstants: []interface{}{},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpTrue),
				code.Make(code.OpBang),
				code.Make(code.OpPop),
			},
		},
	}

	runCompilerTests(t, tests)
}

func TestConditions(t *testing.T) {
	runCompilerTests(t, []compilerTestCase{
		{
			`if (true) { 10 }; 3333;`,
			[]interface{}{10, 3333},
			[]code.Instructions{
				// 0
				code.Make(code.OpTrue),
				// 1
				code.Make(code.OpJumpNotTruthy, 7),
				// 4
				code.Make(code.OpConstant, 0),
				// 7
				code.Make(code.OpPop),
				// 8
				code.Make(code.OpConstant, 1),
				// 11
				code.Make(code.OpPop),
			},
		},
		{
			`if (true) { 10 } else { 20 }; 3333;`,
			[]interface{}{10, 20, 3333},
			[]code.Instructions{
				// 0
				code.Make(code.OpTrue),
				// 1
				code.Make(code.OpJumpNotTruthy, 10),
				// 4
				code.Make(code.OpConstant, 0),
				// 7
				code.Make(code.OpJump, 13),
				// 10
				code.Make(code.OpConstant, 1),
				// 13
				code.Make(code.OpPop),
				// 14
				code.Make(code.OpConstant, 2),
				// 17
				code.Make(code.OpPop),
			},
		},
	})
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		compiler := New()
		err := compiler.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %v", err)
		}

		bytecode := compiler.ByteCode()

		err = testInstructions(tt.expectedInstructions, bytecode.Instructions)
		if err != nil {
			t.Fatalf("testInstructions failed: %v", err)
		}

		err = testConstants(tt.expectedConstants, bytecode.Constants)
		if err != nil {
			t.Fatalf("testConstants failed: %v", err)
		}
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testInstructions(expected []code.Instructions, actual code.Instructions) error {
	concatted := concatInstructions(expected)

	if len(actual) != len(concatted) {
		return fmt.Errorf("instructions have wrong length.\nwant=%q\ngot=%q", concatted, actual)
	}

	for i, ins := range concatted {
		if actual[i] != ins {
			return fmt.Errorf("wrong instruction at pos %d.\nwant=%q\ngot=%q", i, concatted, actual)
		}
	}

	return nil
}

func concatInstructions(instructions []code.Instructions) code.Instructions {
	out := code.Instructions{}

	for _, ins := range instructions {
		out = append(out, ins...)
	}

	return out
}

func testConstants(expected []interface{}, actual []object.Object) error {
	if len(expected) != len(actual) {
		return fmt.Errorf("wrong number of constants. want=%d, got=%d", len(expected), len(actual))
	}

	for i, constant := range expected {
		switch constant := constant.(type) {
		case int:
			err := testIntegerObject(int64(constant), actual[i])
			if err != nil {
				return fmt.Errorf("constant %d, testIntegerObject failed: %v", i, err)
			}
		}
	}

	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. want=%d, got=%d", expected, result.Value)
	}

	return nil
}
