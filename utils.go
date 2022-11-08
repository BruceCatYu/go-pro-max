package gopromax

type Ternary struct {
	Expr1, Expr2, Expr3 any
}

func (t *Ternary) Calculate() any {
	var e1Res bool
	e1, ok := t.Expr1.(Ternary)
	if ok {
		e1Res = e1.Calculate().(bool)
	} else {
		e1Res = t.Expr1.(bool)
	}

	if e1Res {
		e2, ok := t.Expr2.(Ternary)
		if ok {
			return e2.Calculate()
		} else {
			return t.Expr2
		}
	} else {
		e3, ok := t.Expr3.(Ternary)
		if ok {
			return e3.Calculate()
		} else {
			return t.Expr3
		}
	}
}
