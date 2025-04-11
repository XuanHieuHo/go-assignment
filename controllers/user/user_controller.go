package controllers

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/XuanHieuHo/go-assignment/importers"
	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/responses"
	userService "github.com/XuanHieuHo/go-assignment/services/user"
	"github.com/XuanHieuHo/go-assignment/utils"
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

func (controller *UserController) CreateUser(ctx *gin.Context) error {
	var req requests.CreateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return responses.BadRequest("Input are invalid", err)
	}

	user, err := controller.UserService.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	rsp := NewUserResponse(*user)
	responses.OK(ctx, rsp)
	return nil
}

func (controller *UserController) ListUser(ctx *gin.Context) error {
	var req requests.ListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		return responses.BadRequest("Input are invalid", err)
	}

	users, err := controller.UserService.GetListUser(ctx, req)
	if err != nil {
		return err
	}
	var rsp []responses.UserResponse
	for _, user := range *users {
		// rsp := newUserResponse(user)
		rsp = append(rsp, NewUserResponse(user))
	}

	responses.OK(ctx, rsp)
	return nil
}

func (controller *UserController) Create1000Users(ctx *gin.Context) error {
	var req requests.CreateUserRequest
	for range 1000 {
		req.Email = faker.Email()
		req.Name = faker.Name()
		_, err := controller.UserService.CreateUser(ctx, req)
		if err != nil {
			return err
		}
	}

	responses.OK(ctx, "Successfully")
	return nil
}

func (controller *UserController) UploadUserFromCSV(ctx *gin.Context) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return err
	}

	dst := fmt.Sprintf("./uploads/%s", faker.TimeString()+file.Filename)
	fileCopy := *file

	go func() {
		backgroundCtx := context.Background()
		req := requests.CreateUserRequest{}
		saver := importers.NewUserCSVSaver(req, controller.UserService)

		if err := utils.UploadFile(&fileCopy, dst); err != nil {
			log.Printf("Error: %+v", err)
			return
		}

		file, err := os.Open(dst)
		if err != nil {
			log.Printf("Error: %+v", err)
			return
		}
		defer file.Close()

		err = importers.ImportCSV(backgroundCtx, file, saver)
		if err != nil {
			log.Printf("Error: %+v", err)
			return
		}

		log.Println("Successfully import data")
	}()

	responses.Accepted(ctx, "File is being processed...")
	return nil
}

func (controller *UserController) CreateFakeUser(ctx *gin.Context) error {
	var data [][]string
	for range 100 {
		data = append(data, []string{faker.Name(), faker.Email()})
	}
	path := fmt.Sprintf("%s.csv", data[0][0])
	err := utils.WriteCSVFile(path, data)
	if err != nil {
		return err
	}
	ctx.File(path)
	return nil
}
