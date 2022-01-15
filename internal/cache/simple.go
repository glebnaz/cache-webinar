package cache

import (
	"context"
	"github.com/glebnaz/cache-webinar/internal/model"
	"sync"
)

type SimpleCache struct {
	m    sync.RWMutex
	data map[string][]model.Post
}

func (s *SimpleCache) WriteToSubs(ctx context.Context, post model.Post, subs []string) error {
	s.m.Lock()
	defer s.m.Unlock()

	for i := range subs {
		if len(s.data[subs[i]]) == 0 {
			s.data[subs[i]] = []model.Post{}
		}
		s.data[subs[i]] = append(s.data[subs[i]], post)
	}
	return nil
}

func (s *SimpleCache) ReadFeed(ctx context.Context, id string) ([]model.Post, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	posts := s.data[id]
	return posts, nil
}

func NewSimpleCache() *SimpleCache {
	return &SimpleCache{data: make(map[string][]model.Post)}
}
