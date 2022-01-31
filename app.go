package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const APP_PORT = ":8000"

type (
	/**
	 * Struct Response default
	 */
	ResponseStatus struct {
		Code    int
		Message string
	}
)

// Handler logs using MiddleWare to do all requests
func LoggerMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//fmt.Printf("Request received: %v\n", r)
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
		nextHandler.ServeHTTP(w, r)
		//fmt.Println("Request handled: OK")
	})
}

/**
 * Response: Resource Not Implemented for All methods
 */
func ResponseResourceNotImplemented(w http.ResponseWriter) {
	resp := ResponseStatus{Code: http.StatusNotImplemented, Message: "Resource not implemented"}
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
		respStatus := ResponseStatus{Code: 200, Message: "pong"}
		respJson, err := json.MarshalIndent(respStatus, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(respJson)
	}
}

// Package MAIN
func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	mux := http.NewServeMux()
	mux.Handle("/api/ping", LoggerMiddleware(http.HandlerFunc(Ping)))
	fmt.Println("Service is running on:", APP_PORT)
	http.ListenAndServe(APP_PORT, mux)
}
