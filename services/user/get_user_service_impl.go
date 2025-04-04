package services

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
)

// GetListUser implements UserService.
func (u *UserServiceImpl) GetListUser(ctx context.Context, req requests.ListRequest) (*[]models.User, error) {
	return u.Manager.New(ctx).UserRepo().ListUser(ctx, req)
}
