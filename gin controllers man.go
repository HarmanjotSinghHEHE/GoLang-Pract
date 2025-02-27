package main

import (
	"GinTut/controllers"
	//"GinTut/services"
	"github.com/gin-gonic/gin"
	//"net/http"
)

func main() {
	router := gin.Default()
	/*r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	r.Run()
	*/

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router)

	router.Run(":6969")
}
