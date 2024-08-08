package house

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func (s *service) HouseDesc(ctx context.Context, uuid uint64) (*models.House, error) {
	house, err := s.houseStorage.HouseDesc(ctx, uuid)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	return house, nil
}

func (s *service) CreateHouse(ctx context.Context, builder models.HouseBuilder) (*models.House, error) {
	house, err := s.houseStorage.CreateHouse(ctx, builder)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	return house, nil
}

func (s *service) HouseFlats(ctx context.Context, uuid uint64) ([]*models.Flat, error) {
	idStr := strconv.FormatUint(uuid, 10)
	role := ctx.Value(models.User{}).(models.User).Role
	roleBytes, err := json.Marshal(role)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	casheKey := fmt.Sprintf("%s%s", idStr, roleBytes)
	data, err := s.cache.Get(casheKey)
	flats := make([]*models.Flat, 0)
	if err == nil {
		err = json.Unmarshal(data, &flats)
		if err == nil {
			return flats, nil
		}
	}
	flats, err = s.houseStorage.FlatsByHouseId(ctx, uuid)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	log.Println(flats)
	data, err = json.Marshal(flats)
	if err != nil {
		return nil, serviceErrors.ServerError{Err: err}
	}
	s.cache.Set(casheKey, data)
	return flats, nil
}
