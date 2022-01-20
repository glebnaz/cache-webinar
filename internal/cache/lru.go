package cache

import (
	"context"
	"errors"
	"github.com/glebnaz/cache-webinar/internal/model"
	lru "github.com/hashicorp/golang-lru"
	"sync"
)

type LRU struct {
	m    sync.RWMutex
	data *lru.Cache
}

func NewLRU() (*LRU, error) {
	data, err := lru.New(100)
	if err != nil {
		return nil, err
	}
	return &LRU{data: data}, nil
}

func (l *LRU) WriteToSubs(ctx context.Context, post model.Post, subs []string) error {
	l.m.Lock()
	defer l.m.Unlock()

	for i := range subs {
		toWrite := []model.Post{}
		val, ok := l.data.Get(subs[i])
		if !ok {
			toWrite = []model.Post{post}
		} else {
			toWrite, ok = val.([]model.Post)
			if !ok {
				return errors.New("cache is bad type")
			}
			if toWrite != nil {
				toWrite = []model.Post{post}
			}
			toWrite = append(toWrite, post)
		}
		l.data.Add(subs[i], toWrite)
	}

	return nil
}

func (l *LRU) ReadFeed(ctx context.Context, id string) ([]model.Post, error) {
	l.m.Lock()
	defer l.m.Unlock()

	val, ok := l.data.Get(id)
	if !ok {
		return nil, nil
	}

	feed, ok := val.([]model.Post)
	if !ok {
		return nil, errors.New("cache is bad type")
	}
	return feed, nil
}
