package api

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	CheckToken(ctx *gin.Context)
}
