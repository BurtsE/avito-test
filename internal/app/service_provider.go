package app

import (
	"avito-test/internal/config"
	router "avito-test/internal/router"
	"avito-test/internal/service"
	houseStorage "avito-test/internal/storage"
	houseStorageInstance "avito-test/internal/storage/house"
	houseService "avito-test/internal/service/house"
	"log"

	"github.com/sirupsen/logrus"
)

type serviceProvider struct {
	cfg          *config.Config
	houseStorage houseStorage.HouseStorage
	houseService service.HouseService
	router       *router.Router
}

func NewSericeProvider() *serviceProvider {
	s := &serviceProvider{}
	s.Router()
	return s
}

func (s *serviceProvider) Config() *config.Config {
	if s.cfg == nil {
		cfg, err := config.InitConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.cfg = cfg
	}
	return s.cfg
}

func (s *serviceProvider) MessageRepository() houseStorage.HouseStorage {
	if s.houseStorage == nil {
		storage, err := houseStorageInstance.NewRepository(s.cfg)
		if err != nil {
			log.Fatalf("could not init storage: %s", err.Error())
		}
		s.houseStorage = storage
	}
	return s.houseStorage
}

func (s *serviceProvider) HouseService() service.HouseService {
	if s.houseService == nil {
		s.houseService = houseService.NewService(s.MessageRepository(), s.Config())
	}
	return s.houseService
}

func (s *serviceProvider) Router() *router.Router {
	if s.router == nil {
		s.router = router.NewRouter(logrus.New(), s.Config(), s.HouseService())
	}
	return s.router
}
