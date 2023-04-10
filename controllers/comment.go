package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *commentController {
	return &commentController{service}
}

func (c *commentController) CreateComment(ctx *gin.Context) {
	var input models.CreateCommentInput
	// TODO: fix input ID
	input.PhotoID = 1
	input.UserID = 1

	// TODO: validate userID and photoID

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	newComment, err := c.service.CreateComment(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newComment)
}
