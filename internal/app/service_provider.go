package app

import (
	"avito-test/internal/config"
	router "avito-test/internal/router"
	"avito-test/internal/service"
	userService "avito-test/internal/service/auth"
	houseService "avito-test/internal/service/house"
	validationService "avito-test/internal/service/validation"
	storages "avito-test/internal/storage"
	userStorageInstance "avito-test/internal/storage/auth"
	houseStorageInstance "avito-test/internal/storage/house"
	"log"

	"github.com/sirupsen/logrus"
)

// TODO auth service, repo init

type serviceProvider struct {
	cfg               *config.Config
	houseStorage      storages.HouseStorage
	validationStorage storages.ValidationStorage
	houseService      service.HouseService
	validationService service.ValidationService
	userService       service.AuthentificationService
	userStorage       storages.UserStorage
	router            *router.Router
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

func (s *serviceProvider) HouseStorage() storages.HouseStorage {
	if s.houseStorage == nil {
		storage, err := houseStorageInstance.NewRepository(s.Config())
		if err != nil {
			log.Fatalf("could not init house storage: %s", err.Error())
		}
		s.houseStorage = storage
		s.validationStorage = storage
	}
	return s.houseStorage
}

func (s *serviceProvider) ValidationStorage() storages.ValidationStorage {
	if s.validationStorage == nil {
		storage, err := houseStorageInstance.NewRepository(s.Config())
		if err != nil {
			log.Fatalf("could not init house storage: %s", err.Error())
		}
		s.houseStorage = storage
		s.validationStorage = storage
	}
	return s.validationStorage
}

func (s *serviceProvider) HouseService() service.HouseService {
	if s.houseService == nil {
		s.houseService = houseService.NewService(s.HouseStorage(), s.Config())
	}
	return s.houseService
}

func (s *serviceProvider) ValidationService() service.ValidationService {
	if s.validationService == nil {
		s.validationService = validationService.NewService(s.ValidationStorage(), s.Config())
	}
	return s.validationService
}

func (s *serviceProvider) UserStorage() storages.UserStorage {
	if s.userStorage == nil {
		storage, err := userStorageInstance.NewRepository(s.Config())
		if err != nil {
			log.Fatalf("could not init user storage: %s", err.Error())
		}
		s.userStorage = storage
	}
	return s.userStorage
}

func (s *serviceProvider) UserService() service.AuthentificationService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserStorage(), s.Config())
	}
	return s.userService
}

func (s *serviceProvider) Router() *router.Router {
	if s.router == nil {
		s.router = router.NewRouter(logrus.New(), s.Config(), s.HouseService(), s.ValidationService(), s.UserService()) // nil
	}
	return s.router
}
