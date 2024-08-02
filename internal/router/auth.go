package router

import "github.com/valyala/fasthttp"

type authImpl struct {
	r *Router
}

func registerAuthApi(r *Router) {
	authImpl := authImpl{r: r}
	r.router.GET("/dummyLogin", authImpl.dummyLogin)
	r.router.GET("/login", authImpl.login)
	r.router.GET("/register", authImpl.register)

}

func (a *authImpl) dummyLogin(ctx *fasthttp.RequestCtx) {

}

func (a *authImpl) login(ctx *fasthttp.RequestCtx) {

}
func (a *authImpl) register(ctx *fasthttp.RequestCtx) {

}
