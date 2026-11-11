package webconsole

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"kronos/models"
)

// events streams the event bus as Server-Sent Events.
//
// The stream starts with an event named "hello" ({"started", "version",
// "replay"}): a GUI that reconnects tells a restarted process by "started".
// Then come the kept events after Last-Event-ID (or ?since=), only the last
// ?tail= of them if given, then the live ones. Each event has its seq as id,
// so the browser resumes where it was after a reconnect. A client that reads
// too slowly loses events (the bus never waits); the gap shows in the seq.
func (s *Server) events(w http.ResponseWriter, r *http.Request) {
	fl, ok := w.(http.Flusher)
	if !ok {
		writeError(w, http.StatusInternalServerError, "this server cannot stream")
		return
	}
	// Where to resume: Last-Event-ID takes precedence over ?since=.
	var since uint64
	for _, v := range []string{r.Header.Get("Last-Event-ID"), r.URL.Query().Get("since")} {
		if v == "" {
			continue
		}
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, "since must be an event number (seq)")
			return
		}
		since = n
		break
	}
	tail, err := intParam(r, "tail", 0, maxListLimit)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	// The kept events and the live channel come together, so none is lost in between.
	hist, ch, cancel := s.b.Events().SubscribeSince(since, tail, eventBufferSize)
	defer cancel()

	// Headers, the hello event and the replayed history.
	h := w.Header()
	h.Set("Content-Type", "text/event-stream; charset=utf-8")
	h.Set("Cache-Control", "no-cache")
	h.Set("X-Accel-Buffering", "no")
	w.WriteHeader(http.StatusOK)
	hello, _ := json.Marshal(map[string]any{"started": s.b.Started(), "version": s.opts.Version, "replay": len(hist)})
	fmt.Fprintf(w, "retry: %d\nevent: hello\ndata: %s\n\n", sseRetryMillis, hello)
	for _, e := range hist {
		writeEvent(w, e)
	}
	fl.Flush()

	// Live events, with a comment line every 15 s to keep idle connections open.
	heartbeat := time.NewTicker(keepAliveInterval)
	defer heartbeat.Stop()
	for {
		select {
		case e, ok := <-ch:
			if !ok {
				return
			}
			writeEvent(w, e)
		drain: // send what is already queued with one flush
			for i := 0; i < maxDrainEvents; i++ {
				select {
				case e, ok := <-ch:
					if !ok {
						fl.Flush()
						return
					}
					writeEvent(w, e)
				default:
					break drain
				}
			}
			fl.Flush()
		case <-heartbeat.C:
			io.WriteString(w, ": ping\n\n")
			fl.Flush()
		case <-r.Context().Done():
			return
		case <-s.stop:
			return
		}
	}
}

// writeEvent writes one event as an SSE message with its seq as id.
func writeEvent(w io.Writer, e models.Event) {
	b, err := json.Marshal(e)
	if err != nil {
		return
	}
	fmt.Fprintf(w, "id: %d\ndata: %s\n\n", e.Seq, b)
}
