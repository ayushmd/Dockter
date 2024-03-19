package internal

import (
	"time"
)

type Background struct {
	Timer    time.Duration
	Callback func()
}

func (b *Background) Run() {

	t := time.NewTicker(b.Timer)
	for {
		select {
		case <-t.C:
			b.Callback()
		}
	}
}
