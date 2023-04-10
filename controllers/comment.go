package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"net/http"
	"strconv"

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

func (c *commentController) UpdateComment(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	var inputData models.UpdateCommentInput
	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	updatedComment, err := c.service.UpdateComment(uint(inputID), inputData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

func (c *commentController) GetCommentByID(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	comment, err := c.service.GetCommentByID(uint(inputID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (c *commentController) GetComments(ctx *gin.Context) {
	comments, err := c.service.GetComments()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}
