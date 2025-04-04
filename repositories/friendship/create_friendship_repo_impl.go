package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
)

// Create implements FriendshipRepository.
func (f *FriendshipRepositoryImpl) Create(ctx context.Context, friendship models.FriendShip) (*models.FriendShip, error) {
	err := f.db.WithContext(ctx).Create(&friendship).Error
	return &friendship, err
}
