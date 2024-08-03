package house

import "avito-test/internal/models"

// CreateHouse implements service.HouseService.
func (s *service) CreateHouse(builder models.HouseBuilder) (*models.House, error) {
	house, err := s.houseStorage.CreateHouse(builder)
	if err != nil {
		return nil, err
	}
	return house, nil

}
