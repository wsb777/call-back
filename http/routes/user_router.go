package routes

import (
	"net/http"

	"github.com/wsb777/call-back/http/controllers/user"
	"github.com/wsb777/call-back/http/middleware"
	services "github.com/wsb777/call-back/internal/services/user"
	_jwt "github.com/wsb777/call-back/pkg/jwt"
)

func UserRoutes(router *http.ServeMux, userSignUpService services.UserSignUpService, jwtEncoder _jwt.Encoder) {
	signUpController := user.NewUserSignUpController(userSignUpService)

	protectedHandler := middleware.AuthMiddleware(
		http.HandlerFunc(signUpController.Register),
		jwtEncoder,
	)

	router.Handle("/api/v1/users/register", protectedHandler)
}
