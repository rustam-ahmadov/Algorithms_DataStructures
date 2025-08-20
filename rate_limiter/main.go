package main

import (
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// 2 types of clients
// Prime / Others
// Prime 	: 100 rpm
// Others	:   1 rpm

type clientType int

const (
	regular clientType = iota
	premium
)
const (
	regularRMP = 1
	premiumRpm = 100
)

type requestPerMin struct {
	time  time.Time
	count int
	limit int
}

type Client struct {
	id int
	clientType
}

func getClient(token string) Client {
	return Client{
		id: rand.Int(),
	}
}

func getLimit(c Client) int {
	if c.clientType == premium {
		return premiumRpm
	}

	return regularRMP
}

// FixedWindow limiter

type RateLimiter struct {
	rw sync.RWMutex
	m  map[int]requestPerMin
}

func (r *RateLimiter) Allow(client Client) bool {
	// first time

	id := client.id

	r.rw.RLock()
	reqPerMin, ok := r.m[id]
	r.rw.Unlock()

	if !ok || time.Since(reqPerMin.time) > time.Minute {
		rpm := requestPerMin{
			time:  time.Now(),
			count: 1, // first request
			limit: getLimit(client),
		}

		r.rw.Lock()
		r.m[id] = rpm
		r.rw.Unlock()

		return true
	}

	switch client.clientType {
	case premium:
		if reqPerMin.count < premiumRpm {
			reqPerMin.count++
		}
	case regular:
		return false
	}

	return true
}

func respErr(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusTooManyRequests)
	_, _ = w.Write([]byte(msg))
}

type TokenBucketLimiter struct {
	fillRateSec int
	rwm         sync.RWMutex
	m           map[int]UserBucket
}

func NewTokenBucketLimiter() *TokenBucketLimiter {
	return &TokenBucketLimiter{
		fillRateSec: 1, //perServ
		rwm:         sync.RWMutex{},
		m:           make(map[int]UserBucket),
	}
}

type UserBucket struct {
	limit      int
	tokens     int
	lastRefill time.Time
}

func (limiter *TokenBucketLimiter) Allow(client Client) bool {
	limiter.rwm.RLock()
	bucket, ok := limiter.m[client.id]
	limiter.rwm.RUnlock()

	if !ok {
		limiter.rwm.Lock()
		limiter.m[client.id] = UserBucket{
			limit:      getLimit(client),
			tokens:     getLimit(client),
			lastRefill: time.Now(),
		}

		limiter.rwm.Unlock()
		return true
	}

	elapsed := time.Now().Sub(bucket.lastRefill)
	newTokens := min(int(elapsed.Seconds()) / limiter.fillRateSec)

	bucket.tokens += newTokens
	if bucket.tokens > 0 {
		bucket.tokens--

		limiter.rwm.Lock()
		limiter.m[client.id] = bucket
		limiter.rwm.Unlock()

		return true
	}

	return false
}

func rateLimiterMiddleware(r *RateLimiter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authentication")

		client := getClient(token)
		if !r.Allow(client) {
			respErr(w, "too many requests")
		}

	})
}

func main() {

}
