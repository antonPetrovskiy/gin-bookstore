package main

import (
	"github.com/RazorEdgexD/gin-bookstore/controllers"
	"github.com/RazorEdgexD/gin-bookstore/models"

	"github.com/gin-gonic/gin"

	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
)

func main() {
	r := gin.Default()

	store := persistence.NewInMemoryStore(time.Second)

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

	// Cached leaderbord
	r.GET("/leaderbordAllCache", cache.CachePage(store, time.Minute, controllers.FindBooks))

	// Run the server
	r.Run()
}
