package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/common"
	"github.com/nlsh710599/still-practice/internal/service"
)

func HealthRoute(healthSrv service.HealthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		result, err := healthSrv.Health(c.Request.Context())
		if err != nil {
			_ = c.Error(err)
			return
		}
		c.JSON(http.StatusOK, ServerResponse[HealthResponse]{
			Code: common.Success,
			Data: HealthResponse{
				Status: result.Status,
			},
		})
	}
}
