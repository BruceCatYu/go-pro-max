package gopromax

import (
	"os"
	"testing"
)

func TestTailN(t *testing.T) {
	f, err := os.Open("./go.mod")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	line := ReadTailN(f, 5)
	t.Log(line.Count, line.Str)
}
