package gnb

import (
	"bytes"
	"encoding/binary"
	"net/netip"
	"strings"
	"testing"

	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
	"kronos/scenario"
)

// gtp is an N3 tunnel end.
func gtp(addr [4]byte, teid uint32) ngapType.UPTransportLayerInformation {
	var t ngapType.UPTransportLayerInformation
	t.Present = ngapType.UPTransportLayerInformationPresentGTPTunnel
	t.GTPTunnel.TransportLayerAddress.Value = ngapType.BitString{Bytes: addr[:], BitLength: 32}
	t.GTPTunnel.GTPTEID.Value = binary.BigEndian.AppendUint32(nil, teid)
	return t
}

// teidOf reads the TEID of a tunnel end.
func teidOf(t ngapType.UPTransportLayerInformation) uint32 {
	return binary.BigEndian.Uint32(t.GTPTunnel.GTPTEID.Value)
}

// setupReqTransfer is a PDU Session Resource Setup Request Transfer (also the
// Handover Request Transfer) with the UPF end and one QoS flow.
func setupReqTransfer(t *testing.T, ul uint32, qfi int64) []byte {
	t.Helper()
	reject := ngapType.Criticality{Value: ngapType.CriticalityPresentReject}
	var q ngapType.QosFlowSetupRequestItem
	q.QosFlowIdentifier.Value = qfi
	q.QosFlowLevelQosParameters.QosCharacteristics.Present = ngapType.QosCharacteristicsPresentNonDynamic5QI
	q.QosFlowLevelQosParameters.QosCharacteristics.NonDynamic5QI.FiveQI.Value = 9
	q.QosFlowLevelQosParameters.AllocationAndRetentionPriority.PriorityLevelARP.Value = 8
	b, err := ngapcodec.EncodeParameter("PDUSessionResourceSetupRequestTransfer", ngapType.PDUSessionResourceSetupRequestTransfer{
		ProtocolIEs: []ngapType.PDUSessionResourceSetupRequestTransferIEsEntry{
			{ID: ngapType.ProtocolIEID{Value: 139}, Criticality: reject, Present: ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentULNGUUPTNLInformation,
				ULNGUUPTNLInformation: gtp([4]byte{10, 0, 0, 9}, ul)},
			{ID: ngapType.ProtocolIEID{Value: 134}, Criticality: reject, Present: ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentPDUSessionType,
				PDUSessionType: ngapType.PDUSessionType{Value: ngapType.PDUSessionTypePresentIpv4}},
			{ID: ngapType.ProtocolIEID{Value: 136}, Criticality: reject, Present: ngapType.PDUSessionResourceSetupRequestTransferIEsEntryPresentQosFlowSetupRequestList,
				QosFlowSetupRequestList: ngapType.QosFlowSetupRequestList{List: []ngapType.QosFlowSetupRequestItem{q}}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return b
}

// targetUE is a UE that took a handover context over: registered, PDU
// session 1 with its UE IP, the source's AMF-UE-NGAP-ID 77.
func targetUE(t *testing.T) *UE {
	t.Helper()
	u := autoRole(t, "").NewUE(0, "001010000000001").(*UE)
	amf := int64(77)
	c := models.HOContext{AMFUENGAPID: &amf, Registered: true, GUTI: "f202f839cafe0000000001", NgKSI: 1,
		Sessions: []models.HOSession{{ID: 1, UEIP: "10.60.0.7", UPF: "10.0.0.9", ULTEID: 0x11, QFIs: []int{1}}}}
	if _, err := u.AdoptHandover(c); err != nil {
		t.Fatal(err)
	}
	return u
}

// buildNGAP builds the minimal message of a type, changed by edit, on u.
func buildNGAP(t *testing.T, u *UE, name string, edit func(any), off models.MapOff) any {
	t.Helper()
	b, err := scenario.NewSignal(name)
	if err != nil {
		t.Fatal(err)
	}
	msg, _ := libngap.Decode(b)
	if edit != nil {
		edit(msg)
		if b, err = libngap.Encode(msg); err != nil {
			t.Fatal(err)
		}
	}
	out, err := u.Build(&scenario.Signal{Name: name, Bytes: b, Mode: scenario.ModeAuto}, off)
	if err != nil {
		t.Fatalf("%s: %v", name, err)
	}
	m, err := libngap.Decode(out.PDU)
	if m == nil {
		t.Fatalf("%s: %v", name, err)
	}
	return m
}

func TestHandoverAtTheTarget(t *testing.T) {
	u := targetUE(t)
	// HandoverRequest: a new association, the UPF end from the source's
	// session, a DL TEID here and the tunnel up.
	var req ngapMessage.HandoverRequest
	req.AMFUENGAPID.Value = 200
	req.PDUSessionResourceSetupListHOReq.List = []ngapType.PDUSessionResourceSetupItemHOReq{{PDUSessionID: ngapType.PDUSessionID{Value: 1}, HandoverRequestTransfer: setupReqTransfer(t, 0x11, 1)}}
	if teids := scenario.HandoverULTEIDs(&req); len(teids) != 1 || teids[0] != 0x11 {
		t.Fatalf("UL TEIDs %v", teids)
	}
	in := u.Receive(&req)
	if !u.hasRAN || u.amfID != 200 || len(in.TunnelsUp) != 1 {
		t.Fatalf("association %+v, inbound %+v", u, in)
	}
	tun := in.TunnelsUp[0]
	if tun.UEIP != netip.MustParseAddr("10.60.0.7") || tun.ULTEID != 0x11 || tun.DLTEID == 0 {
		t.Fatalf("tunnel %+v", tun)
	}
	ack := buildNGAP(t, u, "HandoverRequestAcknowledge", nil, 0).(*ngapMessage.HandoverRequestAcknowledge)
	if ack.RANUENGAPID.Value != u.ranID || ack.AMFUENGAPID.Value != 200 || len(ack.PDUSessionResourceAdmittedList.List) != 1 {
		t.Fatalf("ack %+v", ack)
	}
	var tr ngapType.HandoverRequestAcknowledgeTransfer
	if err := ngapcodec.DecodeParameter("HandoverRequestAcknowledgeTransfer", ack.PDUSessionResourceAdmittedList.List[0].HandoverRequestAcknowledgeTransfer, &tr); ngapcodec.IsFatal(err) {
		t.Fatal(err)
	}
	if teidOf(tr.DLNGUUPTNLInformation) != tun.DLTEID || !bytes.Equal(tr.DLNGUUPTNLInformation.GTPTunnel.TransportLayerAddress.Value.Bytes, []byte{10, 0, 0, 2}) ||
		len(tr.QosFlowSetupResponseList.List) != 1 || tr.QosFlowSetupResponseList.List[0].QosFlowIdentifier.Value != 1 || u.pendingHO != nil {
		t.Fatalf("admitted transfer %+v", tr)
	}

	// Xn handover instead: PathSwitchRequest names the source's AMF-UE-NGAP-ID.
	u = targetUE(t)
	ps := buildNGAP(t, u, "PathSwitchRequest", nil, 0).(*ngapMessage.PathSwitchRequest)
	l := ps.PDUSessionResourceToBeSwitchedDLList.List
	if ps.SourceAMFUENGAPID.Value != 77 || ps.RANUENGAPID.Value != u.ranID || len(l) != 1 || l[0].PDUSessionID.Value != 1 ||
		ps.UESecurityCapabilities.NRencryptionAlgorithms.Value.Bytes[0] != u.r.ea<<1&nrAlgMask {
		t.Fatalf("path switch %+v", ps)
	}
	var pst ngapType.PathSwitchRequestTransfer
	ngapcodec.DecodeParameter("PathSwitchRequestTransfer", l[0].PathSwitchRequestTransfer, &pst)
	if teidOf(pst.DLNGUUPTNLInformation) != u.sessions[1].dlTEID || u.sessions[1].dlTEID == 0 {
		t.Fatalf("path switch transfer %+v", pst)
	}
	ackTr, _ := ngapcodec.EncodeParameter("PathSwitchRequestAcknowledgeTransfer", ngapType.PathSwitchRequestAcknowledgeTransfer{HasULNGUUPTNLInformation: true, ULNGUUPTNLInformation: gtp([4]byte{10, 0, 0, 10}, 0x22)})
	var psa ngapMessage.PathSwitchRequestAcknowledge
	psa.AMFUENGAPID.Value, psa.RANUENGAPID.Value = 300, u.ranID
	psa.PDUSessionResourceSwitchedList.List = []ngapType.PDUSessionResourceSwitchedItem{{PDUSessionID: ngapType.PDUSessionID{Value: 1}, PathSwitchRequestAcknowledgeTransfer: ackTr}}
	in = u.Receive(&psa)
	if u.amfID != 300 || len(in.TunnelsUp) != 1 || in.TunnelsUp[0].ULTEID != 0x22 || in.TunnelsUp[0].UPF.Addr() != netip.MustParseAddr("10.0.0.10") {
		t.Fatalf("path switched %+v", in)
	}
}

// established is a UE with an association and PDU session 1 up.
func established(t *testing.T) *UE {
	t.Helper()
	u := autoRole(t, "").NewUE(0, "001010000000001").(*UE)
	u.ranID, u.hasRAN, u.amfID, u.hasAMF = 1, true, 100, true
	s := u.session(1)
	s.ueIP, s.upf, s.ulTEID, s.dlTEID, s.qfis = netip.MustParseAddr("10.60.0.7"), netip.MustParseAddr("10.0.0.9"), 0x11, 5, []uint8{1}
	return u
}

func TestHandoverAtTheSourceAndLists(t *testing.T) {
	u := established(t)
	hr := buildNGAP(t, u, "HandoverRequired", nil, 0).(*ngapMessage.HandoverRequired)
	if l := hr.PDUSessionResourceListHORqd.List; len(l) != 1 || l[0].PDUSessionID.Value != 1 {
		t.Fatalf("HandoverRequired %+v", l)
	}
	// tunnel off: the template's list.
	hr = buildNGAP(t, u, "HandoverRequired", nil, models.MapTunnel.Of()).(*ngapMessage.HandoverRequired)
	if l := hr.PDUSessionResourceListHORqd.List; len(l) != 1 || l[0].PDUSessionID.Value != 0 {
		t.Fatalf("tunnel off %+v", l)
	}
	// Session 3 has no QFI of its own: the default flow stands for it.
	s3 := u.session(3)
	s3.ueIP, s3.dlTEID = netip.MustParseAddr("10.60.0.8"), 6
	mi := buildNGAP(t, u, "PDUSessionResourceModifyIndication", nil, 0).(*ngapMessage.PDUSessionResourceModifyIndication)
	if l := mi.PDUSessionResourceModifyListModInd.List; len(l) != 2 || l[1].PDUSessionID.Value != 3 {
		t.Fatalf("modify indication list %+v", l)
	}
	var mi3 ngapType.PDUSessionResourceModifyIndicationTransfer
	ngapcodec.DecodeParameter("PDUSessionResourceModifyIndicationTransfer", mi.PDUSessionResourceModifyListModInd.List[1].PDUSessionResourceModifyIndicationTransfer, &mi3)
	if f := mi3.DLQosFlowPerTNLInformation.AssociatedQosFlowList.List; len(f) != 1 || f[0].QosFlowIdentifier.Value != defaultQFI {
		t.Fatalf("default QoS flow %+v", f)
	}
	delete(u.sessions, 3)
	var mit ngapType.PDUSessionResourceModifyIndicationTransfer
	ngapcodec.DecodeParameter("PDUSessionResourceModifyIndicationTransfer", mi.PDUSessionResourceModifyListModInd.List[0].PDUSessionResourceModifyIndicationTransfer, &mit)
	if mi.PDUSessionResourceModifyListModInd.List[0].PDUSessionID.Value != 1 || teidOf(mit.DLQosFlowPerTNLInformation.UPTransportLayerInformation) != 5 {
		t.Fatalf("modify indication %+v", mit)
	}
	ctt := buildNGAP(t, u, "CellTrafficTrace", nil, 0).(*ngapMessage.CellTrafficTrace)
	if !bytes.Equal(ctt.NGRANCGI.NRCGI.PLMNIdentity.Value, []byte{0x00, 0xf1, 0x10}) || ctt.NGRANCGI.NRCGI.NRCellIdentity.Value.BitLength != 36 {
		t.Fatalf("cell %+v", ctt.NGRANCGI)
	}
	if m := u.Receive(&ngapMessage.UEContextModificationRequest{AMFUENGAPID: ngapType.AMFUENGAPID{Value: 100}, RANUENGAPID: ngapType.RANUENGAPID{Value: 1},
		HasNewAMFUENGAPID: true, NewAMFUENGAPID: ngapType.AMFUENGAPID{Value: 555}}); u.amfID != 555 || len(m.Notes) != 1 {
		t.Fatalf("new AMF-UE-NGAP-ID %d", u.amfID)
	}

	// Modify: a new uplink end moves the tunnel, a QoS flow is added; the response answers the session.
	reject := ngapType.Criticality{Value: ngapType.CriticalityPresentReject}
	mrt, err := ngapcodec.EncodeParameter("PDUSessionResourceModifyRequestTransfer", ngapType.PDUSessionResourceModifyRequestTransfer{ProtocolIEs: []ngapType.PDUSessionResourceModifyRequestTransferIEsEntry{
		{ID: ngapType.ProtocolIEID{Value: 140}, Criticality: reject, Present: ngapType.PDUSessionResourceModifyRequestTransferIEsEntryPresentULNGUUPTNLModifyList,
			ULNGUUPTNLModifyList: ngapType.ULNGUUPTNLModifyList{List: []ngapType.ULNGUUPTNLModifyItem{{ULNGUUPTNLInformation: gtp([4]byte{10, 0, 0, 10}, 0x33), DLNGUUPTNLInformation: gtp([4]byte{10, 0, 0, 2}, 5)}}}},
		{ID: ngapType.ProtocolIEID{Value: 135}, Criticality: reject, Present: ngapType.PDUSessionResourceModifyRequestTransferIEsEntryPresentQosFlowAddOrModifyRequestList,
			QosFlowAddOrModifyRequestList: ngapType.QosFlowAddOrModifyRequestList{List: []ngapType.QosFlowAddOrModifyRequestItem{{QosFlowIdentifier: ngapType.QosFlowIdentifier{Value: 2}}, {QosFlowIdentifier: ngapType.QosFlowIdentifier{Value: 3}}}}},
		{ID: ngapType.ProtocolIEID{Value: 137}, Criticality: reject, Present: ngapType.PDUSessionResourceModifyRequestTransferIEsEntryPresentQosFlowToReleaseList,
			QosFlowToReleaseList: ngapType.QosFlowListWithCause{List: []ngapType.QosFlowWithCauseItem{{QosFlowIdentifier: ngapType.QosFlowIdentifier{Value: 1},
				Cause: ngapType.Cause{Present: ngapType.CausePresentNas, Nas: ngapType.CauseNas{Value: ngapType.CauseNasPresentNormalRelease}}}}}},
	}})
	if err != nil {
		t.Fatal(err)
	}
	var mr ngapMessage.PDUSessionResourceModifyRequest
	mr.AMFUENGAPID.Value, mr.RANUENGAPID.Value = 555, 1
	mr.PDUSessionResourceModifyListModReq.List = []ngapType.PDUSessionResourceModifyItemModReq{{PDUSessionID: ngapType.PDUSessionID{Value: 1}, PDUSessionResourceModifyRequestTransfer: mrt}}
	in := u.Receive(&mr)
	if len(in.TunnelsDown) != 1 || in.TunnelsDown[0] != 5 || len(in.TunnelsUp) != 1 || in.TunnelsUp[0].ULTEID != 0x33 || !bytes.Equal(u.sessions[1].qfis, []uint8{2, 3}) {
		t.Fatalf("modified %+v %v", in, u.sessions[1].qfis)
	}
	resTr, _ := ngapcodec.EncodeParameter("PDUSessionResourceModifyResponseTransfer", ngapType.PDUSessionResourceModifyResponseTransfer{
		HasDLNGUUPTNLInformation: true, DLNGUUPTNLInformation: gtp([4]byte{1, 1, 1, 1}, 9),
		HasQosFlowAddOrModifyResponseList: true, QosFlowAddOrModifyResponseList: ngapType.QosFlowAddOrModifyResponseList{List: []ngapType.QosFlowAddOrModifyResponseItem{{}}}})
	res := buildNGAP(t, u, "PDUSessionResourceModifyResponse", func(m any) {
		r := m.(*ngapMessage.PDUSessionResourceModifyResponse)
		r.HasPDUSessionResourceModifyListModRes = true
		r.PDUSessionResourceModifyListModRes.List = []ngapType.PDUSessionResourceModifyItemModRes{{PDUSessionResourceModifyResponseTransfer: resTr}}
	}, 0).(*ngapMessage.PDUSessionResourceModifyResponse)
	var rt ngapType.PDUSessionResourceModifyResponseTransfer
	ngapcodec.DecodeParameter("PDUSessionResourceModifyResponseTransfer", res.PDUSessionResourceModifyListModRes.List[0].PDUSessionResourceModifyResponseTransfer, &rt)
	if res.PDUSessionResourceModifyListModRes.List[0].PDUSessionID.Value != 1 || teidOf(rt.DLNGUUPTNLInformation) != 5 || len(rt.QosFlowAddOrModifyResponseList.List) != 2 || u.pendingModify != nil {
		t.Fatalf("modify response %+v", rt)
	}

	// The requested sessions fail: InitialContextSetupFailure names them, and they are gone.
	u.pendingSetup = []uint8{1}
	icf := buildNGAP(t, u, "InitialContextSetupFailure", func(m any) {
		f := m.(*ngapMessage.InitialContextSetupFailure)
		f.HasPDUSessionResourceFailedToSetupListCxtFail = true
		f.PDUSessionResourceFailedToSetupListCxtFail.List = []ngapType.PDUSessionResourceFailedToSetupItemCxtFail{{PDUSessionID: ngapType.PDUSessionID{Value: 9}, PDUSessionResourceSetupUnsuccessfulTransfer: []byte{0x00, 0x00}}}
	}, 0).(*ngapMessage.InitialContextSetupFailure)
	if l := icf.PDUSessionResourceFailedToSetupListCxtFail.List; len(l) != 1 || l[0].PDUSessionID.Value != 1 || u.sessions[1] != nil || u.pendingSetup != nil {
		t.Fatalf("setup failure %+v", l)
	}
}

func TestNodeProcedures(t *testing.T) {
	r := autoRole(t, "")
	// A partial NG Reset names the current UE associations.
	b, _ := scenario.NewSignal("NGReset")
	msg, _ := libngap.Decode(b)
	reset := msg.(*ngapMessage.NGReset)
	reset.ResetType.Present = ngapType.ResetTypePresentPartOfNGInterface
	reset.ResetType.PartOfNGInterface.List = []ngapType.UEAssociatedLogicalNGConnectionItem{{}}
	as := []models.UEAssociation{{RANUENGAPID: 1, AMFUENGAPID: 100, HasAMF: true}, {RANUENGAPID: 2}}
	if err := r.PatchNode(reset, models.NodeBuild{Associations: as}); err != nil {
		t.Fatal(err)
	}
	if l := reset.ResetType.PartOfNGInterface.List; len(l) != 2 || l[0].AMFUENGAPID.Value != 100 || l[1].HasAMFUENGAPID || l[1].RANUENGAPID.Value != 2 {
		t.Fatalf("reset list %+v", l)
	}
	if _, err := libngap.Encode(reset); err != nil {
		t.Fatal(err)
	}
	// The acknowledgement names what the AMF's NG Reset named.
	text := r.NodeReceived(&ngapMessage.NGReset{ResetType: ngapType.ResetType{Present: ngapType.ResetTypePresentPartOfNGInterface,
		PartOfNGInterface: ngapType.UEAssociatedLogicalNGConnectionList{List: []ngapType.UEAssociatedLogicalNGConnectionItem{{HasRANUENGAPID: true, RANUENGAPID: ngapType.RANUENGAPID{Value: 5}}}}}}, &models.NodeStatus{})
	ack := &ngapMessage.NGResetAcknowledge{}
	r.PatchNode(ack, models.NodeBuild{})
	if !strings.Contains(text, "1 UE association") || !ack.HasUEAssociatedLogicalNGConnectionList || ack.UEAssociatedLogicalNGConnectionList.List[0].RANUENGAPID.Value != 5 {
		t.Fatalf("%q %+v", text, ack)
	}
	r.NodeReceived(&ngapMessage.NGReset{ResetType: ngapType.ResetType{Present: ngapType.ResetTypePresentNGInterface}}, &models.NodeStatus{})
	ack = &ngapMessage.NGResetAcknowledge{HasUEAssociatedLogicalNGConnectionList: true}
	r.PatchNode(ack, models.NodeBuild{})
	if ack.HasUEAssociatedLogicalNGConnectionList {
		t.Fatal("a whole reset is acknowledged without a list")
	}
	// RANConfigurationUpdate: the node information of the config, in the IEs it has.
	b, _ = scenario.NewSignal("RANConfigurationUpdate")
	msg, _ = libngap.Decode(b)
	rcu := msg.(*ngapMessage.RANConfigurationUpdate)
	rcu.HasSupportedTAList = true
	if err := r.PatchNode(rcu, models.NodeBuild{}); err != nil {
		t.Fatal(err)
	}
	if ta := rcu.SupportedTAList.List[0]; !bytes.Equal(ta.TAC.Value, []byte{0, 0, 1}) || !bytes.Equal(ta.BroadcastPLMNList.List[0].PLMNIdentity.Value, []byte{0x00, 0xf1, 0x10}) || rcu.HasGlobalRANNodeID {
		t.Fatalf("RANConfigurationUpdate %+v", rcu)
	}
}

// The signal editor marks the fields the new mappings fill.
func TestAutoFieldsOfProcedures(t *testing.T) {
	r := autoRole(t, "")
	for _, c := range []struct {
		ngap, nas, path string
		auto            bool
		item            string
	}{
		{"PathSwitchRequest", "", "SourceAMFUENGAPID/Value", true, "ngapid"},
		{"PathSwitchRequest", "", "PDUSessionResourceToBeSwitchedDLList/List/0/PathSwitchRequestTransfer/>/DLNGUUPTNLInformation/GTPTunnel/GTPTEID/Value", true, "tunnel"},
		{"HandoverRequired", "", "PDUSessionResourceListHORqd", false, "tunnel"},
		{"HandoverRequestAcknowledge", "", "PDUSessionResourceAdmittedList/List/0/PDUSessionID/Value", true, "tunnel"},
		{"RANConfigurationUpdate", "", "SupportedTAList/List/0/TAC/Value", true, "tac"},
		{"NGResetAcknowledge", "", "UEAssociatedLogicalNGConnectionList", true, "ngapid"},
		{"CellTrafficTrace", "", "NGRANCGI/NRCGI/PLMNIdentity/Value", true, "plmn"},
		{"UplinkNASTransport", "IdentityResponse", "MobileIdentity", true, "identity"},
		{"UplinkNASTransport", "AuthenticationFailure", "AuthenticationFailureParameter/Value", true, "res"},
		{"HandoverNotify", "", "TargetToSourceTransparentContainer", false, ""},
	} {
		auto, note := r.AutoField(scenario.FieldRef{Mode: scenario.ModeAuto, NGAP: c.ngap, NAS: c.nas, Path: c.path})
		if auto != c.auto || c.item != "" && !strings.Contains(note, "[mapping: ") || c.item != "" && !strings.Contains(note, c.item) {
			t.Errorf("%s %s %s: %v %q", c.ngap, c.nas, c.path, auto, note)
		}
	}
}
