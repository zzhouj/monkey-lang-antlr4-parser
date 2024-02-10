package evaluator

import (
	"fmt"
	"monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len":   {Fn: builtinLen},
	"first": {Fn: builtinFirst},
	"last":  {Fn: builtinLast},
	"rest":  {Fn: builtinRest},
	"push":  {Fn: builtinPush},
	"puts":  {Fn: builtinPuts},
}

func builtinLen(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.String:
		return &object.Integer{Value: int64(len(arg.Value))}

	case *object.Array:
		return &object.Integer{Value: int64(len(arg.Elements))}

	default:
		return newError("argument to `len` not supported, got %s", args[0].Type())
	}
}

func builtinFirst(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		if len(arg.Elements) > 0 {
			return arg.Elements[0]
		} else {
			return NULL
		}

	default:
		return newError("argument to `first` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinLast(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		length := len(arg.Elements)
		if length > 0 {
			return arg.Elements[length-1]
		} else {
			return NULL
		}

	default:
		return newError("argument to `last` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinRest(args ...object.Object) object.Object {
	if len(args) != 1 {
		return newError("wrong number of arguments. got=%d, want=1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		length := len(arg.Elements)
		if length > 0 {
			newElements := make([]object.Object, length-1)
			copy(newElements, arg.Elements[1:length])
			return &object.Array{Elements: newElements}
		} else {
			return NULL
		}

	default:
		return newError("argument to `rest` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinPush(args ...object.Object) object.Object {
	if len(args) < 2 {
		return newError("wrong number of arguments. got=%d, want>1", len(args))
	}

	switch arg := args[0].(type) {
	case *object.Array:
		length := len(arg.Elements)
		pushLen := len(args) - 1
		newElements := make([]object.Object, length+pushLen)
		copy(newElements, arg.Elements)
		copy(newElements[length:], args[1:])
		return &object.Array{Elements: newElements}

	default:
		return newError("argument to `push` should be ARRAY, got %s", args[0].Type())
	}
}

func builtinPuts(args ...object.Object) object.Object {
	for _, arg := range args {
		fmt.Println(arg.Inspect())
	}

	return NULL
}
