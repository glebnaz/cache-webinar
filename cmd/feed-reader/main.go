package main

import (
	"github.com/glebnaz/cache-webinar/internal/app"
	"github.com/glebnaz/cache-webinar/internal/cache"
	"github.com/glebnaz/cache-webinar/internal/router"
)

var server app.App

func main() {

	cache, err := cache.NewRedisCache("localhost:6379")
	if err != nil {
		panic(err)
	}
	server = app.NewApp(cache)

	ech := router.InitRouters(Router)

	ech.Start(":8888")
}
