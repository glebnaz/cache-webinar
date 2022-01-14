package router

import "github.com/labstack/echo/v4"

const AnyMethods = "ANY"

//Router route one endpoint
type Router struct {
	Path       string
	Method     string
	Handler    echo.HandlerFunc
	Middleware []echo.MiddlewareFunc
}

//InitRouters init router from array Router
func InitRouters(routers []Router) *echo.Echo {
	server := echo.New()
	for _, v := range routers {
		if v.Method == AnyMethods {
			server.Any(v.Path, v.Handler, v.Middleware...)
		} else {
			server.Add(v.Method, v.Path, v.Handler, v.Middleware...)
		}
	}
	return server
}
