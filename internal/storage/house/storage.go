package house

import (
	"avito-test/internal/config"
	def "avito-test/internal/storage"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var _ def.HouseStorage = (*repository)(nil)
var _ def.ValidationStorage = (*repository)(nil)

type repository struct {
	db *sql.DB
}

func NewRepository(cfg *config.Config) (*repository, error) {
	DSN := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s port=%s sslmode=%s",
		cfg.HouseDB.DB,
		cfg.HouseDB.User,
		cfg.HouseDB.Password,
		cfg.HouseDB.Host,
		cfg.HouseDB.Port,
		cfg.HouseDB.Sslmode,
	)
	db, _ := sql.Open("postgres", DSN)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &repository{
		db: db,
	}, nil
}
