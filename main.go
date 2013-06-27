package main

import (
  "fmt"
  "os"
  "bufio"
  "github.com/rickbutton/goscheme/scheme"
  "github.com/rickbutton/goscheme/lexer"
  "github.com/rickbutton/goscheme/parser"
  "github.com/rickbutton/goscheme/eval"
)

func main() {
  repl()
}

func repl() {
  in := os.Stdin
  r := bufio.NewReader(in)
  _, c := lexer.Lex(r)
  global :=  scheme.NewScope(nil)
  for {
    fmt.Printf("->")
    expr, _ := parser.Parse(c)
    e := eval.Eval(global, expr)
    fmt.Printf("=>%s\n", e)
  }
}
