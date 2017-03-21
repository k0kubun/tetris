package main

import (
	"time"
)

type Clock struct {
	ticker   *time.Ticker
	callback func()
	stop     chan bool
	paused   bool
	gameover bool
	lock     bool
}

func NewClock(callback func()) *Clock {
	clock := &Clock{}
	clock.callback = callback
	clock.paused = false
	clock.gameover = false
	clock.lock = false
	return clock
}

func (c *Clock) start() {
	if c.stop != nil {
		c.stop <- true
	}
	c.stop = make(chan bool)

	go func(c *Clock) {
		c.ticker = time.NewTicker(time.Duration(10*(50-2*engine.level)) * time.Millisecond)
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
	c.gameover = false
}

func (c *Clock) updateInterval() {
	c.pause()
	c.start()
}

func (c *Clock) over() {
	c.ticker.Stop()
	c.gameover = true
	c.paused = true
}

func (c *Clock) pause() {
	c.ticker.Stop()
	c.paused = true
}
