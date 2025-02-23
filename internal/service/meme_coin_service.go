package service

import (
	"context"

	m "github.com/nlsh710599/still-practice/internal/database/model"
	"github.com/nlsh710599/still-practice/internal/result"
)

type MemeCoinService interface {
	GetMemeCoinById(ctx context.Context, id uint) (result.GetMemeCoinResult, error)
	CreateMemeCoin(ctx context.Context, name string, description string) error
	UpdateMemeCoin(ctx context.Context, id uint, description string) error
}

type MemeCoinServiceServiceImpl struct {
	MemeCoinRepo m.MemeCoinRepository
}

func (service *MemeCoinServiceServiceImpl) GetMemeCoinById(ctx context.Context, id uint) (result.GetMemeCoinResult, error) {
	memeCoin, err := service.MemeCoinRepo.GetMemeCoin(ctx, id)
	if err != nil {
		return result.GetMemeCoinResult{}, err
	}
	return result.GetMemeCoinResult{
		MemeCoinEntity: *memeCoin,
	}, nil
}

func (service *MemeCoinServiceServiceImpl) CreateMemeCoin(ctx context.Context, name string, description string) error {
	memeCoin := &m.MemeCoinEntity{
		Name:        name,
		Description: description,
	}
	return service.MemeCoinRepo.CreateMemeCoin(ctx, memeCoin)
}

func (service *MemeCoinServiceServiceImpl) UpdateMemeCoin(ctx context.Context, id uint, description string) error {
	return service.MemeCoinRepo.UpdateMemeCoin(ctx, id, description)
}
