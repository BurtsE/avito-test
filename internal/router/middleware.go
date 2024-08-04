package router

import (
	"avito-test/internal/models"
	serviceErrors "avito-test/internal/service_errors"
	"context"
	"errors"

	"github.com/valyala/fasthttp"
)

func (r *Router) UserAccess(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := ctx.Request.Header.Peek("Authorization")
		role, err := r.authService.CheckAuthorization(context.Background(), token)
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
	return func(apiCtx *fasthttp.RequestCtx) {
		token := apiCtx.Request.Header.Peek("Authorization")
		serviceCtx := context.Background()
		role, err := r.authService.CheckAuthorization(serviceCtx, token)
		if errors.As(err, &serviceErrors.ServerError{}) {
			r.logger.Println(err)
			internalServerErrorResponce(apiCtx)
			return
		}
		if errors.As(err, &serviceErrors.AuthError{}) {
			r.logger.Println(err)
			unAuthorized(apiCtx)
			return
		}
		if role != models.ModeratorRole {
			unAuthorized(apiCtx)
			return
		}
		apiCtx.SetUserValue("role", role)
		handler(apiCtx)
	}
}
