package routes

import (
	"GoGym/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func HandleRequest(db *gorm.DB) {
	r := gin.Default()

	userController := controllers.UserController{DB: db}
	actionController := controllers.ActionController{DB: db}	

	userRoutes := r.Group("/user")
	{
		userRoutes.GET("/:id", userController.GetUser)       
		userRoutes.GET("", userController.GetAllUsers)       
		userRoutes.POST("", userController.CreateUser)       
		userRoutes.PATCH("/:id", userController.UpdateUser) 
		userRoutes.DELETE("/:id", userController.DeleteUser) 
	}

	actionRoutes := r.Group("/action")
	{		
		actionRoutes.GET("/:id", actionController.GetAction)			
		actionRoutes.GET("", actionController.GetAllActions)
		actionRoutes.POST("", actionController.CreateAction)
		actionRoutes.PATCH("/:id", actionController	.UpdateAction)
		actionRoutes.DELETE("/:id", actionController.DeleteAction)	
	}

	r.Run()
}
