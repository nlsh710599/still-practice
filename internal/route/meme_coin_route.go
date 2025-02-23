package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/common"
	"github.com/nlsh710599/still-practice/internal/database"
	"github.com/nlsh710599/still-practice/internal/result"
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
			if database.IsNotFoundError(err) {
				c.JSON(http.StatusOK, ServerResponse[error]{
					Code: common.InvalidArgument,
					Data: common.ErrNotFound,
				})
			} else {
				c.JSON(http.StatusOK, ServerResponse[error]{
					Code: common.InternalServerError,
					Data: common.ErrInternalServer,
				})
			}
			return
		}

		c.JSON(http.StatusOK, ServerResponse[GetMemeCoinResponse]{
			Code: common.Success,
			Data: parseIntoGetMemeCoinResponse(result),
		})
	}
}

func parseIntoGetMemeCoinResponse(result result.GetMemeCoinResult) GetMemeCoinResponse {
	return GetMemeCoinResponse{
		Name:            result.Name,
		Description:     result.Description,
		PopularityScore: result.PopularityScore,
	}
}

func CreateMemeCoinRoute(memeCoinSrv service.MemeCoinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreateMemeCoinRequest
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusOK, ServerResponse[error]{
				Code: common.InvalidArgument,
				Data: common.ErrIDRequired,
			})
			return
		}

		err := memeCoinSrv.CreateMemeCoin(c.Request.Context(), request.Name, request.Description)
		if err != nil {
			if database.IsDuplicatedKeyError(err) {
				c.JSON(http.StatusOK, ServerResponse[error]{
					Code: common.InvalidArgument,
					Data: common.ErrDuplicatedKey,
				})
			} else {
				c.JSON(http.StatusOK, ServerResponse[error]{
					Code: common.InternalServerError,
					Data: common.ErrInternalServer,
				})
			}
			return
		}

		c.JSON(http.StatusOK, SuccessResponse())
	}
}
