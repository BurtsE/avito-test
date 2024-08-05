package house

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"

	"github.com/pkg/errors"
)

func (s *service) HouseDesc(ctx context.Context, uuid uint64) (*models.House, error) {
	house, err := s.houseStorage.HouseDesc(ctx, uuid)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	return house, nil
}

func (s *service) CreateHouse(ctx context.Context, builder models.HouseBuilder) (*models.House, error) {
	house, err := s.houseStorage.CreateHouse(ctx, builder)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	return house, nil
}

func (s *service) HouseFlats(ctx context.Context, uuid uint64) ([]*models.Flat, error) {
	flats, err := s.houseStorage.FlatsByHouseId(ctx, uuid)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	role := ctx.Value(models.Role{})
	if role == models.UserRole {
		return s.filterFlats(flats), nil
	}
	return flats, nil
}

func (s *service) filterFlats(flats []*models.Flat) []*models.Flat {
	filterdFlats := make([]*models.Flat, 0, 100)
	for _, flat := range flats {
		if flat.Status == models.OnModerate {
			continue
		}
		filterdFlats = append(filterdFlats, flat)
	}
	return filterdFlats
}
