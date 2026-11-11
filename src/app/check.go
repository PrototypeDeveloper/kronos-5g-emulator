package app

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"

	"kronos/command"
	"kronos/common/diag"
	"kronos/models"
	"kronos/scenario"
	"kronos/simulator/engine"
	"kronos/state"
)

// load reads and validates the three input files.
func load(cfgPath string) (*state.Config, *scenario.SignalSet, *scenario.File, diag.List) {
	var all diag.List
	cfg, ds := state.Load(cfgPath)
	all = append(all, ds...)
	if cfg == nil {
		return nil, nil, nil, all
	}
	sigs, ds := scenario.ReadSignals(cfg.SignalPath())
	all = append(all, ds...)
	all = append(all, sigs.Compare(engine.Reference(cfg))...)
	f, ds := scenario.ReadScenario(cfg.ScenarioPath())
	all = append(all, ds...)
	if f != nil {
		all = append(all, scenario.Validate(f, sigs, models.ValidateOptions{SnapEnabled: cfg.Base.Snap, Sync: cfg.SyncListen() != nil})...)
	}
	return cfg, sigs, f, all
}

// check prints every diagnostic and exits with 1 when there is an error.
func check(cfgPath string, verbose bool, w io.Writer) int {
	cfg, sigs, f, all := load(cfgPath)
	if verbose && cfg != nil {
		describe(w, cfg, sigs, f)
	}
	all.Sort()
	all.Print(w)
	fmt.Fprintln(w, all.Summary())
	if all.HasErrors() {
		return 1
	}
	return 0
}

// describe prints what was understood from the files (check -v).
func describe(w io.Writer, cfg *state.Config, sigs *scenario.SignalSet, f *scenario.File) {
	first, _ := cfg.SubscriberIMSI(0)
	last, _ := cfg.SubscriberIMSI(cfg.UE.Subscription - 1)
	fmt.Fprintf(w, "subscribers: %d (IMSI %s to %s), snap: %v\n", cfg.UE.Subscription, first, last, cfg.Base.Snap)
	if sigs != nil {
		fmt.Fprintf(w, "\nsignals (%s): %d\n", sigs.Path, len(sigs.List))
		tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
		for _, s := range sigs.List {
			scope := "?"
			switch s.Scope {
			case command.ScopeNode:
				scope = "node"
			case command.ScopeUE:
				scope = "UE"
			}
			fmt.Fprintf(tw, "  %s\t%s\t%s\t%s\n", s.Name, scope, s.Mode, s.Label())
		}
		tw.Flush()
	}
	if f == nil {
		return
	}
	fmt.Fprintf(w, "\nscenario (%s)\n", f.Path)
	for _, v := range f.Vars {
		kind := "private"
		if v.Public {
			kind = "public"
		}
		fmt.Fprintf(w, "  var %s %s = %s\n", kind, v.Name, v.Init.Literal())
	}
	for _, b := range f.Blocks() {
		fmt.Fprintf(w, "  %s\n", b.Name)
		if b.Entry != nil {
			fmt.Fprintf(w, "    entry → %s\n", calls(b.Entry.Body))
		}
		for _, h := range b.Handlers {
			fmt.Fprintf(w, "    receive(%s) → %s\n", h.Trigger, calls(h.Body))
		}
		for _, x := range b.Expires {
			fmt.Fprintf(w, "    expire(%s) → %s\n", x.Var, calls(x.Body))
		}
		for _, x := range b.Syncs {
			fmt.Fprintf(w, "    sync_receive(%s) → %s\n", x.Kind, calls(x.Body))
		}
	}
	fmt.Fprintln(w)
}

// calls joins the items with ", ", or returns "(none)" for an empty list.
func calls[T fmt.Stringer](cs []T) string {
	if len(cs) == 0 {
		return "(none)"
	}
	out := make([]string, len(cs))
	for i, c := range cs {
		out[i] = c.String()
	}
	return strings.Join(out, ", ")
}
