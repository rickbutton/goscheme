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
  if len(os.Args) != 2 {
    usage()
  }

  fileName := os.Args[1]
  fi, err := os.Open(fileName)
  if err != nil {
    fmt.Printf("Error: %s\n", err.Error())
    os.Exit(2)
  }
  r := bufio.NewReader(fi)
  _, c := lexer.Lex(r)
  data := eval.GlobalData()
  global := scheme.NewGlobalWithData(data)
  expr, _ := parser.Parse(c)
  _, err = eval.Eval(global, expr)
  if err != nil {
    fmt.Printf("%s\n", err.Error())
  }
}

func usage() {
  fmt.Printf("usage: %s [file.scm]\n", os.Args[0])
  os.Exit(2)
}

func repl() {
  in := os.Stdin
  r := bufio.NewReader(in)
  _, c := lexer.Lex(r)
  data := eval.GlobalData()
  global :=  scheme.NewGlobalWithData(data)
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
