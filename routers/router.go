package routers

import (
	"my-gram/controllers"
	"my-gram/repositories"
	"my-gram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *gin.Engine {

	photoRepository := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepository)
	photoController := controllers.NewPhotoController(photoService)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	commentRepository := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepository)
	commentController := controllers.NewCommentController(commentService)

	socialMediaRepository := repositories.NewSocialMediaRepository(db)
	socialMediaService := services.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", userController.RegisterUser)
	}

	photoRouter := r.Group("photos")
	{
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.GET("/", photoController.GetPhotos)
		// photoRouter.GET("/", photoController.GetPhotosByUserID)
		photoRouter.PUT("/:photoID", photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoID", photoController.DeletePhoto)
	}

	commentRouter := r.Group("comments")
	{
		commentRouter.POST("/", commentController.CreateComment)
		commentRouter.GET("/", commentController.GetComments)
		commentRouter.GET("/:commentID", commentController.GetCommentByID)
		commentRouter.GET("/all/:photoID", commentController.GetCommentsByPhotoID)
		commentRouter.PUT("/:commentID", commentController.UpdateComment)
		commentRouter.DELETE("/:commentID", commentController.DeleteComment)

	}

	socialMediaRouter := r.Group("social-medias")
	{
		socialMediaRouter.POST("/", socialMediaController.CreateSocialMedia)
		// socialMediaRouter.GET("/", socialMediaController.GetComments)
		// socialMediaRouter.GET("/:commentID", socialMediaController.GetCommentByID)
		// socialMediaRouter.GET("/all/:photoID", socialMediaController.GetCommentsByPhotoID)
		// socialMediaRouter.PUT("/:commentID", socialMediaController.UpdateComment)
		// socialMediaRouter.DELETE("/:commentID", socialMediaController.DeleteComment)

	}
	return r
}
