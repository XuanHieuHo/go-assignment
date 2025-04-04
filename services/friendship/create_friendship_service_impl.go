package services

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
	"github.com/XuanHieuHo/go-assignment/uow"

	pkgerrors "github.com/pkg/errors"
)

func (n *FriendshipServiceImpl) CreateFriendship(ctx context.Context, req requests.CreateFriendshipRequest) (*models.FriendShip, error) {

	var createdFriendship *models.FriendShip
	if err := n.Manager.Do(ctx, func(uow uow.UnitOfWork) error {
		user1, err := uow.UserRepo().GetUserByEmail(ctx, req.FristEmail)
		if err != nil {
			return err
		}

		user2, err := uow.UserRepo().GetUserByEmail(ctx, req.SecondEmail)
		if err != nil {
			return err
		}

		if user1.ID == user2.ID {
			return pkgerrors.New("cannot create friendship with yourself")
		}

		fristUser, secondUser := user1, user2
		if user1.ID > user2.ID {
			fristUser, secondUser = user2, user1
		}

		isFriend, err := uow.FriendshipRepo().IsFriendshipExists(ctx, fristUser.ID, secondUser.ID)
		if err != nil {
			return err
		}
		if isFriend {
			return pkgerrors.New("friendship already exists")
		}

		friendship := models.FriendShip{
			UserID:   fristUser.ID,
			User:     *fristUser,
			FriendID: secondUser.ID,
			Friend:   *secondUser,
		}

		createdFriendship, err = uow.FriendshipRepo().Create(ctx, friendship)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return createdFriendship, nil
}
