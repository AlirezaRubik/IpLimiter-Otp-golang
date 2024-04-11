package middleware

import (
	"errors"
	"information/api/helper"
	"information/pkg/limiter"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter(/*cfg *config.Config*/) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(/*In Here We Use The Otp Limiter Time */120*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseReseponseWithError(nil, false, 400, errors.New("Not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
