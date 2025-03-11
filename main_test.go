package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
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
				results[taskID] = computeSum(n)
			}(i)
		}

		wg.Wait()
		c.JSON(http.StatusOK, gin.H{
			"message": "Toutes les tâches sont terminées 🚀",
			"results": results,
		})
	})

	return r
}

func TestCreateTask(t *testing.T) {
	router := setupRouter()

	task := Task{Title: "Nouvelle tâche"}
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var createdTask Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)
	assert.Equal(t, "Nouvelle tâche", createdTask.Title)
}

func TestGetTasks(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestUpdateTask(t *testing.T) {
	router := setupRouter()

	task := Task{Title: "Tâche à mettre à jour"}
	jsonValue, _ := json.Marshal(task)
	req, _ := http.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var createdTask Task
	json.Unmarshal(w.Body.Bytes(), &createdTask)

	updatedTask := Task{Title: "Tâche mise à jour"}
	jsonValue, _ = json.Marshal(updatedTask)
	req, _ = http.NewRequest("PUT", "/tasks/"+strconv.Itoa(createdTask.ID), bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var modifiedTask Task
	json.Unmarshal(w.Body.Bytes(), &modifiedTask)
	assert.Equal(t, "Tâche mise à jour", modifiedTask.Title)
}

func TestParallelProcessing(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/tasks/parallel", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Toutes les tâches sont terminées 🚀", response["message"])
	assert.NotNil(t, response["results"])
}
