package ratelimit

import (
	"fmt"
	"log"

	"time"

	"github.com/gofiber/fiber/v2"
)

// RateLimitByEmail - rate limit ตาม Email
func RateLimitByEmail(maxRequests int, duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody map[string]string
		if err := c.BodyParser(&requestBody); err != nil {
			return c.Status(400).JSON(fiber.Map{"success": false, "message": "ข้อมูลไม่ถูกต้อง"})
		}

		email := requestBody["email"]
		if email == "" {
			return c.Status(400).JSON(fiber.Map{"success": false, "message": "ไม่พบ Email"})
		}

		now := time.Now()

		limitsLock.Lock()
		emailLimit, exists := cidLimits[email]
		if !exists {
			emailLimit = &UserRateLimit{count: 0, resetTime: now.Add(duration)}
			cidLimits[email] = emailLimit
		}
		limitsLock.Unlock()

		emailLimit.mu.Lock()
		defer emailLimit.mu.Unlock()

		if now.After(emailLimit.resetTime) {
			emailLimit.count = 0
			emailLimit.resetTime = now.Add(duration)
		}

		if emailLimit.count >= maxRequests {
			log.Printf("Rate limit exceeded: Email=%s, Path=%s, Limit=%d/%v", email, c.Path(), maxRequests, duration)
			return c.Status(429).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("เกินขีดจำกัด %d ครั้งต่อ %v", maxRequests, duration),
			})
		}

		emailLimit.count++
		return c.Next()
	}
}
