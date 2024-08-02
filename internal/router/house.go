package router

import "github.com/valyala/fasthttp"

type houseImpl struct {
	r *Router
}

func registerHouseApi(r *Router) {
	houseImpl := houseImpl{r}
	r.router.POST("/house/create", houseImpl.createHouse)
	r.router.GET("/house/{id}", houseImpl.getHouseData)
	r.router.POST("/house/{id}/subscribe", houseImpl.subscrie)
}

func (h *houseImpl) createHouse(ctx *fasthttp.RequestCtx) {

}
func (h *houseImpl) getHouseData(ctx *fasthttp.RequestCtx) {

}
func (h *houseImpl) subscrie(ctx *fasthttp.RequestCtx) {

}
