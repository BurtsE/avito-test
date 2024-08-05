package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
	"context"
)

func (r *repository) Flat(ctx context.Context, id uint64) (*models.Flat, error) {
	query := `
		SELECT price, room_number,house_id, moderation_status 
		FROM flats
		WHERE true
			AND id=$1
	`
	flat := models.Flat{Id: id}
	status := ""
	row := r.db.QueryRow(query, id)
	err := row.Scan(&flat.Price, &flat.RoomNumber, &flat.HouseId, &status)
	if err != nil {
		return nil, err
	}
	flat.Status, err = converter.ModerationValueFromString(status)
	if err != nil {
		return nil, err
	}
	return &flat, nil
}

func (r *repository) CreateFlat(ctx context.Context, builder models.FlatBuilder, status string) (*models.Flat, error) {
	flat := converter.FlatFromFlatBuilder(builder)
	query := `
		INSERT INTO flats(price, room_number, house_id, moderation_status)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	row := r.db.QueryRow(query, &flat.Price, &flat.RoomNumber, &flat.HouseId, &status)
	err := row.Scan(&flat.Id)
	if err != nil {
		return nil, err
	}
	return &flat, nil
}

func (r *repository) UpdateFlatStatus(ctx context.Context, id uint64, status string) (*models.Flat, error) {
	query := `
		UPDATE flats
		SET moderation_status = $2
		WHERE id = $1
	`
	flat := &models.Flat{Id: id}
	_, err := r.db.Exec(query, id, status)
	if err != nil {
		return nil, err
	}
	return flat, nil
}

func (r *repository) FlatsByHouseId(ctx context.Context, uuid uint64) ([]*models.Flat, error) {
	query := `
		SELECT id, price, room_number, moderation_status
		FROM flats
		WHERE true
			AND house_id=$1
	`
	rows, err := r.db.Query(query, &uuid)
	if err != nil {
		return nil, err
	}
	var status string
	flats := make([]*models.Flat, 0, 100)
	for rows.Next() {
		flat := models.Flat{}
		err = rows.Scan(&flat.Id, &flat.Price, &flat.RoomNumber, &status)
		if err != nil {
			return nil, err
		}
		flat.Status, err = converter.ModerationValueFromString(status)
		if err != nil {
			return nil, err
		}
		flats = append(flats, &flat)
	}
	return flats, nil
}
