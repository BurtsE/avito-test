package house

import "avito-test/internal/models"

func (r *repository) GetHouseDesc(uuid uint64) (*models.House, error) {
	query := `
		SELECT adress, construction_date, developer, initialization_date, last_update_time
		FROM houses
		WHERE uuid=$1
	`
	house := &models.House{}
	row := r.db.QueryRow(query, uuid)
	if row.Err() != nil {
		return nil, row.Err()
	}
	row.Scan(&house.UUID, &house.Address, &house.ConstructionDate,
		&house.Developer, &house.InitializationDate, &house.LastUpdateTime)
	return house, nil
}
