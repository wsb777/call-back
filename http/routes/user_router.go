package routes

import (
	services "github.com/wsb777/call-back/internal/services/user"
	"net/http"

	"github.com/wsb777/call-back/http/controllers/user"
)

func UserRoutes(router *http.ServeMux, userSignUpService services.UserSignUpService, userSignInService services.UserSignInService) {
	signInController := user.NewUserSignInController(userSignInService)
	signUpController := user.NewUserSignUpController(userSignUpService)
	router.HandleFunc("/api/v1/users/register", signUpController.Register)
	router.HandleFunc("/api/v1/users/login", signInController.SignIn)
}