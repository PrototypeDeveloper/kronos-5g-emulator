package engine

import (
	"errors"
	"time"

	"kronos/command"
)

// newControls returns the turns of control(), none given yet.
func newControls() *controls {
	return &controls{last: map[string]time.Time{}, cut: make(chan struct{})}
}

// reserve gives the next turn at key: now, or interval after the turn
// before, whichever is later. It also returns the channel that ends the wait.
func (c *controls) reserve(key string, interval time.Duration) (time.Time, <-chan struct{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	at := time.Now()
	if last, ok := c.last[key]; ok && last.Add(interval).After(at) {
		at = last.Add(interval)
	}
	c.last[key] = at
	return at, c.cut
}

// reset ends every wait and forgets the turns (the link went down).
func (c *controls) reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	close(c.cut)
	c.cut = make(chan struct{})
	c.last = map[string]time.Time{}
}

// Control implements command.Env: the subscriber waits for its turn.
func (s *subscriber) Control(key string, interval time.Duration) error {
	at, cut := s.e.controls.reserve(key, interval)
	d := time.Until(at)
	if d <= 0 {
		return nil
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return nil
	case <-cut:
	case <-s.e.runCtx().Done():
	}
	return command.ErrInterrupted
}

// Control implements command.Env; only subscribers pace their scenario.
func (n *node) Control(string, time.Duration) error {
	return errors.New("control is a subscriber command")
}

// Control implements command.Env; control is written in the scenario only.
func (consoleEnv) Control(string, time.Duration) error {
	return errors.New("control is a scenario command")
}
