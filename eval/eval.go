package eval

import (
  "errors"
  "fmt"
  "github.com/rickbutton/goscheme/scheme"
)

func Eval(sc *scheme.Scope, e scheme.Sexpr) (scheme.Sexpr, error) {
  //e = transform(sc, e)
  switch e := e.(type) {
  case *scheme.Cons:
    cons := e
    car, err := Eval(sc, cons.Car)
    if err != nil {
      return nil, err
    }
    if !isFunction(car) && !isPrim(car) {
      return nil, evalError("Attempted application on non-function")
    }
    cdr := cons.Cdr
    args, err := flatten(cdr)
    if err != nil {
      return nil, err
    }
    if isPrim(cdr) {
      return car.(*scheme.Primitive).Procedure()(sc, args)
    }
    f := car.(*scheme.Function).Procedure()
    for i, a := range args {
      args[i], err = Eval(sc, a)
      if err != nil {
        return nil, err
      }
    }
    return f(sc, args)
  case *scheme.Symbol:
    return sc.Lookup(e)
  }
  return e, nil
}

func flatten(s scheme.Sexpr) ([]scheme.Sexpr, error) {
	_, ok := s.(*scheme.Cons)
  ss := make([]scheme.Sexpr, 0)
	for ok {
    ss = append(ss, s.(*scheme.Cons).Car)
		s = s.(*scheme.Cons).Cdr
		_, ok = s.(*scheme.Cons)
	}
	if s != scheme.Nil {
    return nil, evalError("list isn't flat")
	}
	return ss, nil
}

func isFunction(e scheme.Sexpr) bool {
  _, ok := e.(*scheme.Function)
  return ok
}

func isPrim(e scheme.Sexpr) bool {
  _, ok := e.(*scheme.Primitive)
  return ok
}

func evalError(str string) error {
  return errors.New(fmt.Sprintf("Eval error: %s", str))
}
