package limiter

type MethodLimiter struct {
	*Limiter
}

func newMethodLimiter() LimiterInterface {
	return MethodLimiter{
		Limiter: &Limiter{limiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}

func (limiter MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

func (limiter MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := limiter.limiterBuckets[key]
	return bucket, ok
}

func (limiter MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
	for _, rule := range rules {
		if _, ok := limiter.limiterBuckets[rule.Key]; !ok {
			limiter.limiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}

	return limiter
}
