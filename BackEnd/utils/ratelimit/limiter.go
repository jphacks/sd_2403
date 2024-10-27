package ratelimit

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

const (
	RATE_PER_SECOND = 60
	BURST           = 180
	CLEAN_UP_PERIOD = 1 * time.Minute
)

// Create a custom visitor struct which holds the rate limiter for each
// visitor and the last time that the visitor was seen.
type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Change the the map to hold values of the type visitor.
var visitors = make(map[string]*visitor)
var mu sync.RWMutex

// Run a background goroutine to remove old entries from the visitors map.
func init() {
	go cleanupVisitors()
}

func getVisitor(ip string) *rate.Limiter {
	mu.RLock()

	v, exists := visitors[ip]
	if !exists {
		mu.RUnlock()
		mu.Lock()
		limiter := rate.NewLimiter(RATE_PER_SECOND, BURST)
		// Include the current time when creating a new visitor.
		visitors[ip] = &visitor{limiter, time.Now()}
		mu.Unlock()
		return limiter
	}

	// Update the last seen time for the visitor.
	v.lastSeen = time.Now()
	mu.RUnlock()
	return v.limiter
}

// Every minute check the map for visitors that haven't been seen for
// more than 3 minutes and delete the entries.
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)

		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > CLEAN_UP_PERIOD {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

// RateLimit limit over the raw http handler independent to rest framework
func RateLimit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal server Error", http.StatusInternalServerError)
			return
		}

		limiter := getVisitor(ip)
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// GinMiddleware limit over gin router as middleware
func GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if len(ip) < 1 {
			c.AbortWithStatusJSON(500, gin.H{"message": "Internal server Error"})
			return
		}

		limiter := getVisitor(ip)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{"message": "Too Many Requests"})
			return
		}
		c.Next()
	}
}
