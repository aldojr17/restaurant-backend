package middleware

import (
	"log"
	"math"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func LatencyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		responseTime := time.Now()
		latency := responseTime.Sub(startTime).Seconds() * 1000
		latency = math.Floor(latency*100) / 100

		pathName := strings.Replace(ctx.FullPath(), "/", "_", -1)
		pathName = strings.Replace(pathName, ":", "_", -1)
		log.Printf("%s_%s_api_latency: %v\n", pathName, ctx.Request.Method, latency)
	}
}
