package main

import (
	"net/http" // Package pour gérer les requêtes HTTP
	"strconv"  // Package pour convertir les chaînes de caractères en nombres
	"sync"     // Package pour gérer la concurrence avec un mutex

	"github.com/gin-gonic/gin" // Importation du framework Gin pour gérer les routes HTTP
)

// Définition de la structure Task
type Task struct {
	ID    int    `json:"id"`    // Identifiant unique de la tâche, sérialisé en JSON sous le nom "id"
	Title string `json:"title"` // Titre de la tâche, sérialisé en JSON sous le nom "title"
}

// Stockage en mémoire des tâches
var (
	tasks  []Task     // Liste des tâches stockées en mémoire
	mutex  sync.Mutex // Mutex pour éviter les problèmes d'accès concurrentiels
	nextID = 1        // Variable pour attribuer un identifiant unique aux tâches
)

func main() {
	// Créer un routeur Gin
	r := gin.Default()

	// Route GET /tasks pour récupérer la liste de toutes les tâches
	r.GET("/tasks", func(c *gin.Context) {
		mutex.Lock()                 // Verrouillage du mutex pour éviter les conflits d'accès
		defer mutex.Unlock()         // Déverrouillage automatique à la fin de la fonction
		c.JSON(http.StatusOK, tasks) // Envoi de la liste des tâches en JSON avec un code 200 OK
	})

	// Route POST /tasks pour ajouter une nouvelle tâche
	r.POST("/tasks", func(c *gin.Context) {
		var newTask Task // Création d'une nouvelle tâche

		// Vérification de la validité des données JSON envoyées par le client
		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"}) // Retourne une erreur 400 si le JSON est incorrect
			return
		}

		// Ajouter un identifiant unique à la nouvelle tâche
		mutex.Lock()                   // Verrouillage pour modifier la liste en toute sécurité
		newTask.ID = nextID            // Attribution d'un identifiant unique
		nextID++                       // Incrémentation du prochain ID
		tasks = append(tasks, newTask) // Ajout de la tâche à la liste
		mutex.Unlock()                 // Déverrouillage après modification

		// Retourner la tâche ajoutée avec un code 201 Created
		c.JSON(http.StatusCreated, newTask)
	})

	// Route PUT /tasks/:id pour modifier une tâche existante
	r.PUT("/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")             // Récupération du paramètre ID de l'URL
		taskID, err := strconv.Atoi(idParam) // Conversion de l'ID en entier
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"}) // Retourne une erreur 400 si l'ID est invalide
			return
		}

		var updatedTask Task // Création d'une nouvelle structure pour stocker les données mises à jour
		if err := c.ShouldBindJSON(&updatedTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"}) // Retourne une erreur 400 si le JSON est incorrect
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		// Parcourir la liste des tâches pour trouver celle à modifier
		for i, task := range tasks {
			if task.ID == taskID { // Si l'ID correspond
				tasks[i].Title = updatedTask.Title // Mise à jour du titre
				c.JSON(http.StatusOK, tasks[i])    // Retourne la tâche mise à jour avec un code 200 OK
				return
			}
		}

		// Si la tâche n'est pas trouvée, retourner une erreur 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche non trouvée"})
	})

	// Route DELETE /tasks/:id pour supprimer une tâche existante
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")             // Récupération du paramètre ID de l'URL
		taskID, err := strconv.Atoi(idParam) // Conversion de l'ID en entier
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"}) // Retourne une erreur 400 si l'ID est invalide
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		// Parcourir la liste des tâches pour trouver celle à supprimer
		for i, task := range tasks {
			if task.ID == taskID { // Si l'ID correspond
				tasks = append(tasks[:i], tasks[i+1:]...)                  // Supprime la tâche de la liste
				c.JSON(http.StatusOK, gin.H{"message": "Tâche supprimée"}) // Retourne un message de confirmation
				return
			}
		}

		// Si la tâche n'est pas trouvée, retourner une erreur 404
		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche non trouvée"})
	})

	// Lancer le serveur sur le port 8080
	r.Run(":8080")
}
