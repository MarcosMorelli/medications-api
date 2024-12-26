package config

import (
	"log/slog"
	"os"
)

type Env string

const (
	DevelopmentEnv Env = "development"
	ProductionEnv  Env = "production"
)

func (e Env) IsProduction() bool {
	return e == ProductionEnv
}

type LogLevel string

const (
	DebugLogLevel LogLevel = "debug"
	InfoLogLevel  LogLevel = "info"
	ErrorLogLevel LogLevel = "error"
)

type Config struct {
	Env             Env
	LogLevel        LogLevel
	Port            string
	MongoUser       string
	MongoPassword   string
	MongoHost       string
	MongoPort       string
	MongoDatabase   string
	MongoCollection string
}

func (c *Config) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("env", string(c.Env)),
		slog.String("log_level", string(c.LogLevel)),
		slog.String("port", c.Port),
		slog.Group("mongo",
			slog.String("user", c.MongoUser),
			slog.String("password", "REDACTED"),
			slog.String("host", c.MongoHost),
			slog.String("port", c.MongoPort),
			slog.String("database", c.MongoDatabase),
			slog.String("collection", c.MongoCollection),
		),
	)
}

func NewConfig() *Config {
	return &Config{
		Env:             GetEnv(),
		LogLevel:        GetLogLevel(),
		Port:            GetEnvVar("PORT", ":8080"),
		MongoUser:       GetEnvVar("MONGO_USER", "admin"),
		MongoPassword:   GetEnvVar("MONGO_PASSWORD", "password"),
		MongoHost:       GetEnvVar("MONGO_HOST", "localhost"),
		MongoPort:       GetEnvVar("MONGO_PORT", "27017"),
		MongoDatabase:   GetEnvVar("MONGO_DB", "medical"),
		MongoCollection: GetEnvVar("MONGO_COLLECTION", "medications"),
	}
}

func GetEnvVar(env, defaultValue string) string {
	envVarContent := os.Getenv(env)
	if envVarContent == "" {
		return defaultValue
	}
	return envVarContent
}

func GetEnv() Env {
	env := Env(GetEnvVar("ENV", string(DevelopmentEnv)))
	valid := IsValidEnv(env)
	if !valid {
		panic("Invalid environment config: " + env)
	}
	return env
}

func IsValidEnv(env Env) bool {
	switch env {
	case DevelopmentEnv, ProductionEnv:
		return true
	default:
		return false
	}
}

func GetLogLevel() LogLevel {
	level := LogLevel(GetEnvVar("LOG_LEVEL", string(ErrorLogLevel)))
	valid := IsValidLogLevel(level)
	if !valid {
		panic("Invalid log level config: " + level)
	}
	return level
}

func IsValidLogLevel(level LogLevel) bool {
	switch level {
	case DebugLogLevel, InfoLogLevel, ErrorLogLevel:
		return true
	default:
		return false
	}
}
