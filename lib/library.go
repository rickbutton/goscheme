package lib

import (
  "github.com/rickbutton/goscheme/scheme"
)

func Definitions() map[*scheme.Symbol]scheme.Sexpr {
  d := make(map[*scheme.Symbol]scheme.Sexpr)

  addFunction(d, "read", read)
  addFunction(d, "eval", primEval)
  addFunction(d, "print", print)
  addFunction(d, "display", display)

  addFunction(d, "boolean?", booleanCheck)
  addFunction(d, "pair?", pairCheck)
  addFunction(d, "symbol?", symbolCheck)
  addFunction(d, "number?", numberCheck)
  addFunction(d, "char?", charCheck)
  addFunction(d, "string?", stringCheck)
  addFunction(d, "procedure?", procCheck)
  addFunction(d, "null?", nullCheck)

  //Arithmetic
  addFunction(d, "+", plus)
  addFunction(d, "*", mul)
  addFunction(d, "-", sub)
  addFunction(d, "/", div)

  //Pairs
  addFunction(d, "cons", cons)
  addFunction(d, "car", car)
  addFunction(d, "cdr", cdr)
  addFunction(d, "set-car!", setcar)
  addFunction(d, "set-cdr!", setcdr)

  addPrim(d, "lambda", lambda)
  addPrim(d, "let", let)
  addPrim(d, "define", define)
  addPrim(d, "quote", quote)
  addPrim(d, "begin", begin)

  addPrim(d, "for", primFor)

  d[scheme.SymbolFromString("nil")] = scheme.Nil

  return d
}

func addFunction(d map[*scheme.Symbol]scheme.Sexpr, name string, p scheme.Proc) {
  d[scheme.SymbolFromString(name)] = scheme.CreateFunction(p, name)
}

func addPrim(d map[*scheme.Symbol]scheme.Sexpr, name string, p scheme.Proc) {
  d[scheme.SymbolFromString(name)] = scheme.CreatePrim(p, name)
}
