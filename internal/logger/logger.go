package logger

import (
	"log"
	"log/slog"
	"os"

	"github.com/MarcosMorelli/medication-api/internal/config"
)

func Init(config *config.Config) {
	levelValue, err := parseLevel(config.LogLevel)
	if err != nil {
		panic(err)
	}

	level := new(slog.LevelVar)
	level.Set(levelValue)

	var logger *slog.Logger
	options := &slog.HandlerOptions{
		Level: level,
	}

	if config.Env.IsProduction() {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, options))
	} else {
		logger = slog.New(slog.NewTextHandler(os.Stdout, options))
		log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	}

	slog.SetDefault(logger)
}

func parseLevel(s config.LogLevel) (slog.Level, error) {
	var level slog.Level
	var err = level.UnmarshalText([]byte(s))
	return level, err
}
