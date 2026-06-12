package ratelimit

import (
	"sync"
	"time"
)

type UserRateLimit struct {
	count     int
	resetTime time.Time
	mu        sync.Mutex
}
