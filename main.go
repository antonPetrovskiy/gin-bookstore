package main

import (
	"github.com/RazorEdgexD/gin-bookstore/controllers"
	"github.com/RazorEdgexD/gin-bookstore/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	r.GET("/leaderbordAll", controllers.FindBooks)
	r.GET("/leaderbord", controllers.FindBooksTen)
	r.GET("/leaderbord/:id", controllers.FindBook)
	r.GET("/playerPlace/:author", controllers.FindBookByPlayer)
	r.POST("/leaderbord", controllers.CreateBook)
	r.PATCH("/leaderbord/:id", controllers.UpdateBook)
	r.PATCH("/leaderbordUUID/:author", controllers.UpdateBookByPlayer)
	r.DELETE("/leaderbord/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}
