package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	service services.PhotoService
}

func NewPhotoController(service services.PhotoService) *photoController {
	return &photoController{service}
}

func (c *photoController) CreatePhoto(ctx *gin.Context) {
	var input models.Photo

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	newPhoto, err := c.service.CreatePhoto(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newPhoto)
}
