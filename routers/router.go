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
		userRouter.POST("/login", userController.LoginUser)
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
		socialMediaRouter.GET("/", socialMediaController.GetSocialMedias)
		socialMediaRouter.GET("/:id", socialMediaController.GetSocialMediaByID)
		// socialMediaRouter.GET("/all/:photoID", socialMediaController.GetCommentsByPhotoID)
		socialMediaRouter.PUT("/:id", socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", socialMediaController.DeleteSocialMedia)

	}
	return r
}
