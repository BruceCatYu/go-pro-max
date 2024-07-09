package gopromax

import (
	"testing"
)

func TestTree(t *testing.T) {
	t1 := NewTree[int](10)
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	t.Log(t1.Contains(15))
	t.Log(t1.Contains(40))
}
