package gnb

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net/netip"
	"sort"
	"strings"

	libnas "libnas"
	"libnas/nasMessage"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
	"kronos/state"
)

// loadSessions converts the ue.session entries of config.
func loadSessions(u state.UE) ([]sessionConf, error) {
	var out []sessionConf
	for _, s := range u.Session {
		c := sessionConf{id: uint8(s.ID)}
		switch strings.ToLower(s.Type) {
		case "ipv4":
			c.pduType = pduTypeIPv4
		case "ipv6":
			c.pduType = pduTypeIPv6
		case "ipv4v6":
			c.pduType = pduTypeIPv4v6
		}
		for _, label := range strings.Split(s.DNN, ".") {
			c.dnn = append(append(c.dnn, byte(len(label))), label...)
		}
		if s.SNSSAI != nil {
			item, err := sliceItem(*s.SNSSAI)
			if err != nil {
				return nil, err
			}
			c.snssai = append([]byte{item.SNSSAI.SST.Value[0]}, item.SNSSAI.SD.Value...)
		}
		out = append(out, c)
	}
	return out, nil
}

// sessionConf picks the configured session with the template's ID, or the
// first one. Without ue.session the template values stay.
func (r *Role) sessionConf(templateID uint8) (sessionConf, bool) {
	for _, c := range r.sessions {
		if c.id == templateID {
			return c, true
		}
	}
	if len(r.sessions) > 0 {
		return r.sessions[0], true
	}
	return sessionConf{id: templateID}, false
}

// session returns the state of PDU session id, creating it if needed.
func (u *UE) session(id uint8) *pdu {
	if u.sessions == nil {
		u.sessions = map[uint8]*pdu{}
	}
	s, ok := u.sessions[id]
	if !ok {
		s = &pdu{id: id}
		u.sessions[id] = s
	}
	return s
}

// allocPTI returns the next procedure transaction identity.
func (u *UE) allocPTI() uint8 {
	u.nextPTI = u.nextPTI%maxPTI + 1 // 1..254 (0 and 255 are reserved)
	return u.nextPTI
}

// patchULNASTransport writes the session of config into the 5GSM message the
// transport carries and into its PDU session ID, S-NSSAI and DNN IEs.
func (u *UE) patchULNASTransport(t *nasMessage.ULNASTransport) error {
	if t.PayloadContainerType != payloadN1SMInfo || len(t.PayloadContainer) < 4 {
		return nil // not N1 SM information
	}
	inner, err := libnas.Decode(t.PayloadContainer, nil)
	if err != nil {
		return fmt.Errorf("cannot decode the 5GSM of ULNASTransport: %v", err)
	}
	conf, configured := u.r.sessionConf(t.PayloadContainer[1])
	id := conf.id
	// With session off the IDs and PTI are sent as written, and the
	// procedure is followed with them.
	on := u.off.On(models.MapSession)
	switch m := inner.(type) {
	case *nasMessage.PDUSessionEstablishmentRequest:
		if on {
			m.PDUSessionID, m.ProcedureTransactionID = id, u.allocPTI()
			if m.HasPDUSessionType && conf.pduType != 0 {
				m.PDUSessionType.Value = conf.pduType
			}
		}
		u.session(m.PDUSessionID).pti = m.ProcedureTransactionID
	case *nasMessage.PDUSessionReleaseRequest:
		if on {
			m.PDUSessionID, m.ProcedureTransactionID = id, u.allocPTI()
		}
		u.session(m.PDUSessionID).pti = m.ProcedureTransactionID
	case *nasMessage.PDUSessionReleaseComplete:
		if on {
			m.PDUSessionID, m.ProcedureTransactionID = id, u.session(id).relPTI
		}
	case *nasMessage.Status5GSM:
		if on {
			m.PDUSessionID = id // the PTI stays: it names the message the status is about
		}
	}
	if on {
		if t.PayloadContainer, err = libnas.Encode(inner, nil); err != nil {
			return err
		}
		if t.HasPDUSessionID {
			t.PDUSessionID.Value = id
		}
		if configured && t.HasDNN && conf.dnn != nil {
			t.DNN.Value = append([]byte(nil), conf.dnn...)
		}
	}
	if configured && t.HasSNSSAI && conf.snssai != nil && u.off.On(models.MapSNSSAI) {
		t.SNSSAI.Value = append([]byte(nil), conf.snssai...)
	}
	return nil
}

// receiveNGAP handles what the NGAP layer of a received UE message says about
// PDU sessions.
func (u *UE) receiveNGAP(msg any, in *models.Inbound) {
	switch m := msg.(type) {
	case *ngapMessage.PDUSessionResourceSetupRequest:
		u.pendingSetup = nil
		for _, it := range m.PDUSessionResourceSetupListSUReq.List {
			u.setupSession(uint8(it.PDUSessionID.Value), it.PDUSessionResourceSetupRequestTransfer, in)
		}
	case *ngapMessage.PDUSessionResourceReleaseCommand:
		u.pendingRelease = nil
		for _, it := range m.PDUSessionResourceToReleaseListRelCmd.List {
			id := uint8(it.PDUSessionID.Value)
			u.pendingRelease = append(u.pendingRelease, id)
			if s := u.sessions[id]; s != nil && s.dlTEID != 0 {
				in.TunnelsDown = append(in.TunnelsDown, s.dlTEID)
			}
			in.Notes = append(in.Notes, fmt.Sprintf("PDU session %d released (its N3 tunnel stops)", id))
		}
	case *ngapMessage.InitialContextSetupRequest:
		u.pendingSetup = nil
		if m.HasPDUSessionResourceSetupListCxtReq {
			for _, it := range m.PDUSessionResourceSetupListCxtReq.List {
				u.setupSession(uint8(it.PDUSessionID.Value), it.PDUSessionResourceSetupRequestTransfer, in)
			}
		}
	}
}

// setupSession reads the UPF side of the tunnel from the transfer and
// allocates the downlink TEID.
func (u *UE) setupSession(id uint8, transfer []byte, in *models.Inbound) {
	var tr ngapType.PDUSessionResourceSetupRequestTransfer
	if err := ngapcodec.DecodeParameter("PDUSessionResourceSetupRequestTransfer", transfer, &tr); ngapcodec.IsFatal(err) {
		in.Errors = append(in.Errors, fmt.Sprintf("cannot decode the Transfer of PDU session %d: %v", id, err))
		return
	}
	s := u.session(id)
	s.qfis = nil
	for _, ie := range tr.ProtocolIEs {
		switch ie.Present {
		case ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentULNGUUPTNLInformation:
			if ie.ULNGUUPTNLInformation.Present == ngapType.UPTransportLayerInformationPresentGTPTunnel {
				g := ie.ULNGUUPTNLInformation.GTPTunnel
				s.upf = transportAddress(g.TransportLayerAddress.Value)
				if len(g.GTPTEID.Value) == models.TEIDLen {
					s.ulTEID = binary.BigEndian.Uint32(g.GTPTEID.Value)
				}
			}
		case ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentQosFlowSetupRequestList:
			for _, q := range ie.QosFlowSetupRequestList.List {
				s.qfis = append(s.qfis, uint8(q.QosFlowIdentifier.Value))
			}
		}
	}
	if s.dlTEID == 0 {
		s.dlTEID = uint32(u.r.nextTEID.Add(1))
	}
	u.pendingSetup = append(u.pendingSetup, id)
	in.Notes = append(in.Notes, fmt.Sprintf("PDU session %d N3: UPF %s (UL TEID %#x), DL TEID %#x, QFI %v", id, s.upf, s.ulTEID, s.dlTEID, s.qfis))
	if o := u.r.upfOverride; o.IsValid() && o.Addr() != s.upf {
		in.Notes = append(in.Notes, fmt.Sprintf("the UPF address is replaced by gtpu.dst of config (%s)", o))
	}
}

// transportAddress takes the IPv4 address of an NGAP TransportLayerAddress
// (32 bits, or 160 bits for IPv4 followed by IPv6), else the IPv6 one.
func transportAddress(bs ngapType.BitString) netip.Addr {
	switch {
	case bs.BitLength >= 32 && bs.BitLength != 128 && len(bs.Bytes) >= 4:
		return netip.AddrFrom4([4]byte(bs.Bytes[:4]))
	case bs.BitLength == 128 && len(bs.Bytes) == 16:
		return netip.AddrFrom16([16]byte(bs.Bytes))
	}
	return netip.Addr{}
}

// acceptSession records what PDU Session Establishment Accept assigned.
func (u *UE) acceptSession(m *nasMessage.PDUSessionEstablishmentAccept, in *models.Inbound) {
	s := u.session(m.PDUSessionID)
	if a := m.PDUAddress.Value; m.HasPDUAddress && len(a) >= 1 {
		switch a[0] & pduTypeMask {
		case pduTypeIPv4: // IPv4
			if len(a) >= 5 {
				s.ueIP = netip.AddrFrom4([4]byte(a[1:5]))
			}
		case pduTypeIPv4v6: // IPv4v6: the IPv6 interface identifier, then IPv4
			if len(a) >= 13 {
				s.ueIP = netip.AddrFrom4([4]byte(a[9:13]))
			}
		}
	}
	in.Notes = append(in.Notes, fmt.Sprintf("PDU session %d established: UE IP %s", m.PDUSessionID, s.ueIP))
}

// tunnelsUp lists the sessions answered by the current setup whose user plane
// is complete.
func (u *UE) tunnelsUp(in *models.Inbound) { u.tunnelsFor(u.pendingSetup, in) }

// tunnelsFor lists the sessions of ids whose user plane is complete.
func (u *UE) tunnelsFor(ids []uint8, in *models.Inbound) {
	for _, id := range ids {
		s := u.sessions[id]
		if s == nil || !s.ueIP.IsValid() || !s.upf.IsValid() || s.ulTEID == 0 {
			continue
		}
		qfi := uint8(defaultQFI)
		if len(s.qfis) > 0 {
			qfi = s.qfis[0]
		}
		peer := netip.AddrPortFrom(s.upf, gtpuPort)
		if o := u.r.upfOverride; o.IsValid() {
			peer = o
		}
		in.TunnelsUp = append(in.TunnelsUp, models.Tunnel{SessionID: id, UEIP: s.ueIP, UPF: peer, ULTEID: s.ulTEID, DLTEID: s.dlTEID, QFI: qfi})
	}
}

// responseTransfer builds a PDUSessionResourceSetupResponseTransfer from the
// template's (kept for any other IEs) with the gNB's N3 address, the
// downlink TEID and the session's QoS flows.
func (r *Role) responseTransfer(template []byte, s *pdu) ([]byte, error) {
	var tr ngapType.PDUSessionResourceSetupResponseTransfer
	if len(template) > 0 {
		if err := ngapcodec.DecodeParameter("PDUSessionResourceSetupResponseTransfer", template, &tr); ngapcodec.IsFatal(err) {
			tr = ngapType.PDUSessionResourceSetupResponseTransfer{}
		}
	}
	info := &tr.DLQosFlowPerTNLInformation
	info.UPTransportLayerInformation.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel
	g := &info.UPTransportLayerInformation.GTPTunnel
	g.TransportLayerAddress.Value = ngapType.BitString{Bytes: r.n3.AsSlice(), BitLength: uint64(r.n3.BitLen())}
	g.GTPTEID.Value = binary.BigEndian.AppendUint32(nil, s.dlTEID)
	var pattern ngapType.AssociatedQosFlowItem
	if l := info.AssociatedQosFlowList.List; len(l) > 0 {
		pattern = l[0]
	}
	qfis := s.qfis
	if len(qfis) == 0 {
		qfis = []uint8{defaultQFI}
	}
	info.AssociatedQosFlowList.List = nil
	for _, q := range qfis {
		it := pattern
		it.QosFlowIdentifier.Value = int64(q)
		info.AssociatedQosFlowList.List = append(info.AssociatedQosFlowList.List, it)
	}
	tr.HasAdditionalDLQosFlowPerTNLInformation = false
	tr.HasQosFlowFailedToSetupList = false
	return ngapcodec.EncodeParameter("PDUSessionResourceSetupResponseTransfer", tr)
}

// patchSetupLists answers the sessions of the last setup request in
// PDUSessionResourceSetupResponse and InitialContextSetupResponse (one item
// per requested session, the template's first as pattern).
func (u *UE) patchSetupLists(msg any) error {
	if !u.off.On(models.MapTunnel) {
		// The lists as written; the requested sessions count as answered.
		switch msg.(type) {
		case *ngapMessage.PDUSessionResourceSetupResponse, *ngapMessage.InitialContextSetupResponse:
			u.pendingSetup = nil
		}
		return nil
	}
	switch m := msg.(type) {
	case *ngapMessage.PDUSessionResourceSetupResponse:
		if !m.HasPDUSessionResourceSetupListSURes {
			// A failure answer: the requested sessions failed and are gone.
			if m.HasPDUSessionResourceFailedToSetupListSURes && len(u.pendingSetup) > 0 {
				var pattern ngapType.PDUSessionResourceFailedToSetupItemSURes
				if l := m.PDUSessionResourceFailedToSetupListSURes.List; len(l) > 0 {
					pattern = l[0]
				}
				var items []ngapType.PDUSessionResourceFailedToSetupItemSURes
				for _, id := range u.pendingSetup {
					it := pattern
					it.PDUSessionID.Value = int64(id)
					items = append(items, it)
					delete(u.sessions, id)
				}
				m.PDUSessionResourceFailedToSetupListSURes.List = items
				u.pendingSetup = nil
			}
			return nil
		}
		if len(u.pendingSetup) == 0 {
			return errors.New("no PDU session to answer (no PDUSessionResourceSetupRequest received)")
		}
		var pattern ngapType.PDUSessionResourceSetupItemSURes
		if l := m.PDUSessionResourceSetupListSURes.List; len(l) > 0 {
			pattern = l[0]
		}
		var items []ngapType.PDUSessionResourceSetupItemSURes
		for _, id := range u.pendingSetup {
			tr, err := u.r.responseTransfer(pattern.PDUSessionResourceSetupResponseTransfer, u.sessions[id])
			if err != nil {
				return err
			}
			it := pattern
			it.PDUSessionID.Value, it.PDUSessionResourceSetupResponseTransfer = int64(id), tr
			items = append(items, it)
		}
		m.PDUSessionResourceSetupListSURes.List = items
		m.HasPDUSessionResourceFailedToSetupListSURes = false
	case *ngapMessage.InitialContextSetupResponse:
		if len(u.pendingSetup) == 0 { // no session requested: the optional list is left out
			m.HasPDUSessionResourceSetupListCxtRes = false
			return nil
		}
		var pattern ngapType.PDUSessionResourceSetupItemCxtRes
		if l := m.PDUSessionResourceSetupListCxtRes.List; m.HasPDUSessionResourceSetupListCxtRes && len(l) > 0 {
			pattern = l[0]
		}
		var items []ngapType.PDUSessionResourceSetupItemCxtRes
		for _, id := range u.pendingSetup {
			tr, err := u.r.responseTransfer(pattern.PDUSessionResourceSetupResponseTransfer, u.sessions[id])
			if err != nil {
				return err
			}
			it := pattern
			it.PDUSessionID.Value, it.PDUSessionResourceSetupResponseTransfer = int64(id), tr
			items = append(items, it)
		}
		m.HasPDUSessionResourceSetupListCxtRes = true
		m.PDUSessionResourceSetupListCxtRes.List = items
		m.HasPDUSessionResourceFailedToSetupListCxtRes = false
	default:
		return nil
	}
	u.pendingSetup = nil // answered
	return nil
}

// established returns the IDs of the PDU sessions the network has accepted, in order.
func (u *UE) established() []uint8 {
	var ids []uint8
	for id, s := range u.sessions {
		if s.ueIP.IsValid() {
			ids = append(ids, id)
		}
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
}

// patchReleaseLists fills the session lists of the release messages: the
// sessions of the last release command in PDUSessionResourceReleaseResponse,
// and the established sessions in UEContextReleaseRequest / Complete. A list
// that would be empty is left out (it is optional).
func (u *UE) patchReleaseLists(msg any) error {
	if !u.off.On(models.MapTunnel) {
		// The lists as written; the released sessions are gone all the same.
		if _, ok := msg.(*ngapMessage.PDUSessionResourceReleaseResponse); ok {
			for _, id := range u.pendingRelease {
				delete(u.sessions, id)
			}
			u.pendingRelease = nil
		}
		return nil
	}
	switch m := msg.(type) {
	case *ngapMessage.PDUSessionResourceReleaseResponse:
		if len(u.pendingRelease) == 0 {
			return errors.New("no PDU session release to answer (no PDUSessionResourceReleaseCommand received)")
		}
		var pattern ngapType.PDUSessionResourceReleasedItemRelRes
		if l := m.PDUSessionResourceReleasedListRelRes.List; len(l) > 0 {
			pattern = l[0]
		}
		if pattern.PDUSessionResourceReleaseResponseTransfer == nil {
			pattern.PDUSessionResourceReleaseResponseTransfer = []byte{0x00} // empty transfer
		}
		var items []ngapType.PDUSessionResourceReleasedItemRelRes
		for _, id := range u.pendingRelease {
			it := pattern
			it.PDUSessionID.Value = int64(id)
			items = append(items, it)
			delete(u.sessions, id)
		}
		m.PDUSessionResourceReleasedListRelRes.List = items
		u.pendingRelease = nil
	case *ngapMessage.UEContextReleaseRequest:
		if m.HasPDUSessionResourceListCxtRelReq {
			m.PDUSessionResourceListCxtRelReq.List = nil
			for _, id := range u.established() {
				m.PDUSessionResourceListCxtRelReq.List = append(m.PDUSessionResourceListCxtRelReq.List,
					ngapType.PDUSessionResourceItemCxtRelReq{PDUSessionID: ngapType.PDUSessionID{Value: int64(id)}})
			}
			m.HasPDUSessionResourceListCxtRelReq = len(m.PDUSessionResourceListCxtRelReq.List) > 0
		}
	case *ngapMessage.UEContextReleaseComplete:
		if m.HasPDUSessionResourceListCxtRelCpl {
			m.PDUSessionResourceListCxtRelCpl.List = nil
			for _, id := range u.established() {
				m.PDUSessionResourceListCxtRelCpl.List = append(m.PDUSessionResourceListCxtRelCpl.List,
					ngapType.PDUSessionResourceItemCxtRelCpl{PDUSessionID: ngapType.PDUSessionID{Value: int64(id)}})
			}
			m.HasPDUSessionResourceListCxtRelCpl = len(m.PDUSessionResourceListCxtRelCpl.List) > 0
		}
	}
	return nil
}
