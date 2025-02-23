package route

import (
	"github.com/nlsh710599/still-practice/internal/common"
)

func SuccessResponse() ServerResponse[string] {
	return ServerResponse[string]{
		Code: common.Success,
		Data: "Success",
	}
}

type ServerResponse[T any] struct {
	Code int `json:"code"`
	Data T   `json:"data"`
}

type HealthResponse struct {
	Status bool `json:"status"`
}

type GetMemeCoinResponse struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	PopularityScore int    `json:"popularity_score"`
}
