package controller

import "github.com/gin-gonic/gin"

func ErrorHandle(handle func(c *gin.Context) error) gin.HandlerFunc {
	return func(context *gin.Context) {
		err := handle(context)
		if err != nil {
			context.JSON(400, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
		}
	}
}
