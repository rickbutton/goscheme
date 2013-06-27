package scheme

var data map[*Symbol]Sexpr = nil

func GlobalData() map[*Symbol]Sexpr {
  if data == nil {
    data = setupData()
  }
  return data
}

func setupData() map[*Symbol]Sexpr {
  d := make(map[*Symbol]Sexpr)

  addFunction(d, "+", plus)
  addFunction(d, "*", mul)
  addFunction(d, "-", sub)
  addFunction(d, "/", div)

  return d
}

func addFunction(d map[*Symbol]Sexpr, name string, p proc) {
  d[SymbolFromString(name)] = &Function{p, name}
}
