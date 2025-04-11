package routers

import (
	friendshipController "github.com/XuanHieuHo/go-assignment/controllers/friendship"
	"github.com/XuanHieuHo/go-assignment/handler"
	"github.com/gin-gonic/gin"
)

func FriendshipRouter(router gin.IRouter, friendshipController *friendshipController.FriendshipController) {
	router.POST("/friendship", handler.ErrorHandler(friendshipController.CreateFriendship))
	router.POST("/friends", handler.ErrorHandler(friendshipController.GetFriendOfUser))
}
