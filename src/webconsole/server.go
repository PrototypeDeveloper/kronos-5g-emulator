// Package webconsole is the web console: the GUI, embedded in
// the binary, and the API it uses (REST and Server-Sent Events under
// /api/v1). Every operation goes through the same command line as the CUI,
// so the GUI has no function of its own.
//
// The console has no authentication and no TLS: it is a test tool, to be
// opened from any browser that can reach it.
package webconsole

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"net"
	"net/http"
	"time"

	"kronos/models"
	"kronos/simulator/engine"
)

//go:embed static
var static embed.FS

// Backend is what the web console uses of the engine (*engine.Engine).
type Backend interface {
	Exec(line string) (string, error)
	Events() *engine.Bus
	Started() time.Time
	Overview() models.Overview
	UEs(offset, limit int, state string) ([]models.UEStatus, int)
	UE(imsi string) (models.UEStatus, bool)
	Signals() []models.SignalInfo
	DecodeSignal(name string) (models.SignalInfo, *models.TreeNode, error)
	Scenario() models.ScenarioInfo
	Vars() []models.VarValue
	DecodeEvent(seq uint64) (models.Event, *models.TreeNode, error)
	SignalFile() (models.SignalFile, error)
	EditSignal(name, hex, mode string, op *models.EditOp) (models.SignalEdit, error)
	SaveSignal(name, hex, mode string, add bool) (models.SignalFile, error)
	SignalTypes() models.SignalTypes
}

var _ Backend = (*engine.Engine)(nil)

// New prepares the console for b.
func New(b Backend, opts models.WebOptions) *Server {
	s := &Server{b: b, opts: opts, mux: http.NewServeMux(), stop: make(chan struct{})}
	s.routes()
	return s
}

// Handler is the HTTP handler of the console (GUI and /api/v1).
func (s *Server) Handler() http.Handler { return s.mux }

// routes registers the GUI files and the API handlers.
func (s *Server) routes() {
	// The GUI: the embedded static files, GET and HEAD only.
	files, _ := fs.Sub(static, "static")
	fileServer := http.FileServerFS(files)
	s.mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			w.Header().Set("Allow", "GET, HEAD")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		w.Header().Set("Cache-Control", "no-cache") // a rebuilt binary shows its own GUI
		fileServer.ServeHTTP(w, r)
	}))
	// The API.
	s.mux.HandleFunc("GET /api/v1/status", s.status)
	s.mux.HandleFunc("GET /api/v1/ues", s.ues)
	s.mux.HandleFunc("GET /api/v1/ues/{imsi}", s.ue)
	s.mux.HandleFunc("GET /api/v1/signals", s.signals)
	s.mux.HandleFunc("GET /api/v1/signals/{name}", s.signal)
	s.mux.HandleFunc("GET /api/v1/signalfile", s.signalFile)
	s.mux.HandleFunc("POST /api/v1/signalfile", s.addSignal)
	s.mux.HandleFunc("PUT /api/v1/signalfile/{name}", s.updateSignal)
	s.mux.HandleFunc("POST /api/v1/signaledit", s.editSignal)
	s.mux.HandleFunc("GET /api/v1/signaltypes", s.signalTypes)
	s.mux.HandleFunc("GET /api/v1/scenario", s.scenario)
	s.mux.HandleFunc("GET /api/v1/vars", s.vars)
	s.mux.HandleFunc("GET /api/v1/commands", s.commandList)
	s.mux.HandleFunc("POST /api/v1/commands", s.command)
	s.mux.HandleFunc("GET /api/v1/trace", s.trace)
	s.mux.HandleFunc("GET /api/v1/trace/{seq}", s.traceEvent)
	s.mux.HandleFunc("GET /api/v1/events", s.events)
	// Any other API path gets a JSON 404.
	s.mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		writeError(w, http.StatusNotFound, "no such API path: "+r.Method+" "+r.URL.Path)
	})
}

// Serve answers on l until ctx is cancelled, then shuts down: the event
// streams end at once, other requests get a moment to finish.
func (s *Server) Serve(ctx context.Context, l net.Listener) error {
	hs := &http.Server{Handler: s.mux, ReadHeaderTimeout: readHeaderTimeout}
	hs.RegisterOnShutdown(s.shutdown)
	done := make(chan error, 1)
	go func() { done <- hs.Serve(l) }()
	select {
	case err := <-done:
		s.shutdown()
		return err
	case <-ctx.Done():
	}
	sctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	err := hs.Shutdown(sctx)
	if e := <-done; !errors.Is(e, http.ErrServerClosed) && err == nil {
		err = e
	}
	return err
}

// shutdown ends the event streams; it may be called more than once.
func (s *Server) shutdown() { s.stopOnce.Do(func() { close(s.stop) }) }
