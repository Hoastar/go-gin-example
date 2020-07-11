package costTime

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CostTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求前获取当前时间
		nowTime := time.Now()

		// 责任链 请求处理
		c.Next()

		// 处理后获取消耗时间
		costTime := time.Since(nowTime)
		url := c.Request.URL.String()
		fmt.Printf("%v %v\n", url, costTime)
	}
}
