package gnb

import (
	"encoding/binary"
	"fmt"

	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
)

// The NGAP procedures besides registration and session setup: handover
// (N2 and Xn), PDU session modification and notification, failures and
// trace. Their PDU session lists name the sessions the procedure is about,
// with the template's first item as the pattern.

// gtpTunnel is the gNB's end of an N3 tunnel: gtpu.src and the DL TEID.
func (r *Role) gtpTunnel(teid uint32) ngapType.UPTransportLayerInformation {
	var t ngapType.UPTransportLayerInformation
	t.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel
	t.GTPTunnel.TransportLayerAddress.Value = ngapType.BitString{Bytes: r.n3.AsSlice(), BitLength: uint64(r.n3.BitLen())}
	t.GTPTunnel.GTPTEID.Value = binary.BigEndian.AppendUint32(nil, teid)
	return t
}

// flows returns the QFIs of a session (the default one when it has none).
func (s *pdu) flows() []uint8 {
	if len(s.qfis) == 0 {
		return []uint8{defaultQFI}
	}
	return s.qfis
}

// dl returns the DL TEID of a session, allocating one the first time.
func (u *UE) dl(s *pdu) uint32 {
	if s.dlTEID == 0 {
		s.dlTEID = uint32(u.r.nextTEID.Add(1))
	}
	return s.dlTEID
}

// associated returns the associated QoS flow list of the QFIs, items like pattern.
func associated(pattern []ngapType.AssociatedQosFlowItem, qfis []uint8) []ngapType.AssociatedQosFlowItem {
	var p ngapType.AssociatedQosFlowItem
	if len(pattern) > 0 {
		p = pattern[0]
	}
	out := make([]ngapType.AssociatedQosFlowItem, 0, len(qfis))
	for _, q := range qfis {
		it := p
		it.QosFlowIdentifier.Value = int64(q)
		out = append(out, it)
	}
	return out
}

// transfer decodes the template's transfer of an item into v (left empty when there is none).
func transfer(name string, b []byte, v any) {
	if len(b) > 0 {
		ngapcodec.DecodeParameter(name, b, v)
	}
}

// patchProcedures fills the lists and identities of the other UE-associated procedures.
func (u *UE) patchProcedures(msg any) error {
	lists := u.off.On(models.MapTunnel)
	switch m := msg.(type) {
	case *ngapMessage.HandoverRequired:
		if lists {
			return u.handoverRequired(m)
		}
	case *ngapMessage.HandoverRequestAcknowledge:
		defer func() { u.pendingHO = nil }()
		if lists {
			return u.handoverAcknowledge(m)
		}
	case *ngapMessage.PathSwitchRequest:
		if u.off.On(models.MapNGAPID) && u.hasSourceAMF {
			m.SourceAMFUENGAPID.Value = u.sourceAMF
		}
		if u.off.On(models.MapSecCap) {
			// NR algorithms: one bit each from NEA1 / NIA1 on (5G-EA0 / 5G-IA0 have none).
			m.UESecurityCapabilities.NRencryptionAlgorithms.Value = ngapType.BitString{Bytes: []byte{u.r.ea << 1 & nrAlgMask, 0}, BitLength: nrAlgBits}
			m.UESecurityCapabilities.NRintegrityProtectionAlgorithms.Value = ngapType.BitString{Bytes: []byte{u.r.ia << 1 & nrAlgMask, 0}, BitLength: nrAlgBits}
		}
		if lists {
			return u.pathSwitch(m)
		}
	case *ngapMessage.PDUSessionResourceModifyResponse:
		defer func() { u.pendingModify = nil }()
		if lists && m.HasPDUSessionResourceModifyListModRes {
			return u.modifyResponse(m)
		}
	case *ngapMessage.PDUSessionResourceModifyIndication:
		if lists {
			return u.modifyIndication(m)
		}
	case *ngapMessage.PDUSessionResourceNotify:
		if lists && m.HasPDUSessionResourceNotifyList {
			var pattern ngapType.PDUSessionResourceNotifyItem
			if l := m.PDUSessionResourceNotifyList.List; len(l) > 0 {
				pattern = l[0]
			}
			var items []ngapType.PDUSessionResourceNotifyItem
			for _, id := range u.established() {
				it := pattern
				it.PDUSessionID.Value = int64(id)
				items = append(items, it)
			}
			if len(items) > 0 {
				m.PDUSessionResourceNotifyList.List = items
			}
		}
	case *ngapMessage.InitialContextSetupFailure:
		// The requested sessions failed: they are answered and gone.
		failed := u.pendingSetup
		u.pendingSetup = nil
		for _, id := range failed {
			delete(u.sessions, id)
		}
		if lists && m.HasPDUSessionResourceFailedToSetupListCxtFail && len(failed) > 0 {
			var pattern ngapType.PDUSessionResourceFailedToSetupItemCxtFail
			if l := m.PDUSessionResourceFailedToSetupListCxtFail.List; len(l) > 0 {
				pattern = l[0]
			}
			var items []ngapType.PDUSessionResourceFailedToSetupItemCxtFail
			for _, id := range failed {
				it := pattern
				it.PDUSessionID.Value = int64(id)
				items = append(items, it)
			}
			m.PDUSessionResourceFailedToSetupListCxtFail.List = items
		}
	case *ngapMessage.CellTrafficTrace:
		// The cell is this gNB's.
		if u.off.On(models.MapPLMN) || u.off.On(models.MapNCI) {
			m.NGRANCGI.Present = ngapType.NGRANCGIPresentNRCGI
		}
		if u.off.On(models.MapPLMN) {
			m.NGRANCGI.NRCGI.PLMNIdentity.Value = clone(u.r.plmn)
		}
		if u.off.On(models.MapNCI) {
			m.NGRANCGI.NRCGI.NRCellIdentity.Value = ngapType.BitString{Bytes: clone(u.r.nci.Bytes), BitLength: u.r.nci.BitLength}
		}
	}
	return nil
}

// handoverRequired lists the established sessions (at the source).
func (u *UE) handoverRequired(m *ngapMessage.HandoverRequired) error {
	var pattern ngapType.PDUSessionResourceItemHORqd
	if l := m.PDUSessionResourceListHORqd.List; len(l) > 0 {
		pattern = l[0]
	}
	if len(pattern.HandoverRequiredTransfer) == 0 {
		b, err := ngapcodec.EncodeParameter("HandoverRequiredTransfer", ngapType.HandoverRequiredTransfer{})
		if err != nil {
			return err
		}
		pattern.HandoverRequiredTransfer = b
	}
	var items []ngapType.PDUSessionResourceItemHORqd
	for _, id := range u.established() {
		it := pattern
		it.PDUSessionID.Value = int64(id)
		items = append(items, it)
	}
	if len(items) > 0 { // the list is mandatory: without a session the template's stays
		m.PDUSessionResourceListHORqd.List = items
	}
	return nil
}

// handoverAcknowledge admits the sessions of the HandoverRequest (at the
// target) with the DL TEIDs allocated for them.
func (u *UE) handoverAcknowledge(m *ngapMessage.HandoverRequestAcknowledge) error {
	var pattern ngapType.PDUSessionResourceAdmittedItem
	if l := m.PDUSessionResourceAdmittedList.List; len(l) > 0 {
		pattern = l[0]
	}
	var items []ngapType.PDUSessionResourceAdmittedItem
	for _, id := range u.pendingHO {
		s := u.sessions[id]
		if s == nil {
			continue
		}
		var tr ngapType.HandoverRequestAcknowledgeTransfer
		transfer("HandoverRequestAcknowledgeTransfer", pattern.HandoverRequestAcknowledgeTransfer, &tr)
		tr.DLNGUUPTNLInformation = u.r.gtpTunnel(u.dl(s))
		var fp ngapType.QosFlowItemWithDataForwarding
		if l := tr.QosFlowSetupResponseList.List; len(l) > 0 {
			fp = l[0]
		}
		tr.QosFlowSetupResponseList.List = nil
		for _, q := range s.flows() {
			it := fp
			it.QosFlowIdentifier.Value = int64(q)
			tr.QosFlowSetupResponseList.List = append(tr.QosFlowSetupResponseList.List, it)
		}
		tr.HasQosFlowFailedToSetupList = false
		b, err := ngapcodec.EncodeParameter("HandoverRequestAcknowledgeTransfer", tr)
		if err != nil {
			return err
		}
		it := pattern
		it.PDUSessionID.Value, it.HandoverRequestAcknowledgeTransfer = int64(id), b
		items = append(items, it)
	}
	if len(items) > 0 {
		m.PDUSessionResourceAdmittedList.List = items
		m.HasPDUSessionResourceFailedToSetupListHOAck = false
	}
	return nil
}

// pathSwitch lists the sessions to switch to this gNB (Xn handover) with their DL TEIDs.
func (u *UE) pathSwitch(m *ngapMessage.PathSwitchRequest) error {
	var pattern ngapType.PDUSessionResourceToBeSwitchedDLItem
	if l := m.PDUSessionResourceToBeSwitchedDLList.List; len(l) > 0 {
		pattern = l[0]
	}
	var items []ngapType.PDUSessionResourceToBeSwitchedDLItem
	for _, id := range u.established() {
		s := u.sessions[id]
		var tr ngapType.PathSwitchRequestTransfer
		transfer("PathSwitchRequestTransfer", pattern.PathSwitchRequestTransfer, &tr)
		tr.DLNGUUPTNLInformation = u.r.gtpTunnel(u.dl(s))
		var fp ngapType.QosFlowAcceptedItem
		if l := tr.QosFlowAcceptedList.List; len(l) > 0 {
			fp = l[0]
		}
		tr.QosFlowAcceptedList.List = nil
		for _, q := range s.flows() {
			it := fp
			it.QosFlowIdentifier.Value = int64(q)
			tr.QosFlowAcceptedList.List = append(tr.QosFlowAcceptedList.List, it)
		}
		b, err := ngapcodec.EncodeParameter("PathSwitchRequestTransfer", tr)
		if err != nil {
			return err
		}
		it := pattern
		it.PDUSessionID.Value, it.PathSwitchRequestTransfer = int64(id), b
		items = append(items, it)
	}
	if len(items) > 0 {
		m.PDUSessionResourceToBeSwitchedDLList.List = items
		m.HasPDUSessionResourceFailedToSetupListPSReq = false
	}
	return nil
}

// modifyResponse answers the sessions of the last PDUSessionResourceModifyRequest.
func (u *UE) modifyResponse(m *ngapMessage.PDUSessionResourceModifyResponse) error {
	var pattern ngapType.PDUSessionResourceModifyItemModRes
	if l := m.PDUSessionResourceModifyListModRes.List; len(l) > 0 {
		pattern = l[0]
	}
	var items []ngapType.PDUSessionResourceModifyItemModRes
	for _, id := range u.pendingModify {
		s := u.sessions[id]
		if s == nil {
			continue
		}
		var tr ngapType.PDUSessionResourceModifyResponseTransfer
		transfer("PDUSessionResourceModifyResponseTransfer", pattern.PDUSessionResourceModifyResponseTransfer, &tr)
		if tr.HasDLNGUUPTNLInformation {
			tr.DLNGUUPTNLInformation = u.r.gtpTunnel(u.dl(s))
		}
		if tr.HasQosFlowAddOrModifyResponseList {
			tr.QosFlowAddOrModifyResponseList.List = nil
			for _, q := range s.flows() {
				tr.QosFlowAddOrModifyResponseList.List = append(tr.QosFlowAddOrModifyResponseList.List, ngapType.QosFlowAddOrModifyResponseItem{QosFlowIdentifier: ngapType.QosFlowIdentifier{Value: int64(q)}})
			}
		}
		tr.HasQosFlowFailedToAddOrModifyList = false
		b, err := ngapcodec.EncodeParameter("PDUSessionResourceModifyResponseTransfer", tr)
		if err != nil {
			return err
		}
		it := pattern
		it.PDUSessionID.Value, it.PDUSessionResourceModifyResponseTransfer = int64(id), b
		items = append(items, it)
	}
	if len(items) > 0 {
		m.PDUSessionResourceModifyListModRes.List = items
		m.HasPDUSessionResourceFailedToModifyListModRes = false
	}
	return nil
}

// modifyIndication tells the AMF the gNB's N3 end of each established session.
func (u *UE) modifyIndication(m *ngapMessage.PDUSessionResourceModifyIndication) error {
	var pattern ngapType.PDUSessionResourceModifyItemModInd
	if l := m.PDUSessionResourceModifyListModInd.List; len(l) > 0 {
		pattern = l[0]
	}
	var items []ngapType.PDUSessionResourceModifyItemModInd
	for _, id := range u.established() {
		s := u.sessions[id]
		var tr ngapType.PDUSessionResourceModifyIndicationTransfer
		transfer("PDUSessionResourceModifyIndicationTransfer", pattern.PDUSessionResourceModifyIndicationTransfer, &tr)
		tr.DLQosFlowPerTNLInformation.UPTransportLayerInformation = u.r.gtpTunnel(u.dl(s))
		tr.DLQosFlowPerTNLInformation.AssociatedQosFlowList.List = associated(tr.DLQosFlowPerTNLInformation.AssociatedQosFlowList.List, s.flows())
		b, err := ngapcodec.EncodeParameter("PDUSessionResourceModifyIndicationTransfer", tr)
		if err != nil {
			return err
		}
		it := pattern
		it.PDUSessionID.Value, it.PDUSessionResourceModifyIndicationTransfer = int64(id), b
		items = append(items, it)
	}
	if len(items) > 0 {
		m.PDUSessionResourceModifyListModInd.List = items
	}
	return nil
}

// receiveProcedures handles what the other procedures tell the UE context.
func (u *UE) receiveProcedures(msg any, in *models.Inbound) {
	switch m := msg.(type) {
	case *ngapMessage.HandoverRequest:
		// A new UE association at this gNB, the target: its own RAN-UE-NGAP-ID
		// and the sessions of the source, now with DL TEIDs here.
		u.ranID, u.hasRAN = u.r.nextRAN.Add(1), true
		u.pendingSetup = nil
		for _, it := range m.PDUSessionResourceSetupListHOReq.List {
			u.setupSession(uint8(it.PDUSessionID.Value), it.HandoverRequestTransfer, in)
		}
		u.pendingHO, u.pendingSetup = u.pendingSetup, nil
		u.tunnelsFor(u.pendingHO, in)
		in.Notes = append(in.Notes, fmt.Sprintf("handover request: RAN-UE-NGAP-ID %d, %d PDU session(s) to admit", u.ranID, len(u.pendingHO)))
	case *ngapMessage.HandoverCommand:
		in.Notes = append(in.Notes, "handover command: the UE moves to the target")
	case *ngapMessage.PathSwitchRequestAcknowledge:
		// The path is switched: the UPF may name a new uplink end.
		var ids []uint8
		for _, it := range m.PDUSessionResourceSwitchedList.List {
			id := uint8(it.PDUSessionID.Value)
			s := u.sessions[id]
			if s == nil {
				continue
			}
			var tr ngapType.PathSwitchRequestAcknowledgeTransfer
			transfer("PathSwitchRequestAcknowledgeTransfer", it.PathSwitchRequestAcknowledgeTransfer, &tr)
			if tr.HasULNGUUPTNLInformation && tr.ULNGUUPTNLInformation.Present == ngapType.UPTransportLayerInformationPresentGTPTunnel {
				u.setUL(s, tr.ULNGUUPTNLInformation.GTPTunnel)
			}
			ids = append(ids, id)
		}
		u.tunnelsFor(ids, in)
		in.Notes = append(in.Notes, fmt.Sprintf("path switched: %d PDU session(s)", len(ids)))
	case *ngapMessage.UEContextModificationRequest:
		if m.HasNewAMFUENGAPID {
			u.amfID, u.hasAMF = m.NewAMFUENGAPID.Value, true
			in.Notes = append(in.Notes, fmt.Sprintf("the AMF-UE-NGAP-ID is now %d", u.amfID))
		}
	case *ngapMessage.PDUSessionResourceModifyRequest:
		u.pendingModify = nil
		for _, it := range m.PDUSessionResourceModifyListModReq.List {
			id := uint8(it.PDUSessionID.Value)
			u.pendingModify = append(u.pendingModify, id)
			u.modifySession(u.session(id), it.PDUSessionResourceModifyRequestTransfer, in)
		}
	}
}

// setUL takes the UPF's end of a tunnel.
func (u *UE) setUL(s *pdu, g ngapType.GTPTunnel) {
	s.upf = transportAddress(g.TransportLayerAddress.Value)
	if len(g.GTPTEID.Value) == models.TEIDLen {
		s.ulTEID = binary.BigEndian.Uint32(g.GTPTEID.Value)
	}
}

// modifySession applies a PDUSessionResourceModifyRequestTransfer: a new
// uplink end, QoS flows added or released. A running tunnel is set up again.
func (u *UE) modifySession(s *pdu, b []byte, in *models.Inbound) {
	var tr ngapType.PDUSessionResourceModifyRequestTransfer
	if err := ngapcodec.DecodeParameter("PDUSessionResourceModifyRequestTransfer", b, &tr); ngapcodec.IsFatal(err) {
		in.Errors = append(in.Errors, fmt.Sprintf("cannot decode the modify Transfer of PDU session %d: %v", s.id, err))
		return
	}
	moved := false
	for _, ie := range tr.ProtocolIEs {
		switch ie.Present {
		case ngapType.PDUSessionResourceModifyRequestTransferIEsEntryPresentULNGUUPTNLModifyList:
			if l := ie.ULNGUUPTNLModifyList.List; len(l) > 0 && l[0].ULNGUUPTNLInformation.Present == ngapType.UPTransportLayerInformationPresentGTPTunnel {
				u.setUL(s, l[0].ULNGUUPTNLInformation.GTPTunnel)
				moved = true
			}
		case ngapType.PDUSessionResourceModifyRequestTransferIEsEntryPresentQosFlowAddOrModifyRequestList:
			for _, q := range ie.QosFlowAddOrModifyRequestList.List {
				if qfi := uint8(q.QosFlowIdentifier.Value); !hasQFI(s.qfis, qfi) {
					s.qfis = append(s.qfis, qfi)
				}
			}
		case ngapType.PDUSessionResourceModifyRequestTransferIEsEntryPresentQosFlowToReleaseList:
			for _, q := range ie.QosFlowToReleaseList.List {
				s.qfis = removeQFI(s.qfis, uint8(q.QosFlowIdentifier.Value))
			}
		}
	}
	in.Notes = append(in.Notes, fmt.Sprintf("PDU session %d modified: UPF %s (UL TEID %#x), QFI %v", s.id, s.upf, s.ulTEID, s.qfis))
	if moved && s.dlTEID != 0 && s.ueIP.IsValid() {
		in.TunnelsDown = append(in.TunnelsDown, s.dlTEID)
		u.tunnelsFor([]uint8{s.id}, in)
	}
}

// hasQFI reports whether qfis holds q.
func hasQFI(qfis []uint8, q uint8) bool {
	for _, x := range qfis {
		if x == q {
			return true
		}
	}
	return false
}

// removeQFI returns qfis without q.
func removeQFI(qfis []uint8, q uint8) []uint8 {
	out := qfis[:0]
	for _, x := range qfis {
		if x != q {
			out = append(out, x)
		}
	}
	return out
}
