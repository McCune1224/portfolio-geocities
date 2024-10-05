package georedis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type GeoDB struct {
	client *redis.Client
}

func NewClientMust(redisURL string) *GeoDB {
	redisOpt, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}
	return &GeoDB{redis.NewClient(redisOpt)}

}

func (gr *GeoDB) Ping(ctx context.Context) error {
	return gr.client.Ping(ctx).Err()
}

func (gr *GeoDB) GetSiteViewCount(ctx context.Context, key string) (int64, error) {
	return gr.client.Incr(ctx, key).Result()
}

func (gr *GeoDB) GetRedisClient() *redis.Client {
	return gr.client
}
