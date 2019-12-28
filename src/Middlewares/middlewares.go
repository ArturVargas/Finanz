package middlewares

import (
	"net/http"

	utils "../Utils"
	"github.com/gin-gonic/gin"
)

//SetMiddlewareToken
func SetMiddleWareToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Next()
	}
}
