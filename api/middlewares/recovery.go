package middlewares

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func (m *MiddleWare) Recovery(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				fmt.Println(string(httpRequest))

				logger.Error().
					Str("message", "[Recovery from panic]").
					Time("time", time.Now()).
					Interface("error", err).
					Str("request", string(httpRequest)).
					Str("stack", string(debug.Stack())).
					Msg("")
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
