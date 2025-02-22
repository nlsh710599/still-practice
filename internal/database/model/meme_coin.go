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
	CreatedAt       time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;not null"`
	PopularityScore int       `json:"popularity_score"`
}

type MemeCoinRepository interface {
	InitTable() error
	GetMemeCoin(ctx context.Context, id uint) (*MemeCoinEntity, error)
}

type MemeCoinRepositoryImpl struct {
	Client *gorm.DB
}

func (repo *MemeCoinRepositoryImpl) GetMemeCoin(ctx context.Context, id uint) (*MemeCoinEntity, error) {
	var memeCoin MemeCoinEntity
	err := repo.Client.WithContext(ctx).Where("id = ?", id).First(&memeCoin).Error
	if err != nil {
		return nil, err
	}
	return &memeCoin, nil
}

func (repo *MemeCoinRepositoryImpl) InitTable() error {
	return repo.Client.AutoMigrate(&MemeCoinEntity{})
}
