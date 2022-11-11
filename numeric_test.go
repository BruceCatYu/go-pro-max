package gopromax

import "testing"

func TestMaxNumeric(t *testing.T) {
	n := MaxNumber[uint](2, 1)
	t.Log(n)
}
