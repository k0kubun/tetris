package main

import (
	"time"
)

type Clock struct {
	ticker   *time.Ticker
	callback func()
	stop     chan bool
	paused   bool
}

func NewClock(callback func()) *Clock {
	clock := &Clock{}
	clock.callback = callback
	clock.paused = false
	return clock
}

func (c *Clock) start() {
	if c.stop != nil {
		c.stop <- true
	}
	c.stop = make(chan bool)

	go func(c *Clock) {
		c.ticker = time.NewTicker(time.Duration(10*(50-2*level)) * time.Millisecond)
		for {
			select {
			case <-c.ticker.C:
				c.callback()
			case <-c.stop:
				return
			}
		}
	}(c)
	c.paused = false
}

func (c *Clock) updateInterval() {
	c.pause()
	c.start()
}

func (c *Clock) pause() {
	c.ticker.Stop()
	c.paused = true
}
