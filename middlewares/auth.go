package middlewares

import (
	"ApiRest/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	err := utils.VerifyJwt(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	} else {
		c.Next()
	}
}
