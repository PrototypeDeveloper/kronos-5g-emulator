package webconsole_test

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"

	"kronos/models"
	"kronos/scenario"
	"kronos/simulator/engine"
	"kronos/simulator/gnb"
	"kronos/state"
	"kronos/webconsole"
)

// newConsole serves the console for an engine running the sample files
// (read only; nothing is written to the source tree) with no link.
func newConsole(t *testing.T) (*httptest.Server, *engine.Engine) {
	t.Helper()
	return newConsoleOn(t, "../../config")
}

// newConsoleCopy runs the console on a copy of the sample files, which the
// test may change.
func newConsoleCopy(t *testing.T) (*httptest.Server, *engine.Engine, string) {
	t.Helper()
	dir := t.TempDir()
	for _, f := range []string{"config.yaml", "signal.txt", "scenario.txt"} {
		b, err := os.ReadFile(filepath.Join("../../config", f))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, f), b, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	ts, eng := newConsoleOn(t, dir)
	return ts, eng, dir
}

func newConsoleOn(t *testing.T, dir string) (*httptest.Server, *engine.Engine) {
	t.Helper()
	cfg, ds := state.Load(filepath.Join(dir, "config.yaml"))
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	sigs, ds := scenario.ReadSignals(cfg.SignalPath())
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	f, ds := scenario.ReadScenario(cfg.ScenarioPath())
	ds = append(ds, scenario.Validate(f, sigs, models.ValidateOptions{SnapEnabled: cfg.Base.Snap})...)
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	role, err := gnb.New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	eng := engine.New(cfg, sigs, f, role)
	ctx, cancel := context.WithCancel(context.Background())
	eng.Start(ctx)
	ts := httptest.NewServer(webconsole.New(eng, models.WebOptions{Version: "test"}).Handler())
	t.Cleanup(func() {
		ts.Close()
		cancel()
		eng.Wait()
	})
	return ts, eng
}

func get(t *testing.T, ts *httptest.Server, path string, want int, v any) string {
	t.Helper()
	r, err := http.Get(ts.URL + path)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	if r.StatusCode != want {
		t.Fatalf("GET %s: %d %s", path, r.StatusCode, b)
	}
	if v != nil {
		if err := json.Unmarshal(b, v); err != nil {
			t.Fatalf("GET %s: %v\n%s", path, err, b)
		}
	}
	return string(b)
}

func post(t *testing.T, ts *httptest.Server, line string) (code int, reply map[string]string) {
	t.Helper()
	body, _ := json.Marshal(map[string]string{"line": line})
	r, err := http.Post(ts.URL+"/api/v1/commands", "application/json", strings.NewReader(string(body)))
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	json.NewDecoder(r.Body).Decode(&reply)
	return r.StatusCode, reply
}

func TestGUIFiles(t *testing.T) {
	ts, _ := newConsole(t)
	if b := get(t, ts, "/", 200, nil); !strings.Contains(b, "KRONOS") || !strings.Contains(b, "/app.js") || !strings.Contains(b, `data-page="license"`) {
		t.Fatalf("index:\n%s", b)
	}
	// The License page: the copyright holder and the license.
	if b := get(t, ts, "/app.js", 200, nil); !strings.Contains(b, "Copyright 2026 Kronos Systems.") || !strings.Contains(b, "Apache License, Version 2.0") {
		t.Fatal("app.js has no License page")
	}
	r, err := http.Get(ts.URL + "/kronos-mark.svg")
	if err != nil {
		t.Fatal(err)
	}
	r.Body.Close()
	if r.StatusCode != 200 || !strings.HasPrefix(r.Header.Get("Content-Type"), "image/svg+xml") {
		t.Fatalf("logo: %d %s", r.StatusCode, r.Header.Get("Content-Type"))
	}
	for _, p := range []string{"/app.js", "/style.css"} {
		get(t, ts, p, 200, nil)
	}
	var e map[string]string
	get(t, ts, "/api/v1/nothing", 404, &e)
	if e["error"] == "" {
		t.Fatal("no error message for an unknown API path")
	}
}

func TestSnapshots(t *testing.T) {
	ts, _ := newConsole(t)
	var st struct {
		models.Overview
		Version string `json:"version"`
	}
	get(t, ts, "/api/v1/status", 200, &st)
	if st.Version != "test" || st.Subscribers != 1 || st.Node.Connected || st.States["start"] != 1 || st.RM[models.RMDeregistered] != 1 {
		t.Fatalf("status %+v", st)
	}
	var ues struct {
		Total int               `json:"total"`
		Items []models.UEStatus `json:"items"`
	}
	get(t, ts, "/api/v1/ues", 200, &ues)
	if ues.Total != 1 || len(ues.Items) != 1 || ues.Items[0].IMSI != "208930000000001" || ues.Items[0].CM != models.CMIdle {
		t.Fatalf("ues %+v", ues)
	}
	get(t, ts, "/api/v1/ues?state=nowhere", 200, &ues)
	if ues.Total != 0 || len(ues.Items) != 0 {
		t.Fatalf("filtered ues %+v", ues)
	}
	var ue models.UEStatus
	get(t, ts, "/api/v1/ues/208930000000001", 200, &ue)
	if ue.State != "start" || ue.Vars["s"] != `"test"` {
		t.Fatalf("ue %+v", ue)
	}
	get(t, ts, "/api/v1/ues/1", 404, nil)

	var sigs []models.SignalInfo
	get(t, ts, "/api/v1/signals", 200, &sigs)
	if len(sigs) != 18 || sigs[0].Name != "NGSetupRequest" || sigs[0].Scope != "node" || sigs[0].Hex == "" {
		t.Fatalf("signals %+v", sigs)
	}
	var tr struct {
		Signal models.SignalInfo `json:"signal"`
		Tree   models.TreeNode   `json:"tree"`
	}
	get(t, ts, "/api/v1/signals/InitialUEMeesageRegistrationRequest", 200, &tr)
	if tr.Tree.Name != "InitialUEMessage" || !strings.Contains(get(t, ts, "/api/v1/signals/InitialUEMeesageRegistrationRequest", 200, nil), `"RegistrationRequest"`) {
		t.Fatalf("signal tree %+v", tr.Tree)
	}
	get(t, ts, "/api/v1/signals/nothing", 404, nil)

	var sc models.ScenarioInfo
	get(t, ts, "/api/v1/scenario", 200, &sc)
	kinds := map[string]bool{}
	for _, b := range sc.Blocks {
		kinds[b.Kind+":"+b.Name] = true
	}
	if !kinds["node:node"] || !kinds["any:any"] || !kinds["state:start"] || !kinds["state:sessionestablishment"] || len(sc.Vars) != 4 {
		t.Fatalf("scenario %+v", sc)
	}
	var vars []models.VarValue
	get(t, ts, "/api/v1/vars", 200, &vars)
	if len(vars) != 4 || vars[0].Name != "a" || !vars[0].Public || vars[0].Value != "1" {
		t.Fatalf("vars %+v", vars)
	}

	var cmds []struct {
		Name  string   `json:"name"`
		Scope []string `json:"scope"`
		Args  []struct {
			Name string `json:"name"`
			Kind string `json:"kind"`
		} `json:"args"`
	}
	get(t, ts, "/api/v1/commands", 200, &cmds)
	found := map[string]string{}
	for _, c := range cmds {
		s := strings.Join(c.Scope, ",")
		for _, a := range c.Args {
			s += " " + a.Name + ":" + a.Kind
		}
		found[c.Name] = s
	}
	if found["send"] != "signal signal:signal" || found["ping"] != "ue dst:value count:value interval:value size:value" ||
		found["iperf"] != "ue dst:value port:value rate:rate time:value dir:choice len:value" || found["stop"] != "ue kind:choice" || found["linkup"] != "console" {
		t.Fatalf("commands %v", found)
	}
}

func TestCommandsAndTrace(t *testing.T) {
	ts, eng := newConsole(t)
	if code, r := post(t, ts, "help"); code != 200 || !strings.Contains(r["output"], "send(signal)") || r["error"] != "" {
		t.Fatalf("help: %d %v", code, r)
	}
	if code, r := post(t, ts, "bogus("); code != 200 || r["error"] == "" {
		t.Fatalf("bad line: %d %v", code, r)
	}
	if code, r := post(t, ts, "linkdown"); code != 200 || !strings.Contains(r["error"], "not connected") {
		t.Fatalf("linkdown without a link: %d %v", code, r)
	}
	if code, _ := post(t, ts, " "); code != 400 {
		t.Fatalf("empty line: %d", code)
	}
	r, err := http.Post(ts.URL+"/api/v1/commands", "text/plain", strings.NewReader(`{"line":"help"}`))
	if err != nil {
		t.Fatal(err)
	}
	r.Body.Close()
	if r.StatusCode != http.StatusUnsupportedMediaType {
		t.Fatalf("text/plain accepted: %d", r.StatusCode)
	}

	// The lines run from the GUI and their errors are events of source "web".
	var tr struct {
		Events []models.Event `json:"events"`
	}
	get(t, ts, "/api/v1/trace?ue=web", 200, &tr)
	var texts []string
	for _, e := range tr.Events {
		texts = append(texts, e.Kind+":"+e.Text)
	}
	if len(tr.Events) != 5 || texts[0] != "command:help" || !strings.HasPrefix(texts[2], "error:bogus(: ") {
		t.Fatalf("web events %q", texts)
	}
	get(t, ts, "/api/v1/trace?limit=2", 200, &tr)
	if len(tr.Events) != 2 {
		t.Fatalf("limit: %d events", len(tr.Events))
	}

	// A message event decodes with its PDU.
	sig := eng.Signals()[0]
	eng.Events().Publish(models.Event{Kind: models.KindMessage, Source: models.SourceNode, Dir: "tx", NGAP: sig.NGAP, Signal: sig.Name, Hex: sig.Hex})
	get(t, ts, "/api/v1/trace?limit=1", 200, &tr)
	var one struct {
		Event models.Event    `json:"event"`
		Tree  models.TreeNode `json:"tree"`
	}
	get(t, ts, "/api/v1/trace/"+itoa(tr.Events[0].Seq), 200, &one)
	if one.Event.Signal != "NGSetupRequest" || one.Tree.Name != "NGSetupRequest" || len(one.Tree.Kids) == 0 {
		t.Fatalf("decoded event %+v", one)
	}
	get(t, ts, "/api/v1/trace/999999", 404, nil)
}

func itoa(n uint64) string { return strconv.FormatUint(n, 10) }

// The event stream starts with hello, replays the kept events and follows.
func TestEventStream(t *testing.T) {
	ts, _ := newConsole(t)
	post(t, ts, "help")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", ts.URL+"/api/v1/events?tail=1", nil)
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	if ct := r.Header.Get("Content-Type"); !strings.HasPrefix(ct, "text/event-stream") {
		t.Fatalf("Content-Type %s", ct)
	}
	sc := bufio.NewScanner(r.Body)
	next := func() (event, id, data string) {
		t.Helper()
		for sc.Scan() {
			line := sc.Text()
			switch {
			case line == "":
				if data != "" {
					return
				}
			case strings.HasPrefix(line, "event: "):
				event = line[7:]
			case strings.HasPrefix(line, "id: "):
				id = line[4:]
			case strings.HasPrefix(line, "data: "):
				data = line[6:]
			}
		}
		t.Fatalf("stream ended: %v", sc.Err())
		return
	}
	ev, _, data := next()
	var hello map[string]any
	json.Unmarshal([]byte(data), &hello)
	if ev != "hello" || hello["version"] != "test" || hello["replay"] != float64(1) {
		t.Fatalf("hello %s %s", ev, data)
	}
	// tail=1: only the last kept event (the "help" line from the GUI).
	_, id, data := next()
	var e models.Event
	json.Unmarshal([]byte(data), &e)
	if e.Source != models.SourceWeb || e.Text != "help" || id != itoa(e.Seq) {
		t.Fatalf("replayed %s %s", id, data)
	}
	post(t, ts, "wait(1ms)")
	_, _, data = next()
	json.Unmarshal([]byte(data), &e)
	if e.Text != "wait(1ms)" {
		t.Fatalf("live %s", data)
	}
}

// send makes a JSON request and decodes the reply.
func send(t *testing.T, ts *httptest.Server, method, path string, body any, v any) int {
	t.Helper()
	b, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, ts.URL+path, strings.NewReader(string(b)))
	req.Header.Set("Content-Type", "application/json")
	r, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Body.Close()
	raw, _ := io.ReadAll(r.Body)
	if v != nil {
		if err := json.Unmarshal(raw, v); err != nil {
			t.Fatalf("%s %s: %v\n%s", method, path, err, raw)
		}
	}
	return r.StatusCode
}

// TestSignalEditor goes through the editor's API: the file, an edit, an
// update, an addition, the errors, and reload.
func TestSignalEditor(t *testing.T) {
	ts, _, dir := newConsoleCopy(t)
	var f models.SignalFile
	get(t, ts, "/api/v1/signalfile", 200, &f)
	if f.Pending || len(f.Signals) < 10 || f.Signals[0].Name != "NGSetupRequest" || f.Signals[0].Line == 0 || !f.Signals[0].Loaded {
		t.Fatalf("file %+v", f)
	}
	const name = "InitialUEMeesageRegistrationRequest"
	var hexText string
	for _, s := range f.Signals {
		if s.Name == name {
			hexText = s.Hex
		}
	}

	var ed models.SignalEdit
	if code := send(t, ts, "POST", "/api/v1/signaledit", map[string]any{"name": name, "hex": hexText, "mode": "auto"}, &ed); code != 200 || ed.Hex != hexText || ed.Tree.Name != "InitialUEMessage" {
		t.Fatalf("decode: %d %+v", code, ed)
	}
	op := map[string]any{"kind": "set", "path": "RRCEstablishmentCause/Value", "value": "4"}
	if code := send(t, ts, "POST", "/api/v1/signaledit", map[string]any{"name": name, "hex": hexText, "mode": "auto", "op": op}, &ed); code != 200 || ed.Hex == hexText {
		t.Fatalf("edit: %d %+v", code, ed)
	}
	var reply map[string]string
	bad := map[string]any{"kind": "set", "path": "RANUENGAPID/Value", "value": "x"}
	if code := send(t, ts, "POST", "/api/v1/signaledit", map[string]any{"hex": hexText, "op": bad}, &reply); code != 400 || !strings.Contains(reply["error"], "not an integer") {
		t.Fatalf("bad edit: %d %v", code, reply)
	}

	if code := send(t, ts, "PUT", "/api/v1/signalfile/"+name, map[string]any{"hex": ed.Hex, "mode": "auto"}, &f); code != 200 || !f.Pending {
		t.Fatalf("update: %d %+v", code, f)
	}
	if code := send(t, ts, "POST", "/api/v1/signalfile", map[string]any{"name": "Copy", "hex": ed.Hex, "mode": "raw"}, &f); code != 200 {
		t.Fatalf("add: %d", code)
	}
	for _, c := range []struct {
		method, path string
		body         map[string]any
		code         int
	}{
		{"POST", "/api/v1/signalfile", map[string]any{"name": "Copy", "hex": ed.Hex}, 409},
		{"PUT", "/api/v1/signalfile/Nothing", map[string]any{"hex": ed.Hex}, 404},
		{"POST", "/api/v1/signalfile", map[string]any{"name": "a b", "hex": ed.Hex}, 400},
		{"POST", "/api/v1/signalfile", map[string]any{"name": "Bad", "hex": "zz"}, 400},
	} {
		if code := send(t, ts, c.method, c.path, c.body, nil); code != c.code {
			t.Errorf("%s %s %v: %d (want %d)", c.method, c.path, c.body, code, c.code)
		}
	}
	r, err := http.Post(ts.URL+"/api/v1/signalfile", "text/plain", strings.NewReader(`{"name":"X","hex":"00"}`))
	if err != nil || r.StatusCode != http.StatusUnsupportedMediaType {
		t.Fatalf("text/plain: %v %v", err, r)
	}
	r.Body.Close()
	src, _ := os.ReadFile(filepath.Join(dir, "signal.txt"))
	if !strings.Contains(string(src), "\nCopy [raw] {\n  "+ed.Hex+"\n}\n") {
		t.Fatalf("signal.txt:\n%s", src)
	}

	// A new signal from its type, with a NAS message chosen from the list.
	var types models.SignalTypes
	get(t, ts, "/api/v1/signaltypes", 200, &types)
	if !contains(types.NGAP, "InitialUEMessage") || !contains(types.GMM, "RegistrationRequest") || !contains(types.GSM, "PDUSessionEstablishmentRequest") {
		t.Fatalf("types %+v", types)
	}
	newOp := map[string]any{"kind": "new", "value": "InitialUEMessage"}
	if code := send(t, ts, "POST", "/api/v1/signaledit", map[string]any{"hex": "", "mode": "raw", "op": newOp}, &ed); code != 200 || ed.Tree.Name != "InitialUEMessage" {
		t.Fatalf("new: %d %+v", code, ed)
	}
	nasOp := map[string]any{"kind": "nas", "path": "NASPDU/Value", "value": "RegistrationRequest"}
	if code := send(t, ts, "POST", "/api/v1/signaledit", map[string]any{"hex": ed.Hex, "mode": "raw", "op": nasOp}, &ed); code != 200 {
		t.Fatalf("NAS: %d %+v", code, ed)
	}
	if code := send(t, ts, "POST", "/api/v1/signalfile", map[string]any{"name": "MyRegistration", "hex": ed.Hex, "mode": "raw"}, &f); code != 200 {
		t.Fatalf("add the new one: %d", code)
	}

	if code, reply := post(t, ts, "reload"); code != 200 || reply["error"] != "" {
		t.Fatalf("reload: %d %v", code, reply)
	}
	get(t, ts, "/api/v1/signalfile", 200, &f)
	if f.Pending {
		t.Fatalf("pending after reload %+v", f)
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
