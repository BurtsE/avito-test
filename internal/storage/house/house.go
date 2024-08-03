package house

import (
	"avito-test/internal/config"
	def "avito-test/internal/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var _ def.HouseStorage = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*repository, error) {
	DSN := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		cfg.DB,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Postgres.Port,
		cfg.Sslmode,
	)
	db, _ := sql.Open("postgres", DSN)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &repository{
		db: db,
	}, nil
}
