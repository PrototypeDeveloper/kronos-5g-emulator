package gnb

import (
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"kronos/models"
	"kronos/scenario"
	"kronos/state"
)

func autoRole(t *testing.T, extra string) *Role {
	t.Helper()
	path := filepath.Join(t.TempDir(), "c.yaml")
	cfg := `
ue: { imsi: "001010000000001", credential: { key: 8baf473f2f8fd09487cccbd7097c6862, op: 8e27b6af0e692e750f32667a3b14605d, type: OPC, amf: "8000" }` + extra + ` }
gnodeb:
  gnodebid: "000001"
  ngap: { dst: { address: 10.0.0.1, port: 38412 } }
  gtpu: { src: { address: 10.0.0.2, port: 2152 } }
  plmn: { mcc: "001", mnc: "01" }
  snssai: [ { sst: 1 } ]
  tac: 1
`
	if err := os.WriteFile(path, []byte(cfg), 0o644); err != nil {
		t.Fatal(err)
	}
	c, ds := state.Load(path)
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	r, err := New(c)
	if err != nil {
		t.Fatal(err)
	}
	return r
}

func findEdit(n *models.EditNode, path string) *models.EditNode {
	if n.Path == path {
		return n
	}
	for _, k := range n.Kids {
		if f := findEdit(k, path); f != nil {
			return f
		}
	}
	return nil
}

// The auto-fix rules mark the fields of the sample templates that the gNB
// replaces, by send mode.
func TestAutoFields(t *testing.T) {
	sigs, ds := scenario.ReadSignals("../../../config/signal.txt")
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	r := autoRole(t, "")
	const in = "NASPDU/Value/>/"
	for _, c := range []struct {
		signal string
		mode   scenario.Mode
		path   string
		auto   bool
		note   string
	}{
		{"InitialUEMeesageRegistrationRequest", scenario.ModeAuto, "RANUENGAPID/Value", true, "the NGAP ID of the UE association"},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeAuto, "UserLocationInformation/UserLocationInformationNR/TAI/TAC/Value", true, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeAuto, "RRCEstablishmentCause/Value", false, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeAuto, in + "MobileIdentity", true, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeAuto, in + "RegistrationTypeAndNgKSI", false, "the ngKSI (high 4 bits) is replaced"},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeAuto, in + "@SecurityHeaderType", true, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeNASKeep, in + "MobileIdentity", false, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeNASKeep, in + "@SecurityHeaderType", true, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeNASKeep, "RANUENGAPID/Value", true, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeRaw, "RANUENGAPID/Value", false, ""},
		{"InitialUEMeesageRegistrationRequest", scenario.ModeRaw, in + "@SecurityHeaderType", false, ""},
		{"UpLinkNASTransportSecurityModeComplete", scenario.ModeAuto, in + "IMEISV/Value", false, "replaced by ue.imeisv when the Security Mode Command asks for it"},
		{"UpLinkNASTransportSecurityModeComplete", scenario.ModeAuto, in + "NASMessageContainer/Value/>/MobileIdentity", true, ""},
		{"UpLinkNASTransportSecurityModeComplete", scenario.ModeAuto, in + "NASMessageContainer/Value/>/RegistrationTypeAndNgKSI", false, ""},
		{"UpLinkNASTransportSecurityModeComplete", scenario.ModeAuto, in + "NASMessageContainer/Value/>/RequestedNSSAI/Value", false, ""},
		{"UpLinkNASTransportAuthenticationResponse", scenario.ModeAuto, in + "AuthenticationResponseParameter/Value", true, ""},
		{"UplinkNASTransportPDUSessionEstablishmentRequest", scenario.ModeAuto, in + "PayloadContainer/>/ProcedureTransactionID", true, ""},
		{"UplinkNASTransportPDUSessionEstablishmentRequest", scenario.ModeAuto, in + "PayloadContainer/>/SSCMode/Value", false, ""},
		{"UplinkNASTransportPDUSessionEstablishmentRequest", scenario.ModeAuto, in + "PDUSessionID/Value", true, ""},
		{"UplinkNASTransportPDUSessionEstablishmentRequest", scenario.ModeAuto, in + "DNN/Value", false, "replaced when ue.session gives one"},
		{"PDUSessionResourceSetupResponse", scenario.ModeAuto, "PDUSessionResourceSetupListSURes/List", false, "one item per PDU session answered; the first item is the pattern"},
		{"PDUSessionResourceSetupResponse", scenario.ModeAuto, "PDUSessionResourceSetupListSURes/List/0/PDUSessionID/Value", true, ""},
		{"PDUSessionResourceSetupResponse", scenario.ModeAuto,
			"PDUSessionResourceSetupListSURes/List/0/PDUSessionResourceSetupResponseTransfer/>/DLQosFlowPerTNLInformation/UPTransportLayerInformation/GTPTunnel/GTPTEID/Value", true, ""},
		{"NGSetupRequest", scenario.ModeAuto, "SupportedTAList/List/0/TAC/Value", true, ""},
		{"NGSetupRequest", scenario.ModeAuto, "RANNodeName/Value", false, ""}, // no gnodeb.name
		{"NGSetupRequest", scenario.ModeAuto, "DefaultPagingDRX/Value", false, ""},
	} {
		sig, _ := sigs.Lookup(c.signal)
		n := findEdit(scenario.EditTree(sig.Bytes, c.mode, r.AutoField), c.path)
		if n == nil || n.Auto != c.auto || (c.note != "" && n.Note != c.note) {
			t.Errorf("%s %s %s: %+v", c.signal, c.mode, c.path, n)
		}
	}
}

// A plain NAS-PDU whose spare bits are set is not taken for a protected one:
// building it reports the spare bits (libnas does not decode them).
func TestBuildNASSpareBits(t *testing.T) {
	sigs, ds := scenario.ReadSignals("../../../config/signal.txt")
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	reg, _ := sigs.Lookup("InitialUEMeesageRegistrationRequest")
	r := autoRole(t, "")
	if _, err := r.NewUE(0, "001010000000001").Build(reg); err != nil {
		t.Fatalf("the sample: %v", err)
	}
	h := hex.EncodeToString(reg.Bytes)
	set, _ := scenario.ParseSignals("x", []byte("Spare { "+strings.Replace(h, "7e004179", "7e104179", 1)+" }"))
	spare, _ := set.Lookup("Spare")
	if _, err := r.NewUE(0, "001010000000001").Build(spare); err == nil || !strings.Contains(err.Error(), "spare bits") {
		t.Fatalf("spare bits: %v", err)
	}
}
