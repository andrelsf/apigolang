package game

import (
	"database/sql"
	"fmt"

	gameRepository "github.com/andrelsf/apigolang/repository"
)

type mysqlGameRepository struct {
	Conn *sql.DB
}

// Retorna a interface Game implementado pelo repository
func NewSQLGameRepository(Conn *sql.DB) gameRepository.GameRepository {
	fmt.Println("")
}
