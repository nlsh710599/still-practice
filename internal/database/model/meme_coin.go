package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type MemeCoinEntity struct {
	ID              uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name            string    `gorm:"unique;not null" json:"name"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	PopularityScore int       `json:"popularity_score"`
}

type MemeCoinRepository interface {
	CreateMemeCoin(ctx context.Context, memeCoin *MemeCoinEntity) error
	GetMemeCoin(ctx context.Context, id uint) (*MemeCoinEntity, error)
	UpdateMemeCoin(ctx context.Context, memeCoin *MemeCoinEntity) error
	DeleteMemeCoin(ctx context.Context, id uint) error
	PokeMemeCoin(ctx context.Context, id uint) error
}

type MemeCoinRepositoryImpl struct {
	Client *gorm.DB
}
