package limiter

import "time"

type LimiterInterface interface {
	Key (c *gin.Context) string

	GetBucket (key string) (*ratelimit.Bucket, book)

	AddBuckets(rules ...LimiterBucketRule) LimiterInterface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key string
	FillInterval time.Duration
	Capacity int64
	Quantum int64
}
