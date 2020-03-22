package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"wirsindcorona/database"
	"wirsindcorona/route"
)

func main() {
	database.ConnectMongoDb()
	defer database.Disconnect()

	r := gin.Default()
	r.Use(cors.Default())

	route.InitQuestionRoute(r.Group("/questions"))

	rerr := r.Run()
	if rerr != nil {
		log.Panic(rerr)
	}
}
