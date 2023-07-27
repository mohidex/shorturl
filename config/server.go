package config

type Env struct {
	Debug        bool
	MaxIdleConns int
	MaxOpenConns int
	ServerPort   string
}

func GetEnv() *Env {
	return &Env{
		Debug:        true,
		ServerPort:   "5000",
		MaxIdleConns: 50,
		MaxOpenConns: 100,
	}
}
