package eval

import (
  "github.com/rickbutton/goscheme/scheme"
)

func IsNumber(s scheme.Sexpr) bool {
  _, ok := s.(*scheme.Number)
  return ok
}
