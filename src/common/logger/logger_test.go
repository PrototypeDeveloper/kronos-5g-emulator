package logger

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestLogCaller(t *testing.T) {
	var output bytes.Buffer
	mu.Lock()
	previousWriter, previousLevel := writer, currentLevel
	writer, currentLevel = &output, LevelTrace
	mu.Unlock()
	t.Cleanup(func() {
		mu.Lock()
		writer, currentLevel = previousWriter, previousLevel
		mu.Unlock()
	})
	for _, tt := range []struct {
		label string
		log   func(string, ...any)
	}{
		{"TRACE", Trace}, {"DEBUG", Debug}, {"INFO", Info},
		{"WARN", Warn}, {"ERROR", Error},
	} {
		t.Run(tt.label, func(t *testing.T) {
			output.Reset()
			pc, file, line, _ := runtime.Caller(0)
			tt.log("caller marker %d", 42)
			function := shortFunctionName(runtime.FuncForPC(pc).Name())
			want := fmt.Sprintf("[%s][%s][%s:%d] caller marker 42\n", tt.label, function, filepath.Base(file), line+1)
			if !strings.HasSuffix(output.String(), want) {
				t.Fatalf("log = %q, want suffix %q", output.String(), want)
			}
		})
	}
}

func TestParseLevel(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Level
	}{
		{"trace", "trace", LevelTrace},
		{"debug", "debug", LevelDebug},
		{"info", "info", LevelInfo},
		{"warn", "warn", LevelWarn},
		{"warning", "warning", LevelWarn},
		{"error", "error", LevelError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLevel(tt.input)
			if err != nil {
				t.Fatalf("ParseLevel() error = %v", err)
			}
			if got != tt.want {
				t.Fatalf("ParseLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigureSharedThreshold(t *testing.T) {
	t.Setenv("CORE_LOG_OUTPUT", "file")
	oldDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(oldDir); err != nil {
			t.Error(err)
		}
	})
	previous := slog.Default()
	t.Cleanup(func() { _ = Close(); SetLevel(LevelDebug); slog.SetDefault(previous) })
	if err := Configure("test-service", "warn"); err != nil {
		t.Fatal(err)
	}
	Debug("hidden-common")
	slog.Info("hidden-slog")
	Warn("visible-common")
	slog.Warn("visible-slog")
	if err := Close(); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile("log/test-service.txt")
	if err != nil {
		t.Fatal(err)
	}
	text := string(data)
	if strings.Contains(text, "hidden-") || !strings.Contains(text, "visible-common") || !strings.Contains(text, "visible-slog") {
		t.Fatalf("unexpected logs: %s", text)
	}
}

func TestConfigureOutputs(t *testing.T) {
	for _, tt := range []struct {
		name, level, mode string
		file, console     bool
	}{
		{"debug-auto", "debug", "auto", true, true},
		{"trace-auto", "trace", "auto", true, true},
		{"info-auto-in-test", "info", "auto", false, true},
		{"debug-file", "debug", "file", true, false},
		{"debug-console", "debug", "console", false, true},
		{"info-both", "info", "both", true, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("CORE_LOG_OUTPUT", tt.mode)
			oldDir, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}
			if err := os.Chdir(t.TempDir()); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if err := os.Chdir(oldDir); err != nil {
					t.Error(err)
				}
			})
			capture, err := os.CreateTemp(t.TempDir(), "stdout")
			if err != nil {
				t.Fatal(err)
			}
			previousStdout, previousSlog := os.Stdout, slog.Default()
			os.Stdout = capture
			t.Cleanup(func() {
				os.Stdout = previousStdout
				_ = Close()
				writer = previousStdout
				SetLevel(LevelDebug)
				slog.SetDefault(previousSlog)
				_ = capture.Close()
			})
			if err := Configure("output-test", tt.level); err != nil {
				t.Fatal(err)
			}
			Info("visible-common")
			slog.Info("visible-slog")
			if err := Close(); err != nil {
				t.Fatal(err)
			}
			console, err := os.ReadFile(capture.Name())
			if err != nil {
				t.Fatal(err)
			}
			file, err := os.ReadFile("log/output-test.txt")
			if tt.file && err != nil {
				t.Fatal(err)
			}
			if !tt.file && !os.IsNotExist(err) {
				t.Fatalf("unexpected file: %v", err)
			}
			for _, marker := range []string{"visible-common", "visible-slog"} {
				if strings.Contains(string(console), marker) != tt.console {
					t.Errorf("console: %q", console)
				}
				if strings.Contains(string(file), marker) != tt.file {
					t.Errorf("file: %q", file)
				}
			}
		})
	}
}
