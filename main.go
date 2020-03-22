package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"wirsindcorona/config"
	"wirsindcorona/database"
	"wirsindcorona/route"
)

func main() {
	config.Init()

	database.ConnectMongoDb()
	defer database.Disconnect()

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AddAllowHeaders(
		"Access-Control-Allow-Origin",
		"Access-Control-Allow-Credentials",
		"Authorization",
	)
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:4200", "https://unsgehtscorona.de"}
	config.AddAllowMethods("GET", "POST", "PUT", "DELETE", "PATCH")
	r.Use(cors.New(config))

	route.InitCollections()
	route.InitQuestionRoute(r.Group("/questions"))
	route.InitAnswerRoute(r.Group("/answer"))

	rerr := r.Run("0.0.0.0:8080")
	if rerr != nil {
		log.Panic(rerr)
	}
}
