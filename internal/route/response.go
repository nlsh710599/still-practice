package route

import (
	"github.com/nlsh710599/still-practice/internal/database/model"
)

type ServerResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

type HealthResponse struct {
	Status bool `json:"status"`
}

type GetMemeCoinResponse struct {
	model.MemeCoinEntity
}
