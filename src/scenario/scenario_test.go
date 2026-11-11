package scenario

import (
	"bytes"
	"encoding/hex"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"testing"

	libnas "libnas"
	libngap "libngap"

	"kronos/command"
	"kronos/common/diag"
	"kronos/models"
)

const (
	sampleSignals  = "../../config/signal.txt"
	sampleScenario = "../../config/scenario.txt"
)

func TestNGAPTableMatchesLibrary(t *testing.T) {
	files, err := filepath.Glob("../lib/ngap/ngapMessage/*.go")
	if err != nil || len(files) == 0 {
		t.Fatalf("glob: %v", err)
	}
	var want []string
	for _, f := range files {
		if !strings.HasSuffix(f, "_test.go") {
			want = append(want, strings.TrimSuffix(filepath.Base(f), ".go"))
		}
	}
	sort.Strings(want)
	if got := NGAPNames(); !reflect.DeepEqual(got, want) {
		t.Fatalf("ngap_messages.go is stale; run `go generate ./scenario` (have %d, library has %d)", len(got), len(want))
	}
}

func TestNGAPClassification(t *testing.T) {
	ue := 0
	for _, m := range NGAPMessages() {
		if m.Assoc != AssocNode {
			ue++
		}
	}
	if ue != 69 { // 69 of the 144 messages carry UE NGAP IDs
		t.Fatalf("UE-associated messages: %d", ue)
	}
	for name, want := range map[string]Assoc{
		"NGSetupRequest": AssocNode, "Paging": AssocNode, "DownlinkNASTransport": AssocUE,
		"InitialUEMessage": AssocUE, "UEContextReleaseCommand": AssocUE, "ErrorIndication": AssocEither,
	} {
		if m, ok := LookupNGAP(name); !ok || m.Assoc != want {
			t.Fatalf("%s: %v", name, m)
		}
	}
	for name, want := range map[string]bool{
		"DownlinkNASTransport": true, "PDUSessionResourceSetupRequest": true, "InitialContextSetupRequest": true,
		"NGSetupResponse": false, "UEContextReleaseCommand": false,
	} {
		if m, _ := LookupNGAP(name); m.CarriesNAS != want {
			t.Fatalf("%s: CarriesNAS %v", name, m.CarriesNAS)
		}
	}
	if m, ok := LookupNGAP("DownLinkNasTransport"); !ok || m.Name != "DownlinkNASTransport" {
		t.Fatal("case-insensitive lookup")
	}
	if n, ok := LookupNAS("5gmmstatus"); !ok || n.Name != "Status5GMM" {
		t.Fatal("NAS alias")
	}
}

func mustHex(t *testing.T, s string) []byte {
	t.Helper()
	set, ds := ParseSignals("t", []byte("X [raw] {"+s+"}"))
	if ds.HasErrors() || len(set.List) != 1 {
		t.Fatalf("hex %s: %v", s, ds)
	}
	return set.List[0].Bytes
}

func readSample(t *testing.T) *SignalSet {
	t.Helper()
	set, ds := ReadSignals(sampleSignals)
	if ds.HasErrors() {
		t.Fatalf("sample signals: %v", ds)
	}
	return set
}

func TestSampleSignals(t *testing.T) {
	set := readSample(t)
	if len(set.List) != 18 {
		t.Fatalf("signals: %d", len(set.List))
	}
	for _, s := range set.List {
		// The builder relies on templates re-encoding byte for byte.
		re, err := libngap.Encode(s.Msg)
		if err != nil || !bytes.Equal(re, s.Bytes) {
			t.Fatalf("%s: NGAP round trip differs (%v)", s.Name, err)
		}
		for _, f := range findNAS(s.Msg) {
			plain := f.b
			if plain[0] == models.EPD5GMM && plain[1] != 0 {
				plain = plain[7:]
			}
			m, err := libnas.Decode(plain, nil)
			if err != nil {
				t.Fatalf("%s: NAS: %v", s.Name, err)
			}
			if re, err := libnas.Encode(m, nil); err != nil || !bytes.Equal(re, plain) {
				t.Fatalf("%s: NAS round trip differs (%v)", s.Name, err)
			}
		}
	}
	for name, want := range map[string]string{
		"NGSetupRequest":                                   "NGSetupRequest",
		"UpLinkNASTransportSecurityModeComplete":           "UplinkNASTransport:SecurityModeComplete(4)",
		"UplinkNASTransportPDUSessionEstablishmentRequest": "UplinkNASTransport:ULNASTransport>PDUSessionEstablishmentRequest(2)",
		"InitialUEMeesageServiceRequest":                   "InitialUEMessage:ServiceRequest(1)",
	} {
		if s, _ := set.Lookup(name); s.Label() != want {
			t.Fatalf("%s: %s", name, s.Label())
		}
	}
	if s, _ := set.Lookup("NGSetupRequest"); s.Scope != command.ScopeNode {
		t.Fatal("NGSetupRequest scope")
	}
	if s, _ := set.Lookup("UEContextReleaseComplete"); s.Scope != command.ScopeUE {
		t.Fatal("UEContextReleaseComplete scope")
	}
}

// compareSamples are captures whose values differ from the config: the
// sample signals had them before they were made to match it.
const compareSamples = `
NGSetupRequest { 00150035000004001b00090002f839500000000100524008028035672d72616e0066001000000000010002f839000010080102030015400140 }
UEContextReleaseComplete { 20290024000003000a40020001005540048001000c0079400f402f83900000389d102f8390000001 }
UplinkNASTransportPDUSessionReleaseRequest {
  002e403e000004000a00020003005500048001000e00260016157e024aaf3ac2057e00670100062e0109d1592412010079400f409909200000389d10990920000001
}
`

func TestCompareFindsProblems(t *testing.T) {
	ref := models.SignalReference{MCC: "208", MNC: "93", TAC: 1, HasTAC: true, NCI: 0x10, HasNCI: true, GNBID: 1, GNBIDLen: 32, HasGNBID: true}
	set, ds := ParseSignals("signal.txt", []byte(compareSamples))
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	text := joinDiags(set.Compare(ref))
	for _, want := range []string{
		"UEContextReleaseComplete: PLMN 2f8390 is not valid BCD",
		"UplinkNASTransportPDUSessionReleaseRequest: PLMN differs from 208/93 of config: 999/020",
		"NR cell ID 0000389d1",
		"RANNodeName is sent as \"5g-ran\"",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("missing %q in\n%s", want, text)
		}
	}
	if strings.Contains(text, "TAC") || strings.Contains(text, "gNB ID") {
		t.Fatalf("unexpected TAC / gNB ID difference:\n%s", text)
	}
	// The sample signals match the sample config.
	if ds := readSample(t).Compare(ref); ds.Count(diag.Warning) > 0 {
		t.Fatalf("the sample signals differ from the config:\n%s", joinDiags(ds))
	}
}

func TestParseSignalErrors(t *testing.T) {
	for src, want := range map[string]string{
		"A { 0 }":                        "odd number",
		"A { 00zz }":                     "is not a hex digit",
		"A { }":                          "is empty",
		"A { 00":                         "not closed",
		"A [fast] { 00 }":                "unknown option",
		"A [raw, nas=keep] { 00 }":       "cannot go together",
		"A 00 }":                         "must be followed by {",
		"A [raw] { 00 }\nA [raw] { 00 }": "used twice",
		"A { 00 }":                       "cannot decode as NGAP",
	} {
		_, ds := ParseSignals("t", []byte(src))
		if !ds.HasErrors() || !strings.Contains(joinDiags(ds), want) {
			t.Fatalf("%q: want %q, got %v", src, want, ds)
		}
	}
	// raw signals need not decode: they are sent as written.
	if _, ds := ParseSignals("t", []byte("A [raw] { 00 }")); ds.HasErrors() {
		t.Fatalf("raw: %v", ds)
	}
}

func joinDiags(ds diag.List) string {
	var b strings.Builder
	for _, d := range ds {
		b.WriteString(d.String() + "\n")
	}
	return b.String()
}

func TestSampleScenario(t *testing.T) {
	set := readSample(t)
	f, ds := ReadScenario(sampleScenario)
	if ds.HasErrors() {
		t.Fatalf("parse: %v", ds)
	}
	ds = Validate(f, set, models.ValidateOptions{SnapEnabled: true})
	if ds.HasErrors() {
		t.Fatalf("the sample scenario has errors:\n%s", joinDiags(ds))
	}
	// The PDU session is its own state, started by its entry block.
	if b, ok := f.State("sessionestablishment"); !ok || b.Entry == nil || b.Entry.Calls[0].Args[0].Name != "UplinkNASTransportPDUSessionEstablishmentRequest" {
		t.Fatal("sessionestablishment state with its entry is missing")
	}
	h := f.States[0].Handlers[0]
	if h.Trigger.NGAPInfo.Name != "DownlinkNASTransport" || h.Trigger.NASInfo.Name != "AuthenticationRequest" {
		t.Fatalf("trigger not resolved: %+v", h.Trigger)
	}
	if !strings.Contains(joinDiags(ds), "signals the scenario does not send") {
		t.Fatal("unused signals not reported")
	}
}

const goodScenario = `
var {
	public calls = 0
	private label = "ue"
}

node {
	entry {
		send(NGSetupRequest)
	}
	receive(NGSetupResponse) {
		snap("NG Setup ok, calls=%d", calls)
	}
	receive(ErrorIndication) {}
}

any {
	receive(UEContextReleaseCommand) {
		send(UEContextReleaseComplete)
		move(idle)
	}
}

start {
	entry {
		send(InitialUEMeesageRegistrationRequest)   // starts after NG Setup succeeds
	}
	receive(DownlinkNASTransport:AuthenticationRequest) {
		send(UpLinkNASTransportAuthenticationResponse)
	}
	receive(DownlinkNASTransport:SecurityModeCommand) {
		send(UpLinkNASTransportSecurityModeComplete)
	}
	receive(InitialContextSetupRequest:RegistrationAccept) {
		send(InitialContextSetupResponse)
		send(UplinkNASTransportRegistrationComplete)
		move(registered)
	}
}

registered {
	entry {
		snap("%s %s registered", label, $imsi)
		send(UplinkNASTransportPDUSessionEstablishmentRequest)
	}
	receive(PDUSessionResourceSetupRequest:PDUSessionEstablishmentAccept) {
		send(PDUSessionResourceSetupResponse)
	}
}

idle {
	receive(Paging) {
		send(InitialUEMeesageServiceRequest)
	}
}
`

func TestValidScenarioWithEntry(t *testing.T) {
	set := readSample(t)
	f, ds := ParseScenario("good.txt", []byte(goodScenario))
	ds = append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...)
	if ds.HasErrors() || ds.Count(diag.Warning) > 0 {
		t.Fatalf("unexpected diagnostics:\n%s", joinDiags(ds))
	}
	if f.Node.Entry == nil || len(f.Node.Entry.Calls) != 1 || f.States[0].Entry.Calls[0].Args[0].Name != "InitialUEMeesageRegistrationRequest" {
		t.Fatal("entry blocks not parsed")
	}
	if n := f.StripCalls("snap"); n != 2 {
		t.Fatalf("stripped %d snap calls", n)
	}
	if len(f.Node.Handlers[0].Calls) != 0 || len(f.Node.Handlers) != 2 {
		t.Fatal("handler should stay after its snap is stripped")
	}
}

func TestScenarioErrors(t *testing.T) {
	set := readSample(t)
	for src, want := range map[string]string{
		"node { receive(DownlinkNASTransport) {} }":                            "is UE-associated, so node cannot receive it",
		"start { receive(NGSetupResponse) {} }":                                "belongs to the node",
		"node { receive(NGSetupResponse) { send(UEContextReleaseComplete) } }": "is a UE signal, so the node block cannot send it",
		"start { receive(Paging) { send(NGSetupRequest) } }":                   "is a node signal",
		"start { receive(DownlinkNasTransprt) {} }":                            "did you mean DownlinkNASTransport?",
		"start { receive(DownlinkNASTransport:AuthRequest) {} }":               "unknown NAS message",
		"start { receive(UEContextReleaseCommand:ServiceAccept) {} }":          "carries no NAS",
		"start { receive(Paging) {} receive(paging) {} }":                      "the same trigger",
		"any { entry { } }":           "any cannot have an entry",
		"start { entry {} entry {} }": "only one entry per block",
		"start { receive(Paging) { target(208930000000001).send(InitialUEMeesageServiceRequest) } }": "target() cannot be used in a scenario",
		"start { receive(Paging) { snap(\"%d %d\", 1) } }":                                           "2 conversion(s) in the format but 1 value(s)",
		"node { receive(NGSetupResponse) { move(start) } } start {}":                                 "cannot be used in the node block",
		"start { timeout(5s) { } }":                                                                  "timeout is not supported",
		"start { receive(*) { } }":                                                                   "receive(*) is not supported",
		"start { send(NGSetupRequest) }":                                                             "cannot be written directly in a block",
		"idle { }":                                                                                   "there is no initial state start",
		"start {} start {}":                                                                          "a second state start",
		"receive {}":                                                                                 "the reserved word receive cannot name a state",
		"var { public a = 1 private a = 2 }":                                                         "declared twice",
		"var { public a = b }":                                                                       "must be a literal",
		"node { receive(NGSetupResponse) { snap(\"%s\", s) } } var { private s = \"x\" }": "private variable s cannot be used in node",
		"start { receive(Paging) { send(NGSetupRequest) }":                                "not closed",
		"start { receive(Paging) { send(NGSetupRequest).move(start) } }":                  "cannot be joined",
	} {
		f, ds := ParseScenario("t.txt", []byte(src))
		ds = append(ds, Validate(f, set, models.ValidateOptions{SnapEnabled: true})...)
		if !ds.HasErrors() || !strings.Contains(joinDiags(ds), want) {
			t.Fatalf("%s\nwant %q, got:\n%s", src, want, joinDiags(ds))
		}
	}
}

func TestSnapDisabledNote(t *testing.T) {
	f, _ := ParseScenario("t.txt", []byte(`node { receive(NGSetupResponse) { snap("x") } }`))
	ds := Validate(f, readSample(t), models.ValidateOptions{SnapEnabled: false})
	if !strings.Contains(joinDiags(ds), "snap (1 place(s)) does not run") {
		t.Fatal(joinDiags(ds))
	}
}

func TestReadMissingFiles(t *testing.T) {
	dir := t.TempDir()
	if _, ds := ReadSignals(filepath.Join(dir, "none.txt")); !ds.HasErrors() {
		t.Fatal("missing signal file")
	}
	if _, ds := ReadScenario(filepath.Join(dir, "none.txt")); !ds.HasErrors() {
		t.Fatal("missing scenario file")
	}
	p := filepath.Join(dir, "s.txt")
	if err := os.WriteFile(p, []byte("start { receive(Paging) { send(Anything) } }"), 0o644); err != nil {
		t.Fatal(err)
	}
	f, _ := ReadScenario(p)
	// Without signal.txt, signal names cannot be checked and must not cascade into errors.
	if ds := Validate(f, nil, models.ValidateOptions{}); ds.HasErrors() {
		t.Fatal(joinDiags(ds))
	}
}

// The security header type of a 5GMM NAS-PDU is the low half of its second
// octet; the high half is spare (TS 24.501 §9.3).
func TestNASSpareBits(t *testing.T) {
	set := readSample(t)
	reg, _ := set.Lookup("InitialUEMeesageRegistrationRequest")
	smc, _ := set.Lookup("UpLinkNASTransportSecurityModeComplete")
	flip := func(b []byte, from, to string) string {
		h := hex.EncodeToString(b)
		if !strings.Contains(h, from) {
			t.Fatalf("no %s in %s", from, h)
		}
		return strings.Replace(h, from, to, 1)
	}
	// Plain, with spare bits: reported as such, not read as protected.
	src := "Plain { " + flip(reg.Bytes, "7e004179", "7e104179") + " }\n" +
		// Protected, with spare bits: the header is still skipped.
		"Protected { " + flip(smc.Bytes, "7e04144a17fc", "7e14144a17fc") + " }\n"
	s, ds := ParseSignals("signal.txt", []byte(src))
	text := joinDiags(ds)
	if !strings.Contains(text, "Plain: the NAS of NASPDU: the spare bits of the security header octet are set (0x10)") {
		t.Fatalf("plain:\n%s", text)
	}
	if strings.Contains(text, "Protected") {
		t.Fatalf("protected:\n%s", text)
	}
	if p, _ := s.Lookup("Protected"); len(p.NAS) != 1 || p.NAS[0].Header != 4 || p.NAS[0].Name != "SecurityModeComplete" {
		t.Fatalf("protected NAS %+v", p.NAS)
	}
}
