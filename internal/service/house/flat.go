package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
)

func (s *service) UpdateFlatStatus(flatStatus models.FlatStatus) (*models.Flat, error) {
	status, err := converter.ModerationValueFromString(*flatStatus.Value)
	if err != nil {
		return nil, err
	}
	flat, err := s.houseStorage.UpdateFlatStatus(*flatStatus.Id, *flatStatus.Value)
	if err != nil {
		return nil, err
	}
	flat.Status = status
	return flat, nil
}

func (s *service) CreateFlat(flatBuilder models.FlatBuilder) (*models.Flat, error) {
	status, err := converter.StringFromModerationValue(models.OnModerate)
	if err != nil {
		return nil, err
	}
	flat, err := s.houseStorage.CreateFlat(flatBuilder, status)
	if err != nil {
		return nil, err
	}
	flat.Status = models.OnModerate
	return flat, nil
}
