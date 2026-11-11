package engine

import (
	"sync"
	"time"

	"kronos/models"
)

// NewBus keeps up to history events for late subscribers.
func NewBus(history int) *Bus {
	if history < 1 {
		history = 1
	}
	return &Bus{subs: map[int]chan models.Event{}, ring: make([]models.Event, history),
		count: models.Counters{Messages: map[string]*models.MessageCount{}}}
}

// Publish stamps the event with a sequence number and time and delivers it.
func (b *Bus) Publish(e models.Event) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.seq++
	e.Seq = b.seq
	if e.Time.IsZero() {
		e.Time = time.Now()
	}
	b.ring[b.head] = e
	b.head = (b.head + 1) % len(b.ring)
	if b.head == 0 {
		b.full = true
	}
	b.tally(e)
	for _, ch := range b.subs {
		select {
		case ch <- e:
		default:
			b.dropped++
		}
	}
}

// tally updates the counters (b.mu held).
func (b *Bus) tally(e models.Event) {
	c := &b.count
	switch e.Kind {
	case models.KindError:
		c.Errors++
	case models.KindWarn:
		c.Warnings++
	}
	if e.Handler == "unhandled" {
		c.Unhandled++
	}
	if e.NGAP == "" || (e.Dir != "tx" && e.Dir != "rx") {
		return
	}
	m := c.Messages[e.NGAP]
	if m == nil {
		m = &models.MessageCount{}
		c.Messages[e.NGAP] = m
	}
	if e.Dir == "tx" {
		c.TX++
		m.TX++
	} else {
		c.RX++
		m.RX++
	}
}

// Counters returns a copy of the counters.
func (b *Bus) Counters() models.Counters {
	b.mu.Lock()
	defer b.mu.Unlock()
	c := b.count
	c.Dropped = b.dropped
	c.Messages = make(map[string]*models.MessageCount, len(b.count.Messages))
	for k, v := range b.count.Messages {
		m := *v
		c.Messages[k] = &m
	}
	return c
}

// Subscribe returns a channel of future events and a function that ends the
// subscription and closes the channel.
func (b *Bus) Subscribe(buffer int) (<-chan models.Event, func()) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.subscribe(buffer)
}

// SubscribeSince is Subscribe that also returns the kept events after since
// (only the last tail of them when tail > 0), with no gap or overlap between
// the history and the channel. A since beyond the last event (a number from
// an earlier run) counts from the start.
func (b *Bus) SubscribeSince(since uint64, tail, buffer int) ([]models.Event, <-chan models.Event, func()) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if since > b.seq {
		since = 0
	}
	hist := b.history(since)
	if tail > 0 && len(hist) > tail {
		hist = hist[len(hist)-tail:]
	}
	ch, cancel := b.subscribe(buffer)
	return hist, ch, cancel
}

// Event returns the kept event with sequence number seq.
func (b *Bus) Event(seq uint64) (models.Event, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	kept := uint64(b.head)
	if b.full {
		kept = uint64(len(b.ring))
	}
	if seq == 0 || seq > b.seq || b.seq-seq >= kept {
		return models.Event{}, false
	}
	back := int(b.seq - seq) // 0 is the newest
	return b.ring[(b.head-1-back+len(b.ring))%len(b.ring)], true
}

// subscribe registers a channel (b.mu held).
func (b *Bus) subscribe(buffer int) (<-chan models.Event, func()) {
	id := b.next
	b.next++
	ch := make(chan models.Event, buffer)
	b.subs[id] = ch
	var once sync.Once
	return ch, func() {
		once.Do(func() {
			b.mu.Lock()
			defer b.mu.Unlock()
			delete(b.subs, id)
			close(ch)
		})
	}
}

// History returns the kept events with a sequence number above since.
func (b *Bus) History(since uint64) []models.Event {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.history(since)
}

// history returns the kept events with a sequence number above since, oldest first (b.mu held).
func (b *Bus) history(since uint64) []models.Event {
	var all []models.Event
	if b.full {
		all = append(all, b.ring[b.head:]...)
	}
	all = append(all, b.ring[:b.head]...)
	out := all[:0:0]
	for _, e := range all {
		if e.Seq > since {
			out = append(out, e)
		}
	}
	return out
}

// Dropped returns how many deliveries were lost to slow subscribers.
func (b *Bus) Dropped() uint64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.dropped
}
