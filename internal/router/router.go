package router

import (
	"avito-test/internal/config"
	"avito-test/internal/service"
	"encoding/json"

	"github.com/fasthttp/router"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Router struct {
	router       *router.Router
	srv          *fasthttp.Server
	logger       *logrus.Logger
	houseService service.HouseService
	port         string
}

func NewRouter(logger *logrus.Logger, cfg *config.Config, houseService service.HouseService) *Router {
	router := router.New()
	srv := &fasthttp.Server{}
	r := &Router{
		router:       router,
		logger:       logger,
		houseService: houseService,
		port:         cfg.Service.Port,
		srv:          srv,
	}
	srv.Handler = r.loggerDecorator(router.Handler)

	registerHouseApi(r)
	registerFlatApi(r)
	registerAuthApi(r)
	r.router.GET("/status", statusHandler)
	return r
}

func (r *Router) Start() error {
	return r.srv.ListenAndServe(r.port)
}

func (r *Router) Shutdown() error {
	return r.srv.Shutdown()
}

func statusHandler(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
}

func (r *Router) loggerDecorator(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		handler(ctx)
		r.logger.Printf("api request: %s ;status code: %d", ctx.Path(), ctx.Response.StatusCode())
	}
}

func (r *Router) sendResponce(ctx *fasthttp.RequestCtx, a interface{}) {
	responce, err := json.Marshal(a)
	if err != nil {
		r.logger.Println(err)
		internalServerErrorResponce(ctx)
		return
	}
	ctx.Response.AppendBody(responce)
}
