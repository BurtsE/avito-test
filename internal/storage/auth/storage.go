package auth

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	def "avito-test/internal/storage"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var _ def.UserStorage = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*repository, error) {
	DSN := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		cfg.UserDB.DB,
		cfg.UserDB.User,
		cfg.UserDB.Password,
		cfg.UserDB.Host,
		cfg.UserDB.Port,
		cfg.UserDB.Sslmode,
	)
	db, _ := sql.Open("postgres", DSN)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &repository{
		db: db,
	}, nil
}

// GetUser implements storage.UserStorage.
func (r *repository) User(ctx context.Context) {
	panic("unimplemented")
}

// RegisterUser implements storage.UserStorage.
func (r *repository) RegisterUser(ctx context.Context, user models.User) {
	query := `
		INSERT INTO users(user_id, email, pass, user_type)
		VALUES($1,$2,$3,$4)
	`
	r.db.Exec(query, user.Id, user.Password, user.Role)
}
