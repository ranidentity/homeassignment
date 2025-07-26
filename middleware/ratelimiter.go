package middleware

import (
	"net/http"
	"sync"
)

type IPRateLimiter struct {
	ips           map[string]int
	mu            sync.Mutex
	maxConcurrent int
}

func NewIPRateLimiter(maxConcurrent int) *IPRateLimiter {
	return &IPRateLimiter{
		ips:           make(map[string]int),
		maxConcurrent: maxConcurrent,
	}
}

func (rl *IPRateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // For simplicity

		rl.mu.Lock()
		defer rl.mu.Unlock()

		if rl.ips[ip] >= rl.maxConcurrent {
			http.Error(w, "Too many concurrent requests", http.StatusTooManyRequests)
			return
		}

		rl.ips[ip]++
		defer func() {
			rl.mu.Lock()
			rl.ips[ip]--
			rl.mu.Unlock()
		}()

		next.ServeHTTP(w, r)
	})
}
