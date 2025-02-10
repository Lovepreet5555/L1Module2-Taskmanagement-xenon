package controllers

import (
	"net/http"
	"strconv"
	"taskmanagementnew/Database"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Due_Date    string `json:"due_date" binding:"required"`
	Status      string `json:"status" binding:"oneof=pending in_progress completed"`
}

func CreateTask(c *gin.Context) {
	var newTask Task

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := database.ExecDB("INSERT INTO tasks (title, description, due_date, status) VALUES (?, ?, ?, ?)", newTask.Title, newTask.Description, newTask.Due_Date, newTask.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	taskID, _ := result.LastInsertId()
	newTask.ID = int(taskID)

	c.JSON(http.StatusCreated, newTask)
}

func GetTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var task Task
	row := database.DB.QueryRow("SELECT id, title, description, due_date, status FROM tasks WHERE id = ?", taskID)

	if err = row.Scan(&task.ID, &task.Title, &task.Description, &task.Due_Date, &task.Status); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, task)
}

func UpdateTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask Task
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = database.ExecDB("UPDATE tasks SET title = ?, description = ?, due_date = ?, status = ? WHERE id = ?",
		updatedTask.Title, updatedTask.Description, updatedTask.Due_Date, updatedTask.Status, taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	updatedTask.ID = taskID
	c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	_, err = database.ExecDB("DELETE FROM tasks WHERE id = ?", taskID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func ListTask(c *gin.Context) {
	rows, err := database.QueryDB("SELECT id, title, description, due_date, status FROM tasks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Due_Date, &task.Status); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, tasks)
}