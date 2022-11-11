package gopromax

import "sync"

type WaitGroup struct {
	sync.WaitGroup
}

func (w *WaitGroup) Wrap(f func()) {
	w.Add(1)
	go func() {
		f()
		w.Done()
	}()
}
