package config

type PostgresConfig struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

// LoadConfig loads the configuration from environment variables.
func LoadPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		Name:     getEnv("DB_NAME", "url_db"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "short_url"),
		Password: getEnv("DB_PASSWORD", "password"),
	}
}
