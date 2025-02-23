package model

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/nlsh710599/still-practice/internal/database"
	"github.com/stretchr/testify/assert"
)

func Test_MemeCoinRepository(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pg, err := database.NewPostgres(os.Getenv("DSN"))
	assert.NoError(t, err)

	pg.Migrator().DropTable(
		&MemeCoinEntity{},
	)

	repo := &MemeCoinRepositoryImpl{
		Client: pg,
	}
	assert.NoError(t, repo.InitTable())

	memeCoin := &MemeCoinEntity{
		Name:      "testCoin",
		CreatedAt: time.Now(),
	}

	t.Run("CreateMemeCoin", func(t *testing.T) {
		err = repo.CreateMemeCoin(ctx, memeCoin)
		assert.NoError(t, err)
		assert.Equal(t, memeCoin.ID, uint(1))
		assert.Equal(t, memeCoin.PopularityScore, int(0))

		err = repo.CreateMemeCoin(ctx, memeCoin)
		assert.Error(t, err)
		assert.Equal(t, database.IsDuplicatedKeyError(err), true)

	})

	t.Run("GetMemeCoin", func(t *testing.T) {
		retrieved, err := repo.GetMemeCoin(ctx, memeCoin.ID)
		assert.NoError(t, err)
		assert.Equal(t, memeCoin.Name, retrieved.Name)
		var emptyStr string
		assert.Equal(t, memeCoin.Description, emptyStr)

		var NotFoundID uint = 999
		retrieved, err = repo.GetMemeCoin(ctx, NotFoundID)
		assert.Error(t, err)
		assert.Equal(t, database.IsNotFoundError(err), true)
		assert.Nil(t, retrieved)
	})

	t.Run("UpdateMemeCoin", func(t *testing.T) {
		updatedDescription := "avada kedavra"
		err = repo.UpdateMemeCoin(ctx, memeCoin.ID, updatedDescription)
		assert.NoError(t, err)

		retrieved, err := repo.GetMemeCoin(ctx, memeCoin.ID)
		assert.NoError(t, err)
		assert.Equal(t, updatedDescription, retrieved.Description)

		var NotFoundID uint = 999
		err = repo.UpdateMemeCoin(ctx, NotFoundID, updatedDescription)
		assert.Error(t, err)
		assert.Equal(t, database.IsNotFoundError(err), true)
	})

	t.Run("PokeMemeCoin", func(t *testing.T) {
		err = repo.PokeMemeCoin(ctx, memeCoin.ID)
		assert.NoError(t, err)

		retrieved, err := repo.GetMemeCoin(ctx, memeCoin.ID)
		assert.NoError(t, err)
		assert.Equal(t, 1, retrieved.PopularityScore)

		var NotFoundID uint = 999
		err = repo.PokeMemeCoin(ctx, NotFoundID)
		assert.Error(t, err)
		assert.Equal(t, database.IsNotFoundError(err), true)
	})

	t.Run("DeleteMemeCoin", func(t *testing.T) {
		err = repo.DeleteMemeCoin(ctx, memeCoin.ID)
		assert.NoError(t, err)

		retrieved, err := repo.GetMemeCoin(ctx, memeCoin.ID)
		assert.Error(t, err)
		assert.Equal(t, database.IsNotFoundError(err), true)
		assert.Nil(t, retrieved)

		var NotFoundID uint = 999
		err = repo.DeleteMemeCoin(ctx, NotFoundID)
		assert.Error(t, err)
		assert.Equal(t, database.IsNotFoundError(err), true)
	})
}
