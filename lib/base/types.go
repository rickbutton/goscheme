package base

import (
  "github.com/rickbutton/goscheme/scheme"
)

func booleanCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("boolean? requires exactly 1 argument")
  }
  _, ok := args[0].(*scheme.Boolean)
  if ok {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}

func pairCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("pair? requires exactly 1 argument")
  }
  _, ok := args[0].(*scheme.Cons)
  if ok {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}

func symbolCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("symbol? requires exactly 1 argument")
  }
  _, ok := args[0].(*scheme.Symbol)
  if ok {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}

func numberCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("number? requires exactly 1 argument")
  }
  _, ok := args[0].(*scheme.Number)
  if ok {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}

func charCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("char? requires exactly 1 argument")
  }
  _, ok := args[0].(*scheme.Char)
  if ok {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}

func stringCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("string? requires exactly 1 argument")
  }
  _, ok := args[0].(*scheme.String)
  if ok {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}

func procCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("proc? requires exactly 1 argument")
  }
  _, okF := args[0].(*scheme.Function)
  _, okP := args[0].(*scheme.Primitive)
  if okF || okP {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}

func nullCheck(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("null? requires exactly 1 argument")
  }
  _, ok := args[0].(*scheme.NilPrim)
  if ok {
    return scheme.True, nil
  } else {
    return scheme.False, nil
  }
  return nil, nil
}
