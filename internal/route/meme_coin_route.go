package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/common"
	"github.com/nlsh710599/still-practice/internal/database"
	m "github.com/nlsh710599/still-practice/internal/database/model"
)

func GetMemeCoinRoute(memeCoinRepo m.MemeCoinRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params getMemeCoinParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		result, err := memeCoinRepo.GetMemeCoin(c.Request.Context(), params.ID)
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

func parseIntoGetMemeCoinResponse(result *m.MemeCoinEntity) GetMemeCoinResponse {
	return GetMemeCoinResponse{
		Name:            result.Name,
		Description:     result.Description,
		PopularityScore: result.PopularityScore,
	}
}

func CreateMemeCoinRoute(memeCoinRepo m.MemeCoinRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request createMemeCoinRequest
		if err := c.ShouldBind(&request); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		memeCoin := &m.MemeCoinEntity{
			Name:        request.Name,
			Description: request.Description,
		}

		err := memeCoinRepo.CreateMemeCoin(c.Request.Context(), memeCoin)
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

func UpdateMemeCoinRoute(memeCoinRepo m.MemeCoinRepository) gin.HandlerFunc {
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

		err := memeCoinRepo.UpdateMemeCoin(c.Request.Context(), params.ID, request.Description)
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

func DeleteMemeCoinRoute(memeCoinRepo m.MemeCoinRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params deleteMemeCoinParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		err := memeCoinRepo.DeleteMemeCoin(c.Request.Context(), params.ID)
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

func PokeMemeCoinRoute(memeCoinRepo m.MemeCoinRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params pokeMemeCoinParams
		if err := c.ShouldBindUri(&params); err != nil {
			c.JSON(http.StatusOK, ServerResponse[string]{
				Code: common.InvalidArgument,
				Data: common.ErrMissingField.Error(),
			})
			return
		}

		err := memeCoinRepo.PokeMemeCoin(c.Request.Context(), params.ID)
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
