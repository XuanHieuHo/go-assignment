package services

import (
	"context"
	"time"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/uow"
)

// CreateUser implements UserService.
func (u *UserServiceImpl) CreateUser(ctx context.Context, req requests.CreateUserRequest) (*models.User, error) {
	user := models.User{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return uow.New(u.db, ctx).UserRepo().Create(ctx, user)
}
