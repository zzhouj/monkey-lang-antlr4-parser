package object

import "fmt"

var Builtins = []struct {
	Name    string
	Builtin *Builtin
}{
	{"len", &Builtin{Fn: builtinLen}},
	{"first", &Builtin{Fn: builtinFirst}},
	{"last", &Builtin{Fn: builtinLast}},
	{"rest", &Builtin{Fn: builtinRest}},
	{"push", &Builtin{Fn: builtinPush}},
	{"puts", &Builtin{Fn: builtinPuts}},
}

func GetBuiltinByName(name string) *Builtin {
	for _, def := range Builtins {
		if def.Name == name {
			return def.Builtin
		}
	}
	return nil
}

func builtinLen(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *String:
		return &Integer{Value: int64(len(arg.Value))}

	case *Array:
		return &Integer{Value: int64(len(arg.Elements))}

	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func builtinFirst(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *Array:
		if len(arg.Elements) > 0 {
			return arg.Elements[0]
		} else {
			return nil
		}

	default:
		return newError("argument to `first` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinLast(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *Array:
		length := len(arg.Elements)
		if length > 0 {
			return arg.Elements[length-1]
		} else {
			return nil
		}

	default:
		return newError("argument to `last` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinRest(args ...Object) Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *Array:
		length := len(arg.Elements)
		if length > 0 {
			newElements := make([]Object, length-1)
			copy(newElements, arg.Elements[1:length])
			return &Array{Elements: newElements}
		} else {
			return nil
		}

	default:
		return newError("argument to `rest` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinPush(args ...Object) Object {
	if len(args) < 2 {
		return newError("wrong number of arguments. got=%d, want>1", len(args))
	}

	switch arg := args[0].(type) {
	case *Array:
		length := len(arg.Elements)
		pushLen := len(args) - 1
		newElements := make([]Object, length+pushLen)
		copy(newElements, arg.Elements)
		copy(newElements[length:], args[1:])
		return &Array{Elements: newElements}

	default:
		return newError("argument to `push` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinPuts(args ...Object) Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return nil
}

func newError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}
