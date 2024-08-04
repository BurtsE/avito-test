package router

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"errors"

	"github.com/valyala/fasthttp"
)

func (r *Router) UserAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := ctx.Request.Header.Peek("Authorization")
		role, err := r.authService.CheckAuthorization(token)
		if errors.As(err, &serviceErrors.ServerError{}) {
			r.logger.Println(err)
			internalServerErrorResponce(ctx)
			return
		}
		if errors.As(err, &serviceErrors.AuthError{}) {
			r.logger.Println(err)
			unAuthorized(ctx)
			return
		}
		if role != models.ModeratorRole && role != models.UserRole {
			unAuthorized(ctx)
			return
		}
		handler(ctx)
	}
}

func (r *Router) ModeratorAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := ctx.Request.Header.Peek("Authorization")
		role, err := r.authService.CheckAuthorization(token)
		if errors.As(err, &serviceErrors.ServerError{}) {
			r.logger.Println(err)
			internalServerErrorResponce(ctx)
			return
		}
		if errors.As(err, &serviceErrors.AuthError{}) {
			r.logger.Println(err)
			unAuthorized(ctx)
			return
		}
		if role != models.ModeratorRole {
			unAuthorized(ctx)
			return
		}
		handler(ctx)
	}
}
