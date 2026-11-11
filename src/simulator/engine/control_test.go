package engine_test

import (
	"sort"
	"strings"
	"testing"
	"time"

	"libngap/ngapMessage"

	"kronos/models"
)

// Each subscriber loops through start three times; control(20) lets one
// pass every 50 ms, counted over both subscribers.
const controlScenario = `
var { private n = 0 }
node { entry { send(NGSetupRequest) } }
start {
	entry {
		control(20)
		snap("pass")
		set(n = n + 1)
		if (n < 3) { move(start) }
	}
}
`

func TestControlPaces(t *testing.T) {
	hs := start(t, true, 2, controlScenario)
	hs.linkup(t)
	var at []time.Time
	deadline := time.After(5 * time.Second)
	for len(at) < 6 {
		select {
		case e := <-hs.events:
			if e.Kind == models.KindSnap && e.Text == "pass" {
				at = append(at, e.Time)
			}
		case <-deadline:
			t.Fatalf("%d passes", len(at))
		}
	}
	// The turns are 50 ms apart; a late wake-up may shorten one gap, but
	// none comes close to none, and the six take about 250 ms.
	sort.Slice(at, func(i, j int) bool { return at[i].Before(at[j]) })
	for i := 1; i < len(at); i++ {
		if gap := at[i].Sub(at[i-1]); gap < 25*time.Millisecond {
			t.Fatalf("pass %d came %v after the one before (want 50 ms)", i, gap)
		}
	}
	if total := at[5].Sub(at[0]); total < 200*time.Millisecond || total > time.Second {
		t.Fatalf("6 passes took %v (want about 250 ms)", total)
	}
}

// linkdown ends the waits: the rest of the entry does not run, and the
// turns start afresh at the next linkup.
func TestControlLinkdown(t *testing.T) {
	hs := start(t, true, 2, strings.Replace(controlScenario, "control(20)", "control(0.5)", 1))
	hs.linkup(t)
	hs.wait(t, "the first pass", func(e models.Event) bool { return e.Kind == models.KindSnap })
	// The next turn is 2 s away; linkdown cuts the wait short.
	begin := time.Now()
	if _, err := hs.eng.Exec("linkdown"); err != nil {
		t.Fatal(err)
	}
	hs.wait(t, "the reset", func(e models.Event) bool { return strings.Contains(e.Text, "back to the initial state") })
	if d := time.Since(begin); d > time.Second {
		t.Fatalf("linkdown took %v", d)
	}
	// linkup as hs.linkup does, but without skipping the events after NG Setup.
	if err := hs.eng.LinkUp(); err != nil {
		t.Fatal(err)
	}
	hs.amf.accept(t)
	recvAs[*ngapMessage.NGSetupRequest](t, hs.amf)
	hs.amf.send(t, setupResponse())
	got := hs.wait(t, "a pass after linkup", func(e models.Event) bool {
		return e.Kind == models.KindSnap || e.Kind == models.KindError
	})
	if got.Kind != models.KindSnap || time.Since(begin) > 1500*time.Millisecond {
		t.Fatalf("after linkup: %+v (%v)", got, time.Since(begin))
	}
}
