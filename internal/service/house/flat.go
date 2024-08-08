package house

import (
	"avito-test/internal/converter"
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"
)

var initialStatus = models.Created

func (s *service) UpdateFlatStatus(ctx context.Context, flatStatus models.FlatStatus) (*models.Flat, error) {

	flat, err := s.houseStorage.Flat(ctx, *flatStatus.Id)
	if err != nil {
		return flat, serviceErrors.ServerError{Err: err}
	}
	userId := ctx.Value(models.User{}).(models.User).Id
	if flat.Status == models.OnModerate && userId != &flat.ModeratorId {
		return flat, nil
	}

	err = s.houseStorage.UpdateFlatStatus(ctx, *flatStatus.Id, *flatStatus.Value)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}

	flat.Status, err = converter.ModerationValueFromString(*flatStatus.Value)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	return flat, nil
}

func (s *service) CreateFlat(ctx context.Context, flatBuilder models.FlatBuilder) (*models.Flat, error) {
	flat, err := s.houseStorage.CreateFlat(ctx, flatBuilder, initialStatus.String())
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	flat.Status = initialStatus
	return flat, nil
}
