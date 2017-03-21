package main

type Clock struct {
	stop chan bool
	lock bool
}

func NewClock() *Clock {
	clock := &Clock{}
	return clock
}

func (c *Clock) start() {
	if c.stop != nil {
		c.stop <- true
	}
	c.stop = make(chan bool)

	go func(c *Clock) {
		for {
			select {
			case <-engine.ticker.C:
				engine.tick()
			case <-c.stop:
				return
			}
		}
	}(c)
}
