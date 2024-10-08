package auth

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	def "avito-test/internal/service"
	serviceErrors "avito-test/internal/service_errors"
	"avito-test/internal/storage"
	"context"
	"encoding/json"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// TODO change private key

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

func (s *service) DummyAuthorize(ctx context.Context, role models.EnumRole) (string, error) {
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
		return "", serviceErrors.ServerError{Err: errors.New("role does not exist")}
	}
	jstr := fmt.Sprintf(`{"role":"%s","id":"%s"}`, roleStr, uuid.New().String())
	if err := json.Unmarshal([]byte(jstr), &claims); err != nil {
		return "", errors.Wrap(serviceErrors.ServerError{}, err.Error())
	}
	token := jwt.NewWithClaims(s.method, claims)
	tokenString, err := token.SignedString(s.jwtSecretKey)
	if err != nil {
		return "", serviceErrors.ServerError{Err: err}
	}

	return fmt.Sprintf(`{"token":"%s"}`, tokenString), nil
}

func (s *service) RegisterUser(ctx context.Context, user models.User) (string, error) {
	panic("unimplemented")
}
