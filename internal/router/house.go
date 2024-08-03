package router

import (
	"avito-test/internal/converter"
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
	r.router.POST("/house/create", houseImpl.createHouse)
	r.router.GET("/house/{id}", houseImpl.getHouseData)
	r.router.POST("/house/{id}/subscribe", houseImpl.subscribe)
}

func (h *houseImpl) createHouse(ctx *fasthttp.RequestCtx) {
	var (

	)
	data := ctx.Request.Body()
	builder, err := converter.HouseBuilderFromRawData(data)
	if err != nil {
		h.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}
	
	house, err := h.r.houseService.CreateHouse(builder)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		h.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}

	if errors.As(err, &serviceErrors.DatabaseError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	if errors.As(err, &serviceErrors.AuthError{}) {
		h.r.logger.Println(err)
		unAuthorized(ctx)
		return
	}

	h.r.sendResponce(ctx, house)
}
func (h *houseImpl) getHouseData(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	uuid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}

	house, err := h.r.houseService.GetHouseDesc(uuid)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		h.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}

	if errors.As(err, &serviceErrors.DatabaseError{}) {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	if errors.As(err, &serviceErrors.AuthError{}) {
		h.r.logger.Println(err)
		unAuthorized(ctx)
		return
	}
	
	h.r.sendResponce(ctx, house)

}
func (h *houseImpl) subscribe(ctx *fasthttp.RequestCtx) {

}
