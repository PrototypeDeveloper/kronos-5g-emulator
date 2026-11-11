// Package sctp is a small SCTP transport on the Linux kernel stack: one-to-one
// style sockets (RFC 6458) driven through the standard syscall package, with
// no cgo and no extra modules. Sockets are non-blocking and
// registered with the Go runtime poller, so Close unblocks a pending Recv.
package sctp

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"kronos/models"
	"net/netip"
	"os"
	"syscall"
	"time"
	"unsafe"
)

// withDefaults fills the zero fields with the defaults.
func (o Options) withDefaults() Options {
	if o.OutStreams == 0 {
		o.OutStreams = defaultStreams
	}
	if o.InStreams == 0 {
		o.InStreams = defaultStreams
	}
	if o.Timeout == 0 {
		o.Timeout = defaultTimeout
	}
	return o
}

// LocalAddr and RemoteAddr return the primary addresses of the association.
func (c *Conn) LocalAddr() netip.AddrPort  { return c.local }
func (c *Conn) RemoteAddr() netip.AddrPort { return c.remote }

// Streams returns the negotiated number of outbound and inbound streams.
func (c *Conn) Streams() (out, in uint16) { return c.out, c.in }

// Dial sets up an association to remote. A valid local address or a nonzero
// local port is bound first (the gNB uses a fixed source address and port).
func Dial(local, remote netip.AddrPort, opt Options) (*Conn, error) {
	opt = opt.withDefaults()
	if !remote.IsValid() {
		return nil, fmt.Errorf("sctp: invalid remote address %v", remote)
	}
	fd, err := newSocket(remote.Addr(), opt)
	if err != nil {
		return nil, err
	}
	if local.Addr().IsValid() || local.Port() != 0 {
		if err := syscall.Bind(fd, sockaddr(local, remote.Addr())); err != nil {
			syscall.Close(fd)
			return nil, fmt.Errorf("sctp: bind %v: %w", local, err)
		}
	}
	cerr := syscall.Connect(fd, sockaddr(remote, remote.Addr()))
	if cerr != nil && cerr != syscall.EINPROGRESS {
		syscall.Close(fd)
		return nil, fmt.Errorf("sctp: connect %v: %w", remote, cerr)
	}
	c, err := newConn(fd)
	if err != nil {
		return nil, err
	}
	if cerr == syscall.EINPROGRESS {
		if err := c.waitConnected(opt.Timeout); err != nil {
			c.f.Close()
			return nil, fmt.Errorf("sctp: connect %v: %w", remote, err)
		}
	}
	c.loadInfo()
	return c, nil
}

// waitConnected waits until a non-blocking connect completes or fails.
func (c *Conn) waitConnected(timeout time.Duration) error {
	if err := c.f.SetWriteDeadline(time.Now().Add(timeout)); err != nil {
		return err
	}
	defer c.f.SetWriteDeadline(time.Time{})
	var result error
	err := c.rc.Write(func(fd uintptr) bool {
		v, err := syscall.GetsockoptInt(int(fd), syscall.SOL_SOCKET, syscall.SO_ERROR)
		switch {
		case err != nil:
			result = err
		case v != 0:
			result = syscall.Errno(v)
		default:
			if _, err := syscall.Getpeername(int(fd)); err == syscall.ENOTCONN {
				return false // still setting up: wait until writable
			}
		}
		return true
	})
	if errors.Is(err, os.ErrDeadlineExceeded) {
		return fmt.Errorf("timed out after %v", timeout)
	}
	if err != nil {
		return err
	}
	return result
}

// newSocket opens a non-blocking one-to-one SCTP socket of peer's family.
func newSocket(peer netip.Addr, opt Options) (int, error) {
	domain := syscall.AF_INET6
	if peer.Unmap().Is4() {
		domain = syscall.AF_INET
	}
	fd, err := syscall.Socket(domain, syscall.SOCK_STREAM|syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC, syscall.IPPROTO_SCTP)
	if err != nil {
		return -1, fmt.Errorf("sctp: socket: %w (is the sctp kernel module loaded?)", err)
	}
	if err := setOptions(fd, opt); err != nil {
		syscall.Close(fd)
		return -1, err
	}
	return fd, nil
}

// setOptions sets the stream counts, receive info, no delay and address reuse.
func setOptions(fd int, opt Options) error {
	var init [initMsgLen]byte // struct sctp_initmsg; attempts and timeout keep the system defaults
	binary.NativeEndian.PutUint16(init[0:], opt.OutStreams)
	binary.NativeEndian.PutUint16(init[2:], opt.InStreams)
	for _, o := range []struct {
		name string
		set  func() error
	}{
		{"SCTP_INITMSG", func() error { return syscall.SetsockoptString(fd, solSCTP, sctpInitMsg, string(init[:])) }},
		{"SCTP_RECVRCVINFO", func() error { return syscall.SetsockoptInt(fd, solSCTP, sctpRecvRcvInfo, 1) }},
		{"SCTP_NODELAY", func() error { return syscall.SetsockoptInt(fd, solSCTP, sctpNoDelay, 1) }},
		{"SO_REUSEADDR", func() error { return syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1) }},
	} {
		if err := o.set(); err != nil {
			return fmt.Errorf("sctp: setsockopt %s: %w", o.name, err)
		}
	}
	return nil
}

// newConn wraps fd in a Conn registered with the runtime poller.
func newConn(fd int) (*Conn, error) {
	f := os.NewFile(uintptr(fd), "sctp")
	rc, err := f.SyscallConn()
	if err != nil {
		f.Close()
		return nil, err
	}
	return &Conn{f: f, rc: rc, buf: make([]byte, recvBufLen), oob: make([]byte, syscall.CmsgSpace(rcvInfoLen)+oobExtra)}, nil
}

// loadInfo reads the negotiated streams and the addresses of the association.
func (c *Conn) loadInfo() {
	c.rc.Control(func(fd uintptr) {
		var st [statusLen]byte // struct sctp_status
		if n, err := getsockopt(int(fd), solSCTP, sctpStatus, st[:]); err == nil && n >= 20 {
			c.in = binary.NativeEndian.Uint16(st[16:])
			c.out = binary.NativeEndian.Uint16(st[18:])
		}
		if sa, err := syscall.Getsockname(int(fd)); err == nil {
			c.local = addrPort(sa)
		}
		if sa, err := syscall.Getpeername(int(fd)); err == nil {
			c.remote = addrPort(sa)
		}
	})
}

// getsockopt reads a socket option into buf and returns its length.
func getsockopt(fd, level, name int, buf []byte) (int, error) {
	l := uint32(len(buf))
	_, _, e := syscall.Syscall6(syscall.SYS_GETSOCKOPT, uintptr(fd), uintptr(level), uintptr(name),
		uintptr(unsafe.Pointer(&buf[0])), uintptr(unsafe.Pointer(&l)), 0)
	if e != 0 {
		return 0, e
	}
	return int(l), nil
}

// Send sends one message on the given stream with the given PPID.
func (c *Conn) Send(b []byte, stream uint16, ppid uint32) error {
	oob := make([]byte, syscall.CmsgSpace(sndInfoLen))
	h := (*syscall.Cmsghdr)(unsafe.Pointer(&oob[0]))
	h.Level, h.Type = solSCTP, cmsgSndInfo
	h.SetLen(syscall.CmsgLen(sndInfoLen))
	info := oob[syscall.CmsgLen(0):]
	binary.NativeEndian.PutUint16(info[0:], stream)
	binary.BigEndian.PutUint32(info[4:], ppid) // passed to the peer as is: network byte order
	var serr error
	err := c.rc.Write(func(fd uintptr) bool {
		serr = syscall.Sendmsg(int(fd), b, oob, nil, 0)
		return serr != syscall.EAGAIN
	})
	if err != nil {
		return c.mapErr(err)
	}
	return serr
}

// mapErr reports any failure after Close as os.ErrClosed (the poller returns
// its own internal error for a closed descriptor).
func (c *Conn) mapErr(err error) error {
	if c.closed.Load() {
		return os.ErrClosed
	}
	return err
}

// Recv returns the next complete message. io.EOF means the peer shut the
// association down; os.ErrClosed means Close was called.
func (c *Conn) Recv() ([]byte, models.SCTPInfo, error) {
	// A large message comes in several reads; MSG_EOR marks its end.
	var msg []byte
	var info models.SCTPInfo
	first := true
	for {
		var n, oobn, flags int
		var rerr error
		err := c.rc.Read(func(fd uintptr) bool {
			n, oobn, flags, _, rerr = syscall.Recvmsg(int(fd), c.buf, c.oob, 0)
			return rerr != syscall.EAGAIN
		})
		if err != nil {
			return nil, info, c.mapErr(err)
		}
		if rerr != nil {
			return nil, info, rerr
		}
		if flags&msgNotification != 0 {
			continue // notifications are not subscribed; skip any that arrive
		}
		if n == 0 {
			return nil, info, io.EOF
		}
		if first {
			info = parseRcvInfo(c.oob[:oobn])
			first = false
		}
		msg = append(msg, c.buf[:n]...)
		if flags&msgEOR != 0 {
			return msg, info, nil
		}
	}
}

// parseRcvInfo reads the stream and PPID of the SCTP_RCVINFO control message.
func parseRcvInfo(oob []byte) models.SCTPInfo {
	msgs, err := syscall.ParseSocketControlMessage(oob)
	if err != nil {
		return models.SCTPInfo{}
	}
	for _, m := range msgs {
		if m.Header.Level == solSCTP && m.Header.Type == cmsgRcvInfo && len(m.Data) >= rcvInfoLen {
			return models.SCTPInfo{Stream: binary.NativeEndian.Uint16(m.Data[0:]), PPID: binary.BigEndian.Uint32(m.Data[8:])}
		}
	}
	return models.SCTPInfo{}
}

// Close shuts the association down gracefully and unblocks Recv.
func (c *Conn) Close() error {
	c.closed.Store(true)
	return c.f.Close()
}

// Listen binds local and listens. Port 0 picks a free port (see Addr).
func Listen(local netip.AddrPort, opt Options) (*Listener, error) {
	opt = opt.withDefaults()
	fd, err := newSocket(local.Addr(), opt)
	if err != nil {
		return nil, err
	}
	if err := syscall.Bind(fd, sockaddr(local, local.Addr())); err != nil {
		syscall.Close(fd)
		return nil, fmt.Errorf("sctp: bind %v: %w", local, err)
	}
	if err := syscall.Listen(fd, listenBacklog); err != nil {
		syscall.Close(fd)
		return nil, fmt.Errorf("sctp: listen %v: %w", local, err)
	}
	f := os.NewFile(uintptr(fd), "sctp-listener")
	rc, err := f.SyscallConn()
	if err != nil {
		f.Close()
		return nil, err
	}
	l := &Listener{f: f, rc: rc, opt: opt}
	rc.Control(func(fd uintptr) {
		if sa, err := syscall.Getsockname(int(fd)); err == nil {
			l.addr = addrPort(sa)
		}
	})
	return l, nil
}

// Addr returns the bound address.
func (l *Listener) Addr() netip.AddrPort { return l.addr }

// Accept waits for the next association.
func (l *Listener) Accept() (*Conn, error) {
	var nfd int
	var aerr error
	err := l.rc.Read(func(fd uintptr) bool {
		nfd, _, aerr = syscall.Accept4(int(fd), syscall.SOCK_NONBLOCK|syscall.SOCK_CLOEXEC)
		return aerr != syscall.EAGAIN
	})
	if err != nil {
		if l.closed.Load() {
			return nil, os.ErrClosed
		}
		return nil, err
	}
	if aerr != nil {
		return nil, aerr
	}
	if err := setOptions(nfd, l.opt); err != nil {
		syscall.Close(nfd)
		return nil, err
	}
	c, err := newConn(nfd)
	if err != nil {
		return nil, err
	}
	c.loadInfo()
	return c, nil
}

// Close stops listening and unblocks Accept.
func (l *Listener) Close() error {
	l.closed.Store(true)
	return l.f.Close()
}

// sockaddr converts ap for a socket whose family follows peer.
func sockaddr(ap netip.AddrPort, peer netip.Addr) syscall.Sockaddr {
	a := ap.Addr()
	if peer.Unmap().Is4() {
		if !a.IsValid() {
			a = netip.IPv4Unspecified()
		}
		return &syscall.SockaddrInet4{Port: int(ap.Port()), Addr: a.Unmap().As4()}
	}
	if !a.IsValid() {
		a = netip.IPv6Unspecified()
	}
	return &syscall.SockaddrInet6{Port: int(ap.Port()), Addr: a.As16()}
}

// addrPort converts a socket address.
func addrPort(sa syscall.Sockaddr) netip.AddrPort {
	switch s := sa.(type) {
	case *syscall.SockaddrInet4:
		return netip.AddrPortFrom(netip.AddrFrom4(s.Addr), uint16(s.Port))
	case *syscall.SockaddrInet6:
		return netip.AddrPortFrom(netip.AddrFrom16(s.Addr), uint16(s.Port))
	}
	return netip.AddrPort{}
}
