package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wsb777/call-back/internal/dto"
	services "github.com/wsb777/call-back/internal/services/user"
)

type UserSignInController struct {
	service services.UserSignInService
}

func NewUserSignInController(service services.UserSignInService) *UserSignInController {
	return &UserSignInController{service: service}
}

func (c *UserSignInController) SignIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.UserSignInDto
	log.Println("Запрос на авторизацию")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	token, err := c.service.SignIn(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	log.Printf("Generated token: %s", token)
	fmt.Fprintf(w, `{"token": "%s"}`, token)
}
