package router

import (
	serviceErrors "avito-test/internal/service_errors"
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

func (h *houseImpl) createHouse(ctx *fasthttp.RequestCtx) {
	data := ctx.Request.Body()
	builder, err := h.r.validationService.ValidateHouseData(data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		h.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	house, err := h.r.houseService.CreateHouse(builder)

	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	h.r.sendResponce(ctx, house)
}
func (h *houseImpl) getHouseData(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	uuid, _ := strconv.ParseUint(idStr, 10, 64)

	err := h.r.validationService.ValidateHouse(uuid)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		h.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	house, err := h.r.houseService.GetHouseDesc(uuid)
	if errors.As(err, &serviceErrors.ServerError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	h.r.sendResponce(ctx, house)

}
func (h *houseImpl) subscribe(ctx *fasthttp.RequestCtx) {

}
