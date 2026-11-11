// Package iperf is an iperf3 client for the emulated UEs, against an
// ordinary iperf3 server (iperf3 -s) on the DN, uplink or downlink (iperf3
// -R). Run makes a UDP test through the UE's PDU session, with the control
// connection on the UE's small TCP (package tcp). RunTCP makes a TCP test
// with host sockets bound to the UE address, which the UE's TUN device
// carries through the session. Serve is a minimal UDP server for tests.
package iperf

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net"
	"time"

	"kronos/traffic/tcp"
)

// The UDP "connect" handshake: iperf3 writes these ints in host byte order,
// so on the usual little-endian hosts the octets come reversed.
var (
	udpConnectMsg   = []byte{0x39, 0x38, 0x37, 0x36} // 0x36373839 ("6789")
	udpConnectReply = []byte{0x36, 0x37, 0x38, 0x39} // 0x39383736 ("9876")
)

// stateName names a control state as iperf_api.h does.
func stateName(s int8) string {
	switch s {
	case testStart:
		return "TEST_START"
	case testRunning:
		return "TEST_RUNNING"
	case testEnd:
		return "TEST_END"
	case paramExchange:
		return "PARAM_EXCHANGE"
	case createStreams:
		return "CREATE_STREAMS"
	case serverTerminate:
		return "SERVER_TERMINATE"
	case clientTerminate:
		return "CLIENT_TERMINATE"
	case exchangeResults:
		return "EXCHANGE_RESULTS"
	case displayResults:
		return "DISPLAY_RESULTS"
	case iperfDone:
		return "IPERF_DONE"
	case accessDenied:
		return "ACCESS_DENIED"
	case serverError:
		return "SERVER_ERROR"
	}
	return fmt.Sprintf("state %d", s)
}

// smallControl uses a connection of the UE's small TCP.
func smallControl(c *tcp.Conn) control { return control{c, c.SetReadDeadline} }

// hostControl uses a host socket.
func hostControl(c net.Conn) control {
	return control{c, func(t time.Time) { c.SetReadDeadline(t) }}
}

// readState reads one control state, waiting at most timeout.
func (k control) readState(timeout time.Duration) (int8, error) {
	k.deadline(time.Now().Add(timeout))
	var b [1]byte
	if _, err := io.ReadFull(k.c, b[:]); err != nil {
		return 0, fmt.Errorf("cannot read the control connection: %w", err)
	}
	return int8(b[0]), nil
}

// writeState sends one control state.
func (k control) writeState(s int8) error {
	_, err := k.c.Write([]byte{byte(s)})
	return err
}

// writeJSON sends a 4-octet length (network order) and the JSON text.
func (k control) writeJSON(v any) error {
	j, err := json.Marshal(v)
	if err != nil {
		return err
	}
	b := binary.BigEndian.AppendUint32(nil, uint32(len(j)))
	_, err = k.c.Write(append(b, j...))
	return err
}

// readJSON reads a length-prefixed JSON text into v, waiting at most timeout.
func (k control) readJSON(v any, timeout time.Duration) error {
	k.deadline(time.Now().Add(timeout))
	var n [4]byte
	if _, err := io.ReadFull(k.c, n[:]); err != nil {
		return fmt.Errorf("cannot read the control connection: %w", err)
	}
	size := binary.BigEndian.Uint32(n[:])
	if size > maxJSONLen {
		return fmt.Errorf("the JSON on the control connection is too large (%d octets)", size)
	}
	b := make([]byte, size)
	if _, err := io.ReadFull(k.c, b); err != nil {
		return fmt.Errorf("cannot read the control connection: %w", err)
	}
	return json.Unmarshal(b, v)
}

// serverErrorText reads what follows SERVER_ERROR: i_errno and errno.
func (k control) serverErrorText() string {
	k.deadline(time.Now().Add(promptReadTimeout))
	var b [8]byte
	if _, err := io.ReadFull(k.c, b[:]); err != nil {
		return "server error"
	}
	return fmt.Sprintf("server error (i_errno %d, errno %d)", int32(binary.BigEndian.Uint32(b[:4])), int32(binary.BigEndian.Uint32(b[4:])))
}

// expect reads one state and fails unless it is want; an error state ends the test.
func (k control) expect(want int8, timeout time.Duration) error {
	s, err := k.readState(timeout)
	if err != nil {
		return err
	}
	switch s {
	case want:
		return nil
	case accessDenied:
		return fmt.Errorf("the server is busy with another test (ACCESS_DENIED)")
	case serverError:
		return fmt.Errorf("%s", k.serverErrorText())
	case serverTerminate:
		return fmt.Errorf("the server ended the test (SERVER_TERMINATE)")
	}
	return fmt.Errorf("expected %s, got %s", stateName(want), stateName(s))
}

// take accounts one received datagram p that arrived at at.
func (r *receiver) take(p []byte, at time.Time) {
	if len(p) < udpHeader {
		return
	}
	sec, usec := binary.BigEndian.Uint32(p[0:]), binary.BigEndian.Uint32(p[4:])
	pcount := uint64(binary.BigEndian.Uint32(p[8:]))
	sent := float64(sec) + float64(usec)/1e6
	transit := math.Abs(float64(at.UnixNano())/1e9 - sent)
	r.mu.Lock()
	defer r.mu.Unlock()
	r.bytes += uint64(len(p))
	r.datagrams++
	// A gap counts as lost; a late datagram is out of order and undoes one loss.
	if pcount >= r.count+1 {
		if pcount > r.count+1 {
			r.lost += pcount - 1 - r.count
		}
		r.count = pcount
	} else {
		r.outOfOrder++
		if r.lost > 0 {
			r.lost--
		}
	}
	// Jitter: J += (|D| - J) / 16 over the transit times.
	if !r.started {
		r.prevTransit, r.started = transit, true
	}
	d := math.Abs(transit - r.prevTransit)
	r.prevTransit = transit
	r.jitter += (d - r.jitter) / 16
}

// snapshot returns the counters so far.
func (r *receiver) snapshot() receiverSnapshot {
	r.mu.Lock()
	defer r.mu.Unlock()
	return receiverSnapshot{r.bytes, r.datagrams, r.count, r.lost, r.outOfOrder, r.jitter}
}

// putHeader writes the sender time and the packet counter.
func putHeader(b []byte, now time.Time, pcount uint32) {
	binary.BigEndian.PutUint32(b[0:], uint32(now.Unix()))
	binary.BigEndian.PutUint32(b[4:], uint32(now.Nanosecond()/1000))
	binary.BigEndian.PutUint32(b[8:], pcount)
}
