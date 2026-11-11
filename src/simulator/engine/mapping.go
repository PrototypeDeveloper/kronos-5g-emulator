package engine

import (
	"sync/atomic"

	"kronos/models"
)

// setMapBits turns items on (clears their bits) or off (sets them) in a.
func setMapBits(a *atomic.Uint32, items models.MapOff, on bool) {
	for {
		old := a.Load()
		v := old | uint32(items)
		if on {
			v = old &^ uint32(items)
		}
		if a.CompareAndSwap(old, v) {
			return
		}
	}
}

// offCount returns 1 if item is off in a, else 0.
func offCount(a *atomic.Uint32, item models.MapItem) int {
	if models.MapOff(a.Load()).On(item) {
		return 0
	}
	return 1
}

// SetMapping, Mapping and Info implement command.Env for a subscriber.
func (s *subscriber) SetMapping(items models.MapOff, on bool) error {
	setMapBits(&s.mapOff, items, on)
	return nil
}
func (s *subscriber) Mapping(item models.MapItem) (int, int) { return offCount(&s.mapOff, item), 1 }
func (s *subscriber) Info(text string) {
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: text})
}

// SetMapping, Mapping and Info implement command.Env for the node: its
// switches are the engine's, so they outlive the association.
func (n *node) SetMapping(items models.MapOff, on bool) error {
	setMapBits(&n.e.nodeMapOff, items, on)
	return nil
}
func (n *node) Mapping(item models.MapItem) (int, int) { return offCount(&n.e.nodeMapOff, item), 1 }
func (n *node) Info(text string) {
	n.e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Text: text})
}

// SetMapping, Mapping and Info implement command.Env for the console: the
// node and every subscriber.
func (c consoleEnv) SetMapping(items models.MapOff, on bool) error {
	setMapBits(&c.e.nodeMapOff, items, on)
	for _, s := range c.e.subs {
		setMapBits(&s.mapOff, items, on)
	}
	return nil
}
func (c consoleEnv) Mapping(item models.MapItem) (int, int) {
	off := offCount(&c.e.nodeMapOff, item)
	for _, s := range c.e.subs {
		off += offCount(&s.mapOff, item)
	}
	return off, 1 + len(c.e.subs)
}
func (c consoleEnv) Info(text string) { c.Print(text) }
