package lib

import (
  "github.com/rickbutton/goscheme/scheme"
  "github.com/rickbutton/goscheme/eval"
)

func primFor(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, scheme.ProcError("for requires exactly two arguments")
  }
  cond := args[0]
  expr := args[1]
  var val scheme.Sexpr = scheme.Nil
  cv, err := eval.Eval(s, cond)
  if err != nil {
    return nil, err
  }
  for cv != scheme.Nil && cv != scheme.False {
    val, err = eval.Eval(s, expr)
    if err != nil {
      return nil, err
    }
    cv, err = eval.Eval(s, cond)
    if err != nil {
      return nil, err
    }
  }
  return val, nil
}

func lambda(s *scheme.Scope, ss []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(ss) < 2 {
    return nil, scheme.ProcError("lambda at least 2 arguments")
  }
  exprs := ss[1:]
  evalScopeParent := scheme.NewScope(s)
  lArgs := ss[0]
  f := func (callScope *scheme.Scope, ss []scheme.Sexpr) (scheme.Sexpr, error) {
    args := lArgs
    evalScope := scheme.NewScope(evalScopeParent)
    aC, ok :=args.(*scheme.Cons)
    for args != scheme.Nil {
      if len(ss) == 0 {
        return nil, scheme.ProcError("invalid number of arguments")
      }
      if !ok {
        val := scheme.Unflatten(ss)
        s, k := args.(*scheme.Symbol)
        if !k {
          return nil, scheme.ProcError("Invalid parameter specification")
        }
        evalScope.Define(s, val)
        goto done
      }
      arg := aC.Car
      val := ss[0]
      s, k := arg.(*scheme.Symbol)
      if !k {
        return nil, scheme.ProcError("invalid parameter specification")
      }
      evalScope.Define(s, val)

      ss = ss[1:]
      args = aC.Cdr
      aC, ok = args.(*scheme.Cons)
    }
    if len(ss) > 0 {
      return nil, scheme.ProcError("Invalid number of arguments")
    }
  done:
    for _, e := range exprs[0:len(exprs) - 1] {
      eval.Eval(evalScope, e)
    }
    return eval.Eval(evalScope, exprs[len(exprs) - 1])
  }
  return scheme.CreateFunction(f, "lamba"), nil
}

func let(s *scheme.Scope, ss []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(ss) < 1 {
    return nil, scheme.ProcError("let requires at least one argument")
  }
  evalScope := scheme.NewScope(s)
  bindings, err := scheme.Flatten(ss[0])
  if err != nil {
    return nil, err
  }
  for _, b := range bindings {
    bs, err := scheme.Flatten(b)
    if err != nil {
      return nil, err
    }
    if len(bs) != 2 {
      return nil, scheme.ProcError("invalid binding on let")
    }
    sym, ok := bs[0].(*scheme.Symbol)
    if !ok {
      return nil, scheme.ProcError("invalid binding on let")
    }
    val, err := eval.Eval(s, bs[1])
    if err != nil {
      return nil , err
    }
    evalScope.Define(sym, val)
  }
  prog := ss[1:]
  var last scheme.Sexpr = scheme.Nil
  for _, l := range prog {
    last, err = eval.Eval(evalScope, l)
    if err != nil {
      return nil, err
    }
  }
  return last, nil
}

func define(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, scheme.ProcError("define requires exactly 2 arguments")
  }
  idSym, ok := args[0].(*scheme.Symbol)
  if !ok {
    return nil, scheme.ProcError("invalid argument to define")
  }
  val, err := eval.Eval(s, args[1])
  if err != nil {
    return nil, err
  }
  s.DefineHigh(idSym, val)
  return scheme.Nil, nil
}

func quote(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("quote requires exactly one argument")
  }
  return args[0], nil
}

func begin(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  var last scheme.Sexpr = scheme.Nil
  var err error = nil
  for _, l := range args {
    last, err = eval.Eval(s, l)
    if err != nil {
      return nil, err
    }
  }
  return last, nil
}
