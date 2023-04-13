package routers

import (
	"my-gram/controllers"
	"my-gram/middlewares"
	"my-gram/repositories"
	"my-gram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "my-gram/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title My-gram
// @version 1.0
// @description Service to post photo
// @termsOfService https://google.com
// @contact.name API Support
// @contact.email nurussalamahqonaah@gmail.com
// @lisence.name Apache 2.0
// @lisence.url https://google.com
// @host localhost:3000
// @BasePath /
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
		photoRouter.Use(middlewares.AuthMiddleware())
		photoRouter.POST("/", photoController.CreatePhoto)
		photoRouter.GET("/", photoController.GetPhotos)
		photoRouter.GET("/user", photoController.GetPhotosByUserID)
		photoRouter.GET("/:photoID", photoController.GetPhotoByID)
		photoRouter.PUT("/:photoID", photoController.UpdatePhoto)
		photoRouter.DELETE("/:photoID", photoController.DeletePhoto)
	}

	commentRouter := r.Group("comments")
	{
		commentRouter.Use(middlewares.AuthMiddleware())
		commentRouter.POST("/", commentController.CreateComment)
		commentRouter.GET("/", commentController.GetComments)
		commentRouter.GET("/:commentID", commentController.GetCommentByID)
		commentRouter.GET("/photo/:photoID", commentController.GetCommentsByPhotoID)
		commentRouter.PUT("/:commentID", commentController.UpdateComment)
		commentRouter.DELETE("/:commentID", commentController.DeleteComment)
	}

	socialMediaRouter := r.Group("social-medias")
	{
		socialMediaRouter.Use(middlewares.AuthMiddleware())
		socialMediaRouter.POST("/", socialMediaController.CreateSocialMedia)
		socialMediaRouter.GET("/", socialMediaController.GetSocialMedias)
		socialMediaRouter.GET("/:id", socialMediaController.GetSocialMediaByID)
		socialMediaRouter.PUT("/:id", socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", socialMediaController.DeleteSocialMedia)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
