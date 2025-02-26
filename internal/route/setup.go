package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
	m "github.com/nlsh710599/still-practice/internal/database/model"
	"github.com/nlsh710599/still-practice/internal/service"

	"gorm.io/gorm"
)

func Setup(r *gin.Engine, db *gorm.DB) error {
	memeCoinRepo := &m.MemeCoinRepositoryImpl{Client: db}
	err := memeCoinRepo.InitTable()
	if err != nil {
		fmt.Printf("Failed to init meme coin table: %v\n", err)
		return err
	}

	healthService := &service.HealthServiceImpl{}

	r.GET("/health", HealthRoute(healthService))

	memeCoinGroup := r.Group("/meme-coin")
	memeCoinGroup.POST("", CreateMemeCoinRoute(memeCoinRepo))
	memeCoinGroup.GET("/:id", GetMemeCoinRoute(memeCoinRepo))
	memeCoinGroup.PUT("/:id", UpdateMemeCoinRoute(memeCoinRepo))
	memeCoinGroup.DELETE("/:id", DeleteMemeCoinRoute(memeCoinRepo))
	memeCoinGroup.PUT("/:id/poke", PokeMemeCoinRoute(memeCoinRepo))

	return nil
}
