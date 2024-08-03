package router

import "github.com/valyala/fasthttp"

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
	role := ctx.QueryArgs().Peek("role")
	a.r.logger.Println(string(role))
}

func (a *authImpl) login(ctx *fasthttp.RequestCtx) {

}
func (a *authImpl) register(ctx *fasthttp.RequestCtx) {

}
