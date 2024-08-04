package house

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"

	"github.com/pkg/errors"
)

func (s *service) HouseDesc(uuid uint64) (*models.House, error) {
	house, err := s.houseStorage.HouseDesc(uuid)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	return house, nil
}

func (s *service) CreateHouse(builder models.HouseBuilder) (*models.House, error) {
	house, err := s.houseStorage.CreateHouse(builder)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	return house, nil
}

func (s *service) HouseFlats(uuid uint64) ([]*models.Flat, error) {
	flats, err := s.houseStorage.FlatsByHouseId(uuid)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	return flats, nil
}
