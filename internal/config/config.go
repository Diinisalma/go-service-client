package config

import (
	"os"
	"time"
)

type HttpClientConfig struct {
	BaseURL string
	Timeout time.Duration
}

type GrpcClientConfig struct {
	Address string
}

type TimeoutConfig struct {
	Addition    time.Duration
	Multiply    time.Duration
	Subtraction time.Duration
}

type Config struct {
	HttpClient HttpClientConfig
	Timeout    TimeoutConfig
	GrpcClient GrpcClientConfig
}

func Load() *Config {
	return &Config{
		HttpClient: HttpClientConfig{
			BaseURL: getEnv("BASE_URL", "http://localhost:1323"),
			Timeout: getEnvDuration("TIMEOUT_CLIENT", 30*time.Second),
		},
		Timeout: TimeoutConfig{
			Addition:    getEnvDuration("TIMEOUT_ADDITION", 60*time.Second),
			Multiply:    getEnvDuration("TIMEOUT_MULTIPLY", 3*time.Second),
			Subtraction: getEnvDuration("TIMEOUT_SUBTRACTION", 3*time.Second),
		},
		GrpcClient: GrpcClientConfig{
			Address: getEnv("GRPC_ADDR", "localhost:50052"),
		},
	}
}

// getEnv membaca environment variable, return fallback kalau kosong.
func getEnv(key string, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// getEnvDuration membaca environment variable sebagai durasi.
// Format value harus valid menurut time.ParseDuration ("60s", "1m30s", "500ms").
// Kalau value kosong atau tidak valid, return fallback.
func getEnvDuration(key string, fallback time.Duration) time.Duration {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	d, err := time.ParseDuration(v)
	if err != nil {
		return fallback
	}
	return d
}
