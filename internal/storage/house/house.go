package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
	"context"
)

func (r *repository) CreateHouse(ctx context.Context, builder models.HouseBuilder) (models.House, error) {
	house := converter.HouseFromHouseBuilder(builder)
	query := `
		INSERT INTO houses (street, construction_date, developer, initialization_date, last_update_time)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING uuid
	`
	row := r.db.QueryRow(query, house.Address, house.ConstructionDate,
		house.Developer, house.InitializationDate, house.LastUpdateTime)
	err := row.Scan(&house.UUID)
	if err != nil {
		return house, err
	}
	return house, nil
}

func (r *repository) HouseDesc(ctx context.Context, uuid uint64) (models.House, error) {
	query := `
		SELECT street, construction_date, developer, initialization_date, last_update_time
		FROM houses
		WHERE uuid=$1
	`
	house := models.House{UUID: uuid}
	row := r.db.QueryRow(query, uuid)
	err := row.Scan(&house.Address, &house.ConstructionDate,
		&house.Developer, &house.InitializationDate, &house.LastUpdateTime)
	if err != nil {
		return house, err
	}
	return house, nil
}
