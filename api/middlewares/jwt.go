package middlewares

import (
	"flick_tickets/common/log"
	"flick_tickets/core/usecase"

	"github.com/gin-gonic/gin"
)

type MiddleWare struct {
	jwtUseCase     *usecase.UseCaseJwt
	getUserUseCase *usecase.UseCaseUser
}

func NewMiddleware(
	jwtUseCase *usecase.UseCaseJwt,
	getUserUseCase *usecase.UseCaseUser,

) *MiddleWare {
	return &MiddleWare{
		jwtUseCase:     jwtUseCase,
		getUserUseCase: getUserUseCase,
	}
}

func (m *MiddleWare) Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		data, err := m.jwtUseCase.Decrypt(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}

		context.Set("username", data.UserName)
		log.Infof("user name", data.UserName)
		context.Next()
	}
}
