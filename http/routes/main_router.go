package routes

import (
	"fmt"
	"net/http"

	"github.com/wsb777/call-back/http/middleware"
	services "github.com/wsb777/call-back/internal/services/user"
	"github.com/wsb777/call-back/pkg/jwt"
)

func NewHTTPServer(userSignUpService services.UserSignUpService, userSignInService services.UserSignInService, jwtEncoder jwt.Encoder) http.Handler {
	mux := http.NewServeMux()

	// Корневой обработчик
	mux.HandleFunc("/", (func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, "Database version: %s", "Ураура")
	}))
	// Регистрация пользовательских маршрутов
	UserRoutes(mux, userSignUpService, jwtEncoder)
	AuthRoutes(mux, userSignInService)
	middleServer := middleware.AllInfoMiddleware(mux)
	return middleServer
}
