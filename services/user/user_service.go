package services

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
)

type UserService interface {
	CreateUser(ctx context.Context, user requests.CreateUserRequest) (*models.User, error)
	GetListUser(ctx context.Context, req requests.ListRequest) (*[]models.User, error)
}
