package config

type Config struct {
	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDatabase string `env:"POSTGRES_DB"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`

	BotToken string `env:"BOT_TOKEN"`

	LogLevel string `env:"LOG_LEVEL" envDefault:"DEBUG"`
	LogType  string `env:"LOG_TYPE" envDefault:"TEXT"`

	FleetHost     string `env:"FLEET_HOST"`
	FleetParkID   string `env:"FLEET_PARK_ID"`
	FleetClientID string `env:"FLEET_CLIENT_ID"`
	FleetAPIKey   string `env:"FLEET_API_KEY"`

	AdminID string `env:"ADMIN_ID"`

	JumpTaxiHost      string `env:"JUMP_TAXI_HOST"`
	JumpTaxiClientKey string `env:"JUMP_TAXI_CLIENT_KEY"`

	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisPort     string `env:"REDIS_PORT" envDefault:"6379"`
	RedisHost     string `env:"REDIS_HOST" envDefault:"localhost"`

	AppAddr string `env:"APP_ADDR" envDefault:"localhost:8080"`
}
