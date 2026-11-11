package logger

import (
	"bytes"
	"context"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestCompatibilityFormatCallerAndLevels(t *testing.T) {
	t.Setenv("CORE_LOG_OUTPUT", "console")
	previous, previousLog, flags := slog.Default(), log.Writer(), log.Flags()
	mu.Lock()
	oldWriter, oldLevel := writer, currentLevel
	mu.Unlock()
	t.Cleanup(func() {
		slog.SetDefault(previous)
		log.SetOutput(previousLog)
		log.SetFlags(flags)
		mu.Lock()
		writer, currentLevel = oldWriter, oldLevel
		mu.Unlock()
	})
	if err := Configure("bridge", "trace"); err != nil {
		t.Fatal(err)
	}
	var output bytes.Buffer
	mu.Lock()
	writer = &output
	mu.Unlock()
	_, file, line, _ := runtime.Caller(0)
	slog.Warn("slog marker", "peer", "127.0.0.1:8805", "error", "two words")
	want := fmtCaller(file, line+1)
	if text := output.String(); !strings.Contains(text, "[WARN][logger.TestCompatibilityFormatCallerAndLevels]"+want+" slog marker peer=127.0.0.1:8805 error=\"two words\"") || strings.Contains(text, "level=WARN") {
		t.Fatal(text)
	}
	output.Reset()
	_, file, line, _ = runtime.Caller(0)
	log.Print("standard marker")
	if text := output.String(); !strings.Contains(text, "[INFO][logger.TestCompatibilityFormatCallerAndLevels]"+fmtCaller(file, line+1)+" standard marker") {
		t.Fatal(text)
	}
	output.Reset()
	child := slog.Default().With("nf", "smf").WithGroup("session").With("id", 1)
	child.Info("group marker", slog.Group("", slog.String("ip", "10.60.0.1")), slog.Group("empty"))
	if text := output.String(); !strings.Contains(text, "group marker nf=smf session.id=1 session.ip=10.60.0.1") || strings.Contains(text, "empty=") {
		t.Fatal(text)
	}
	output.Reset()
	SetLevel(LevelWarn)
	slog.Debug("hidden")
	log.Print("hidden")
	child.Info("hidden")
	slog.Warn("visible")
	if strings.Contains(output.String(), "hidden") || !strings.Contains(output.String(), "[WARN]") {
		t.Fatal(output.String())
	}
	output.Reset()
	SetLevel(LevelTrace)
	slog.Log(context.Background(), slog.LevelDebug-4, "trace marker")
	if !strings.Contains(output.String(), "[TRACE]") {
		t.Fatal(output.String())
	}
	// Logging and dynamic thresholds share the native wrapper's lock.
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 20; j++ {
				child.Warn("parallel")
				Warn("parallel")
			}
		}()
	}
	wg.Wait()
}

func fmtCaller(file string, line int) string {
	return "[" + filepath.Base(file) + ":" + strconv.Itoa(line) + "]"
}

func TestBridgeUsesCurrentWriterAfterClose(t *testing.T) {
	t.Setenv("CORE_LOG_OUTPUT", "console")
	oldDefault, oldLog, flags := slog.Default(), log.Writer(), log.Flags()
	oldWriter, oldLevel, oldStdout := writer, currentLevel, os.Stdout
	capture, err := os.CreateTemp(t.TempDir(), "stdout")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = Close()
		slog.SetDefault(oldDefault)
		log.SetOutput(oldLog)
		log.SetFlags(flags)
		writer, currentLevel, os.Stdout = oldWriter, oldLevel, oldStdout
		_ = capture.Close()
	})
	os.Stdout = capture
	if err := Configure("bridge", "info"); err != nil {
		t.Fatal(err)
	}
	f, err := os.CreateTemp(t.TempDir(), "closed-log")
	if err != nil {
		t.Fatal(err)
	}
	mu.Lock()
	logFile, writer = f, f
	mu.Unlock()
	if err := Close(); err != nil {
		t.Fatal(err)
	}
	slog.Error("after-close")
	data, err := os.ReadFile(capture.Name())
	if err != nil || !bytes.Contains(data, []byte("[ERROR]")) || !bytes.Contains(data, []byte("after-close")) {
		t.Fatalf("%s: %v", data, err)
	}
}
