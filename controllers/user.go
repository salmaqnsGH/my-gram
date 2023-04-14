package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"my-gram/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *userController {
	return &userController{service}
}

func (c *userController) RegisterUser(ctx *gin.Context) {
	var input models.User

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	newUser, err := c.service.CreateUser(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	responseUser := models.RegisterUserResponse{}
	responseUser.ID = newUser.ID
	responseUser.Username = newUser.Username
	responseUser.Email = newUser.Email
	responseUser.Age = newUser.Age
	responseUser.CreatedAt = newUser.CreatedAt
	responseUser.UpdatedAt = newUser.UpdatedAt

	ctx.JSON(http.StatusCreated, responseUser)
}

func (c *userController) LoginUser(ctx *gin.Context) {
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	existUser, err := c.service.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, "Invalid email/password", "Bad Request"))
		return
	}

	if !utils.PasswordValid(existUser.Password, user.Password) {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, "Invalid email/password", "Bad Request"))
		return
	}

	token, err := utils.GenerateToken(existUser.ID, existUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.ErrResponse(http.StatusInternalServerError, "Failed to generate token", "Bad Request"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
