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
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	newUser, err := c.service.CreateUser(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	password := user.Password

	user, err = c.service.GetUserByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messsage": "Failed to login",
			"error":    err.Error(),
		})
		return
	}

	if !utils.PasswordValid(user.Password, password) {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"messsage": "Bad request",
				"error":    err.Error(),
			})
			return
		}
	}
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
