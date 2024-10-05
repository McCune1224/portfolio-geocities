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

func (gr *GeoDB) GetSiteViewCount(ctx context.Context) (int64, error) {
	return gr.client.Incr(ctx, "site_view_count").Result()
}

func (gr *GeoDB) SetSiteViewCount(ctx context.Context, value int64) error {
	_, err := gr.client.Set(ctx, "site_view_count", value, 0).Result()
	return err
}

func (gr *GeoDB) IncrementSiteViewCount(ctx context.Context) (int64, error) {
	return gr.client.Incr(ctx, "site_view_count").Result()
}

func (gr *GeoDB) GetRedisClient() *redis.Client {
	return gr.client
}
