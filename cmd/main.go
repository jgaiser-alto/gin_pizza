package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"pizza/pkg/common/db"
	"pizza/pkg/pizzas"
)

func main() {
	config := viper.New()
	config.SetConfigName("local")
	config.AddConfigPath("./config")
	if err := config.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	run(config)
}

func run(config *viper.Viper) {
	var host = config.GetString("server.host")
	var port = config.GetString("server.port")
	var dbURL = config.GetString("database.url")

	router := gin.Default()
	dbHandler := db.Init(dbURL)

	pizzas.RegisterRoutes(router, dbHandler)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"port":  port,
			"dbURL": dbURL,
		})
	})

	router.Run(fmt.Sprintf("%s:%s", host, port))
}
