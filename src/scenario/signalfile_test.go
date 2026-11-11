package scenario

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

const fileSample = `# signals
First {
  0011
}

# the second one
Second [nas=keep] {
  # a note in the body
  00 22
  33
}
Third{44}`

func TestReplaceAndAppendSignal(t *testing.T) {
	out, err := ReplaceSignal([]byte(fileSample), "Second", ModeRaw, []byte{0xab, 0xcd})
	if err != nil {
		t.Fatal(err)
	}
	want := `# signals
First {
  0011
}

# the second one
Second [raw] {
  # a note in the body
  abcd
}
Third{44}`
	if string(out) != want {
		t.Fatalf("replace:\n%s", out)
	}
	out, err = ReplaceSignal(out, "Third", ModeAuto, []byte{0x55})
	if err != nil || !bytes.HasSuffix(out, []byte("}\n\n# the second one\nSecond [raw] {\n  # a note in the body\n  abcd\n}\nThird {\n  55\n}")) {
		t.Fatalf("replace the last: %v\n%s", err, out)
	}
	if _, err := ReplaceSignal(out, "Fourth", ModeAuto, []byte{1}); !errors.Is(err, ErrNoSignal) {
		t.Fatalf("replace a missing signal: %v", err)
	}

	out, err = AppendSignal(out, "Fourth", ModeNASKeep, []byte{0x66, 0x77})
	if err != nil || !bytes.HasSuffix(out, []byte("Third {\n  55\n}\n\nFourth [nas=keep] {\n  6677\n}\n")) {
		t.Fatalf("append: %v\n%s", err, out)
	}
	if _, err := AppendSignal(out, "First", ModeAuto, []byte{1}); !errors.Is(err, ErrSignalExists) {
		t.Fatalf("append a taken name: %v", err)
	}
	if _, err := AppendSignal(out, "Bad-Name", ModeAuto, []byte{1}); err == nil {
		t.Fatal("append with a bad name")
	}
	set, ds := ParseSignals("x", out)
	if len(set.List) != 4 {
		t.Fatalf("parsed %d signals: %v", len(set.List), ds)
	}
	if s, _ := set.Lookup("Second"); s.Mode != ModeRaw || !bytes.Equal(s.Bytes, []byte{0xab, 0xcd}) {
		t.Fatalf("Second %+v", s)
	}
	if s, _ := set.Lookup("Fourth"); s.Mode != ModeNASKeep || s.Pos.Line != 15 {
		t.Fatalf("Fourth %+v", s)
	}
}

func TestParseMode(t *testing.T) {
	for in, want := range map[string]Mode{"": ModeAuto, "auto": ModeAuto, "nas=keep": ModeNASKeep, "RAW": ModeRaw} {
		if m, err := ParseMode(in); err != nil || m != want {
			t.Errorf("%q: %v %v", in, m, err)
		}
	}
	if _, err := ParseMode("keep"); err == nil {
		t.Error("keep accepted")
	}
}

func TestWriteFileAtomic(t *testing.T) {
	path := filepath.Join(t.TempDir(), "signal.txt")
	if err := os.WriteFile(path, []byte("old"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := WriteFileAtomic(path, []byte("new")); err != nil {
		t.Fatal(err)
	}
	b, _ := os.ReadFile(path)
	st, _ := os.Stat(path)
	entries, _ := os.ReadDir(filepath.Dir(path))
	if string(b) != "new" || st.Mode().Perm() != 0o600 || len(entries) != 1 {
		t.Fatalf("%q %v %d files", b, st.Mode(), len(entries))
	}
}
