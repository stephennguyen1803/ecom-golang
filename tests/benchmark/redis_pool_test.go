package benchmark

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var ctxSet = context.Background()

func setDataToRedis(b *testing.B, rdb *redis.Client) {
	var key, value string
	uuid := uuid.New().String()
	key = "key_" + uuid
	value = "value_" + uuid
	if err := rdb.Set(ctxSet, key, value, 0).Err(); err != nil {
		b.Fatal(err, "Failed to set data")
	}
}

func BenchmarkPoolSize1(b *testing.B) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Protocol: 3,
		PoolSize: 1,
	})
	defer rdb.Close()
	b.ResetTimer()
	rdb.FlushAll(ctxSet)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			setDataToRedis(b, rdb)
		}
	})
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
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			setDataToRedis(b, rdb)
		}
	})
}
