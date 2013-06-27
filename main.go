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
  global :=  scheme.NewGlobalScope()
  fmt.Printf("%v\n", global)
  for {
    fmt.Printf("->")
    expr, _ := parser.Parse(c)
    e, err := eval.Eval(global, expr)
    if err != nil {
      fmt.Printf("%s\n", err)
    } else {
      fmt.Printf("=>%s\n", e)
    }
  }
}
