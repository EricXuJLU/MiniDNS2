package dao

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Dao struct {
	db	*gorm.DB
	r	*redis.Client
}

func NewDao(db *gorm.DB, r *redis.Client) *Dao {
	return &Dao{
		db: db,
		r: r,
	}
}