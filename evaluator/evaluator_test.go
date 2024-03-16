package evaluator

import (
	"monkey/object"
	parser "monkey/parser_antlr4"
	"testing"
)

func TestIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		{"-5", -5},
		{"-10", -10},
		{"5 + 5 + 5 + 5 - 10", 10},
		{"2 * 2 * 2 * 2 * 2", 32},
		{"-50 + 100 + -50", 0},
		{"5 * 2 + 10", 20},
		{"5 + 2 * 10", 25},
		{"50 / 2 * 2 + 10", 60},
		{"2 * (5 + 10)", 30},
		{"3 * 3 * 3 + 10", 37},
		{"3 * 3 * (3 + 10)", 117},
		{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	// l := lexer.New(input)
	p := parser.New(input)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object should be Integer. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("value of object should be %d. got=%d", expected, result.Value)
		return false
	}

	return true
}

func TestStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`"hello world"`, "hello world"},
		{`"hello" + " " + "world"`, "hello world"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testStringObject(t, evaluated, tt.expected)
	}
}

func testStringObject(t *testing.T, obj object.Object, expected string) bool {
	result, ok := obj.(*object.String)
	if !ok {
		t.Errorf("object should be String. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("value of object should be %q. got=%q", expected, result.Value)
		return false
	}
	return true
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true", true},
		{"false", false},
		{"1 < 2", true},
		{"1 > 2", false},
		{"1 < 1", false},
		{"1 > 1", false},
		{"1 == 2", false},
		{"1 != 2", true},
		{"1 == 1", true},
		{"1 != 1", false},
		{"true == true", true},
		{"false == false", true},
		{"true == false", false},
		{"true != false", true},
		{"false != true", true},
		{"(1 < 2) == true", true},
		{"(1 < 2) == false", false},
		{"(1 > 2) == true", false},
		{"(1 > 2) == false", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object should be Boolean. got=%T (%+v)", obj, obj)
		return false
	}

	if result.Value != expected {
		t.Errorf("value of object should be %t. got=%t", expected, result.Value)
		return false
	}

	return true
}

func TestBangOperator(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"!true", false},
		{"!false", true},
		{"!5", false},
		{"!!true", true},
		{"!!false", false},
		{"!!5", true},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testBooleanObject(t, evaluated, tt.expected)
	}
}

func TestIfElseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"if (true) { 10 }", 10},
		{"if (false) { 10 }", nil},
		{"if (1) { 10 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 > 2) { 10 }", nil},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		integer, ok := tt.expected.(int)
		if ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. got=%T (%+v)", obj, obj)
		return true
	}
	return false
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"return 10;", 10},
		{"return 10; 9;", 10},
		{"return 2 * 5; 9;", 10},
		{"9; return 2 * 5; 9;", 10},
		{`
		if (10 > 1) {
			if (10 > 1) {
				return 10;
			}
			return 1;
		}
		`, 10},
		{`
		let add = fn(x, y) {
			return x + y;
		};
		add(1, 2);
		add(3, 4);
		`, 7},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func TestErrorHandling(t *testing.T) {
	tests := []struct {
		input           string
		expectedMessage string
	}{
		{
			"5 + true;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"5 + true; 5;",
			"type mismatch: INTEGER + BOOLEAN",
		},
		{
			"-true",
			"unknown operator: -BOOLEAN",
		},
		{
			"true + false;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"5; true + false; 5;",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"if (10 > 1) { true + false; }",
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			`
			if (10 > 1) {
				if (10 > 1) {
					return true + false;
				}

				return 1;
			}
			`,
			"unknown operator: BOOLEAN + BOOLEAN",
		},
		{
			"foobar",
			"identifier not found: foobar",
		},
		{
			`"hello" - "world"`,
			"unknown operator: STRING - STRING",
		},
		{
			`len(1)`,
			"argument to `len` not supported, got INTEGER",
		},
		{
			`len("one", "two")`,
			"wrong number of arguments. got=2, want=1",
		},
		{
			`[1, 2, 3][fn() {}]`,
			"index operator not supported: ARRAY[FUNCTION]",
		},
		{
			`1[0]`,
			"index operator not supported: INTEGER[INTEGER]",
		},
		{
			`first(1)`,
			"argument to `first` should be ARRAY, got INTEGER",
		},
		{
			`first("one", "two")`,
			"wrong number of arguments. got=2, want=1",
		},
		{
			`last(1)`,
			"argument to `last` should be ARRAY, got INTEGER",
		},
		{
			`last("one", "two")`,
			"wrong number of arguments. got=2, want=1",
		},
		{
			`rest(1)`,
			"argument to `rest` should be ARRAY, got INTEGER",
		},
		{
			`rest("one", "two")`,
			"wrong number of arguments. got=2, want=1",
		},
		{
			`push(1, 2)`,
			"argument to `push` should be ARRAY, got INTEGER",
		},
		{
			`push("one")`,
			"wrong number of arguments. got=1, want>1",
		},
		{
			`{{}: {}}`,
			"unusable as hash key: HASH",
		},
		{
			`{"name": "monkey"}[fn(x) { x }]`,
			`unusable as hash key: FUNCTION`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		if errObj, ok := evaluated.(*object.Error); !ok {
			t.Errorf("no error object returned. got=%T (%+v)", evaluated, evaluated)
			continue
		} else if errObj.Message != tt.expectedMessage {
			t.Errorf("wrong error message. expected=%q, got=%q", tt.expectedMessage, errObj.Message)
		}
	}
}

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let a = 5; a;", 5},
		{"let a = 5 * 5; a;", 25},
		{"let a = 5; let b = a; b;", 5},
		{"let a = 5; let b = a; let c = a + b + 5; c;", 15},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}

func TestFunctionObject(t *testing.T) {
	input := "fn(x) { x + 2; };"

	evaluated := testEval(input)
	fn, ok := evaluated.(*object.Function)
	if !ok {
		t.Fatalf("object is not Function. got=%T (%+v)", evaluated, evaluated)
	}

	if len(fn.Parameters) != 1 {
		t.Fatalf("function has wrong number of parameters. got=%d", len(fn.Parameters))
	}

	parameter := fn.Parameters[0].String()
	if parameter != "x" {
		t.Fatalf("parameter should be %q. got=%q", "x", parameter)
	}

	expectedBody := "(x + 2)"
	body := fn.Body.String()
	if body != expectedBody {
		t.Fatalf("body should be %q. got=%q", expectedBody, body)
	}
}

func TestFunctionApplication(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"let identity = fn(x) { x; }; identity(5);", 5},
		{"let identity = fn(x) { return x; }; identity(5);", 5},
		{"let double = fn(x) { x * 2; }; double(5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5, 5);", 10},
		{"let add = fn(x, y) { x + y; }; add(5 + 5, add(5, 5));", 20},
		{"fn(x) { x; }(5)", 5},
		{"fn(x) { x * 2; }(5)", 10},
	}

	for _, tt := range tests {
		testIntegerObject(t, testEval(tt.input), tt.expected)
	}
}
func TestClosures(t *testing.T) {
	input := `
	let newAdder = fn(x) {
		return fn(y) {
			return x + y;
		};
	};
	let addTwo = newAdder(2);
	addTwo(3)
	`
	testIntegerObject(t, testEval(input), 5)
}

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world!")`, 12},
		{`len([])`, 0},
		{`len([1, 2, 3])`, 3},
		{
			`rest([1, 2, 3, 4])`,
			`[2, 3, 4]`,
		},
		{
			`rest([4])`,
			`[]`,
		},
		{
			`rest([])`,
			`null`,
		},
		{
			`push([],1)`,
			`[1]`,
		},
		{
			`push([1], 2, 3, 4)`,
			`[1, 2, 3, 4]`,
		},
		{
			`push([1], [2, 3, 4])`,
			`[1, [2, 3, 4]]`,
		},
		{
			`let a = [1]; let b = push(a, 2); a`,
			`[1]`,
		},
		{
			`let a = [1]; let b = push(a, 2); b`,
			`[1, 2]`,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch v := tt.expected.(type) {
		case int:
			testIntegerObject(t, evaluated, int64(v))
		case string:
			inspect := evaluated.Inspect()
			if inspect != v {
				t.Errorf("object inspect should be %q. got=%q", v, inspect)
			}
		}
	}
}

func TestArraylExpression(t *testing.T) {
	input := `[1, 2 * 3, 4 + 5]`

	evaluated := testEval(input)

	result, ok := evaluated.(*object.Array)
	if !ok {
		t.Fatalf("object should be Array. got=%T (%+v)", evaluated, evaluated)
	}

	if len(result.Elements) != 3 {
		t.Fatalf("number of elements of array should be 3. got=%d", len(result.Elements))
	}

	testIntegerObject(t, result.Elements[0], 1)
	testIntegerObject(t, result.Elements[1], 6)
	testIntegerObject(t, result.Elements[2], 9)
}

func TestArrayIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			"[1, 2, 3][0]",
			1,
		},
		{
			"[1, 2, 3][1]",
			2,
		},
		{
			"[1, 2, 3][2]",
			3,
		},
		{
			"let i = 0; [1][i]",
			1,
		},
		{
			"[1, 2, 3][1 + 1]",
			3,
		},
		{
			"let arr = [1, 2, 3]; arr[2]",
			3,
		},
		{
			"let arr = [1, 2, 3]; arr[0] + arr[1] + arr[2]",
			6,
		},
		{
			"let arr = [1, 2, 3]; let i = arr[0]; arr[i]",
			2,
		},
		{
			"[1, 2, 3][3]",
			nil,
		},
		{
			"[1, 2, 3][-1]",
			nil,
		},
		{
			"first([1, 2, 3])",
			1,
		},
		{
			"first([])",
			nil,
		},
		{
			"last([1, 2, 3])",
			3,
		},
		{
			"last([])",
			nil,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if integer, ok := tt.expected.(int); ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}

func TestHashExpression(t *testing.T) {
	input := `
	let two = "two";
	{
		"one": 10 - 9,
		two: 1 + 1,
		"thr" + "ee": 6 / 2,
		4: 4,
		true: 5,
		false: 6
	}
	`

	evaluated := testEval(input)
	hash, ok := evaluated.(*object.Hash)
	if !ok {
		t.Fatalf("object is not Hash. got=%T (%+v)", evaluated, evaluated)
	}

	expected := map[object.HashKey]int64{
		(&object.String{Value: "one"}).HashKey():   1,
		(&object.String{Value: "two"}).HashKey():   2,
		(&object.String{Value: "three"}).HashKey(): 3,
		(&object.Integer{Value: 4}).HashKey():      4,
		TRUE.HashKey():                             5,
		FALSE.HashKey():                            6,
	}

	if len(hash.Pairs) != len(expected) {
		t.Fatalf("Hash has wrong number of parameters. got=%d", len(hash.Pairs))
	}

	for key, value := range expected {
		pair, ok := hash.Pairs[key]
		if !ok {
			t.Errorf("no pair for given key in Pairs")
		}

		testIntegerObject(t, pair.Value, value)
	}
}

func TestHashIndexExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{
			`{"foo": 5}["foo"]`,
			5,
		},
		{
			`{"foo": 5}["bar"]`,
			nil,
		},
		{
			`let key = "foo"; {"foo": 5}[key]`,
			5,
		},
		{
			`{5: 5}[5]`,
			5,
		},
		{
			`{true: 5}[true]`,
			5,
		},
		{
			`{false: 5}[false]`,
			5,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		if integer, ok := tt.expected.(int); ok {
			testIntegerObject(t, evaluated, int64(integer))
		} else {
			testNullObject(t, evaluated)
		}
	}
}
