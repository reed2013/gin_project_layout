package routers

import (
	"github.com/reed2013/gin_project_layout/api/controllers"
	"github.com/gin-gonic/gin"
	"github.com/reed2013/gin_project_layout/api/internel/mymiddlewares"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(mymiddlewares.Trace())

	api := r.Group("")
	{
		user := api.Group("/user")
		{
			user.GET("/detail/:id", controllers.UserController{}.GetDetail)
		}
	}

	return r
}
