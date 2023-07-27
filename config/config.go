package config

type DatabaseConfig struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

type Env struct {
	Debug        bool
	MaxIdleConns int
	MaxOpenConns int
	ServerPort   string

	DB DatabaseConfig
}

func GetEnv() *Env {
	return &env
}
