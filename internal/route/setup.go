package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nlsh710599/still-practice/internal/service"
)

// Setup is a function that sets up the routes and services for the admin api
func Setup(r *gin.Engine) error {
	healthSrv := &service.HealthServiceImpl{}

	r.GET("/health", HealthCheckRoute(healthSrv))

	return nil
}
