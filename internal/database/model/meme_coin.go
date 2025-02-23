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
	PopularityScore int       `gorm:"default:0" json:"popularity_score"`
}

type MemeCoinRepository interface {
	InitTable() error
	GetMemeCoin(ctx context.Context, id uint) (*MemeCoinEntity, error)
	CreateMemeCoin(ctx context.Context, memeCoin *MemeCoinEntity) error
	UpdateMemeCoin(ctx context.Context, id uint, description string) error
	DeleteMemeCoin(ctx context.Context, id uint) error
}

type MemeCoinRepositoryImpl struct {
	Client *gorm.DB
}

func (repo *MemeCoinRepositoryImpl) InitTable() error {
	return repo.Client.AutoMigrate(&MemeCoinEntity{})
}

func (repo *MemeCoinRepositoryImpl) GetMemeCoin(ctx context.Context, id uint) (*MemeCoinEntity, error) {
	var memeCoin MemeCoinEntity
	err := repo.Client.WithContext(ctx).Where("id = ?", id).First(&memeCoin).Error
	if err != nil {
		return nil, err
	}
	return &memeCoin, nil
}

func (repo *MemeCoinRepositoryImpl) CreateMemeCoin(ctx context.Context, memeCoin *MemeCoinEntity) error {
	return repo.Client.WithContext(ctx).Create(memeCoin).Error
}

func (repo *MemeCoinRepositoryImpl) UpdateMemeCoin(ctx context.Context, id uint, description string) error {
	res := repo.Client.WithContext(ctx).Model(&MemeCoinEntity{}).Where("id = ?", id).Update("description", description)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (repo *MemeCoinRepositoryImpl) DeleteMemeCoin(ctx context.Context, id uint) error {
	res := repo.Client.WithContext(ctx).Where("id = ?", id).Delete(&MemeCoinEntity{})
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
