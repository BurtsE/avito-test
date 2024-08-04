package router

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"
	"errors"
	"strconv"

	"github.com/valyala/fasthttp"
)

type houseImpl struct {
	r *Router
}

func registerHouseApi(r *Router) {
	houseImpl := houseImpl{r}
	r.router.POST("/house/create", r.ModeratorAccess(houseImpl.createHouse))
	r.router.GET("/house/{id:[0-9]*}", r.UserAccess(houseImpl.getHouseData))
	r.router.POST("/house/{id}/subscribe", r.UserAccess(houseImpl.subscribe))
}

func (h *houseImpl) createHouse(apiCtx *fasthttp.RequestCtx) {
	serviceCtx := context.WithValue(context.Background(), models.Role{}, apiCtx.Value("role"))
	data := apiCtx.Request.Body()
	builder, err := h.r.validationService.ValidateHouseData(serviceCtx, data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		h.r.logger.Println(err)
		invalidDataResponce(apiCtx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}

	house, err := h.r.houseService.CreateHouse(serviceCtx, builder)

	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}

	h.r.sendResponce(apiCtx, house)
}
func (h *houseImpl) getHouseData(apiCtx *fasthttp.RequestCtx) {
	serviceCtx := context.WithValue(context.Background(), models.Role{}, apiCtx.Value("role"))
	idStr := apiCtx.UserValue("id").(string)
	uuid, _ := strconv.ParseUint(idStr, 10, 64)
	err := h.r.validationService.ValidateHouse(serviceCtx, uuid)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		h.r.logger.Println(err)
		invalidDataResponce(apiCtx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}

	flats, err := h.r.houseService.HouseFlats(context.Background(), uuid)
	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}

	h.r.sendResponce(apiCtx, flats)

}
func (h *houseImpl) subscribe(ctx *fasthttp.RequestCtx) {
	
}
