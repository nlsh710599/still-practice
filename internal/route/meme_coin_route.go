package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/common"
	"github.com/nlsh710599/still-practice/internal/service"
)

func GetMemeCoinRoute(memeCoinSrv service.MemeCoinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params getMemeCoinParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusOK, ServerResponse[error]{
				Code: common.InvalidArgument,
				Data: common.ErrIDRequired,
			})
			return
		}

		result, err := memeCoinSrv.GetMemeCoinById(c.Request.Context(), params.ID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, ServerResponse[GetMemeCoinResponse]{
			Code: common.Success,
			Data: GetMemeCoinResponse{
				MemeCoinEntity: result.MemeCoinEntity,
			},
		})
	}
}
