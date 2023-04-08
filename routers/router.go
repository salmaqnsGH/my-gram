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

	r := gin.Default()

	photoRouter := r.Group("photos")
	{
		photoRouter.POST("/", photoController.CreatePhoto)
	}

	return r
}
