package auth

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	def "avito-test/internal/service"
	serviceErrors "avito-test/internal/service_errors"
	"avito-test/internal/storage"
	"encoding/json"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

var _ def.AuthentificationService = (*service)(nil)

type service struct {
	storage      storage.UserStorage
	jwtSecretKey []byte
	method       jwt.SigningMethod
}

func NewService(userStorage storage.UserStorage, cfg *config.Config) *service {
	return &service{
		storage:      userStorage,
		jwtSecretKey: []byte(`some-private-key`),
		method:       jwt.GetSigningMethod("HS256"),
	}
}

func (s *service) DummyAuthorize(role models.EnumRole) (string, error) {
	var (
		roleStr string
		claims  jwt.MapClaims
	)
	switch role {
	case models.ModeratorRole:
		roleStr = "moderator"
	case models.UserRole:
		roleStr = "user"
	default:
		return "", errors.Wrap(serviceErrors.ServerError{}, "role does not exist")
	}
	jstr := fmt.Sprintf(`{"role":"%s"}`, roleStr)
	if err := json.Unmarshal([]byte(jstr), &claims); err != nil {
		return "", errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	log.Println(claims)
	token := jwt.NewWithClaims(s.method, claims)
	tokenString, err := token.SignedString(s.jwtSecretKey)
	if err != nil {
		return "", errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}

	return fmt.Sprintf(`{"token":"%s"}`, tokenString), nil
}

func (s *service) CheckAuthorization(data []byte) (models.EnumRole, error) {
	token, err := jwt.Parse(string(data), func(t *jwt.Token) (interface{}, error) {
		return s.jwtSecretKey, nil
	})
	if err != nil {
		return nil, errors.Wrap(serviceErrors.AuthError{}, err.Error())
	}
	val, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.Wrap(serviceErrors.AuthError{}, err.Error())
	}

	switch val["role"] {
	case "user":
		return models.UserRole, nil
	case "moderator":
		return models.ModeratorRole, nil
	default:
		return nil, errors.Wrap(serviceErrors.AuthError{}, err.Error())
	}
}
