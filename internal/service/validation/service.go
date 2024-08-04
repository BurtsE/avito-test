package validation

import (
	"avito-test/internal/config"
	"avito-test/internal/converter"
	"avito-test/internal/models"
	def "avito-test/internal/service"
	serviceErrors "avito-test/internal/service_errors"
	"avito-test/internal/storage"
	"context"

	"github.com/pkg/errors"
)

var _ def.ValidationService = (*service)(nil)

type service struct {
	validationStorage storage.ValidationStorage
}

func NewService(validationStorage storage.ValidationStorage, cfg *config.Config) *service {
	return &service{
		validationStorage: validationStorage,
	}
}

func (s *service) ValidateFlatStatusData(ctx context.Context, data []byte) (models.FlatStatus, error) {
	status, err := converter.FlatStatusFromRawData(data)
	if err != nil {
		return models.FlatStatus{}, errors.Wrap(serviceErrors.ValidationError{}, err.Error())
	}

	err = s.ValidateFlat(ctx, *status.Id)
	if err != nil {
		return models.FlatStatus{}, err
	}
	return status, nil
}

func (s *service) ValidateHouseData(ctx context.Context, data []byte) (models.HouseBuilder, error) {
	builder, err := converter.HouseBuilderFromRawData(data)
	if err != nil {
		return models.HouseBuilder{}, errors.Wrap(serviceErrors.ValidationError{}, err.Error())
	}
	return builder, nil
}

func (s *service) ValidateFlatBuilderData(ctx context.Context, data []byte) (models.FlatBuilder, error) {
	builder, err := converter.FlatBuilderFromRawData(data)
	if err != nil {
		return models.FlatBuilder{}, errors.Wrap(serviceErrors.ValidationError{}, err.Error())
	}
	err = s.ValidateHouse(ctx ,*builder.HouseId)
	if err != nil {
		return models.FlatBuilder{}, err
	}
	return builder, nil
}

func (s *service) ValidateHouse(ctx context.Context, uuid uint64) error {
	exists, err := s.validationStorage.HouseExists(uuid)
	if err != nil {
		return errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	if !exists {
		return errors.Wrap(serviceErrors.ValidationError{}, "house does not exist")
	}
	return nil
}

func (s *service) ValidateFlat(ctx context.Context, uuid uint64) error {
	exists, err := s.validationStorage.FlatExists(uuid)
	if err != nil {
		return errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	if !exists {
		return errors.Wrap(serviceErrors.ValidationError{}, "flat does not exist")
	}
	return nil
}

func (s *service) ValidateDummyUserData(ctx context.Context, data []byte) (models.EnumRole, error) {
	switch string(data) {
	case "user":
		return models.UserRole, nil
	case "moderator":
		return models.ModeratorRole, nil
	default:
		return nil, errors.Wrap(serviceErrors.ValidationError{}, "role does not exist")
	}
}
