package api

import "github.com/gin-gonic/gin"

type TaskController interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	GetList(ctx *gin.Context)
	GetById(ctx *gin.Context)
}
