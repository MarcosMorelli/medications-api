package logger

import (
	"log/slog"
	"testing"

	"github.com/MarcosMorelli/medication-api/internal/config"
)

func TestParseLevel(t *testing.T) {
	tests := []struct {
		logLevel config.LogLevel
		expected slog.Level
	}{
		{config.DebugLogLevel, slog.LevelDebug},
		{config.InfoLogLevel, slog.LevelInfo},
		{config.ErrorLogLevel, slog.LevelError},
	}

	for _, test := range tests {
		level, err := parseLevel(test.logLevel)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if level != test.expected {
			t.Errorf("expected %v, got %v", test.expected, level)
		}
	}

	_, err := parseLevel("invalid")
	if err == nil {
		t.Errorf("expected error for invalid log level")
	}
}
