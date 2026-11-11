package traffic

import (
	"kronos/models"
	"net/netip"
	"testing"
)

var (
	ue = netip.MustParseAddr("10.60.0.7")
	dn = netip.MustParseAddr("10.10.10.20")
)

func TestUDPRoundTrip(t *testing.T) {
	pkt := BuildUDP(netip.AddrPortFrom(ue, 50000), netip.AddrPortFrom(dn, 5201), 7, []byte("hello"))
	ip, ok := ParseIPv4(pkt)
	if !ok || ip.Src != ue || ip.Dst != dn || ip.Proto != models.ProtoUDP || ip.ID != 7 {
		t.Fatalf("ip %+v %v", ip, ok)
	}
	u, ok := ParseUDP(ip)
	if !ok || u.Src.Port() != 50000 || u.Dst.Port() != 5201 || string(u.Payload) != "hello" {
		t.Fatalf("udp %+v %v", u, ok)
	}
	pkt[len(pkt)-1] ^= 1 // corrupt the payload: the checksum catches it
	ip, _ = ParseIPv4(pkt)
	if _, ok := ParseUDP(ip); ok {
		t.Fatal("bad UDP checksum accepted")
	}
	pkt[10] ^= 1 // and the header checksum
	if _, ok := ParseIPv4(pkt); ok {
		t.Fatal("bad IP checksum accepted")
	}
}

func TestEcho(t *testing.T) {
	pkt := BuildEcho(ue, dn, 1, 8, 0x1234, 9, []byte("abcdefgh"))
	ip, _ := ParseIPv4(pkt)
	e, ok := ParseEcho(ip)
	if !ok || e.Type != 8 || e.ID != 0x1234 || e.Seq != 9 || string(e.Data) != "abcdefgh" || pkt[6] != 0x40 {
		t.Fatalf("echo %+v %v (flags %x)", e, ok, pkt[6])
	}
	if big := BuildEcho(ue, dn, 1, 8, 1, 1, make([]byte, 2000)); big[6] != 0 {
		t.Fatal("a packet over 1500 octets must be fragmentable")
	}
}

// TestEphemeralStart: each stack starts its ephemeral ports at a random
// place, so a new process does not reuse the last one's ports.
func TestEphemeralStart(t *testing.T) {
	seen := map[uint16]bool{}
	for i := 0; i < 8; i++ {
		p, err := NewStack().Bind(models.ProtoTCP, 0, func(models.IPv4) {})
		if err != nil || p < 49152 {
			t.Fatalf("port %d %v", p, err)
		}
		seen[p] = true
	}
	if len(seen) < 4 {
		t.Fatalf("the first ports of 8 stacks: %v", seen)
	}
}

func TestStackDelivery(t *testing.T) {
	st := NewStack()
	var got []string
	p1, err := st.Bind(models.ProtoUDP, 0, func(ip models.IPv4) { got = append(got, "udp1") })
	if err != nil || p1 < 49152 {
		t.Fatalf("first ephemeral port %d %v", p1, err)
	}
	p2, _ := st.Bind(models.ProtoUDP, 0, func(ip models.IPv4) { got = append(got, "udp2") })
	if p2 != p1+1 && !(p1 == 65535 && p2 == 49152) {
		t.Fatalf("second ephemeral port %d after %d", p2, p1)
	}
	if _, err := st.Bind(models.ProtoUDP, p1, nil); err == nil {
		t.Fatal("a taken port was given twice")
	}
	id, _ := st.Bind(models.ProtoICMP, 0, func(ip models.IPv4) { got = append(got, "icmp") })
	st.Deliver(BuildUDP(netip.AddrPortFrom(dn, 5201), netip.AddrPortFrom(ue, p2), 1, nil))
	st.Deliver(BuildEcho(dn, ue, 1, 0, id, 1, nil))
	st.Deliver(BuildEcho(dn, ue, 1, 8, id, 1, nil)) // a request has no owner here
	st.Deliver(BuildUDP(netip.AddrPortFrom(dn, 5201), netip.AddrPortFrom(ue, 40000), 1, nil))
	st.Unbind(models.ProtoUDP, p1)
	st.Deliver(BuildUDP(netip.AddrPortFrom(dn, 5201), netip.AddrPortFrom(ue, p1), 1, nil))
	if len(got) != 2 || got[0] != "udp2" || got[1] != "icmp" || st.Unknown() != 3 {
		t.Fatalf("delivered %v, unknown %d", got, st.Unknown())
	}
}

func TestParseRate(t *testing.T) {
	for in, want := range map[string]uint64{"100M": 100e6, "1G": 1e9, "1.5g": 1.5e9, "500k": 500e3, "2000": 2000} {
		if got, err := ParseRate(in); err != nil || got != want {
			t.Errorf("%s: %d %v", in, got, err)
		}
	}
	for _, bad := range []string{"", "M", "-1M", "x", "0"} {
		if _, err := ParseRate(bad); err == nil {
			t.Errorf("%q accepted", bad)
		}
	}
	if FormatRate(99.84e6) != "99.8 Mbit/s" || FormatBytes(11_900_000) != "11.9 MB" {
		t.Fatal(FormatRate(99.84e6), FormatBytes(11_900_000))
	}
}
