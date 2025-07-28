package repo

import (
	"context"
	"database/sql"
	"time"

	"github.com/wsb777/call-back/internal/models"
)

type JWTRepo interface {
	IsTokenRevoked(ctx context.Context, tokenID string) (bool, error)
	RevokeToken(ctx context.Context, token *models.JwtToken) error
	CleanupExpiredTokens(ctx context.Context) error
	IsTokenRevokedBatch(ctx context.Context, tokenIDs []string) (map[string]bool, error)
}

type jwtRepo struct {
	db *sql.DB
}

func NewJWTRepo(db *sql.DB) JWTRepo {
	return &jwtRepo{db: db}
}

// Проверка отзыва токена
func (r *jwtRepo) IsTokenRevoked(ctx context.Context, tokenID string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM revoked_tokens WHERE id = $1)`
	var exists bool

	// Используем контекст для управления таймаутами
	err := r.db.QueryRowContext(ctx, query, tokenID).Scan(&exists)
	if err != nil {
		// Обрабатываем специальный случай отсутствия строк
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return exists, nil
}

// Отзыв токена
func (r *jwtRepo) RevokeToken(ctx context.Context, token *models.JwtToken) error {
	query := `INSERT INTO revoked_tokens (id, user_id, expires_at) 
              VALUES ($1, $2, $3)
              ON CONFLICT (id) DO NOTHING`

	_, err := r.db.ExecContext(
		ctx,
		query,
		token.ID,
		token.UserId,
		token.ExpiresAt.UTC(), // Всегда сохраняем время в UTC
	)
	return err
}

// Очистка устаревших токенов
func (r *jwtRepo) CleanupExpiredTokens(ctx context.Context) error {
	query := `DELETE FROM revoked_tokens WHERE expires_at < $1`
	_, err := r.db.ExecContext(ctx, query, time.Now().UTC())
	return err
}

// Дополнительный метод для пакетной проверки токенов (оптимизация)
func (r *jwtRepo) IsTokenRevokedBatch(ctx context.Context, tokenIDs []string) (map[string]bool, error) {
	if len(tokenIDs) == 0 {
		return map[string]bool{}, nil
	}

	query := `SELECT id FROM revoked_tokens WHERE id = ANY($1)`
	rows, err := r.db.QueryContext(ctx, query, tokenIDs)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	revokedMap := make(map[string]bool)
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		revokedMap[id] = true
	}

	return revokedMap, nil
}
