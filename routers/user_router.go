package routers

import (
	userController "github.com/XuanHieuHo/go-assignment/controllers/user"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup, userController *userController.UserController) {
	router.POST("/users", userController.CreateUser)
	router.GET("/1000users", userController.Create1000Users)
	router.GET("/users", userController.ListUser)
}
