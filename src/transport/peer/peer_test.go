package peer

import (
	"context"
	"net"
	"net/http"
	"net/netip"
	"strings"
	"testing"

	"kronos/models"
)

func TestServerAndClient(t *testing.T) {
	got := make(chan models.SyncMessage, 1)
	srv := NewServer(func(m models.SyncMessage) error {
		switch m.IMSI {
		case "208930000000099":
			return ErrNotFound
		case "208930000000098":
			return ErrUnavailable
		}
		got <- m
		return nil
	})
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go srv.Serve(ctx, l)
	dst := l.Addr().String()

	cli := NewClient(netip.MustParseAddrPort("127.0.0.1:12345"))
	ho := &models.HOContext{Registered: true, Sessions: []models.HOSession{{ID: 1, UEIP: "10.60.0.7", QFIs: []int{1}}}}
	if err := cli.Send(ctx, dst, models.SyncMessage{Kind: models.SyncKindHO, IMSI: "208930000000001", Handover: ho}); err != nil {
		t.Fatal(err)
	}
	m := <-got
	if m.From != "127.0.0.1:12345" || m.Handover == nil || m.Handover.Sessions[0].UEIP != "10.60.0.7" {
		t.Fatalf("delivered %+v", m)
	}
	for _, tc := range []struct {
		m    models.SyncMessage
		want string
	}{
		{models.SyncMessage{Kind: "SYNC", IMSI: "208930000000099"}, "404"},
		{models.SyncMessage{Kind: "SYNC", IMSI: "208930000000098"}, "503"},
		{models.SyncMessage{Kind: "HO", IMSI: "208930000000001"}, "carries the handover context"},
		{models.SyncMessage{Kind: "HO"}, "imsi is needed"},
		{models.SyncMessage{Kind: "SYNC", IMSI: "12"}, "imsi must be 15 digits"},
	} {
		if err := cli.Send(ctx, dst, tc.m); err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Errorf("%+v: want %q, got %v", tc.m, tc.want, err)
		}
	}
	// The body is JSON with known fields only.
	for body, want := range map[string]int{`{"kind":"SYNC","bogus":1}`: http.StatusBadRequest, `not json`: http.StatusBadRequest, `{"kind":"SYNC"}`: http.StatusAccepted} {
		resp, err := http.Post("http://"+dst+SyncPath, "application/json", strings.NewReader(body))
		if err != nil {
			t.Fatal(err)
		}
		resp.Body.Close()
		if resp.StatusCode != want {
			t.Errorf("%s: %s", body, resp.Status)
		}
	}
	<-got
	if resp, err := http.Get("http://" + dst + SyncPath); err != nil || resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("GET: %v %v", resp, err)
	}
	// Nothing listens there.
	if err := cli.Send(ctx, "127.0.0.1:1", models.SyncMessage{Kind: "SYNC"}); err == nil || !strings.Contains(err.Error(), "cannot reach") {
		t.Errorf("no peer: %v", err)
	}
}
