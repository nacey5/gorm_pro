package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// 限流
type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

// @Param Key：自定义键值对名称。
// @Param FillInterval：间隔多久时间放 N 个令牌。
// @Param Capacity：令牌桶的容量。
// @Param Quantum：每次到达间隔时间后所放的具体令牌数量。
type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
