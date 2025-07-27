package user

import (
	"encoding/json"
	"log"
	"net/http"

	services "github.com/wsb777/call-back/internal/services/user"

	"github.com/wsb777/call-back/internal/dto"
)

type UserSignUpController struct {
	userService services.UserSignUpService
}

func NewUserSignUpController(userService services.UserSignUpService) *UserSignUpController {
	return &UserSignUpController{
		userService: userService,
	}
}

func (c *UserSignUpController) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.UserSignUpDto
	log.Println("Запрос на регистрацию")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	err := c.userService.CreateUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
