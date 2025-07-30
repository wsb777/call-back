//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	authServices "github.com/wsb777/call-back/internal/services/auth"
	userServices "github.com/wsb777/call-back/internal/services/user"

	"github.com/google/wire"
	"github.com/wsb777/call-back/http/routes"
	"github.com/wsb777/call-back/internal/config"
	"github.com/wsb777/call-back/internal/db"
	"github.com/wsb777/call-back/internal/db/repo"
	"github.com/wsb777/call-back/pkg/hasher"
	_jwt "github.com/wsb777/call-back/pkg/jwt"
)

// Application собирает необходимые компоненты приложения
type Application struct {
	HTTPServer http.Handler
	AdminInit  *AdminInitializer
}

// Провайдер для Application
func NewApplication(
	server http.Handler,
	adminInit *AdminInitializer,
) *Application {
	return &Application{
		HTTPServer: server,
		AdminInit:  adminInit,
	}
}

func InitApplication() (*Application, error) {
	wire.Build(
		// Конфиг
		config.NewConfig,
		// БД
		db.NewDatabasePG,
		db.ConnectDBProvider,
		// Репозитории
		repo.NewUserRepo,
		repo.NewJWTRepo,
		// Утилиты
		wire.Bind(new(hasher.PasswordHasher), new(*hasher.BCryptHasher)),
		hasher.NewBCryptHasher,
		wire.Bind(new(_jwt.Encoder), new(*_jwt.JWTEncoder)),
		_jwt.NewJWTEncoder,
		// Сервисы
		authServices.NewAuthService,
		authServices.NewRefreshService,
		userServices.NewUserSignUpService,
		// Инициализация админа
		NewAdminInitializer,
		// Роутеры
		routes.NewHTTPServer,
		// Создаем Application
		NewApplication,
	)
	return nil, nil
}
