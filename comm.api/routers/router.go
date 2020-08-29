package routers

import (
	"comm.api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("")
	{
		user := api.Group("/user")
		{
			user.GET("/detail/:id", controllers.UserController{}.GetDetail)
		}
	}

	return r
}
