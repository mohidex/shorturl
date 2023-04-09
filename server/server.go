package server

import (
	"fmt"

	"github.com/mohidex/shorturl/config"
)

func Init() {
	serverPort := config.GetEnv().ServerPort
	r := NewRouter()

	serverAddr := fmt.Sprintf(":%s", serverPort)
	r.Run(serverAddr)
}
