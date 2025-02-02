package electronics

type ConfigDatabase struct {
	AppName  string `env:"APP_NAME" env-default:"ELECTRONICS"`
	AppEnv   string `env:"APP_ENV" env-default:"DEV"`
	Port     string `env:"APP_PORT" env-default:"8080"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL" env-default:"ERROR"`
}

var config ConfigDatabase
