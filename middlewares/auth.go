package middlewares

import (
	"my-gram/services"
	"my-gram/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := utils.VerifyToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("userData", claims)
		c.Next()
	}
}

func CommentAuthorization(commentService services.CommentService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["user_id"].(float64))

		inputID, err := strconv.Atoi(ctx.Param("commentID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
			return
		}

		comment, err := commentService.GetCommentByID(uint(inputID))

		if comment.UserID != userID {
			if ctx.Request.Method == http.MethodDelete {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(http.StatusUnauthorized, "cannot access this data", "Unauthorized"))
				return
			} else if ctx.Request.Method == http.MethodPut {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(http.StatusUnauthorized, "cannot access this data", "Unauthorized"))
				return
			}

			ctx.Next()
		}

		ctx.Next()
	}
}

func PhotoAuthorization(photoService services.PhotoService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["user_id"].(float64))

		inputID, err := strconv.Atoi(ctx.Param("photoID"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
			return
		}

		photo, err := photoService.GetPhotoByID(uint(inputID))

		if photo.UserID != userID {
			if ctx.Request.Method == http.MethodDelete {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(http.StatusUnauthorized, "cannot access this data", "Unauthorized"))
				return
			} else if ctx.Request.Method == http.MethodPut {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(http.StatusUnauthorized, "cannot access this data", "Unauthorized"))
				return
			}

			ctx.Next()
		}

		ctx.Next()
	}
}

func SocialMediaAuthorization(socialMediaService services.SocialMediaService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["user_id"].(float64))

		inputID, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, utils.ErrResponse(http.StatusBadRequest, err.Error(), "Bad Request"))
			return
		}

		socialMedia, err := socialMediaService.GetSocialMediaByID(uint(inputID))

		if socialMedia.UserID != userID {
			if ctx.Request.Method == http.MethodDelete {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(http.StatusUnauthorized, "cannot access this data", "Unauthorized"))
				return
			} else if ctx.Request.Method == http.MethodPut {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrResponse(http.StatusUnauthorized, "cannot access this data", "Unauthorized"))
				return
			}

			ctx.Next()
		}

		ctx.Next()
	}
}
