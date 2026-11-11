package app

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/netip"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	libngap "libngap"
	"libngap/ngapMessage"
	"libngap/ngapType"

	"kronos/models"
	"kronos/simulator/gnb"
	"kronos/transport/sctp"
)

// sampleDir is the absolute path of the sample config/ (the tests do not run
// in the package directory, see TestMain).
var sampleDir string

func sample(name string) string { return filepath.Join(sampleDir, name) }

// TestMain runs the tests in a temporary directory: gnodeb writes its log/
// (application log and trace) under the working directory, and the source
// tree must not get any log.
func TestMain(m *testing.M) {
	dir, err := filepath.Abs("../../config")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	sampleDir = dir
	tmp, err := os.MkdirTemp("", "gnodeb-cmd-test-")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := os.Chdir(tmp); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	code := m.Run()
	os.RemoveAll(tmp)
	os.Exit(code)
}

func runArgs(args ...string) (int, string) {
	var out bytes.Buffer
	code := Run(context.Background(), "test", args, strings.NewReader(""), &out, &out)
	return code, out.String()
}

func TestCheckSample(t *testing.T) {
	code, out := runArgs("check", "-c", sample("config.yaml"))
	if code != 0 {
		t.Fatalf("exit %d\n%s", code, out)
	}
	// The sample signals carry the values of the sample config.
	for _, want := range []string{"0 error(s), 0 warning(s)"} {
		if !strings.Contains(out, want) {
			t.Fatalf("missing %q in\n%s", want, out)
		}
	}
}

func copyFile(t *testing.T, from, to string) {
	t.Helper()
	b, err := os.ReadFile(from)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(to, b, 0o644); err != nil {
		t.Fatal(err)
	}
}

// brokenSample copies the sample into a temporary directory with a move to an
// undefined state in the scenario, and returns the config path.
func brokenSample(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	copyFile(t, sample("config.yaml"), filepath.Join(dir, "config.yaml"))
	copyFile(t, sample("signal.txt"), filepath.Join(dir, "signal.txt"))
	src, err := os.ReadFile(sample("scenario.txt"))
	if err != nil {
		t.Fatal(err)
	}
	broken := strings.Replace(string(src), "move(sessionestablishment)", "move(nowhere)", 1)
	if broken == string(src) {
		t.Fatal("the sample no longer moves to sessionestablishment")
	}
	if err := os.WriteFile(filepath.Join(dir, "scenario.txt"), []byte(broken), 0o644); err != nil {
		t.Fatal(err)
	}
	return filepath.Join(dir, "config.yaml")
}

func TestCheckBrokenScenario(t *testing.T) {
	code, out := runArgs("-c", brokenSample(t), "check", "-v")
	if code != 1 || !strings.Contains(out, "error: state nowhere is not defined in scenario.txt") ||
		!strings.Contains(out, "entry → send(NGSetupRequest)") || !strings.Contains(out, "1 error(s)") {
		t.Fatalf("exit %d\n%s", code, out)
	}
}

func TestUsage(t *testing.T) {
	if code, _ := runArgs("bogus"); code != 2 {
		t.Fatal("unknown subcommand accepted")
	}
	if code, out := runArgs("version"); code != 0 || !strings.HasPrefix(out, "gnodeb ") {
		t.Fatal(out)
	}
	// --exec runs a batch and exits; the sample config is valid and the batch
	// does not connect.
	dir := t.TempDir()
	cfg := filepath.Join(dir, "config.yaml")
	for _, f := range []string{"config.yaml", "signal.txt", "scenario.txt"} {
		copyFile(t, sample(f), filepath.Join(dir, f))
	}
	ok, bad := filepath.Join(dir, "ok.txt"), filepath.Join(dir, "bad.txt")
	os.WriteFile(ok, []byte("help\nwait(10ms)\n"), 0o644)
	os.WriteFile(bad, []byte("help\nbogus(\n"), 0o644)
	if code, out := runArgs("-c", cfg, "--exec", ok); code != 0 || !strings.Contains(out, "send(signal)") || !strings.Contains(out, "(2 lines)") {
		t.Fatalf("--exec: %d %s", code, out)
	}
	if code, out := runArgs("-c", cfg, "--exec", bad); code != 1 || !strings.Contains(out, "bad.txt:2") {
		t.Fatalf("--exec with an error: %d %s", code, out)
	}
	// A scenario with an error must not start (and must not wait for input).
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var out bytes.Buffer
	if code := Run(ctx, "test", []string{"-c", brokenSample(t)}, strings.NewReader(""), &out, &out); code != 1 || !strings.Contains(out.String(), "not starting") {
		t.Fatalf("%d %s", code, out.String())
	}
}

func TestFormatEvent(t *testing.T) {
	at := time.Date(2026, 11, 10, 12, 0, 0, 101e6, time.Local)
	for _, tc := range []struct {
		e    models.Event
		want string
	}{
		{models.Event{Time: at, Kind: models.KindMessage, Source: "node", Dir: "tx", NGAP: "NGSetupRequest", Signal: "NGSetupRequest"},
			"12:00:00.101 node            -> NGSetupRequest                               (NGSetupRequest)"},
		{models.Event{Time: at, Kind: models.KindMessage, Source: "node", Dir: "rx", NGAP: "NGSetupResponse", Handler: "node"},
			"12:00:00.101 node            <- NGSetupResponse                              [node]"},
		{models.Event{Time: at, Kind: models.KindSnap, Source: "208930000000001", Text: "hello"},
			"12:00:00.101 208930000000001 hello"},
		{models.Event{Time: at, Kind: models.KindError, Source: "node", Handler: "node/entry", Text: "send(X): boom"},
			"12:00:00.101 node            error: [node/entry] send(X): boom"},
	} {
		if got := formatEvent(tc.e); got != tc.want {
			t.Fatalf("\n got %q\nwant %q", got, tc.want)
		}
	}
}

type syncBuffer struct {
	mu sync.Mutex
	b  bytes.Buffer
}

func (s *syncBuffer) Write(p []byte) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.Write(p)
}

func (s *syncBuffer) String() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.b.String()
}

func waitOutput(t *testing.T, out *syncBuffer, want string) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)
	for !strings.Contains(out.String(), want) {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %q in\n%s", want, out.String())
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func chdir(t *testing.T, dir string) {
	t.Helper()
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chdir(old) })
}

// TestRunWithStubAMF starts gnodeb against a loopback AMF: NG Setup from the
// node entry block, a CUI command from stdin, the trace file and the log.
func TestRunWithStubAMF(t *testing.T) {
	l, err := sctp.Listen(netip.MustParseAddrPort("127.0.0.1:0"), sctp.Options{})
	if err != nil {
		t.Skipf("SCTP is not available here: %v", err)
	}
	defer l.Close()
	signals := sample("signal.txt")
	dir := t.TempDir()
	copyFile(t, signals, filepath.Join(dir, "signal.txt"))
	cfg := fmt.Sprintf(`
base: { snap: true }
ue:
  imsi: "208930000000001"
  credential: { key: 8baf473f2f8fd09487cccbd7097c6862, op: 8e27b6af0e692e750f32667a3b14605d, type: OPC, amf: "8000" }
gnodeb:
  gnodebid: "000001"
  ngap: { src: { address: 127.0.0.1, port: 0 }, dst: { address: 127.0.0.1, port: %d } }
  gtpu: { src: { address: 127.0.0.1, port: 2152 } }
  plmn: { mcc: "208", mnc: "93" }
  snssai: [ { sst: 1, sd: "010203" } ]
  tac: 1
`, l.Addr().Port())
	scen := `node {
	entry { send(NGSetupRequest) }
	receive(NGSetupResponse) { snap("NG Setup OK") }
}`
	for name, body := range map[string]string{"config.yaml": cfg, "scenario.txt": scen} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(body), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	chdir(t, dir)

	pr, pw := io.Pipe()
	defer pw.Close()
	out := &syncBuffer{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan int, 1)
	go func() { done <- Run(ctx, "test", []string{"-c", "config.yaml"}, pr, out, out) }()
	waitOutput(t, out, "linkup connects SCTP")
	io.WriteString(pw, "linkup\n") // startup does not connect; linkup does

	amf, err := l.Accept()
	if err != nil {
		t.Fatal(err)
	}
	defer amf.Close()
	recvSetup := func() {
		b, _, err := amf.Recv()
		if err != nil {
			t.Fatal(err)
		}
		if m, _ := libngap.Decode(b); m == nil {
			t.Fatal("undecodable request")
		} else if _, ok := m.(*ngapMessage.NGSetupRequest); !ok {
			t.Fatalf("got %T", m)
		}
	}
	recvSetup()
	p := []byte{0x02, 0xf8, 0x39}
	resp, err := libngap.Encode(&ngapMessage.NGSetupResponse{
		AMFName: ngapType.AMFName{Value: "stub-amf"},
		ServedGUAMIList: ngapType.ServedGUAMIList{List: []ngapType.ServedGUAMIItem{{GUAMI: ngapType.GUAMI{
			PLMNIdentity: ngapType.PLMNIdentity{Value: p}, AMFRegionID: ngapType.AMFRegionID{Value: gnb.BitString(0xca, 8)},
			AMFSetID: ngapType.AMFSetID{Value: gnb.BitString(0x3f8, 10)}, AMFPointer: ngapType.AMFPointer{Value: gnb.BitString(0, 6)}}}}},
		RelativeAMFCapacity: ngapType.RelativeAMFCapacity{Value: 255},
		PLMNSupportList: ngapType.PLMNSupportList{List: []ngapType.PLMNSupportItem{{PLMNIdentity: ngapType.PLMNIdentity{Value: p},
			SliceSupportList: ngapType.SliceSupportList{List: []ngapType.SliceSupportItem{{SNSSAI: ngapType.SNSSAI{SST: ngapType.SST{Value: []byte{1}}}}}}}}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := amf.Send(resp, 0, models.PPIDNGAP); err != nil {
		t.Fatal(err)
	}
	waitOutput(t, out, "NG Setup succeeded: AMF stub-amf, GUAMI 208/93-cafe00")
	waitOutput(t, out, "node            NG Setup OK")

	io.WriteString(pw, "help()\nsend(NGSetupRequest)\n")
	waitOutput(t, out, "send(signal)")
	recvSetup()

	cancel()
	select {
	case code := <-done:
		if code != 0 {
			t.Fatalf("exit %d\n%s", code, out.String())
		}
	case <-time.After(5 * time.Second):
		t.Fatal("gnodeb did not stop")
	}
	t.Log("CUI output:\n" + out.String())
	traces, _ := filepath.Glob(filepath.Join(dir, "log", "trace-*.jsonl"))
	if len(traces) != 1 {
		t.Fatalf("trace files: %v", traces)
	}
	if b, _ := os.ReadFile(traces[0]); !bytes.Contains(b, []byte(`"ngap":"NGSetupResponse"`)) {
		t.Fatalf("trace:\n%s", b)
	}
	if b, err := os.ReadFile(filepath.Join(dir, "log", "gnodeb.txt")); err != nil || !bytes.Contains(b, []byte("NG Setup succeeded")) {
		t.Fatalf("application log: %v\n%s", err, b)
	}
}
