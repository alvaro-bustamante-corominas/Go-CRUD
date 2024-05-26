package handlers

import (
	"database/sql"
	"go-crud-tareas/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StatusUpdate struct {
	Status string `json:"status"`
}

func CreateTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := models.CreateTask(db, task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)
	}
}

func DeleteTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		err = models.DeleteTask(db, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	}
}

func GetTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tareas, err := models.GetTaskList(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, tareas)
	}

}

func UpdateTaskHandler(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}
		var task models.Task
		if err := c.ShouldBindJSON(&task); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = models.UpdateTask(db, task, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, task)

	}

}

func UpdateStatusTaskHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
			return
		}

		var statusUpdate StatusUpdate
		if err := c.ShouldBindJSON(&statusUpdate); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		status, err := models.ParseStatus(statusUpdate.Status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err = models.UpdateTaskStatus(db, status, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Status updated"})
	}

}
