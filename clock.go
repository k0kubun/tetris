package main

import (
	"time"
)

type Clock struct {
	ticker *time.Ticker
	stop   chan bool
	paused bool
	lock   bool
}

func NewClock() *Clock {
	clock := &Clock{}
	clock.paused = false
	clock.lock = false
	engine.NewGame()
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
				engine.tick()
			case <-c.stop:
				return
			}
		}
	}(c)
	c.paused = false
	engine.gameover = false
}

func (c *Clock) updateInterval() {
	c.pause()
	c.start()
}

func (c *Clock) over() {
	c.ticker.Stop()
	c.paused = true
}

func (c *Clock) pause() {
	c.ticker.Stop()
	c.paused = true
}
