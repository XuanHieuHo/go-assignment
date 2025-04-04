package controllers

import (
	"net/http"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/responses"
	userService "github.com/XuanHieuHo/go-assignment/services/user"
	"github.com/XuanHieuHo/go-assignment/util"
	"github.com/gin-gonic/gin"
	"github.com/go-faker/faker/v4"
)

type UserController struct {
	UserService userService.UserService
}

func NewUserController(userService userService.UserService) *UserController {
	return &UserController{UserService: userService}
}

func NewUserResponse(user models.User) responses.UserResponse {
	return responses.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (controller *UserController) CreateUser(ctx *gin.Context) {
	var req requests.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
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

	user, err := controller.UserService.CreateUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.OutlineResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server",
			Data: util.ErrorResponse{
				Error: err.Error(),
			},
		},
		)
		return
	}
	rsp := NewUserResponse(*user)
	ctx.JSON(http.StatusOK, util.OutlineResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rsp,
	})
}

func (controller *UserController) ListUser(ctx *gin.Context) {
	var req requests.ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
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

	users, err := controller.UserService.GetListUser(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.OutlineResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server",
			Data: util.ErrorResponse{
				Error: err.Error(),
			},
		},
		)
		return
	}
	var rsp []responses.UserResponse
	for _, user := range *users {
		// rsp := newUserResponse(user)
		rsp = append(rsp, NewUserResponse(user))
	}
	ctx.JSON(http.StatusOK, util.OutlineResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rsp,
	})

}

func (controller *UserController) Create1000Users(ctx *gin.Context) {
	var req requests.CreateUserRequest
	for range 1000 {
		req.Email = faker.Email()
		req.Name = faker.Name()
		_, err := controller.UserService.CreateUser(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, util.OutlineResponse{
				Code:   http.StatusInternalServerError,
				Status: "Internal Server",
				Data: util.ErrorResponse{
					Error: err.Error(),
				},
			},
			)
			return
		}
	}
	ctx.JSON(http.StatusOK, util.OutlineResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Create 1000 Users Successfully",
	})

}
