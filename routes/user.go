package routes

import "github.com/gin-gonic/gin"

func HandleRequest() {
	r := gin.Default()

	userRoutes := r.Group("/user")
	{
		userRoutes.GET("", )	
		userRoutes.GET("", )	
		userRoutes.POST("", )	
		userRoutes.PATCH("", )	
		userRoutes.DELETE("", )	
	}

	r.Run()
}