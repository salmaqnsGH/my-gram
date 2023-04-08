package controllers

import (
	"fmt"
	"my-gram/models"
	"my-gram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type photoController struct {
	service services.PhotoService
}

func NewPhotoController(service services.PhotoService) *photoController {
	return &photoController{service}
}

func (c *photoController) CreatePhoto(ctx *gin.Context) {
	file, err := ctx.FormFile("photo_url")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	path := fmt.Sprintf("images/%s", file.Filename)
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	var input models.Photo

	input.PhotoUrl = path
	title := ctx.PostForm("title")
	caption := ctx.PostForm("caption")
	userID := ctx.PostForm("user_id")

	input.Title = title
	input.Caption = caption
	userIDUint, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		fmt.Println("Failed to convert string to uint:", err)
		return
	}
	input.UserID = uint(userIDUint)

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

func (h *photoController) GetPhotos(ctx *gin.Context) {
	photos, err := h.service.GetPhotos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}
