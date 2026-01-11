package middleware

import (
	"net/http"
	"sync"
	"time"

	"lpmaarifnu-site-api/internal/config"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	requests map[string]*clientRate
	mu       sync.RWMutex
	config   *config.RateLimitConfig
}

type clientRate struct {
	count     int
	lastReset time.Time
}

var limiter *rateLimiter

func RateLimit(cfg *config.RateLimitConfig) gin.HandlerFunc {
	if !cfg.Enabled {
		return func(c *gin.Context) {
			c.Next()
		}
	}

	limiter = &rateLimiter{
		requests: make(map[string]*clientRate),
		config:   cfg,
	}

	// Clean up old entries every minute
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			limiter.cleanup()
		}
	}()

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if !limiter.allow(clientIP) {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Rate limit exceeded. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (r *rateLimiter) allow(clientIP string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()

	if rate, exists := r.requests[clientIP]; exists {
		// Check if window has expired
		if now.Sub(rate.lastReset) > time.Duration(r.config.Window)*time.Second {
			rate.count = 1
			rate.lastReset = now
			return true
		}

		// Check if limit exceeded
		if rate.count >= r.config.Requests {
			return false
		}

		rate.count++
		return true
	}

	// New client
	r.requests[clientIP] = &clientRate{
		count:     1,
		lastReset: now,
	}
	return true
}

func (r *rateLimiter) cleanup() {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	for ip, rate := range r.requests {
		if now.Sub(rate.lastReset) > time.Duration(r.config.Window)*time.Second*2 {
			delete(r.requests, ip)
		}
	}
}
