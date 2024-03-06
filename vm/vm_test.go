package vm

import (
	"fmt"
	"monkey/ast"
	"monkey/compiler"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"testing"
)

type vmTestCase struct {
	input    string
	expected interface{}
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %v", err)
		}

		vm := New(comp.ByteCode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("vm error: %v", err)
		}

		lastPopped := vm.LastPopped()
		testExpectedObject(t, tt.expected, lastPopped)
	}
}

func testExpectedObject(t *testing.T, expected interface{}, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(int64(expected), actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %v", err)
		}

	case []int:
		err := testIntegerArrayObject(expected, actual)
		if err != nil {
			t.Errorf("testIntegerArrayObject failed: %v", err)
		}

	case map[object.HashKey]int64:
		err := testIntegerHashObject(expected, actual)
		if err != nil {
			t.Errorf("testIntegerHashObject failed: %v", err)
		}

	case string:
		err := testStringObject(expected, actual)
		if err != nil {
			t.Errorf("testStringObject failed: %v", err)
		}

	case bool:
		err := testBooleanObject(expected, actual)
		if err != nil {
			t.Errorf("testBooleanObject failed: %v", err)
		}

	case *object.Null:
		if actual != NULL {
			t.Errorf("object is not Null: %T (%+v)", actual, actual)
		}

	case *object.Error:
		errObj, ok := actual.(*object.Error)
		if !ok {
			t.Errorf("object is not Error: %T (%+v)", actual, actual)
			return
		}
		if errObj.Message != expected.Message {
			t.Errorf("wrong error message: want=%q, got=%q", expected.Message, errObj.Message)
		}
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
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

func testIntegerArrayObject(expected []int, actual object.Object) error {
	result, ok := actual.(*object.Array)
	if !ok {
		return fmt.Errorf("object is not Array. got=%T (%+v)", actual, actual)
	}

	if len(result.Elements) != len(expected) {
		return fmt.Errorf("wrong num of elements. want=%d, got=%d", len(expected), len(result.Elements))
	}

	for i, el := range result.Elements {
		err := testIntegerObject(int64(expected[i]), el)
		if err != nil {
			return fmt.Errorf("[%d] testIntegerObject failed: %v", i, err)
		}
	}

	return nil
}

func testIntegerHashObject(expected map[object.HashKey]int64, actual object.Object) error {
	result, ok := actual.(*object.Hash)
	if !ok {
		return fmt.Errorf("object is not Hash. got=%T (%+v)", actual, actual)
	}

	if len(result.Pairs) != len(expected) {
		return fmt.Errorf("wrong num of pairs. want=%d, got=%d", len(expected), len(result.Pairs))
	}

	for key, pair := range result.Pairs {
		err := testIntegerObject(expected[key], pair.Value)
		if err != nil {
			return fmt.Errorf("[%s] testIntegerObject failed: %v", pair.Key.Inspect(), err)
		}
	}

	return nil
}

func testStringObject(expected string, actual object.Object) error {
	result, ok := actual.(*object.String)
	if !ok {
		return fmt.Errorf("object is not String. got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. want=%q, got=%q", expected, result.Value)
	}

	return nil
}

func testBooleanObject(expected bool, actual object.Object) error {
	result, ok := actual.(*object.Boolean)
	if !ok {
		return fmt.Errorf("object is not Boolean. got=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. want=%t, got=%t", expected, result.Value)
	}

	return nil
}

func TestIntegerArithmetic(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"1 * 2", 2},
		{"4 / 2", 2},
		{"50 / 2 * 2 + 10 - 5", 55},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"5 * (2 + 10)", 60},
		{"-5", -5},
		{"-10", -10},
		{"-50 + 100 + -50", 0},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	})
}

func TestBooleanExpression(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 1", true},
		{"1 != 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"true == true", true},
		{"true != true", false},
		{"true == false", false},
		{"true != false", true},
		{"1 < 2 == true", true},
		{"1 < 2 == false", false},
		{"1 > 2 == true", false},
		{"1 > 2 == false", true},
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
		{"!(if (false) { 10 })", true},
	})
}

func TestConditions(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{"if (true) { 10 }", 10},
		{"if (true) { 10 } else { 20 }", 10},
		{"if (false) { 10 } else { 20 }", 20},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
		{"if (1 > 2) { 10 }", NULL},
		{"if (false) { 10 }", NULL},
		{"if (if (false) { 10 }) { 10 } else { 20 }", 20},
	})
}

func TestGlobalLetStatements(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{"let one = 1; one", 1},
		{"let one = 1; let two = 2; one + two", 3},
		{"let one = 1; let two = one + one; one + two", 3},
	})
}

func TestStringExpressions(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{`"monkey"`, "monkey"},
		{`"mon" + "key"`, "monkey"},
		{`"mon" + "key" + "banana"`, "monkeybanana"},
	})
}

func TestArrayLiterals(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{"[]", []int{}},
		{"[1, 2, 3]", []int{1, 2, 3}},
		{"[1 + 2, 3 - 4, 5 * 6]", []int{3, -1, 30}},
	})
}

func TestHashLiterals(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{
			"{}",
			map[object.HashKey]int64{},
		},
		{
			"{1: 2, 3: 4}",
			map[object.HashKey]int64{
				(&object.Integer{Value: 1}).HashKey(): 2,
				(&object.Integer{Value: 3}).HashKey(): 4,
			},
		},
		{
			"{1: 2 + 3, 4 * 5: 6}",
			map[object.HashKey]int64{
				(&object.Integer{Value: 1}).HashKey():  5,
				(&object.Integer{Value: 20}).HashKey(): 6,
			},
		},
	})
}

func TestIndexExpressions(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{"[1, 2, 3][1]", 2},
		{"[1, 2, 3][0 + 2]", 3},
		{"[[1, 2, 3]][0][0]", 1},
		{"[][0]", NULL},
		{"[1, 2, 3][99]", NULL},
		{"[1][-1]", NULL},
		{"{1: 1, 2: 2}[1]", 1},
		{"{1: 1, 2: 2}[2]", 2},
		{"{1: 1}[0]", NULL},
		{"{}[0]", NULL},
	})
}

func TestFunctionCalls(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{
			`let fivePlusTen = fn() { 5 + 10 }; fivePlusTen();`,
			15,
		},
		{
			`let one = fn() { 1; }; let two = fn() { 2; }; one() + two();`,
			3,
		},
		{
			`let earlyExit = fn() { return 99; 100; }; earlyExit();`,
			99,
		},
		{
			`let earlyExit = fn() { return 99; return 100; }; earlyExit();`,
			99,
		},
		{
			`let noRet = fn() {}; noRet();`,
			NULL,
		},
		{
			`let noRet = fn() {}; let noRetTwo = fn() { noRet(); }; noRet(); noRetTwo();`,
			NULL,
		},
		{
			`let a = fn(){1}; let b = fn(){a}; b()();`,
			1,
		},
		{
			`let one = fn() { let one = 1; one}; one()`,
			1,
		},
		{
			`let oneAndTwo = fn() { let one = 1; let two = 2; one + two }; oneAndTwo()`,
			3,
		},
		{
			`
			let oneAndTwo = fn() { let one = 1; let two = 2; one + two };
			let threeAndFour = fn() { let three = 3; let four = 4; three + four };
			oneAndTwo() + threeAndFour()`,
			10,
		},
		{
			`
			let first = fn() { let foobar = 50; foobar };
			let second = fn() { let foobar = 100; foobar };
			first() + second()`,
			150,
		},
		{
			`
			let globalSeed = 50;
			let first = fn() { let num = 1; globalSeed - num };
			let second = fn() { let num = 2; globalSeed - num };
			first() + second()`,
			97,
		},
		{
			`let oneArg = fn(a) { a }; oneArg(24);`,
			24,
		},
		{
			`let manyArg = fn(a, b, c) { a + b + c }; manyArg(24, 25, 26);`,
			75,
		},
		{
			`let sum = fn(a, b) { let c = a + b; c }; sum(1, 2)`,
			3,
		},
		{
			`let sum = fn(a, b) { let c = a + b; c }; sum(1, 2) + sum(3, 4)`,
			10,
		},
		{
			`let sum = fn(a, b) { let c = a + b; c }; let outer = fn() { sum(1, 2) + sum(3, 4) }; outer()`,
			10,
		},
		{
			`let globalNum = 10; let sum = fn(a, b) { let c = a + b; c + globalNum }; let outer = fn() { sum(1, 2) + sum(3, 4) + globalNum }; outer() + globalNum`,
			50,
		},
		{
			`
			let a = 1;
			fn() {
				let b = 2;
				fn() {
					let c = 3;
					a + b + c;
				} 
			}()()`,
			6,
		},
	})
}

func TestFunctionCallsWithWrongArguments(t *testing.T) {
	tests := []vmTestCase{
		{
			`fn() { 1; }(1)`,
			`wrong number of arguments: want=0, got=1`,
		},
		{
			`fn(a) { a; }()`,
			`wrong number of arguments: want=1, got=0`,
		},
		{
			`fn(a, b) { a + b; }(1)`,
			`wrong number of arguments: want=2, got=1`,
		},
		{
			`fn(a, b) { a + b; }(1)`,
			`wrong number of arguments: want=2, got=1`,
		},
	}

	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.ByteCode())
		err = vm.Run()
		if err == nil {
			t.Fatalf("expected VM error but resulted in one.")
		}

		if err.Error() != tt.expected {
			t.Fatalf("wrong VM error: want=%q, got=%q", tt.expected, err.Error())
		}
	}
}

func TestBuiltinFunctions(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world!")`, 12},
		{`len([])`, 0},
		{`len([1, 2, 3])`, 3},
		{`first([1, 2, 3])`, 1},
		{`last([1, 2, 3])`, 3},
		{
			`rest([1, 2, 3, 4])`,
			[]int{2, 3, 4},
		},
		{
			`rest([4])`,
			[]int{},
		},
		{
			`rest([])`,
			NULL,
		},
		{
			`push([],1)`,
			[]int{1},
		},
		{
			`push([1], 2, 3, 4)`,
			[]int{1, 2, 3, 4},
		},
		{
			`let a = [1]; let b = push(a, 2); a`,
			[]int{1},
		},
		{
			`let a = [1]; let b = push(a, 2); b`,
			[]int{1, 2},
		},
		{
			`len(1)`,
			&object.Error{Message: "argument to `len` not supported, got INTEGER"},
		},
		{
			`len("one", "two")`,
			&object.Error{Message: "wrong number of arguments. got=2, want=1"},
		},

		{
			`first(1)`,
			&object.Error{Message: "argument to `first` should be ARRAY, got INTEGER"},
		},
		{
			`first("one", "two")`,
			&object.Error{Message: "wrong number of arguments. got=2, want=1"},
		},
		{
			`last(1)`,
			&object.Error{Message: "argument to `last` should be ARRAY, got INTEGER"},
		},
		{
			`last("one", "two")`,
			&object.Error{Message: "wrong number of arguments. got=2, want=1"},
		},
		{
			`rest(1)`,
			&object.Error{Message: "argument to `rest` should be ARRAY, got INTEGER"},
		},
		{
			`rest("one", "two")`,
			&object.Error{Message: "wrong number of arguments. got=2, want=1"},
		},
		{
			`push(1, 2)`,
			&object.Error{Message: "argument to `push` should be ARRAY, got INTEGER"},
		},
		{
			`push("one")`,
			&object.Error{Message: "wrong number of arguments. got=1, want>1"},
		},
	})
}

func TestClosures(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{
			`
			let newClosure = fn(a) {
				fn() { a; };
			};
			let closure = newClosure(99);
			closure();
			`,
			99,
		},
		{
			`
			let newAdder = fn(a, b) {
				fn(c) { a + b + c };
			};
			let adder = newAdder(1, 2);
			adder(8);
			`,
			11,
		},
		{
			`
			let newAdder = fn(a, b) {
				let c = a + b;
				fn(d) {
					c + d;
				};
			};
			let adder = newAdder(1, 2);
			adder(8);
			`,
			11,
		},
		{
			`
			let newAdderOuter = fn(a, b) {
				let c = a + b;
				fn(d) {
					let e = c + d;
					fn(f) {
						e + f;
					};
				};
			};
			let newAdder = newAdderOuter(1, 2);
			let adder = newAdder(3);
			adder(8);
			`,
			14,
		},
		{
			`
			let a = 1;
			let newAdderOuter = fn(b) {
				fn(c) {
					fn(d) {
						a + b + c + d;
					};
				};
			};
			let newAdder = newAdderOuter(2);
			let adder = newAdder(3);
			adder(8);
			`,
			14,
		},
		{
			`
			let newClosure = fn(a, b) {
				let one = fn() { a }
				let two = fn() { b }
				fn() { one() + two() };
			};
			let closure = newClosure(9, 90);
			closure();
			`,
			99,
		},
	})
}

func TestRecursiveFunctions(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{
			`
			let countDown = fn(x) {
				if (x == 0) {
					return 0;
				} else {
					countDown(x - 1);
				}
			};
			countDown(1);
			`,
			0,
		},
		{
			`
			let countDown = fn(x) {
				if (x == 0) {
					return 0;
				} else {
					countDown(x - 1);
				}
			};
			let wrapper = fn() {
				countDown(1);
			}
			wrapper();
			`,
			0,
		},
		{
			`
			let wrapper = fn() {
				let countDown = fn(x) {
					if (x == 0) {
						return 0;
					} else {
						countDown(x - 1);
					}
				};
				countDown(1);
			}
			wrapper();
			`,
			0,
		},
	})
}

func TestRecursiveFibonacci(t *testing.T) {
	runVmTests(t, []vmTestCase{
		{
			`
			let fibonacci = fn(x) {
				if (x == 0) {
					return 0;
				} else {
					if (x == 1) {
						return 1;
					} else {
						fibonacci(x - 1) + fibonacci(x - 2);
					}
				}
			};
			fibonacci(15);
			`,
			610,
		},
	})
}
