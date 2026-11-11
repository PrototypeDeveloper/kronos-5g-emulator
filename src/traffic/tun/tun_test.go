package tun

import (
	"encoding/hex"
	"errors"
	"net/netip"
	"os"
	"strings"
	"testing"
)

// The requests carry what iproute2 sends for the same settings (seen with
// strace): ip link set dev X mtu 1456 up, ip addr replace 10.60.0.1/32 dev X,
// ip route replace default dev X src 10.60.0.1 table 10001 proto static
// scope link, ip rule add from 10.60.0.1 lookup 10001 pref 1000.
func TestMessages(t *testing.T) {
	if native.Uint16([]byte{1, 0}) != 1 {
		t.Skip("the expected octets are little-endian")
	}
	ue := netip.MustParseAddr("10.60.0.1")
	for _, c := range []struct {
		name string
		got  []byte
		want string
	}{
		{"link", linkUpMsg(7, 1456), "00000000" + "07000000" + "01000000" + "01000000" + "08000400" + "b0050000"},
		{"addr", addrMsg(7, ue), "02200000" + "07000000" + "08000200" + "0a3c0001" + "08000100" + "0a3c0001"},
		{"route", routeMsg(7, ue, 10001), "02000000" + "0004fd01" + "00000000" + "08000f00" + "11270000" + "08000400" + "07000000" + "08000700" + "0a3c0001"},
		{"rule", ruleMsg(ue, 10001, 1000), "02002000" + "00000001" + "00000000" + "08000200" + "0a3c0001" + "08000f00" + "11270000" + "08000600" + "e8030000"},
	} {
		if g := hex.EncodeToString(c.got); g != c.want {
			t.Errorf("%s:\n got %s\nwant %s", c.name, g, c.want)
		}
	}
	m := nlMessage(0x20, 0x605, 9, ruleMsg(ue, 10001, 1000))
	if g := hex.EncodeToString(m[:16]); g != "34000000"+"2000"+"0506"+"09000000"+"00000000" {
		t.Errorf("header %s", g)
	}
}

func TestConfigCheck(t *testing.T) {
	ok := Config{Name: "kronos-ue1", Addr: netip.MustParseAddr("10.60.0.1"), MTU: 1456, Table: 10001, Priority: 1000}
	if err := ok.check(); err != nil {
		t.Fatal(err)
	}
	for _, c := range []struct {
		edit func(*Config)
		want string
	}{
		{func(c *Config) { c.Name = "" }, "name"},
		{func(c *Config) { c.Name = "kronos-ue1234567" }, "name"},
		{func(c *Config) { c.Addr = netip.MustParseAddr("fd00::1") }, "not IPv4"},
		{func(c *Config) { c.MTU = 100 }, "MTU"},
		{func(c *Config) { c.Table = 254 }, "reserved"},
		{func(c *Config) { c.Table = 0 }, "reserved"},
	} {
		x := ok
		c.edit(&x)
		if err := x.check(); err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%+v: %v", x, err)
		}
	}
}

// Without CAP_NET_ADMIN, Open says what is missing. With it, the test
// leaves the host alone.
func TestOpenWithoutPermission(t *testing.T) {
	if os.Geteuid() == 0 {
		t.Skip("running as root")
	}
	if _, err := os.Stat("/dev/net/tun"); err != nil {
		t.Skip("no /dev/net/tun")
	}
	d, err := Open(Config{Name: "kronos-test0", Addr: netip.MustParseAddr("10.255.255.254"), MTU: 1456, Table: 10999, Priority: 1000})
	if err == nil {
		d.Close()
		t.Skip("this process has CAP_NET_ADMIN")
	}
	if !errors.Is(err, ErrPermission) {
		t.Fatalf("%v", err)
	}
}

func TestHostAddressRefused(t *testing.T) {
	_, err := Open(Config{Name: "kronos-test0", Addr: netip.MustParseAddr("127.0.0.1"), MTU: 1456, Table: 10999, Priority: 1000})
	if err == nil || !strings.Contains(err.Error(), "already an address of this host") {
		t.Fatalf("%v", err)
	}
}

func TestCapEff(t *testing.T) {
	status := "Name:\tgnodeb\nCapInh:\t0000000000000000\nCapPrm:\t0000000000001000\nCapEff:\t0000000000001000\n"
	if !capEffHas(status, capNetAdmin) || capEffHas(status, 13) || capEffHas("CapEff:\t0000000000000000\n", capNetAdmin) || capEffHas("", capNetAdmin) {
		t.Fatal("CapEff")
	}
}
