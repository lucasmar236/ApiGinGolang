package routes

import (
	"ApiRest/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users/:user", controllers.GetUser)
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.PostUser)
	r.PUT("/users/:user", controllers.PutUser)
	return r
}
