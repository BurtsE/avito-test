package router

import (
	"fmt"

	"github.com/valyala/fasthttp"
)


// TODO 500 code fix
func internalServerErrorResponce(ctx *fasthttp.RequestCtx) []byte {
	ctx.SetStatusCode(500)
	return fmt.Appendf([]byte{}, `{
		"message": %s,
		"request_id": %s,
		"code": %d
	}`, "что-то пошло не так", "g12ugs67gqw67yu12fgeuqwd", 12345)
}
