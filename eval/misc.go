package eval

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "github.com/rickbutton/goscheme/scheme"
  "github.com/rickbutton/goscheme/lexer"
  "github.com/rickbutton/goscheme/parser"
)

func read(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 0 {
    return nil, procError("read must have no arguments")
  }

  br := bufio.NewReader(os.Stdin)
  str, err := br.ReadString('\n')

  r := strings.NewReader(str)
  _, c := lexer.Lex(r)
  expr, err := parser.Parse(c)
  if err != nil {
    return nil, err
  }
  return expr, nil
}

func eval(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, procError("eval requires exactly one argument")
  }
  return Eval(s, args[0])
}

func print(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, procError("print requires exactly one argument")
  }
  fmt.Printf("%s\n", args[0])
  return scheme.Nil, nil
}

func display(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, procError("display requires exactly one argument")
  }
  fmt.Printf("%s", args[0])
  return scheme.Nil, nil
}
