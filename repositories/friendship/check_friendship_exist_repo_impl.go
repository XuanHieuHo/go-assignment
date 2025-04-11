package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	pkgerrors "github.com/pkg/errors"
)

// IsFriendshipExists implements FriendshipRepository.
func (f *FriendshipRepositoryImpl) IsFriendshipExists(ctx context.Context, userID uint, friendID uint) (bool, error) {
	var count int64

	err := f.db.WithContext(ctx).Model(&models.FriendShip{}).Where("user_id = ? AND friend_id = ?", userID, friendID).Count(&count).Error
	return count > 0, pkgerrors.WithStack(err)
}
