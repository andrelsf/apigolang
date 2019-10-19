package repository

import (
	"context"

	"github.com/andrelsf/apigolang/models"
)

type GameRepository interface {
	Fetch(ctx context.Context, num int64) ([]*models.Game, error)
	GetById(ctx context.Context, id int64) (*models.Game, error)
	Create(ctx context.Context, g *models.Game) (int64, error)
	Update(ctx context.Context, g *models.Game) (*models.Game, error)
	Delete(ctx context.Context, id int64) (bool, error)
}
