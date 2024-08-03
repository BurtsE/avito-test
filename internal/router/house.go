package router

import (
	"avito-test/internal/converter"
	"encoding/json"
	"log"
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
	data := ctx.Request.Body()
	builder, err := converter.BuilderFromRawData(data)
	if err != nil {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	house, err := h.r.houseService.CreateHouse(builder)
	if err != nil {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	responce, err := json.Marshal(house)
	if err != nil {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	ctx.Response.AppendBody(responce)

}
func (h *houseImpl) getHouseData(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	uuid, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	log.Println(uuid)

	house, err := h.r.houseService.GetHouseDesc(uuid)
	if err != nil {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	responce, err := json.Marshal(house)
	if err != nil {
		h.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	ctx.Response.AppendBody(responce)

}
func (h *houseImpl) subscribe(ctx *fasthttp.RequestCtx) {

}
