package converter

import (
	"avito-test/internal/models"
	"encoding/json"
	"errors"
	"time"
)

var invalidJsonError error = errors.New("invalid json")

func HouseFromHouseBuilder(builder models.HouseBuilder) models.House {
	house := models.House{
		Address:   *builder.Address,
		Developer: *builder.Developer,
	}
	initIime := time.Now()
	house.InitializationDate = initIime
	house.LastUpdateTime = initIime

	constructionDate := time.Date(*builder.ConstructionDate, 0, 0, 0, 0, 0, 0, time.Local)
	house.ConstructionDate = constructionDate
	return house
}

func HouseBuilderFromRawData(data []byte) (models.HouseBuilder, error) {
	builder := models.HouseBuilder{}
	err := json.Unmarshal(data, &builder)
	if err != nil {
		return builder, err
	}
	if builder.ConstructionDate == nil || builder.Address == nil {
		return builder, invalidJsonError
	}
	if builder.Developer == nil {
		builder.Developer = new(string)
	}
	return builder, nil
}

func FlatBuilderFromRawData(data []byte) (models.FlatBuilder, error) {
	builder := models.FlatBuilder{}
	err := json.Unmarshal(data, &builder)
	if err != nil {
		return builder, err
	}
	if builder.HouseId == nil || builder.Price == nil || builder.Rooms == nil {
		return builder, invalidJsonError
	}
	return builder, nil
}

func FlatFromFlatBuilder(builder models.FlatBuilder) models.Flat {
	flat := models.Flat{
		HouseId:    *builder.HouseId,
		Price:      *builder.Price,
		RoomNumber: *builder.Rooms,
	}
	return flat
}

func FlatStatusFromRawData(data []byte) (models.FlatStatus, error) {
	status := models.FlatStatus{}
	err := json.Unmarshal(data, &status)
	if err != nil {
		return status, err
	}
	if status.Id == nil || status.Value == nil {
		return status, invalidJsonError
	}
	return status, nil
}

func ModerationValueFromString(s string) (models.ModerationStatus, error) {
	switch s {
	case "approved":
		return models.Approved, nil
	case "created":
		return models.Created, nil
	case "on moderate":
		return models.OnModerate, nil
	case "declined":
		return models.Declined, nil
	default:
		return nil, errors.New("unknown status")
	}
}

func StringFromModerationValue(status models.ModerationStatus) (string, error) {
	switch status {
	case models.Approved:
		return "approved", nil
	case models.Created:
		return "created", nil
	case models.OnModerate:
		return "on moderate", nil
	case models.Declined:
		return "declined", nil
	default:
		return "", errors.New("unknown status")
	}
}
