package router

import (
	serviceErrors "avito-test/internal/service_errors"
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

func (a *authImpl) dummyLogin(ctx *fasthttp.RequestCtx) {
	data := ctx.QueryArgs().Peek("role")
	role, err := a.r.validationService.ValidateDummyUserData(data)
	if errors.As(err, &serviceErrors.ValidationError{}) {
		a.r.logger.Println(err)
		invalidDataResponce(ctx)
		return
	}
	if errors.As(err, &serviceErrors.ServerError{}) {
		a.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}

	responce, err := a.r.authService.DummyAuthorize(role)
	if errors.As(err, &serviceErrors.ServerError{}) {
		a.r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	a.r.sendResponce(ctx, responce)
}

func (a *authImpl) login(ctx *fasthttp.RequestCtx) {

}
func (a *authImpl) register(ctx *fasthttp.RequestCtx) {

}
