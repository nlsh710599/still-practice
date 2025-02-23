package service

import (
	"context"
	"testing"

	"github.com/nlsh710599/still-practice/internal/result"
	"github.com/stretchr/testify/assert"
)

func TestHealthService(t *testing.T) {
	t.Run("Test Health", func(t *testing.T) {
		healthService := &HealthServiceImpl{}
		healthResult, err := healthService.Health(context.Background())

		assert.Nil(t, err)
		assert.Equal(t, result.HealthResult{Status: true}, healthResult)
	})
}
