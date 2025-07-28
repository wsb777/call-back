package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/wsb777/call-back/internal/app"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	application, err := app.InitApplication()

	if err != nil {
		log.Fatal(err)
	}
	application.AdminInit.InitAdmin()

	port := ":8080"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, application.HTTPServer); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
