package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wsb777/call-back/internal/dto"
	services "github.com/wsb777/call-back/internal/services/auth"
)

type RefreshController struct {
	service services.RefreshService
}

func NewRefreshController(service services.RefreshService) *RefreshController {
	return &RefreshController{service: service}
}

func (c *RefreshController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.RefreshTokenDto
	log.Println("Запрос на генерацию нового токена")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := c.service.RefreshToken(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"access_token": "%s", "refresh_token": "%s"}`, accessToken, refreshToken)
}
