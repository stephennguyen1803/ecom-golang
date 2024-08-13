package benchmark

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

var ctxSet = context.Background()

func BenchmarkPoolSize1(b *testing.B) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Protocol: 3,
		PoolSize: 1,
	})
	defer rdb.Close()
	b.ResetTimer()
	rdb.FlushAll(ctxSet)
	for i := 0; i < b.N; i++ {
		rdb.Set(ctxSet, "key", "value", 0)
	}
}

func BenchmarkPoolSize10(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Protocol: 3,
		PoolSize: 10,
	})
	defer rdb.Close()
	b.ResetTimer()
	rdb.FlushAll(ctxSet)
	for i := 0; i < b.N; i++ {
		rdb.Set(ctxSet, "key", "value", 0)
	}
}
