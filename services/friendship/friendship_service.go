package services

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
)

type FriendshipService interface {
	GetFriendOfUser(ctx context.Context, email string, reqList requests.ListRequest) (*[]models.User, error)
	CreateFriendship(ctx context.Context, req requests.CreateFriendshipRequest) (*models.FriendShip, error)
}
