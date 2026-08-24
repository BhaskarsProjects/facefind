package config

import (
	"os"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	App         AppConfig         `koanf:"primary" validate:"required"`
	Server      ServerConfig      `koanf:"server" validate:"required"`
	Database    DatabaseConfig    `koanf:"database" validate:"required"`
	Auth        AuthConfig        `koanf:"auth" validate:"required"`
	Redis       RedisConfig       `koanf:"redis" validate:"required"`
	Integration IntegrationConfig `koanf:"integration" validate:"required"`
	R2          R2Config          `koanf:"r2" validate:"required"`
	AIWorker    AIWorkerConfig    `koanf:"ai_worker" validate:"required"`
	// Observability *ObservabilityConfig `koanf:"observability"`
}

type AppConfig struct {
	Env string `koanf:"env" validate:"required"`
}

type AIWorkerConfig struct {
	BaseURL string `koanf:"base_url" validate:"required"`
	Token   string `koanf:"token" validate:"required"`
}

type ServerConfig struct {
	Port               string   `koanf:"port" validate:"required"`
	ReadTimeout        int      `koanf:"read_timeout" validate:"required"`
	WriteTimeout       int      `koanf:"write_timeout" validate:"required"`
	IdleTimeout        int      `koanf:"idle_timeout" validate:"required"`
	CORSAllowedOrigins []string `koanf:"cors_allowed_origins" validate:"required"`
}

type DatabaseConfig struct {
	Host     string `koanf:"host" validate:"required"`
	Port     int    `koanf:"port" validate:"required"`
	User     string `koanf:"user" validate:"required"`
	Password string `koanf:"password"`
	Name     string `koanf:"name" validate:"required"`
	SSLMode  string `koanf:"ssl_mode" validate:"required"`

	MinConns        int32         `koanf:"min_conns" validate:"required"`
	MaxConns        int32         `koanf:"max_conns" validate:"required"`
	ConnMaxLifetime time.Duration `koanf:"conn_max_lifetime" validate:"required"`
	ConnMaxIdleTime time.Duration `koanf:"conn_max_idle_time" validate:"required"`
}

type RedisConfig struct {
	Address string `koanf:"address" validate:"required"`
}

type IntegrationConfig struct {
	ResendAPIKey string `koanf:"resend_api_key" validate:"required"`
}

// Authenticator wraps the go-auth0 authentication client.
type AuthConfig struct {
	Audience           string `koanf:"audience" validate:"required"`
	Domain             string `koanf:"domain" validate:"required"`
	ClientID           string `koanf:"client_id" validate:"required"`
	ClientSecret       string `koanf:"client_secret" validate:"required"`
	CallbackURL        string `koanf:"callback_url" validate:"required"`
	SessionSecret      string `koanf:"session_secret" validate:"required"`
	DatabaseConnection string `koanf:"database_connection" validate:"required"`
}

type R2Config struct {
	AccountID       string `koanf:"account_id" validate:"required"`
	AccessKeyID     string `koanf:"access_key_id" validate:"required"`
	SecretAccessKey string `koanf:"secret_access_key" validate:"required"`
	Bucket          string `koanf:"bucket" validate:"required"`
}

func LoadConfig() (*Config, error) {

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()

	k := koanf.New(".")

	err := k.Load(env.Provider("FACEFIND_", ".", func(s string) string {
		return strings.ToLower(strings.TrimPrefix(s, "FACEFIND_"))
	}), nil)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not load initial env variables")
	}

	mainConfig := &Config{}

	err = k.Unmarshal("", mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("could not unmarshal main config")
	}

	// Convert comma-separated CORS origins into a slice.
	if len(mainConfig.Server.CORSAllowedOrigins) > 0 {
		var origins []string

		for _, value := range mainConfig.Server.CORSAllowedOrigins {
			origins = append(
				origins,
				strings.Split(value, ",")...,
			)
		}

		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
		}

		mainConfig.Server.CORSAllowedOrigins = origins
	}

	validate := validator.New()

	err = validate.Struct(mainConfig)
	if err != nil {
		logger.Fatal().Err(err).Msg("config validation failed")
	}

	// Set default observability config if not provided
	// if mainConfig.Observability == nil {
	// 	mainConfig.Observability = DefaultObservabilityConfig()
	// }

	// // Override service name and environment from primary config
	// mainConfig.Observability.ServiceName = "facefind"
	// mainConfig.Observability.Environment = mainConfig.App.Env

	// // Validate observability config
	// if err := mainConfig.Observability.Validate(); err != nil {
	// 	logger.Fatal().Err(err).Msg("invalid observability config")
	// }

	return mainConfig, nil
}
