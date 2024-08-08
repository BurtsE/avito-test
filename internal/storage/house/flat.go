package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
	"context"
	"time"
)

func (r *repository) Flat(ctx context.Context, id uint64) (*models.Flat, error) {
	query := `
		SELECT unit_number, price, room_number,house_id, moderation_status 
		FROM flats
		WHERE true
			AND id=$1
	`
	flat := models.Flat{Id: id}
	status := ""
	row := r.db.QueryRow(query, id)
	err := row.Scan(&flat.UnitNumber, &flat.Price, &flat.RoomNumber, &flat.HouseId, &status)
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
		INSERT INTO flats(unit_number, price, room_number, house_id, moderation_status)
		VALUES (
			(SELECT flats_number + 1
			FROM HOUSES h 
			WHERE h.uuid = $3
			), $1, $2, $3, $4)
		RETURNING id, unit_number
	`
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	row := tx.QueryRow(query, &flat.Price, &flat.RoomNumber, &flat.HouseId, &status)
	err = row.Scan(&flat.Id, &flat.UnitNumber)
	if err != nil {
		return nil, err
	}
	query = `
		UPDATE houses
		SET last_update_time = $2, flats_number = flats_number + 1
		WHERE uuid = $1
	`
	_, err = tx.Exec(query, builder.HouseId, time.Now())
	if err != nil {
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return &flat, nil
}

func (r *repository) UpdateFlatStatus(ctx context.Context, id uint64, status string) error {
	query := `
		UPDATE flats
		SET moderation_status = $2
		AND moderator_id = $3
		WHERE id = $1
	`
	userId := ctx.Value(models.User{}).(models.User).Id
	_, err := r.db.Exec(query, id, status, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FlatsByHouseId(ctx context.Context, uuid uint64) ([]*models.Flat, error) {
	query := `
		SELECT id, unit_number, price, room_number, moderation_status
		FROM flats
		WHERE true
			AND house_id=$1
	`
	if ctx.Value(models.Role{}) == models.UserRole {
		query += "AND moderation_status == 'approved'"
	}
	rows, err := r.db.Query(query, &uuid)
	if err != nil {
		return nil, err
	}
	var status string
	flats := make([]*models.Flat, 0, 100)
	for rows.Next() {
		flat := models.Flat{HouseId: uuid}
		err = rows.Scan(&flat.Id, &flat.UnitNumber, &flat.Price, &flat.RoomNumber, &status)
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
