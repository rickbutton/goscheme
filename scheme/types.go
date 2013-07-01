package scheme

import (
  "strconv"
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
func IsNumber(s Sexpr) bool {
  _, ok := s.(*Number)
  return ok
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

type Char struct {
  r rune
}
func (s *Char) String() string {
  return string(s.r)
}
func CharFromRune(r rune) *Char {
  return &Char{r}
}

type Cons struct {
  Car, Cdr Sexpr
}
func (c *Cons) String() string {
  return "(" + c.Car.String() + " . " + c.Cdr.String() + ")"
}

