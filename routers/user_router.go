package routers

import (
	userController "github.com/XuanHieuHo/go-assignment/controllers/user"
	"github.com/XuanHieuHo/go-assignment/handler"
	"github.com/gin-gonic/gin"
)

func UserRouter(router gin.IRouter, userController *userController.UserController) {
	router.POST("/users", handler.ErrorHandler(userController.CreateUser))
	router.GET("/1000users", handler.ErrorHandler(userController.Create1000Users))
	router.GET("/fake_user_csv", handler.ErrorHandler(userController.CreateFakeUser))
	router.POST("/users/upload", handler.ErrorHandler(userController.UploadUserFromCSV))
	router.GET("/users", handler.ErrorHandler(userController.ListUser))
}
