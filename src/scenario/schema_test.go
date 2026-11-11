package scenario

import (
	"strings"
	"testing"

	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"
	"libngap/ngapcodec"

	"kronos/models"
)

func editOp(t *testing.T, b []byte, o models.EditOp) ([]byte, *models.EditNode) {
	t.Helper()
	out, tree, err := EditSignal(b, ModeRaw, &o, nil)
	if err != nil {
		t.Fatalf("%+v: %v", o, err)
	}
	return out, tree
}

// ENUMERATED values come with their names from the libngap schema; NAS
// fields, whose types share names with NGAP ones, get none.
func TestEnumNames(t *testing.T) {
	b := sampleSignal(t, "InitialUEMeesageRegistrationRequest").Bytes
	tree := EditTree(b, ModeAuto, nil)
	n := find(tree, "RRCEstablishmentCause/Value")
	if n == nil || len(n.Enum) != 12 || n.Enum[3] != (models.EnumOption{Value: 3, Name: "mo-Signalling"}) || !n.Enum[10].Ext || n.Enum[10].Value != 10 {
		t.Fatalf("RRCEstablishmentCause %+v", n)
	}
	est := sampleSignal(t, "UplinkNASTransportPDUSessionEstablishmentRequest").Bytes
	if n := find(EditTree(est, ModeAuto, nil), "NASPDU/Value/>/PayloadContainer/>/PDUSessionType/Value"); n == nil || n.Enum != nil {
		t.Fatalf("NAS PDUSessionType %+v", n)
	}
	// A cause: the enumerated value of a CHOICE alternative.
	rel := sampleSignal(t, "AN_UEContextReleaseRequest").Bytes
	if n := find(EditTree(rel, ModeAuto, nil), "Cause/RadioNetwork/Value"); n == nil || len(n.Enum) < 40 || n.Enum[0].Name != "unspecified" {
		t.Fatalf("cause %+v", n)
	}
}

// An IE container (a Transfer of protocol IEs) is made with its mandatory
// IEs, and its entries are edited: the IE ID follows the IE chosen.
func TestIEContainer(t *testing.T) {
	b, err := NewSignal("PDUSessionResourceSetupRequest")
	if err != nil {
		t.Fatal(err)
	}
	b, tree := editOp(t, b, models.EditOp{Kind: "add", Path: "PDUSessionResourceSetupListSUReq/List"})
	const tr = "PDUSessionResourceSetupListSUReq/List/0/PDUSessionResourceSetupRequestTransfer"
	if n := find(tree, tr+"/>/ProtocolIEs"); n == nil || n.Kind != "list" || len(n.Kids) != 3 || n.Kids[0].Kind != "choice" || n.Kids[0].Hint != "id 139, reject" {
		t.Fatalf("container %+v", n)
	}
	msg, _ := libngap.Decode(b)
	raw := msg.(*ngapMessage.PDUSessionResourceSetupRequest).PDUSessionResourceSetupListSUReq.List[0].PDUSessionResourceSetupRequestTransfer
	var v ngapType.PDUSessionResourceSetupRequestTransfer
	if err := ngapcodec.DecodeParameter("PDUSessionResourceSetupRequestTransfer", raw, &v); ngapcodec.IsFatal(err) || len(v.ProtocolIEs) != 3 ||
		v.ProtocolIEs[0].ID.Value != 139 || v.ProtocolIEs[1].ID.Value != 134 || v.ProtocolIEs[2].ID.Value != 136 {
		t.Fatalf("transfer %+v %v", v, err)
	}
	// One more entry, then another IE in it.
	b, tree = editOp(t, b, models.EditOp{Kind: "add", Path: tr + "/>/ProtocolIEs"})
	if n := find(tree, tr+"/>/ProtocolIEs"); len(n.Kids) != 4 || n.Kids[3].Del == "" {
		t.Fatalf("added %+v", n)
	}
	_, tree = editOp(t, b, models.EditOp{Kind: "choose", Path: tr + "/>/ProtocolIEs/3", Value: "SecurityIndication"})
	if n := find(tree, tr+"/>/ProtocolIEs/3"); n.Alt != "SecurityIndication" || n.Hint != "id 138, reject" {
		t.Fatalf("chosen %+v", n)
	}
	// A mandatory IE cannot be replaced: the encoder refuses the message.
	if _, _, err := EditSignal(b, ModeRaw, &models.EditOp{Kind: "choose", Path: tr + "/>/ProtocolIEs/1", Value: "SecurityIndication"}, nil); err == nil ||
		!strings.Contains(err.Error(), "mandatory IE PDUSessionType (134) missing") {
		t.Fatalf("replacing a mandatory IE: %v", err)
	}
}

// The extension alternative of a CHOICE and extension containers are edited.
func TestExtensions(t *testing.T) {
	b := sampleSignal(t, "InitialUEMeesageRegistrationRequest").Bytes
	tree := EditTree(b, ModeRaw, nil)
	if n := find(tree, "UserLocationInformation"); !contains(n.Alts, "ChoiceExtensions") {
		t.Fatalf("alternatives %v", n.Alts)
	}
	_, tree = editOp(t, b, models.EditOp{Kind: "choose", Path: "UserLocationInformation", Value: "ChoiceExtensions"})
	if n := find(tree, "UserLocationInformation/ChoiceExtensions"); n == nil || n.Alt != "UserLocationInformationTNGF" || n.Hint != "id 244, ignore" {
		t.Fatalf("choice extension %+v", n)
	}
	const nr = "UserLocationInformation/UserLocationInformationNR"
	if n := find(tree, nr); n != nil {
		t.Fatalf("NR still there %+v", n)
	}
	nrNode := find(EditTree(b, ModeRaw, nil), nr)
	if !contains(nrNode.Add, "IEExtensions") {
		t.Fatalf("NR extensions not addable: %v", nrNode.Add)
	}
	b2, tree := editOp(t, b, models.EditOp{Kind: "add", Path: nr, Value: "IEExtensions"})
	if n := find(tree, nr+"/IEExtensions"); n == nil || n.Kind != "list" || len(n.Kids) != 1 || n.Del != nr+"/IEExtensions" {
		t.Fatalf("NR extensions %+v", n)
	}
	if _, err := libngap.Decode(b2); err != nil && !strings.Contains(err.Error(), "notify") {
		t.Fatalf("decode with extensions: %v", err)
	}
	// An extension container of IEs libngap does not know stays read-only.
	sn := find(EditTree(sampleSignal(t, "NGSetupRequest").Bytes, ModeRaw, nil), "SupportedTAList/List/0/BroadcastPLMNList/List/0/TAISliceSupportList/List/0/SNSSAI")
	if contains(sn.Add, "IEExtensions") {
		t.Fatalf("S-NSSAI extensions addable: %v", sn.Add)
	}
}
