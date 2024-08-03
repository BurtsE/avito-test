package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
)

func (r *repository) CreateHouse(builder models.HouseBuilder) (*models.House, error) {
	house := converter.HouseFromBuilder(builder)
	query := `
		INSERT INTO houses (adress, construction_date, developer, initialization_date, last_update_time)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING uuid
	`
	row := r.db.QueryRow(query, house.Address, house.ConstructionDate,
		house.Developer, house.InitializationDate, house.LastUpdateTime)
	err := row.Scan(&house.UUID)
	if err != nil {
		return nil, err
	}
	return &house, nil
}
