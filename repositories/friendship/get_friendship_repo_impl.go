package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
)

// GetFriendOfUser implements FriendshipRepository.
func (f *FriendshipRepositoryImpl) GetFriendOfUser(ctx context.Context, userID uint, reqList requests.ListRequest) (*[]models.FriendShip, error) {
	limit := reqList.PageSize
	offset := (reqList.PageID - 1) * reqList.PageSize

	var friends []models.FriendShip

	err := f.db.WithContext(ctx).Limit(int(limit)).Offset(int(offset)).Where("user_id = ? OR friend_id = ?", userID, userID).Find(&friends).Error
	return &friends, err
}
