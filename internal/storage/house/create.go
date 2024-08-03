package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
	"log"

	"github.com/google/uuid"
)

func (r *repository) CreateHouse(builder models.HouseBuilder) (*models.House, error) {
	house := converter.HouseFromBuilder(builder)
	log.Println(len(house.UUID))
	query := `
		INSERT INTO houses (uuid, adress, construction_date, developer, initialization_date, last_update_time)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.Exec(query, uuid.New(), house.Address, house.ConstructionDate,
		house.Developer, house.InitializationDate, house.LastUpdateTime)
	if err != nil {
		return nil, err
	}
	return &house, nil
}
