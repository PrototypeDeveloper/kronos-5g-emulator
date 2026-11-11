package engine

import (
	"errors"
	"fmt"
	"net"
	"net/netip"

	libgtp "libgtp"
	"libgtp/gtpMessage"
	"libgtp/gtpType"

	"kronos/models"
	"kronos/traffic"
)

// errNoTunnel is the error of an uplink packet while the session's user
// plane is down (the UE is idle, or the session was released).
var errNoTunnel = errors.New("no N3 tunnel (the UE is idle, or has no session)")

// newUplane returns a user plane with no tunnels; bind opens the socket.
func newUplane(e *Engine) *uplane {
	return &uplane{e: e, byDL: map[uint32]dlRoute{}, bySession: map[sessionKey]models.Tunnel{}}
}

// bind routes downlink packets of the tunnel to s, opening the socket if needed.
func (u *uplane) bind(t models.Tunnel, s *subscriber) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.conn == nil {
		src := u.e.cfg.GNodeB.GTPU.Src
		if src == nil {
			return fmt.Errorf("gnodeb.gtpu.src is missing")
		}
		a, err := netip.ParseAddr(src.Address)
		if err != nil {
			return err
		}
		conn, err := net.ListenUDP("udp", net.UDPAddrFromAddrPort(netip.AddrPortFrom(a, uint16(src.Port))))
		if err != nil {
			return fmt.Errorf("cannot open the GTP-U socket: %v", err)
		}
		conn.SetReadBuffer(socketBuffer)
		conn.SetWriteBuffer(socketBuffer)
		u.conn = conn
		go u.read(conn)
		u.e.publish(models.Event{Kind: models.KindLink, Source: models.SourceNode, Text: "GTP-U: sending and receiving on " + conn.LocalAddr().String()})
	}
	u.byDL[t.DLTEID] = dlRoute{s: s, t: t}
	u.bySession[sessionKey{s, t.SessionID}] = t
	return nil
}

// unbind stops routing a downlink TEID, and its session if the TEID is still that session's tunnel.
func (u *uplane) unbind(dlTEID uint32) {
	u.mu.Lock()
	defer u.mu.Unlock()
	if r, ok := u.byDL[dlTEID]; ok {
		k := sessionKey{r.s, r.t.SessionID}
		if u.bySession[k].DLTEID == dlTEID {
			delete(u.bySession, k)
		}
	}
	delete(u.byDL, dlTEID)
}

// close drops every tunnel and the socket (the SCTP association went down).
func (u *uplane) close() {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.conn != nil {
		u.conn.Close()
		u.conn = nil
	}
	u.byDL = map[uint32]dlRoute{}
	u.bySession = map[sessionKey]models.Tunnel{}
}

// read handles the GTP-U messages received on conn until it is closed.
func (u *uplane) read(conn *net.UDPConn) {
	buf := make([]byte, packetBuffer)
	for {
		n, from, err := conn.ReadFromUDPAddrPort(buf)
		if err != nil {
			return // closed
		}
		msg, err := libgtp.Decode(buf[:n])
		if err != nil {
			continue
		}
		// G-PDUs go downlink, Echo Requests are answered and Error Indications reported.
		switch m := msg.(type) {
		case *gtpMessage.GPDU:
			u.downlink(m.Header.TEID, m.TPDU)
		case *gtpMessage.EchoRequest:
			if b, err := gtpMessage.NewEchoResponse(m.Header.SequenceNumber, gtpType.NewRecovery()).Encode(); err == nil {
				conn.WriteToUDPAddrPort(b, from)
			}
		case *gtpMessage.ErrorIndication:
			u.e.publish(models.Event{Kind: models.KindWarn, Source: models.SourceNode, Text: fmt.Sprintf("GTP-U Error Indication received (%s)", from)})
		}
	}
}

// downlink takes a G-PDU payload. An ICMP echo request to the UE is answered
// here (a ping from the DN pages an idle UE and gets its replies);
// everything else goes to the subscriber's IP stack, and what it does not
// take to the subscriber's TUN device if there is one.
func (u *uplane) downlink(teid uint32, pkt []byte) {
	u.mu.Lock()
	r, known := u.byDL[teid]
	u.mu.Unlock()
	if !known {
		return
	}
	if ip, ok := traffic.ParseIPv4(pkt); ok && ip.Proto == models.ProtoICMP && ip.Dst == r.t.UEIP {
		if e, ok := traffic.ParseEcho(ip); ok && e.Type == icmpEchoRequest {
			u.answerEcho(r, ip, e)
			return
		}
	}
	if !r.s.stack.Deliver(pkt) {
		r.s.toTUN(pkt) // with ue.tun, the host's TCP/IP gets the rest
	}
}

// answerEcho sends the echo reply of an ICMP echo request to the UE.
func (u *uplane) answerEcho(r dlRoute, ip models.IPv4, e models.Echo) {
	reply := traffic.BuildEcho(ip.Dst, ip.Src, r.s.stack.IPID(), icmpEchoReply, e.ID, e.Seq, e.Data)
	text := fmt.Sprintf("answered an ICMP Echo Request (%s → %s, id=%d seq=%d)", ip.Src, ip.Dst, e.ID, e.Seq)
	kind := models.KindInfo
	if err := u.sendUL(r.t, reply); err != nil {
		kind, text = models.KindError, fmt.Sprintf("cannot send the ICMP Echo Reply (→ %s): %v", ip.Src, err)
	}
	u.e.publish(models.Event{Kind: kind, Source: r.s.imsi, Text: text})
}

// sendFor sends an uplink IP packet through the current tunnel of a
// subscriber's PDU session (it changes when the UE comes back from idle).
func (u *uplane) sendFor(s *subscriber, session uint8, pkt []byte) error {
	u.mu.Lock()
	t, ok := u.bySession[sessionKey{s, session}]
	u.mu.Unlock()
	if !ok {
		return errNoTunnel
	}
	return u.sendUL(t, pkt)
}

// sendUL wraps an IP packet in a G-PDU for the tunnel (PDU Session
// Container: UL, QFI) and sends it to the UPF.
func (u *uplane) sendUL(t models.Tunnel, pkt []byte) error {
	g, err := gtpMessage.NewGPDUWithPDUSessionContainer(t.ULTEID, pkt, gtpType.PDUSessionInformationUL, t.QFI)
	if err != nil {
		return err
	}
	b, err := g.Encode()
	if err != nil {
		return err
	}
	u.mu.Lock()
	conn := u.conn
	u.mu.Unlock()
	if conn == nil {
		return fmt.Errorf("no GTP-U socket")
	}
	_, err = conn.WriteToUDPAddrPort(b, t.UPF)
	return err
}
