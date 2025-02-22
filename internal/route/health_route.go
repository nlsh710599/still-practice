package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/common"
	"github.com/nlsh710599/still-practice/internal/service"
)

type HealthCheckResponse struct {
	Status bool `json:"status"`
}

func HealthCheckRoute(healthCheckSrv service.HealthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := healthCheckSrv.Health(c.Request.Context())
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(http.StatusOK, ServerResponse[HealthCheckResponse]{
			Code: common.Success,
			Data: HealthCheckResponse{
				Status: result.Status,
			},
		})
	}
}
