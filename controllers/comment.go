package controllers

import (
	"my-gram/models"
	"my-gram/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type commentController struct {
	service services.CommentService
}

func NewCommentController(service services.CommentService) *commentController {
	return &commentController{service}
}

// CreateComment godoc
// @Summary Create comment
// @Description Create comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param models.Comment body models.CreateCommentInput true "Create comment"
// @Success 201 {object} models.Comment
// @Router /comments [post]
func (c *commentController) CreateComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["user_id"].(float64))

	var input models.CreateCommentInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	newComment, err := c.service.CreateComment(input, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, newComment)
}

// UpdateComment godoc
// @Summary Update comment
// @Description Update comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param models.Comment body models.UpdateCommentInput true "Update comment"
// @Param id path uint true "ID of comment"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [put]
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

// GetCommentByID godoc
// @Summary Get one comment
// @Description Get comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path uint true "ID of comment"
// @Success 200 {object} models.Comment
// @Router /comments/{id} [get]
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

// GetComments godoc
// @Summary Get all comment
// @Description Get all comment
// @Tags Comment
// @Accept json
// @Produce json
// @Success 200 {object} []models.Comment
// @Router /comments [get]
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

// GetCommentsByPhotoID godoc
// @Summary Get comments by photoID
// @Description Get comments by photoID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path uint true "ID of photo"
// @Success 200 {object} []models.Comment
// @Router /comments/photo/{photoID} [get]
func (c *commentController) GetCommentsByPhotoID(ctx *gin.Context) {
	inputID, err := strconv.Atoi(ctx.Param("photoID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	comments, err := c.service.GetCommentsByPhotoID(uint(inputID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

// DeleteComment godoc
// @Summary Delete comment
// @Description Delete one comment by ID
// @Tags Comment
// @Accept json
// @Produce json
// @Param id path uint true "ID of comment"
// @Success 200 "Deleted"
// @Router /comments/{id} [delete]
func (c *commentController) DeleteComment(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"messsage": "Bad request",
			"error":    err.Error(),
		})
		return
	}

	err = c.service.DeleteComment(uint(commentID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messsage": "Internal server error",
			"error":    err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "Deleted")
}
