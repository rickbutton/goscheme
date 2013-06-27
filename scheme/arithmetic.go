package scheme

func plus(s *Scope, args []Sexpr) (Sexpr, error) {
  if len(args) < 2 {
    return nil, procError("+ requires at least 2 arguments")
  }
  var sum int64 = 0
  for _, a := range args {
    if IsNumber(a) {
      sum += a.(*Number).n
    } else {
      return nil, procError("+ requires all number arguments")
    }
  }
  return NumberFromInt(sum), nil
}

func mul(s *Scope, args []Sexpr) (Sexpr, error) {
  if len(args) < 2 {
    return nil, procError("* requires at least 2 arguments")
  }
  var product int64 = 1
  for _, a := range args {
    if IsNumber(a) {
      product *= a.(*Number).n
    } else {
      return nil, procError("* requires all number arguments")
    }
  }
  return NumberFromInt(product), nil
}

func sub(s *Scope, args []Sexpr) (Sexpr, error) {
  if len(args) != 2 {
    return nil, procError("- requires exactly 2 arguments")
  }
  if !IsNumber(args[0]) || !IsNumber(args[1]) {
    return nil, procError("- requires all number arguments")
  }
  left := args[0].(*Number).n
  right := args[1].(*Number).n
  return NumberFromInt(left - right), nil
}

func div(s *Scope, args []Sexpr) (Sexpr, error) {
  if len(args) != 2 {
    return nil, procError("/ requires exactly 2 arguments")
  }
  if !IsNumber(args[0]) || !IsNumber(args[1]) {
    return nil, procError("/ requires all number arguments")
  }
  left := args[0].(*Number).n
  right := args[1].(*Number).n
  return NumberFromInt(left / right), nil
}
