package main

import (
	"github.com/glebnaz/cache-webinar/internal/app"
	"github.com/glebnaz/cache-webinar/internal/router"
)

var server app.App

func main() {
	server = app.NewApp(nil)

	ech := router.InitRouters(Router)

	ech.Start(":8888")
}
