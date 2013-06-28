package base

import (
  "fmt"
  "strings"
  "bufio"
  "os"
  "github.com/rickbutton/goscheme/scheme"
  "github.com/rickbutton/goscheme/lexer"
  "github.com/rickbutton/goscheme/parser"
  "github.com/rickbutton/goscheme/eval"
)

func read(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 0 {
    return nil, scheme.ProcError("read must have no arguments")
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

func primEval(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("eval requires exactly one argument")
  }
  return eval.Eval(s, args[0])
}

func print(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("print requires exactly one argument")
  }
  fmt.Printf("%s\n", args[0])
  return scheme.Nil, nil
}

func display(s *scheme.Scope, args []scheme.Sexpr) (scheme.Sexpr, error) {
  if len(args) != 1 {
    return nil, scheme.ProcError("display requires exactly one argument")
  }
  fmt.Printf("%s", args[0])
  return scheme.Nil, nil
}
