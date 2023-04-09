package main

import (
	"log"

	"github.com/HynDuf/tasks-go-clean-architecture/bootstrap"
	"github.com/HynDuf/tasks-go-clean-architecture/database/mysql"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/delivery/http/route"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/logger"
)

func main() {
	env := bootstrap.NewEnv()
	router := gin.Default()

	db, err := mysql.Connect(env.DBHost, env.DBPort, env.DBUser, env.DBPass, env.DBName, logger.LogLevel(env.DBLogMode))
	if err != nil {
		log.Panicln("Failed to Initialized mysql DB:", err)
	}
	route.Setup(env, db, router)
	router.Run(env.ServerAddress)
}
