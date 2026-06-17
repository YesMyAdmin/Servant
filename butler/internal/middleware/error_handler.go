package middleware

import (
	"butler/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandlerFunc 将返回 error 的 controller 函数适配为 gin.HandlerFunc
func HandlerFunc(fn func(*gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := fn(c); err != nil {
			if svcErr, ok := err.(*pkg.ServiceError); ok {
				c.AbortWithStatusJSON(svcErr.HttpStatus(), svcErr.ErrorResp())
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				pkg.InternalServerError(err.Error()))
		}
	}
}
