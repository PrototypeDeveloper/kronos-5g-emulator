package engine_test

import (
	"bytes"
	"testing"

	libngap "libngap"
	"libngap/ngapMessage"

	"kronos/models"
	"kronos/scenario"
)

const mapScenario = `
node {
	entry { send(NGSetupRequest) }
}
start {
	entry { mapoff(ngapid) mapget(ngapid) }
}
`

// templateTAC returns the TAC of the sample NGSetupRequest.
func templateTAC(t *testing.T) []byte {
	t.Helper()
	sigs, ds := scenario.ReadSignals("../../../config/signal.txt")
	if ds.HasErrors() {
		t.Fatal(ds)
	}
	sig, _ := sigs.Lookup("NGSetupRequest")
	msg, _ := libngap.Decode(sig.Bytes)
	return msg.(*ngapMessage.NGSetupRequest).SupportedTAList.List[0].TAC.Value
}

func TestMappingCommands(t *testing.T) {
	hs := start(t, true, 2, mapScenario)
	exec := func(line, want string) {
		t.Helper()
		out, err := hs.eng.Exec(line)
		if err != nil || out != want {
			t.Fatalf("%s: %q, %v (want %q)", line, out, err, want)
		}
	}
	// Without target: the node and every subscriber; kept until linkup.
	exec("mapget", "mapping: every item on")
	exec("mapoff(tac)", "mapping off: tac")
	exec("mapget(tac)", "tac: off")
	req := hs.linkup(t)
	tmpl := templateTAC(t)
	if got := req.SupportedTAList.List[0].TAC.Value; !bytes.Equal(got, tmpl) || bytes.Equal(tmpl, []byte{0, 0, 7}) {
		t.Fatalf("tac off: TAC %x (template %x)", got, tmpl)
	}
	// The scenario turns ngapid off on each subscriber, which mapget reports.
	pending := map[string]bool{"208930000000001": true, "208930000000002": true}
	hs.wait(t, "mapget of both subscribers", func(e models.Event) bool {
		if e.Kind == models.KindInfo && e.Text == "ngapid: off" {
			delete(pending, e.Source)
		}
		return len(pending) == 0
	})
	exec("mapget(ngapid)", "ngapid: mixed (off on 2 of 3)")
	exec("mapget", "mapping: tac off, ngapid mixed (off on 2 of 3); the others on")
	// With target: that subscriber only.
	exec("target(208930000000001).mapon(ngapid)", "")
	exec("target(208930000000001).mapget(ngapid)", "")
	hs.wait(t, "mapget with target", func(e models.Event) bool {
		return e.Kind == models.KindInfo && e.Source == "208930000000001" && e.Text == "ngapid: on"
	})
	exec("mapget(ngapid)", "ngapid: mixed (off on 1 of 3)")
	exec("mapon", "mapping on: every item")
	exec("mapget", "mapping: every item on")
	if _, err := hs.eng.Exec("mapoff(nope)"); err == nil {
		t.Fatal("mapoff(nope) ran")
	}
	// linkdown keeps the switches.
	exec("mapoff(gnodebid)", "mapping off: gnodebid")
	if _, err := hs.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	exec("mapget(gnodebid)", "gnodebid: off")
}
