package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/pemako/gva/server/global"
	"github.com/pemako/gva/server/model/common/response"
)

// 处理跨域请求,支持options访问
func NeedInit() gin.HandlerFunc {
	return func(c *gin.Context) {
		if global.GvaDb == nil {
			response.OkWithDetailed(gin.H{
				"needInit": true,
			}, "前往初始化数据库", c)
			c.Abort()
		} else {
			c.Next()
		}
		// 处理请求
	}
}
