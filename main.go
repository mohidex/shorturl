package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohidex/shorturl/config"
	"github.com/mohidex/shorturl/database"
)

func main() {
	database.InitDB()
	database.InitRedis()
	fmt.Println(database.GetDB())
	fmt.Println(database.GetRedis())
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":  "pong",
			"dbname":   config.GetEnv().DB.Name,
			"user":     config.GetEnv().DB.User,
			"password": config.GetEnv().DB.Password,
			"port":     config.GetEnv().DB.Port,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
