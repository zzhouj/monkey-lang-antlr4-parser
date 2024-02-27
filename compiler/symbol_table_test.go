package compiler

import "testing"

func TestDefine(t *testing.T) {
	tests := map[string]Symbol{
		"a": {"a", GlobalScope, 0},
		"b": {"b", GlobalScope, 1},
	}

	st := NewSymbolTable()

	for _, name := range []string{"a", "b"} {
		expected := tests[name]
		actual := st.Define(name)
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}

	for name, expected := range tests {
		actual, ok := st.Resolve(name)
		if !ok {
			t.Errorf("name %q not resolvable", name)
			continue
		}
		if actual != expected {
			t.Errorf("want=%+v, got=%+v", expected, actual)
		}
	}
}
