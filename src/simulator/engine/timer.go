package engine

import (
	"errors"
	"fmt"
	"time"

	"kronos/command"
	"kronos/models"
)

// Scenario timers: timer(t) starts the timer of a duration variable (again
// from the start if it runs), stoptimer(t) stops it. When it runs out, the
// actor gets a timerMsg and runs expire(t): a subscriber's from the state it
// is in then, or any's; the node's from the node block. A linkdown and a
// reset stop every timer.

// start (re)starts the timer name for d; when it runs out, post gets its message.
func (ts *timers) start(name string, d time.Duration, post func(timerMsg)) {
	if ts.runs == nil {
		ts.runs = map[string]*timerRun{}
	}
	ts.stop(name)
	ts.gen++
	m := timerMsg{name: name, gen: ts.gen}
	ts.runs[name] = &timerRun{t: time.AfterFunc(d, func() { post(m) }), gen: m.gen}
}

// stop stops the timer name; it reports whether it ran.
func (ts *timers) stop(name string) bool {
	r, ok := ts.runs[name]
	if ok {
		r.t.Stop()
		delete(ts.runs, name)
	}
	return ok
}

// stopAll stops every timer.
func (ts *timers) stopAll() {
	for name := range ts.runs {
		ts.stop(name)
	}
}

// expired takes the message of a timer that ran out; false for a late one
// of a timer since restarted or stopped.
func (ts *timers) expired(m timerMsg) bool {
	r, ok := ts.runs[m.name]
	if !ok || r.gen != m.gen {
		return false
	}
	delete(ts.runs, m.name)
	return true
}

// timerDuration reads the duration of a timer variable.
func timerDuration(name string, v command.Value, err error) (time.Duration, error) {
	switch {
	case err != nil:
		return 0, err
	case v.Type != command.TypeDuration:
		return 0, fmt.Errorf("%s is not a duration (%s)", name, v.Literal())
	case v.Dur <= 0:
		return 0, fmt.Errorf("the duration of %s must be positive (%s)", name, v.Literal())
	}
	return v.Dur, nil
}

// Timer implements command.Env for a subscriber.
func (s *subscriber) Timer(name string) error {
	v, err := s.Var(name)
	d, err := timerDuration(name, v, err)
	if err != nil {
		return err
	}
	s.timers.start(name, d, func(m timerMsg) {
		select {
		case s.inbox <- m:
		case <-s.e.runCtx().Done():
		}
	})
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: fmt.Sprintf("timer %s started (%s)", name, d)})
	return nil
}

// StopTimer implements command.Env for a subscriber.
func (s *subscriber) StopTimer(name string) error {
	if s.timers.stop(name) {
		s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: "timer " + name + " stopped"})
	}
	return nil
}

// timerExpired runs expire(name) of the current state, else of any.
func (s *subscriber) timerExpired(m timerMsg) {
	if !s.timers.expired(m) {
		return
	}
	block := s.state
	st, _ := s.e.scen.State(s.state)
	x, ok := st.Expire(m.name)
	if !ok {
		block = "any"
		x, ok = s.e.scen.Any.Expire(m.name)
	}
	if !ok {
		s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, State: s.state,
			Text: fmt.Sprintf("timer %s expired: no expire(%s) in %s or any", m.name, m.name, s.state)})
		return
	}
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, State: s.state, Handler: block, Text: "timer " + m.name + " expired"})
	s.run(block+"/expire("+m.name+")", x.Body)
	s.afterHandler()
}

// Timer implements command.Env for the node (a public variable).
func (n *node) Timer(name string) error {
	v, err := n.Var(name)
	d, err := timerDuration(name, v, err)
	if err != nil {
		return err
	}
	n.timers.start(name, d, func(m timerMsg) {
		select {
		case n.inbox <- m:
		case <-n.done:
		}
	})
	n.e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Text: fmt.Sprintf("timer %s started (%s)", name, d)})
	return nil
}

// StopTimer implements command.Env for the node.
func (n *node) StopTimer(name string) error {
	if n.timers.stop(name) {
		n.e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Text: "timer " + name + " stopped"})
	}
	return nil
}

// timerExpired runs expire(name) of the node block.
func (n *node) timerExpired(m timerMsg) {
	if !n.timers.expired(m) {
		return
	}
	x, ok := n.e.scen.Node.Expire(m.name)
	if !ok {
		n.e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Text: fmt.Sprintf("timer %s expired: no expire(%s) in node", m.name, m.name)})
		return
	}
	n.e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Handler: "node", Text: "timer " + m.name + " expired"})
	n.run("node/expire("+m.name+")", x.Body)
}

// Timer and StopTimer implement command.Env for the console: timers belong
// to subscribers (and to the node in a scenario).
func (c consoleEnv) Timer(string) error { return errors.New("timer is a subscriber command") }
func (c consoleEnv) StopTimer(string) error {
	return errors.New("stoptimer is a subscriber command")
}
