package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	service services.SocialMediaService
}

func NewSocialMediaController(service services.SocialMediaService) *socialMediaController {
	return &socialMediaController{service}
}

func (c *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	var input models.CreateSocialMediaInput
	// TODO: fix input ID
	input.UserID = 1

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	newSocialMedia, err := c.service.CreateSocialMedia(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newSocialMedia)
}
