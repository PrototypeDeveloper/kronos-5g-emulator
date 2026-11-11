package gnb

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/models"
	"kronos/scenario"
	"kronos/state"
)

func TestBitString(t *testing.T) {
	for _, tc := range []struct {
		v    uint64
		n    int
		want []byte
	}{{1, 32, []byte{0, 0, 0, 1}}, {1, 22, []byte{0, 0, 4}}, {0x3f8, 10, []byte{0xfe, 0}}, {0xca, 8, []byte{0xca}}} {
		bs := BitString(tc.v, tc.n)
		if !bytes.Equal(bs.Bytes, tc.want) || bs.BitLength != uint64(tc.n) || BitValue(bs) != tc.v {
			t.Fatalf("%x/%d: %x", tc.v, tc.n, bs.Bytes)
		}
	}
}

// The sample NGSetupRequest (captured from UERANSIM) is rewritten from config.
func TestPatchSampleNGSetupRequest(t *testing.T) {
	sigs, ds := scenario.ReadSignals("../../../config/signal.txt")
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	sig, _ := sigs.Lookup("NGSetupRequest")
	path := filepath.Join(t.TempDir(), "c.yaml")
	cfg := `
ue: { imsi: "001010000000001", credential: { key: 8baf473f2f8fd09487cccbd7097c6862, op: 8e27b6af0e692e750f32667a3b14605d, type: OPC, amf: "8000" } }
gnodeb:
  gnodebid: "000001"
  ngap: { dst: { address: 10.0.0.1, port: 38412 } }
  gtpu: { src: { address: 10.0.0.2, port: 2152 } }
  plmn: { mcc: "001", mnc: "01" }
  snssai: [ { sst: 1 } ]
  tac: 0x000102
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
	msg, _ := libngap.Decode(sig.Bytes)
	name := msg.(*ngapMessage.NGSetupRequest).RANNodeName.Value
	if err := r.PatchNode(msg, models.NodeBuild{}); err != nil {
		t.Fatal(err)
	}
	b, err := libngap.Encode(msg)
	if err != nil {
		t.Fatal(err)
	}
	back, _ := libngap.Decode(b)
	m := back.(*ngapMessage.NGSetupRequest)
	if got := m.GlobalRANNodeID.GlobalGNBID.PLMNIdentity.Value; !bytes.Equal(got, []byte{0x00, 0xf1, 0x10}) {
		t.Fatalf("PLMN %x", got)
	}
	// gnodeb.name is empty, so the template's RANNodeName stays.
	if m.RANNodeName.Value != name || name == "" {
		t.Fatalf("RANNodeName %q (template %q)", m.RANNodeName.Value, name)
	}
	ta := m.SupportedTAList.List[0]
	if !bytes.Equal(ta.TAC.Value, []byte{0, 1, 2}) || ta.BroadcastPLMNList.List[0].TAISliceSupportList.List[0].SNSSAI.HasSD {
		t.Fatalf("TA %+v", ta)
	}
	if m.DefaultPagingDRX.Value != msg.(*ngapMessage.NGSetupRequest).DefaultPagingDRX.Value {
		t.Fatal("unrelated IE changed")
	}
}

func TestNodeReceivedFailure(t *testing.T) {
	r := &Role{}
	var st models.NodeStatus
	text := r.NodeReceived(&ngapMessage.NGSetupFailure{Cause: ngapType.Cause{Present: ngapType.CausePresentProtocol, Protocol: ngapType.CauseProtocol{Value: 4}}}, &st)
	if st.Setup != models.SetupFailed || !strings.Contains(text, "protocol(4) semantic-error") {
		t.Fatalf("%q %+v", text, st)
	}
	if r.NodeReceived(&ngapMessage.ErrorIndication{}, &st) != "" {
		t.Fatal("ErrorIndication should not change the setup state")
	}
}
