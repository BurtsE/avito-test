package house

import "avito-test/internal/models"

func (s *service) GetHouseDesc(uuid uint64) (*models.House, error) {
	return s.houseStorage.GetHouseDesc(uuid)
}

func (s *service) CreateHouse(builder models.HouseBuilder) (*models.House, error) {
	return s.houseStorage.CreateHouse(builder)
}
