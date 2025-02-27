package main

import (
	"net/http"
	"strconv"
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

	// Route PUT /tasks/:id pour modifier une tâche existante
	r.PUT("/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		taskID, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}

		var updatedTask Task
		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"})
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		// Chercher et modifier la tâche
		for i, task := range tasks {
			if task.ID == taskID {
				tasks[i].Title = updatedTask.Title
				c.JSON(http.StatusOK, tasks[i])
				return
			}
		}

		// Si la tâche n'est pas trouvée
		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche non trouvée"})
	})

	// Route DELETE /tasks/:id pour supprimer une tâche existante
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		taskID, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		// Chercher la tâche et la supprimer
		for i, task := range tasks {
			if task.ID == taskID {
				tasks = append(tasks[:i], tasks[i+1:]...) // Supprime la tâche
				c.JSON(http.StatusOK, gin.H{"message": "Tâche supprimée"})
				return
			}
		}

		// Si la tâche n'est pas trouvée
		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche non trouvée"})
	})

	// Lancer le serveur sur le port 8080
	r.Run(":8080")
}
