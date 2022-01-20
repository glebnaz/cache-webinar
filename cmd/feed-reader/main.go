package main

import (
	"github.com/glebnaz/cache-webinar/internal/app"
	"github.com/glebnaz/cache-webinar/internal/cache"
	"github.com/glebnaz/cache-webinar/internal/router"
)

var server app.App

func main() {

	simpleCache := cache.NewSimple()
	server = app.NewApp(simpleCache)

	ech := router.InitRouters(Router)

	ech.Start(":8888")
}
