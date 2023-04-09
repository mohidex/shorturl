package config

type DatabaseConfig struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

type RedisConfig struct {
	Addr      string
	Password  string
	RedisDb   int
	SessionDb int
	CacheDb   int
}

type Env struct {
	Debug        bool
	MaxIdleConns int
	MaxOpenConns int
	ServerPort   string

	DB    DatabaseConfig
	Redis RedisConfig
}

func GetEnv() *Env {
	return &env
}
