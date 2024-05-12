package utils

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yazmeyaa/todo_app_backend/data/response"
)

func GetIntQuery(ctx *gin.Context, query string, defaultValue *int) (int, error) {
	q, exist := ctx.GetQuery(query)
	if !exist && defaultValue != nil {
		return *defaultValue, nil
	}
	parsed, fail := strconv.Atoi(q)

	if fail != nil {
		ctx.JSON(400, response.ApiErrorResponse{
			Error: fmt.Sprintf("Unexpected value in query param \"%s\".", query),
		})

		return parsed, fail
	}

	return parsed, nil
}
