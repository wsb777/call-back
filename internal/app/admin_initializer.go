package app

import (
	"log"

	"github.com/wsb777/call-back/internal/config"
	"github.com/wsb777/call-back/internal/dto"
	services "github.com/wsb777/call-back/internal/services/user"
)

type AdminInitializer struct {
	userService services.UserSignUpService
	cfg         *config.Config
}

func NewAdminInitializer(us services.UserSignUpService, cfg *config.Config) *AdminInitializer {
	return &AdminInitializer{userService: us, cfg: cfg}
}

func (ai *AdminInitializer) InitAdmin() {
	var adminInfo dto.UserSignUpDto
	adminInfo.Login = "admin"
	adminInfo.Password = ai.cfg.AdminPassword
	err := ai.userService.CreateUser(adminInfo)

	if err != nil {
		log.Println("У сервиса уже есть созданный администратор")
	} else {
		log.Println("✅ Админ создан")
	}
}
