package app

import (
	"github.com/glebnaz/cache-webinar/internal/cache"
	"github.com/glebnaz/cache-webinar/internal/model"
	"github.com/labstack/echo/v4"
	"net/http"
)

type App struct {
	cache cache.Cache
}

type CreatePostRequest struct {
	post model.Post
}

type GetFeedRequest struct {
	id string
}

func (a *App) NewCreatePost() echo.HandlerFunc {
	subs := []string{"1"}
	return func(c echo.Context) error {
		var resp CreatePostRequest
		err := c.Bind(&resp)
		if err != nil {
			c.Error(err)
			return err
		}

		err = a.cache.WriteToSubs(c.Request().Context(), resp.post, subs)
		if err != nil {
			c.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, struct {
			Status string
			Subs   []string
		}{Status: "ok", Subs: subs})
	}
}

func (a *App) NewGetFeed() echo.HandlerFunc {
	return func(c echo.Context) error {
		var resp GetFeedRequest
		err := c.Bind(&resp)
		if err != nil {
			c.Error(err)
			return err
		}

		feed, err := a.cache.ReadFeed(c.Request().Context(), resp.id)
		if err != nil {
			c.Error(err)
			return err
		}

		return c.JSON(http.StatusOK, feed)
	}
}

func NewApp(cache cache.Cache) App {
	return App{cache: cache}
}
