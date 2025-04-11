package controller

import (
	userController "github.com/XuanHieuHo/go-assignment/controllers/user"
	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/responses"
	friendService "github.com/XuanHieuHo/go-assignment/services/friendship"
	"github.com/gin-gonic/gin"
)

type FriendshipController struct {
	FriendshipService friendService.FriendshipService
}

func NewFriendshipController(friendshipService friendService.FriendshipService) *FriendshipController {
	return &FriendshipController{FriendshipService: friendshipService}
}

func (controller *FriendshipController) CreateFriendship(ctx *gin.Context) error {
	var req requests.CreateFriendshipRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return responses.BadRequest("Input are invalid", err)
	}

	friendship, err := controller.FriendshipService.CreateFriendship(ctx, req)
	if err != nil {
		return err
	}

	responses.OK(ctx, friendship)
	return nil
}

func (controller *FriendshipController) GetFriendOfUser(ctx *gin.Context) error {
	var reqList requests.ListRequest
	var reqEmail requests.EmailRequest
	if err := ctx.ShouldBindJSON(&reqEmail); err != nil {
		return responses.BadRequest("Input are invalid", err)
	}

	if err := ctx.ShouldBindQuery(&reqList); err != nil {
		return responses.BadRequest("Input are invalid", err)
	}

	friends, err := controller.FriendshipService.GetFriendOfUser(ctx, reqEmail.Email, reqList)
	if err != nil {
		return err
	}

	var friendsRes []responses.UserResponse
	if friends != nil {
		for _, f := range *friends {
			friendsRes = append(friendsRes, userController.NewUserResponse(f))
		}
	}

	responses.OK(ctx, friendsRes)
	return nil
}
