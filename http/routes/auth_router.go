package routes

import (
	"net/http"

	services "github.com/wsb777/call-back/internal/services/auth"

	"github.com/wsb777/call-back/http/controllers/auth"
)

func AuthRoutes(router *http.ServeMux, AuthService services.AuthService, RefreshService services.RefreshService) {
	authController := auth.NewAuthController(AuthService)
	refreshController := auth.NewRefreshController(RefreshService)
	router.HandleFunc("/api/v1/auth", authController.SignIn)
	router.HandleFunc("/api/v1/refresh_token", refreshController.RefreshToken)
	router.HandleFunc("/api/v1/logout", authController.SignIn)
}
