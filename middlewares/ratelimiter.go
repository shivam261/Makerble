package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github/shivam261/ClinicManagement/initializers"

	"github.com/gin-gonic/gin"
)

func RateLimiter(limit int, duration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		key := fmt.Sprintf("rl:%s", clientIP)

		// Increment the count
		count, err := initializers.RedisClient.Incr(initializers.Ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter failed"})
			return
		}

		// Set expiry only when it's first request
		if count == 1 {
			initializers.RedisClient.Expire(initializers.Ctx, key, duration)
		}

		if count > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Please try again later.",
			})
			return
		}

		c.Next()
	}
}
