package gnb

import (
	"bytes"
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	libngap "libngap"
	"libngap/ngapMessage"

	"kronos/models"
	"kronos/scenario"
)

// With every item off, the signals go as signal.txt has them.
func TestMapOffAll(t *testing.T) {
	sigs, ds := scenario.ReadSignals("../../../config/signal.txt")
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	r := autoRole(t, "")
	setup, _ := sigs.Lookup("NGSetupRequest")
	msg, _ := libngap.Decode(setup.Bytes)
	if err := r.PatchNode(msg, models.NodeBuild{Off: models.MapAll}); err != nil {
		t.Fatal(err)
	}
	if b, err := libngap.Encode(msg); err != nil || !bytes.Equal(b, setup.Bytes) {
		t.Fatalf("NGSetupRequest changed: %x %v", b, err)
	}
	reg, _ := sigs.Lookup("InitialUEMeesageRegistrationRequest")
	u := r.NewUE(0, "001010000000001").(*UE)
	built, err := u.Build(reg, models.MapAll)
	if err != nil || !bytes.Equal(built.PDU, reg.Bytes) {
		t.Fatalf("InitialUEMessage changed: %x %v", built.PDU, err)
	}
	// The template's RAN-UE-NGAP-ID is the one of the association.
	tmpl, _ := libngap.Decode(reg.Bytes)
	if id, ok := u.RANID(); !ok || id != tmpl.(*ngapMessage.InitialUEMessage).RANUENGAPID.Value {
		t.Fatalf("RAN-UE-NGAP-ID %d %v", id, ok)
	}
}

// One item off: its fields keep the template's values, the others are fixed.
func TestMapOffItems(t *testing.T) {
	sigs, _ := scenario.ReadSignals("../../../config/signal.txt")
	r := autoRole(t, "")
	setup, _ := sigs.Lookup("NGSetupRequest")
	tmpl, _ := libngap.Decode(setup.Bytes)
	msg, _ := libngap.Decode(setup.Bytes)
	if err := r.PatchNode(msg, models.NodeBuild{Off: models.MapPLMN.Of()}); err != nil {
		t.Fatal(err)
	}
	m, tm := msg.(*ngapMessage.NGSetupRequest), tmpl.(*ngapMessage.NGSetupRequest)
	ta, tta := m.SupportedTAList.List[0], tm.SupportedTAList.List[0]
	switch {
	case !bytes.Equal(m.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value, tm.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value):
		t.Fatal("plmn off: the PLMN of GlobalRANNodeID changed")
	case !bytes.Equal(ta.BroadcastPLMNList.List[0].PLMNIdentity.Value, tta.BroadcastPLMNList.List[0].PLMNIdentity.Value):
		t.Fatal("plmn off: the broadcast PLMN changed")
	case !bytes.Equal(ta.TAC.Value, []byte{0, 0, 1}):
		t.Fatalf("tac on: TAC %x", ta.TAC.Value)
	}

	reg, _ := sigs.Lookup("InitialUEMeesageRegistrationRequest")
	treg, _ := libngap.Decode(reg.Bytes)
	tnas, _ := libnas.Decode(treg.(*ngapMessage.InitialUEMessage).NASPDU.Value, nil)
	u := r.NewUE(0, "001010000000001").(*UE)
	built, err := u.Build(reg, models.MapIdentity.Of()|models.MapNCI.Of())
	if err != nil {
		t.Fatal(err)
	}
	out, _ := libngap.Decode(built.PDU)
	im := out.(*ngapMessage.InitialUEMessage)
	nas, _ := libnas.Decode(im.NASPDU.Value, nil)
	nr, tnr := im.UserLocationInformation.UserLocationInformationNR, treg.(*ngapMessage.InitialUEMessage).UserLocationInformation.UserLocationInformationNR
	switch {
	case !bytes.Equal(nas.(*nasMessage.RegistrationRequest).MobileIdentity, tnas.(*nasMessage.RegistrationRequest).MobileIdentity):
		t.Fatal("identity off: the SUCI changed")
	case !bytes.Equal(nr.NRCGI.NRCellIdentity.Value.Bytes, tnr.NRCGI.NRCellIdentity.Value.Bytes):
		t.Fatal("nci off: the cell changed")
	case !bytes.Equal(nr.TAI.PLMNIdentity.Value, []byte{0x00, 0xf1, 0x10}):
		t.Fatalf("plmn on: TAI PLMN %x", nr.TAI.PLMNIdentity.Value)
	case im.RANUENGAPID.Value != u.ranID:
		t.Fatal("ngapid on: the allocated RAN-UE-NGAP-ID is not sent")
	}
}

// A handover context goes out and comes in; a broken one is refused whole.
func TestHandoverContext(t *testing.T) {
	r := autoRole(t, "")
	src := r.NewUE(0, "001010000000001").(*UE)
	src.registered, src.guti, src.ngKSI, src.kamf = true, []byte{0xf2, 0x02, 0xf8, 0x39, 0xca, 0xfe, 0, 0, 0, 0, 1}, 1, bytes.Repeat([]byte{7}, 32)
	src.sec = &libnas.SecurityContext{IntegrityAlgorithm: 2, CipheringAlgorithm: 2, UplinkCount: 5, DownlinkCount: 3}
	src.sec.KnasInt[0], src.sec.KnasEnc[0] = 0xaa, 0xbb
	src.session(1).qfis = []uint8{1, 9}
	c := src.HandoverContext()
	dst := r.NewUE(0, "001010000000001").(*UE)
	dst.ranID, dst.hasRAN = 9, true
	in, err := dst.AdoptHandover(c)
	if err != nil {
		t.Fatal(err)
	}
	switch {
	case !dst.registered || !bytes.Equal(dst.guti, src.guti) || dst.ngKSI != 1 || !bytes.Equal(dst.kamf, src.kamf):
		t.Fatalf("identity %+v", dst)
	case *dst.sec != *src.sec:
		t.Fatalf("security %+v", dst.sec)
	case dst.hasRAN || len(dst.sessions) != 1 || !bytes.Equal(dst.sessions[1].qfis, []uint8{1, 9}):
		t.Fatalf("association and sessions %+v", dst)
	case in.PagingKey == "" || len(in.Notes) != 1:
		t.Fatalf("inbound %+v", in)
	}
	for name, bad := range map[string]func(c *models.HOContext){
		"guti":    func(c *models.HOContext) { c.GUTI = "f2" },
		"knasInt": func(c *models.HOContext) { s := *c.Security; s.KnasInt = "zz"; c.Security = &s },
		"ueIP":    func(c *models.HOContext) { c.Sessions = []models.HOSession{{ID: 1, UEIP: "x"}} },
		"qfi":     func(c *models.HOContext) { c.Sessions = []models.HOSession{{ID: 1, QFIs: []int{64}}} },
	} {
		cc := src.HandoverContext()
		bad(&cc)
		u := r.NewUE(0, "001010000000001").(*UE)
		if _, err := u.AdoptHandover(cc); err == nil || !strings.Contains(err.Error(), name) || u.registered {
			t.Errorf("%s: %v", name, err)
		}
	}
}
