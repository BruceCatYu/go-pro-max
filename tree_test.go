package gopromax

import (
	"strings"
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func OrderPerson(p1, p2 Person) int {
	out := strings.Compare(p1.Name, p2.Name)
	if out == 0 {
		out = p1.Age - p2.Age
	}
	return out
}

func TestTree(t *testing.T) {
	t1 := NewTree(NumericOrderable[int])
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	t.Log(t1.Contains(15))
	t.Log(t1.Contains(40))

	t2 := NewTree(OrderPerson)
	t2.Add(Person{"Bob", 30})
	t2.Add(Person{"Maria", 35})
	t2.Add(Person{"Aob", 50})
	t.Log(t2.Contains(Person{"Bob", 30}))
}
