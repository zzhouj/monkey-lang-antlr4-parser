package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/compiler"
	"monkey/object"
	parser "monkey/parser_antlr4"
	"monkey/vm"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	constants := []object.Object{}
	symbolTable := compiler.NewSymbolTable()
	for i, builtin := range object.Builtins {
		symbolTable.DefineBuiltin(i, builtin.Name)
	}
	globals := make([]object.Object, vm.GlobalsSize)

	for {
		fmt.Fprint(out, PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		// l := lexer.New(line)
		p := parser.New(line)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		comp := compiler.NewWithState(constants, symbolTable)
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "compilation failed: %v\n", err)
			continue
		}

		bc := comp.ByteCode()
		constants = bc.Constants

		vm := vm.NewWithState(bc, globals)
		err = vm.Run()
		if err != nil {
			fmt.Fprintf(out, "executing bytecode failed: %v\n", err)
			continue
		}

		lastPopped := vm.LastPopped()
		if lastPopped != nil {
			io.WriteString(out, lastPopped.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
