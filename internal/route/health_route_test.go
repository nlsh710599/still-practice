package route

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/common"
	"github.com/nlsh710599/still-practice/internal/result"
	"github.com/nlsh710599/still-practice/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_health_route(t *testing.T) {
	mockHealthService := mocks.HealthService{}
	healthRoute := HealthRoute(&mockHealthService)

	t.Run("health", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)

		mockHealthService.On("Health", c.Request.Context()).Return(result.HealthResult{
			Status: true,
		}, nil)

		healthRoute(c)

		result := w.Result()
		assert.Equal(t, http.StatusOK, result.StatusCode)

		var response ServerResponse[HealthResponse]
		body, _ := io.ReadAll(result.Body)
		json.Unmarshal(body, &response)
		assert.Equal(t, common.Success, response.Code)
	})
}
