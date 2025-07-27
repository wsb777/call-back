package routes

import (
	services "github.com/wsb777/call-back/internal/services/user"
	"fmt"
	"net/http"
	"github.com/wsb777/call-back/pkg/middleware"
)

func NewHTTPServer(userSignUpService services.UserSignUpService, userSignInService services.UserSignInService) http.Handler {
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
	UserRoutes(mux, userSignUpService, userSignInService)
	middleServer := middleware.AllInfoMiddleware(mux)
	return middleServer
}