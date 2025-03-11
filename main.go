package main

import (
	"encoding/json" 
	"fmt"          
	"net/http"     
	"os"            
	"strconv"       
	"sync"         
	"time"

	"github.com/gin-gonic/gin" 
	"math/rand"
	
)


type Task struct {
	ID    int    `json:"id"`    
	Title string `json:"title"` 
}

var (
	tasks  []Task     // Liste des tâches stockées en mémoire
	mutex  sync.Mutex // Mutex pour gérer l'accès concurrentiel aux tâches
	nextID = 1        // ID incrémental pour attribuer un identifiant unique aux nouvelles tâches
)

const taskFile = "tasks.json"

func saveTasksToFile() {
	data, err := json.MarshalIndent(tasks, "", "  ") 
	if err != nil {
		fmt.Println("Erreur d'encodage JSON :", err)
		return
	}
	err = os.WriteFile(taskFile, data, 0644)
	if err != nil {
		fmt.Println("Erreur d'écriture dans le fichier :", err)
	}
}

func computeSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func loadTasksFromFile() {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return
	}

	data, err := os.ReadFile(taskFile)
	if err != nil {
		fmt.Println("Erreur de lecture du fichier :", err)
		return
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Println("Erreur de décodage JSON :", err)
		return
	}

	for _, task := range tasks {
		if task.ID >= nextID {
			nextID = task.ID + 1
		}
	}
}

func main() {

	loadTasksFromFile()

	r := gin.Default()

	r.GET("/tasks", func(c *gin.Context) {
		mutex.Lock()
		defer mutex.Unlock()
		c.JSON(http.StatusOK, tasks)
	})

	r.POST("/tasks", func(c *gin.Context) {
		var newTask Task 


		if err := c.ShouldBindJSON(&newTask); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides"}) 
			return
		}

		mutex.Lock()
		newTask.ID = nextID 
		nextID++           
		tasks = append(tasks, newTask)
		saveTasksToFile()
		mutex.Unlock()

		c.JSON(http.StatusCreated, newTask) 
	})

	
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

		for i, task := range tasks {
			if task.ID == taskID { 
				tasks[i].Title = updatedTask.Title 
				saveTasksToFile()                 
				c.JSON(http.StatusOK, tasks[i])   
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche non trouvée"})
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		idParam := c.Param("id")             
		taskID, err := strconv.Atoi(idParam) 
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"}) 
			return
		}

		mutex.Lock()
		defer mutex.Unlock()

		for i, task := range tasks {
			if task.ID == taskID { 
				tasks = append(tasks[:i], tasks[i+1:]...)                  
				saveTasksToFile()                                          
				c.JSON(http.StatusOK, gin.H{"message": "Tâche supprimée"}) 
				return
			}
		}

		c.JSON(http.StatusNotFound, gin.H{"error": "Tâche non trouvée"})
	})

	r.GET("/tasks/process", func(c *gin.Context) {
		go func() {
			fmt.Println("Démarrage du traitement en arrière-plan...")
			time.Sleep(5 * time.Second)
			fmt.Println("Traitement terminé.")
		}()
		c.JSON(http.StatusAccepted, gin.H{"message": "Traitement lancé en arrière-plan"})
	})

	r.GET("/tasks/parallel", func(c *gin.Context) {
		var wg sync.WaitGroup
		taskCount := 5
		results := make([]int, taskCount) 
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < taskCount; i++ {
			wg.Add(1)
			go func(taskID int) {
				defer wg.Done()
			
				sleepTime := time.Duration(2+rand.Intn(4)) * time.Second
				time.Sleep(sleepTime)

				n := 10 + rand.Intn(91)
				sumResult := computeSum(n)
				results[taskID] = sumResult
				fmt.Printf("Tâche %d terminée après %v - Somme calculée: %d\n", taskID+1, sleepTime, sumResult)
			}(i)
		}

		wg.Wait()
		c.JSON(http.StatusOK, gin.H{
			"message": "Toutes les tâches sont terminées 🚀",
			"results": results,
		})
	})

	r.Run(":8080")
}
