package service

import (
	"context"

	"github.com/nlsh710599/still-practice/internal/result"
)

type HealthService interface {
	Health(ctx context.Context) (result.HealthResult, error)
}

type HealthServiceImpl struct {
}

func (service *HealthServiceImpl) Health(ctx context.Context) (result.HealthResult, error) {
	return result.HealthResult{Status: true}, nil
}
