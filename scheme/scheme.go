package scheme

import (
  "strconv"
  "fmt"
  "errors"
)

var (
  Nil = &nilPrim{}
  True = &Boolean{true}
  False = &Boolean{false}
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
  Val int64
}
func (n *Number) String() string {
  return strconv.FormatInt(n.Val, 10)
}
func NumberFromInt(n int64) *Number {
  return &Number{n}
}

type Boolean struct {
  Val bool
}
func (b *Boolean) String() string {
  if b.Val {
    return "#t"
  } else {
    return "#f"
  }
  return "#t"
}
func BooleanFromBool(b bool) *Boolean {
  if b {
    return True
  }
  return False
}
func BooleanFromString(str string) *Boolean {
  if str == "#f" {
    return False
  }
  return True
}

type String struct {
  str string
}
func (s *String) String() string {
  return s.str
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
func (n *nilPrim) String() string { return "()" }

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

type Proc func(*Scope, []Sexpr) (Sexpr, error)
type Function struct {
  p Proc
  name string
}
func (f *Function) String() string {
  return f.name
}
func (f *Function) Procedure() Proc {
  return f.p
}
func CreateFunction(p Proc, name string) *Function {
  return &Function{p, name}
}
type Primitive struct {
  p Proc
  name string
}
func (p *Primitive) Procedure() Proc {
  return p.p
}
func CreatePrim(p Proc, name string) *Primitive {
  return &Primitive{p, name}
}
func (p *Primitive) String() string {
  return p.name
}

