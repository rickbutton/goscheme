package parser

import (
  "strconv"
  "fmt"
  "errors"
  "github.com/rickbutton/goscheme/lexer"
  "github.com/rickbutton/goscheme/scheme"
)

var (
  _LPAREN = lexer.Token{"(", lexer.ParensToken}
  _RPAREN = lexer.Token{")", lexer.ParensToken}
  _PROTECT = lexer.Token{"'", lexer.SymbolToken}
)

func Parse(ch chan lexer.Token) (scheme.Sexpr, error) {
  return parse(ch)
}

func parse(ch chan lexer.Token) (scheme.Sexpr, error) {
  tok := <-ch
  if (tok.T == lexer.ErrorToken) {
    return nil, errors.New(tok.Val)
  }
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
  if (tok.T == lexer.ErrorToken) {
    return nil, errors.New(tok.Val)
  }
  if tok == _RPAREN {
    return scheme.Nil, nil
  }
  if tok.Val == "." {
    tok := <-ch
    if (tok.T == lexer.ErrorToken) {
      return nil, errors.New(tok.Val)
    }
    ret, err := parseNext(ch, tok)
    if err != nil {
      return nil, err
    }
    tok = <-ch
    if (tok.T == lexer.ErrorToken) {
      return nil, errors.New(tok.Val)
    }
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
  if isString(tok) {
    return scheme.StringFromString(string(tok.Val[1 : len(tok.Val) - 1]))
  } else if isBoolean(tok) {
    return scheme.BooleanFromString(tok.Val)
  } else if isNumber(tok) {
    n, _ := strconv.ParseInt(tok.Val, 10, 64)
    return scheme.NumberFromInt(n)
  }
  return scheme.SymbolFromString(tok.Val)
}

func isNumber(tok lexer.Token) bool {
  _, err := strconv.ParseInt(tok.Val, 10, 64)
  return err == nil
}
func isString(tok lexer.Token) bool {
  return tok.Val[0] =='"' && tok.Val[len(tok.Val)-1] == '"'
}
func isBoolean(tok lexer.Token) bool {
  return tok.Val == "#t" || tok.Val == "#f"
}

func parseError(str string) error {
  return errors.New(fmt.Sprintf("Parse error: %s", str))
}
