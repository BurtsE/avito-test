package auth

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func (s *service) CheckAuthorization(ctx context.Context, data []byte) (models.User, error) {
	token, err := jwt.Parse(string(data), func(t *jwt.Token) (interface{}, error) {
		return s.jwtSecretKey, nil
	})
	if err != nil {
		return models.User{}, serviceErrors.AuthError{Err: err}
	}
	val, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return models.User{}, serviceErrors.AuthError{Err: err}
	}
	var role models.EnumRole
	id := ""
	if id, ok = val["id"].(string); !ok {
		return models.User{}, serviceErrors.AuthError{Err: err}
	}
	switch val["role"] {
	case "user":
		role = models.UserRole
	case "moderator":
		role = models.ModeratorRole
	default:
		return models.User{}, serviceErrors.AuthError{Err: errors.New("invalid token")}
	}
	return models.User{Id: &id, Role: role}, nil
}
