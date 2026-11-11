package state

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"kronos/common/diag"
)

func TestSampleConfig(t *testing.T) {
	cfg, ds := Load("../../config/config.yaml")
	if cfg == nil || ds.HasErrors() {
		t.Fatalf("sample config: %v", ds)
	}
	if n := ds.Count(diag.Warning); n != 0 {
		t.Fatalf("warnings: %v", ds)
	}
	// Digit strings keep their text (YAML would read 010203 as octal 4227).
	if cfg.UE.IMSI != "208930000000001" || cfg.GNodeB.SNSSAI[0].SD != "010203" || cfg.UE.SUCI.RoutingIndicator != "0000" ||
		cfg.GNodeB.PLMN.MNC != "93" || cfg.GNodeB.GNodeBID != "000001" || cfg.UE.Credential.AMF != "8000" {
		t.Fatalf("digit strings changed: %+v", cfg)
	}
	if s := cfg.UE.Session; len(s) != 1 || s[0].DNN != "internet" || s[0].Type != "IPv4" || s[0].SNSSAI.SD != "010203" {
		t.Fatalf("legacy ue.type/apn: %+v", s)
	}
	if w := cfg.WebListen(); w == nil || w.Address != "10.10.10.30" || w.Port != 8080 {
		t.Fatalf("gnodeb.web: %+v", w)
	}
	if h := cfg.GNodeB.HTTP; h == nil || h.Src == nil || h.Src.Port != 12345 {
		t.Fatalf("gnodeb.http (between pseudo-calls): %+v", h)
	}
	if nci, err := cfg.GNodeB.NCIValue(); err != nil || nci != 0x10 {
		t.Fatalf("default NCI %x %v", nci, err)
	}
	if cfg.ScenarioPath() != filepath.Join("../../config", "scenario.txt") || !cfg.Base.Snap {
		t.Fatalf("paths/snap: %s %v", cfg.ScenarioPath(), cfg.Base.Snap)
	}
	if !strings.Contains(joinDiags(ds), "not quoted: imsi (line 6)") {
		t.Fatalf("quoting hint missing:\n%s", joinDiags(ds))
	}
}

func joinDiags(ds diag.List) string {
	var b strings.Builder
	for _, d := range ds {
		b.WriteString(d.String() + "\n")
	}
	return b.String()
}

const base = `
ue:
  imsi: "208930000000001"
  subscription: 3
  credential: { key: 8baf473f2f8fd09487cccbd7097c6862, op: 8e27b6af0e692e750f32667a3b14605d, type: OPC, amf: "8000" }
gnodeb:
  gnodebid: "000001"
  ngap: { dst: { address: 10.10.10.40, port: 38412 } }
  gtpu: { src: { address: 10.10.10.30, port: 2152 } }
  plmn: { mcc: "208", mnc: "93" }
  snssai: [ { sst: 1, sd: "010203" } ]
  tac: 1
`

func load(t *testing.T, yaml string) (*Config, diag.List) {
	t.Helper()
	p := filepath.Join(t.TempDir(), "config.yaml")
	if err := os.WriteFile(p, []byte(yaml), 0o644); err != nil {
		t.Fatal(err)
	}
	return Load(p)
}

func TestSubscribers(t *testing.T) {
	cfg, ds := load(t, base)
	if ds.HasErrors() || ds.Count(diag.Warning) != 0 {
		t.Fatalf("%v", ds)
	}
	if imsi, _ := cfg.SubscriberIMSI(2); imsi != "208930000000003" {
		t.Fatal(imsi)
	}
	if i, ok := cfg.SubscriberIndex("208930000000002"); !ok || i != 1 {
		t.Fatal(i, ok)
	}
	for _, imsi := range []string{"208930000000004", "208930000000000", "20893000000001", "x"} {
		if _, ok := cfg.SubscriberIndex(imsi); ok {
			t.Fatalf("%s accepted", imsi)
		}
	}
	if cfg.WebListen() != nil || cfg.UE.Integrity == nil || !cfg.UE.Integrity.NIA2 || cfg.UE.Integrity.NIA3 {
		t.Fatal("defaults")
	}
}

func TestConfigErrors(t *testing.T) {
	for edit, want := range map[[2]string]string{
		{`imsi: "208930000000001"`, `imsi: "20893"`}:                            "ue.imsi must be 15 digits",
		{`imsi: "208930000000001"`, `imsi: "208939999999999"`}:                  "overflow the digits",
		{`imsi: "208930000000001"`, `imsi: "001010000000001"`}:                  "does not start with gnodeb.plmn (20893)",
		{"subscription: 3", "subscription: 0"}:                                  "1 or more",
		{"tac: 1", "tac: abc"}:                                                  "3 octets",
		{"tac: 1", "tac: 0x1000000"}:                                            "3 octets",
		{"type: OPC", "type: XYZ"}:                                              "OP or OPC",
		{`amf: "8000"`, `amf: "0000"`}:                                          "separation bit",
		{`gnodebid: "000001"`, `gnodebid: "000001"` + "\n  nci: \"000000020\""}: "top 32 bits",
		{`gnodebid: "000001"`, `gnodebid: "zz"`}:                                "hex digits",
		{"address: 10.10.10.40", "address: amf.local"}:                          "not an IP address",
		{"port: 38412", "port: 0"}:                                              "1 to 65535",
		{"subscription: 3", "subscription: 3\n  subscriptions: 3"}:              "unknown key subscriptions",
		{"subscription: 3", "subscription: many"}:                               "as an integer",
		{`mnc: "93"`, `mnc: "9"`}:                                               "2 or 3 digits",
		{`sd: "010203"`, `sd: "0102"`}:                                          "6 hex digits",
		{"gnodeb:", "base: { role: amf }\ngnodeb:"}:                             "not implemented",
		{"gnodeb:", "ue2: 1\ngnodeb:"}:                                          "unknown key ue2",
	} {
		src := strings.Replace(base, edit[0], edit[1], 1)
		if src == base {
			t.Fatalf("edit %q did not apply", edit[0])
		}
		_, ds := load(t, src)
		if !strings.Contains(joinDiags(ds), want) {
			t.Fatalf("%s → want %q, got:\n%s", edit[1], want, joinDiags(ds))
		}
	}
}

func TestTACIsNotOctal(t *testing.T) {
	cfg, ds := load(t, strings.Replace(base, "tac: 1", "tac: 000010", 1))
	if ds.HasErrors() || cfg.GNodeB.TAC.Value != 10 {
		t.Fatalf("tac %d %v", cfg.GNodeB.TAC.Value, ds)
	}
	cfg, _ = load(t, strings.Replace(base, "tac: 1", "tac: 0x10", 1))
	if cfg.GNodeB.TAC.Value != 16 {
		t.Fatal(cfg.GNodeB.TAC.Value)
	}
}

func TestBrokenYAML(t *testing.T) {
	if cfg, ds := load(t, "ue: [\n"); cfg != nil || !ds.HasErrors() {
		t.Fatalf("%v", ds)
	}
	if cfg, ds := load(t, ""); cfg != nil || !ds.HasErrors() {
		t.Fatalf("empty: %v", ds)
	}
	if _, ds := Load(filepath.Join(t.TempDir(), "missing.yaml")); !ds.HasErrors() {
		t.Fatal("missing file")
	}
}

// SUCI Profile A and B take the home network public key of their curve.
func TestSUCIProfileConfig(t *testing.T) {
	const a = "5a8d38864820197c3394b92613b20b91633cbd897119273bf8e4a6f4eec0a650"
	const b = "0272da71976234ce833a6907425867b82e074d44ef907dfb4b3e21c1c2256ebcd1"
	for suci, want := range map[string]string{
		"{ protectionscheme: 1, keyid: 1, key: " + a + " }": "",
		"{ protectionscheme: 2, keyid: 2, key: " + b + " }": "",
		"{ protectionscheme: 1, keyid: 1 }":                 "Profile A/B needs the home network public key",
		"{ protectionscheme: 1, key: zz }":                  "ue.suci.key is not hex",
		"{ protectionscheme: 1, key: " + b + " }":           "Profile A takes a 32-octet X25519 public key",
		"{ protectionscheme: 2, key: " + a + " }":           "Profile B takes a secp256r1 public key",
		"{ protectionscheme: 3 }":                           "must be 0 to 2",
	} {
		src := strings.Replace(base, `  imsi: "208930000000001"`, `  imsi: "208930000000001"`+"\n  suci: "+suci, 1)
		_, ds := load(t, src)
		if text := joinDiags(ds); want == "" && strings.Contains(text, "suci") || want != "" && !strings.Contains(text, want) {
			t.Errorf("%s → want %q, got:\n%s", suci, want, text)
		}
	}
}

// ue.tun of older configs is read, only to say it is no longer needed.
func TestTUNKeyIsNoLongerNeeded(t *testing.T) {
	src := strings.Replace(base, `  imsi: "208930000000001"`, `  imsi: "208930000000001"`+"\n  tun: true", 1)
	_, ds := load(t, src)
	if text := joinDiags(ds); ds.HasErrors() || !strings.Contains(text, "info: ue.tun is no longer needed") {
		t.Fatal(text)
	}
}
