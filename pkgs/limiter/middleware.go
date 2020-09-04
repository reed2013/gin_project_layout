package limiter

func RateLimiter(limit LimiterInterface) gin.HandleFunc {
	return func(c *gin.Context) {
		key := limit.Key(c)
		if bucket, ok := limit.GetBucket(key) {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequest)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
