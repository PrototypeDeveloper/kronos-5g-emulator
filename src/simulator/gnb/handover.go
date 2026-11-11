package gnb

import (
	"encoding/hex"
	"fmt"
	"net/netip"
	"sort"

	libnas "libnas"

	"kronos/models"
)

// HandoverContext implements role.UE: the mapping context the target of a
// handover needs (sync_send(HO)). The target makes its own UE association;
// the AMF-UE-NGAP-ID goes along for its PathSwitchRequest (Xn handover).
func (u *UE) HandoverContext() models.HOContext {
	c := models.HOContext{Registered: u.registered, GUTI: hex.EncodeToString(u.guti), NgKSI: u.ngKSI,
		Kamf: hex.EncodeToString(u.kamf), NextPTI: u.nextPTI}
	if u.hasAMF {
		id := u.amfID
		c.AMFUENGAPID = &id
	}
	if u.sec != nil {
		c.Security = &models.HOSecurity{Integrity: u.sec.IntegrityAlgorithm, Ciphering: u.sec.CipheringAlgorithm,
			KnasInt: hex.EncodeToString(u.sec.KnasInt[:]), KnasEnc: hex.EncodeToString(u.sec.KnasEnc[:]),
			UplinkCount: uint32(u.sec.UplinkCount), DownlinkCount: uint32(u.sec.DownlinkCount)}
	}
	ids := make([]int, 0, len(u.sessions))
	for id := range u.sessions {
		ids = append(ids, int(id))
	}
	sort.Ints(ids)
	for _, id := range ids {
		s := u.sessions[uint8(id)]
		hs := models.HOSession{ID: s.id, ULTEID: s.ulTEID}
		for _, q := range s.qfis {
			hs.QFIs = append(hs.QFIs, int(q))
		}
		if s.ueIP.IsValid() {
			hs.UEIP = s.ueIP.String()
		}
		if s.upf.IsValid() {
			hs.UPF = s.upf.String()
		}
		c.Sessions = append(c.Sessions, hs)
	}
	return c
}

// AdoptHandover implements role.UE: the subscriber takes over the mapping
// context of the source (sync_receive(HO)). Its own NGAP association and
// PDU sessions are dropped; the target makes a new association.
func (u *UE) AdoptHandover(c models.HOContext) (models.Inbound, error) {
	var in models.Inbound
	guti, err := hex.DecodeString(c.GUTI)
	if err != nil || len(guti) != 0 && len(guti) != gutiLen {
		return in, fmt.Errorf("guti: want %d octets in hex (%q)", gutiLen, c.GUTI)
	}
	kamf, err := hex.DecodeString(c.Kamf)
	if err != nil {
		return in, fmt.Errorf("kamf: not hex (%q)", c.Kamf)
	}
	var sec *libnas.SecurityContext
	if s := c.Security; s != nil {
		sec = &libnas.SecurityContext{IntegrityAlgorithm: s.Integrity, CipheringAlgorithm: s.Ciphering,
			UplinkCount: libnas.Count(s.UplinkCount), DownlinkCount: libnas.Count(s.DownlinkCount)}
		if err := hexKey(sec.KnasInt[:], "knasInt", s.KnasInt); err != nil {
			return in, err
		}
		if err := hexKey(sec.KnasEnc[:], "knasEnc", s.KnasEnc); err != nil {
			return in, err
		}
	}
	sessions := map[uint8]*pdu{}
	for _, hs := range c.Sessions {
		p := &pdu{id: hs.ID, ulTEID: hs.ULTEID}
		for _, q := range hs.QFIs {
			if q < 0 || q > maxQFI {
				return in, fmt.Errorf("session %d: qfi %d is out of range (0 to %d)", hs.ID, q, maxQFI)
			}
			p.qfis = append(p.qfis, uint8(q))
		}
		if hs.UEIP != "" {
			if p.ueIP, err = netip.ParseAddr(hs.UEIP); err != nil {
				return in, fmt.Errorf("session %d: ueIP: %v", hs.ID, err)
			}
		}
		if hs.UPF != "" {
			if p.upf, err = netip.ParseAddr(hs.UPF); err != nil {
				return in, fmt.Errorf("session %d: upf: %v", hs.ID, err)
			}
		}
		sessions[hs.ID] = p
	}
	// The subscriber's own tunnels go away with its sessions.
	for _, s := range u.sessions {
		if s.dlTEID != 0 {
			in.TunnelsDown = append(in.TunnelsDown, s.dlTEID)
		}
	}
	u.ReleaseNGAP()
	u.registered, u.guti, u.ngKSI, u.kamf, u.sec, u.sessions, u.nextPTI = c.Registered, guti, c.NgKSI, kamf, sec, sessions, c.NextPTI
	u.resStar, u.imeisvRequested = nil, false
	u.sourceAMF, u.hasSourceAMF = 0, c.AMFUENGAPID != nil
	if c.AMFUENGAPID != nil {
		u.sourceAMF = *c.AMFUENGAPID
	}
	in.PagingKey = u.pagingKey()
	state := "not registered"
	if u.registered {
		state = "registered"
	}
	in.Notes = append(in.Notes, fmt.Sprintf("handover context taken over: %s, NAS security %s, %d PDU session(s)", state, onOff(sec != nil), len(sessions)))
	return in, nil
}

// hexKey decodes a 16-octet key written in hex into dst.
func hexKey(dst []byte, name, s string) error {
	b, err := hex.DecodeString(s)
	if err != nil || len(b) != len(dst) {
		return fmt.Errorf("%s: want %d octets in hex", name, len(dst))
	}
	copy(dst, b)
	return nil
}

// onOff returns "on" or "off".
func onOff(on bool) string {
	if on {
		return "on"
	}
	return "off"
}
