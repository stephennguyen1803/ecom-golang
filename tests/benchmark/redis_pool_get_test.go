package benchmark

import (
	"context"
	"strconv"
	"testing"

	"github.com/redis/go-redis/v9"
)

var ctxGet = context.Background()

func initDataRedis(b *testing.B, ctx context.Context, rdb *redis.Client) {
	rdb.FlushAll(ctx)
	var key, value string
	for i := 0; i < b.N; i++ {
		key = "key" + strconv.Itoa(i)
		value = "value" + strconv.Itoa(i)
		rdb.Set(ctx, key, value, 0)
	}
}

func BenchmarkPoolSizeGet1(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Protocol: 3,
		PoolSize: 1,
	})
	defer rdb.Close()
	b.ResetTimer()
	initDataRedis(b, ctxGet, rdb)
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			rdb.Get(ctxGet, "key"+strconv.Itoa(i))
		}
	})
}

func BenchmarkPoolSizeGet15(b *testing.B) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Protocol: 3,
		PoolSize: 15,
	})
	defer rdb.Close()
	b.ResetTimer()
	initDataRedis(b, ctxGet, rdb)
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			rdb.Get(ctxGet, "key"+strconv.Itoa(i))
		}
	})
}
