package router

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"
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

func (f *flatImpl) createFlat(apiCtx *fasthttp.RequestCtx) {
	serviceCtx := context.WithValue(context.Background(), models.Role{}, apiCtx.UserValue("role"))
	data := apiCtx.Request.Body()
	flatBuilder, err := f.r.validationService.ValidateFlatBuilderData(serviceCtx, data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		f.r.logger.Println(err)
		invalidDataResponce(apiCtx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}

	flat, err := f.r.houseService.CreateFlat(serviceCtx, flatBuilder)

	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}
	f.r.sendResponce(apiCtx, flat)
}

func (f *flatImpl) changeModerationType(apiCtx *fasthttp.RequestCtx) {
	serviceCtx := context.WithValue(context.Background(), models.Role{}, apiCtx.UserValue("role"))
	data := apiCtx.Request.Body()
	status, err := f.r.validationService.ValidateFlatStatusData(serviceCtx, data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		f.r.logger.Println(err)
		invalidDataResponce(apiCtx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}

	flat, err := f.r.houseService.UpdateFlatStatus(serviceCtx, status)
	if errors.As(err, &serviceErrors.ServerError{}) {
		f.r.logger.Println(err)
		internalServerErrorResponce(apiCtx)
		return
	}

	f.r.sendResponce(apiCtx, flat)
}
