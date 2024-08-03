package router

import (
	"avito-test/internal/converter"

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
	data := ctx.Request.Body()
	flatBuilder, err := converter.FlatBuilderFromRawData(data)
	if err != nil {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	flat, err := f.r.houseService.CreateFlat(flatBuilder)
	if err != nil {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	f.r.sendResponce(ctx, flat)
}

func (f *flatImpl) changeModerationType(ctx *fasthttp.RequestCtx) {
	data := ctx.Request.Body()
	status, err := converter.FlatStatusFromRawData(data)
	if err != nil {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	flat, err := f.r.houseService.UpdateFlatStatus(status)
	if err != nil {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	f.r.sendResponce(ctx, flat)
}
