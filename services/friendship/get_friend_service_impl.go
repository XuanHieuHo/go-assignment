package services

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
)

// GetFriendOfUser implements FriendshipService.
func (n *FriendshipServiceImpl) GetFriendOfUser(ctx context.Context, email string, reqList requests.ListRequest) (*[]models.User, error) {
	user, err := n.uow.UserRepo().GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	friendships, err := n.uow.FriendshipRepo().GetFriendOfUser(ctx, user.ID, reqList)
	if err != nil {
		return nil, err
	}

	if len(*friendships) == 0 {
		return nil, nil
	}

	var friendIDs []uint
	for _, f := range *friendships {
		if f.UserID == user.ID {
			friendIDs = append(friendIDs, f.FriendID)
		} else {
			friendIDs = append(friendIDs, f.UserID)
		}
	}

	friends, err := n.uow.UserRepo().GetUserByIDs(ctx, friendIDs)
	if err != nil {
		return nil, err
	}

	return friends, nil

}
