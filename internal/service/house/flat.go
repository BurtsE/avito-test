package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"

	"github.com/pkg/errors"
)

func (s *service) UpdateFlatStatus(ctx context.Context, flatStatus models.FlatStatus) (*models.Flat, error) {
	status, err := converter.ModerationValueFromString(*flatStatus.Value)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}

	flat, err := s.houseStorage.Flat(ctx, *flatStatus.Id)
	if err != nil {
		return flat, err
	}
	if flat.Status == models.OnModerate {
		return flat, nil
	}

	flat, err = s.houseStorage.UpdateFlatStatus(ctx, *flatStatus.Id, *flatStatus.Value)
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	flat.Status = status
	return flat, nil
}

func (s *service) CreateFlat(ctx context.Context, flatBuilder models.FlatBuilder) (*models.Flat, error) {
	// status, err := converter.StringFromModerationValue(models.Created)
	// if err != nil {
	// 	return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	// }
	flat, err := s.houseStorage.CreateFlat(ctx, flatBuilder, models.Created.String())
	if err != nil {
		return nil, errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	flat.Status = models.OnModerate
	return flat, nil
}
