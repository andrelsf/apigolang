package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andrelsf/apigolang/driver"
	models "github.com/andrelsf/apigolang/models"
	repository "github.com/andrelsf/apigolang/repository"
	game "github.com/andrelsf/apigolang/repository/game"
)

// Game ...
type Game struct {
	repo repository.GameRepository
}

// Manager connection for new games
func NewGameHandler(db *driver.DB) *Game {
	return &Game{
		repo: game.NewSQLGameRepo(db.SQL),
	}
}

// Write JSON response format
func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.MarshalIndent(payload, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Return error messages
func responseWithError(w http.ResponseWriter, code int, message string) {
	responseWithJSON(w, code, map[string]string{"message": message})
}

// Fetch all games data
func (g *Game) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := g.repo.Fetch(r.Context(), 10)
	responseWithJSON(w, http.StatusOK, payload)
}

// Create new GAME
func (g *Game) Create(w http.ResponseWriter, r *http.Request) {
	game := models.Game{}

	err := json.NewDecoder(r.Body).Decode(&game)

	_, err := g.repo.Create(r.Context(), &game) //newID
	if err != nil {
		log.Printf("%s", err)
		responseWithError(w, http.StatusInternalServerError, "Server error")
	}

	responseWithJSON(w, http.StatusCreated, map[string]string{"message": "Successfully created"})
}

// Update a game by id
func (g *Game) Update(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusNotImplemented, map[string]string{"message": "Resource not implemented"})
}

// Retorna a game details by id
func (g *Game) GetById(w http.ResponseWriter, r *http.Request) {
	//TODO
	payload, err := g.repo.GetByID(r.Context(), int64(1))

	if err != nil {
		responseWithError(w, http.StatusNoContent, "Content not found")
	}

	responseWithJSON(w, http.StatusOK, payload)
}

// Remove a game by id
func (g *Game) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO
	_, err := g.repo.Delete(r.Context(), int64(1))

	if err != nil {
		responseWithError(w, http.StatusInternalServerError, "Error try delete by id")
	}

	responseWithJSON(w, http.StatusMovedPermanently, map[string]string{"message": "Delete successfully"})
}
