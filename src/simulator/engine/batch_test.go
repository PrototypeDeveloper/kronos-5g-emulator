package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"libngap/ngapMessage"

	"kronos/models"
)

func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
}

// TestBatch runs "@file" lines: in order, nested, with comments, and
// stopping at the first error with its file and line.
func TestBatch(t *testing.T) {
	hs := start(t, true, 2, "node { receive(NGSetupResponse) { snap(\"ok\") } }")
	dir := t.TempDir()
	main := filepath.Join(dir, "main.txt")
	writeFile(t, main, "# connect, then send NG Setup from a nested batch\nlinkup\n\nwait(100ms)\n@inner.txt   // relative to main.txt\nhelp\n")
	writeFile(t, filepath.Join(dir, "inner.txt"), "send(NGSetupRequest)\n")

	out, err := hs.eng.Exec("@" + main)
	if err != nil || !strings.Contains(out, "(4 lines)") {
		t.Fatalf("batch: %q %v", out, err)
	}
	hs.amf.accept(t)
	recvAs[*ngapMessage.NGSetupRequest](t, hs.amf)
	for _, want := range []string{main + ":2: linkup", main + ":5: @inner.txt", filepath.Join(dir, "inner.txt") + ":1: send(NGSetupRequest)", "send(signal)"} {
		hs.wait(t, want, func(e models.Event) bool { return e.Kind == models.KindCommand && strings.Contains(e.Text, want) })
	}

	// The first error stops the batch, with its position.
	bad := filepath.Join(dir, "bad.txt")
	writeFile(t, bad, "help\nsend(NoSuchSignal)\nhelp\n")
	if _, err := hs.eng.Exec("@" + bad); err == nil || !strings.Contains(err.Error(), bad+":2") {
		t.Fatalf("bad batch: %v", err)
	}
	// A command for every subscriber waits for all of them and fails if any failed.
	fan := filepath.Join(dir, "fan.txt")
	writeFile(t, fan, "send(UpLinkNASTransportAuthenticationResponse)\n")
	if _, err := hs.eng.Exec("@" + fan); err == nil || !strings.Contains(err.Error(), "2 subscriber(s) failed") {
		t.Fatalf("fan-out batch: %v", err)
	}
	self := filepath.Join(dir, "self.txt")
	writeFile(t, self, "@self.txt\n")
	if _, err := hs.eng.Exec("@" + self); err == nil || !strings.Contains(err.Error(), "nested too deep") {
		t.Fatalf("recursive batch: %v", err)
	}
	if _, err := hs.eng.Exec("@" + filepath.Join(dir, "missing.txt")); err == nil || !strings.Contains(err.Error(), "cannot read") {
		t.Fatalf("missing batch: %v", err)
	}
	if _, err := hs.eng.Exec("@"); err == nil {
		t.Fatal("@ without a file")
	}
}
