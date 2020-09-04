package routers

import (
	"github.com/reed2013/gin_project_layout/controllers"
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
