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
		commentRouter.PUT("/:commentID", commentController.UpdateComment)
	}

	return r
}
