package house

import "avito-test/internal/models"

func (r *repository) GetHouseDesc(uuid uint64) (*models.House, error) {
	query := `
		SELECT adress, construction_date, developer, initialization_date, last_update_time
		FROM houses
		WHERE uuid=$1
	`
	house := &models.House{UUID: uuid}
	row := r.db.QueryRow(query, uuid)
	err := row.Scan(&house.Address, &house.ConstructionDate,
		&house.Developer, &house.InitializationDate, &house.LastUpdateTime)
	if err != nil {
		return nil, err
	}
	return house, nil
}
