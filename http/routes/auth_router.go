package routes

import (
	"net/http"

	services "github.com/wsb777/call-back/internal/services/user"

	"github.com/wsb777/call-back/http/controllers/user"
)

func AuthRoutes(router *http.ServeMux, userSignInService services.UserSignInService) {
	signInController := user.NewUserSignInController(userSignInService)
	router.HandleFunc("/api/v1/auth", signInController.SignIn)
}
