package eval

import (
  "fmt"
  "errors"
  "github.com/rickbutton/goscheme/scheme"
)

func IsNumber(s scheme.Sexpr) bool {
  _, ok := s.(*scheme.Number)
  return ok
}

func procError(str string) error {
  return errors.New(fmt.Sprintf("Procedure error: %s", str))
}
