package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"net/http"
	"strconv"

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

func (c *socialMediaController) GetSocialMediaByID(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	socialMedia, err := c.service.GetSocialMediaByID(uint(inputID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func (c *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	var inputData models.UpdateSocialMediaInput
	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	updatedComment, err := c.service.UpdateSocialMedia(uint(inputID), inputData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

func (c *socialMediaController) GetSocialMedias(ctx *gin.Context) {
	socialMedias, err := c.service.GetSocialMedias()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

func (c *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	err = c.service.DeleteSocialMedia(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Deleted")
}
