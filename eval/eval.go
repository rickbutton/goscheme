package eval

import (
  "github.com/rickbutton/goscheme/scheme"
)

func Eval(sc *scheme.Scope, e scheme.Sexpr) scheme.Sexpr {
  //e = transform(sc, e)
  switch e := e.(type) {
  case *scheme.Cons:
    cons := e
    car := Eval(sc, cons.Car)
    if !isFunction(car) && !isPrim(car) {
      panic("Attempted application on non-function")
    }
    cdr := cons.Cdr
    args := flatten(cdr)
    if isPrim(cdr) {
      return (*(car.(*scheme.Primitive)))(sc, args)
    }
    f := *car.(*scheme.Function)
    for i, a := range args {
      args[i] = Eval(sc, a)
    }
    return f(sc, args)
  case *scheme.Symbol:
    return sc.Lookup(e)
  }
  return e
}

func flatten(s scheme.Sexpr) (ss []scheme.Sexpr) {
	_, ok := s.(*scheme.Cons)
	for ok {
		ss = append(ss, s.(*scheme.Cons).Car)
		s = s.(*scheme.Cons).Cdr
		_, ok = s.(*scheme.Cons)
	}
	if s != nil {
		panic("List isn't flat")
	}
	return
}

func isFunction(e scheme.Sexpr) bool {
  _, ok := e.(*scheme.Function)
  return ok
}

func isPrim(e scheme.Sexpr) bool {
  _, ok := e.(*scheme.Primitive)
  return ok
}
