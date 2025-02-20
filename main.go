package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Créer un routeur Gin
	r := gin.Default()

	// route GET /tasks
	r.GET("/tasks", func(c *gin.Context) {
		// Pour l'instant, on renvoie une liste vide
		tasks := []string{}

		c.JSON(http.StatusOK, tasks)
	})

	// serveur port 8080
	r.Run(":8080")
}
