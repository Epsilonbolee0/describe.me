package config

type (
	Config struct {
		App      `yaml:"app"`
		HTTP     `yaml:"http"`
		Logger   `yaml:"logger"`
		Postgres `yaml:"postgres"`
		Shelter  `yaml:"shelter"`
	}

	// Shelter -.
	Shelter struct {
		Secret string `env-required:"true" yaml:"name"    env:"secret"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"app_name"`
		Version string `env-required:"true" yaml:"version" env:"app_version"`
	}

	// HTTP -.
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"http_port"`
	}

	// Logger -.
	Logger struct {
		Level string `env-required:"true" yaml:"log_level" env:"logging_level"`
	}

	// Postgres -.
	Postgres struct {
		MaxIdleConnections int `env-required:"true" yaml:"idle_conns" env:"pg_idle_conns"`
		MaxOpenConnections int `env-required:"true" yaml:"open_conns" env:"pg_open_conns"`

		DSN `env-required:"true" yaml:"dsn"`
	}

	// DSN -.
	DSN struct {
		Unauthorized string `env-required:"true" yaml:"unauthorized" env:"dsn_unauthorized"`
		Student      string `env-required:"true" yaml:"student" env:"dsn_student"`
		Teacher      string `env-required:"true" yaml:"teacher" env:"dsn_teacher"`
	}
)
