package router

import (
	"github.com/valyala/fasthttp"
)

type flatImpl struct {
	r *Router
}

func registerFlatApi(r *Router) {
	flatImpl := flatImpl{r}
	r.router.POST("/flat/create", flatImpl.createFlat)
	r.router.POST("/flat/update", flatImpl.changeModerationType)
}

func (f *flatImpl) createFlat(ctx *fasthttp.RequestCtx) {

}

func (f *flatImpl) changeModerationType(ctx *fasthttp.RequestCtx) {


}
