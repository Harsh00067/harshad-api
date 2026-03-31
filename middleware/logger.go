package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {

	return func(c *gin.Context) {

		requestID := time.Now().UnixNano()
		start := time.Now()

		fmt.Printf("[START] ID:%d %s %s\n",
			requestID,
			c.Request.Method,
			c.Request.URL.Path,
		)

		c.Next()

		duration := time.Since(start)

		fmt.Printf("[END] ID:%d STATUS:%d TIME:%v\n",
			requestID,
			c.Writer.Status(),
			duration,
		)
	}
}
