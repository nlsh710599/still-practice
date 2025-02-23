package service

import (
	"context"
	"errors"
	"testing"
	"time"

	m "github.com/nlsh710599/still-practice/internal/database/model"
	"github.com/nlsh710599/still-practice/internal/result"
	"github.com/nlsh710599/still-practice/mocks"

	"github.com/stretchr/testify/assert"
)

func TestMemeCoinService(t *testing.T) {
	mockMemeCoinRepository := mocks.MemeCoinRepository{}

	memeCoinService := &MemeCoinServiceServiceImpl{
		MemeCoinRepo: &mockMemeCoinRepository,
	}

	t.Run("Test GetMemeCoinById - success ", func(t *testing.T) {
		ctx := context.Background()
		var ID uint = 1
		record := &m.MemeCoinEntity{
			ID:              ID,
			Name:            "testCoin",
			Description:     "",
			CreatedAt:       time.Now(),
			PopularityScore: 0,
		}

		mockMemeCoinRepository.On("GetMemeCoin", ctx, ID).Return(record, nil)

		res, err := memeCoinService.GetMemeCoinById(ctx, ID)
		assert.Nil(t, err)
		assert.Equal(t, res, result.GetMemeCoinResult{
			MemeCoinEntity: *record,
		})
		mockMemeCoinRepository.AssertExpectations(t)
	})

	t.Run("Test GetMemeCoinById - error turn ", func(t *testing.T) {
		errFromMemeCoinRepo := errors.New("whatever error from meme coin repo")
		ctx := context.Background()
		var ID uint = 2
		mockMemeCoinRepository.On("GetMemeCoin", ctx, ID).Return(nil, errFromMemeCoinRepo)

		_, err := memeCoinService.GetMemeCoinById(ctx, ID)

		assert.True(t, errors.Is(err, errFromMemeCoinRepo))
		mockMemeCoinRepository.AssertExpectations(t)
	})

	t.Run("Test CreateMemeCoin ", func(t *testing.T) {
		errFromMemeCoinRepo := errors.New("whatever error from meme coin repo")
		ctx := context.Background()
		name := "testCoin2"
		description := "testCoin2description"

		mockMemeCoinRepository.On("CreateMemeCoin", ctx, &m.MemeCoinEntity{
			Name:        name,
			Description: description,
		}).Return(errFromMemeCoinRepo)

		err := memeCoinService.CreateMemeCoin(ctx, name, description)

		assert.True(t, errors.Is(err, errFromMemeCoinRepo))
		mockMemeCoinRepository.AssertExpectations(t)
	})

	t.Run("Test UpdateMemeCoin ", func(t *testing.T) {
		errFromMemeCoinRepo := errors.New("whatever error from meme coin repo")
		ctx := context.Background()
		var ID uint = 3
		description := "testCoin2description"

		mockMemeCoinRepository.On("UpdateMemeCoin", ctx, ID, description).Return(errFromMemeCoinRepo)

		err := memeCoinService.UpdateMemeCoin(ctx, ID, description)

		assert.True(t, errors.Is(err, errFromMemeCoinRepo))
		mockMemeCoinRepository.AssertExpectations(t)
	})

	t.Run("Test DeleteMemeCoin ", func(t *testing.T) {
		errFromMemeCoinRepo := errors.New("whatever error from meme coin repo")
		ctx := context.Background()
		var ID uint = 4

		mockMemeCoinRepository.On("DeleteMemeCoin", ctx, ID).Return(errFromMemeCoinRepo)

		err := memeCoinService.DeleteMemeCoin(ctx, ID)

		assert.True(t, errors.Is(err, errFromMemeCoinRepo))
		mockMemeCoinRepository.AssertExpectations(t)
	})

	t.Run("Test PokeMemeCoin ", func(t *testing.T) {
		errFromMemeCoinRepo := errors.New("whatever error from meme coin repo")
		ctx := context.Background()
		var ID uint = 5

		mockMemeCoinRepository.On("PokeMemeCoin", ctx, ID).Return(errFromMemeCoinRepo)

		err := memeCoinService.PokeMemeCoin(ctx, ID)

		assert.True(t, errors.Is(err, errFromMemeCoinRepo))
		mockMemeCoinRepository.AssertExpectations(t)
	})

}
