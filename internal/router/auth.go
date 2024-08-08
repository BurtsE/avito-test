package router

import (
	serviceErrors "avito-test/internal/service_errors"
	"context"
	"errors"

	"github.com/valyala/fasthttp"
)

type authImpl struct {
	r *Router
}

func registerAuthApi(r *Router) {
	authImpl := authImpl{r: r}
	r.router.GET("/dummyLogin", authImpl.dummyLogin)
	r.router.POST("/login", authImpl.login)
	r.router.POST("/register", authImpl.register)

}

func (a *authImpl) dummyLogin(apiCtx *fasthttp.RequestCtx) {
	data := apiCtx.QueryArgs().Peek("role")
	serviceCtx := context.Background()
	role, err := a.r.validationService.ValidateDummyUserData(serviceCtx, data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		a.r.logger.Println(err)
		invalidDataResponce(apiCtx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		a.r.logger.Println(err)
		apiCtx.SetUserValue("errorMessage", "неправильный формат json")
		internalServerErrorResponce(apiCtx)
		return
	}

	responce, err := a.r.authService.DummyAuthorize(serviceCtx, role)
	if errors.As(err, &serviceErrors.ServerError{}) {
		a.r.logger.Println(err)
		apiCtx.SetUserValue("errorMessage", "ошибка авторизации")
		internalServerErrorResponce(apiCtx)
		return
	}

	a.r.sendResponce(apiCtx, responce)
}

func (a *authImpl) login(ctx *fasthttp.RequestCtx) {

}
func (a *authImpl) register(ctx *fasthttp.RequestCtx) {

}
