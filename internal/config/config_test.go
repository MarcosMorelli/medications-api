package config

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {
	tests := []struct {
		env             string
		logLevel        string
		port            string
		mongoUser       string
		mongoPassword   string
		mongoHost       string
		mongoPort       string
		mongoDatabase   string
		MongoCollection string
		expected        Config
	}{
		{"development", "info", ":8080", "admin", "password", "localhost", "27017", "medical", "medications",
			Config{DevelopmentEnv, InfoLogLevel, ":8080", "admin", "password", "localhost", "27017", "medical", "medications"}},
		{"production", "debug", ":9090", "admin2", "password2", "mongo", "37017", "medical2", "medications2",
			Config{ProductionEnv, DebugLogLevel, ":9090", "admin2", "password2", "mongo", "37017", "medical2", "medications2"}},
	}

	for _, test := range tests {
		os.Setenv("ENV", test.env)
		os.Setenv("LOG_LEVEL", test.logLevel)
		os.Setenv("PORT", test.port)
		os.Setenv("MONGO_USER", test.mongoUser)
		os.Setenv("MONGO_PASSWORD", test.mongoPassword)
		os.Setenv("MONGO_HOST", test.mongoHost)
		os.Setenv("MONGO_PORT", test.mongoPort)
		os.Setenv("MONGO_DB", test.mongoDatabase)
		os.Setenv("MONGO_COLLECTION", test.MongoCollection)

		config := NewConfig()

		if config.Env != test.expected.Env {
			t.Errorf("expected Env %s, got %s", test.expected.Env, config.Env)
		}
		if config.LogLevel != test.expected.LogLevel {
			t.Errorf("expected LogLevel %s, got %s", test.expected.LogLevel, config.LogLevel)
		}
		if config.Port != test.expected.Port {
			t.Errorf("expected Port %s, got %s", test.expected.Port, config.Port)
		}
		if config.MongoUser != test.expected.MongoUser {
			t.Errorf("expected MongoUser %s, got %s", test.expected.MongoUser, config.MongoUser)
		}
		if config.MongoPassword != test.expected.MongoPassword {
			t.Errorf("expected MongoPassword %s, got %s", test.expected.MongoPassword, config.MongoPassword)
		}
		if config.MongoHost != test.expected.MongoHost {
			t.Errorf("expected MongoHost %s, got %s", test.expected.MongoHost, config.MongoHost)
		}
		if config.MongoPort != test.expected.MongoPort {
			t.Errorf("expected MongoPort %s, got %s", test.expected.MongoPort, config.MongoPort)
		}
		if config.MongoDatabase != test.expected.MongoDatabase {
			t.Errorf("expected MongoDatabase %s, got %s", test.expected.MongoDatabase, config.MongoDatabase)
		}
		if config.MongoCollection != test.expected.MongoCollection {
			t.Errorf("expected MongoCollection %s, got %s", test.expected.MongoCollection, config.MongoCollection)
		}
	}
}

func TestGetEnvVar(t *testing.T) {
	os.Setenv("TEST_ENV_VAR", "test_value")
	defer os.Unsetenv("TEST_ENV_VAR")

	result := GetEnvVar("TEST_ENV_VAR", "default_value")
	expected := "test_value"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	result = GetEnvVar("NON_EXISTENT_ENV_VAR", "default_value")
	expected = "default_value"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("ENV", "development")
	defer os.Unsetenv("ENV")

	result := GetEnv()
	expected := DevelopmentEnv
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	os.Setenv("ENV", "invalid")
	defer os.Unsetenv("ENV")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for invalid environment")
		}
	}()
	GetEnv()
}

func TestIsValidEnv(t *testing.T) {
	tests := []struct {
		env      Env
		expected bool
	}{
		{DevelopmentEnv, true},
		{ProductionEnv, true},
		{"invalid", false},
	}

	for _, test := range tests {
		result := IsValidEnv(test.env)
		if result != test.expected {
			t.Errorf("expected %v, got %v for env %s", test.expected, result, test.env)
		}
	}
}

func TestGetLogLevel(t *testing.T) {
	os.Setenv("LOG_LEVEL", "debug")
	defer os.Unsetenv("LOG_LEVEL")

	result := GetLogLevel()
	expected := DebugLogLevel
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}

	os.Setenv("LOG_LEVEL", "invalid")
	defer os.Unsetenv("LOG_LEVEL")

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic for invalid log level")
		}
	}()
	GetLogLevel()
}

func TestIsValidLogLevel(t *testing.T) {
	tests := []struct {
		level    LogLevel
		expected bool
	}{
		{DebugLogLevel, true},
		{InfoLogLevel, true},
		{ErrorLogLevel, true},
		{"invalid", false},
	}

	for _, test := range tests {
		result := IsValidLogLevel(test.level)
		if result != test.expected {
			t.Errorf("expected %v, got %v for log level %s", test.expected, result, test.level)
		}
	}
}
