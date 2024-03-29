package compiler

import "testing"

func TestDefine(t *testing.T) {
	tests := map[string]Symbol{
		"a": {"a", GlobalScope, 0},
		"b": {"b", GlobalScope, 1},
		"c": {"c", LocalScope, 0},
		"d": {"d", LocalScope, 1},
		"e": {"e", LocalScope, 0},
		"f": {"f", LocalScope, 1},
		"g": {"g", BuiltinScope, 0},
		"h": {"h", BuiltinScope, 1},
		"i": {"i", FunctionScope, 0},
		"j": {"j", FunctionScope, 0},
	}

	global := NewSymbolTable()
	local1 := NewEnclosedSymbolTable(global)
	local2 := NewEnclosedSymbolTable(local1)

	for _, name := range []string{"a", "b"} {
		expected := tests[name]
		actual := global.Define(name)
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}

	for i, name := range []string{"g", "h"} {
		expected := tests[name]
		actual := global.DefineBuiltin(i, name)
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}

	for _, name := range []string{"c", "d"} {
		expected := tests[name]
		actual := local1.Define(name)
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}

	for _, name := range []string{"e", "f"} {
		expected := tests[name]
		actual := local2.Define(name)
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}

	for _, name := range []string{"i", "j"} {
		expected := tests[name]
		actual := global.DefineFunction(name)
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}

	tests["c"] = Symbol{"c", FreeScope, 0}
	tests["d"] = Symbol{"d", FreeScope, 1}
	tests["i"] = Symbol{"i", FreeScope, 2}
	tests["j"] = Symbol{"j", FreeScope, 3}

	for _, name := range []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"} {
		expected := tests[name]
		actual, ok := local2.Resolve(name)
		if !ok {
			t.Errorf("name %q not resolvable", name)
			continue
		}
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}
}
