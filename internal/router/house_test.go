package router

import (
	"avito-test/internal/config"
	"avito-test/internal/models"
	"avito-test/internal/test/mocks"
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func Test_getHouseData(t *testing.T) {
	tests := []struct {
		role   models.EnumRole
		flats  []*models.Flat
		result string
	}{
		{
			role: models.UserRole,
			flats: []*models.Flat{
				{Status: models.Approved},
			},
			result: `[{"id":0,"house_id":0,"price":0,"rooms":0,"status":"approved"}]`,
		},
		{
			role: models.ModeratorRole,
			flats: []*models.Flat{
				{Status: models.Approved},
				{Status: models.OnModerate},
				{Status: models.Created},
			},
			result: `[{"id":0,"house_id":0,"price":0,"rooms":0,"status":"approved"},{"id":0,"house_id":0,"price":0,"rooms":0,"status":"on moderate"},{"id":0,"house_id":0,"price":0,"rooms":0,"status":"created"}]`,
		},
	}
	houseService := mocks.NewHouseService(t)
	authService := mocks.NewAuthentificationService(t)
	validationService := mocks.NewValidationService(t)
	id := uint64(2)
	h := houseImpl{
		r: NewRouter(logrus.New(), &config.Config{}, houseService, validationService, authService),
	}

	for _, test := range tests {
		apiCtx := new(fasthttp.RequestCtx)
		apiCtx.SetUserValue("id", "2")
		apiCtx.SetUserValue("user", models.User{Role: test.role})
		ctx := context.WithValue(context.Background(), models.User{}, models.User{Role: test.role})

		validationService.
			On("ValidateHouse", ctx, id).
			Return(nil)
		houseService.
			On("HouseFlats", ctx, id).
			Return(test.flats, nil)
		h.getHouseData(apiCtx)
		responce := apiCtx.Response.Body()
		validationService.AssertCalled(t, "ValidateHouse", ctx, id)
		houseService.AssertCalled(t, "HouseFlats", ctx, id)
		if string(responce) != test.result {
			t.Log(responce)
			t.Fatalf("wrong result. Expected: %s\nGot: %s", test.result, responce)
		}
	}
}
