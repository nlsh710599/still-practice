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
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		result, err := memeCoinSrv.GetMemeCoinById(c.Request.Context(), params.ID)
		if err != nil {
			if database.IsNotFoundError(err) {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InvalidArgument,
					Data: common.ErrNotFound.Error(),
				})
			} else {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InternalServerError,
					Data: common.ErrInternalServer.Error(),
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

func parseIntoGetMemeCoinResponse(result *result.GetMemeCoinResult) GetMemeCoinResponse {
	return GetMemeCoinResponse{
		Name:            result.Name,
		Description:     result.Description,
		PopularityScore: result.PopularityScore,
	}
}

func CreateMemeCoinRoute(memeCoinSrv service.MemeCoinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createMemeCoinRequest
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		err := memeCoinSrv.CreateMemeCoin(c.Request.Context(), request.Name, request.Description)
		if err != nil {
			if database.IsDuplicatedKeyError(err) {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InvalidArgument,
					Data: common.ErrDuplicatedKey.Error(),
				})
			} else {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InternalServerError,
					Data: common.ErrInternalServer.Error(),
				})
			}
			return
		}

		c.JSON(http.StatusOK, SuccessResponse())
	}
}

func UpdateMemeCoinRoute(memeCoinSrv service.MemeCoinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params updateMemeCoinParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		var request UpdateMemeCoinRequest
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		err := memeCoinSrv.UpdateMemeCoin(c.Request.Context(), params.ID, request.Description)
		if err != nil {
			if database.IsNotFoundError(err) {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InvalidArgument,
					Data: common.ErrNotFound.Error(),
				})
			} else {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InternalServerError,
					Data: common.ErrInternalServer.Error(),
				})
			}
			return
		}

		c.JSON(http.StatusOK, SuccessResponse())
	}
}

func DeleteMemeCoinRoute(memeCoinSrv service.MemeCoinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params deleteMemeCoinParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		err := memeCoinSrv.DeleteMemeCoin(c.Request.Context(), params.ID)
		if err != nil {
			if database.IsNotFoundError(err) {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InvalidArgument,
					Data: common.ErrNotFound.Error(),
				})
			} else {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InternalServerError,
					Data: common.ErrInternalServer.Error(),
				})
			}
			return
		}

		c.JSON(http.StatusOK, SuccessResponse())
	}
}

func PokeMemeCoinRoute(memeCoinSrv service.MemeCoinService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params pokeMemeCoinParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		err := memeCoinSrv.PokeMemeCoin(c.Request.Context(), params.ID)
		if err != nil {
			if database.IsNotFoundError(err) {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InvalidArgument,
					Data: common.ErrNotFound.Error(),
				})
			} else {
				c.JSON(http.StatusOK, ServerResponse[string]{
					Code: common.InternalServerError,
					Data: common.ErrInternalServer.Error(),
				})
			}
			return
		}

		c.JSON(http.StatusOK, SuccessResponse())
	}
}
