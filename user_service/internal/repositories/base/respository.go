package base

import (
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"sweete/core/pkg/mysql"
)

type Model interface {
	TableName() string
	//Preload(query *gorm.DB, preload interface{}) *gorm.DB
	//QueryByParam(query *gorm.DB, param interface{}) *gorm.DB
}

type BaseRepository struct {
	Model Model
}

func NewBaseRepository(model Model) *BaseRepository {
	return &BaseRepository{Model: model}
}

const keyDB = "db"

func GetDBFromContext(ctx context.Context) *gorm.DB {
	val, ok := ctx.Value(keyDB).(*gorm.DB)

	if !ok {
		return mysql.GetInstance()
	}
	val.Debug()
	return val
}
func SetDBWithContext(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, keyDB, db)
}
func (r *BaseRepository) GetDB(ctx context.Context) *gorm.DB {
	return GetDBFromContext(ctx)
}
func (r *BaseRepository) BeginTransaction(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Begin()
}
func (r *BaseRepository) BeginTransactionWithContext(ctx context.Context) context.Context {
	db := r.BeginTransaction(ctx)
	return SetDBWithContext(ctx, db)
}
func (r *BaseRepository) Commit(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Commit()
}
func (r *BaseRepository) Rollback(ctx context.Context) *gorm.DB {
	db := r.GetDB(ctx)
	return db.Rollback()
}

type LimitParam struct {
	Limit  *int
	Page   *int
	Offset *int
}

func (r *BaseRepository) LimitOffset(db *gorm.DB, param LimitParam) *gorm.DB {
	if param.Limit != nil {
		db = db.Limit(*param.Limit)
	}
	if param.Offset != nil {
		db = db.Offset(*param.Offset)
	} else if param.Page != nil && param.Limit != nil {
		offset := (*param.Page - 1) * (*param.Limit)
		db = db.Offset(offset)
	}
	return db
}
