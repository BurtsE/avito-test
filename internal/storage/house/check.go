package house

import "context"

func (r *repository) HouseExists(ctx context.Context, id uint64) (bool, error) {
	query := `
		SELECT uuid
		FROM houses
		WHERE uuid = $1
	`
	err := r.db.QueryRow(query, id).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *repository) FlatExists(ctx context.Context, id uint64) (bool, error) {
	query := `
		SELECT id
		FROM flats
		WHERE id = $1
	`
	err := r.db.QueryRow(query, id).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}
