package controllers

import (
	friendshipController "github.com/XuanHieuHo/go-assignment/controllers/friendship"
	userController "github.com/XuanHieuHo/go-assignment/controllers/user"
	friendshipService "github.com/XuanHieuHo/go-assignment/services/friendship"
	userService "github.com/XuanHieuHo/go-assignment/services/user"
	"github.com/XuanHieuHo/go-assignment/uow"
)

type ControllerRegistery struct {
	UserController   *userController.UserController
	FriendController *friendshipController.FriendshipController
}

func NewGormControllerRegistery(uow uow.UnitOfWork) *ControllerRegistery {
	// service
	userService := userService.NewUserServiceImpl(uow)
	friendshipService := friendshipService.NewFriendshipServiceImpl(uow)

	// controller
	return &ControllerRegistery{
		UserController:   userController.NewUserController(userService),
		FriendController: friendshipController.NewFriendshipController(friendshipService),
	}

}
