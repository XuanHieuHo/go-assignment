package controllers

import (
	"net/http"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/responses"
	userService "github.com/XuanHieuHo/go-assignment/services/user"
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
		ctx.Error(err)
		return
	}

	user, err := controller.UserService.CreateUser(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	rsp := NewUserResponse(*user)
	responses.NewResponseBuilder().
		WithCode(http.StatusOK).
		WithData(rsp).
		RespondWithJSON(ctx)
}

func (controller *UserController) ListUser(ctx *gin.Context) {
	var req requests.ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.Error(err)
		return
	}

	users, err := controller.UserService.GetListUser(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}
	var rsp []responses.UserResponse
	for _, user := range *users {
		// rsp := newUserResponse(user)
		rsp = append(rsp, NewUserResponse(user))
	}

	responses.NewResponseBuilder().
		WithCode(http.StatusOK).
		WithData(rsp).
		RespondWithJSON(ctx)

}

func (controller *UserController) Create1000Users(ctx *gin.Context) {
	var req requests.CreateUserRequest
	for range 1000 {
		req.Email = faker.Email()
		req.Name = faker.Name()
		_, err := controller.UserService.CreateUser(ctx, req)
		ctx.Error(err)
	}

	responses.NewResponseBuilder().
		WithCode(http.StatusOK).
		WithData("Create 1000 Users Successfully").
		RespondWithJSON(ctx)

}
