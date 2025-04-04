package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
)

type UserRepository interface {
	Create(ctx context.Context, user models.User) (*models.User, error)
	ListUser(ctx context.Context, req requests.ListRequest) (*[]models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, ID uint) (*models.User, error)
	GetUserByIDs(ctx context.Context, IDs []uint) (*[]models.User, error)
}
