package compiler

type SymbolScope string

const (
	LocalScope   SymbolScope = "LOCAL"
	GlobalScope  SymbolScope = "GLOBAL"
	BuiltinScope SymbolScope = "BUILTIN"
	FreeScope    SymbolScope = "FREE"
)

type Symbol struct {
	Name  string
	Scope SymbolScope
	Index int
}

type SymbolTable struct {
	Outer *SymbolTable

	store          map[string]Symbol
	numDifinitions int

	FreeSymbols []Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		store:       map[string]Symbol{},
		FreeSymbols: []Symbol{},
	}
}

func NewEnclosedSymbolTable(outer *SymbolTable) *SymbolTable {
	st := NewSymbolTable()
	st.Outer = outer
	return st
}

func (st *SymbolTable) Define(name string) Symbol {
	symbol := Symbol{name, GlobalScope, st.numDifinitions}
	if st.Outer != nil {
		symbol.Scope = LocalScope
	}
	st.store[name] = symbol
	st.numDifinitions++
	return symbol
}

func (st *SymbolTable) DefineBuiltin(index int, name string) Symbol {
	symbol := Symbol{name, BuiltinScope, index}
	st.store[name] = symbol
	return symbol
}

func (st *SymbolTable) defineFree(orig Symbol) Symbol {
	st.FreeSymbols = append(st.FreeSymbols, orig)
	s := Symbol{orig.Name, FreeScope, len(st.FreeSymbols) - 1}
	st.store[s.Name] = s
	return s
}

func (st *SymbolTable) Resolve(name string) (Symbol, bool) {
	symbol, ok := st.store[name]
	if !ok && st.Outer != nil {
		symbol, ok = st.Outer.Resolve(name)
		if ok && (symbol.Scope != GlobalScope && symbol.Scope != BuiltinScope) {
			symbol = st.defineFree(symbol)
		}
	}
	return symbol, ok
}
