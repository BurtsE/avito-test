package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
)

// CreateFlat implements house.HouseStorage.
func (r *repository) CreateFlat(builder models.FlatBuilder, status string) (*models.Flat, error) {
	flat := converter.FlatFromFlatBuilder(builder)

	query := `
		INSERT INTO flats(price, room_number, house_id, moderation_status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	row := r.db.QueryRow(query, &flat.Price, &flat.RoomNumber, &flat.HouseId, status)
	err := row.Scan(&flat.Id)
	if err != nil {
		return nil, err
	}
	return &flat, nil
}
