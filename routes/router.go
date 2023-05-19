package routes

import (
	"ApiRest/controllers"
	"ApiRest/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	users := r.Group("/users")
	users.Use(middlewares.Auth)
	users.GET("/:user", controllers.GetUser)
	users.GET("/", controllers.GetUsers)
	users.POST("/", controllers.PostUser)
	users.PUT("/:user", controllers.PutUser)
	r.POST("/login", controllers.Login)
	r.GET("/auth", controllers.Auth)
	return r
}
