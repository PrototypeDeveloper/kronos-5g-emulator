package scenario

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
)

func sampleSignal(t *testing.T, name string) *Signal {
	t.Helper()
	set, ds := ReadSignals("../../config/signal.txt")
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	s, ok := set.Lookup(name)
	if !ok {
		t.Fatalf("no %s", name)
	}
	return s
}

// find returns the node at path.
func find(n *models.EditNode, path string) *models.EditNode {
	if n.Path == path {
		return n
	}
	for _, k := range n.Kids {
		if f := find(k, path); f != nil {
			return f
		}
	}
	return nil
}

func edit(t *testing.T, b []byte, op models.EditOp) ([]byte, *models.EditNode) {
	t.Helper()
	out, tree, err := EditSignal(b, ModeAuto, &op, nil)
	if err != nil {
		t.Fatalf("%+v: %v", op, err)
	}
	return out, tree
}

// Every sample signal shows as a tree, and a value set to itself gives the
// same octets.
func TestEditTreeOfSamples(t *testing.T) {
	set, _ := ReadSignals("../../config/signal.txt")
	for _, s := range set.List {
		out, tree, err := EditSignal(s.Bytes, s.Mode, nil, nil)
		if err != nil || !bytes.Equal(out, s.Bytes) || tree == nil || tree.Kind != "struct" || len(tree.Kids) == 0 {
			t.Fatalf("%s: %v %+v", s.Name, err, tree)
		}
		var leaf *models.EditNode
		var first func(n *models.EditNode)
		first = func(n *models.EditNode) {
			if leaf == nil && n.Kind == "int" && !strings.Contains(n.Path, "@") {
				leaf = n
			}
			for _, k := range n.Kids {
				first(k)
			}
		}
		first(tree)
		if leaf == nil {
			continue
		}
		if again, _ := edit(t, s.Bytes, models.EditOp{Kind: "set", Path: leaf.Path, Value: leaf.Value}); !bytes.Equal(again, s.Bytes) {
			t.Fatalf("%s: setting %s to itself changed the octets", s.Name, leaf.Path)
		}
	}
}

func TestEditNGAPFields(t *testing.T) {
	b := sampleSignal(t, "InitialUEMeesageRegistrationRequest").Bytes
	// A value, a PLMN, a BIT STRING.
	out, tree := edit(t, b, models.EditOp{Kind: "set", Path: "RRCEstablishmentCause/Value", Value: "4"})
	if n := find(tree, "RRCEstablishmentCause/Value"); n.Value != "4" || n.Kind != "int" || *n.Max != 9 || !n.Ext {
		t.Fatalf("cause %+v", n)
	}
	const nr = "UserLocationInformation/UserLocationInformationNR/"
	out, _ = edit(t, out, models.EditOp{Kind: "set", Path: nr + "TAI/PLMNIdentity/Value", Value: "001/01"})
	out, tree = edit(t, out, models.EditOp{Kind: "set", Path: nr + "NRCGI/NRCellIdentity/Value", Value: "0x123456789"})
	if n := find(tree, nr+"TAI/PLMNIdentity/Value"); n.Value != "00f110" || n.Hint != "001/01" {
		t.Fatalf("PLMN %+v", n)
	}
	if n := find(tree, nr+"NRCGI/NRCellIdentity/Value"); n.Value != "0x123456789" || n.Bits != 36 {
		t.Fatalf("NCI %+v", n)
	}
	msg, _ := libngap.Decode(out)
	m := msg.(*ngapMessage.InitialUEMessage)
	if m.RRCEstablishmentCause.Value != 4 || !bytes.Equal(m.UserLocationInformation.UserLocationInformationNR.NRCGI.NRCellIdentity.Value.Bytes, []byte{0x12, 0x34, 0x56, 0x78, 0x90}) {
		t.Fatalf("decoded %+v", m)
	}

	// An optional IE: added with its smallest value, removed again.
	if n := find(tree, ""); !contains(n.Add, "AMFSetID") {
		t.Fatalf("addable %v", n.Add)
	}
	added, tree := edit(t, b, models.EditOp{Kind: "add", Path: "", Value: "AMFSetID"})
	if n := find(tree, "AMFSetID/Value"); n == nil || n.Kind != "bits" || n.Bits != 10 || n.Del != "AMFSetID" {
		t.Fatalf("added %+v", n)
	}
	if back, _ := edit(t, added, models.EditOp{Kind: "remove", Path: "AMFSetID"}); !bytes.Equal(back, b) {
		t.Fatal("adding and removing AMFSetID changed the octets")
	}

	// A CHOICE.
	_, tree = edit(t, b, models.EditOp{Kind: "choose", Path: "UserLocationInformation", Value: "UserLocationInformationEUTRA"})
	if n := find(tree, "UserLocationInformation"); n.Alt != "UserLocationInformationEUTRA" || find(tree, "UserLocationInformation/UserLocationInformationEUTRA/TAI/TAC/Value") == nil {
		t.Fatalf("choice %+v", n)
	}

	// The whole signal as hex.
	if out, _ := edit(t, b, models.EditOp{Kind: "hex", Value: hex.EncodeToString(b[:10]) + " ff"}); len(out) != 11 {
		t.Fatalf("hex %x", out)
	}
}

func TestEditNAS(t *testing.T) {
	b := sampleSignal(t, "UpLinkNASTransportSecurityModeComplete").Bytes
	const in = "NASPDU/Value/>/"
	// The security header, a value, and the message in the container.
	out, _ := edit(t, b, models.EditOp{Kind: "set", Path: in + "@MAC", Value: "01020304"})
	out, _ = edit(t, out, models.EditOp{Kind: "set", Path: in + "IMEISV/Value", Value: "4573806121856151f2"})
	out, tree := edit(t, out, models.EditOp{Kind: "set", Path: in + "NASMessageContainer/Value/>/MobileIdentity", Value: "0102f839f0ff000000000000a0"})
	out, tree = edit(t, out, models.EditOp{Kind: "add", Path: in + "NASMessageContainer/Value/>", Value: "MICOIndication"})
	if n := find(tree, in+"NASMessageContainer/Value"); n.Kind != "container" {
		t.Fatalf("container %+v", n)
	}
	msg, _ := libngap.Decode(out)
	nas := msg.(*ngapMessage.UplinkNASTransport).NASPDU.Value
	if !bytes.Equal(nas[:7], []byte{0x7e, 0x04, 1, 2, 3, 4, 0}) {
		t.Fatalf("header %x", nas[:7])
	}
	m, err := libnas.Decode(nas[7:], nil)
	if err != nil {
		t.Fatal(err)
	}
	smc := m.(*nasMessage.SecurityModeComplete)
	inner, err := libnas.Decode(smc.NASMessageContainer.Value, nil)
	if err != nil || !bytes.Equal(smc.IMEISV.Value, []byte{0x45, 0x73, 0x80, 0x61, 0x21, 0x85, 0x61, 0x51, 0xf2}) {
		t.Fatalf("SMC %+v %v", smc, err)
	}
	rr := inner.(*nasMessage.RegistrationRequest)
	if !rr.HasMICOIndication || rr.MobileIdentity[4] != 0xf0 {
		t.Fatalf("container %+v", rr)
	}

	// A plain NAS-PDU can be protected by its header type.
	reg := sampleSignal(t, "InitialUEMeesageRegistrationRequest").Bytes
	out, tree = edit(t, reg, models.EditOp{Kind: "set", Path: in + "@SecurityHeaderType", Value: "1"})
	if n := find(tree, in+"@MAC"); n == nil || n.Value != "00000000" {
		t.Fatalf("protected %+v", n)
	}
	if back, _ := edit(t, out, models.EditOp{Kind: "set", Path: in + "@SecurityHeaderType", Value: "0"}); !bytes.Equal(back, reg) {
		t.Fatal("protecting and unprotecting changed the octets")
	}

	// 5GSM in the payload container.
	est := sampleSignal(t, "UplinkNASTransportPDUSessionEstablishmentRequest").Bytes
	_, tree = edit(t, est, models.EditOp{Kind: "set", Path: in + "PayloadContainer/>/SSCMode/Value", Value: "2"})
	if n := find(tree, in+"PayloadContainer/>/SSCMode/Value"); n.Value != "2" {
		t.Fatalf("5GSM %+v", n)
	}
}

func TestEditTransferAndLists(t *testing.T) {
	b := sampleSignal(t, "PDUSessionResourceSetupResponse").Bytes
	const list = "PDUSessionResourceSetupListSURes/List"
	const g = list + "/0/PDUSessionResourceSetupResponseTransfer/>/DLQosFlowPerTNLInformation/UPTransportLayerInformation/GTPTunnel/"
	out, _ := edit(t, b, models.EditOp{Kind: "set", Path: g + "TransportLayerAddress/Value", Value: "192.0.2.1"})
	out, tree := edit(t, out, models.EditOp{Kind: "set", Path: g + "GTPTEID/Value", Value: "0000abcd"})
	if n := find(tree, g+"TransportLayerAddress/Value"); n.Kind != "ip" || n.Value != "192.0.2.1" {
		t.Fatalf("address %+v", n)
	}
	msg, _ := libngap.Decode(out)
	tr := msg.(*ngapMessage.PDUSessionResourceSetupResponse).PDUSessionResourceSetupListSURes.List[0].PDUSessionResourceSetupResponseTransfer
	var v ngapType.PDUSessionResourceSetupResponseTransfer
	if err := ngapcodec.DecodeParameter("PDUSessionResourceSetupResponseTransfer", tr, &v); ngapcodec.IsFatal(err) ||
		!bytes.Equal(v.DLQosFlowPerTNLInformation.UPTransportLayerInformation.GTPTunnel.GTPTEID.Value, []byte{0, 0, 0xab, 0xcd}) {
		t.Fatalf("transfer %+v %v", v, err)
	}

	// A list item with a minimal Transfer, then removed.
	more, tree := edit(t, b, models.EditOp{Kind: "add", Path: list})
	if n := find(tree, list); n.Hint != "2 items" || find(tree, list+"/1/PDUSessionResourceSetupResponseTransfer/>") == nil {
		t.Fatalf("list %+v", n)
	}
	if back, _ := edit(t, more, models.EditOp{Kind: "remove", Path: list + "/1"}); !bytes.Equal(back, b) {
		t.Fatal("adding and removing an item changed the octets")
	}
}

func TestEditErrors(t *testing.T) {
	b := sampleSignal(t, "InitialUEMeesageRegistrationRequest").Bytes
	setup := sampleSignal(t, "PDUSessionResourceSetupResponse").Bytes
	for _, c := range []struct {
		b    []byte
		op   models.EditOp
		want string
	}{
		{b, models.EditOp{Kind: "set", Path: "RRCEstablishmentCause/Value", Value: "-1"}, "0 or more"},
		{b, models.EditOp{Kind: "set", Path: "RANUENGAPID/Value", Value: "4294967296"}, "above 4294967295"},
		{b, models.EditOp{Kind: "set", Path: "RANUENGAPID/Value", Value: "x"}, "not an integer"},
		{b, models.EditOp{Kind: "set", Path: "UserLocationInformation", Value: "1"}, "not a value"},
		{b, models.EditOp{Kind: "set", Path: "UserLocationInformation/UserLocationInformationNR/TAI/TAC/Value", Value: "0001"}, "must have 3 octets"},
		{b, models.EditOp{Kind: "set", Path: "UserLocationInformation/UserLocationInformationEUTRA/TAI", Value: "1"}, "not the alternative chosen"},
		{b, models.EditOp{Kind: "set", Path: "AMFSetID/Value", Value: "1"}, "not present"},
		{b, models.EditOp{Kind: "set", Path: "Nothing", Value: "1"}, "no field Nothing"},
		{b, models.EditOp{Kind: "remove", Path: "RANUENGAPID"}, "mandatory"},
		{b, models.EditOp{Kind: "add", Path: "", Value: "RANUENGAPID"}, "not an optional field"},
		{b, models.EditOp{Kind: "add", Path: "", Value: "UEContextRequest"}, "present already"},
		{b, models.EditOp{Kind: "choose", Path: "UserLocationInformation", Value: "Nothing"}, "not an alternative"},
		{b, models.EditOp{Kind: "set", Path: "NASPDU/Value/>/@MAC", Value: "00"}, "set @SecurityHeaderType first"},
		{sampleSignal(t, "UplinkNASTransportPDUSessionEstablishmentRequest").Bytes, models.EditOp{Kind: "set", Path: "NASPDU/Value/>/PayloadContainer/>/@MAC", Value: "00"}, "no security header"},
		{b, models.EditOp{Kind: "set", Path: "UserLocationInformation/UserLocationInformationNR/NRCGI/NRCellIdentity/Value", Value: "0x1", Bits: 30}, "must have 36 bits"},
		{b, models.EditOp{Kind: "set", Path: "UserLocationInformation/UserLocationInformationNR/NRCGI/NRCellIdentity/Value", Value: "0x1000000000"}, "does not fit in 36 bits"},
		{setup, models.EditOp{Kind: "remove", Path: "PDUSessionResourceSetupListSURes/List/0"}, "at least 1 items"},
		{b, models.EditOp{Kind: "frob"}, "unknown edit"},
	} {
		if _, _, err := EditSignal(c.b, ModeAuto, &c.op, nil); err == nil || !strings.Contains(err.Error(), c.want) {
			t.Errorf("%+v: %v (want %q)", c.op, err, c.want)
		}
	}
}

// The role's auto-fix rule marks fields, and the fields under them.
func TestEditAutoFields(t *testing.T) {
	b := sampleSignal(t, "UpLinkNASTransportSecurityModeComplete").Bytes
	var seen []FieldRef
	rule := func(f FieldRef) (bool, string) {
		seen = append(seen, f)
		switch {
		case f.NAS == "" && f.Path == "UserLocationInformation":
			return true, ""
		case f.NAS == "RegistrationRequest" && f.Outer == "SecurityModeComplete" && f.Path == "MobileIdentity":
			return true, "the SUCI"
		case f.NAS == "SecurityModeComplete" && f.Path == "@MAC":
			return f.Mode != ModeRaw, ""
		}
		return false, ""
	}
	tree := EditTree(b, ModeAuto, rule)
	const in = "NASPDU/Value/>/"
	for path, auto := range map[string]bool{
		"UserLocationInformation": true, "UserLocationInformation/UserLocationInformationNR/TAI/TAC/Value": true,
		in + "NASMessageContainer/Value/>/MobileIdentity": true, in + "@MAC": true, "RANUENGAPID/Value": false, in + "IMEISV/Value": false,
	} {
		if n := find(tree, path); n == nil || n.Auto != auto {
			t.Errorf("%s: %+v", path, n)
		}
	}
	if n := find(tree, in+"NASMessageContainer/Value/>/MobileIdentity"); n.Note != "the SUCI" {
		t.Errorf("note %+v", n)
	}
	if n := find(EditTree(b, ModeRaw, rule), in+"@MAC"); n.Auto {
		t.Error("raw: the MAC is auto")
	}
}

func contains(list []string, s string) bool {
	for _, x := range list {
		if x == s {
			return true
		}
	}
	return false
}
