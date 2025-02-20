package repositories

import (
	"context"
	"gorm.io/gorm"
)

type IBaseRepository interface {
	BeginTransaction(ctx context.Context) *gorm.DB
	Commit(ctx context.Context) *gorm.DB
	Rollback(ctx context.Context) *gorm.DB
	BeginTransactionWithContext(ctx context.Context) context.Context
}
