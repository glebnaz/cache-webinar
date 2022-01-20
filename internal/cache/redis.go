package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/glebnaz/cache-webinar/internal/model"
	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	cli *redis.Client
}

func (r *RedisCache) WriteToSubs(ctx context.Context, post model.Post, subs []string) error {
	for i := range subs {
		subFeed, err := r.cli.Get(ctx, subs[i]).Result()
		if err != nil {
			if fmt.Sprintf("%s", err) != "redis: nil" {
				return err
			}
		}

		var posts []model.Post

		if len(subFeed) != 0 {
			err = json.Unmarshal([]byte(subFeed), &posts)
			if err != nil {
				return err
			}
		}

		posts = append(posts, post)

		res, err := json.Marshal(posts)
		if err != nil {
			return err
		}

		err = r.cli.Set(ctx, subs[i], res, 0).Err()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RedisCache) ReadFeed(ctx context.Context, id string) ([]model.Post, error) {
	res, err := r.cli.Get(ctx, id).Result()
	if err != nil {
		return nil, err
	}

	var posts []model.Post

	err = json.Unmarshal([]byte(res), &posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func NewRedisCache(redisHost string) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	return &RedisCache{cli: rdb}, err
}
