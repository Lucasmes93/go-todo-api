package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// Définition de la structure Task
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
 
// Stockage en mémoire des tâches
var (
	tasks  []Task
	mutex  sync.Mutex // Mutex pour éviter les problèmes de concurrence
	nextID = 1        // ID incrémental
)
 
func main() {
	// Créer un routeur Gin
	r := gin.Default()
 
	// Route GET /tasks pour récupérer la liste des tâches
	r.GET("/tasks", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		c.JSON(http.StatusOK, tasks)
	})
 
	// Route POST /tasks pour ajouter une nouvelle tâche
	r.POST("/tasks", func(c *gin.Context) {
		var newTask Task
 
		// Vérifier si le JSON envoyé est valide
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}
 
		// Ajouter un ID unique à la tâche
		mutex.Lock()
		newTask.ID = nextID
		nextID++
		tasks = append(tasks, newTask)
		mutex.Unlock()
 
		// Retourner la tâche ajoutée
		c.JSON(http.StatusCreated, newTask)
	})
 
	// Lancer le serveur sur le port 8080
	r.Run(":8080")
}