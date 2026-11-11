package scenario

import (
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"kronos/common/diag"
	"kronos/models"
)

// Writing signal.txt from the signal editor: one signal is replaced in
// place (the rest of the file, its comments and its order stay), or a new
// one is added at the end.

// ValidName reports whether name can be a signal name: letters, digits and
// '_', as the parser reads them.
func ValidName(name string) bool {
	if name == "" {
		return false
	}
	for _, c := range []byte(name) {
		if !(c == '_' || (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')) {
			return false
		}
	}
	return true
}

// ParseMode reads a send mode as the editor names it: auto (or ""),
// nas=keep, raw.
func ParseMode(s string) (Mode, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "", "auto":
		return ModeAuto, nil
	case "nas=keep":
		return ModeNASKeep, nil
	case "raw":
		return ModeRaw, nil
	}
	return 0, fmt.Errorf("unknown send mode %q (auto, nas=keep or raw)", s)
}

// signalBlock writes a signal as signal.txt has it, with the comment lines
// the old body had.
func signalBlock(name string, mode Mode, b []byte, comments []string) string {
	var sb strings.Builder
	sb.WriteString(name)
	switch mode {
	case ModeRaw:
		sb.WriteString(" [raw]")
	case ModeNASKeep:
		sb.WriteString(" [nas=keep]")
	}
	sb.WriteString(" {\n")
	for _, c := range comments {
		sb.WriteString("  " + c + "\n")
	}
	sb.WriteString("  " + hex.EncodeToString(b) + "\n}")
	return sb.String()
}

// bodyComments returns the comment lines inside a signal's braces.
func bodyComments(block string) []string {
	open := strings.IndexByte(block, '{')
	if open < 0 {
		return nil
	}
	var out []string
	for _, line := range strings.Split(block[open+1:], "\n") {
		if i := strings.IndexByte(line, '#'); i >= 0 {
			out = append(out, strings.TrimSpace(line[i:]))
		}
	}
	return out
}

// ErrNoSignal and ErrSignalExists are returned when the name is missing
// (update) or taken (add).
var (
	ErrNoSignal     = errors.New("no signal of that name in the file")
	ErrSignalExists = errors.New("a signal of that name is already in the file")
)

// ReplaceSignal returns src with the signal name rewritten to mode and b.
func ReplaceSignal(src []byte, name string, mode Mode, b []byte) ([]byte, error) {
	set, _ := ParseSignals("", src)
	sig, ok := set.Lookup(name)
	if !ok {
		return nil, fmt.Errorf("%s: %w", name, ErrNoSignal)
	}
	block := signalBlock(name, mode, b, bodyComments(string(src[sig.Start:sig.End])))
	out := make([]byte, 0, len(src)+len(block))
	out = append(out, src[:sig.Start]...)
	out = append(out, block...)
	return append(out, src[sig.End:]...), nil
}

// AppendSignal returns src with a new signal at the end.
func AppendSignal(src []byte, name string, mode Mode, b []byte) ([]byte, error) {
	if !ValidName(name) {
		return nil, fmt.Errorf("the signal name %q may have letters, digits and _ only", name)
	}
	set, _ := ParseSignals("", src)
	if _, dup := set.Lookup(name); dup {
		return nil, fmt.Errorf("%s: %w", name, ErrSignalExists)
	}
	out := append([]byte(nil), src...)
	if len(out) > 0 && out[len(out)-1] != '\n' {
		out = append(out, '\n')
	}
	if len(out) > 0 {
		out = append(out, '\n')
	}
	out = append(out, signalBlock(name, mode, b, nil)...)
	return append(out, '\n'), nil
}

// WriteFileAtomic replaces path with data through a temporary file in the
// same directory, keeping the file's permissions.
func WriteFileAtomic(path string, data []byte) error {
	perm := os.FileMode(newFilePerm)
	if st, err := os.Stat(path); err == nil {
		perm = st.Mode().Perm()
	}
	tmp, err := os.CreateTemp(filepath.Dir(path), "."+filepath.Base(path)+".*")
	if err != nil {
		return err
	}
	defer os.Remove(tmp.Name()) // after the rename there is nothing to remove
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	if err := os.Chmod(tmp.Name(), perm); err != nil {
		return err
	}
	return os.Rename(tmp.Name(), path)
}

// CheckSignal runs the checks of signal.txt on one signal: its decoding,
// its NAS and the comparison with the config.
func CheckSignal(name string, mode Mode, b []byte, ref models.SignalReference) diag.List {
	sig := &Signal{Name: name, Mode: mode, Bytes: b}
	var ds diag.List
	analyze(sig, &ds)
	set := &SignalSet{List: []*Signal{sig}, byName: map[string]*Signal{name: sig}}
	return append(ds, set.Compare(ref)...)
}
