// Package peer is the communication between pseudo-calls (sync_send /
// sync_receive): JSON over HTTP/REST, e.g. to hand a subscriber's mapping
// context over to the target of a handover.
package peer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/netip"

	"kronos/models"
)

// Errors of Deliver that the server answers with their own status.
var (
	ErrNotFound    = errors.New("no such subscriber")       // 404
	ErrUnavailable = errors.New("cannot take it right now") // 503
)

// NewServer returns a server that hands each valid message to deliver.
func NewServer(deliver func(models.SyncMessage) error) *Server {
	s := &Server{mux: http.NewServeMux(), deliver: deliver}
	s.mux.HandleFunc("POST "+SyncPath, s.sync)
	return s
}

// Serve answers on l until ctx ends.
func (s *Server) Serve(ctx context.Context, l net.Listener) error {
	hs := &http.Server{Handler: s.mux, ReadHeaderTimeout: readHeaderTimeout}
	go func() {
		<-ctx.Done()
		hs.Close()
	}()
	if err := hs.Serve(l); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// sync takes one message: 202 when it was handed over, else an error status.
func (s *Server) sync(w http.ResponseWriter, r *http.Request) {
	var m models.SyncMessage
	dec := json.NewDecoder(io.LimitReader(r.Body, maxBody))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&m); err != nil {
		reply(w, http.StatusBadRequest, fmt.Errorf("the body must be a sync message in JSON: %v", err))
		return
	}
	if err := Check(m); err != nil {
		reply(w, http.StatusBadRequest, err)
		return
	}
	err := s.deliver(m)
	switch {
	case err == nil:
		reply(w, http.StatusAccepted, nil)
	case errors.Is(err, ErrNotFound):
		reply(w, http.StatusNotFound, err)
	case errors.Is(err, ErrUnavailable):
		reply(w, http.StatusServiceUnavailable, err)
	default:
		reply(w, http.StatusBadRequest, err)
	}
}

// reply writes the JSON answer.
func reply(w http.ResponseWriter, status int, err error) {
	ans := models.SyncReply{Accepted: err == nil}
	if err != nil {
		ans.Error = err.Error()
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ans)
}

// Check reports what is wrong with a message, or nil.
func Check(m models.SyncMessage) error {
	switch {
	case m.Kind != models.SyncKindSync && m.Kind != models.SyncKindHO:
		return fmt.Errorf("kind must be %s or %s (found %q)", models.SyncKindSync, models.SyncKindHO, m.Kind)
	case m.IMSI != "" && !isIMSI(m.IMSI):
		return fmt.Errorf("imsi must be %d digits (found %q)", models.IMSIDigits, m.IMSI)
	case m.Kind == models.SyncKindHO && m.IMSI == "":
		return errors.New("a HO message is for a subscriber: imsi is needed")
	case m.Kind == models.SyncKindHO && m.Handover == nil:
		return errors.New("a HO message carries the handover context")
	}
	return nil
}

// isIMSI reports whether s is an IMSI: models.IMSIDigits digits.
func isIMSI(s string) bool {
	if len(s) != models.IMSIDigits {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// NewClient returns a client whose connections come from src (gnodeb.http.src; its port is not used).
func NewClient(src netip.AddrPort) *Client {
	d := &net.Dialer{Timeout: sendTimeout, LocalAddr: &net.TCPAddr{IP: src.Addr().AsSlice()}}
	tr := &http.Transport{DialContext: d.DialContext, MaxIdleConnsPerHost: 1}
	return &Client{http: &http.Client{Transport: tr, Timeout: sendTimeout}, from: src.String()}
}

// Send posts m to the pseudo-call at dst (host:port) and waits for its answer.
func (c *Client) Send(ctx context.Context, dst string, m models.SyncMessage) error {
	m.From = c.from
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://"+dst+SyncPath, bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("cannot reach %s: %v", dst, err)
	}
	defer resp.Body.Close()
	var ans models.SyncReply
	if err := json.NewDecoder(io.LimitReader(resp.Body, maxBody)).Decode(&ans); err != nil {
		return fmt.Errorf("%s answered %s, not a sync reply", dst, resp.Status)
	}
	if resp.StatusCode != http.StatusAccepted || !ans.Accepted {
		return fmt.Errorf("%s refused it (%s): %s", dst, resp.Status, ans.Error)
	}
	return nil
}
