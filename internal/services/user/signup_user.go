package services

import (
	"errors"
	"log"

	"github.com/wsb777/call-back/internal/db/repo"
	"github.com/wsb777/call-back/internal/dto"
	"github.com/wsb777/call-back/internal/models"
	"github.com/wsb777/call-back/pkg/hasher"
)

type UserSignUpService interface {
	CreateUser(req dto.UserSignUpDto) error
}

type userSignUpService struct {
	userRepo repo.UserRepo
	hasher   hasher.PasswordHasher
}

func NewUserSignUpService(
	userRepo repo.UserRepo,
	hasher hasher.PasswordHasher,
) UserSignUpService {
	return &userSignUpService{
		userRepo: userRepo,
		hasher:   hasher,
	}
}

func (s *userSignUpService) CreateUser(req dto.UserSignUpDto) error {
	log.Println("Запрос на существование пользователя")
	if existing, _ := s.userRepo.FindByLogin(req.Login); existing != nil {
		log.Println("Пользователь существует")
		return errors.New("Пользователь существует")
	}
	log.Println("Такого пользователя не существует")

	hashedPassword, err := s.hasher.HashPassword(req.Password)
	if err != nil {
		return err
	}

	log.Println("Пароль успешно захеширован")

	user := &models.User{
		Login:    req.Login,
		Password: hashedPassword,
	}

	log.Println("Данные засэчены в модель")

	err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil
	}
	log.Println("Модель зарегистрирована")

	return nil
}
