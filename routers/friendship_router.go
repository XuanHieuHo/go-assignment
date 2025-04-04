package routers

import (
	friendshipController "github.com/XuanHieuHo/go-assignment/controllers/friendship"
	"github.com/gin-gonic/gin"
)

func FriendshipRouter(router *gin.RouterGroup, friendshipController *friendshipController.FriendshipController) {
	router.POST("/friendship", friendshipController.CreateFriendship)
	router.POST("/friends", friendshipController.GetFriendOfUser)
}
