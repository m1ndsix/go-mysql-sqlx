package main

import (
	"go-mysql-sqlx/controllers"
	"go-mysql-sqlx/db_client"

	"github.com/gin-gonic/gin"
)

func main() {
	db_client.InitialiseDBConnection()

	r := gin.Default()

	r.POST("/", controllers.CreatePost)
	r.GET("/", controllers.GetPosts)
	r.GET("/:id", controllers.GetPost)

	if err := r.Run(":5000"); err != nil {
		panic(err.Error())
	}
}
