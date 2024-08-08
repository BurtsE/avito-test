package house

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"
)

func (s *service) HouseDesc(ctx context.Context, uuid uint64) (*models.House, error) {
	house, err := s.houseStorage.HouseDesc(ctx, uuid)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	return house, nil
}

func (s *service) CreateHouse(ctx context.Context, builder models.HouseBuilder) (*models.House, error) {
	house, err := s.houseStorage.CreateHouse(ctx, builder)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	return house, nil
}

func (s *service) HouseFlats(ctx context.Context, uuid uint64) ([]*models.Flat, error) {
	flats, err := s.houseStorage.FlatsByHouseId(ctx, uuid)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	return flats, nil
}
