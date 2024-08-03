package converter

import (
	"avito-test/internal/models"
	"encoding/json"
	"errors"
	"time"
)

func HouseFromBuilder(builder models.HouseBuilder) models.House {
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

func BuilderFromRawData(data []byte) (models.HouseBuilder, error) {
	builder := models.HouseBuilder{}
	err := json.Unmarshal(data, &builder)
	if err != nil {
		return builder, err
	}
	if builder.ConstructionDate == nil || builder.Address == nil {
		return builder, errors.New("invalid json")
	}
	if builder.Developer == nil {
		builder.Developer = new(string)
	}
	return builder, nil
}
