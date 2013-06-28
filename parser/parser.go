package parser

import (
  "strconv"
  "fmt"
  "errors"
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
  return parseNext(ch, tok)
}

func parseNext(ch chan lexer.Token, tok lexer.Token) (scheme.Sexpr, error) {
  switch tok {
  case _LPAREN:
    return parseCons(ch)
  case _RPAREN:
    return nil, parseError("unmatched )")
  case _PROTECT:
    s, err := parse(ch)
    if err != nil {
      return nil, err
    }
    return &scheme.Cons{scheme.SymbolFromString("quote"), &scheme.Cons{s, scheme.Nil}}, nil
  }
  return parseAtom(tok), nil
}

func parseCons(ch chan lexer.Token) (scheme.Sexpr, error) {
  tok := <-ch
  if tok == _RPAREN {
    return scheme.Nil, nil
  }
  if tok == lexer.Token(".") {
    tok := <-ch
    ret, err := parseNext(ch, tok)
    if err != nil {
      return nil, err
    }
    tok = <-ch
    if tok != _RPAREN {
      return nil, parseError("expected (")
    }
    return ret, nil
  }
  car, err := parseNext(ch, tok)
  if err != nil {
    return nil ,err
  }
  cdr, err := parseCons(ch)
  return &scheme.Cons{car, cdr}, err
}

func parseAtom(tok lexer.Token) scheme.Sexpr {
  var e scheme.Sexpr = scheme.SymbolFromString(string(tok))
  if tok[0] == '"' {
    e = scheme.StringFromString(string(tok[1 : len(tok) - 1]))
  }
  if tok == "#t" || tok == "#f" {
    e = scheme.BooleanFromString(string(tok))
  }

  n, err := strconv.ParseInt(string(tok), 10, 64)
  if err == nil {
    e = scheme.NumberFromInt(n)
  }
  return e
}

func parseError(str string) error {
  return errors.New(fmt.Sprintf("Parse error: %s", str))
}
