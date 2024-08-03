package service

import "avito-test/internal/models"

type HouseService interface {
	CreateHouse(models.HouseBuilder) (*models.House, error)
	GetHouseDesc() (models.House, error)
}
