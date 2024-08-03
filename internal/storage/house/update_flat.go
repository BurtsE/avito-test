package house

import (
	"avito-test/internal/models"
)

func (r *repository) UpdateFlatStatus(id uint64, status string) (*models.Flat, error) {
	query := `
		SELECT price, room_number, house_id
		FROM flats
		WHERE id = $1
	`
	flat := &models.Flat{Id: id}
	err := r.db.QueryRow(query, id).Scan(&flat.Price, &flat.RoomNumber, &flat.HouseId)
	if err != nil {
		return nil, err
	}

	query = `
		UPDATE flats
		SET status = $2
		WHERE id = $1
	`
	_, err = r.db.Exec(query, id, status)
	if err != nil {
		return nil, err
	}
	return flat, nil
}
