package scheme

import (
  "fmt"
  "errors"
)

func IsNumber(s Sexpr) bool {
  _, ok := s.(*Number)
  return ok
}

func procError(str string) error {
  return errors.New(fmt.Sprintf("Procedure error: %s", str))
}
