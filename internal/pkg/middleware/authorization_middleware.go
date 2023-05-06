package middleware

import (
	"net/http"

	"github.com/14mdzk/dev-store/internal/pkg/handler"
	"github.com/14mdzk/dev-store/internal/pkg/reason"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func AuthorizationMiddleware(sub string, obj string, act string, enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok, err := enforcer.Enforce(sub, obj, act)
		if err != nil || !ok {
			handler.ResponseError(ctx, http.StatusForbidden, reason.InvalidAccess)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
