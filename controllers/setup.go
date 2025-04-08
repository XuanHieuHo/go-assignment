package controllers

import (
	friendshipController "github.com/XuanHieuHo/go-assignment/controllers/friendship"
	userController "github.com/XuanHieuHo/go-assignment/controllers/user"
	friendshipService "github.com/XuanHieuHo/go-assignment/services/friendship"
	userService "github.com/XuanHieuHo/go-assignment/services/user"
	"gorm.io/gorm"
)

type ControllerRegistery struct {
	UserController   *userController.UserController
	FriendController *friendshipController.FriendshipController
}

func NewGormControllerRegistery(db *gorm.DB) *ControllerRegistery {
	// service
	userService := userService.NewUserServiceImpl(db)
	friendshipService := friendshipService.NewFriendshipServiceImpl(db)

	// controller
	return &ControllerRegistery{
		UserController:   userController.NewUserController(userService),
		FriendController: friendshipController.NewFriendshipController(friendshipService),
	}

}
