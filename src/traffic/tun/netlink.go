package tun

import (
	"encoding/binary"
	"fmt"
	"net/netip"
	"syscall"
)

// A small rtnetlink client: one request at a time, each acknowledged.

// native is the byte order of netlink messages (the host's).
var native = binary.NativeEndian

// dialNetlink opens an rtnetlink socket.
func dialNetlink() (*netlink, error) {
	fd, err := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_RAW|syscall.SOCK_CLOEXEC, syscall.NETLINK_ROUTE)
	if err != nil {
		return nil, err
	}
	if err := syscall.Bind(fd, &syscall.SockaddrNetlink{Family: syscall.AF_NETLINK}); err != nil {
		syscall.Close(fd)
		return nil, err
	}
	return &netlink{fd: fd}, nil
}

func (n *netlink) close() { syscall.Close(n.fd) }

// do sends one request and waits for its acknowledgement.
func (n *netlink) do(typ uint16, flags uint16, body []byte) error {
	n.seq++
	msg := nlMessage(typ, syscall.NLM_F_REQUEST|syscall.NLM_F_ACK|flags, n.seq, body)
	if err := syscall.Sendto(n.fd, msg, 0, &syscall.SockaddrNetlink{Family: syscall.AF_NETLINK}); err != nil {
		return err
	}
	buf := make([]byte, nlBufLen)
	for {
		k, _, err := syscall.Recvfrom(n.fd, buf, 0)
		if err != nil {
			return err
		}
		msgs, err := syscall.ParseNetlinkMessage(buf[:k])
		if err != nil {
			return err
		}
		for _, m := range msgs {
			if m.Header.Seq != n.seq || m.Header.Type != syscall.NLMSG_ERROR {
				continue
			}
			if len(m.Data) < 4 {
				return fmt.Errorf("short netlink acknowledgement")
			}
			if e := int32(native.Uint32(m.Data)); e != 0 {
				return syscall.Errno(-e)
			}
			return nil
		}
	}
}

// nlMessage puts a netlink header before body.
func nlMessage(typ, flags uint16, seq uint32, body []byte) []byte {
	b := make([]byte, syscall.NLMSG_HDRLEN, syscall.NLMSG_HDRLEN+len(body))
	native.PutUint32(b[0:], uint32(syscall.NLMSG_HDRLEN+len(body)))
	native.PutUint16(b[4:], typ)
	native.PutUint16(b[6:], flags)
	native.PutUint32(b[8:], seq)
	return append(b, body...)
}

// attr appends a route attribute (padded to 4 octets).
func attr(b []byte, typ uint16, data []byte) []byte {
	var h [4]byte
	native.PutUint16(h[0:], uint16(4+len(data)))
	native.PutUint16(h[2:], typ)
	b = append(append(b, h[:]...), data...)
	for len(b)%4 != 0 {
		b = append(b, 0)
	}
	return b
}

// u32 and ip4 encode attribute values.
func u32(v uint32) []byte {
	var b [4]byte
	native.PutUint32(b[:], v)
	return b[:]
}

func ip4(a netip.Addr) []byte { b := a.As4(); return b[:] }

// linkUpMsg is an ifinfomsg that sets the MTU and brings the link up.
func linkUpMsg(index, mtu int) []byte {
	b := make([]byte, syscall.SizeofIfInfomsg)
	b[0] = syscall.AF_UNSPEC
	native.PutUint32(b[4:], uint32(index))
	native.PutUint32(b[8:], syscall.IFF_UP)  // flags
	native.PutUint32(b[12:], syscall.IFF_UP) // change
	return attr(b, syscall.IFLA_MTU, u32(uint32(mtu)))
}

// addrMsg is an ifaddrmsg giving the device addr/32.
func addrMsg(index int, addr netip.Addr) []byte {
	b := make([]byte, syscall.SizeofIfAddrmsg)
	b[0] = syscall.AF_INET
	b[1] = hostPrefixLen // prefix length
	native.PutUint32(b[4:], uint32(index))
	b = attr(b, syscall.IFA_LOCAL, ip4(addr))
	return attr(b, syscall.IFA_ADDRESS, ip4(addr))
}

// routeMsg is an rtmsg for "default dev <index> src <addr> table <table>".
func routeMsg(index int, addr netip.Addr, table uint32) []byte {
	b := make([]byte, syscall.SizeofRtMsg)
	b[0] = syscall.AF_INET
	b[4] = syscall.RT_TABLE_UNSPEC // the table is in RTA_TABLE
	b[5] = rtprotStatic
	b[6] = rtScopeLink
	b[7] = rtnUnicast
	b = attr(b, syscall.RTA_TABLE, u32(table))
	b = attr(b, syscall.RTA_OIF, u32(uint32(index)))
	return attr(b, syscall.RTA_PREFSRC, ip4(addr))
}

// ruleMsg is a fib_rule_hdr for "from <addr> lookup <table> priority <prio>".
func ruleMsg(addr netip.Addr, table, prio uint32) []byte {
	b := make([]byte, fibRuleHdrLen) // family, dst_len, src_len, tos, table, res1, res2, action, flags(4)
	b[0] = syscall.AF_INET
	b[2] = hostPrefixLen
	b[7] = frActToTbl
	b = attr(b, fraSrc, ip4(addr))
	b = attr(b, fraTable, u32(table))
	return attr(b, fraPriority, u32(prio))
}
