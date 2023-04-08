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

	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", userController.RegisterUser)
	}

	photoRouter := r.Group("photos")
	{
		photoRouter.POST("/", photoController.CreatePhoto)
		// photoRouter.GET("/", photoController.GetPhotos)
		photoRouter.GET("/", photoController.GetPhotosByUserID)
	}

	return r
}
