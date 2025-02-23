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
	memeCoinService := &service.MemeCoinServiceServiceImpl{MemeCoinRepo: memeCoinRepo}

	healthService := &service.HealthServiceImpl{}

	r.GET("/health", HealthRoute(healthService))

	memeCoinGroup := r.Group("/meme-coin")
	memeCoinGroup.POST("", CreateMemeCoinRoute(memeCoinService))
	memeCoinGroup.GET("/:id", GetMemeCoinRoute(memeCoinService))
	memeCoinGroup.PUT("/:id", UpdateMemeCoinRoute(memeCoinService))
	memeCoinGroup.DELETE("/:id", DeleteMemeCoinRoute(memeCoinService))
	memeCoinGroup.PUT("/:id/poke", PokeMemeCoinRoute(memeCoinService))

	return nil
}
