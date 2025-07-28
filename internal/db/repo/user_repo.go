package repo

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/wsb777/call-back/internal/models"
)

type UserRepo interface {
	CreateUser(user *models.User) error
	FindByLogin(login string) (*models.User, error)
	FindById(id string) (*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	query := `
	INSERT INTO users (login, password, created_at, updated_at, system_role_id)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
`
	now := time.Now().UTC()
	err := r.db.QueryRow(
		query,
		user.Login,
		user.Password,
		now,
		now,
		1,
	).Scan(&user.ID)

	if err != nil {
		// Логируем ошибку для отладки
		log.Printf("Ошибка при создании пользователя: %v", err)
		return fmt.Errorf("ошибка при создании пользователя: %w", err)
	}

	log.Println("Пользователь создан")

	return err
}

func (r *userRepo) FindByLogin(login string) (*models.User, error) {
	query := "SELECT id, login, password, system_role_id FROM users WHERE login = $1"
	row := r.db.QueryRow(query, login)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.SystemRole)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}

func (r *userRepo) FindById(id string) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	row := r.db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.ID, &user.Login, &user.Password, &user.SystemRole)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	return user, nil
}
