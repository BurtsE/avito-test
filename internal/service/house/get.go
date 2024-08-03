package house

import "avito-test/internal/models"

// GetHouseDesc implements service.HouseService.
func (s *service) GetHouseDesc(uuid uint64) (*models.House, error) {
	return s.houseStorage.GetHouseDesc(uuid)
}
