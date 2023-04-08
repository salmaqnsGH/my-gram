package controllers

import (
	"my-gram/models"
	"my-gram/services"
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

	ctx.JSON(http.StatusCreated, newUser)
}
