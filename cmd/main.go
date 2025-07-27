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

	handler, err := app.InitHttpServer()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server is running on :8080")

	log.Fatal(http.ListenAndServe(":8080", handler))
}
