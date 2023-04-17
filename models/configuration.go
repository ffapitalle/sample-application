package models

import "github.com/pedidosya/peya-go/newrelic"

// Configuration for service
type Configuration struct {
	Vault    ExternalServiceConfig `json:"vault"`
	NewRelic *newrelic.Config      `json:"new_relic"`
	App      AppConfig             `json:"server"`
	Auth     ExternalServiceConfig `json:"auth"`
}

// ExternalServiceConfig representation
type ExternalServiceConfig struct {
	URL     string `json:"url"`
	Key     string `json:"key"`
	Timeout int64  `json:"timeout"`
	Enabled bool   `json:"enabled"`
}

type AppConfig struct {
	AppVersion string `json:"version"`
	AppName    string `json:"name"`
	AppPort    int    `json:"port"`
	AppEnv     string
}
