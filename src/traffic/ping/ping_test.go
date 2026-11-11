package ping

import (
	"strings"
	"testing"
	"time"
)

// The summary follows the ping of Linux: mdev is the standard deviation.
func TestStats(t *testing.T) {
	var s Stats
	for _, ms := range []int{1, 2, 3, 4} {
		s.add(time.Duration(ms) * time.Millisecond)
	}
	s.Sent, s.Elapsed = 5, 4005*time.Millisecond
	want := "5 transmitted, 4 received, 20% loss, time 4005 ms, rtt min/avg/max/mdev = 1.000/2.500/4.000/1.118 ms"
	if got := s.String(); got != want {
		t.Fatalf("\n got %s\nwant %s", got, want)
	}
	if got := (Stats{Sent: 2}).String(); !strings.HasPrefix(got, "2 transmitted, 0 received, 100% loss") || strings.Contains(got, "rtt") {
		t.Fatal(got)
	}
}
