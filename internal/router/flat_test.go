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

func Test_createFlat(t *testing.T) {
	houseService := mocks.NewHouseService(t)
	authService := mocks.NewAuthentificationService(t)
	validationService := mocks.NewValidationService(t)
	flatData := []byte(
		`{
  			"house_id": 12345,
 			"price": 10000,
  			"rooms": 4
		}`)
	houseId := uint64(12345)
	price := uint64(10000)
	rooms := byte(4)
	flatBuilder := models.FlatBuilder{
		HouseId: &houseId,
		Price:   &price,
		Rooms:   &rooms,
	}
	flat := models.Flat{
		HouseId:    houseId,
		Price:      price,
		RoomNumber: rooms,
		Status:     models.OnModerate,
	}
	apiCtx := new(fasthttp.RequestCtx)
	apiCtx.Request.AppendBody(flatData)
	apiCtx.SetUserValue("user", models.User{Role: models.UserRole})
	ctx := context.WithValue(context.Background(), models.User{}, apiCtx.UserValue("user"))

	expectedResponce :=
		`{"id":0,"unit_number":0,"house_id":12345,"price":10000,"rooms":4,"status":{"ModerationStatus":null},"moderator_id":""}`
	validationService.
		On("ValidateFlatBuilderData", ctx, flatData).
		Return(flatBuilder, nil)

	houseService.
		On("CreateFlat", ctx, flatBuilder).
		Return(flat, nil)

	router := NewRouter(logrus.New(), &config.Config{}, houseService, validationService, authService)
	f := &flatImpl{
		r: router,
	}
	f.createFlat(apiCtx)
	responce := apiCtx.Response.Body()
	t.Log(string(responce))
	validationService.AssertCalled(t, "ValidateFlatBuilderData", ctx, flatData)
	houseService.AssertCalled(t, "CreateFlat", ctx, flatBuilder)
	if string(responce) != expectedResponce {
		t.Fatalf("wrong result. Expected: %s\nGot: %s", expectedResponce, responce)
	}
}
