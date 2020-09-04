package controllers

import (
	"github.com/reed2013/gin_project_layout/pkgs/common/models"
	"github.com/reed2013/gin_project_layout/pkgs/db"
	"github.com/gin-gonic/gin"
	"log"
)

type UserController struct {

}

func (user UserController) GetDetail(ctx *gin.Context) {
	log.Println(db.Master())
	/*
	db.Master().Create(&models.UserModel{
		Name: "reed",
		Age: 20,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	})
	*/

	var userModel models.UserModel
	db.Slave().First(&userModel, 3)
	log.Println(userModel)
	ctx.JSON(200, userModel)
}
