package scheme

import (
  "errors"
  "fmt"
)

var (
  Nil = &NilPrim{}
  True = &Boolean{true}
  False = &Boolean{false}
)

type NilPrim struct {}
func (n *NilPrim) String() string { return "()" }

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

func ProcError(str string) error {
  return errors.New(fmt.Sprintf("Procedure error: %s", str))
}

func Flatten(s Sexpr) ([]Sexpr, error) {
	_, ok := s.(*Cons)
  ss := make([]Sexpr, 0)
	for ok {
    ss = append(ss, s.(*Cons).Car)
		s = s.(*Cons).Cdr
		_, ok = s.(*Cons)
	}
	if s != Nil {
    return nil, errors.New("list is not flat")
	}
	return ss, nil
}

func Unflatten(s []Sexpr) Sexpr {
  var c Sexpr = Nil
  for i := len(s) - 1; i >= 0; i-- {
    c = &Cons{s[i], c}
  }
  return c
}
