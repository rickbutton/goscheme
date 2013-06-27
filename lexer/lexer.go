package lexer

import (
  "io"
  "fmt"
  "strings"
  "errors"
  "bytes"
)

const (
  PARENS = "()"
  WS     = "\t\r\n "
  SPLIT  = PARENS + WS + ";"
  PROTECT= '\''
)

type Lexer struct {
  r io.RuneScanner
  tokens chan Token
  tmp *bytes.Buffer
}
type lexState func (l *Lexer) (lexState, error)

func newLexer(r io.RuneScanner) *Lexer {
  return &Lexer{r, make(chan Token), new(bytes.Buffer)}
}
func (l *Lexer) run() {
  for state := readyState; state != nil; {

    newState, _ := state(l)
    state = newState

  }
  close(l.tokens)
}

type Token string

func Lex(r io.RuneScanner) (*Lexer, chan Token) {
  l := newLexer(r)
  go l.run()
  return l, l.tokens
}

func (l *Lexer) nextRune() (rune, error) {
  ch, _, err := l.r.ReadRune()
  return ch, err
}

func readyState(l *Lexer) (lexState, error) {
  ch, err := l.nextRune()
  if err != nil {
    return nil, err
  }
  if strings.ContainsRune(PARENS, ch) {
    //Just emit the parens character
    l.tokens <- Token(ch)
  } else if strings.ContainsRune(WS, ch) {
    //ignore whitespace
  } else if ch == ';' {
    l.tmp.WriteRune(ch)
    return commentState, nil
  } else if ch == '"' {
    return stringState, nil
  } else if ch == PROTECT {
    l.tokens <- Token(ch)
  } else {
    l.tmp.WriteRune(ch)
    return readingState, nil
  }
  return readyState, nil
}

func readingState(l *Lexer) (lexState, error) {
  ch, err := l.nextRune()
  if err == io.EOF {
    tok := Token(l.tmp.String())
    l.tmp.Reset()
    l.r.UnreadRune()
    l.tokens <- tok
    return nil, nil
  } else if err != nil {
    return nil, err
  }
  if strings.ContainsRune(SPLIT, ch) {
    //current token ened
    tok := Token(l.tmp.String())
    l.tmp.Reset()
    l.r.UnreadRune()
    l.tokens <- tok
    return readyState, nil
  } else {
    l.tmp.WriteRune(ch)
  }
  return readingState, nil
}

func stringState(l *Lexer) (lexState, error) {
  ch, err := l.nextRune()
  if err != nil {
    return nil, err
  }
  if ch == '\\' {
    return escapeState, nil
  } else {
    if ch == '"' {
      tok := Token("\"" + l.tmp.String() + "\"")
      l.tmp.Reset()
      l.tokens <- tok
      return readyState, nil
    } else {
      l.tmp.WriteRune(ch)
    }
    return stringState, nil
  }
  return stringState, nil
}

func escapeState(l *Lexer) (lexState, error) {
  ch, err := l.nextRune()
  if err != nil {
    return nil, err
  }
  switch ch {
  case 'n':
    l.tmp.WriteRune('\n')
  case 't':
    l.tmp.WriteRune('\t')
  default:
    return nil, lexError(fmt.Sprintf("invalid escape %c", ch))
  }
  return stringState, nil
}

func commentState(l *Lexer) (lexState, error) {
  ch, err := l.nextRune()
  if err != nil {
    return nil, err
  }
  if ch == '\n' {
    return readyState, nil
  }
  return commentState, nil
}

func lexError(str string) error {
  return errors.New(fmt.Sprintf("Lexer error: %s", str))
}
