package entity

type (
	// Config -.
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		LOG  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// HTTP -.
	HTTP struct {
		Port            string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		CookieStoreKey  string `env-required:"true" yaml:"cookie_store_key" env:"COOKIE_STORE_KEY"`
		SessionStoreKey string `env-required:"true" yaml:"session_store_key" env:"SESSION_STORE_KEY"`
	}

	// PG -.
	PG struct {
		PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		URL     string `env-required:"true"                 env:"PG_URL"`
	}

	// Log -.
	LOG struct {
		// Format specifies the output log format.
		// Accepted values are: json, logfmt
		Format string `env-required:"true" yaml:"log_format" env:"LOG_FORMAT"`

		// Level is the minimum log level that should appear on the output.
		Level string `env-required:"true" yaml:"log_level" env:"LOG_LEVEL"`

		// NoColor makes sure that no log output gets colorized.
		NoColor bool `env-required:"true" yaml:"log_no_color" env:"LOG_NO_COLOR"`
	}
)
