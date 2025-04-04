package controller

import (
	"net/http"

	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/responses"
	friendService "github.com/XuanHieuHo/go-assignment/services/friendship"
	"github.com/XuanHieuHo/go-assignment/util"
	"github.com/gin-gonic/gin"
	userController "github.com/XuanHieuHo/go-assignment/controllers/user"
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
		ctx.JSON(http.StatusBadRequest, util.OutlineResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data: util.ErrorResponse{
				Error: err.Error(),
			},
		})
		return
	}

	friendship, err := controller.FriendshipService.CreateFriendship(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.OutlineResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server",
			Data: util.ErrorResponse{
				// Error: fmt.Sprintf("%+v", err),
				Error: err.Error(),
			},
		},
		)
		return
	}
	ctx.JSON(http.StatusOK, util.OutlineResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   friendship,
	})
}


func (controller *FriendshipController) GetFriendOfUser(ctx *gin.Context) {
	var reqList requests.ListRequest
	var reqEmail requests.EmailRequest
	if err := ctx.ShouldBindJSON(&reqEmail); err != nil {
		ctx.JSON(http.StatusBadRequest, util.OutlineResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data: util.ErrorResponse{
				Error: err.Error(),
			},
		})
		return
	}

	if err := ctx.ShouldBindQuery(&reqList); err != nil {
		ctx.JSON(http.StatusBadRequest, util.OutlineResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data: util.ErrorResponse{
				Error: err.Error(),
			},
		},
		)
		return
	}

	friends, err := controller.FriendshipService.GetFriendOfUser(ctx, reqEmail.Email, reqList)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.OutlineResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server",
			Data: util.ErrorResponse{
				// Error: fmt.Sprintf("%+v", err),
				Error: err.Error(),
			},
		},
		)
		return
	}

	var friendsRes []responses.UserResponse
	for _, f := range *friends {
		friendsRes = append(friendsRes, userController.NewUserResponse(f))
	}
	ctx.JSON(http.StatusOK, util.OutlineResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   friendsRes,
	})
}