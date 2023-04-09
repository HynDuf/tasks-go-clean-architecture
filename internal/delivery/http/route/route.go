package route

import (
	"github.com/HynDuf/tasks-go-clean-architecture/bootstrap"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/delivery/http/handler"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/delivery/http/middleware"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/repository/taskrepo"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/repository/userrepo"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/usecase/login"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/usecase/refreshtkn"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/usecase/signup"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/usecase/tasks"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, db *gorm.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	userRP := userrepo.NewUserRepository(db)
	taskRP := taskrepo.NewTaskRepository(db)
	loginUsecase := login.NewLoginUsecase(userRP)
	signUpUsecase := signup.NewSignupUsecase(userRP)
	tasksUsecase := tasks.NewTaskUsecase(taskRP)
	refreshTokenUsecase := refreshtkn.NewRefreshTokenUsecase(userRP)
	h := handler.NewHandler(loginUsecase, signUpUsecase, tasksUsecase, refreshTokenUsecase, env)
	publicRouter.POST("/api/signup", h.SignUp)
	publicRouter.POST("/api/login", h.LogIn)
	publicRouter.POST("/api/refresh", h.RefreshToken)

	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	protectedRouter.GET("/api/task", h.FetchTasks)
	protectedRouter.POST("/api/task/add", h.CreateTask)
}
