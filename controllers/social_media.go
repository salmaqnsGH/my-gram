package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"my-gram/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type socialMediaController struct {
	service services.SocialMediaService
}

func NewSocialMediaController(service services.SocialMediaService) *socialMediaController {
	return &socialMediaController{service}
}

// CreateSocialMedia godoc
// @Summary Create social media
// @Description Create social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Param models.CreateSocialMediaInput body models.CreateSocialMediaInput true "Create Social Media"
// @Success 201 {object} models.SocialMedia
// @Router /social-medias [post]
func (c *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	var input models.CreateSocialMediaInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	newSocialMedia, err := c.service.CreateSocialMedia(input, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	ctx.JSON(http.StatusCreated, newSocialMedia)
}

// GetSocialMediaByID godoc
// @Summary Get one social media
// @Description Get social media by ID
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path uint true "ID of Social Media"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [get]
func (c *socialMediaController) GetSocialMediaByID(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	socialMedia, err := c.service.GetSocialMediaByID(uint(inputID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrResponse(http.StatusNotFound, err.Error(), "Not Found"))
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

// UpdateSocialMedia godoc
// @Summary Update social media
// @Description Update social media by ID
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path uint true "ID of Social Media"
// @Param models.UpdateSocialMediaInput body models.UpdateSocialMediaInput true "Update social media"
// @Success 200 {object} models.SocialMedia
// @Router /social-medias/{id} [put]
func (c *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	existingSocialMedia, err := c.service.GetSocialMediaByID(uint(inputID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrResponse(http.StatusNotFound, err.Error(), "Not Found"))
		return
	}

	var inputData models.UpdateSocialMediaInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	if inputData.Name == "" {
		inputData.Name = existingSocialMedia.Name
	}
	if inputData.SocialMediaUrl == "" {
		inputData.SocialMediaUrl = existingSocialMedia.SocialMediaUrl
	}

	updatedComment, err := c.service.UpdateSocialMedia(uint(inputID), inputData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

// GetSocialMedias godoc
// @Summary Get all social media
// @Description Get all social media
// @Tags Social Media
// @Accept json
// @Produce json
// @Success 200 {object} []models.SocialMedia
// @Router /social-medias [get]
func (c *socialMediaController) GetSocialMedias(ctx *gin.Context) {
	socialMedias, err := c.service.GetSocialMedias()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

// DeleteSocialMedia godoc
// @Summary Delete social media
// @Description Delete one social media by ID
// @Tags Social Media
// @Accept json
// @Produce json
// @Param id path uint true "ID of Social Media"
// @Success 200 "Deleted"
// @Router /social-medias/{id} [delete]
func (c *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
		return
	}

	err = c.service.DeleteSocialMedia(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.ErrResponse(http.StatusNotFound, err.Error(), "Not Found"))
		return
	}

	ctx.JSON(http.StatusOK, "Deleted")
}
