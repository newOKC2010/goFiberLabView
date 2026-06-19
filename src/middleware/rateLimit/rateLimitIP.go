package ratelimit

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// RateLimitByIP - rate limit ตาม IP address
func RateLimitByIP(maxRequests int, duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		clientIP := getRealIP(c)
		now := time.Now()

		limitsLock.Lock()
		ipLimit, exists := ipLimits[clientIP]
		if !exists {
			ipLimit = &UserRateLimit{count: 0, resetTime: now.Add(duration)}
			ipLimits[clientIP] = ipLimit
		}
		limitsLock.Unlock()

		ipLimit.mu.Lock()
		defer ipLimit.mu.Unlock()

		if now.After(ipLimit.resetTime) {
			ipLimit.count = 0
			ipLimit.resetTime = now.Add(duration)
		}

		if ipLimit.count >= maxRequests {
			log.Printf("Rate limit exceeded: IP=%s, Path=%s, Limit=%d/%v", clientIP, c.Path(), maxRequests, duration)
			return c.Status(429).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("เกินขีดจำกัด %d ครั้งต่อ %v", maxRequests, duration),
			})
		}

		ipLimit.count++
		return c.Next()
	}
}

// getRealIP - ดึง real IP จาก proxy headers
func getRealIP(c *fiber.Ctx) string {
	// ลำดับความสำคัญ headers
	headers := []string{
		"X-Forwarded-For",
		"X-Real-IP",
		"X-Client-IP",
		"CF-Connecting-IP", // Cloudflare
	}

	for _, header := range headers {
		if ip := c.Get(header); ip != "" {
			// X-Forwarded-For อาจมีหลาย IP คั่นด้วย comma
			if header == "X-Forwarded-For" {
				if firstIP := strings.Split(ip, ",")[0]; firstIP != "" {
					return strings.TrimSpace(firstIP)
				}
			}
			return ip
		}
	}

	// ถ้าไม่มี header ให้ใช้ IP default
	return c.IP()
}
