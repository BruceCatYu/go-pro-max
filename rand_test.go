package gopromax

import "testing"

func TestRand(t *testing.T) {
	a := RandomNumeric(0, 10)
	t.Log(a)
	b := RandomNumeric(2.1, 9.2)
	t.Log(b)
}
