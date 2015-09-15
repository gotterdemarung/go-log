package log

import (
	"sync"
)

type AsyncDispatcher struct {
	delivery []chan *Packet
	wait sync.WaitGroup
}

func NewAsyncDispatcher() *AsyncDispatcher {
	return &AsyncDispatcher{
		delivery: []chan *Packet{},
	}
}

func (ad *AsyncDispatcher) Register(a Appender) {
	// Building channel for listening
	ch := make(chan *Packet)

	// Goroutine to listen
	go func() {
		for val := range ch {
			a(val)
			ad.wait.Done()
		}
	}()

	ad.delivery = append(ad.delivery, ch)
}

func (ad *AsyncDispatcher) Dispatch(l *Packet) {
	for _, ch := range ad.delivery {
		ad.wait.Add(1)
		ch <- l
	}
}

func (ad *AsyncDispatcher) Wait() {
	ad.wait.Wait()
}


