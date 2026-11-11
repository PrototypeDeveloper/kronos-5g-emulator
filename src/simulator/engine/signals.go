package engine

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"

	"kronos/common/diag"
	"kronos/models"
	"kronos/scenario"
	"kronos/state"
)

// The signals in use, the signal editor (signal.txt on disk) and reload.

// signals returns the signals in use (reload replaces them).
func (e *Engine) signals() *scenario.SignalSet { return e.sigs.Load() }

// Reference returns what the signals are compared with (gnodeb check).
func Reference(cfg *state.Config) models.SignalReference {
	g := cfg.GNodeB
	ref := models.SignalReference{MCC: g.PLMN.MCC, MNC: g.PLMN.MNC, TAC: g.TAC.Value, HasTAC: g.TAC.Set, Name: g.Name}
	if id, err := g.GNBIDValue(); err == nil {
		ref.GNBID, ref.GNBIDLen, ref.HasGNBID = id, g.GNodeBIDLength, true
	}
	if nci, err := g.NCIValue(); err == nil {
		ref.NCI, ref.HasNCI = nci, true
	}
	return ref
}

// diagTexts converts diagnostics for the web console.
func diagTexts(ds diag.List) []models.DiagText {
	out := make([]models.DiagText, 0, len(ds))
	for _, d := range ds {
		out = append(out, models.DiagText{Sev: d.Sev.String(), Line: d.Pos.Line, Msg: d.Msg})
	}
	return out
}

// errorsOf joins the error messages of ds.
func errorsOf(ds diag.List) string {
	var msgs []string
	for _, d := range ds {
		if d.Sev == diag.Error {
			msgs = append(msgs, d.String())
		}
	}
	return strings.Join(msgs, "; ")
}

// Reload reads signal.txt again and uses its signals from now on. With an
// error in the file, or a signal the scenario needs that no longer fits,
// nothing changes. A signal being sent keeps its old template.
func (e *Engine) Reload() error {
	e.fileMu.Lock()
	defer e.fileMu.Unlock()
	path := e.cfg.SignalPath()
	set, ds := scenario.ReadSignals(path)
	if ds.HasErrors() {
		return fmt.Errorf("%s has errors, so nothing is reloaded: %s", path, errorsOf(ds))
	}
	if vs := scenario.Validate(e.scen, set, models.ValidateOptions{SnapEnabled: e.snap}); vs.HasErrors() {
		return fmt.Errorf("the scenario does not fit the signals of %s, so nothing is reloaded: %s", path, errorsOf(vs))
	}
	warnings := append(ds, set.Compare(Reference(e.cfg))...)
	old := e.signals()
	changed, added, removed := 0, 0, 0
	for _, s := range set.List {
		o, ok := old.Lookup(s.Name)
		switch {
		case !ok:
			added++
		case o.Mode != s.Mode || !bytes.Equal(o.Bytes, s.Bytes):
			changed++
		}
	}
	for _, o := range old.List {
		if _, ok := set.Lookup(o.Name); !ok {
			removed++
		}
	}
	e.sigs.Store(set)
	text := fmt.Sprintf("signals reloaded from %s: %d (%d changed, %d added, %d removed)", path, len(set.List), changed, added, removed)
	if n := warnings.Count(diag.Warning); n > 0 {
		text += fmt.Sprintf(", %d warning(s) (see gnodeb check)", n)
	}
	e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode, Text: text})
	return nil
}

// SignalFile reads signal.txt as it is on disk now, for the editor.
func (e *Engine) SignalFile() (models.SignalFile, error) {
	path := e.cfg.SignalPath()
	src, err := os.ReadFile(path)
	if err != nil {
		return models.SignalFile{}, err
	}
	set, ds := scenario.ParseSignals(path, src)
	ds = append(ds, set.Compare(Reference(e.cfg))...)
	ds.Sort()
	f := models.SignalFile{Path: path, Signals: []models.SignalFileEntry{}, Diags: diagTexts(ds)}
	inUse := e.signals()
	for _, s := range set.List {
		u, ok := inUse.Lookup(s.Name)
		loaded := ok && u.Mode == s.Mode && bytes.Equal(u.Bytes, s.Bytes)
		f.Pending = f.Pending || !loaded
		f.Signals = append(f.Signals, models.SignalFileEntry{Name: s.Name, Mode: s.Mode.String(), Line: s.Pos.Line,
			Hex: hex.EncodeToString(s.Bytes), Label: s.Label(), Loaded: loaded})
	}
	f.Pending = f.Pending || len(set.List) != len(inUse.List)
	return f, nil
}

// EditSignal decodes a signal for the editor and applies op when there is
// one. name only labels the checks.
func (e *Engine) EditSignal(name, hexText, modeText string, op *models.EditOp) (models.SignalEdit, error) {
	mode, err := scenario.ParseMode(modeText)
	if err != nil {
		return models.SignalEdit{}, err
	}
	b, err := hex.DecodeString(strings.Join(strings.Fields(hexText), ""))
	if err != nil {
		return models.SignalEdit{}, fmt.Errorf("the signal is not hex octets: %v", err)
	}
	b, tree, err := scenario.EditSignal(b, mode, op, e.role.AutoField)
	if err != nil {
		return models.SignalEdit{}, err
	}
	if name == "" {
		name = "signal"
	}
	return models.SignalEdit{Hex: hex.EncodeToString(b), Tree: tree, Diags: diagTexts(scenario.CheckSignal(name, mode, b, Reference(e.cfg)))}, nil
}

// Errors of SaveSignal that are the request's fault.
var (
	ErrNoSignal     = scenario.ErrNoSignal
	ErrSignalExists = scenario.ErrSignalExists
	ErrBadSignal    = errors.New("bad signal")
)

// SaveSignal writes a signal into signal.txt: it replaces the signal of
// that name, or with add, appends a new one (the name must be new). The
// signals in use stay until reload.
func (e *Engine) SaveSignal(name, hexText, modeText string, add bool) (models.SignalFile, error) {
	mode, err := scenario.ParseMode(modeText)
	if err != nil {
		return models.SignalFile{}, fmt.Errorf("%w: %v", ErrBadSignal, err)
	}
	b, err := hex.DecodeString(strings.Join(strings.Fields(hexText), ""))
	if err != nil || len(b) == 0 {
		return models.SignalFile{}, fmt.Errorf("%w: the signal must be hex octets", ErrBadSignal)
	}
	if !scenario.ValidName(name) {
		return models.SignalFile{}, fmt.Errorf("%w: the signal name %q may have letters, digits and _ only", ErrBadSignal, name)
	}
	if ds := scenario.CheckSignal(name, mode, b, Reference(e.cfg)); ds.HasErrors() {
		return models.SignalFile{}, fmt.Errorf("%w: %s", ErrBadSignal, errorsOf(ds))
	}
	e.fileMu.Lock()
	path := e.cfg.SignalPath()
	src, err := os.ReadFile(path)
	if err == nil {
		if add {
			src, err = scenario.AppendSignal(src, name, mode, b)
		} else {
			src, err = scenario.ReplaceSignal(src, name, mode, b)
		}
	}
	if err == nil {
		err = scenario.WriteFileAtomic(path, src)
	}
	e.fileMu.Unlock()
	if err != nil {
		return models.SignalFile{}, err
	}
	verb := "updated"
	if add {
		verb = "added"
	}
	e.publish(models.Event{Kind: models.KindInfo, Source: models.SourceNode,
		Text: fmt.Sprintf("%s: %s %s (reload sends it)", path, name, verb)})
	return e.SignalFile()
}
