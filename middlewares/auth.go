package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yazmeyaa/todo_app_backend/data/response"
	"github.com/yazmeyaa/todo_app_backend/services"
)

func AuthJWTMiddleware(jwtService services.JWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) != 2 {
			webResponse := response.NewApiErrorResponse("Missed auth token")
			ctx.JSON(http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		authToken := t[1]
		valid := jwtService.Verify(authToken)

		if !valid {
			webResponse := response.NewApiErrorResponse("Not valid token")
			ctx.JSON(http.StatusUnauthorized, webResponse)
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
