package scheme

import (
  "errors"
  "fmt"
)

type Scope struct {
  data map[*Symbol]Sexpr
  parent *Scope
}
func (s *Scope) Lookup(sym *Symbol) (Sexpr, error) {
  v, ok := s.data[sym]
  if ok {
    return v, nil
  }
  if s.parent != nil {
    return s.parent.Lookup(sym)
  }
  return nil, errors.New(fmt.Sprintf("undefined symbol %s", sym))
}

func (s *Scope) isDefinedHere(sym *Symbol) bool {
  _, ok := s.data[sym]
  return ok
}

func (s *Scope) Define(sym *Symbol, val Sexpr) {
    s.data[sym] = val
}

func (s *Scope) DefineHigh(sym *Symbol, val Sexpr) {
  if s.parent == nil || s.isDefinedHere(sym) {
    s.Define(sym, val)
  } else {
    s.parent.DefineHigh(sym, val)
  }
}

func NewScope(parent *Scope) *Scope {
  s := new(Scope)
  s.data = make(map[*Symbol]Sexpr)
  s.parent = parent
  return s
}

func NewGlobalWithData(data map[*Symbol]Sexpr) *Scope {
  s := NewScope(nil)
  s.data = data
  return s
}

func (s *Scope) GetOuterScope() *Scope {
  if s.parent == nil {
    return s
  }
  return s.parent.GetOuterScope()
}
