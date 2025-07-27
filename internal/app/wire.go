//go:build wireinject
// +build wireinject

package app

import (
	"net/http"

	services "github.com/wsb777/call-back/internal/services/user"

	"github.com/google/wire"
	"github.com/wsb777/call-back/http/routes"
	"github.com/wsb777/call-back/internal/config"
	"github.com/wsb777/call-back/internal/db"
	"github.com/wsb777/call-back/internal/db/repo"
	"github.com/wsb777/call-back/pkg/hasher"
	_jwt "github.com/wsb777/call-back/pkg/jwt"
)

func InitHttpServer() (http.Handler, error) {
	wire.Build(
		// Конфиг
		config.NewConfig,
		// БД
		db.NewDatabasePG,
		db.ConnectDBProvider,
		// Репозитории
		repo.NewUserRepo,
		// Утилиты
		wire.Bind(new(hasher.PasswordHasher), new(*hasher.BCryptHasher)),
		hasher.NewBCryptHasher,
		wire.Bind(new(_jwt.Encoder), new(*_jwt.JWTEncoder)),
		_jwt.NewJWTEncoder,
		// Сервисы
		services.NewUserSignUpService,
		services.NewUserSignInService,
		// Роутеры
		routes.NewHTTPServer,
	)
	return nil, nil
}
