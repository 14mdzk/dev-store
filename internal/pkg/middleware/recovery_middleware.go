package middleware

import (
	"net/http"

	"github.com/14mdzk/dev-store/internal/pkg/reason"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err != nil {
				log.Error(err)
				ctx.JSON(http.StatusInternalServerError, gin.H{"message": reason.InternalServerError})
			}
		}()

		ctx.Next()
	}
}
