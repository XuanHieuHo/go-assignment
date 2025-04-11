package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	pkgerrors "github.com/pkg/errors"
)

// Create implements FriendshipRepository.
func (f *FriendshipRepositoryImpl) Create(ctx context.Context, friendship models.FriendShip) (*models.FriendShip, error) {
	err := f.db.WithContext(ctx).Create(&friendship).Error
	return &friendship, pkgerrors.WithStack(err)
}
