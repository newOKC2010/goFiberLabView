package ratelimit

import (
	"fmt"
	"log"
	"sync"
	"time"

	middleware "view_lab/src/middleware"

	"github.com/gofiber/fiber/v2"
)

var usersLimits = make(map[string]*UserRateLimit)
var ipLimits = make(map[string]*UserRateLimit)
var cidLimits = make(map[string]*UserRateLimit)
var limitsLock sync.RWMutex

// RateLimitByUser - rate limit ตาม user ID
func RateLimitByUsers(maxRequests int, duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userER := c.Locals("user_view_lab").(*middleware.UserViewLabInfo)
		if userER == nil {
			return c.Status(401).JSON(fiber.Map{"success": false, "message": "ไม่พบข้อมูล user"})
		}

		key := fmt.Sprintf("%d:%s", userER.ID, c.Path())
		now := time.Now()

		limitsLock.Lock()
		usersLimit, exists := usersLimits[key]
		if !exists {
			usersLimit = &UserRateLimit{count: 0, resetTime: now.Add(duration)}
			usersLimits[key] = usersLimit
		}
		limitsLock.Unlock()

		usersLimit.mu.Lock()
		defer usersLimit.mu.Unlock()

		// Reset counter หมดเวลา
		if now.After(usersLimit.resetTime) {
			usersLimit.count = 0
			usersLimit.resetTime = now.Add(duration)
		}

		// ตรวจสอบ limit
		if usersLimit.count >= maxRequests {
			log.Printf("Rate limit exceeded: UserID=%d, Path=%s, Limit=%d/%v", userER.ID, c.Path(), maxRequests, duration)
			return c.Status(429).JSON(fiber.Map{
				"success": false,
				"message": fmt.Sprintf("เกินขีดจำกัด %d ครั้งต่อ %v", maxRequests, duration),
			})
		}

		usersLimit.count++
		return c.Next()
	}
}
