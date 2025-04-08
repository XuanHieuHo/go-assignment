package services

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/uow"
)

// GetListUser implements UserService.
func (u *UserServiceImpl) GetListUser(ctx context.Context, req requests.ListRequest) (*[]models.User, error) {
	return uow.New(u.db, ctx).UserRepo().ListUser(ctx, req)
}
