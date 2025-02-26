package route

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/common"
	"github.com/nlsh710599/still-practice/internal/database/model"
	"github.com/nlsh710599/still-practice/mocks"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_parseIntoGetMemeCoinResponse(t *testing.T) {
	memeCoin := model.MemeCoinEntity{
		ID:              uint(1),
		Name:            "testCoin",
		Description:     "testDescription",
		PopularityScore: 0,
	}

	response := parseIntoGetMemeCoinResponse(&memeCoin)

	assert.Equal(t, memeCoin.Name, response.Name)
	assert.Equal(t, memeCoin.Description, response.Description)
	assert.Equal(t, memeCoin.PopularityScore, response.PopularityScore)
}

func Test_GetMemeCoinRoute(t *testing.T) {
	mockMemeCoinRepo := &mocks.MemeCoinRepository{}

	t.Run("GetMemeCoinRoute - missing field", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)

		GetMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrMissingField.Error(), response.Data)
	})

	t.Run("GetMemeCoinRoute - record not found", func(t *testing.T) {
		var ID uint = 1
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("GetMemeCoin", c.Request.Context(), ID).Return(nil, gorm.ErrRecordNotFound)

		GetMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrNotFound.Error(), response.Data)
	})

	t.Run("GetMemeCoinRoute - internal server error", func(t *testing.T) {
		var ID uint = 2
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("GetMemeCoin", c.Request.Context(), ID).Return(nil, errors.New("whatever error from memecoin service"))

		GetMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrInternalServer.Error(), response.Data)
	})

	t.Run("GetMemeCoinRoute - success", func(t *testing.T) {
		memeCoin := &model.MemeCoinEntity{
			ID:              uint(3),
			Name:            "testCoin",
			Description:     "testDescription",
			PopularityScore: 0,
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", memeCoin.ID),
			},
		}

		mockMemeCoinRepo.On("GetMemeCoin", c.Request.Context(), memeCoin.ID).Return(memeCoin, nil)

		GetMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[GetMemeCoinResponse]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.Success, response.Code)
		assert.Equal(t, parseIntoGetMemeCoinResponse(memeCoin), response.Data)

	})
}

func Test_CreateMemeCoinRoute(t *testing.T) {
	mockMemeCoinRepo := &mocks.MemeCoinRepository{}

	t.Run("CreateMemeCoinRoute - missing field", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPost, "", nil)

		CreateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrMissingField.Error(), response.Data)
	})

	t.Run("CreateMemeCoinRoute - duplicate key", func(t *testing.T) {
		memeCoin := createMemeCoinRequest{
			Name:        "testCoin2",
			Description: "",
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		requestJson, _ := json.Marshal(memeCoin)
		c.Request, _ = http.NewRequest(http.MethodPost, "", bytes.NewBuffer(requestJson))
		c.Request.Header.Add("Content-Type", "application/json;charset=UTF-8")

		mockMemeCoinRepo.On("CreateMemeCoin", c.Request.Context(), &model.MemeCoinEntity{
			Name:        memeCoin.Name,
			Description: memeCoin.Description,
		}).Return(errors.New("duplicate key value violates unique"))

		CreateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrDuplicatedKey.Error(), response.Data)
	})

	t.Run("CreateMemeCoinRoute - internal server error", func(t *testing.T) {
		memeCoin := createMemeCoinRequest{
			Name:        "testCoin3",
			Description: "",
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		requestJson, _ := json.Marshal(memeCoin)
		c.Request, _ = http.NewRequest(http.MethodPost, "", bytes.NewBuffer(requestJson))
		c.Request.Header.Add("Content-Type", "application/json;charset=UTF-8")

		mockMemeCoinRepo.On("CreateMemeCoin", c.Request.Context(), &model.MemeCoinEntity{
			Name:        memeCoin.Name,
			Description: memeCoin.Description,
		}).Return(errors.New("whatever error from memecoin service"))

		CreateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrInternalServer.Error(), response.Data)
	})

	t.Run("CreateMemeCoinRoute - success", func(t *testing.T) {
		memeCoin := createMemeCoinRequest{
			Name:        "testCoin4",
			Description: "",
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		requestJson, _ := json.Marshal(memeCoin)
		c.Request, _ = http.NewRequest(http.MethodPost, "", bytes.NewBuffer(requestJson))
		c.Request.Header.Add("Content-Type", "application/json;charset=UTF-8")

		mockMemeCoinRepo.On("CreateMemeCoin", c.Request.Context(), &model.MemeCoinEntity{
			Name:        memeCoin.Name,
			Description: memeCoin.Description,
		}).Return(nil)

		CreateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.Success, response.Code)
	})
}

func Test_UpdateMemeCoinRoute(t *testing.T) {
	mockMemeCoinRepo := &mocks.MemeCoinRepository{}

	t.Run("UpdateMemeCoinRoute - missing field on route params", func(t *testing.T) {
		memeCoin := createMemeCoinRequest{
			Description: "testDescription",
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		requestJson, _ := json.Marshal(memeCoin)
		c.Request, _ = http.NewRequest(http.MethodPut, "", bytes.NewBuffer(requestJson))
		c.Request.Header.Add("Content-Type", "application/json;charset=UTF-8")

		UpdateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrMissingField.Error(), response.Data)
	})

	t.Run("UpdateMemeCoinRoute - missing field on request body", func(t *testing.T) {
		var ID uint = 3
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPut, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		UpdateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrMissingField.Error(), response.Data)
	})

	t.Run("UpdateMemeCoinRoute - record not found", func(t *testing.T) {
		var ID uint = 4
		memeCoin := UpdateMemeCoinRequest{
			Description: "testDescription2",
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		requestJson, _ := json.Marshal(memeCoin)
		c.Request, _ = http.NewRequest(http.MethodPut, "", bytes.NewBuffer(requestJson))
		c.Request.Header.Add("Content-Type", "application/json;charset=UTF-8")
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("UpdateMemeCoin", c.Request.Context(), ID, memeCoin.Description).Return(gorm.ErrRecordNotFound)

		UpdateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrNotFound.Error(), response.Data)
	})

	t.Run("UpdateMemeCoinRoute - internal server error", func(t *testing.T) {
		var ID uint = 5
		memeCoin := UpdateMemeCoinRequest{
			Description: "testDescription3",
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		requestJson, _ := json.Marshal(memeCoin)
		c.Request, _ = http.NewRequest(http.MethodPut, "", bytes.NewBuffer(requestJson))
		c.Request.Header.Add("Content-Type", "application/json;charset=UTF-8")
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("UpdateMemeCoin", c.Request.Context(), ID, memeCoin.Description).Return(errors.New("whatever error from memecoin service"))

		UpdateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrInternalServer.Error(), response.Data)
	})

	t.Run("UpdateMemeCoinRoute - success", func(t *testing.T) {
		var ID uint = 6
		memeCoin := UpdateMemeCoinRequest{
			Description: "testDescription4",
		}
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		requestJson, _ := json.Marshal(memeCoin)
		c.Request, _ = http.NewRequest(http.MethodPut, "", bytes.NewBuffer(requestJson))
		c.Request.Header.Add("Content-Type", "application/json;charset=UTF-8")
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("UpdateMemeCoin", c.Request.Context(), ID, memeCoin.Description).Return(nil)

		UpdateMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.Success, response.Code)
	})
}

func Test_DeleteMemeCoinRoute(t *testing.T) {
	mockMemeCoinRepo := &mocks.MemeCoinRepository{}

	t.Run("DeleteMemeCoinRoute - missing field on route params", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodDelete, "", nil)

		DeleteMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrMissingField.Error(), response.Data)
	})

	t.Run("DeleteMemeCoinRoute - record not found", func(t *testing.T) {
		var ID uint = 7
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodDelete, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("DeleteMemeCoin", c.Request.Context(), ID).Return(gorm.ErrRecordNotFound)

		DeleteMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrNotFound.Error(), response.Data)
	})

	t.Run("DeleteMemeCoinRoute - internal server error", func(t *testing.T) {
		var ID uint = 8
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodDelete, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("DeleteMemeCoin", c.Request.Context(), ID).Return(errors.New("whatever error from memecoin service"))

		DeleteMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrInternalServer.Error(), response.Data)
	})

	t.Run("DeleteMemeCoinRoute - success", func(t *testing.T) {
		var ID uint = 9
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodDelete, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("DeleteMemeCoin", c.Request.Context(), ID).Return(nil)

		DeleteMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.Success, response.Code)
	})
}

func Test_PokeMemeCoinRoute(t *testing.T) {
	mockMemeCoinRepo := &mocks.MemeCoinRepository{}

	t.Run("PokeMemeCoinRoute - missing field", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPut, "", nil)

		PokeMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrMissingField.Error(), response.Data)
	})

	t.Run("PokeMemeCoinRoute - record not found", func(t *testing.T) {
		var ID uint = 10
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPut, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("PokeMemeCoin", c.Request.Context(), ID).Return(gorm.ErrRecordNotFound)

		PokeMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrNotFound.Error(), response.Data)
	})

	t.Run("PokeMemeCoinRoute - internal server error", func(t *testing.T) {
		var ID uint = 11
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPut, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("PokeMemeCoin", c.Request.Context(), ID).Return(errors.New("whatever error from memecoin service"))

		PokeMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.ErrInternalServer.Error(), response.Data)
	})

	t.Run("PokeMemeCoinRoute - success", func(t *testing.T) {
		var ID uint = 12
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodPut, "", nil)
		c.Params = gin.Params{
			{
				Key:   "id",
				Value: fmt.Sprintf("%v", ID),
			},
		}

		mockMemeCoinRepo.On("PokeMemeCoin", c.Request.Context(), ID).Return(nil)

		PokeMemeCoinRoute(mockMemeCoinRepo)(c)

		res := w.Result()
		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response ServerResponse[string]
		body, _ := io.ReadAll(res.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.Success, response.Code)
	})
}
