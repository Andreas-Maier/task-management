package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func handleGetTasks(c *gin.Context) {
	var loadedTasks, err = GetAllTasks()
	if err != nil {
		c.JSON(404, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"tasks": loadedTasks})
}

func handleGetTask(c *gin.Context) {
	var task Task
	if err := c.BindUri(&task); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	var loadedTask, err = GetTaskByID(task.ID)
	if err != nil {
		c.JSON(404, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"ID": loadedTask.ID, "Body": loadedTask.Body})
}

func handleCreateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"msg": err})
		return
	}
	id, err := Create(&task)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func handleUpdateTask(c *gin.Context) {
	var task Task
	if err := c.ShouldBindJSON(&task); err != nil {
		log.Print(err)
		c.JSON(400, gin.H{"msg": err})
		return
	}
	savedTask, err := Update(&task)
	if err != nil {
		c.JSON(500, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"task": savedTask})
}

func main() {
	r := gin.Default()
	r.GET("/tasks/:id", handleGetTask)
	r.GET("/tasks/", handleGetTasks)
	r.PUT("/tasks/", handleCreateTask)
	r.POST("/tasks/", handleUpdateTask)
	r.Run() // listen and serve on 0.0.0.0:8080
}
