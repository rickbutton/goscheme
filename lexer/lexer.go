package lexer

import (
  "io"
  "fmt"
  "strings"
  "strconv"
  "unicode"
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
  escape *bytes.Buffer
  done bool
}
func (l *Lexer) Done() bool {
  return l.done
}
type lexState func (l *Lexer) (lexState, error)

func newLexer(r io.RuneScanner) *Lexer {
  return &Lexer{r, make(chan Token), new(bytes.Buffer), new(bytes.Buffer), false}
}
func (l *Lexer) run() {
  for state := readyState; state != nil; {

    newState, err := state(l)
    if err == io.EOF {
      close(l.tokens)
      l.done = true
      return
    } else if err != nil {
      l.tokens <- Token{err.Error(), ErrorToken}
      close(l.tokens)
      l.done = true
      return
    }
    state = newState

  }
  close(l.tokens)
  l.done = true
}

type Token struct {
  Val string
  T TokenType
}

type TokenType int
const (
  SymbolToken TokenType = iota
  ParensToken
  StringToken
  BooleanToken
  NumberToken
  ErrorToken
)

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
    l.tokens <- Token{string(ch), ParensToken}
  } else if strings.ContainsRune(WS, ch) {
    //ignore whitespace
  } else if ch == ';' {
    l.tmp.WriteRune(ch)
    return commentState, nil
  } else if ch == '"' {
    return stringState, nil
  } else if ch == PROTECT {
    l.tokens <- Token{string(ch), SymbolToken}
  } else {
    l.tmp.WriteRune(ch)
    return readingState, nil
  }
  return readyState, nil
}

func readingState(l *Lexer) (lexState, error) {
  ch, err := l.nextRune()
  if err == io.EOF {
    tok := Token{l.tmp.String(), SymbolToken}
    l.tmp.Reset()
    l.r.UnreadRune()
    l.tokens <- tok
    return nil, nil
  } else if err != nil {
    return nil, err
  }
  if strings.ContainsRune(SPLIT, ch) {
    //current token ened
    tok := Token{l.tmp.String(), SymbolToken}
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
      tok := Token{"\"" + l.tmp.String() + "\"", StringToken}
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
  case 'a':
    l.tmp.WriteRune('\a')
  case 'b':
    l.tmp.WriteRune('\b')
  case 't':
    l.tmp.WriteRune('\t')
  case 'n':
    l.tmp.WriteRune('\n')
  case 'v':
    l.tmp.WriteRune('\v')
  case 'f':
    l.tmp.WriteRune('\f')
  case 'r':
    l.tmp.WriteRune('\r')
  case '"':
    l.tmp.WriteRune('"')
  case '\\':
    l.tmp.WriteRune('\\')
  case 'x':
    return hexEscapeState, nil
  default:
    return nil, lexError(fmt.Sprintf("invalid escape %c", ch))
  }
  return stringState, nil
}

func hexEscapeState(l *Lexer) (lexState, error) {
  ch, err := l.nextRune()
  if err != nil {
    return nil, err
  }
  if unicode.IsDigit(ch) {
    l.escape.WriteRune(ch)
  } else if ch == ';' {
    s := l.escape.String()
    n, _ := strconv.ParseInt(s, 16, 64)
    l.tmp.WriteRune(rune(n))
    return stringState, nil
  } else if !unicode.IsDigit(ch) {
    return nil, lexError("invalid hex escape")
  }
  return hexEscapeState, nil
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
