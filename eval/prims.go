package eval

import (
  "github.com/rickbutton/goscheme/scheme"
)

func lambda(s *scheme.Scope, ss []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(ss) != 2 {
    return nil, procError("lambda requires exactly 2 arguments")
  }
  expr := ss[1]
  evalScopeParent := scheme.NewScope(s)
  lArgs := ss[0]
  f := func (callScope *scheme.Scope, ss []scheme.Sexpr) (scheme.Sexpr, error) {
    args := lArgs
    evalScope := scheme.NewScope(evalScopeParent)
    aC, ok :=args.(*scheme.Cons)
    for args != scheme.Nil {
      if len(ss) == 0 {
        panic("TEST")
        return nil, procError("invalid number of arguments")
      }
      if !ok {
        val := unflatten(ss)
        s, k := args.(*scheme.Symbol)
        if !k {
          return nil, procError("Invalid parameter specification")
        }
        evalScope.Define(s, val)
        goto done
      }
      arg := aC.Car
      val := ss[0]
      s, k := arg.(*scheme.Symbol)
      if !k {
        return nil, procError("invalid parameter specification")
      }
      evalScope.Define(s, val)

      ss = ss[1:]
      args = aC.Cdr
      aC, ok = args.(*scheme.Cons)
    }
    if len(ss) > 0 {
      return nil, procError("Invalid number of arguments")
    }
  done:
    return Eval(evalScope, expr)
  }
  return scheme.CreateFunction(f, "lamba"), nil
}

func let(s *scheme.Scope, ss []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(ss) < 1 {
    return nil, procError("let requires at least one argument")
  }
  evalScope := scheme.NewScope(s)
  bindings, err := flatten(ss[0])
  if err != nil {
    return nil, err
  }
  for _, b := range bindings {
    bs, err := flatten(b)
    if err != nil {
      return nil, err
    }
    if len(bs) != 2 {
      return nil, procError("invalid binding on let")
    }
    sym, ok := bs[0].(*scheme.Symbol)
    if !ok {
      return nil, procError("invalid binding on let")
    }
    val, err := Eval(s, bs[1])
    if err != nil {
      return nil , err
    }
    evalScope.Define(sym, val)
  }
  prog := ss[1:]
  var last scheme.Sexpr = scheme.Nil
  for _, l := range prog {
    last, err = Eval(evalScope, l)
    if err != nil {
      return nil, err
    }
  }
  return last, nil
}

func define(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, procError("define requires exactly 2 arguments")
  }
  idSym, ok := args[0].(*scheme.Symbol)
  if !ok {
    return nil, procError("invalid argument to define")
  }
  val, err := Eval(s, args[1])
  if err != nil {
    return nil, err
  }
  s.DefineHigh(idSym, val)
  return scheme.Nil, nil
}

func quote(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, procError("quote requires exactly one argument")
  }
  return args[0], nil
}
