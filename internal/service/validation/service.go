package validation

import (
	"avito-test/internal/config"
	"avito-test/internal/converter"
	"avito-test/internal/models"
	def "avito-test/internal/service"
	serviceErrors "avito-test/internal/service_errors"
	"avito-test/internal/storage"

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

func (s *service) ValidateFlatStatusData(data []byte) (models.FlatStatus, error) {
	status, err := converter.FlatStatusFromRawData(data)
	if err != nil {
		return models.FlatStatus{}, errors.Wrap(serviceErrors.ValidationError{}, err.Error())
	}

	err = s.ValidateFlat(*status.Id)
	if err != nil {
		return models.FlatStatus{}, err
	}
	return status, nil
}

func (s *service) ValidateHouseData(data []byte) (models.HouseBuilder, error) {
	builder, err := converter.HouseBuilderFromRawData(data)
	if err != nil {
		return models.HouseBuilder{}, errors.Wrap(serviceErrors.ValidationError{}, err.Error())
	}
	return builder, nil
}

func (s *service) ValidateFlatBuilderData(data []byte) (models.FlatBuilder, error) {
	builder, err := converter.FlatBuilderFromRawData(data)
	if err != nil {
		return models.FlatBuilder{}, errors.Wrap(serviceErrors.ValidationError{}, err.Error())
	}
	err = s.ValidateHouse(*builder.HouseId)
	if err != nil {
		return models.FlatBuilder{}, err
	}
	return builder, nil
}

func (s *service) ValidateHouse(uuid uint64) error {
	exists, err := s.validationStorage.HouseExists(uuid)
	if err != nil {
		return errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	if !exists {
		return errors.Wrap(serviceErrors.ValidationError{}, "house does not exist")
	}
	return nil
}

func (s *service) ValidateFlat(uuid uint64) error {
	exists, err := s.validationStorage.FlatExists(uuid)
	if err != nil {
		return errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	if !exists {
		return errors.Wrap(serviceErrors.ValidationError{}, "flat does not exist")
	}
	return nil
}
