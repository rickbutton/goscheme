package eval

import (
  "github.com/rickbutton/goscheme/scheme"
)

func plus(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) < 2 {
    return nil, procError("+ requires at least 2 arguments")
  }
  var sum int64 = 0
  for _, a := range args {
    if IsNumber(a) {
      sum += a.(*scheme.Number).Val
    } else {
      return nil, procError("+ requires all number arguments")
    }
  }
  return scheme.NumberFromInt(sum), nil
}

func mul(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) < 2 {
    return nil, procError("* requires at least 2 arguments")
  }
  var product int64 = 1
  for _, a := range args {
    if IsNumber(a) {
      product *= a.(*scheme.Number).Val
    } else {
      return nil, procError("* requires all number arguments")
    }
  }
  return scheme.NumberFromInt(product), nil
}

func sub(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, procError("- requires exactly 2 arguments")
  }
  if !IsNumber(args[0]) || !IsNumber(args[1]) {
    return nil, procError("- requires all number arguments")
  }
  left := args[0].(*scheme.Number).Val
  right := args[1].(*scheme.Number).Val
  return scheme.NumberFromInt(left - right), nil
}

func div(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, procError("/ requires exactly 2 arguments")
  }
  if !IsNumber(args[0]) || !IsNumber(args[1]) {
    return nil, procError("/ requires all number arguments")
  }
  left := args[0].(*scheme.Number).Val
  right := args[1].(*scheme.Number).Val
  return scheme.NumberFromInt(left / right), nil
}
