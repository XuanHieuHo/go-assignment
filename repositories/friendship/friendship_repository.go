package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
)

type FriendshipRepository interface {
	GetFriendOfUser(ctx context.Context, userID uint, reqList requests.ListRequest) (*[]models.FriendShip, error)
	Create(ctx context.Context, friendship models.FriendShip) (*models.FriendShip, error)
	IsFriendshipExists(ctx context.Context, userID, friendID uint) (bool, error)
}
