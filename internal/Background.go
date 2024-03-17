package internal

import "time"

type Background struct {
	Callback func()
	timer    time.Duration
}

func (b *Background) Run() {
	t := time.NewTicker(b.timer)
	for {
		select {
		case <-t.C:
			b.Callback()
		}
	}
}
