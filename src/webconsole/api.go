package webconsole

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"kronos/command"
	"kronos/models"
	"kronos/simulator/engine"
)

// The REST part of the API. Reads return snapshots; the only
// write is POST /commands, which runs a command line like the CUI does.

// writeJSON sends v as an uncached JSON reply with status code.
func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

// writeError sends {"error": msg} with status code.
func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

// intParam reads a non-negative integer query parameter.
func intParam(r *http.Request, name string, def, max int) (int, error) {
	s := r.URL.Query().Get(name)
	if s == "" {
		return def, nil
	}
	n, err := strconv.Atoi(s)
	if err != nil || n < 0 {
		return 0, errors.New(name + " must be an integer of 0 or more")
	}
	if max > 0 && n > max {
		n = max
	}
	return n, nil
}

// status serves GET /api/v1/status: the dashboard overview and the version.
func (s *Server) status(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, statusReply{Overview: s.b.Overview(), Version: s.opts.Version})
}

// ues serves GET /api/v1/ues: a page of subscribers (?offset, ?limit, ?state).
func (s *Server) ues(w http.ResponseWriter, r *http.Request) {
	offset, err := intParam(r, "offset", 0, 0)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	limit, err := intParam(r, "limit", defaultListLimit, maxListLimit)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	items, total := s.b.UEs(offset, limit, r.URL.Query().Get("state"))
	if items == nil {
		items = []models.UEStatus{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"total": total, "offset": offset, "items": items})
}

// ue serves GET /api/v1/ues/{imsi}: the status of one subscriber.
func (s *Server) ue(w http.ResponseWriter, r *http.Request) {
	st, ok := s.b.UE(r.PathValue("imsi"))
	if !ok {
		writeError(w, http.StatusNotFound, "there is no subscriber with IMSI "+r.PathValue("imsi"))
		return
	}
	writeJSON(w, http.StatusOK, st)
}

// signals serves GET /api/v1/signals: the signals in use, in file order.
func (s *Server) signals(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.b.Signals())
}

// signal serves GET /api/v1/signals/{name}: one signal with its PDU decoded.
func (s *Server) signal(w http.ResponseWriter, r *http.Request) {
	si, tree, err := s.b.DecodeSignal(r.PathValue("name"))
	if si.Name == "" {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	reply := treeReply{Signal: &si, Tree: tree}
	if err != nil {
		reply.Error = err.Error()
	}
	writeJSON(w, http.StatusOK, reply)
}

// scenario serves GET /api/v1/scenario: the variables and blocks of the scenario.
func (s *Server) scenario(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.b.Scenario())
}

// vars serves GET /api/v1/vars: the scenario variables and their values.
func (s *Server) vars(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, s.b.Vars())
}

var kindNames = map[command.Kind]string{
	command.KindSignal: "signal", command.KindState: "state", command.KindFormat: "format",
	command.KindValue: "value", command.KindCommand: "command", command.KindRate: "rate", command.KindChoice: "choice",
}

// commandList serves GET /api/v1/commands: every registered command.
func (s *Server) commandList(w http.ResponseWriter, r *http.Request) {
	out := []commandInfo{}
	for _, sp := range command.All() {
		ci := commandInfo{Name: sp.Name, Help: sp.Help, Args: []argInfo{}}
		for _, a := range sp.Args {
			ci.Args = append(ci.Args, argInfo{Name: a.Name, Kind: kindNames[a.Kind], Optional: a.Optional, Variadic: a.Variadic, Choices: a.Choices, Default: a.Default})
		}
		for _, sc := range []struct {
			bit  command.Scope
			name string
		}{{command.ScopeNode, "node"}, {command.ScopeUE, "ue"}, {command.ScopeConsole, "console"}, {command.ScopeBySignal, "signal"}} {
			if sp.Scope&sc.bit != 0 {
				ci.Scope = append(ci.Scope, sc.name)
			}
		}
		if sp.Where&command.InScenario != 0 {
			ci.Where = append(ci.Where, "scenario")
		}
		if sp.Where&command.Interactive != 0 {
			ci.Where = append(ci.Where, "interactive")
		}
		out = append(out, ci)
	}
	writeJSON(w, http.StatusOK, out)
}

// readJSON reads a JSON request body into v. A JSON body cannot be sent
// cross-site without a CORS preflight, which this server never grants:
// other sites' pages cannot run commands or write signal.txt.
func readJSON(w http.ResponseWriter, r *http.Request, v any, shape string) bool {
	if ct := r.Header.Get("Content-Type"); !strings.HasPrefix(ct, "application/json") {
		writeError(w, http.StatusUnsupportedMediaType, "Content-Type must be application/json")
		return false
	}
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxBodySize)).Decode(v); err != nil {
		writeError(w, http.StatusBadRequest, "send JSON "+shape+": "+err.Error())
		return false
	}
	return true
}

// command runs one command line (POST {"line": "..."}). Parse and check
// errors and the result of console commands come back in the reply; what
// subscribers do comes as events. The line and any error are also events,
// so the trace records what was done from the GUI.
func (s *Server) command(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Line string `json:"line"`
	}
	if !readJSON(w, r, &req, `{"line": "…"}`) {
		return
	}
	line := strings.TrimSpace(req.Line)
	switch {
	case line == "":
		writeError(w, http.StatusBadRequest, "the command is empty")
		return
	case strings.ContainsAny(line, "\r\n"):
		writeError(w, http.StatusBadRequest, "send one line (several lines go in a batch file, @file)")
		return
	}
	bus := s.b.Events()
	bus.Publish(models.Event{Kind: models.KindCommand, Source: models.SourceWeb, Text: line})
	out, err := s.b.Exec(line)
	reply := commandReply{Line: line, Output: out}
	if err != nil {
		reply.Error = err.Error()
		bus.Publish(models.Event{Kind: models.KindError, Source: models.SourceWeb, Text: line + ": " + err.Error()})
	}
	writeJSON(w, http.StatusOK, reply)
}

// trace returns kept events: those after since, of one source (an IMSI,
// node or web) if ue is given, the last limit of them.
func (s *Server) trace(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	var since uint64
	if v := q.Get("since"); v != "" {
		n, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, "since must be an event number (seq)")
			return
		}
		since = n
	}
	limit, err := intParam(r, "limit", defaultListLimit, maxListLimit)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	evs := s.b.Events().History(since)
	if ue := q.Get("ue"); ue != "" {
		kept := evs[:0]
		for _, e := range evs {
			if e.Source == ue {
				kept = append(kept, e)
			}
		}
		evs = kept
	}
	if len(evs) > limit {
		evs = evs[len(evs)-limit:]
	}
	if evs == nil {
		evs = []models.Event{}
	}
	writeJSON(w, http.StatusOK, map[string]any{"events": evs})
}

// traceEvent returns one kept event with its PDU decoded.
func (s *Server) traceEvent(w http.ResponseWriter, r *http.Request) {
	seq, err := strconv.ParseUint(r.PathValue("seq"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "give an event number (seq)")
		return
	}
	ev, tree, err := s.b.DecodeEvent(seq)
	if ev.Seq == 0 {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}
	reply := treeReply{Event: &ev, Tree: tree}
	if err != nil {
		reply.Error = err.Error()
	}
	writeJSON(w, http.StatusOK, reply)
}

// The signal editor: signal.txt on disk, one edit at a time (the API keeps
// no state), and saving. reload (a command) makes the file's signals used.

// signalFile serves GET /api/v1/signalfile: signal.txt as it is on disk.
func (s *Server) signalFile(w http.ResponseWriter, r *http.Request) {
	f, err := s.b.SignalFile()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, f)
}

// editSignal serves POST /api/v1/signaledit: decodes a signal and applies an edit, without saving.
func (s *Server) editSignal(w http.ResponseWriter, r *http.Request) {
	var req signalRequest
	if !readJSON(w, r, &req, `{"name", "hex", "mode", "op": {"kind", "path", "value", "bits"}}`) {
		return
	}
	ed, err := s.b.EditSignal(req.Name, req.Hex, req.Mode, req.Op)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, ed)
}

// saveSignal writes a signal to signal.txt and maps the engine's errors to HTTP statuses.
func (s *Server) saveSignal(w http.ResponseWriter, name, hex, mode string, add bool) {
	f, err := s.b.SaveSignal(name, hex, mode, add)
	switch {
	case err == nil:
		writeJSON(w, http.StatusOK, f)
	case errors.Is(err, engine.ErrNoSignal):
		writeError(w, http.StatusNotFound, err.Error())
	case errors.Is(err, engine.ErrSignalExists):
		writeError(w, http.StatusConflict, err.Error())
	case errors.Is(err, engine.ErrBadSignal):
		writeError(w, http.StatusBadRequest, err.Error())
	default:
		writeError(w, http.StatusInternalServerError, err.Error())
	}
}

// updateSignal serves PUT /api/v1/signalfile/{name}: replaces a signal in signal.txt.
func (s *Server) updateSignal(w http.ResponseWriter, r *http.Request) {
	var req signalRequest
	if !readJSON(w, r, &req, `{"hex", "mode"}`) {
		return
	}
	s.saveSignal(w, r.PathValue("name"), req.Hex, req.Mode, false)
}

// addSignal serves POST /api/v1/signalfile: appends a new signal to signal.txt.
func (s *Server) addSignal(w http.ResponseWriter, r *http.Request) {
	var req signalRequest
	if !readJSON(w, r, &req, `{"name", "hex", "mode"}`) {
		return
	}
	s.saveSignal(w, req.Name, req.Hex, req.Mode, true)
}
