package game

import (
	"context"
	"database/sql"

	models "github.com/andrelsf/apigolang/models"
	gameRepository "github.com/andrelsf/apigolang/repository"
)

type mysqlGameRepository struct {
	Connection *sql.DB
}

// Retorna a interface Game implementado pelo repository
func NewSQLGameRepository(Conn *sql.DB) gameRepository.GameRepository {
	return &mysqlGameRepository{
		Connection: Conn,
	}
}

func (m *mysqlGameRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Game, error) {
	rows, err := m.Connection.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Game, 0)
	for rows.Next() {
		data := new(models.Game)
		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.Platform,
			&data.Description,
			&data.Price,
			&data.CreateAt,
			&data.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlGameRepository) Fetch(ctx context.Context, num int64) ([]*models.Game, error) {
	query := "SELECT * FROM games LIMIT ?"
	return m.fetch(ctx, query, num)
}

func (m *mysqlGameRepository) Create(ctx context.Context, g *models.Game) (int64, error) {
	query := "INSERT INTO games (name, platform, description, price, createat, updateat) VALUES ('?', '?', '?', '?', '?', '?')"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, g.Name, g.Platform, g.Description, g.Price, g.CreateAt, g.UpdateAt)
	defer stmt.Close()

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlGameRepository) GetByID(ctx context.Context, id int64) (*models.Game, error) {
	query := "SELECT * FROM games WHERE id = ?"

	row, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Game{}
	if len(row) > 0 {
		payload = row[0]
	}
	return payload, nil
}

func (m *mysqlGameRepository) Update(ctx context.Context, g *models.Game) (*models.Game, error) {
	query := "UPDATE games SET name = '?', platform = '?', description = '?', price = '?', updateat = '?' WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.ExecContext(
		ctx,
		g.Name,
		g.Platform,
		g.Description,
		g.Price,
		g.UpdateAt,
		g.Id,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	return g, nil
}

func (m *mysqlGameRepository) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM games WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
