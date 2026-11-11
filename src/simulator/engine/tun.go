package engine

import (
	"fmt"
	"io"
	"net/netip"

	"kronos/models"
	"kronos/traffic/tun"
)

// With ue.tun, a subscriber whose PDU session is up gets a TUN device on the
// host (kronos-ue1 for the first subscriber) that holds the UE address. What
// the host sends from that address (iperf3 -B, curl --interface) is read
// from the device and sent uplink through the session; downlink packets that
// no built-in endpoint (ping, iperf) takes are written to it. The device
// stays while the UE is idle (uplink packets are dropped then) and goes with
// the session, a linkdown or the end of the engine.

// tunDevice is an open TUN device (tun.Device; a fake in tests).
type tunDevice interface {
	io.ReadWriteCloser
	Name() string
}

// tunOpener creates the device of subscriber index for the UE address.
type tunOpener func(index int, ueIP netip.Addr) (tunDevice, error)

// openHostTUN creates the host TUN device of a subscriber with the UE address and its source routing.
func openHostTUN(index int, ueIP netip.Addr) (tunDevice, error) {
	return tun.Open(tun.Config{
		Name:     fmt.Sprintf("kronos-ue%d", index+1),
		Addr:     ueIP,
		MTU:      tunMTU,
		Table:    uint32(tunTable + 1 + index),
		Priority: tunPriority,
	})
}

// tunUp opens the device for a session that came up, unless it is already
// there for the same address (the UE came back from idle). One device per
// subscriber: it carries the first session that brought it up.
func (s *subscriber) tunUp(t models.Tunnel) {
	if s.e.openTUN == nil || !t.UEIP.Is4() {
		return
	}
	if cur := s.tun.Load(); cur != nil {
		if cur.session != t.SessionID || cur.ueIP == t.UEIP {
			return
		}
		s.tunDown() // the session has a new address
	}
	dev, err := s.e.openTUN(s.index, t.UEIP)
	if err != nil {
		s.e.publish(models.Event{Kind: models.KindError, Source: s.imsi, Text: "TUN: " + err.Error()})
		return
	}
	x := &ueTUN{dev: dev, session: t.SessionID, ueIP: t.UEIP}
	s.tun.Store(x)
	go s.readTUN(x)
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: fmt.Sprintf(
		"TUN %s up: %s, MTU %d (the host's packets from %s go through PDU session %d, e.g. iperf3 -B %s)",
		dev.Name(), t.UEIP, tunMTU, t.UEIP, t.SessionID, t.UEIP)})
}

// dropGone stops the N3 tunnels and the TUN device of the PDU sessions the
// UE no longer has (released, or the UE deregistered; an idle UE keeps
// them). It runs after each message of the actor, with its snapshot.
func (s *subscriber) dropGone(st models.UEStatus) {
	x := s.tun.Load()
	if len(s.tunnels) == 0 && x == nil {
		return
	}
	has := func(id uint8) bool {
		for _, ss := range st.Sessions {
			if ss.ID == id {
				return true
			}
		}
		return false
	}
	for id, t := range s.tunnels {
		if !has(id) {
			s.e.uplane.unbind(t.DLTEID)
			delete(s.tunnels, id)
		}
	}
	if x != nil && !has(x.session) {
		s.tunDown()
	}
}

// tunDown removes the device, if any.
func (s *subscriber) tunDown() {
	x := s.tun.Swap(nil)
	if x == nil {
		return
	}
	x.dev.Close()
	text := "TUN " + x.dev.Name() + " removed"
	if n := x.dropped.Load(); n > 0 {
		text += fmt.Sprintf(" (%d uplink packet(s) dropped: no tunnel or not from the UE)", n)
	}
	s.e.publish(models.Event{Kind: models.KindInfo, Source: s.imsi, Text: text})
}

// readTUN sends what the host sends through the device uplink, until the
// device is closed.
func (s *subscriber) readTUN(x *ueTUN) {
	buf := make([]byte, packetBuffer)
	ue := x.ueIP.As4()
	for {
		n, err := x.dev.Read(buf)
		if err != nil {
			return
		}
		pkt := buf[:n]
		if n < ipv4HeaderLen || pkt[0]>>4 != ipv4Version || [4]byte(pkt[12:16]) != ue {
			x.dropped.Add(1)
			continue
		}
		if s.e.uplane.sendFor(s, x.session, pkt) != nil {
			x.dropped.Add(1)
		}
	}
}

// toTUN gives the host a downlink packet that no built-in endpoint took.
func (s *subscriber) toTUN(pkt []byte) {
	if x := s.tun.Load(); x != nil {
		x.dev.Write(pkt)
	}
}
