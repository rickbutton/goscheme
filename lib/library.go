package lib

import (
  "strings"
  "github.com/rickbutton/goscheme/scheme"
  "github.com/rickbutton/goscheme/lib/base"
)

var (
  goLibs = map[string]*Library {
    "rnrs/base":&Library{"rnrs/base", base.Definitions()},
  }
)

type Library struct {
  name string
  data map[*scheme.Symbol]scheme.Sexpr
}

func LoadLibrary(s *scheme.Scope, names ...string) error {
  path := strings.Join(names, "/")

  lib, ok := goLibs[path]
  if ok {
    injectGoDefs(s, lib)
  }
  return nil
}

func injectGoDefs(s *scheme.Scope, l *Library) {
  for k, v := range l.data {
    s.Define(k, v)
  }
}
