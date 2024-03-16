package main

import (
	"flag"
	"fmt"
	"monkey/compiler"
	"monkey/evaluator"
	"monkey/object"
	parser "monkey/parser_antlr4"
	"monkey/vm"
	"time"
)

var engine = flag.String("engine", "vm", "use 'vm' or 'eval'")

var input = `
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
fibonacci(35);
`

func main() {
	flag.Parse()

	var duration time.Duration
	var result object.Object

	// l := lexer.New(input)
	p := parser.New(input)
	program := p.ParseProgram()

	if len(p.Errors()) > 0 {
		for _, err := range p.Errors() {
			fmt.Printf("parser error: %s", err)
		}
		return
	}

	switch *engine {
	case "vm":
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Printf("compiler error: %s", err)
			return
		}

		machine := vm.New(comp.ByteCode())

		start := time.Now()

		err = machine.Run()
		if err != nil {
			fmt.Printf("vm error: %s", err)
			return
		}

		duration = time.Since(start)
		result = machine.LastPopped()

	case "eval":
		env := object.NewEnvironment()

		start := time.Now()
		result = evaluator.Eval(program, env)
		duration = time.Since(start)

	default:
		fmt.Printf("unknown engine: %q", *engine)
		return
	}

	fmt.Printf("engine=%s, result=%s, duration=%s\n", *engine, result.Inspect(), duration)
}
