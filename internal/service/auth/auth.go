package auth

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	def "avito-test/internal/service"
	serviceErrors "avito-test/internal/service_errors"
	"avito-test/internal/storage"
	"encoding/json"
	"fmt"

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
		jwtSecretKey: []byte("javainuse-secret-key"),
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

	token := jwt.NewWithClaims(s.method, claims)
	signedToken, err := token.SignedString(s.jwtSecretKey)
	if err != nil {
		return "", errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	return fmt.Sprintf(`{"token":"%s"}`, signedToken), nil
}

func (s *service) CheckAuthorization(token []byte) error {
	return nil
}
