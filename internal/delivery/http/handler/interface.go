package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	SignUp(c *gin.Context)
	LogIn(c *gin.Context)
	FetchTasks(c *gin.Context)
	FetchTasksByStatus(c *gin.Context)
	CreateTask(c *gin.Context)
    ToggleTask(c *gin.Context)
    DeleteTask(c *gin.Context)
}
