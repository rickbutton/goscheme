package scheme

import (
  "strconv"
)

var (
  Nil = &nilPrim{}
)

type Sexpr interface {
  String() string
}

type Symbol struct {
  str string
}
func (s *Symbol) String() string {
  return s.str
}
func SymbolFromString(str string) *Symbol {
  return &Symbol{str}
}

type Number struct {
  n int64
}
func (n *Number) String() string {
  return strconv.FormatInt(n.n, 10)
}
func NumberFromInt(n int64) *Number {
  return &Number{n}
}

type String struct {
  str string
}
func (s *String) String() string {
  return "\"" + s.str + "\""
}
func StringFromString(str string) *String {
  return &String{str}
}

type Cons struct {
  Car, Cdr Sexpr
}
func (c *Cons) String() string {
  return "(" + c.Car.String() + " . " + c.Cdr.String() + ")"
}

type nilPrim struct {}
func (n *nilPrim) String() string { return "nil" }

type Scope struct {
  data map[*Symbol]Sexpr
  parent *Scope
}
func (s *Scope) Lookup(sym *Symbol) Sexpr {
  v, ok := s.data[sym]
  if ok {
    return v
  }
  if s.parent != nil {
    return s.parent.Lookup(sym)
  }
  panic("undefined symbol")
}
func NewScope(parent *Scope) *Scope {
  s := new(Scope)
  s.data = make(map[*Symbol]Sexpr)
  s.parent = parent
  return s
}

type Function func(*Scope, []Sexpr) Sexpr
func (f *Function) String() string {
  return "func"
}
type Primitive func(*Scope, []Sexpr) Sexpr
func (p *Primitive) String() string {
  return "primitive"
}
