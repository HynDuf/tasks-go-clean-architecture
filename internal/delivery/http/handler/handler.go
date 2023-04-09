package handler

import (
	"net/http"
	"strconv"

	"github.com/HynDuf/tasks-go-clean-architecture/bootstrap"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/entity"
	"github.com/HynDuf/tasks-go-clean-architecture/internal/domain/interface/usecase"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SuccessResponse struct {
	Data interface{} `json:"data"`
}
type ErrorResponse struct {
	Message string `json:"error"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupRequest struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type restHandler struct {
	loginUsecase        usecase.LoginUsecase
	signUpUsecase       usecase.SignupUsecase
	taskUsecase         usecase.TaskUsecase
	refreshTokenUsecase usecase.RefreshTokenUsecase
	env                 *bootstrap.Env
}

func NewHandler(loginUsecase usecase.LoginUsecase, signUpUsecase usecase.SignupUsecase, taskUsecase usecase.TaskUsecase, refreshTokenUsecase usecase.RefreshTokenUsecase) RestHandler {
	return &restHandler{
		loginUsecase:        loginUsecase,
		signUpUsecase:       signUpUsecase,
		taskUsecase:         taskUsecase,
		refreshTokenUsecase: refreshTokenUsecase,
	}
}

func (h *restHandler) SignUp(c *gin.Context) {
	var request SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	_, err = h.signUpUsecase.GetUserByEmail(request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = h.signUpUsecase.Create(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := h.signUpUsecase.CreateAccessToken(&user, h.env.AccessTokenSecret, h.env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.signUpUsecase.CreateRefreshToken(&user, h.env.RefreshTokenSecret, h.env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}

func (h *restHandler) LogIn(c *gin.Context) {
	var request LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	user, err := h.loginUsecase.GetUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := h.loginUsecase.CreateAccessToken(&user, h.env.AccessTokenSecret, h.env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.loginUsecase.CreateRefreshToken(&user, h.env.RefreshTokenSecret, h.env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}

func (h *restHandler) RefreshToken(c *gin.Context) {
	var request RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: err.Error()})
		return
	}

	id, err := h.refreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, h.env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not found"})
		return
	}

	user, err := h.refreshTokenUsecase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := h.refreshTokenUsecase.CreateAccessToken(&user, h.env.AccessTokenSecret, h.env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := h.refreshTokenUsecase.CreateRefreshToken(&user, h.env.RefreshTokenSecret, h.env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}

func (h *restHandler) FetchTasks(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("x-user-id"))
	tasks, err := h.taskUsecase.FetchByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *restHandler) FetchTasksByStatus(c *gin.Context) {
	userID, _ := strconv.Atoi(c.GetString("x-user-id"))
	completedStatus, _ := strconv.ParseBool(c.Param("completed"))
	tasks, err := h.taskUsecase.FetchByUserIDAndStatus(userID, completedStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
func (h *restHandler) CreateTask(c *gin.Context) {
	var task entity.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	task.UserID, _ = strconv.Atoi(c.GetString("x-user-id"))
	err = h.taskUsecase.Create(&task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: err.Error()})
	}
	c.JSON(http.StatusOK, SuccessResponse{Data: "Task created successfully"})
}

func (h *restHandler) ToggleTask(c *gin.Context) {

}

func (h *restHandler) DeleteTask(c *gin.Context) {

}
