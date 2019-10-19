package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	models "github.com/andrelsf/apigolang/models"
)

const (
	APP_PORT = ":8000"
	LOG_PATH = "./data/log/app.log"
)

type (
	/**
	 * Struct Response default
	 */
	ResponseStatus struct {
		StatusCode int
		Message    string
	}
)

// Handler logs using MiddleWare to do all requests
func LoggerMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %v\n", r)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		nextHandler.ServeHTTP(w, r)
		fmt.Println("Request handled: OK")
	})
}

// Open Log File
func OpenLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)
		if err != nil {
			log.Fatal("OpenLogFile: os.OpenFile:", err)
		}
		log.SetOutput(lf)
	}
}

/**
 * Response: Resource Not Implemented for All methods
 */
func ResponseResourceNotImplemented(w http.ResponseWriter) {
	resp := ResponseStatus{StatusCode: http.StatusNotImplemented, Message: "Resource not implemented"}
	respJson, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	w.Write(respJson)
}

/**
 * GET: /api/ping
 * @CURL:
 * 		curl -X GET -H "Content-type: application/json" localhost:8000/api/ping
 */
func Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		respStatus := ResponseStatus{StatusCode: 200, Message: "pong"}
		respJson, err := json.MarshalIndent(respStatus, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	}
}

/***
 * POST: /api/game
 * @CURL:
 * 		curl -X POST -H "Content-type: application/json" \
 * 		-d '{"name":"GOD OF WAR IV","platform":"PS4","description":"Kratos adventure in Nordic lands with his son Atreus","price":"99.90"}' \
 * 		localhost:8000
 */
func Games(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		game := models.Game{}

		err := json.NewDecoder(r.Body).Decode(&game)
		if err != nil {
			log.Fatal(err)
		}
		game.CreateAt = time.Now().UTC()

		// Handler Database start
		// TODO
		// fim database
		gameJson, err := json.MarshalIndent(&game, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(gameJson)
	} else {
		ResponseResourceNotImplemented(w)
	}
}

// Package MAIN
func main() {
	/*dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := mysql.ConnectSQL(dbHost, dbPort, dbUser, dbPass, dbName)
	*/

	OpenLogFile(LOG_PATH)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mux := http.NewServeMux()
	mux.Handle("/api/ping", LoggerMiddleware(http.HandlerFunc(Ping)))
	mux.Handle("/api/game", LoggerMiddleware(http.HandlerFunc(Games)))
	fmt.Println("Service is running on:", APP_PORT)
	http.ListenAndServe(APP_PORT, mux)
}
