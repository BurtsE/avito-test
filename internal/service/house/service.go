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
	cacheCfg := bigcache.DefaultConfig(10 * time.Second)
	cacheCfg.CleanWindow = 1 * time.Second
	cache, _ := bigcache.NewBigCache(cacheCfg)

	return &service{
		houseStorage: houseStorage,
		cache:        cache,
	}
}
