package parser

import (
  "strconv"
  "github.com/rickbutton/goscheme/lexer"
  "github.com/rickbutton/goscheme/scheme"
)

const (
  _LPAREN = lexer.Token("(")
  _RPAREN = lexer.Token(")")
  _PROTECT = lexer.Token("'")
)

func Parse(ch chan lexer.Token) (scheme.Sexpr, error) {
  return parse(ch)
}

func parse(ch chan lexer.Token) (scheme.Sexpr, error) {
  tok := <-ch
  return parseNext(ch, tok), nil
}

func parseNext(ch chan lexer.Token, tok lexer.Token) scheme.Sexpr {
  switch tok {
  case _LPAREN:
    return parseCons(ch)
  case _RPAREN:
    panic("Unmatched )")
  case _PROTECT:
    s, err := parse(ch)
    if err != nil {
      panic(err)
    }
    return &scheme.Cons{scheme.SymbolFromString("quote"), &scheme.Cons{s, nil}}
  }
  return parseAtom(tok)
}

func parseCons(ch chan lexer.Token) scheme.Sexpr {
  tok := <-ch
  if tok == _RPAREN {
    return scheme.Nil
  }
  if tok == lexer.Token(".") {
    tok := <-ch
    ret := parseNext(ch, tok)
    tok = <-ch
    if tok != _RPAREN {
      panic("Expected )")
    }
    return ret
  }
  car := parseNext(ch, tok)
  cdr := parseCons(ch)
  return &scheme.Cons{car, cdr}
}

func parseAtom(tok lexer.Token) scheme.Sexpr {
  var e scheme.Sexpr = scheme.SymbolFromString(string(tok))
  if tok[0] == '"' {
    e = scheme.StringFromString(string(tok[1 : len(tok) - 1]))
  }

  n, err := strconv.ParseInt(string(tok), 10, 64)
  if err == nil {
    e = scheme.NumberFromInt(n)
  }
  return e
}
