package engine_test

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"kronos/models"
	"kronos/simulator/engine"
)

func fileEntry(f models.SignalFile, name string) (models.SignalFileEntry, bool) {
	for _, s := range f.Signals {
		if s.Name == name {
			return s, true
		}
	}
	return models.SignalFileEntry{}, false
}

func signalHex(hs *harness, name string) string {
	for _, s := range hs.eng.Signals() {
		if s.Name == name {
			return s.Hex
		}
	}
	return ""
}

// The signal editor: read signal.txt, edit a signal, save it and add one,
// then reload makes them the signals in use.
func TestSignalEditorAndReload(t *testing.T) {
	hs := start(t, true, 1, sessionScenario)
	const name = "InitialUEMeesageRegistrationRequest"
	f, err := hs.eng.SignalFile()
	if err != nil || f.Pending || len(f.Signals) != len(hs.eng.Signals()) {
		t.Fatalf("file %+v %v", f, err)
	}
	orig, ok := fileEntry(f, name)
	if !ok || !orig.Loaded || orig.Mode != "auto" || orig.Label == "" {
		t.Fatalf("entry %+v", orig)
	}

	// Edit: the role marks what it replaces; the checks come with it.
	ed, err := hs.eng.EditSignal(name, orig.Hex, "auto", &models.EditOp{Kind: "set", Path: "RRCEstablishmentCause/Value", Value: "4"})
	if err != nil || ed.Hex == orig.Hex || ed.Tree == nil {
		t.Fatalf("edit %+v %v", ed, err)
	}
	if n := findNode(ed.Tree, "RANUENGAPID/Value"); n == nil || !n.Auto {
		t.Fatalf("RAN UE NGAP ID %+v", n)
	}
	if _, err := hs.eng.EditSignal(name, orig.Hex, "auto", &models.EditOp{Kind: "set", Path: "RANUENGAPID/Value", Value: "-1"}); err == nil {
		t.Fatal("an out of range value was taken")
	}
	// A mapped field is edited too (for a send without the mapping); it stays marked.
	mapped, err := hs.eng.EditSignal(name, orig.Hex, "auto", &models.EditOp{Kind: "set", Path: "RANUENGAPID/Value", Value: "7"})
	if n := findNode(mapped.Tree, "RANUENGAPID/Value"); err != nil || n.Value != "7" || !n.Auto {
		t.Fatalf("mapped field %+v %v", n, err)
	}

	// Save: the file changes, the signals in use do not.
	f, err = hs.eng.SaveSignal(name, ed.Hex, "auto", false)
	if e, _ := fileEntry(f, name); err != nil || !f.Pending || e.Loaded || e.Hex != ed.Hex {
		t.Fatalf("saved %+v %v", f, err)
	}
	hs.wait(t, "saved", func(e models.Event) bool { return strings.HasSuffix(e.Text, name+" updated (reload sends it)") })
	if signalHex(hs, name) != orig.Hex {
		t.Fatal("saving changed the signal in use")
	}
	if _, err := hs.eng.SaveSignal("NewRegistration", ed.Hex, "nas=keep", true); err != nil {
		t.Fatal(err)
	}
	if _, err := hs.eng.SaveSignal("Junk", "deadbeef", "raw", true); err != nil {
		t.Fatalf("a raw signal: %v", err)
	}
	for _, c := range []struct {
		name, hex, mode string
		add             bool
		want            error
	}{
		{"NewRegistration", ed.Hex, "auto", true, engine.ErrSignalExists},
		{"Missing", ed.Hex, "auto", false, engine.ErrNoSignal},
		{"bad-name", ed.Hex, "auto", true, engine.ErrBadSignal},
		{"Undecodable", "deadbeef", "auto", true, engine.ErrBadSignal},
		{"Odd", "abc", "auto", true, engine.ErrBadSignal},
		{"Mode", ed.Hex, "keep", true, engine.ErrBadSignal},
	} {
		if _, err := hs.eng.SaveSignal(c.name, c.hex, c.mode, c.add); !errors.Is(err, c.want) {
			t.Errorf("%s: %v (want %v)", c.name, err, c.want)
		}
	}
	src, _ := os.ReadFile(filepath.Join(hs.dir, "signal.txt"))
	if !strings.Contains(string(src), "\nNewRegistration [nas=keep] {\n  "+ed.Hex+"\n}\n") {
		t.Fatalf("signal.txt:\n%s", src)
	}

	// reload: the file's signals are the ones in use.
	if _, err := hs.eng.Exec("reload"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "reloaded", func(e models.Event) bool {
		return strings.HasPrefix(e.Text, "signals reloaded from ") && strings.Contains(e.Text, "(1 changed, 2 added, 0 removed)")
	})
	if f, _ := hs.eng.SignalFile(); f.Pending || signalHex(hs, name) != ed.Hex || signalHex(hs, "NewRegistration") != ed.Hex {
		t.Fatalf("after reload %+v", f)
	}

	// A file with errors, or one without a signal the scenario sends, is not taken.
	path := filepath.Join(hs.dir, "signal.txt")
	if err := os.WriteFile(path, append(src, []byte("\nBroken { zz }\n")...), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := hs.eng.Exec("reload"); err == nil || !strings.Contains(err.Error(), "has errors, so nothing is reloaded") {
		t.Fatalf("reload with errors: %v", err)
	}
	without := strings.Replace(string(src), "UplinkNASTransportPDUSessionEstablishmentRequest {", "Renamed {", 1)
	if err := os.WriteFile(path, []byte(without), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := hs.eng.Exec("reload"); err == nil || !strings.Contains(err.Error(), "the scenario does not fit") {
		t.Fatalf("reload without a signal of the scenario: %v", err)
	}
	if signalHex(hs, "UplinkNASTransportPDUSessionEstablishmentRequest") == "" {
		t.Fatal("a failed reload changed the signals")
	}
	if _, err := hs.eng.Exec("target(208930000000001).reload"); err == nil {
		t.Fatal("reload took a target")
	}
}

func findNode(n *models.EditNode, path string) *models.EditNode {
	if n.Path == path {
		return n
	}
	for _, k := range n.Kids {
		if f := findNode(k, path); f != nil {
			return f
		}
	}
	return nil
}
