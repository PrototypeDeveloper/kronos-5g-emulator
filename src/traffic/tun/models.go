// Types of package tun.

package tun

import (
	"net/netip"
	"os"
)

// Config describes the interface of one UE.
type Config struct {
	Name     string     // device name, at most 15 octets (kronos-ue1)
	Addr     netip.Addr // the UE's IPv4 address
	MTU      int        // the device MTU
	Table    uint32     // routing table of the device's default route
	Priority uint32     // priority of the source rule
}

// Device is an open TUN device with its address, route and rule. Read
// returns one IPv4 packet the host sent; Write gives the host one.
type Device struct {
	cfg   Config
	f     *os.File
	index int
}

// netlink is an rtnetlink socket with the sequence number of the last request.
type netlink struct {
	fd  int
	seq uint32
}
