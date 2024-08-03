package house

import (
	"avito-test/internal/config"
	def "avito-test/internal/service"
	storage "avito-test/internal/storage"
)

var _ def.HouseService = (*service)(nil)

type service struct {
	houseStorage storage.HouseStorage
}

func NewService(houseStorage storage.HouseStorage, cfg *config.Config) *service {
	return &service{
		houseStorage: houseStorage,
	}
}
