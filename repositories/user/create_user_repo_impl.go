package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
)

// Create implements UserRepository.
func (u *UserRepositoryImpl) Create(ctx context.Context, user models.User) (*models.User, error) {
	err := u.db.WithContext(ctx).Create(&user).Error
	return &user, err
}
