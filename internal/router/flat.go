package router

import (
	serviceErrors "avito-test/internal/service_errors"
	"errors"

	"github.com/valyala/fasthttp"
)

type flatImpl struct {
	r *Router
}

func registerFlatApi(r *Router) {
	flatImpl := flatImpl{r}
	r.router.POST("/flat/create", r.UserAccess(flatImpl.createFlat))
	r.router.POST("/flat/update", r.ModeratorAccess(flatImpl.changeModerationType))
}

func (f *flatImpl) createFlat(ctx *fasthttp.RequestCtx) {
	data := ctx.Request.Body()
	flatBuilder, err := f.r.validationService.ValidateFlatBuilderData(data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		f.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	flat, err := f.r.houseService.CreateFlat(flatBuilder)

	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	if errors.As(err, &serviceErrors.AuthError{}) {
		f.r.logger.Println(err)
		unAuthorized(ctx)
		return
	}

	f.r.sendResponce(ctx, flat)
}

func (f *flatImpl) changeModerationType(ctx *fasthttp.RequestCtx) {
	data := ctx.Request.Body()
	status, err := f.r.validationService.ValidateFlatStatusData(data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		f.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	flat, err := f.r.houseService.UpdateFlatStatus(status)
	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	if errors.As(err, &serviceErrors.AuthError{}) {
		f.r.logger.Println(err)
		unAuthorized(ctx)
		return
	}

	f.r.sendResponce(ctx, flat)
}
