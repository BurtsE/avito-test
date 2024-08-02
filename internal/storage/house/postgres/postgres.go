package postgres

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	def "avito-test/internal/storage/house"
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

func (r *repository) GetHouseDesc(uuid uint64) (models.House, error) {
	query := `
		SELECT adress, construction_date, developer, initialization_date, last_update_time
		FROM houses
		WHERE uuid=$1
	`
	house := models.House{}
	row := r.db.QueryRow(query, uuid)
	if row.Err() != nil {
		return models.House{}, row.Err()
	}
	row.Scan(&house.UUID, &house.Adress, &house.ConstructionDate,
		&house.Developer, &house.InitializationDate, &house.LastUpdateTime)
	return house, nil
}

func (r *repository) CreateHouse(house models.House) (uint64, error) {
	var id uint64
	query := `
		INSERT INTO houses (uuid, adress, construction_date, developer, initialization_date, last_update_time)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	row := r.db.QueryRow(query, house.UUID, house.Adress, house.ConstructionDate,
		house.Developer, house.InitializationDate, house.LastUpdateTime)
	if row.Err() != nil {
		return 0, row.Err()
	}
	row.Scan(&id)
	return id, nil
}
