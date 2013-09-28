package lib

import (
  "github.com/rickbutton/goscheme/scheme"
)

func cons(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, scheme.ProcError("cons requires exactly 2 arguments")
  }
  return &scheme.Cons{args[0], args[1]}, nil
}

func car(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("car requires exactly 1 argument")
  }
  if !scheme.IsCons(args[0]) {
    return nil, scheme.ProcError("car requires a cons arguments")
  }
  return args[0].(*scheme.Cons).Car, nil
}

func cdr(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("cdr requires exactly 1 argument")
  }
  if !scheme.IsCons(args[0]) {
    return nil, scheme.ProcError("cdr requires a cons argument")
  }
  return args[0].(*scheme.Cons).Cdr, nil
}

func setcar(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, scheme.ProcError("set-car! requires exactly 2 arguments")
  }
  pair := args[0]
  value := args[1]
  if !scheme.IsCons(pair) {
    return nil, scheme.ProcError("set-car! requires a cons argument")
  }
  pair.(*scheme.Cons).Car = value
  return scheme.Nil, nil
}

func setcdr(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 2 {
    return nil, scheme.ProcError("set-cdr! requires exactly 2 arguments")
  }
  pair := args[0]
  value := args[1]
  if !scheme.IsCons(pair) {
    return nil, scheme.ProcError("set-cdr! requires a cons argument")
  }
  pair.(*scheme.Cons).Cdr = value
  return scheme.Nil, nil
}
