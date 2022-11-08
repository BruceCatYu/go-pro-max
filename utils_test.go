package gopromax

import (
	"reflect"
	"testing"
)

func TestTernary(t *testing.T) {
	// a := Ternary{
	// 	Ternary{2 > 1, false, true},
	// 	Ternary{false, 3, 4},
	// 	6,
	// }
	a := Ternary{2 > 1, false, true}
	res := a.Calculate()
	c := reflect.TypeOf(res)

	t.Log(c.Kind())
}
