package scheme

import (
  "strconv"
  "fmt"
  "errors"
)

var (
  Nil = &nilPrim{}
)

type Sexpr interface {
  String() string
}

var symbolCache map[string]*Symbol = make(map[string]*Symbol)
type Symbol struct {
  str string
}
func (s *Symbol) String() string {
  return s.str
}
func SymbolFromString(str string) *Symbol {
  _, ok := symbolCache[str]
  if !ok {
    symbolCache[str] = &Symbol{str}
  }
  return symbolCache[str]
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
func NewScope(parent *Scope) *Scope {
  s := new(Scope)
  s.data = make(map[*Symbol]Sexpr)
  s.parent = parent
  return s
}

func NewGlobalScope() *Scope {
  s := NewScope(nil)
  s.data = GlobalData()
  return s
}

type proc func(*Scope, []Sexpr) (Sexpr, error)
type Function struct {
  p proc
  name string
}
func (f *Function) String() string {
  return f.name
}
func (f *Function) Procedure() proc {
  return f.p
}
type Primitive struct {
  p proc
  name string
}
func (p *Primitive) Procedure() proc {
  return p.p
}
func (p *Primitive) String() string {
  return p.name
}

