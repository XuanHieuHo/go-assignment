package uow

// import (
// 	"context"

// 	"gorm.io/gorm"
// )

// type Manager struct {
// 	db *gorm.DB
// }

// func NewManager(db *gorm.DB) *Manager {
// 	return &Manager{db: db}
// }

// func (m *Manager) Do(ctx context.Context, fn func(uow UnitOfWork) error) error {
// 	return m.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
// 		u := &UnitOfWorkImpl{db: tx}
// 		return fn(u)
// 	})
// }

// func (m *Manager) New(ctx context.Context) UnitOfWork {
// 	return &UnitOfWorkImpl{db: m.db.WithContext(ctx)}
// }