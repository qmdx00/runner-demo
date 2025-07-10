package internal

import (
	"sync"
)

var (
	tickerInstance *Ticker
	tickerOnce     sync.Once
)

// Ticker represents a single tick in the game loop.
type Ticker struct {
	counter int64
}

func NewTicker() *Ticker {
	tickerOnce.Do(func() {
		tickerInstance = &Ticker{counter: 0}
	})
	return tickerInstance
}

// Counter returns the current tick count.
func (t *Ticker) Tick() {
	t.counter++
}

// Count returns the current count of ticks.
func (t *Ticker) Count() int64 {
	return t.counter
}

// Reset resets the ticker's counter to zero.
func (t *Ticker) Reset() {
	t.counter = 0
}
