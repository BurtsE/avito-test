package router

import (
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

// TODO 500 code fix
func internalServerErrorResponce(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(500)
	ctx.Response.AppendBody(fmt.Appendf([]byte{}, `{
		"message": %s,
		"request_id": %d,
		"code": %d
	}`, "что-то пошло не так", ctx.ID(), 500))
	ctx.Response.Header.Add("Retry-After", time.Minute.String())
}

func invalidDataResponce(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(400)
}

func unAuthorized(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(401)
}
