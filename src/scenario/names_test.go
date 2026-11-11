package scenario

import (
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

// TestNGAPMessageTable checks ngapMessageTypes against the messages of
// libngap (one file per message type).
func TestNGAPMessageTable(t *testing.T) {
	files, err := filepath.Glob("../lib/ngap/ngapMessage/*.go")
	if err != nil || len(files) == 0 {
		t.Skipf("libngap sources not found: %v", err)
	}
	have := map[string]bool{}
	for _, p := range ngapMessageTypes {
		have[reflect.TypeOf(p).Elem().Name()] = true
	}
	var missing []string
	for _, f := range files {
		name := strings.TrimSuffix(filepath.Base(f), ".go")
		if strings.HasSuffix(name, "_test") {
			continue
		}
		if !have[name] {
			missing = append(missing, name)
		}
		delete(have, name)
	}
	if len(missing) > 0 || len(have) > 0 {
		t.Fatalf("ngap_messages.go is out of date: add %v, remove %v", missing, have)
	}
}
