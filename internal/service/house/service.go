package house

import (
	"avito-test/internal/config"
	def "avito-test/internal/service"
	"avito-test/internal/storage"
	"time"

	"github.com/allegro/bigcache"
)

var _ def.HouseService = (*service)(nil)

type service struct {
	houseStorage storage.HouseStorage
	cache        *bigcache.BigCache
}

func NewService(houseStorage storage.HouseStorage, cfg *config.Config) *service {
	cache, _ := bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Second))
	return &service{
		houseStorage: houseStorage,
		cache:        cache,
	}
}
