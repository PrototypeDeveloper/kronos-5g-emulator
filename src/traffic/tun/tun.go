// Package tun gives a UE an interface on the host: a TUN device that holds
// the UE's IPv4 address, a default route in a table of its own, and a source
// rule that sends what the host sends from that address to the table.
// Programs on the host (iperf3 -B, curl --interface) then talk through the
// PDU session with the kernel's TCP/IP: the device hands their packets to
// the pseudo-call, which sends them uplink, and takes the downlink ones.
//
// Creating the device and the routing needs CAP_NET_ADMIN (sudo setcap
// cap_net_admin+ep ./gnodeb, or root).
package tun

import (
	"errors"
	"fmt"
	"net"
	"net/netip"
	"os"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

// ErrPermission is returned when the process may not create the device or
// change the routing.
var ErrPermission = errors.New("this needs CAP_NET_ADMIN: run sudo setcap cap_net_admin+ep ./gnodeb after each build, or start gnodeb as root")

// Open creates the device, gives it the address and the MTU, brings it up
// and routes what the host sends from the address through it.
func Open(c Config) (*Device, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	if host, err := isHostAddr(c.Addr); err != nil {
		return nil, err
	} else if host {
		return nil, fmt.Errorf("the UE address %s is already an address of this host", c.Addr)
	}
	fd, err := syscall.Open("/dev/net/tun", syscall.O_RDWR|syscall.O_CLOEXEC|syscall.O_NONBLOCK, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot open /dev/net/tun: %w", permission(err))
	}
	// Create the TUN device (no packet information header).
	var ifr struct {
		name  [syscall.IFNAMSIZ]byte
		flags uint16
		_     [22]byte
	}
	copy(ifr.name[:], c.Name)
	ifr.flags = syscall.IFF_TUN | syscall.IFF_NO_PI
	if _, _, e := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), tunSetIff, uintptr(unsafe.Pointer(&ifr))); e != 0 {
		syscall.Close(fd)
		if e == syscall.EBUSY {
			return nil, fmt.Errorf("cannot create the TUN device %s: a device of that name is in use", c.Name)
		}
		return nil, fmt.Errorf("cannot create the TUN device %s: %w", c.Name, permission(e))
	}
	// Non-blocking, so reads go through the runtime poller and Close ends them.
	d := &Device{cfg: c, f: os.NewFile(uintptr(fd), "/dev/net/tun:"+c.Name)}
	if err := d.setup(); err != nil {
		d.f.Close() // the device goes away with its address and route
		return nil, fmt.Errorf("cannot set up the TUN device %s: %w", c.Name, permission(err))
	}
	return d, nil
}

// check validates the configuration.
func (c Config) check() error {
	switch {
	case c.Name == "" || len(c.Name) >= syscall.IFNAMSIZ:
		return fmt.Errorf("the TUN device name %q must be 1 to %d octets", c.Name, syscall.IFNAMSIZ-1)
	case !c.Addr.Is4():
		return fmt.Errorf("the UE address %s is not IPv4", c.Addr)
	case c.MTU < minMTU || c.MTU > maxMTU:
		return fmt.Errorf("the TUN MTU %d is out of range", c.MTU)
	case c.Table == 0 || c.Table >= firstReservedTable && c.Table <= lastReservedTable:
		return fmt.Errorf("the routing table %d is reserved", c.Table)
	}
	return nil
}

// setup configures the device: no IPv6 (it would send router
// solicitations), the MTU, up, the address, the route and the rule.
func (d *Device) setup() error {
	ifi, err := net.InterfaceByName(d.cfg.Name)
	if err != nil {
		return err
	}
	d.index = ifi.Index
	os.WriteFile("/proc/sys/net/ipv6/conf/"+d.cfg.Name+"/disable_ipv6", []byte("1"), 0) // best effort
	nl, err := dialNetlink()
	if err != nil {
		return err
	}
	defer nl.close()
	if err := nl.do(syscall.RTM_NEWLINK, 0, linkUpMsg(d.index, d.cfg.MTU)); err != nil {
		return fmt.Errorf("MTU and up: %w", err)
	}
	if err := nl.do(syscall.RTM_NEWADDR, syscall.NLM_F_CREATE|nlmFReplace, addrMsg(d.index, d.cfg.Addr)); err != nil {
		return fmt.Errorf("address: %w", err)
	}
	if err := nl.do(syscall.RTM_NEWROUTE, syscall.NLM_F_CREATE|nlmFReplace, routeMsg(d.index, d.cfg.Addr, d.cfg.Table)); err != nil {
		return fmt.Errorf("route: %w", err)
	}
	// A rule left by a process that did not end cleanly is replaced.
	deleteRules(nl, d.cfg.Addr, d.cfg.Table, d.cfg.Priority)
	if err := nl.do(syscall.RTM_NEWRULE, syscall.NLM_F_CREATE|syscall.NLM_F_EXCL, ruleMsg(d.cfg.Addr, d.cfg.Table, d.cfg.Priority)); err != nil {
		return fmt.Errorf("source rule: %w", err)
	}
	return nil
}

// deleteRules removes every copy of the rule (best effort).
func deleteRules(nl *netlink, addr netip.Addr, table, prio uint32) {
	for i := 0; i < maxRuleDeletes; i++ {
		if nl.do(syscall.RTM_DELRULE, 0, ruleMsg(addr, table, prio)) != nil {
			return
		}
	}
}

// Name returns the device name.
func (d *Device) Name() string { return d.cfg.Name }

// Read reads one packet the host sent through the device.
func (d *Device) Read(b []byte) (int, error) { return d.f.Read(b) }

// Write gives the host one IP packet, as if it came in on the device.
func (d *Device) Write(b []byte) (int, error) { return d.f.Write(b) }

// Close removes the rule and the device (its address and route go with
// it). Reads in progress return an error.
func (d *Device) Close() error {
	if nl, err := dialNetlink(); err == nil {
		deleteRules(nl, d.cfg.Addr, d.cfg.Table, d.cfg.Priority)
		nl.close()
	}
	return d.f.Close()
}

// isHostAddr reports whether a is already an address of a host interface.
func isHostAddr(a netip.Addr) (bool, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return false, err
	}
	for _, x := range addrs {
		if n, ok := x.(*net.IPNet); ok {
			if h, ok := netip.AddrFromSlice(n.IP); ok && h.Unmap() == a {
				return true, nil
			}
		}
	}
	return false, nil
}

// permission explains EPERM and EACCES.
func permission(err error) error {
	if errors.Is(err, syscall.EPERM) || errors.Is(err, syscall.EACCES) {
		return fmt.Errorf("%v: %w", err, ErrPermission)
	}
	return err
}

// HasNetAdmin reports whether this process has CAP_NET_ADMIN in effect.
func HasNetAdmin() bool {
	b, err := os.ReadFile("/proc/self/status")
	if err != nil {
		return false
	}
	return capEffHas(string(b), capNetAdmin)
}

// capEffHas reads the CapEff line of /proc/<pid>/status.
func capEffHas(status string, bit uint) bool {
	for _, line := range strings.Split(status, "\n") {
		if v, ok := strings.CutPrefix(line, "CapEff:"); ok {
			caps, err := strconv.ParseUint(strings.TrimSpace(v), 16, 64)
			return err == nil && caps&(1<<bit) != 0
		}
	}
	return false
}
