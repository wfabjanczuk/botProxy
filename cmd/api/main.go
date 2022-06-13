package main

import (
	"github.com/joho/godotenv"
	"github.com/wfabjanczuk/botProxy/internal/handlers"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/install", handlers.Install)
	log.Fatal(http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil))
}
