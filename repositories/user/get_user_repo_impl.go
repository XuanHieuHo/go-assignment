package repositories

import (
	"context"

	"github.com/XuanHieuHo/go-assignment/models"
	"github.com/XuanHieuHo/go-assignment/requests"
	pkgerrors "github.com/pkg/errors"
)

// GetUserByEmail implements UserRepository.
func (u *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, pkgerrors.WithStack(err)
}

// GetUserByID implements UserRepository.
func (u *UserRepositoryImpl) GetUserByID(ctx context.Context, ID uint) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("id = ?", ID).First(&user).Error
	return &user, pkgerrors.WithStack(err)
}

// GetUserByIDs implements UserRepository.
func (u *UserRepositoryImpl) GetUserByIDs(ctx context.Context, IDs []uint) (*[]models.User, error) {
	var users []models.User

	err := u.db.WithContext(ctx).Where("id IN ?", IDs).Find(&users).Error
	return &users, pkgerrors.WithStack(err)
}

// ListUser implements UserRepository.
func (u *UserRepositoryImpl) ListUser(ctx context.Context, req requests.ListRequest) (*[]models.User, error) {
	limit := req.PageSize
	offset := (req.PageID - 1) * req.PageSize
	var users []models.User
	err := u.db.WithContext(ctx).Limit(int(limit)).Offset(int(offset)).Find(&users).Error
	return &users, pkgerrors.WithStack(err)
}
