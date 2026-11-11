package engine_test

import (
	"context"
	"net"
	"strings"
	"testing"
	"time"

	"libngap/ngapMessage"

	"kronos/models"
	"kronos/transport/peer"
)

// waitTexts waits for events with every one of texts, in any order.
func waitTexts(t *testing.T, hs *harness, texts ...string) {
	t.Helper()
	want := map[string]bool{}
	for _, s := range texts {
		want[s] = true
	}
	deadline := time.After(5 * time.Second)
	for len(want) > 0 {
		select {
		case e := <-hs.events:
			delete(want, e.Text)
		case <-deadline:
			t.Fatalf("no event %v", want)
		}
	}
}

// The commands that act on the node, written in the node block; and a
// timer of no length, an error when it runs.
func TestNodeBlockCommands(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	got := make(chan models.SyncMessage, 1)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go peer.NewServer(func(m models.SyncMessage) error { got <- m; return nil }).Serve(ctx, l)
	dst := l.Addr().String()
	src := freeUDPPort(t)
	scen := `
var { public n = 40ms public m = 40ms private t = 0s }
node {
	entry {
		send(NGSetupRequest)
		mapoff(gnodebid)
		mapget(gnodebid)
		timer(n)
		stoptimer(n)
		timer(m)
		sync_send(SYNC, ` + dst + `)
	}
	expire(n) { snap("n expired") }
}
start {
	entry { timer(t) }
}
`
	hs := startN3(t, true, 1, scen, freeUDPPort(t), freeUDPPort(t), httpLine(src))
	// linkup as hs.linkup does, without skipping the node entry's events.
	if err := hs.eng.LinkUp(); err != nil {
		t.Fatal(err)
	}
	hs.amf.accept(t)
	recvAs[*ngapMessage.NGSetupRequest](t, hs.amf)
	hs.amf.send(t, setupResponse())
	waitTexts(t, hs, "mapping off: gnodebid", "gnodebid: off", "timer n started (40ms)", "timer n stopped",
		"sync_send(SYNC) to "+dst+": accepted", "timer m expired: no expire(m) in node",
		"timer(t): the duration of t must be positive (0s)")
	if m := <-got; m.Kind != models.SyncKindSync || m.IMSI != "" || !strings.HasPrefix(m.From, "127.0.0.1:") {
		t.Fatalf("the node's SYNC %+v", m)
	}
	none(t, hs, 150*time.Millisecond, "n expired")
}

// Without gnodeb.http there is no pseudo-call link to send on.
func TestSyncSendWithoutHTTP(t *testing.T) {
	hs := start(t, true, 1, "node { entry { send(NGSetupRequest) } }\nstart { }")
	if _, err := hs.eng.Exec("target(208930000000001).sync_send(SYNC, 127.0.0.1:1)"); err == nil || !strings.Contains(err.Error(), "gnodeb.http.src is not set") {
		t.Fatalf("sync_send without gnodeb.http: %v", err)
	}
}
