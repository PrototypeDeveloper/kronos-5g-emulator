package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"kronos/common/diag"
	"kronos/common/logger"
	"kronos/models"
	"kronos/simulator/engine"
	"kronos/simulator/gnb"
	"kronos/state"
	"kronos/traffic/tun"
	"kronos/webconsole"
)

// start runs the pseudo-call until the user leaves the CUI (Ctrl-D on a
// terminal), ctx is cancelled (Ctrl-C, SIGTERM) or, with --exec, the batch
// has run; a failed batch exits with 1.
func start(ctx context.Context, o options, stdin io.Reader, stdout, stderr io.Writer) int {
	cfg, sigs, scen, ds := load(o.config)
	// Errors are shown in full; warnings and notes stay in `gnodeb check`.
	var errs diag.List
	for _, d := range ds {
		if d.Sev == diag.Error {
			errs = append(errs, d)
		}
	}
	errs.Sort()
	errs.Print(stderr)
	if len(errs) > 0 {
		fmt.Fprintf(stderr, "not starting because of the errors (see %s check -v)\n", progName)
		return 1
	}
	if n := ds.Count(diag.Warning); n > 0 {
		fmt.Fprintf(stderr, "%d warning(s) (see %s check)\n", n, progName)
	}

	closeLog := setupLogger(cfg, o.version)
	defer closeLog()
	role, err := gnb.New(cfg)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	eng := engine.New(cfg, sigs, scen, role)
	bus := eng.Events()

	interactive := !o.noCUI && o.exec == "" && isTerminal(stdin)
	con := &console{out: stdout, interactive: interactive, prompt: progName + "> "}
	events, unsubscribe := bus.Subscribe(consoleEventBuffer)
	printed := make(chan struct{})
	go func() {
		for e := range events {
			con.line(formatEvent(e))
			logEvent(e)
		}
		close(printed)
	}()
	stopTrace := func() {}
	if path, stop, err := startTrace(bus, "log"); err != nil {
		con.line(fmt.Sprintf("cannot write the trace: %v", err))
	} else {
		stopTrace = stop
		con.line("trace: " + path)
	}
	// The web console: not in batch mode, nor with --no-web or
	// without gnodeb.web in the config. Listening first reports a busy port at once.
	var webListener net.Listener
	if w := cfg.WebListen(); w != nil && !o.noWeb && o.exec == "" {
		addr := net.JoinHostPort(w.Address, strconv.Itoa(w.Port))
		l, err := net.Listen("tcp", addr)
		if err != nil {
			con.line(fmt.Sprintf("cannot open the web console on %s: %v (--no-web starts without it)", addr, err))
			unsubscribe()
			<-printed
			stopTrace()
			return 1
		}
		webListener = l
		url := "http://" + l.Addr().String() + "/"
		if ip := net.ParseIP(w.Address); ip != nil && ip.IsUnspecified() {
			url = fmt.Sprintf("http://<this host>:%d/", l.Addr().(*net.TCPAddr).Port)
		}
		con.line("web console: " + url)
		logger.Info("web console: %s", url)
	}
	snap := "off"
	if cfg.Base.Snap {
		snap = "on"
	}
	hint := "linkup connects SCTP, help lists the commands, Ctrl-D quits"
	switch {
	case o.exec != "":
		hint = "running the batch " + o.exec + ", then exiting"
	case o.noCUI && webListener != nil:
		hint = "headless: linkup on the web console connects SCTP, Ctrl-C quits"
	case o.noCUI:
		hint = "headless, so SCTP connects by itself; Ctrl-C quits"
	case !interactive:
		hint = "linkup connects SCTP, Ctrl-C quits"
	}
	con.line(fmt.Sprintf("%s %s: %s (%d subscriber(s), snap %s). %s", progName, o.version, cfg.Path, cfg.UE.Subscription, snap, hint))
	if cfg.UE.TUN && !tun.HasNetAdmin() {
		con.line("ue.tun is on, but " + progName + " does not have CAP_NET_ADMIN, so the UEs get no TUN device: " +
			"run sudo setcap cap_net_admin+ep ./" + progName + " after each build, or start it as root")
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	eng.Start(ctx)
	webDone := make(chan struct{})
	if webListener != nil {
		web := webconsole.New(eng, models.WebOptions{Version: o.version})
		go func() {
			defer close(webDone)
			if err := web.Serve(ctx, webListener); err != nil {
				logger.Error("web console: %v", err)
			}
		}()
	} else {
		close(webDone)
	}
	if o.noCUI && o.exec == "" && webListener == nil {
		// Headless with no web console: nothing could send linkup, so connect at once.
		go func() {
			if err := eng.LinkUp(); err != nil {
				logger.Error("linkup: %v", err)
			}
		}()
	}

	code := 0
	if o.exec != "" {
		// Batch mode: the same runner as "@file" on the command line.
		if out, err := eng.Exec("@" + o.exec); err != nil {
			con.line("error: " + err.Error())
			code = 1
		} else {
			con.line(out)
		}
	} else {
		left := make(chan struct{})
		if !o.noCUI {
			go func() {
				con.readCommands(stdin, eng)
				if interactive { // Ctrl-D ends the session; piped input just ends the commands
					close(left)
				}
			}()
		}
		select {
		case <-left:
		case <-ctx.Done():
		}
	}
	cancel()
	<-webDone // the console stops with ctx
	eng.Wait()
	// The engine has stopped, so every event is already queued.
	unsubscribe()
	<-printed
	stopTrace()
	if n := bus.Dropped(); n > 0 {
		fmt.Fprintf(stderr, "events dropped by the display or the trace: %d\n", n)
	}
	if interactive {
		fmt.Fprintln(stdout)
	}
	return code
}

// setupLogger sends the application log to log/gnodeb.txt only: the console
// belongs to the CUI. CORE_LOG_OUTPUT still overrides the destination.
func setupLogger(cfg *state.Config, version string) func() {
	if os.Getenv("CORE_LOG_OUTPUT") == "" {
		os.Setenv("CORE_LOG_OUTPUT", "file")
	}
	level := "info"
	if cfg.Base.Debug {
		level = "debug"
	}
	if err := logger.Configure(progName, level); err != nil {
		fmt.Fprintln(os.Stderr, "cannot set up the log:", err)
		return func() {}
	}
	logger.Info("%s %s started (config %s)", progName, version, cfg.Path)
	return func() {
		logger.Info("%s stopped", progName)
		logger.Close()
	}
}

// logEvent copies link changes, warnings and errors to the application log.
func logEvent(e models.Event) {
	switch e.Kind {
	case models.KindLink:
		logger.Info("[%s] %s", e.Source, e.Text)
	case models.KindWarn:
		logger.Warn("[%s] %s", e.Source, e.Text)
	case models.KindError:
		logger.Error("[%s] %s %s", e.Source, e.Handler, e.Text)
	}
}

// startTrace writes every event as one JSON line to dir/trace-<time>.jsonl.
// The returned function flushes and closes the file.
func startTrace(bus *engine.Bus, dir string) (string, func(), error) {
	if err := os.MkdirAll(dir, traceDirPerm); err != nil {
		return "", nil, err
	}
	path := filepath.Join(dir, "trace-"+time.Now().Format("20060102-150405")+".jsonl")
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, traceFilePerm)
	if err != nil {
		return "", nil, err
	}
	events, unsubscribe := bus.Subscribe(traceEventBuffer)
	done := make(chan struct{})
	go func() {
		enc := json.NewEncoder(f)
		for e := range events {
			enc.Encode(e)
		}
		close(done)
	}()
	return path, func() {
		unsubscribe()
		<-done
		f.Close()
	}, nil
}

// isTerminal reports whether r is a character device such as a terminal.
func isTerminal(r io.Reader) bool {
	f, ok := r.(*os.File)
	if !ok {
		return false
	}
	fi, err := f.Stat()
	return err == nil && fi.Mode()&os.ModeCharDevice != 0
}
