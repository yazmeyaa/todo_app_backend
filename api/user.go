package api

import "github.com/gin-gonic/gin"

type UserController interface {
	Create(*gin.Context)
	Delete(*gin.Context)
	Find(*gin.Context)
}
