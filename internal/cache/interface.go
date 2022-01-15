package cache

import (
	"context"
	"github.com/glebnaz/cache-webinar/internal/model"
)

//Cache interface
type Cache interface {
	WriteToSubs(ctx context.Context, post model.Post, subs []string) error
	ReadFeed(ctx context.Context, id string) ([]model.Post, error)
}
