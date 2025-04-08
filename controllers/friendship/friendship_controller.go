package controller

import (
	"net/http"

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

func (controller *FriendshipController) CreateFriendship(ctx *gin.Context) {
	var req requests.CreateFriendshipRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.Error(err)
		return
	}

	friendship, err := controller.FriendshipService.CreateFriendship(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	responses.NewResponseBuilder().
		WithCode(http.StatusOK).
		WithData(friendship).
		RespondWithJSON(ctx)
}

func (controller *FriendshipController) GetFriendOfUser(ctx *gin.Context) {
	var reqList requests.ListRequest
	var reqEmail requests.EmailRequest
	if err := ctx.ShouldBindJSON(&reqEmail); err != nil {
		ctx.Error(err)
		return
	}

	if err := ctx.ShouldBindQuery(&reqList); err != nil {
		ctx.Error(err)
		return
	}

	friends, err := controller.FriendshipService.GetFriendOfUser(ctx, reqEmail.Email, reqList)
	if err != nil {
		ctx.Error(err)
		return
	}

	var friendsRes []responses.UserResponse
	for _, f := range *friends {
		friendsRes = append(friendsRes, userController.NewUserResponse(f))
	}

	responses.NewResponseBuilder().
		WithCode(http.StatusOK).
		WithData(friendsRes).
		RespondWithJSON(ctx)
}
