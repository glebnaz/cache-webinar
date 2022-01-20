package cache

import (
	"context"
	"github.com/glebnaz/cache-webinar/internal/model"
	"sync"
)

type Simple struct {
	data map[string][]model.Post
	m    sync.RWMutex
}

func NewSimple() *Simple {
	data := make(map[string][]model.Post)
	return &Simple{data: data}
}

func (s *Simple) WriteToSubs(ctx context.Context, post model.Post, subs []string) error {
	s.m.Lock()
	defer s.m.Unlock()

	for i := range subs {
		if s.data[subs[i]] == nil {
			s.data[subs[i]] = []model.Post{}
		}
		s.data[subs[i]] = append(s.data[subs[i]], post)
	}

	return nil
}

func (s *Simple) ReadFeed(ctx context.Context, id string) ([]model.Post, error) {
	s.m.RLock()
	defer s.m.RUnlock()

	data := s.data[id]
	return data, nil
}
