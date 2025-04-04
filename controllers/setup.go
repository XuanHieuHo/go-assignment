package controllers

import (
	friendshipController "github.com/XuanHieuHo/go-assignment/controllers/friendship"
	userController "github.com/XuanHieuHo/go-assignment/controllers/user"
	friendshipService "github.com/XuanHieuHo/go-assignment/services/friendship"
	userService "github.com/XuanHieuHo/go-assignment/services/user"
	"github.com/XuanHieuHo/go-assignment/uow"
	"gorm.io/gorm"
)

type ControllerRegistery struct {
	//DB               *gorm.DB
	UserController   *userController.UserController
	FriendController *friendshipController.FriendshipController
}

func NewGormControllerRegistery(db *gorm.DB) *ControllerRegistery {
	// create new GORM executor (add or replace executor in future)
	uowManager := uow.NewManager(db)
	// service
	userService := userService.NewUserServiceImpl(uowManager)
	friendshipService := friendshipService.NewFriendshipServiceImpl(uowManager)

	// controller
	return &ControllerRegistery{
		//DB:               db,
		UserController:   userController.NewUserController(userService),
		FriendController: friendshipController.NewFriendshipController(friendshipService),
	}

}
