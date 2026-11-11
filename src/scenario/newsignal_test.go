package scenario

import (
	"reflect"
	"strings"
	"testing"

	libnas "libnas"
	"libnas/nasMessage"
	libngap "libngap"
	"libngap/ngapMessage"

	"kronos/models"
)

// nasMessageTypes holds the NAS messages nasTable marks as supported.
func TestNASMessageTable(t *testing.T) {
	have := map[string]bool{}
	for _, m := range nasMessageTypes {
		have[reflect.TypeOf(m).Elem().Name()] = true
	}
	for _, info := range nasTable {
		if info.Supported != have[info.Name] {
			t.Errorf("%s: supported %v, in nasMessageTypes %v", info.Name, info.Supported, have[info.Name])
		}
	}
}

func TestSignalTypes(t *testing.T) {
	ngap, gmm, gsm := SignalTypes()
	for _, want := range []string{"InitialUEMessage", "UplinkNASTransport", "NGSetupRequest", "PDUSessionResourceSetupResponse", "PDUSessionResourceSetupRequest"} {
		if !contains(ngap, want) {
			t.Errorf("no %s in %v", want, ngap)
		}
	}
	if contains(ngap, "PrivateMessage") || len(ngap) < 140 {
		t.Errorf("%d NGAP messages", len(ngap))
	}
	if len(gmm)+len(gsm) != len(nasMessageTypes) || !contains(gmm, "RegistrationRequest") || !contains(gsm, "PDUSessionEstablishmentRequest") || contains(gmm, "Status5GSM") {
		t.Errorf("5GMM %v, 5GSM %v", gmm, gsm)
	}
	// Every NAS message made is one libnas decodes again.
	for _, name := range append(append([]string{}, gmm...), gsm...) {
		b, err := NewNAS(name)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if _, err := libnas.Decode(b, nil); err != nil {
			t.Errorf("%s: %x does not decode: %v", name, b, err)
		}
	}
	if _, err := NewSignal("Nothing"); err == nil {
		t.Error("an unknown NGAP message")
	}
	if _, err := NewNAS("Nothing"); err == nil {
		t.Error("an unknown NAS message")
	}
}

// A signal made from nothing: an InitialUEMessage with a RegistrationRequest,
// an UplinkNASTransport with 5GSM in its payload container.
func TestNewSignalWithNAS(t *testing.T) {
	op := func(b []byte, o models.EditOp) ([]byte, *models.EditNode) {
		t.Helper()
		out, tree, err := EditSignal(b, ModeRaw, &o, nil)
		if err != nil {
			t.Fatalf("%+v: %v", o, err)
		}
		return out, tree
	}
	b, tree := op(nil, models.EditOp{Kind: "new", Value: "InitialUEMessage"})
	if tree.Name != "InitialUEMessage" || find(tree, "NASPDU/Value").NAS != "5gmm" {
		t.Fatalf("new %+v", tree)
	}
	b, tree = op(b, models.EditOp{Kind: "nas", Path: "NASPDU/Value", Value: "RegistrationRequest"})
	if n := find(tree, "NASPDU/Value/>"); n == nil || n.Name != "RegistrationRequest" {
		t.Fatalf("NAS %+v", tree)
	}
	b, _ = op(b, models.EditOp{Kind: "set", Path: "NASPDU/Value/>/RegistrationTypeAndNgKSI", Value: "0x71"})
	b, _ = op(b, models.EditOp{Kind: "set", Path: "RANUENGAPID/Value", Value: "5"})
	msg, _ := libngap.Decode(b)
	m := msg.(*ngapMessage.InitialUEMessage)
	nas, err := libnas.Decode(m.NASPDU.Value, nil)
	if err != nil || m.RANUENGAPID.Value != 5 || nas.(*nasMessage.RegistrationRequest).RegistrationTypeAndNgKSI != 0x71 {
		t.Fatalf("decoded %+v %+v %v", m, nas, err)
	}

	b, tree = op(nil, models.EditOp{Kind: "new", Value: "UplinkNASTransport"})
	b, tree = op(b, models.EditOp{Kind: "nas", Path: "NASPDU/Value", Value: "ULNASTransport"})
	if n := find(tree, "NASPDU/Value/>/PayloadContainer"); n == nil || n.NAS != "5gsm" {
		t.Fatalf("payload container %+v", n)
	}
	b, tree = op(b, models.EditOp{Kind: "nas", Path: "NASPDU/Value/>/PayloadContainer", Value: "PDUSessionEstablishmentRequest"})
	if n := find(tree, "NASPDU/Value/>/PayloadContainer/>"); n == nil || n.Name != "PDUSessionEstablishmentRequest" {
		t.Fatalf("5GSM %+v", n)
	}
	if _, _, err := EditSignal(b, ModeRaw, &models.EditOp{Kind: "nas", Path: "RANUENGAPID/Value", Value: "RegistrationRequest"}, nil); err == nil ||
		!strings.Contains(err.Error(), "only a NAS-PDU or a NAS container") {
		t.Fatalf("NAS outside a container: %v", err)
	}
}
