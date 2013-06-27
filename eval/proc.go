package eval

import (
  "github.com/rickbutton/goscheme/scheme"
)

var data map[*scheme.Symbol]scheme.Sexpr = nil

func GlobalData() map[*scheme.Symbol]scheme.Sexpr {
  if data == nil {
    data = setupData()
  }
  return data
}

func setupData() map[*scheme.Symbol]scheme.Sexpr {
  d := make(map[*scheme.Symbol]scheme.Sexpr)

  addFunction(d, "+", plus)
  addFunction(d, "*", mul)
  addFunction(d, "-", sub)
  addFunction(d, "/", div)
  
  addPrim(d, "lambda", lambda)
  addPrim(d, "let", let)
  addPrim(d, "define", define)
  addPrim(d, "quote", quote)

  return d
}

func addFunction(d map[*scheme.Symbol]scheme.Sexpr, name string, p scheme.Proc) {
  d[scheme.SymbolFromString(name)] = scheme.CreateFunction(p, name)
}

func addPrim(d map[*scheme.Symbol]scheme.Sexpr, name string, p scheme.Proc) {
  d[scheme.SymbolFromString(name)] = scheme.CreatePrim(p, name)
}
