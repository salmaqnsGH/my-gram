package controllers

import (
	"errors"
	"fmt"
	"my-gram/models"
	"my-gram/services"
	util "my-gram/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type photoController struct {
	service services.PhotoService
}

func NewPhotoController(service services.PhotoService) *photoController {
	return &photoController{service}
}

func (c *photoController) CreatePhoto(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	file, err := ctx.FormFile("photo_url")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	unixTime := time.Now().UnixNano()
	path := fmt.Sprintf("images/%d_%s", unixTime, file.Filename)
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	var input models.CreatePhotoInput
	input.PhotoUrl = path
	input.Title = ctx.PostForm("title")
	input.Caption = ctx.PostForm("caption")
	input.UserID = userID

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

func (c *photoController) GetPhotos(ctx *gin.Context) {
	photos, err := c.service.GetPhotos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (c *photoController) GetPhotosByUserID(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	photos, err := c.service.GetPhotosByUserID(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (c *photoController) UpdatePhoto(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	photoIDInt, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid product ID"))
		return
	}
	photoID := uint(photoIDInt)

	existingPhoto, err := c.service.GetPhotoByID(photoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messsage": "Not found",
			"error":    err.Error(),
		})
		return
	}

	file, err := ctx.FormFile("photo_url")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	unixTime := time.Now().UnixNano()
	path := fmt.Sprintf("images/%d_%s", unixTime, file.Filename)

	if file != nil {
		err = ctx.SaveUploadedFile(file, path)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server error",
				"error":   err.Error(),
			})
			return
		}
	}

	title := ctx.PostForm("title")
	caption := ctx.PostForm("caption")

	if title == "" {
		title = existingPhoto.Title
	}
	if caption == "" {
		caption = existingPhoto.Caption
	}

	var photo models.Photo
	photo.PhotoUrl = path
	photo.Title = title
	photo.Caption = caption
	photo.UserID = userID

	updatedPhoto, err := c.service.UpdatePhoto(photoID, photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (c *photoController) DeletePhoto(ctx *gin.Context) {
	photoIDInt, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid product ID"))
		return
	}
	photoID := uint(photoIDInt)

	photo, err := c.service.GetPhotoByID(photoID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"messsage": "Not found",
			"error":    err.Error(),
		})
		return
	}

	err = c.service.DeletePhoto(photoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	_, err = util.DeleteFile(photo.PhotoUrl)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Deleted")
}

func (c *photoController) GetPhotoByID(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	photo, err := c.service.GetPhotoByID(uint(inputID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}
